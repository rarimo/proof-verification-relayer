package pg

import (
	"database/sql"
	"encoding/json"
	"math/big"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type votingInf struct {
	VotingId               int64  `db:"voting_id"`
	Balance                string `db:"residual_balance"`
	GasLimit               uint64 `db:"gas_limit"`
	CreatorAddress         string `db:"creator_address"`
	ProposalInfoWithConfig string `db:"proposal_info_with_config"`
}

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

func (cq *checkerQ) GetVotingInfo(votingId int64) (*data.VotingInfo, error) {
	var votingInfo votingInf

	query := sq.Select("*").From("voting_contract_accounts").Where(sq.Eq{"voting_id": votingId})

	err := cq.db.Get(&votingInfo, query)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "failed to get voting info from db")
		}
		return nil, nil
	}

	var proposalInfo resources.VotingInfoAttributes
	err = json.Unmarshal([]byte(votingInfo.ProposalInfoWithConfig), &proposalInfo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ProposalInfoWithConfig")
	}

	balance, success := new(big.Int).SetString(votingInfo.Balance, 10)
	if !success {
		return nil, errors.New("error converting string balance to big.Int")
	}

	return &data.VotingInfo{
		GasLimit:               votingInfo.GasLimit,
		VotingId:               votingInfo.VotingId,
		Balance:                balance,
		CreatorAddress:         votingInfo.CreatorAddress,
		ProposalInfoWithConfig: proposalInfo,
	}, nil
}

func (cq *checkerQ) SelectVotes() ([]*data.VotingInfo, error) {
	var votingInfoList []votingInf

	query := sq.Select("*").From("voting_contract_accounts")

	err := cq.db.Select(&votingInfoList, query)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "failed to select all votes from db")
		}
		return nil, nil
	}

	var result []*data.VotingInfo
	for _, vote := range votingInfoList {
		balance, success := new(big.Int).SetString(vote.Balance, 10)
		if !success {
			continue
		}
		var proposalInfo resources.VotingInfoAttributes
		err = json.Unmarshal([]byte(vote.ProposalInfoWithConfig), &proposalInfo)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal ProposalInfoWithConfig")
		}
		result = append(result, &data.VotingInfo{
			GasLimit:               vote.GasLimit,
			VotingId:               vote.VotingId,
			Balance:                balance,
			CreatorAddress:         vote.CreatorAddress,
			ProposalInfoWithConfig: proposalInfo,
		})
	}

	return result, nil
}

func (q *checkerQ) InsertVotingInfo(value *data.VotingInfo) error {
	jsonData, err := json.Marshal(value.ProposalInfoWithConfig)
	if err != nil {
		return errors.Wrap(err, "failed to marshal ProposalInfoWithConfig")
	}

	query := sq.Insert("voting_contract_accounts").
		Columns("voting_id", "residual_balance", "gas_limit", "creator_address", "proposal_info_with_config").
		Values(value.VotingId, value.Balance.String(), value.GasLimit, value.CreatorAddress, jsonData)

	err = q.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to insert voting info to db")
	}
	return nil
}

func (q *checkerQ) UpdateVotingInfo(value *data.VotingInfo) error {
	jsonData, err := json.Marshal(value.ProposalInfoWithConfig)
	if err != nil {
		return errors.Wrap(err, "failed to marshal ProposalInfoWithConfig")
	}

	query := sq.Update("voting_contract_accounts").
		Set("residual_balance", value.Balance.String()).
		Set("gas_limit", value.GasLimit).
		Set("proposal_info_with_config", jsonData).
		Where(sq.Eq{"voting_id": value.VotingId})

	err = q.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to update voting info in db")
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
	if err != nil {
		return 0, err
	}
	return block, nil
}

func (q *checkerQ) CheckProcessedEventExist(value data.ProcessedEvent) (bool, error) {
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
