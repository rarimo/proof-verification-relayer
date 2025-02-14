package proposalsstate

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ProposalConfigMock struct {
	VotingWhitelist []common.Address
}

type ProposalsStateCallerMock struct {
	Address common.Address
	RPC     *ethclient.Client
}

func NewProposalsStateCallerMock(address common.Address, rpc *ethclient.Client) (*ProposalsStateCallerMock, error) {
	return &ProposalsStateCallerMock{Address: address, RPC: rpc}, nil
}

func (p *ProposalsStateCallerMock) GetProposalConfig(opts interface{}, proposalID *big.Int) (*ProposalConfigMock, error) {
	if proposalID == nil {
		return nil, errors.New("invalid proposal ID")
	}
	return &ProposalConfigMock{
		VotingWhitelist: []common.Address{
			common.HexToAddress("0x0000000000000000000000000000000000000001"),
		},
	}, nil
}
