package config

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	vaultapi "github.com/hashicorp/vault/api"
	"gitlab.com/distributed_lab/dig"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
	"sync"
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
	RPC            string         `fig:"rpc,required"`
	Proposer       common.Address `fig:"proposer,required"`
	VotingRegistry common.Address `fig:"voting_registry,required"`
	Address        string         `fig:"vault_address,required"`
	MountPath      string         `fig:"vault_mount_path,required"`

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

		ethClient, err := ethclient.Dial(result.RPC)
		if err != nil {
			panic(errors.Wrap(err, "failed to create connection"))
		}
		defer ethClient.Close()

		chainID, err := ethClient.ChainID(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed to get chain ID")
		}

		result.ChainID = chainID

		if err := dig.Out(&result).
			Where(map[string]interface{}{
				"rpc":              result.RPC,
				"proposer":         result.Proposer,
				"voting_registry":  result.VotingRegistry,
				"vault_address":    result.Address,
				"vault_mount_path": result.MountPath,
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

		result.PrivateKey = vaultRelayerConf.PrivateKey

		nonce, err := ethClient.NonceAt(context.Background(), crypto.PubkeyToAddress(result.PrivateKey.PublicKey), nil)
		if err != nil {
			panic(errors.Wrap(err, "failed to get nonce"))
		}

		result.nonce = nonce

		result.mut = &sync.Mutex{}

		return &result
	}).(*NetworkConfig)
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
