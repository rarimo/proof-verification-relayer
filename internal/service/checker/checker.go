package checker

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

type checker struct {
	log             *logan.Entry
	client          *ethclient.Client
	votingQ         data.VotingQ
	processedEventQ data.ProcessedEventQ
	VotingV2Config  config.VotingV2Config
	pinger          config.Pinger
	cfg             config.Config
	eventHashes     []common.Hash
	parsedABI       abi.ABI
}

// event ProposalConfigChanged(uint256 indexed proposalId);
// event ProposalFunded(uint256 indexed proposalId, uint256 fundAmount);
// event ProposalCreated(uint256 indexed proposalId, address proposalSMT, uint256 fundAmount)
var eventNames = []string{"ProposalCreated", "ProposalFunded", "ProposalConfigChanged"}
var eventHashIdToNames = map[string]string{
	"0x3417b456fad6209c73445d5efd446d686e75e4560f0f50c13b5a5cde976447b4": "ProposalCreated",
	"0xe626ac29dddf069f3357c054c144b90e4e79f25ee6c1020f9142bae435477e6c": "ProposalFunded",
	"0xa9cc646240fc6ba1b4b124e96765839b67cd0e2e698942d5d5948a36c7b998d5": "ProposalConfigChanged",
}

func Run(ctx context.Context, cfg config.Config) {
	NewChecker(cfg).Start(ctx)
}

type Checker interface {
	Start(ctx context.Context)
}

func NewChecker(
	cfg config.Config,
) Checker {
	return &checker{
		log: cfg.Log().WithFields(logan.F{
			"service": "checker",
			"address": cfg.VotingV2Config().Address.Hex(),
		}),
		client:          cfg.VotingV2Config().RPC,
		votingQ:         pg.NewVotingQ(cfg.DB().Clone()),
		VotingV2Config:  *cfg.VotingV2Config(),
		pinger:          cfg.Pinger(),
		cfg:             cfg,
		processedEventQ: pg.NewProcessedEventsQ(cfg.DB().Clone()),
	}
}

func (ch *checker) Start(ctx context.Context) {
	if !ch.VotingV2Config.Enable {
		ch.log.WithField("Enable checker", ch.VotingV2Config.Enable).Info("stop reading events")
		return
	}
	eventHashes, err := ch.getEventHashes()
	if err != nil {
		// If there is no list of events, scanning is pointless
		ch.log.WithError(err).Error("failed to get event hash list")
		return
	}
	ch.eventHashes = eventHashes
	ch.parsedABI, err = abi.JSON(strings.NewReader(contracts.ProposalsStateABI))
	if err != nil {
		ch.log.WithError(err).Error(err, "failed to parse contract ABI")
		return
	}
	running.WithBackOff(
		ctx, ch.log, "checker", ch.check,
		ch.pinger.NormalPeriod, ch.pinger.MinAbnormalPeriod, ch.pinger.MinAbnormalPeriod,
	)
}

func (ch *checker) check(ctx context.Context) error {
	ch.log.Info("Start Checker...")

	startBlock, err := ch.getStartBlockNumber()
	if err != nil {
		return errors.Wrap(err, "Failed get start block")
	}

	ch.readOldEvents(ctx, startBlock)
	startBlock, err = ch.getStartBlockNumber()
	if err != nil {
		return errors.Wrap(err, "Failed get start block")
	}
	go ch.readNewEvents(ctx, ch.VotingV2Config.WithSub)

	time.Sleep(ch.pinger.NormalPeriod)

	ch.readOldEvents(ctx, startBlock)

	<-ctx.Done()
	ch.log.Info("unsubscribe from events")

	return nil
}

// readNewEvents starts scanning for new events
func (ch *checker) readNewEvents(ctx context.Context, isSub bool) {
	if isSub {
		go ch.readNewEventsSub(ctx)
		return
	}
	go ch.readNewEventsWithoutSub(ctx)
}

