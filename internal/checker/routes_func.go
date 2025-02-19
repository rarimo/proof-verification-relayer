package checker

import (
	"context"
	"math/big"

	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"gitlab.com/distributed_lab/logan/v3"
)

func IsEnough(cfg config.Config, votingId int64) (bool, error) {
	countTx, err := GetCountTx(cfg, votingId)
	if err != nil {
		return false, err
	}

	return countTx != 0, nil
}

func GetCountTx(cfg config.Config, votingId int64) (uint64, error) {
	logger := cfg.Log()
	pg := pg.NewMaterDB(cfg.DB())
	txFee, err := getTxFee(cfg, votingId)
	if err != nil {
		logger.Errorf("Failed get fee: %v", err)
		return 0, err
	}

	votingInfo, err := pg.CheckerQ().GetVotingInfo(votingId)
	if err != nil {
		return 0, err
	}
	if txFee == 0 {
		txFee = cfg.VotingV2Config().GasLimit
	}
	voteBalance := votingInfo.Balance
	countTx := voteBalance / txFee
	return countTx, nil
}

func GetPredictCount(cfg config.Config, votingId int64, amount uint64) (uint64, error) {
	logger := cfg.Log()

	txFee, err := getTxFee(cfg, votingId)
	if err != nil {
		logger.Errorf("Failed get fee: %v", err)
		return 0, err
	}
	if txFee == 0 {
		txFee = cfg.VotingV2Config().GasLimit
	}

	return amount / txFee, nil
}

func getTxFee(cfg config.Config, votingId int64) (uint64, error) {
	logger := cfg.Log()
	networkParam := cfg.VotingV2Config()
	client := networkParam.RPC

	feeCap, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Warnf("Failed get gas price")
		return 0, err
	}

	votingInfo, err := pg.NewCheckerQ(cfg.DB()).GetVotingInfo(votingId)
	if err != nil {
		return 0, err
	}
	if feeCap.Uint64() == 0 {
		feeCap = big.NewInt(1)
	}
	totalFee := votingInfo.GasLimit * feeCap.Uint64()
	return totalFee, nil
}

func CheckUpdateGasLimit(value uint64, cfg config.Config, votingId int64) error {
	pgDB := pg.NewMaterDB(cfg.DB())

	voteInfo, err := pgDB.CheckerQ().GetVotingInfo(votingId)
	if err != nil {
		return err
	}

	voteInfo.GasLimit = value
	err = pgDB.CheckerQ().UpdateVotingInfo(voteInfo)
	if err != nil {
		return err
	}
	return nil
}

func AmountForCountTx(cfg config.Config, votingId int64, countTx uint64) (uint64, error) {
	logger := cfg.Log()

	txFee, err := getTxFee(cfg, votingId)
	if err != nil {
		logger.Errorf("Failed get fee: %v", err)
		return 0, err
	}
	if txFee == 0 {
		txFee = cfg.VotingV2Config().GasLimit
	}
	return countTx * txFee, nil
}

func UpdateVotingBalance(cfg config.Config, gasPrice uint64, gas uint64, votingId int64) error {
	pgDB := pg.NewMaterDB(cfg.DB()).CheckerQ()

	voteInfo, err := pgDB.GetVotingInfo(votingId)
	if err != nil {
		cfg.Log().WithFields(logan.F{
			"Error":     err,
			"voting_id": votingId,
		}).Errorf("failed get voting info from db")
		return err
	}
	if gasPrice == 0 {
		gasPrice = 1
	}
	voteInfo.Balance = voteInfo.Balance - (gasPrice * gas)
	err = pgDB.UpdateVotingInfo(voteInfo)
	if err != nil {
		cfg.Log().WithFields(logan.F{
			"Error":     err,
			"voting_id": votingId,
		}).Errorf("failed update voting info from db")
		return err
	}
	return err
}
