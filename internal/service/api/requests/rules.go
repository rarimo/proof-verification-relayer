package requests

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func isAddressRule(value interface{}) error {
	address, _ := value.(string)
	if !common.IsHexAddress(address) {
		return errors.New("should be a valid address")
	}
	return nil
}
