package data

import (
	"github.com/google/uuid"
	"time"
)

type ProofsQ interface {
	New() ProofsQ
	Insert(value Proof) error
	OrderBy(column, order string) ProofsQ
	Select() ([]Proof, error)
	DeleteByID(id uuid.UUID) error
	ResetFilter() ProofsQ
}

type Proof struct { // TODO unique(voting session, nullifier)
	ID        uuid.UUID `db:"id" structs:"-"`
	TxData    string    `db:"tx_data" structs:"tx_data"`
	CreatedAt time.Time `db:"created_at" structs:"-"`
}
