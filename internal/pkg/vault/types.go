package vault

import (
	"errors"

	vaultapi "github.com/hashicorp/vault/api"
)

type Vault interface {
	GetSecret(secretName string, clearSecret bool) (*vaultapi.KVSecret, error)
	GetSecretData(secretName string, clearSecret bool) (map[string]interface{}, error)
	FigureOutSecret(secretName string, dst any, clearSecret bool) error
}

var (
	ErrSecretNotFound = errors.New("secret not found")
)