// readNewEventsWithoutSub check events without a subscription, for example, if the network does not support web sockets
func (ch *checker) readNewEventsWithoutSub(ctx context.Context) {
	client := ch.VotingV2Config.RPC
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		ch.log.WithError(err).Error("failed to get latest block header")
		return
	}

	toBlock := header.Number.Uint64()
	ch.log.WithField("from", toBlock).Info("start reading events")
	block := toBlock
	toBlock = block + 1

	for {
		select {
		case <-ctx.Done():
			ch.log.Info("unsubscribe from events")
			return
		default:
			time.Sleep(ch.pinger.Timeout)
		}

		err := ch.checkFilter(block, toBlock)
		if err != nil {
			ch.log.WithError(err).WithFields(logan.F{
				"from": block,
				"to":   toBlock,
			}).Error("failed check blocks")
			continue
		}
		header, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			ch.log.WithError(err).Error("failed to get latest block header")
			continue
		}

		block = toBlock
		toBlock = header.Number.Uint64()
	}
}

func (ch *checker) readNewEventsSub(ctx context.Context) {
	query := ethereum.FilterQuery{
		Topics:    [][]common.Hash{ch.eventHashes},
		Addresses: []common.Address{ch.VotingV2Config.Address},
	}

	logs := make(chan types.Log)
	sub, err := ch.VotingV2Config.RPC.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		ch.log.WithError(err).Error("failed subscribe event")
		return
	}

	for {
		select {
		case <-ctx.Done():
			ch.log.Info("unsubscribe from events")
			return
		case err := <-sub.Err():
			ch.log.WithError(err).Error("failed subscribe event")
			return
		case vLog := <-logs:
			eventName := eventHashIdToNames[vLog.Topics[0].Hex()]
			err := ch.processLog(vLog, eventName)
			if err != nil {
				ch.log.WithError(err).WithFields(logan.F{
					"log_index":  vLog.Index,
					"hash_tx":    vLog.TxHash.Hex(),
					"event_name": eventName}).
					Error("failed process log")
			}
			ch.log.WithFields(logan.F{
				"log_index":  vLog.Index,
				"hash_tx":    vLog.TxHash.Hex(),
				"event_name": eventName}).
				Info("new event")
		}
	}
}

// Search and process old events starting from a specific block
func (ch *checker) readOldEvents(ctx context.Context, block uint64) {
	client := ch.VotingV2Config.RPC
	ch.log.WithField("from", block).Info("start reading old events")

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		ch.log.WithError(err).Error("failed to get latest block header")
		return
	}

	toBlock := ch.readFromToBlock(ctx, block, header.Number.Uint64())
	ch.log.WithField("to", toBlock).Info("finish reading old events")
}

// readFromToBlock sequential block gap check. It is necessary to check gaps that are larger than the network limits.
func (ch *checker) readFromToBlock(ctx context.Context, fromBlock uint64, toBlock uint64) uint64 {
	for ; fromBlock < toBlock; fromBlock = min(toBlock, fromBlock+ch.pinger.BlocksDistance) {
		select {
		case <-ctx.Done():
			ch.log.Info("unsubscribe from events")
			return toBlock
		default:
		}
		toBlock := min(toBlock, fromBlock+ch.pinger.BlocksDistance)

		err := ch.checkFilter(fromBlock, toBlock)
		if err != nil {
			ch.log.WithError(err).WithFields(logan.F{
				"from": fromBlock,
				"to":   toBlock,
			}).Error("failed check blocks")
			fromBlock = max(0, fromBlock-ch.pinger.BlocksDistance)
		}
	}
	return toBlock
}

// checkFilter filters logs with the required events
func (ch *checker) checkFilter(block, toBlock uint64) error {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(block)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Topics:    [][]common.Hash{ch.eventHashes},
		Addresses: []common.Address{ch.VotingV2Config.Address},
	}

	logs, err := ch.VotingV2Config.RPC.FilterLogs(context.Background(), query)
	if err != nil {
		return errors.Wrap(err, "failed get filter logs")
	}
	events := 0
	for _, vLog := range logs {
		eventName := eventHashIdToNames[vLog.Topics[0].Hex()]
		err := ch.processLog(vLog, eventName)
		if err != nil {
			ch.log.WithError(err).WithFields(logan.F{
				"log_index": vLog.Index,
				"hash_tx":   vLog.TxHash.Hex(),
				"eventName": eventName,
			}).Error("failed process log")
			continue
		}
		ch.log.WithFields(logan.F{
			"log_index":  vLog.Index,
			"hash_tx":    vLog.TxHash.Hex(),
			"event_name": eventName}).
			Info("new event")
		events += 1
	}

	ch.log.WithFields(logan.F{
		"from":   block,
		"to":     toBlock,
		"events": events,
	}).Debug("found events")

	return nil
}
