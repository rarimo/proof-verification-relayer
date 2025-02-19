package handlers

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	configCtxKey
	networkConfigCtxKey
	ethClientCtxKey
	voteVerifierRegisterMethodCtxKey
	votingVoteMethodCtxKey
	votingRegistryCtxKey
	lightweightStateCtxKey
	signedTransitStateCtxKey
	stateQCtxKey
	relayerConfigCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxConfig(cfg config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, configCtxKey, cfg)
	}
}

func Config(r *http.Request) config.Config {
	return r.Context().Value(configCtxKey).(config.Config)
}

func CtxVoteVerifierRegisterMethod(method *abi.Method) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, voteVerifierRegisterMethodCtxKey, method)
	}
}

func VoteVerifierRegisterMethod(r *http.Request) *abi.Method {
	return r.Context().Value(voteVerifierRegisterMethodCtxKey).(*abi.Method)
}

func CtxVotingRegistry(registry *contracts.VotingRegistry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, votingRegistryCtxKey, registry)
	}
}

func VotingRegistry(r *http.Request) *contracts.VotingRegistry {
	return r.Context().Value(votingRegistryCtxKey).(*contracts.VotingRegistry)
}

func CtxLightweightState(registry *contracts.LightweightState) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, lightweightStateCtxKey, registry)
	}
}

func LightweightState(r *http.Request) *contracts.LightweightState {
	return r.Context().Value(lightweightStateCtxKey).(*contracts.LightweightState)
}

func CtxSignedTransitStateMethod(method *abi.Method) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, signedTransitStateCtxKey, method)
	}
}

func SignedTransitStateMethod(r *http.Request) *abi.Method {
	return r.Context().Value(signedTransitStateCtxKey).(*abi.Method)
}

func CtxStateQ(stateQ data.StateQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, stateQCtxKey, stateQ)
	}
}

func StateQ(r *http.Request) data.StateQ {
	return r.Context().Value(stateQCtxKey).(data.StateQ)
}

func CtxVotingV2Config(cfg *config.VotingV2Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, relayerConfigCtxKey, cfg)
	}
}

func VotingV2Config(r *http.Request) *config.VotingV2Config {
	return r.Context().Value(relayerConfigCtxKey).(*config.VotingV2Config)
}
