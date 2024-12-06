package config

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	contractsYamlKey = "contracts"
	Proposer         = "proposer"
	VotingRegistry   = "voting_registry"
	LightweightState = "lightweight_state"
	Register2        = "register2"
)

var (
	errMissingContract = errors.New("missing contract")
	contracts          = []string{Proposer, VotingRegistry, LightweightState, Register2}
)

type ContractsConfiger interface {
	ContractsConfig() ContractsConfig
}

func NewContractsConfiger(getter kv.Getter) ContractsConfiger {
	return &contractsCfg{
		getter: getter,
	}
}

type contractsCfg struct {
	once   comfig.Once
	getter kv.Getter
}

type ContractConfig struct {
	Address common.Address `fig:"address,required"`
	Block   uint64         `fig:"block"`
}

type ContractsConfig map[string]ContractConfig

func (e *contractsCfg) ContractsConfig() ContractsConfig {
	return e.once.Do(func() interface{} {
		var result ContractsConfig

		err := figure.
			Out(&result).
			With(figure.BaseHooks, figure.EthereumHooks).
			From(kv.MustGetStringMap(e.getter, contractsYamlKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ethereum config", logan.F{"key": contractsYamlKey}))
		}

		for _, contract := range contracts {
			if len(result[contract].Address.Bytes()) == 0 {
				panic(errors.Wrap(errMissingContract, "failed to check contract", logan.F{"key": contract}))
			}
		}

		return result
	}).(ContractsConfig)
}
