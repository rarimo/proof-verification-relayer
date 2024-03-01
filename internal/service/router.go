package service

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
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

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxNetworkConfig(s.cfg.NetworkConfig()),
			handlers.CtxMasterQ(pg.NewMasterQ(s.cfg.DB().Clone())),
			handlers.CtxEthClient(ethClient),
		),
	)
	r.Route("/integrations/proof-verification-relayer", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/verify-proof", handlers.VerifyProof)
		})
	})

	return r
}
