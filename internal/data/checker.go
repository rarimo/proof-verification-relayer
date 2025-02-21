package data

import (
	"math/big"
	"time"
)

type VotingInfo struct {
	VotingId int64    `db:"voting_id"`
	Balance  *big.Int `db:"residual_balance"`
	GasLimit uint64   `db:"gas_limit"`
}

type ProcessedEvent struct {
	TransactionHash []byte    `db:"transaction_hash"`
	LogIndex        int64     `db:"log_index"`
	CreatedAt       time.Time `db:"created_at"`
	BlockNumber     int64     `db:"block_number"`
}

type CheckerDB interface {
	New() CheckerDB
	CheckerQ() CheckerQ
}

type CheckerQ interface {
	InsertVotingInfo(value VotingInfo) error
	GetVotingInfo(votingId int64) (VotingInfo, error)
	UpdateVotingInfo(value VotingInfo) error

	InsertProcessedEvent(value ProcessedEvent) error
	GetLastBlock() (uint64, error)
	CheckProcessedEventExist(value ProcessedEvent) error
}
