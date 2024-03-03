package service

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/handlers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	ethClient, err := ethclient.Dial(s.cfg.NetworkConfig().RPC)
	if err != nil {
		panic(errors.Wrap(err, "failed to dial connect"))
	}

	verifierABI, err := contracts.VoteVerifierMetaData.GetAbi()
	if err != nil {
		panic(errors.Wrap(err, "failed to get vote verifier ABI"))
	}

	method, ok := verifierABI.Methods["register"]
	if !ok {
		panic(errors.New("register method not found"))
	}

	votingRegistry, err := contracts.NewVotingRegistry(s.cfg.NetworkConfig().VotingRegistry, ethClient)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize new voting registry"))
	}

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxNetworkConfig(s.cfg.NetworkConfig()),
			handlers.CtxEthClient(ethClient),
			handlers.CtxVoteVerifierRegisterMethod(&method),
			handlers.CtxVotingRegistry(votingRegistry),
		),
	)
	r.Route("/integrations/proof-verification-relayer", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/verify-proof", handlers.VerifyProof)
		})
	})

	return r
}
