package config

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	vaultapi "github.com/hashicorp/vault/api"
	"gitlab.com/distributed_lab/dig"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type NetworkConfiger interface {
	NetworkConfig() *NetworkConfig
}

func NewNetworkConfiger(getter kv.Getter) NetworkConfiger {
	return &ethereum{
		getter: getter,
	}
}

type ethereum struct {
	once   comfig.Once
	getter kv.Getter
}

type NetworkConfig struct {
	Client        *ethclient.Client `fig:"rpc,required"`
	Address       string            `fig:"vault_address,required"`
	MountPath     string            `fig:"vault_mount_path,required"`
	GasMultiplier float64           `fig:"gas_multiplier,required"`

	ChainID    *big.Int          `fig:"chain_id"`
	Token      string            `dig:"VAULT_TOKEN,clear"`
	PrivateKey *ecdsa.PrivateKey `fig:"private_key"`
	nonce      uint64

	mut *sync.Mutex
}

func (e *ethereum) NetworkConfig() *NetworkConfig {
	return e.once.Do(func() interface{} {
		var result NetworkConfig

		err := figure.
			Out(&result).
			With(figure.BaseHooks, figure.EthereumHooks).
			From(kv.MustGetStringMap(e.getter, "network")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ethereum config"))
		}

		chainID, err := result.Client.ChainID(context.Background())
		if err != nil {
			panic(errors.Wrap(err, "failed to get chain ID"))
		}

		result.ChainID = chainID

		if result.PrivateKey == nil {
			result.PrivateKey, err = retrieveVaultPrivateKey(result)
			if err != nil {
				panic(errors.Wrap(err, "failed to retrieve vault private key"))
			}
		}

		nonce, err := result.Client.NonceAt(context.Background(), crypto.PubkeyToAddress(result.PrivateKey.PublicKey), nil)
		if err != nil {
			panic(errors.Wrap(err, "failed to get nonce"))
		}

		result.nonce = nonce

		result.mut = &sync.Mutex{}

		return &result
	}).(*NetworkConfig)
}

func retrieveVaultPrivateKey(result NetworkConfig) (*ecdsa.PrivateKey, error) {
	if err := dig.Out(&result).
		Where(map[string]interface{}{
			"vault_address":    result.Address,
			"vault_mount_path": result.MountPath,
			"gas_multiplier":   result.GasMultiplier,
		}).Now(); err != nil {
		panic(err)
	}

	conf := vaultapi.DefaultConfig()
	conf.Address = result.Address

	vaultClient, err := vaultapi.NewClient(conf)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize new client"))
	}

	vaultClient.SetToken(result.Token)

	secret, err := vaultClient.KVv2(result.MountPath).Get(context.Background(), "relayer")
	if err != nil {
		panic(errors.Wrap(err, "failed to get secret"))
	}

	vaultRelayerConf := struct {
		PrivateKey *ecdsa.PrivateKey `fig:"private_key,required"`
	}{}

	if err := figure.
		Out(&vaultRelayerConf).
		With(figure.BaseHooks, figure.EthereumHooks).
		From(secret.Data).
		Please(); err != nil {
		panic(errors.Wrap(err, "failed to figure out"))
	}

	return vaultRelayerConf.PrivateKey, nil
}

func (n *NetworkConfig) LockNonce() {
	n.mut.Lock()
}

func (n *NetworkConfig) UnlockNonce() {
	n.mut.Unlock()
}

func (n *NetworkConfig) Nonce() uint64 {
	return n.nonce
}

func (n *NetworkConfig) IncrementNonce() {
	n.nonce++
}

// ResetNonce sets nonce to the value received from a node
func (n *NetworkConfig) ResetNonce(client *ethclient.Client) error {
	nonce, err := client.NonceAt(context.Background(), crypto.PubkeyToAddress(n.PrivateKey.PublicKey), nil)
	if err != nil {
		return errors.Wrap(err, "failed to get nonce")
	}
	n.nonce = nonce
	return nil
}
