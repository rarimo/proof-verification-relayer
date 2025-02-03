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
	return &transactionQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

type transactionQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func (cq *transactionQ) GetBalance(address string) (uint64, error) {

	var balance uint64

	query := sq.Select("balance").From("wallets").Where(sq.Eq{"address": address})

	err := cq.db.Get(&balance, query)

	if err == sql.ErrNoRows {
		return 0, err
	}
	if err != nil {
		return 0, errors.Wrap(err, "failed to get wallet balance from db")
	}
	return balance, nil

}

func (q *transactionQ) Insert(value data.Wallet) error {

	query := sq.Insert("wallets").Columns("address", "balance").Values(value.Address, value.Balance)

	err := q.db.Exec(query)
	if err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "failed to insert wallet to db")
	}
	return nil

}

func (q *transactionQ) UpdateBalance(value data.Wallet) error {

	query := sq.Update("wallets").Set("balance", value.Balance).Where(sq.Eq{"address": value.Address})

	err := q.db.Exec(query)
	if err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "failed to update balance to db")
	}
	return nil

}

func (cq *transactionQ) GetBlock() (uint64, error) {

	var block uint64

	query := sq.Select("block").From("blocks").Limit(1)

	err := cq.db.Get(&block, query)

	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, errors.Wrap(err, "failed to get block from db")
	}

	return block, nil

}

func (q *transactionQ) InsertBlock(value uint64) error {

	query := sq.Insert("blocks").Columns("block").Values(value)

	err := q.db.Exec(query)
	if err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "failed to insert block to db")
	}
	return nil

}

func (q *transactionQ) UpdateBlock(value uint64) error {

	query := sq.Update("blocks").Set("block", value)

	err := q.db.Exec(query)
	if err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "failed to update block to db")
	}
	return nil

}

func (cq *transactionQ) GetGasLimit() (uint64, error) {

	var gas_limit uint64

	query := sq.Select("gas_limit").From("gas_limits").Limit(1)

	err := cq.db.Get(&gas_limit, query)

	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, errors.Wrap(err, "failed to get gas_limit from db")
	}
	return gas_limit, nil

}

func (q *transactionQ) InsertGasLimit(value uint64) error {

	query := sq.Insert("gas_limits").Columns("gas_limit").Values(value)

	err := q.db.Exec(query)
	if err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "failed to insert gas_limit to db")
	}
	return nil

}

func (q *transactionQ) SetGasLimit(value uint64, defaultGasLimit uint64) error {

	query := sq.Update("gas_limits").Set("gas_limit", value).Where(sq.Eq{"gas_limit": defaultGasLimit})

	err := q.db.Exec(query)
	if err != nil && err != sql.ErrNoRows {
		return errors.Wrap(err, "failed to update gas_limit to db")
	}
	return nil

}
