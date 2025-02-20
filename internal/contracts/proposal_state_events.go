package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ProposalEvent interface {
	ProposalID() int64
	FundAmount() *big.Int

	BlockNumber() int64
	LogIndex() int64
	Next() bool
	TxHash() common.Hash
}

func (p *ProposalsStateProposalCreatedIterator) ProposalID() int64 {
	return p.Event.ProposalId.Int64()
}

func (p *ProposalsStateProposalCreatedIterator) FundAmount() *big.Int {
	return p.Event.FundAmount
}

func (p *ProposalsStateProposalCreatedIterator) BlockNumber() int64 {
	return int64(p.Event.Raw.BlockNumber)
}

func (p *ProposalsStateProposalCreatedIterator) LogIndex() int64 {
	return int64(p.Event.Raw.Index)
}

func (p *ProposalsStateProposalCreatedIterator) TxHash() common.Hash {
	return p.Event.Raw.TxHash
}

func (p *ProposalsStateProposalFundedIterator) ProposalID() int64 {
	return p.Event.ProposalId.Int64()
}

func (p *ProposalsStateProposalFundedIterator) FundAmount() *big.Int {
	return p.Event.FundAmount
}

func (p *ProposalsStateProposalFundedIterator) BlockNumber() int64 {
	return int64(p.Event.Raw.BlockNumber)
}

func (p *ProposalsStateProposalFundedIterator) LogIndex() int64 {
	return int64(p.Event.Raw.Index)
}

func (p *ProposalsStateProposalFundedIterator) TxHash() common.Hash {
	return p.Event.Raw.TxHash
}
