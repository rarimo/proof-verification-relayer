package pg

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type processedEventsQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func (q *processedEventsQ) New() data.ProcessedEventQ {
	return NewProcessedEventsQ(q.db.Clone())
}

func NewProcessedEventsQ(db *pgdb.DB) data.ProcessedEventQ {
	return &processedEventsQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

func (cq *processedEventsQ) GetLastBlock() (uint64, error) {
	var block uint64
	query := sq.Select("block_number").
		From("processed_events").
		OrderBy("block_number DESC").
		Limit(1)

	err := cq.db.Get(&block, query)
	if err != nil {
		return 0, err
	}
	return block, nil
}

func (q *processedEventsQ) CheckProcessedEventExist(value data.ProcessedEvent) (bool, error) {
	var isExist bool
	query := sq.Select("1").From("processed_events").
		Where(sq.Eq{"transaction_hash": value.TransactionHash, "log_index": value.LogIndex}).Limit(1)

	err := q.db.Get(&isExist, query)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, errors.Wrap(err, "failed to check exist event in db")
		}

		return false, nil
	}
	return true, nil
}

func (q *processedEventsQ) InsertProcessedEvent(value data.ProcessedEvent) error {
	query := sq.Insert("processed_events").
		Columns("transaction_hash", "log_index", "block_number").
		Values(value.TransactionHash, value.LogIndex, value.BlockNumber)

	err := q.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to insert processed event to db")
	}
	return nil
}
