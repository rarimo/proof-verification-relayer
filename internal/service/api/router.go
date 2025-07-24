package api

import (
	"github.com/go-chi/chi"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/handlers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3/errors"
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
			handlers.CtxVotingV2Config(s.cfg.VotingV2Config()),
		),
	)
	r.Route("/integrations/proof-verification-relayer", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/register", handlers.Register)
			r.Post("/vote", handlers.Vote)
			r.Post("/transit-state", handlers.TransitState)
			r.Get("/state", handlers.GetSignedState)
		})
		r.Route("/v2", func(r chi.Router) {
			if s.cfg.VotingV2Config().Enable {
				r.Get("/count-remaining-votes/{voting_id}", handlers.VoteCountHandlers)
				r.Get("/is-enough/{voting_id}", handlers.IsEnoughHandler)
				r.Post("/vote", handlers.VoteV2)
				r.Post("/predict", handlers.PredictHandlers)
				r.Get("/proposals", handlers.ProposalsInfo)
			}

			r.Get("/state", handlers.GetSignedStateV2)
		})

		r.Route("/v3", func(r chi.Router) {
			r.Post("/vote", handlers.VoteV3)
		})

	})

	return r
}
