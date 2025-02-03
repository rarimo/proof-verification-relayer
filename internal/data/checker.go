package data

type Wallet struct {
	Address string `db:"address"`
	Balance uint64 `db:"balance"`
}

type CheckerDB interface {
	New() CheckerDB
	CheckerQ() CheckerQ
}

type CheckerQ interface {
	// Select(toids []string, fromids []string, stIndx uint64, pgLen uint64) ([]Wallet, error)
	Insert(value Wallet) error
	GetBalance(address string) (uint64, error)
	UpdateBalance(value Wallet) error
	GetGasLimit() (uint64, error)
	SetGasLimit(value uint64, defaultGasLimit uint64) error
	InsertGasLimit(value uint64) error
	GetBlock() (uint64, error)
	InsertBlock(value uint64) error
	UpdateBlock(value uint64) error
}
