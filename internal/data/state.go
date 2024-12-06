package data

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
	ID        string `db:"id" structs:"-"`
	Root      string `db:"root" structs:"root"`
	TxHash    string `db:"tx_hash" structs:"tx_hash"`
	Block     uint64 `db:"block" structs:"block"`
	Timestamp uint64 `db:"timestamp" structs:"timestamp"`
}

const (
	ASC  SortOrder = "ASC"
	DESC SortOrder = "DESC"
)
