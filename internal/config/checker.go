package config

import (
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Checker interface {
	NewCheckerCfg() configChecker
}

type checkerCfg struct {
	getter kv.Getter
	once   comfig.Once
}

func NewCheckerCfg(getter kv.Getter) Checker {
	return &checkerCfg{
		getter: getter,
	}
}

type configChecker struct {
	RpcURL       string `fig:"rpc"`
	ContractAddr string `fig:"contract_addr"`
	StartBlock   uint64 `fig:"block"`
	GasLimit     uint64 `fig:"gas_limit"`
}

func (c *checkerCfg) NewCheckerCfg() configChecker {
	return c.once.Do(func() interface{} {

		var csparam configChecker

		if err := figure.Out(&csparam).From(kv.MustGetStringMap(kv.MustFromEnv(), "checker")).Please(); err != nil {
			panic(err)
		}

		return csparam
	}).(configChecker)
}
