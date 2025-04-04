package pg

import (
	"database/sql"
	"encoding/hex"
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
	votingClolumns  = []string{
		"voting_id",
		"residual_balance",
		"gas_limit",
		"creator_address",
		"parsed_whitelist_data_with_metadata",
		"total_balance",
		"min_age",
		"max_age",
		"start_timestamp",
		"end_timestamp",
		"votes_count",
	}

	// CASE to determine the current voting status
	statusCase = fmt.Sprintf(` CASE
	WHEN EXTRACT(EPOCH FROM NOW()) < start_timestamp THEN '%s'
	WHEN EXTRACT(EPOCH FROM NOW()) BETWEEN start_timestamp AND end_timestamp THEN '%s'
	WHEN EXTRACT(EPOCH FROM NOW()) > end_timestamp THEN '%s'
END`,
		data.StatusWaiting, data.StatusStarted, data.StatusEnded,
	)
)

type votingInf struct {
	VotingId                        int64  `db:"voting_id"`
	Balance                         string `db:"residual_balance"`
	GasLimit                        uint64 `db:"gas_limit"`
	CreatorAddress                  string `db:"creator_address"`
	ParsedWhiteListDataWithMetadata string `db:"parsed_whitelist_data_with_metadata"`
	TotalBalance                    string `db:"total_balance"`
	MinAge                          int64  `db:"min_age"`
	MaxAge                          int64  `db:"max_age"`
	StartTimestamp                  int64  `db:"start_timestamp"`
	EndTimestamp                    int64  `db:"end_timestamp"`
	VotesCount                      int64  `db:"votes_count"`
	Status                          string `db:"status"`
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

func (q votingQ) GetVotingBalance() (*big.Int, *big.Int, error) {
	var balancesStr struct {
		Balance      string `db:"residual_balance"`
		TotalBalance string `db:"total_balance"`
	}

	query := q.selector.RemoveColumns().Columns("residual_balance", "total_balance")
	if err := q.db.Get(&balancesStr, query); err != nil {
		if err != sql.ErrNoRows {
			return nil, nil, errors.Wrap(err, "failed to get voting info from db")
		}
		return nil, nil, nil
	}

	balance, success := new(big.Int).SetString(balancesStr.Balance, 10)
	if !success {
		return nil, nil, errors.New("error converting string balance to big.Int")
	}

	totalBalance, success := new(big.Int).SetString(balancesStr.Balance, 10)
	if !success {
		return nil, nil, errors.New("error converting string total balance to big.Int")
	}

	return balance, totalBalance, nil
}

// Get complete information for a particular voting
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

	var proposalInfo resources.ParsedWhiteListDataWithMetadata
	err = json.Unmarshal([]byte(votingInfo.ParsedWhiteListDataWithMetadata), &proposalInfo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal ProposalInfoWithConfig")
	}

	balance, success := new(big.Int).SetString(votingInfo.Balance, 10)
	if !success {
		return nil, errors.New("error converting string balance to big.Int")
	}

	totalBalance, success := new(big.Int).SetString(votingInfo.TotalBalance, 10)
	if !success {
		return nil, errors.New("error converting string total balance to big.Int")
	}

	return &data.VotingInfo{
		GasLimit:                        votingInfo.GasLimit,
		VotingId:                        votingInfo.VotingId,
		Balance:                         balance,
		CreatorAddress:                  votingInfo.CreatorAddress,
		ParsedWhiteListDataWithMetadata: proposalInfo,
		TotalBalance:                    totalBalance,
		MinAge:                          votingInfo.MinAge,
		MaxAge:                          votingInfo.MaxAge,
		StartTimestamp:                  votingInfo.StartTimestamp,
		EndTimestamp:                    votingInfo.EndTimestamp,
		VotesCount:                      votingInfo.VotesCount,
	}, nil
}

// Select the list of votes according to filters, if any
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

		totalBalance, success := new(big.Int).SetString(vote.TotalBalance, 10)
		if !success {
			continue
		}

		var proposalInfo resources.ParsedWhiteListDataWithMetadata
		err = json.Unmarshal([]byte(vote.ParsedWhiteListDataWithMetadata), &proposalInfo)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal ProposalInfoWithConfig")
		}
		result = append(result, &data.VotingInfo{
			GasLimit:                        vote.GasLimit,
			VotingId:                        vote.VotingId,
			Balance:                         balance,
			TotalBalance:                    totalBalance,
			CreatorAddress:                  vote.CreatorAddress,
			ParsedWhiteListDataWithMetadata: proposalInfo,
			MinAge:                          vote.MinAge,
			MaxAge:                          vote.MaxAge,
			StartTimestamp:                  vote.StartTimestamp,
			EndTimestamp:                    vote.EndTimestamp,
			VotesCount:                      vote.VotesCount,
			Status:                          vote.Status,
		})
	}

	return result, nil
}

