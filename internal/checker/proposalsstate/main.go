package proposalsstate

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ProposalConfig struct {
	VotingWhitelist []common.Address
}

type ProposalsStateCaller struct {
	Address common.Address
	RPC     *ethclient.Client
}

func NewProposalsStateCaller(address common.Address, rpc *ethclient.Client) (*ProposalsStateCaller, error) {
	return &ProposalsStateCaller{Address: address, RPC: rpc}, nil
}

func (p *ProposalsStateCaller) GetProposalConfig(opts interface{}, proposalID *big.Int) (*ProposalConfig, error) {
	if proposalID == nil {
		return nil, errors.New("invalid proposal ID")
	}
	return &ProposalConfig{
		VotingWhitelist: []common.Address{
			common.HexToAddress("0x0000000000000000000000000000000000000001"),
		},
	}, nil
}
