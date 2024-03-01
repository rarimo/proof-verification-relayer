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

type Proof struct {
	ID                uuid.UUID `db:"id" structs:"-"`
	VotingSession     string    `db:"voting_session" structs:"voting_session"`
	DocumentNullifier string    `db:"document_nullifier" structs:"document_nullifier"`
	CreatedAt         time.Time `db:"created_at" structs:"-"`
}
