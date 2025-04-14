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
	NetworkConfiger
	ContractsConfiger
	VotingV2Configer
	Pinger() Pinger
	Replicator() Replicator
	Ipfs() Ipfs
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	getter kv.Getter
	NetworkConfiger
	ContractsConfiger
	VotingV2Configer

	pinger     comfig.Once
	replicator comfig.Once
	ipfs       comfig.Once
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
		VotingV2Configer:  NewVotingV2Configer(getter),
	}
}
