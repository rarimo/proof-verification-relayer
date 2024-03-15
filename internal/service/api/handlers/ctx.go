package handlers

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	networkConfigCtxKey
	ethClientCtxKey
	voteVerifierRegisterMethodCtxKey
	votingVoteMethodCtxKey
	votingRegistryCtxKey
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

func CtxVoteVerifierRegisterMethod(method *abi.Method) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, voteVerifierRegisterMethodCtxKey, method)
	}
}

func VoteVerifierRegisterMethod(r *http.Request) *abi.Method {
	return r.Context().Value(voteVerifierRegisterMethodCtxKey).(*abi.Method)
}

func CtxVotingVoteMethod(method *abi.Method) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, votingVoteMethodCtxKey, method)
	}
}

func VotingVoteMethod(r *http.Request) *abi.Method {
	return r.Context().Value(votingVoteMethodCtxKey).(*abi.Method)
}

func CtxVotingRegistry(registry *contracts.VotingRegistry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, votingRegistryCtxKey, registry)
	}
}

func VotingRegistry(r *http.Request) *contracts.VotingRegistry {
	return r.Context().Value(votingRegistryCtxKey).(*contracts.VotingRegistry)
}
