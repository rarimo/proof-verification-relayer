package api

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	verifierABI, err := contracts.VoteVerifierMetaData.GetAbi()
	if err != nil {
		panic(errors.Wrap(err, "failed to get vote verifier ABI"))
	}

	registerMethod, ok := verifierABI.Methods["register"]
	if !ok {
		panic(errors.New("register method not found"))
	}

	votingRegistry, err := contracts.NewVotingRegistry(s.cfg.ContractsConfig()[config.VotingRegistry].Address, s.cfg.NetworkConfig().Client)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize new voting registry"))
	}

	lightweightState, err := contracts.NewLightweightState(s.cfg.ContractsConfig()[config.LightweightState].Address, s.cfg.NetworkConfig().Client)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize new lightweight state"))
	}

	lightweightStateABI, err := contracts.LightweightStateMetaData.GetAbi()
	if err != nil {
		panic(errors.Wrap(err, "failed to get lightweight state ABI"))
	}

	signedTransitState, ok := lightweightStateABI.Methods["signedTransitState"]
	if !ok {
		panic(errors.New("signedTransitState method not found"))
	}

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxConfig(s.cfg),
			handlers.CtxVoteVerifierRegisterMethod(&registerMethod),
			handlers.CtxVotingRegistry(votingRegistry),
			handlers.CtxLightweightState(lightweightState),
			handlers.CtxSignedTransitStateMethod(&signedTransitState),
			handlers.CtxStateQ(pg.NewStateQ(s.cfg.DB().Clone())),
			handlers.CtxRelayerConfig(s.cfg.RelayerConfig()),
		),
	)
	r.Route("/integrations/proof-verification-relayer", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/register", handlers.Register)
			r.Post("/vote", handlers.Vote)
			r.Post("/transit-state", handlers.TransitState)
			r.Get("/state", handlers.GetSignedState)
		})
		if s.cfg.RelayerConfig().Enable {
			r.Route("/v2", func(r chi.Router) {
				r.Get("/predict/{address}", handlers.PredictHandlers)
				r.Get("/is-enough/{address}", handlers.IsEnoughHandler)
				r.Post("/vote", handlers.Voting)

			})
		}
	})

	return r
}
