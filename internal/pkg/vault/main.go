package vault

import (
	"context"
	"fmt"

	vaultapi "github.com/hashicorp/vault/api"

	"gitlab.com/distributed_lab/dig"
	"gitlab.com/distributed_lab/figure/v3"
)

func NewVault(secrets map[string]*vaultapi.KVSecret) Vault {
	return &vault{
		secrets: secrets,
	}
}

type vault struct {
	secrets map[string]*vaultapi.KVSecret
}

func (v *vault) GetSecret(secretName string, clearSecret bool) (*vaultapi.KVSecret, error) {
	secret, ok := v.secrets[secretName]
	if !ok {
		return nil, ErrSecretNotFound
	}

	if clearSecret {
		delete(v.secrets, secretName)
	}

	return secret, nil
}

func (v *vault) GetSecretData(secretName string, clearSecret bool) (map[string]interface{}, error) {
	secret, ok := v.secrets[secretName]
	if !ok {
		return nil, ErrSecretNotFound
	}

	if clearSecret {
		delete(v.secrets, secretName)
	}

	return secret.Data, nil
}

func (v *vault) FigureOutSecret(secretName string, dst any, clearSecret bool) error {
	secret, ok := v.secrets[secretName]
	if !ok {
		return ErrSecretNotFound
	}

	if clearSecret {
		delete(v.secrets, secretName)
	}

	if err := figure.
		Out(dst).
		With(figure.EthereumHooks).
		From(secret.Data).
		Please(); err != nil {
		return fmt.Errorf("failed to figure out secret: %w", err)
	}

	return nil
}

func ExtractSecret(vaultAddress, vaultMountPath, secretName string, dst any) error {
	conf := vaultapi.DefaultConfig()
	conf.Address = vaultAddress

	vaultClient, err := vaultapi.NewClient(conf)
	if err != nil {
		panic(fmt.Errorf("failed to initialize new client: %w", err))
	}

	token := struct {
		Token string `dig:"VAULT_TOKEN"`
	}{}

	err = dig.Out(&token).Now()
	if err != nil {
		return fmt.Errorf("failed to dig out token: %w", err)
	}

	vaultClient.SetToken(token.Token)

	secret, err := vaultClient.KVv2(vaultMountPath).Get(context.Background(), secretName)
	if err != nil {
		return fmt.Errorf("failed to get secret: %w", err)
	}

	if err := figure.
		Out(dst).
		With(figure.EthereumHooks).
		From(secret.Data).
		Please(); err != nil {
		return fmt.Errorf("failed to figure out secret: %w", err)
	}

	return nil
}
