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
	Url         string        `fig:"url"`
	RetryPeriod time.Duration `fig:"retry_period"`
	MaxRetries  uint64        `fig:"max_retries"`
}

func (c *config) Ipfs() Ipfs {
	return c.ipfs.Do(func() interface{} {
		var result Ipfs

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
