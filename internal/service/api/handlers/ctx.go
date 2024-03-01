package handlers

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	networkConfigCtxKey
	ethClientCtxKey
	masterQKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxNetworkConfig(cfg *config.NetworkConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, networkConfigCtxKey, cfg)
	}
}

func NetworkConfig(r *http.Request) *config.NetworkConfig {
	return r.Context().Value(networkConfigCtxKey).(*config.NetworkConfig)
}

func CtxEthClient(client *ethclient.Client) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ethClientCtxKey, client)
	}
}

func EthClient(r *http.Request) *ethclient.Client {
	return r.Context().Value(ethClientCtxKey).(*ethclient.Client)
}

func CtxMasterQ(entry data.MasterQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, masterQKey, entry)
	}
}

func MasterQ(r *http.Request) data.MasterQ {
	return r.Context().Value(masterQKey).(data.MasterQ).New()
}
