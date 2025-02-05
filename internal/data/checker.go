package data

import "time"

type VotingInfo struct {
	Address  string `db:"voting_address"`
	Balance  uint64 `db:"residual_balance"`
	GasLimit uint64 `db:"gas_limit"`
}

type ProcessedEvent struct {
	TransactionHash []byte    `db:"transaction_hash"`
	LogIndex        int64     `db:"log_index"`
	EmittedAt       time.Time `db:"emitted_at"`
	CreatedAt       time.Time `db:"created_at"`
	BlockNumber     int64     `db:"block_number"`
}

type CheckerDB interface {
	New() CheckerDB
	CheckerQ() CheckerQ
}

type CheckerQ interface {
	InsertVotingInfo(value VotingInfo) error
	GetVotingInfo(address string) (VotingInfo, error)
	UpdateVotingInfo(value VotingInfo) error

	InsertProcessedEvent(value ProcessedEvent) error
	GetLastBlock() (uint64, error)
}
