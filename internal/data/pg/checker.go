package pg

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/rarimo/proof-verification-relayer/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

func NewMaterDB(db *pgdb.DB) data.CheckerDB {
	return &masterDB{
		db: db.Clone(),
	}
}

func (m *masterDB) CheckerQ() data.CheckerQ {
	return NewCheckerQ(m.db)
}

type masterDB struct {
	db *pgdb.DB
}

func (mdb *masterDB) New() data.CheckerDB {
	return NewMaterDB(mdb.db)
}

func NewCheckerQ(db *pgdb.DB) data.CheckerQ {
	return &checkerQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

type checkerQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func (cq *checkerQ) GetVotingInfo(votingId int64) (data.VotingInfo, error) {
	var votingInfo data.VotingInfo
	query := sq.Select("*").From("voting_contract_accounts").Where(sq.Eq{"voting_id": votingId})

	err := cq.db.Get(&votingInfo, query)
	if err == sql.ErrNoRows {
		return votingInfo, err
	}
	if err != nil {
		return votingInfo, errors.Wrap(err, "failed to get voting residual balance from db")
	}
	return votingInfo, nil
}

func (cq *checkerQ) GetBalance(votingId int64) (uint64, error) {
	var balance uint64
	query := sq.Select("residual_balance").From("voting_contract_accounts").Where(sq.Eq{"voting_id": votingId})

	err := cq.db.Get(&balance, query)
	if err == sql.ErrNoRows {
		return 0, err
	}
	if err != nil {
		return 0, errors.Wrap(err, "failed to get voting residual balance from db")
	}
	return balance, nil
}

func (q *checkerQ) InsertVotingInfo(value data.VotingInfo) error {
	query := sq.Insert("voting_contract_accounts").
		Columns("voting_id", "residual_balance", "gas_limit").
		Values(value.VotingId, value.Balance, value.GasLimit)

	err := q.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to insert voting info to db")
	}
	return nil
}

func (q *checkerQ) UpdateVotingInfo(value data.VotingInfo) error {
	query := sq.Update("voting_contract_accounts").
		Set("residual_balance", value.Balance).
		Set("gas_limit", value.GasLimit).
		Where(sq.Eq{"voting_id": value.VotingId})

	err := q.db.Exec(query)
	if err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "failed to update balance to db")
	}
	return nil
}

func (cq *checkerQ) GetLastBlock() (uint64, error) {
	var block uint64
	query := sq.Select("block_number").
		From("processed_events").
		OrderBy("block_number DESC").
		Limit(1)

	err := cq.db.Get(&block, query)
	if err == sql.ErrNoRows {
		return 0, err
	}
	if err != nil {
		return 0, errors.Wrap(err, "failed to get block from db")
	}
	return block, nil
}

func (q *checkerQ) CheckProcessedEventExist(value data.ProcessedEvent) (bool, error) {
	var isExist int
	query := sq.Select("1").From("processed_events").
		Where(sq.Eq{"transaction_hash": value.TransactionHash, "log_index": value.LogIndex}).Limit(1)

	err := q.db.Get(&isExist, query)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, errors.Wrap(err, "failed to check exist event in db")
	}
	return true, nil
}

func (q *checkerQ) InsertProcessedEvent(value data.ProcessedEvent) error {
	query := sq.Insert("processed_events").
		Columns("transaction_hash", "log_index", "block_number").
		Values(value.TransactionHash, value.LogIndex, value.BlockNumber)

	err := q.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to insert processed event to db")
	}
	return nil
}
