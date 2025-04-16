package vault

import (
	"context"
	"fmt"

	vaultapi "github.com/hashicorp/vault/api"

	"gitlab.com/distributed_lab/dig"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Vaulter interface {
	Vault() Vault
}

func NewVaulter(getter kv.Getter) Vaulter {
	return &vaulter{
		getter: getter,
	}
}

type vaulter struct {
	once   comfig.Once
	getter kv.Getter
}

func (c *vaulter) Vault() Vault {
	return c.once.Do(func() interface{} {
		var cfg struct {
			Address   string         `fig:"addr,required"`
			MountPath string         `fig:"mount_path,required"`
			Secrets   map[string]any `fig:"secrets,required"`
			Disabled  bool           `fig:"disabled"`
		}

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "vault")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out vault config: %w", err))
		}

		secrets := make(map[string]*vaultapi.KVSecret, len(cfg.Secrets))

		if cfg.Disabled {
			for secretID, secret := range cfg.Secrets {
				secretMap, ok := secret.(map[string]any)
				if !ok {
					panic(fmt.Errorf("secret %s is not a map[string]any", secretID))
				}

				secrets[secretID] = &vaultapi.KVSecret{
					Data: secretMap,
				}
			}

			return NewVault(secrets)
		}

		token := struct {
			Token string `dig:"VAULT_TOKEN,clear"`
		}{}
		if err = dig.Out(&token).Now(); err != nil {
			panic(fmt.Errorf("failed to dig out vault token: %w", err))
		}

		KVv2, err := getKVv2(cfg.Address, cfg.MountPath, token.Token)
		if err != nil {
			panic(fmt.Errorf("failed to get KVv2: %w", err))
		}

		for secretID, secretValue := range cfg.Secrets {
			secretName, ok := secretValue.(string)
			if !ok {
				panic(fmt.Errorf("secret value %s is not a string", secretID))
			}

			secret, err := KVv2.Get(context.Background(), secretName)
			if err != nil {
				panic(fmt.Errorf("failed to get secret %s from vault: %w", secretName, err))
			}

			secrets[secretID] = secret
		}

		return NewVault(secrets)
	}).(Vault)
}

func getKVv2(vaultAddress, vaultMountPath, token string) (*vaultapi.KVv2, error) {
	conf := vaultapi.DefaultConfig()
	conf.Address = vaultAddress

	vaultClient, err := vaultapi.NewClient(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize new client: %w", err)
	}

	vaultClient.SetToken(token)

	return vaultClient.KVv2(vaultMountPath), nil
}
