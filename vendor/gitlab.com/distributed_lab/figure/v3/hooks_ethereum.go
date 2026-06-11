package figure

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	// EthereumHooks set of default hooks for common types of Ethereum
	EthereumHooks = Hooks{
		"common.Address": func(value interface{}) (reflect.Value, error) {
			switch v := value.(type) {
			case string:
				if !common.IsHexAddress(v) {
					// provide value does not look like valid address
					return reflect.Value{}, errors.New("invalid address")
				}
				return reflect.ValueOf(common.HexToAddress(v)), nil
			default:
				return reflect.Value{}, errors.Errorf("unsupported conversion from %T", value)
			}
		},
		"*common.Address": func(value interface{}) (reflect.Value, error) {
			switch v := value.(type) {
			case string:
				if !common.IsHexAddress(v) {
					// provide value does not look like valid address
					return reflect.Value{}, errors.New("invalid address")
				}
				address := common.HexToAddress(v)
				return reflect.ValueOf(&address), nil
			default:
				return reflect.Value{}, errors.Errorf("unsupported conversion from %T", value)
			}
		},
		"*ecdsa.PrivateKey": func(raw interface{}) (reflect.Value, error) {
			switch value := raw.(type) {
			case string:
				kp, err := crypto.HexToECDSA(value)
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to init keypair")
				}
				return reflect.ValueOf(kp), nil
			default:
				return reflect.Value{}, errors.Errorf("cant init keypair from type: %T", value)
			}
		},
		"*ethclient.Client": func(value interface{}) (reflect.Value, error) {
			switch v := value.(type) {
			case string:
				client, err := ethclient.Dial(v)
				if err != nil {
					return reflect.Value{}, err
				}
				return reflect.ValueOf(client), nil
			default:
				return reflect.Value{}, errors.Errorf("unsupported conversion from %T", value)
			}
		},
	}
)