func (q votingQ) InsertVotingInfo(value *data.VotingInfo) error {
	jsonDataProposalInfo, err := json.Marshal(value.ParsedWhiteListDataWithMetadata)
	if err != nil {
		return errors.Wrap(err, "failed to marshal ProposalInfoWithConfig")
	}

	query := sq.Insert(votingTableName).
		Columns(votingClolumns...).
		Values(
			value.VotingId,
			value.Balance.String(),
			value.GasLimit,
			value.CreatorAddress,
			jsonDataProposalInfo,
			value.TotalBalance.String(),
			value.MinAge,
			value.MaxAge,
			value.StartTimestamp,
			value.EndTimestamp,
			value.VotesCount,
		)

	err = q.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to insert voting info to db")
	}
	return nil
}

// Full update of information on a specific voting
func (q votingQ) UpdateVotingInfo(value *data.VotingInfo) error {
	jsonDataProposalInfo, err := json.Marshal(value.ParsedWhiteListDataWithMetadata)
	if err != nil {
		return errors.Wrap(err, "failed to marshal ProposalInfoWithConfig")
	}
	var query sq.UpdateBuilder
	query = sq.Update(votingTableName)
	values := []any{
		value.VotingId,
		value.Balance.String(),
		value.GasLimit,
		value.CreatorAddress,
		jsonDataProposalInfo,
		value.TotalBalance.String(),
		value.MinAge,
		value.MaxAge,
		value.StartTimestamp,
		value.EndTimestamp,
		value.VotesCount,
	}

	for indx, column := range votingClolumns {
		query = query.Set(column, values[indx])
	}

	query = query.Where(sq.Eq{"voting_id": value.VotingId})

	err = q.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to update voting info in db")
	}
	return nil
}

// Update certain columns. Responds to filters
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

	return q.withFilters(sq.Eq{"min_age": minAgeList})
}

func (q votingQ) FilterByMaxAge(maxAgeList ...int64) data.VotingQ {
	if len(maxAgeList) == 0 {
		return q
	}

	return q.withFilters(sq.Eq{"max_age": maxAgeList})
}

func (q votingQ) FilterByCitizenship(сitizenshipList ...string) data.VotingQ {
	if len(сitizenshipList) == 0 {
		return q
	}
	var stmt []string
	for _, citizenship := range сitizenshipList {
		stmt = append(stmt, fmt.Sprintf("parsed_whitelist_data_with_metadata @> '{\"parsed_voting_whitelist_data\": [{\"citizenship_whitelist\": [\"%v\"]}]}'", citizenship))
	}

	return q.withFilters(fmt.Sprintf("(%v)", strings.Join(stmt, " OR ")))
}

func (q votingQ) Page(page *pgdb.OffsetPageParams) data.VotingQ {
	q.selector = applyPage(page, q.selector)
	return q
}

func (q votingQ) FilterBySex(sexList ...string) data.VotingQ {
	if len(sexList) == 0 {
		return q
	}

	var stmt []string
	for _, sex := range sexList {
		queryParam := sex
		if sex != "0" {
			queryParam = hex.EncodeToString([]byte(sex))
		}
		stmt = append(stmt, fmt.Sprintf("parsed_whitelist_data_with_metadata @> '{\"parsed_voting_whitelist_data\": [{\"sex\": \"%s\"}]}'", queryParam))
	}

	return q.withFilters(fmt.Sprintf("(%v)", strings.Join(stmt, " OR ")))
}

func (q votingQ) WithStatus() data.VotingQ {
	status := fmt.Sprintf(` %s AS status`, statusCase)

	q.selector = q.selector.Column(status)
	return q
}

func (q votingQ) FilterBySatus(statusList ...string) data.VotingQ {
	if len(statusList) == 0 {
		return q
	}
	return q.withFilters(sq.Eq{statusCase: statusList})
}

func applyPage(page *pgdb.OffsetPageParams, sql sq.SelectBuilder) sq.SelectBuilder {
	if page.Limit == 0 {
		page.Limit = 15
	}
	if page.Order == "" {
		page.Order = pgdb.OrderTypeDesc
	}

	offset := page.Limit * page.PageNumber

	sql = sql.Limit(page.Limit).Offset(offset)

	if page.Order == pgdb.OrderTypeAsc {
		sql = sql.OrderBy("voting_id asc")
		return sql
	}

	sql = sql.OrderBy("voting_id desc")

	return sql
}

func (q votingQ) withFilters(stmt interface{}) data.VotingQ {
	q.selector = q.selector.Where(stmt)
	q.updater = q.updater.Where(stmt)

	return q
}
