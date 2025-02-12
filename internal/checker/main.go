package checker

import (
	"context"
	"database/sql"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// Instead of "abiusdc", use the ABI package
// that will be generated for the required contract.
func CheckEvents(ctx context.Context, cfg config.Config) {
	if !cfg.VotingV2Config().Enable {
		cfg.Log().WithField("Enable checker", cfg.VotingV2Config().Enable).Info("stop reading events")

		return
	}

	pgDB := pg.NewMaterDB(cfg.DB())
	logger := cfg.Log()
	networkParam := cfg.VotingV2Config()

	logger.Info("Start Checker...")
	startBlock, err := getStartBlockNumber(pgDB, logger, networkParam.Block)
	if err != nil {
		logger.Errorf("Failed get start block: %v", err)
		return
	}

	go readOldEvents(ctx, startBlock, cfg)
	if cfg.VotingV2Config().WithSub {
		readNewEventsSub(ctx, cfg)
	} else {
		readNewEvents(ctx, cfg)
	}

}

// Instead of "abiusdc", use the ABI package
// that will be generated for the required contract.
func processEvent(event *contracts.ProposalsStateProposalCreated, pg data.CheckerDB, logger *logan.Entry, defaultGasLimit uint64) {
	var processedEvent data.ProcessedEvent
	block := int64(event.Raw.BlockNumber)
	processedEvent.BlockNumber = block
	processedEvent.LogIndex = int64(event.Raw.Index)
	processedEvent.TransactionHash = event.Raw.TxHash[:]

	err := pg.CheckerQ().InsertProcessedEvent(processedEvent)
	if err != nil {
		return
	}

	ToAddress := strings.ToLower(event.ProposalSMT.Hex())
	value := event.FundAmount
	votingInfo, err := checkVoteAndGetBalance(pg, ToAddress, logger, defaultGasLimit)
	if err != nil {
		return
	}

	votingInfo.Balance = votingInfo.Balance + value.Uint64()

	err = pg.CheckerQ().UpdateVotingInfo(votingInfo)
	if err != nil {
		logger.Errorf("Failed Update voting balance: %v", err)
		return
	}
}

func checkVoteAndGetBalance(pg data.CheckerDB, address string, logger *logan.Entry, defaultGasL uint64) (data.VotingInfo, error) {
	votingInfo, err := pg.CheckerQ().GetVotingInfo(address)
	if err == sql.ErrNoRows {
		newVote := &data.VotingInfo{
			Address:  address,
			Balance:  0,
			GasLimit: defaultGasL,
		}

		err := pg.CheckerQ().InsertVotingInfo(*newVote)
		if err != nil {
			logger.Errorf("Failed insert new voting info: %v", err)
			return *newVote, err
		}
		return *newVote, nil
	}
	return votingInfo, err
}

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
	txFee, err := getTxFee(cfg)
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

	txFee, err := getTxFee(cfg)
	if err != nil {
		logger.Errorf("Failed get fee: %v", err)
		return 0, err
	}
	return amount / txFee, nil
}

func getTxFee(cfg config.Config) (uint64, error) {
	logger := cfg.Log()
	networkParam := cfg.VotingV2Config()
	client := networkParam.RPC

	feeCap, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Warnf("Failed get gas price")
		return 0, err
	}

	totalFee := networkParam.GasLimit * feeCap.Uint64()
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

func getStartBlockNumber(pgDB data.CheckerDB, logger *logan.Entry, defaultSBlock uint64) (uint64, error) {
	block, err := pgDB.CheckerQ().GetLastBlock()
	if err == sql.ErrNoRows {
		block = defaultSBlock
	} else if err != nil {
		logger.Errorf("Failed get block from db: %v", err)
		return defaultSBlock, err
	}

	return block, nil
}

func checkFilter(block, toBlock uint64, contract *contracts.ProposalsStateFilterer, cfg config.Config) error {
	logger := cfg.Log()
	pgDB := pg.NewMaterDB(cfg.DB())

	contractAddress := cfg.VotingV2Config().Address
	query := bind.FilterOpts{
		Start:   block,
		End:     &toBlock,
		Context: context.Background(),
	}
	filterLogs, err := contract.FilterProposalCreated(&query, nil)
	if err != nil {
		logger.WithField("Error", err).Info("failed to filter logs")
		return errors.Wrap(err, "failed to filter logs", logan.F{"address": contractAddress, "start_block": block})
	}
	events := 0
	for filterLogs.Next() {
		processEvent(filterLogs.Event, pgDB, logger, cfg.VotingV2Config().GasLimit)
		events += 1
	}

	logger.WithFields(logan.F{
		"from":   block,
		"to":     toBlock,
		"events": events,
	}).Debug("found events")
	return nil
}

