package checker

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
)

func IsEnough(cfg config.Config, addr string) (bool, error) {
	countTx, err := GetCountTx(cfg, addr)
	if err != nil {
		return false, err
	}

	return countTx != 0, nil
}

func GetCountTx(cfg config.Config, addr string) (uint64, error) {
	logger := cfg.Log()
	pg := pg.NewMaterDB(cfg.DB())
	txFee, err := getTxFee(cfg, addr)
	if err != nil {
		logger.Errorf("Failed get fee: %v", err)
		return 0, err
	}

	votingInfo, err := pg.CheckerQ().GetVotingInfo(addr)
	if err != nil {
		return 0, err
	}

	voteBalance := votingInfo.Balance
	countTx := voteBalance / txFee
	return countTx, nil
}

func GetPredictCount(cfg config.Config, addr string, amount uint64) (uint64, error) {
	logger := cfg.Log()

	txFee, err := getTxFee(cfg, addr)
	if err != nil {
		logger.Errorf("Failed get fee: %v", err)
		return 0, err
	}
	return amount / txFee, nil
}

func getTxFee(cfg config.Config, addr string) (uint64, error) {
	logger := cfg.Log()
	networkParam := cfg.VotingV2Config()
	client := networkParam.RPC

	feeCap, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Warnf("Failed get gas price")
		return 0, err
	}

	votingInfo, err := pg.NewCheckerQ(cfg.DB()).GetVotingInfo(addr)
	if err != nil {
		return 0, err
	}

	totalFee := votingInfo.GasLimit * feeCap.Uint64()
	return totalFee, nil
}

func CheckUpdateGasLimit(value uint64, cfg config.Config, receiver *common.Address) error {
	pgDB := pg.NewMaterDB(cfg.DB())

	voteInfo, err := pgDB.CheckerQ().GetVotingInfo(receiver.Hex())
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

func AmountForCountTx(cfg config.Config, addr string, countTx uint64) (uint64, error) {
	logger := cfg.Log()

	txFee, err := getTxFee(cfg, addr)
	if err != nil {
		logger.Errorf("Failed get fee: %v", err)
		return 0, err
	}
	return countTx * txFee, nil
}
