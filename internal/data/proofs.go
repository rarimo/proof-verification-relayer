package data

import (
	"github.com/google/uuid"
	"time"
)

type ProofsQ interface {
	New() ProofsQ
	Insert(value Proof) error
	FilterBy(column string, value any) ProofsQ
	OrderBy(column, order string) ProofsQ
	Select() ([]Proof, error)
	DeleteByID(id uuid.UUID) error
	ResetFilter() ProofsQ
}

type Proof struct {
	ID                uuid.UUID `db:"id" structs:"-"`
	Registration      string    `db:"registration" structs:"registration"`
	DocumentNullifier string    `db:"document_nullifier" structs:"document_nullifier"`
	CreatedAt         time.Time `db:"created_at" structs:"-"`
}
