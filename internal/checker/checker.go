package checker

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

type Checker interface {
	Start(ctx context.Context)
}

type checker struct {
	log            *logan.Entry
	client         *ethclient.Client
	checkerQ       data.CheckerQ
	VotingV2Config config.VotingV2Config
	pinger         config.Pinger
	cfg            config.Config
}

func NewChecker(
	cfg config.Config,
) Checker {
	return &checker{
		log: cfg.Log().WithFields(logan.F{
			"service": "checker",
			"address": cfg.VotingV2Config().Address.Hex(),
		}),
		client:         cfg.VotingV2Config().RPC,
		checkerQ:       pg.NewMaterDB(cfg.DB()).CheckerQ(),
		VotingV2Config: *cfg.VotingV2Config(),
		pinger:         cfg.Pinger(),
		cfg:            cfg,
	}
}

func (ch *checker) Start(ctx context.Context) {
	running.WithBackOff(
		ctx, ch.log, "checker", ch.check,
		ch.pinger.NormalPeriod, ch.pinger.MinAbnormalPeriod, ch.pinger.MinAbnormalPeriod,
	)
}

func (ch *checker) check(ctx context.Context) error {
	if !ch.VotingV2Config.Enable {
		ch.log.WithField("Enable checker", ch.VotingV2Config.Enable).Info("stop reading events")
		return nil
	}
	ch.log.Info("Start Checker...")
	startBlock, err := ch.getStartBlockNumber()
	if err != nil {
		ch.log.Errorf("Failed get start block: %v", err)
		return err
	}

	go ch.readNewEvents(ctx, ch.VotingV2Config.WithSub)
	time.Sleep(ch.pinger.Timeout)

	go ch.readOldEvents(ctx, startBlock)
	for range ctx.Done() {
		ch.log.Info("unsubscribe from events")
		return nil
	}
	return nil
}

func (ch *checker) readNewEvents(ctx context.Context, isSub bool) {
	if isSub {
		go ch.readNewEventsSub(ctx)
		return
	}
	go ch.readNewEventsWithoutSub(ctx)
}

func (ch *checker) readNewEventsWithoutSub(ctx context.Context) {
	client := ch.VotingV2Config.RPC
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		ch.log.WithField("Error", err).Info("failed to get latest block header")
		return
	}

	contract, err := contracts.NewProposalsStateFilterer(ch.VotingV2Config.Address, client)

	if err != nil {
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
			time.Sleep(10 * time.Second)
		}

		err := ch.checkFilter(block, toBlock, contract)
		if err != nil {
			ch.log.WithFields(logan.F{
				"from":  block,
				"to":    toBlock,
				"Error": err,
			}).Info("failed check blocks")
			continue
		}
		block = toBlock
		toBlock = block + 1

	}
}

func (ch *checker) readNewEventsSub(ctx context.Context) error {
	parsedABI, err := abi.JSON(strings.NewReader(contracts.ProposalsStateABI))

	if err != nil {
		ch.log.Errorf("Failed to parse contract ABI: %v", err)
		return err
	}

	query := ethereum.FilterQuery{
		Topics: [][]common.Hash{{parsedABI.Events["ProposalCreated"].ID}},

		Addresses: []common.Address{ch.VotingV2Config.Address},
	}
	logs := make(chan types.Log)
	sub, err := ch.VotingV2Config.RPC.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed sub: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			ch.log.Info("unsubscribe from events")
			return nil
		case err := <-sub.Err():
			ch.log.Infof("Failed subscribe event: %v", err)
			time.Sleep(ch.pinger.Timeout)
		case vLog := <-logs:
			err := ch.processLog(vLog)
			if err != nil {
				ch.log.WithFields(logan.F{"Error": err, "log_index": vLog.Index, "hash_tx": vLog.TxHash.Hex()}).Info("failed process log")
			}
		}
	}

}

func (ch *checker) checkFilter(block, toBlock uint64, contract *contracts.ProposalsStateFilterer) error {
	contractAddress := ch.VotingV2Config.Address
	query := bind.FilterOpts{
		Start:   block,
		End:     &toBlock,
		Context: context.Background(),
	}

	filterLogs, err := contract.FilterProposalCreated(&query, nil)
	if err != nil {
		ch.log.WithField("Error", err).Info("failed to filter logs")
		return errors.Wrap(err, "failed to filter logs", logan.F{"address": contractAddress, "start_block": block})
	}
	events := 0
	for filterLogs.Next() {
		err := ch.processEvent(filterLogs.Event)
		// err := ch.processEvent(filterLogs.Event)
		if err != nil {
			ch.log.WithFields(logan.F{"Error": err, "log_index": filterLogs.Event.Raw.Index, "hash_tx": filterLogs.Event.Raw.TxHash.Hex()}).Info("failed process log")
			continue
		}
		events += 1
	}

	ch.log.WithFields(logan.F{
		"from":   block,
		"to":     toBlock,
		"events": events,
	}).Debug("found events")
	return nil
}

func (ch *checker) readOldEvents(ctx context.Context, block uint64) error {
	client := ch.VotingV2Config.RPC
	contractAddress := ch.VotingV2Config.Address
	ch.log.WithField("from", block).Info("start reading old events")

	contract, err := contracts.NewProposalsStateFilterer(contractAddress, client)
	if err != nil {
		return err
	}
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		ch.log.WithField("Error", err).Info("failed to get latest block header")
		return errors.Wrap(err, "failed to get latest block header")
	}

	for ; block < header.Number.Uint64(); block = min(header.Number.Uint64(), block+ch.pinger.BlocksDistance) {
		select {
		case <-ctx.Done():
			ch.log.Info("unsubscribe from events")
			return nil
		default:
		}
		toBlock := min(header.Number.Uint64(), block+ch.pinger.BlocksDistance)

		err := ch.checkFilter(block, toBlock, contract)
		if err != nil {
			ch.log.WithFields(logan.F{
				"from":  block,
				"to":    toBlock,
				"Error": err,
			}).Info("failed check blocks")
			block = max(0, block-ch.pinger.BlocksDistance)
		}
	}

	ch.log.WithField("to", block).Info("finish reading old events")
	return nil
}

func Run(ctx context.Context, cfg config.Config) {
	NewChecker(cfg).Start(ctx)
}
