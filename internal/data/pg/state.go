package pg

import (
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const stateTableName = "state"

func NewStateQ(db *pgdb.DB) data.StateQ {
	return &stateQ{
		db:       db,
		selector: sq.Select("*").From(stateTableName),
		updater:  sq.Update(stateTableName),
		deleter:  sq.Delete(stateTableName),
	}
}

type stateQ struct {
	db       *pgdb.DB
	selector sq.SelectBuilder
	updater  sq.UpdateBuilder
	deleter  sq.DeleteBuilder
}

func (q stateQ) New() data.StateQ {
	return NewStateQ(q.db.Clone())
}

func (q stateQ) Get() (*data.State, error) {
	var result data.State
	err := q.db.Get(&result, q.selector)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}

func (q stateQ) Upsert(value data.State) error {
	return q.db.Exec(sq.Insert(stateTableName).SetMap(structs.Map(value)).Suffix("ON CONFLICT (root, tx_hash) DO NOTHING"))
}

func (q stateQ) FilterByRoot(root ...string) data.StateQ {
	return q.withFilters(sq.Eq{"root": root})
}

func (q stateQ) FilterByBlock(block ...uint64) data.StateQ {
	return q.withFilters(sq.Eq{"block": block})
}

func (q stateQ) SortByBlockHeight(order data.SortOrder) data.StateQ {
	q.selector = q.selector.OrderBy(fmt.Sprintf("block %s", order))
	q.updater = q.updater.OrderBy(fmt.Sprintf("block %s", order))
	q.deleter = q.deleter.OrderBy(fmt.Sprintf("block %s", order))

	return q
}

func (q stateQ) withFilters(stmt interface{}) data.StateQ {
	q.selector = q.selector.Where(stmt)
	q.updater = q.updater.Where(stmt)
	q.deleter = q.deleter.Where(stmt)

	return q
}
