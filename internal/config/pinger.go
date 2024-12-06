package config

import (
	"time"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	pingerYamlKey = "pinger"
)

type Pinger struct {
	BlocksDistance    uint64        `fig:"blocks_distance"`
	Timeout           time.Duration `fig:"timeout"`
	NormalPeriod      time.Duration `fig:"normal_period"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period"`
	MaxAbnormalPeriod time.Duration `fig:"max_abnormal_period"`
}

func (c *config) Pinger() Pinger {
	return c.pinger.Do(func() interface{} {
		var result = Pinger{
			BlocksDistance:    0,
			Timeout:           5 * time.Second,
			NormalPeriod:      30 * time.Second,
			MinAbnormalPeriod: 30 * time.Second,
			MaxAbnormalPeriod: 30 * time.Second,
		}

		err := figure.
			Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, pingerYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out config", logan.F{"key": pingerYamlKey}))
		}

		return result
	}).(Pinger)
}
