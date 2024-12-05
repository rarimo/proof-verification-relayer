package listener

import (
	"context"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
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

const (
	serviceName                = "listener"
	RootTransitionedEventTopic = "0x287d7075e3fdd1ee3cb7eef1d33839a4b50939e7bc33a68d8f6031eb3a1a14c6"
)

type Listener interface {
	Start(ctx context.Context)
}

type listener struct {
	log           *logan.Entry
	client        *ethclient.Client
	stateQ        data.StateQ
	smtReplicator *contracts.RegistrationSMTReplicator
	contractCfg   config.ContractConfig
	pinger        config.Pinger
}

func NewListener(
	cfg config.Config,
) (Listener, error) {
	smtReplicatorCfg := cfg.ContractsConfig()[config.SMTReplicator]
	smtReplicatorContract, err := contracts.NewRegistrationSMTReplicator(
		smtReplicatorCfg.Address,
		cfg.NetworkConfig().Client,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create smtReplicator contractCfg")
	}

	return &listener{
		log: cfg.Log().WithFields(logan.F{
			"service": serviceName,
			"address": smtReplicatorCfg.Address.String(),
		}),
		client:        cfg.NetworkConfig().Client,
		stateQ:        pg.NewStateQ(cfg.DB().Clone()),
		smtReplicator: smtReplicatorContract,
		contractCfg:   smtReplicatorCfg,
		pinger:        cfg.Pinger(),
	}, nil
}

func (l *listener) Start(ctx context.Context) {
	running.WithBackOff(
		ctx, l.log, serviceName, l.listen,
		l.pinger.NormalPeriod, l.pinger.MinAbnormalPeriod, l.pinger.MinAbnormalPeriod,
	)
	return
}

func (l *listener) listen(ctx context.Context) error {
	l.log.Info("start listening events")

	block, err := l.getStartBlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get start block number")
	}

	if err = l.listenEvents(ctx, block); err != nil {
		return errors.Wrap(err, "failed to listen events")
	}

	l.log.Info("finish listening events")
	return nil
}

func (l *listener) getStartBlockNumber(ctx context.Context) (uint64, error) {
	chainID, err := l.client.ChainID(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get chain ID")
	}

	state, err := l.stateQ.SortByBlockHeight(data.DESC).Get()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get states", logan.F{"chain_id": chainID.Uint64(), "address": l.contractCfg.Address.String()})
	}

	if state == nil {
		return l.contractCfg.Block, nil
	}

	return max(l.contractCfg.Block, state.Block), nil
}

func (l *listener) readEvents(ctx context.Context, block uint64) (uint64, error) {
	l.log.WithField("from", block).Info("start reading events")

	header, err := l.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return block, errors.Wrap(err, "failed to get latest block header")
	}

	for ; block < header.Number.Uint64(); block = min(header.Number.Uint64(), block+l.pinger.BlocksDistance) {
		toBlock := min(header.Number.Uint64(), block+l.pinger.BlocksDistance)
		logs, err := l.client.FilterLogs(ctx, ethereum.FilterQuery{
			Topics:    [][]common.Hash{{common.HexToHash(RootTransitionedEventTopic)}},
			Addresses: []common.Address{l.contractCfg.Address},
			FromBlock: new(big.Int).SetUint64(block),
			ToBlock:   new(big.Int).SetUint64(toBlock),
		})
		if err != nil {
			return block, errors.Wrap(err, "failed to filter logs", logan.F{"address": l.contractCfg.Address, "start_block": block})
		}

		for _, event := range logs {
			if err = l.handleLog(event); err != nil {
				return block, errors.Wrap(err, "failed to process old event", logan.F{"event": event.Address.String()})
			}
		}

		l.log.WithFields(logan.F{
			"from":   block,
			"to":     toBlock,
			"events": len(logs),
		}).Debug("found events")
	}

	l.log.WithField("to", block).Info("finish reading events")
	return block, nil
}

func (l *listener) listenEvents(ctx context.Context, block uint64) error {
	l.log.Info("subscribe for events")

	var err error
	ticker := time.NewTicker(l.pinger.Timeout)
	defer ticker.Stop()

	for {
		block, err = l.readEvents(ctx, block)
		if err != nil {
			return errors.Wrap(err, "failed to read events")
		}

		select {
		case <-ctx.Done():
			l.log.Warn("unsubscribe from events")
			return ctx.Err()
		case <-ticker.C:
			continue
		}
	}
}

func (l *listener) handleLog(event types.Log) error {
	rootTransitioned, err := l.smtReplicator.ParseRootTransitioned(event)
	if err != nil {
		return errors.Wrap(err, "failed to parse root transitioned event")
	}

	if err := l.stateQ.Upsert(data.State{
		TxHash:    hex.EncodeToString(event.TxHash.Bytes()),
		Block:     event.BlockNumber,
		Root:      hex.EncodeToString(rootTransitioned.NewRoot[:]),
		Timestamp: rootTransitioned.TransitionTimestamp.Uint64(),
	}); err != nil {
		return errors.Wrap(err, "failed to insert new root state")
	}

	return nil
}

func Run(ctx context.Context, cfg config.Config) {
	pinger, err := NewListener(cfg)
	if err != nil {
		panic(errors.Wrap(err, "failed to create listener"))
	}

	pinger.Start(ctx)
}
