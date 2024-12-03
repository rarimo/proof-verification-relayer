package api

import (
	"context"
	"net"

	"github.com/rarimo/proof-verification-relayer/internal/config"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	cfg      config.Config
}

func (s *service) run(ctx context.Context) {
	s.log.Info("Service started")
	ape.Serve(ctx, s.router(), s.cfg, ape.ServeOpts{})
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log().WithField("service", "api"),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		cfg:      cfg,
	}
}

func Run(ctx context.Context, cfg config.Config) {
	newService(cfg).run(ctx)
}
