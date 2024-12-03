package data

import "time"

type SortOrder string

type StateQ interface {
	New() StateQ

	Get() (*State, error)
	Upsert(data State) error

	FilterByRoot(root ...string) StateQ
	FilterByBlock(block ...uint64) StateQ

	SortByBlockHeight(order SortOrder) StateQ
}

type State struct {
	ID        int64     `db:"id" structs:"-"`
	Root      string    `db:"root" structs:"root"`
	TxHash    string    `db:"tx_hash" structs:"tx_hash"`
	Block     uint64    `db:"block" structs:"block"`
	CreatedAt time.Time `db:"created_at" structs:"created_at"`
}

const (
	ASC  SortOrder = "ASC"
	DESC SortOrder = "DESC"
)
