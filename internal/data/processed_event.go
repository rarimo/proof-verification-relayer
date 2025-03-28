package data

import "time"

type ProcessedEvent struct {
	TransactionHash []byte    `db:"transaction_hash"`
	LogIndex        int64     `db:"log_index"`
	CreatedAt       time.Time `db:"created_at"`
	BlockNumber     int64     `db:"block_number"`
}

type ProcessedEventQ interface {
	New() ProcessedEventQ

	InsertProcessedEvent(value ProcessedEvent) error

	GetLastBlock() (uint64, error)
	CheckProcessedEventExist(value ProcessedEvent) (bool, error)
}
