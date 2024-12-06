package config

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	replicatorYamlKey = "replicator"
)

type Replicator struct {
	Address    common.Address `fig:"address,required"`
	SourceSMT  common.Address `fig:"source_smt,required"`
	RootPrefix string         `fig:"root_prefix,required"`
}

func (c *config) Replicator() Replicator {
	return c.replicator.Do(func() interface{} {
		var result Replicator

		err := figure.
			Out(&result).
			With(figure.BaseHooks, figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, replicatorYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out config", logan.F{"key": replicatorYamlKey}))
		}

		return result
	}).(Replicator)
}
