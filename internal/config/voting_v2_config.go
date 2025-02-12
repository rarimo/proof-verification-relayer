package config

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"
	vaultapi "github.com/hashicorp/vault/api"
	"gitlab.com/distributed_lab/dig"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type VotingV2Configer interface {
	VotingV2Config() *VotingV2Config
}

func NewVotingV2Configer(getter kv.Getter) VotingV2Configer {
	return &ethereumVoting{
		getter: getter,
	}
}

type ethereumVoting struct {
	once   comfig.Once
	getter kv.Getter
}

type VotingV2Config struct {
	RPC        *ethclient.Client
	ChainID    *big.Int
	PrivateKey *ecdsa.PrivateKey
	nonce      uint64
	Address    common.Address
	mut        *sync.Mutex
	GasLimit   uint64
	Block      uint64
	Enable     bool
	WithSub    bool
}

func (e *ethereumVoting) VotingV2Config() *VotingV2Config {
	return e.once.Do(func() interface{} {

		raw := kv.MustGetStringMap(e.getter, "voting_v2")

		var probe struct {
			Enable bool `fig:"enable"`
		}

		if err := figure.Out(&probe).From(raw).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out ethereum probe"))
		}

		if !probe.Enable {
			return &VotingV2Config{
				Enable: false,
			}
		}

		var result VotingV2Config

		networkConfig := struct {
			RPC            *ethclient.Client `fig:"rpc,required"`
			PrivateKey     *ecdsa.PrivateKey `fig:"private_key"`
			VaultAddress   string            `fig:"vault_address"`
			VaultMountPath string            `fig:"vault_mount_path"`
			Address        common.Address    `fig:"proposal_state_address,required"`
			Block          uint64            `fig:"block"`
			GasLimit       uint64            `fig:"gas_limit"`
			Enable         bool              `fig:"enable"`
			WithSub        bool              `fig:"check_with_subscribe"`
		}{}
		err := figure.
			Out(&networkConfig).
			With(figure.EthereumHooks).
			From(raw).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ethereum config"))
		}

		result.RPC = networkConfig.RPC
		result.Enable = networkConfig.Enable
		result.ChainID, err = result.RPC.ChainID(context.Background())
		result.WithSub = networkConfig.WithSub
		if err != nil {
			panic(errors.Wrap(err, "failed to get chain ID"))
		}

		result.PrivateKey = networkConfig.PrivateKey
		if result.PrivateKey == nil {
			result.PrivateKey = extractPrivateKey(networkConfig.VaultAddress, networkConfig.VaultMountPath)
		}

		result.nonce, err = result.RPC.NonceAt(context.Background(), crypto.PubkeyToAddress(result.PrivateKey.PublicKey), nil)
		if err != nil {
			panic(errors.Wrap(err, "failed to get nonce"))
		}
		result.Address = networkConfig.Address
		result.mut = &sync.Mutex{}
		result.Block = networkConfig.Block
		result.GasLimit = networkConfig.GasLimit

		return &result
	}).(*VotingV2Config)
}

func (n *VotingV2Config) LockNonce() {
	n.mut.Lock()
}

func (n *VotingV2Config) UnlockNonce() {
	n.mut.Unlock()
}

func (n *VotingV2Config) Nonce() uint64 {
	return n.nonce
}

func (n *VotingV2Config) IncrementNonce() {
	n.nonce++
}

// ResetNonce sets nonce to the value received from a node
func (n *VotingV2Config) ResetNonce(client *ethclient.Client) error {
	nonce, err := client.NonceAt(context.Background(), crypto.PubkeyToAddress(n.PrivateKey.PublicKey), nil)
	if err != nil {
		return errors.Wrap(err, "failed to get nonce")
	}
	n.nonce = nonce
	return nil
}

func extractPrivateKey(vaultAddress, vaultMountPath string) *ecdsa.PrivateKey {
	conf := vaultapi.DefaultConfig()
	conf.Address = vaultAddress

	vaultClient, err := vaultapi.NewClient(conf)
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize new client"))
	}

	token := struct {
		Token string `dig:"VAULT_TOKEN,clear"`
	}{}

	err = dig.Out(&token).Now()
	if err != nil {
		panic(errors.Wrap(err, "failed to dig out token"))
	}

	vaultClient.SetToken(token.Token)

	secret, err := vaultClient.KVv2(vaultMountPath).Get(context.Background(), "relayer")
	if err != nil {
		panic(errors.Wrap(err, "failed to get secret"))
	}

	vaultRelayerConf := struct {
		PrivateKey *ecdsa.PrivateKey `fig:"private_key,required"`
	}{}

	if err := figure.
		Out(&vaultRelayerConf).
		With(figure.EthereumHooks).
		From(secret.Data).
		Please(); err != nil {
		panic(errors.Wrap(err, "failed to figure out"))
	}

	return vaultRelayerConf.PrivateKey
}
