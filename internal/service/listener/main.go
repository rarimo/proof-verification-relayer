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
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

const (
	serviceName           = "listener"
	RootUpdatedEventTopic = "0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4"
)

var (
	ErrInvalidEventDataLen = errors.New("invalid data length for event")
)

type Listener interface {
	Start(ctx context.Context)
}

type listener struct {
	log      *logan.Entry
	client   *ethclient.Client
	stateQ   data.StateQ
	contract config.ContractConfig
	pinger   config.Pinger
}

func NewListener(
	cfg config.Config,
) Listener {
	register2Contract := cfg.ContractsConfig()[config.Register2]
	return &listener{
		log: cfg.Log().WithFields(logan.F{
			"service": serviceName,
			"address": register2Contract.Address.String(),
		}),
		client:   cfg.NetworkConfig().Client,
		stateQ:   pg.NewStateQ(cfg.DB().Clone()),
		contract: register2Contract,
		pinger:   cfg.Pinger(),
	}
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
		return 0, errors.Wrap(err, "failed to get states", logan.F{"chain_id": chainID.Uint64(), "address": l.contract.Address.String()})
	}

	if state == nil {
		return l.contract.Block, nil
	}

	return max(l.contract.Block, state.Block), nil
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
			Topics:    [][]common.Hash{{common.HexToHash(RootUpdatedEventTopic)}},
			Addresses: []common.Address{l.contract.Address},
			FromBlock: new(big.Int).SetUint64(block),
			ToBlock:   new(big.Int).SetUint64(toBlock),
		})
		if err != nil {
			return block, errors.Wrap(err, "failed to filter logs", logan.F{"address": l.contract.Address, "start_block": block})
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
	var root [32]byte
	if len(event.Data) != 32 {
		return ErrInvalidEventDataLen
	}
	copy(root[:], event.Data[:32])

	if err := l.stateQ.Upsert(data.State{
		TxHash:    hex.EncodeToString(event.TxHash.Bytes()),
		Block:     event.BlockNumber,
		Root:      hex.EncodeToString(root[:]),
		Timestamp: uint64(time.Now().Unix()),
	}); err != nil {
		return errors.Wrap(err, "failed to insert new root state")
	}

	return nil
}

func Run(ctx context.Context, cfg config.Config) {
	NewListener(cfg).Start(ctx)
}
