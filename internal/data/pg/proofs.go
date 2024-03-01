package pg

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const claimsTableName = "proofs"

var (
	claimsSelector = sq.Select("*").From(claimsTableName)
	claimsUpdate   = sq.Update(claimsTableName)
)

func NewProofsQ(db *pgdb.DB) data.ProofsQ {
	return &proofsQ{
		db:  db,
		sql: claimsSelector,
		upd: claimsUpdate,
	}
}

type proofsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
	upd sq.UpdateBuilder
}

func (q *proofsQ) New() data.ProofsQ {
	return NewProofsQ(q.db.Clone())
}

func (q *proofsQ) Insert(value data.Proof) error {
	clauses := structs.Map(value)
	stmt := sq.Insert(claimsTableName).SetMap(clauses)
	err := q.db.Exec(stmt)
	return err
}

func (q *proofsQ) FilterBy(column string, value any) data.ProofsQ {
	q.sql = q.sql.Where(sq.Eq{column: value})
	return q
}

func (q *proofsQ) OrderBy(column, order string) data.ProofsQ {
	q.sql = q.sql.OrderBy(fmt.Sprintf("%s %s", column, order))
	return q
}

func (q *proofsQ) Get() (*data.Proof, error) {
	var result data.Proof
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &result, err
}

func (q *proofsQ) Select() ([]data.Proof, error) {
	var result []data.Proof
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *proofsQ) DeleteByID(id uuid.UUID) error {
	if err := q.db.Exec(sq.Delete(claimsTableName).Where(sq.Eq{"id": id})); err != nil {
		return err
	}
	return nil
}

func (q *proofsQ) ForUpdate() data.ProofsQ {
	q.sql = q.sql.Suffix("FOR UPDATE")
	return q
}

func (q *proofsQ) ResetFilter() data.ProofsQ {
	q.sql = sq.Select("*").From(claimsTableName)
	q.upd = sq.Update(claimsTableName)
	return q
}