func readNewEvents(ctx context.Context, cfg config.Config) error {
	client := cfg.VotingV2Config().RPC
	logger := cfg.Log()
	pinger := cfg.Pinger()
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		logger.WithField("Error", err).Info("failed to get latest block header")
		return errors.Wrap(err, "failed to get latest block header")
	}

	contract, err := contracts.NewProposalsStateFilterer(cfg.VotingV2Config().Address, client)
	if err != nil {
		return err
	}
	toBlock := header.Number.Uint64()
	logger.WithField("from", toBlock).Info("start reading events")

	for {
		select {
		case <-ctx.Done():
			logger.Info("unsubscribe from events")
			return nil
		default:
			time.Sleep(pinger.NormalPeriod)
		}
		block := toBlock
		toBlock = block + 1

		checkFilter(block, toBlock, contract, cfg)

	}
}
func readOldEvents(ctx context.Context, block uint64, cfg config.Config) error {
	logger := cfg.Log()
	client := cfg.VotingV2Config().RPC
	pinger := cfg.Pinger()
	contractAddress := cfg.VotingV2Config().Address
	logger.WithField("from", block).Info("start reading old events")
	contract, err := contracts.NewProposalsStateFilterer(contractAddress, client)

	if err != nil {
		return err
	}
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		logger.WithField("Error", err).Info("failed to get latest block header")
		return errors.Wrap(err, "failed to get latest block header")
	}

	for ; block < header.Number.Uint64(); block = min(header.Number.Uint64(), block+pinger.BlocksDistance) {
		select {
		case <-ctx.Done():
			logger.Info("unsubscribe from events")
			return nil
		default:
		}
		toBlock := min(header.Number.Uint64(), block+pinger.BlocksDistance)

		err := checkFilter(block, toBlock, contract, cfg)
		if err != nil {
			return err
		}

	}

	logger.WithField("to", block).Info("finish reading old events")
	return nil
}

func readNewEventsSub(ctx context.Context, cfg config.Config) error {
	logger := cfg.Log()
	client := cfg.VotingV2Config().RPC
	contractAddress := cfg.VotingV2Config().Address
	pgDB := pg.NewMaterDB(cfg.DB())
	parsedABI, err := abi.JSON(strings.NewReader(contracts.ProposalsStateABI))
	if err != nil {
		logger.Errorf("Failed to parse contract ABI: %v", err)
		return err
	}

	query := ethereum.FilterQuery{
		Topics:    [][]common.Hash{{parsedABI.Events["ProposalCreated"].ID}},
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed sub: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			logger.Info("unsubscribe from events")
			return nil
		case err := <-sub.Err():
			logger.Infof("Failed subscribe event: %v", err)
		case vLog := <-logs:
			logger.Warnf("log from block: %v", vLog.BlockNumber)
			processLog(vLog, pgDB, logger, cfg.VotingV2Config().GasLimit)

		}
	}

}

func processLog(vLog types.Log, pg data.CheckerDB, logger *logan.Entry, defaultGasLimit uint64) {
	var processedEvent data.ProcessedEvent
	parsedABI, err := abi.JSON(strings.NewReader(contracts.ProposalsStateABI))
	if err != nil {
		logger.Errorf("Failed to parse contract ABI: %v", err)
		return
	}

	var transferEvent struct {
		Value *big.Int
	}
	err = parsedABI.UnpackIntoInterface(&transferEvent, "ProposalCreated", vLog.Data)
	if err != nil {
		logger.Errorf("Failed to unpack log data: %v", err)
		return
	}
	block := int64(vLog.BlockNumber)
	processedEvent.BlockNumber = block
	processedEvent.LogIndex = int64(vLog.Index)
	processedEvent.TransactionHash = vLog.TxHash[:]

	err = pg.CheckerQ().InsertProcessedEvent(processedEvent)
	if err != nil {
		return
	}
	to := common.HexToAddress(vLog.Topics[2].Hex())

	ToAddress := strings.ToLower(to.Hex())
	value := transferEvent.Value
	votingInfo, err := checkVoteAndGetBalance(pg, ToAddress, logger, defaultGasLimit)
	if err != nil {
		return
	}

	votingInfo.Balance = votingInfo.Balance + value.Uint64()

	err = pg.CheckerQ().UpdateVotingInfo(votingInfo)
	if err != nil {
		logger.Errorf("Failed Update voting balance: %v", err)
		return
	}
}
