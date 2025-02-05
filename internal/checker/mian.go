package checker

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	usdcabi "github.com/rarimo/proof-verification-relayer/internal/checker/usdc_abi"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"gitlab.com/distributed_lab/logan/v3"
)

// Instead of "abiusdc", use the ABI package
// that will be generated for the required contract.
func CheckEvents(ctx context.Context, cfg config.Config) {
	pgDB := pg.NewMaterDB(cfg.DB())
	logger := cfg.Log()
	networkParam := cfg.RelayerConfig()
	gasLimit := networkParam.GasLimit
	contractAddress := networkParam.Address

	client := networkParam.RPC
	defer client.Close()

	startBlock, err := getStartBlockNumber(pgDB, logger, networkParam.Block)
	if err != nil {
		logger.Errorf("Failed get start block: %v", err)
		return
	}
	if startBlock == 0 {
		startBlock, err = client.BlockNumber(context.Background())
		if err != nil {
			logger.Errorf("Failed get start block: %v", err)
			return
		}
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(startBlock)),
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
			return
		case err := <-sub.Err():
			logger.Infof("Failed subscribe event: %v", err)
		case vLog := <-logs:
			processLog(vLog, pgDB, logger, &gasLimit, client)
			startBlock = startBlock + 1

		}
	}

}

// Instead of "abiusdc", use the ABI package
// that will be generated for the required contract.
func processLog(vLog types.Log, pg data.CheckerDB, logger *logan.Entry, defaultGasLimit *uint64, client *ethclient.Client) {
	if len(vLog.Topics) < 3 {
		logger.Warn("Invalid log format for Transfer event")
		return
	}
	transferEventSignature := []byte("Transfer(address,address,uint256)")
	transferEventID := crypto.Keccak256Hash(transferEventSignature)

	if vLog.Topics[0] == transferEventID {
		var processedEvent data.ProcessedEvent
		to := common.HexToAddress(vLog.Topics[2].Hex())

		parsedABI, err := abi.JSON(strings.NewReader(usdcabi.UsdcAbiABI))
		if err != nil {
			log.Fatalf("Failed to parse contract ABI: %v", err)
		}

		var transferEvent struct {
			Value *big.Int
		}
		err = parsedABI.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
		if err != nil {
			log.Printf("Failed to unpack log data: %v", err)
			return
		}
		block := int64(vLog.BlockNumber)
		processedEvent.BlockNumber = block

		processedEvent.EmittedAt, err = getEmittedAt(client, vLog)
		if err != nil {
			return
		}
		processedEvent.LogIndex = int64(vLog.Index)
		processedEvent.TransactionHash = vLog.TxHash[:]

		err = pg.CheckerQ().InsertProcessedEvent(processedEvent)

		if err != nil {
			return
		}

		ToAddress := to.Hex()
		value := transferEvent.Value

		votingInfo, err := checkVoteAndGetBalance(pg, ToAddress, logger, defaultGasLimit)

		if err != nil {
			return
		}
		votingInfo.Address = ToAddress
		votingInfo.Balance = votingInfo.Balance + value.Uint64()

		err = pg.CheckerQ().UpdateVotingInfo(votingInfo)
		if err != nil {
			logger.Errorf("Failed Update user balance: %v", err)
			return
		}
	}
}

func checkVoteAndGetBalance(pg data.CheckerDB, address string, logger *logan.Entry, defaultGasL *uint64) (data.VotingInfo, error) {
	votingInfo, err := pg.CheckerQ().GetVotingInfo(address)
	if err == sql.ErrNoRows {
		newVote := &data.VotingInfo{
			Address:  address,
			Balance:  0,
			GasLimit: *defaultGasL,
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
	logger := cfg.Log()
	pg := pg.NewMaterDB(cfg.DB())

	txFee, err := getTxFee(cfg)
	if err != nil {
		logger.Errorf("Failed get fee: %v", err)
		return false, err
	}

	votingInfo, err := pg.CheckerQ().GetVotingInfo(addr)
	if err != nil {
		return false, err
	}

	voteBalance := votingInfo.Balance
	if txFee <= voteBalance {
		return true, nil
	}
	return false, nil
}

func GetCountTx(cfg config.Config, addr string) (uint64, error) {
	logger := cfg.Log()
	pg := pg.NewMaterDB(cfg.DB())

	txFee, err := getTxFee(cfg)
	logger.Warnf("txFee: %d", txFee)
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

func getTxFee(cfg config.Config) (uint64, error) {
	logger := cfg.Log()
	networkParam := cfg.RelayerConfig()
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
	if err != nil {
		logger.Errorf("Failed get block: %v", err)
		return 0, err
	}

	switch block {
	case 0:
		block = defaultSBlock
	default:
	}
	return block, nil
}

func getEmittedAt(client *ethclient.Client, log types.Log) (time.Time, error) {
	blockNumber := big.NewInt(int64(log.BlockNumber))

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to get block: %v", err)
	}

	timestamp := block.Time()
	emittedAt := time.Unix(int64(timestamp), 0).UTC()
	return emittedAt, nil
}
