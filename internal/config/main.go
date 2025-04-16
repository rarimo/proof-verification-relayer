package config

import (
	"github.com/rarimo/proof-verification-relayer/internal/pkg/vault"
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
	vault.Vaulter
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
	vault.Vaulter

	pinger     comfig.Once
	replicator comfig.Once
	ipfs       comfig.Once
}

func New(getter kv.Getter) Config {
	vaulter := vault.NewVaulter(getter)
	v := vaulter.Vault()

	networkConfiger := NewNetworkConfiger(getter)
	votingV2Configer := NewVotingV2Configer(getter)

	networkConfiger.(*ethereum).SetVault(v)
	votingV2Configer.(*ethereumVoting).SetVault(v)

	return &config{
		getter:            getter,
		Databaser:         pgdb.NewDatabaser(getter),
		Copuser:           copus.NewCopuser(getter),
		Listenerer:        comfig.NewListenerer(getter),
		Logger:            comfig.NewLogger(getter, comfig.LoggerOpts{}),
		ContractsConfiger: NewContractsConfiger(getter),
		NetworkConfiger:   networkConfiger,
		VotingV2Configer:  votingV2Configer,
		Vaulter:           vaulter,
	}
}
