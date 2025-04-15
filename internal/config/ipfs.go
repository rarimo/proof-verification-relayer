package config

import (
	"time"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	ipfsYamlKey = "ipfs"
)

type Ipfs struct {
	Url               string        `fig:"url"`
	MaxRetries        uint64        `fig:"max_retries"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period"`
	MaxAbnormalPeriod time.Duration `fig:"max_abnormal_period"`
}

func (c *config) Ipfs() Ipfs {
	return c.ipfs.Do(func() interface{} {
		result := Ipfs{
			Url:               "https://ipfs.rarimo.com/ipfs",
			MaxRetries:        5,
			MinAbnormalPeriod: 30 * time.Second,
			MaxAbnormalPeriod: 30 * time.Second,
		}

		err := figure.
			Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, ipfsYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out config", logan.F{"key": ipfsYamlKey}))
		}

		return result
	}).(Ipfs)
}
