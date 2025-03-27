package pg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
)

var (
	votingTableName = "voting_contract_accounts"
)

type votingInf struct {
	VotingId               int64  `db:"voting_id"`
	Balance                string `db:"residual_balance"`
	GasLimit               uint64 `db:"gas_limit"`
	CreatorAddress         string `db:"creator_address"`
	ProposalInfoWithConfig string `db:"proposal_info_with_config"`
}

type votingQ struct {
	db       *pgdb.DB
	sql      sq.StatementBuilderType
	selector sq.SelectBuilder
	updater  sq.UpdateBuilder
}

func (q votingQ) New() data.VotingQ {
	return NewVotingQ(q.db.Clone())
}

func NewVotingQ(db *pgdb.DB) data.VotingQ {
	return &votingQ{
		db:       db,
		sql:      sq.StatementBuilder,
		selector: sq.Select("*").From(votingTableName),
		updater:  sq.Update(votingTableName),
	}
}

func (q votingQ) Get(column string, dest interface{}) error {
	query := q.selector.RemoveColumns().Column(column)

	if err := q.db.Get(dest, query); err != nil {
		if err != sql.ErrNoRows {
			return errors.Wrap(err, "failed to get voting info from db")
		}
		return err
	}
	return nil
}

func (q votingQ) GetVotingBalance() (*big.Int, error) {
	var balanceStr string

	query := q.selector.RemoveColumns().Column("residual_balance")
	if err := q.db.Get(&balanceStr, query); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "failed to get voting info from db")
		}
		return nil, nil
	}

	balance, success := new(big.Int).SetString(balanceStr, 10)
	if !success {
		return nil, errors.New("error converting string balance to big.Int")
	}

	return balance, nil
}

func (q votingQ) GetVotingInfo(votingId int64) (*data.VotingInfo, error) {
	var votingInfo votingInf

	query := sq.Select("*").From(votingTableName).Where(sq.Eq{"voting_id": votingId})

	err := q.db.Get(&votingInfo, query)
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

func (q votingQ) SelectVotingInfo() ([]*data.VotingInfo, error) {
	var votingInfoList []votingInf

	err := q.db.Select(&votingInfoList, q.selector)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "failed to select voting info from db")
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

func (q votingQ) InsertVotingInfo(value *data.VotingInfo) error {
	jsonDataProposalInfo, err := json.Marshal(value.ProposalInfoWithConfig)
	if err != nil {
		return errors.Wrap(err, "failed to marshal ProposalInfoWithConfig")
	}

	query := sq.Insert(votingTableName).
		Columns("voting_id", "residual_balance", "gas_limit", "creator_address", "proposal_info_with_config").
		Values(value.VotingId, value.Balance.String(), value.GasLimit, value.CreatorAddress, jsonDataProposalInfo)

	err = q.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to insert voting info to db")
	}
	return nil
}

func (q votingQ) UpdateVotingInfo(value *data.VotingInfo) error {
	jsonDataProposalInfo, err := json.Marshal(value.ProposalInfoWithConfig)
	if err != nil {
		return errors.Wrap(err, "failed to marshal ProposalInfoWithConfig")
	}
	var query sq.UpdateBuilder
	query = sq.Update(votingTableName)
	columns := []string{"voting_id", "residual_balance", "gas_limit", "creator_address", "proposal_info_with_config"}
	values := []any{value.VotingId, value.Balance.String(), value.GasLimit, value.CreatorAddress, jsonDataProposalInfo}

	for indx, column := range columns {
		query = query.Set(column, values[indx])
	}

	query = query.Where(sq.Eq{"voting_id": value.VotingId})

	err = q.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to update voting info in db")
	}
	return nil
}

func (q votingQ) Update(fields map[string]any) error {
	if err := q.db.Exec(q.updater.SetMap(fields)); err != nil {
		return fmt.Errorf("failed to update voting info: %w", err)
	}

	return nil
}

func (q votingQ) FilterByVotingId(votingIds ...int64) data.VotingQ {
	if len(votingIds) == 0 {
		return q
	}
	return q.withFilters(sq.Eq{"voting_id": votingIds})
}

func (q votingQ) FilterByCreator(creators ...string) data.VotingQ {
	if len(creators) == 0 {
		return q
	}
	return q.withFilters(sq.Eq{"creator_address": creators})
}

func (q votingQ) FilterByMinAge(minAgeList ...int64) data.VotingQ {
	if len(minAgeList) == 0 {
		return q
	}

	var stmt []string
	for _, minAge := range minAgeList {
		stmt = append(stmt, fmt.Sprintf("proposal_info_with_config @> '{\"contract\": {\"config\": {\"parsed_voting_whitelist_data\": [{\"min_age\": %d}]}}}'", minAge))
	}

	return q.withFilters(fmt.Sprintf("(%v)", strings.Join(stmt, " OR ")))
}

func (q votingQ) FilterByCitizenship(сitizenshipList ...string) data.VotingQ {
	if len(сitizenshipList) == 0 {
		return q
	}
	var stmt []string
	for _, citizenship := range сitizenshipList {
		stmt = append(stmt, fmt.Sprintf("proposal_info_with_config @> '{\"contract\": {\"config\": {\"parsed_voting_whitelist_data\": [{\"citizenship_whitelist\": [\"%v\"]}]}}}'", citizenship))
	}

	return q.withFilters(fmt.Sprintf("(%v)", strings.Join(stmt, " OR ")))
}

func (q votingQ) withFilters(stmt interface{}) data.VotingQ {
	q.selector = q.selector.Where(stmt)
	q.updater = q.updater.Where(stmt)

	return q
}
