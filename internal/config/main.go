package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	Checker
	NetworkConfiger
	ContractsConfiger
	RelayerConfiger
	Pinger() Pinger
	Replicator() Replicator
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	getter kv.Getter
	Checker
	NetworkConfiger
	ContractsConfiger
	RelayerConfiger

	pinger     comfig.Once
	replicator comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:            getter,
		Databaser:         pgdb.NewDatabaser(getter),
		Copuser:           copus.NewCopuser(getter),
		Listenerer:        comfig.NewListenerer(getter),
		Logger:            comfig.NewLogger(getter, comfig.LoggerOpts{}),
		NetworkConfiger:   NewNetworkConfiger(getter),
		ContractsConfiger: NewContractsConfiger(getter),
		Checker:           NewCheckerCfg(getter),
		RelayerConfiger:   NewRelayerConfiger(getter),
	}
}
