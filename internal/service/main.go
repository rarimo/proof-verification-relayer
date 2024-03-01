package service

import (
	"context"
	"net"
	"net/http"

	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"github.com/rarimo/proof-verification-relayer/internal/service/proofs_sender"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	cfg      config.Config
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	proofsSender, err := proofs_sender.NewProofsSender(s.log, pg.NewMasterQ(s.cfg.DB().Clone()), s.cfg.NetworkConfig())
	if err != nil {
		return errors.Wrap(err, "failed to initialize new proofs sender")
	}
	go proofsSender.Run(context.Background())

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		cfg:      cfg,
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
