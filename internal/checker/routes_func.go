package checker

import (
	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func IsEnough(cfg config.Config, votingId int64) (bool, error) {
	countTx, err := GetCountTx(cfg, votingId)
	if err != nil {
		return false, err
	}
	return countTx != 0, nil
}

func GetCountTx(cfg config.Config, votingId int64) (uint64, error) {
	pg := pg.NewMaterDB(cfg.DB())
	txFee, err := getTxFee(cfg, votingId)
	if err != nil {
		return 0, errors.Wrap(err, "failed get fee")
	}

	votingInfo, err := pg.CheckerQ().GetVotingInfo(votingId)
	if err != nil {
		return 0, err
	}
	if votingInfo == nil {
		return 0, sql.ErrNoRows
	}
	if txFee.Int64() == 0 {
		txFee = big.NewInt(int64(cfg.VotingV2Config().GasLimit))
	}
	countTx := new(big.Int).Div(votingInfo.Balance, txFee)
	return countTx.Uint64(), nil
}

func GetCountTxForAmount(cfg config.Config, votingId int64, amount *big.Int) (*big.Int, error) {
	txFee, err := getTxFee(cfg, votingId)
	if err != nil {
		return nil, errors.Wrap(err, "failed get fee")
	}
	if txFee.Int64() == 0 {
		txFee = big.NewInt(int64(cfg.VotingV2Config().GasLimit))
	}
	return new(big.Int).Div(amount, txFee), nil
}

func getTxFee(cfg config.Config, votingId int64) (*big.Int, error) {
	networkParam := cfg.VotingV2Config()
	client := networkParam.RPC

	feeCap, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed get gas price")
	}
	if feeCap.Uint64() == 0 {
		feeCap = big.NewInt(1)
	}
	votingInfo, err := pg.NewCheckerQ(cfg.DB()).GetVotingInfo(votingId)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

	}

	gasLimit := int64(cfg.VotingV2Config().GasLimit)
	if votingInfo != nil {
		gasLimit = int64(votingInfo.GasLimit)
	}

	totalFee := new(big.Int).Mul(big.NewInt(gasLimit), feeCap)
	return totalFee, nil
}

func CheckUpdateGasLimit(value uint64, cfg config.Config, votingId int64) error {
	pgDB := pg.NewMaterDB(cfg.DB())

	voteInfo, err := pgDB.CheckerQ().GetVotingInfo(votingId)
	if err != nil {
		return err
	}
	if voteInfo == nil {
		return errors.New("Vote Not Found")
	}

	voteInfo.GasLimit = value
	err = pgDB.CheckerQ().UpdateVotingInfo(voteInfo)
	if err != nil {
		return err
	}
	return nil
}

func GetAmountForCountTx(cfg config.Config, votingId int64, countTx *big.Int) (*big.Int, error) {
	txFee, err := getTxFee(cfg, votingId)
	if err != nil {
		return nil, errors.Wrap(err, "failed get fee")
	}
	if txFee.Int64() == 0 {
		big.NewInt(int64(cfg.VotingV2Config().GasLimit))
	}
	return new(big.Int).Mul(countTx, txFee), nil
}

func UpdateVotingBalance(cfg config.Config, gasPrice *big.Int, gas uint64, votingId int64) error {
	checkIsEnough, err := IsEnough(cfg, votingId)
	if err != nil {
		return errors.Wrap(err, "Insufficient balance check failed")
	}

	if !checkIsEnough {
		return errors.New("Insufficient funds in the voting account")
	}

	pgDB := pg.NewMaterDB(cfg.DB()).CheckerQ()
	voteInfo, err := pgDB.GetVotingInfo(votingId)
	if err != nil {
		return fmt.Errorf("failed get voting info from db: %w", err)
	}
	if voteInfo == nil {
		return errors.New("Vote Not Found")
	}
	if gasPrice.Int64() == 0 {
		gasPrice = big.NewInt(1)
	}
	voteInfo.Balance = new(big.Int).Sub(voteInfo.Balance, new(big.Int).Mul(big.NewInt(int64(gas)), gasPrice))

	err = pgDB.UpdateVotingInfo(voteInfo)
	if err != nil {
		return fmt.Errorf("failed update voting info from db: %w", err)
	}
	return err
}

type ProposalDescription struct {
	Title           string              `json:"title"`
	Description     string              `json:"description"`
	AcceptedOptions []resources.Options `json:"acceptedOptions"`
}

func GetProposalList(cfg config.Config) ([]*resources.VotingInfoAttributes, error) {
	var result []*resources.VotingInfoAttributes

	pgDB := pg.NewMaterDB(cfg.DB()).CheckerQ()
	votingInfoList, err := pgDB.SelectVotes()
	if err != nil {
		return nil, err
	}

	for _, vote := range votingInfoList {
		result = append(result, &vote.ProposalInfoWithConfig)
	}
	return result, nil
}

func GetProposalInfo(proposalId int64, cfg config.Config, creatorAddress string) (*resources.VotingInfoAttributes, error) {
	contractAddress := cfg.VotingV2Config().Address
	client := cfg.VotingV2Config().RPC
	ipfsUrl := cfg.VotingV2Config().IpfsUrl
	caller, err := contracts.NewProposalsStateCaller(contractAddress, client)
	if err != nil {
		return nil, err
	}

	proposalInfo, err := caller.GetProposalInfo(nil, big.NewInt(proposalId))
	if err != nil {
		return nil, err
	}
	startTimeStamp := proposalInfo.Config.StartTimestamp
	endTimestamp := proposalInfo.Config.StartTimestamp + proposalInfo.Config.Duration

	desc, err := getProposalDescFromIpfs(proposalInfo.Config.Description, ipfsUrl)
	if err != nil {
		return nil, err
	}
	votingConfig := proposalInfo.Config
	var votingWhitelist []string

	for _, addr := range votingConfig.VotingWhitelist {
		votingWhitelist = append(votingWhitelist, addr.Hex())
	}
	for _, addr := range votingConfig.VotingWhitelist {
		votingWhitelist = append(votingWhitelist, addr.Hex())
	}
	var votingWhitelistData []string
	for _, data := range votingConfig.VotingWhitelistData {
		votingWhitelistData = append(votingWhitelistData, hex.EncodeToString(data))
	}
	return &resources.VotingInfoAttributes{
		Author:   creatorAddress,
		Metadata: *desc,
		Contract: resources.VotingInfoContractInfo{
			ProposalSMT: proposalInfo.ProposalSMT.Hex(),
			Status:      proposalInfo.Status,
			Config: resources.VotingInfoConfig{
				Description:         votingConfig.Description,
				StartTimestamp:      int64(startTimeStamp),
				EndTimestamp:        int64(endTimestamp),
				Multichoice:         votingConfig.Multichoice.Int64(),
				ProposalId:          proposalId,
				VotingWhitelist:     votingWhitelist,
				VotingWhitelistData: votingWhitelistData,
			},
		},
	}, nil

}

func getProposalDescFromIpfs(desId string, ipfsUrl string) (*resources.VotingInfoAttributesMetadata, error) {
	requestURL := fmt.Sprintf("%s/ipfs/%s", ipfsUrl, desId)
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read body response: %v", err)
	}

	var data ProposalDescription
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed parse JSON: %v", err)
	}
	return &resources.VotingInfoAttributesMetadata{
		AcceptedOptions: data.AcceptedOptions,
		Title:           data.Title,
		Description:     data.Description,
	}, nil
}
