package pg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/big"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
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

type votingQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func (q *votingQ) New() data.VotingQ {
	return NewVotingQ(q.db.Clone())
}

func NewVotingQ(db *pgdb.DB) data.VotingQ {
	return &votingQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

func (cq *votingQ) GetVotingInfo(votingId int64) (*data.VotingInfo, error) {
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

func (cq *votingQ) SelectVotingInfo(req requests.ProposalInfoFilter) ([]*data.VotingInfo, error) {
	var votingInfoList []votingInf

	query := sq.Select("*").From("voting_contract_accounts")
	if len(req.CreatorAddress) > 0 {
		query = query.Where(
			sq.Eq{"creator_address": req.CreatorAddress},
		)
	}
	if len(req.MinAge) > 0 {
		query = query.Where(
			sq.Eq{
				"proposal_info_with_config #> '{contract, config, parsed_voting_whitelist_data, 0, min_age}'": req.MinAge},
		)
	}
	if len(req.ProposalId) > 0 {
		query = query.Where(
			sq.Eq{
				"voting_id": req.ProposalId},
		)
	}

	for _, citizenship := range req.CitizenshipList {
		query = query.Where(
			"proposal_info_with_config #> '{contract, config, parsed_voting_whitelist_data, 0, citizenship_whitelist}' @> ?", fmt.Sprintf("[\"%v\"]", citizenship))
	}

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

func (q *votingQ) InsertVotingInfo(value *data.VotingInfo) error {
	jsonDataProposalInfo, err := json.Marshal(value.ProposalInfoWithConfig)
	if err != nil {
		return errors.Wrap(err, "failed to marshal ProposalInfoWithConfig")
	}

	query := sq.Insert("voting_contract_accounts").
		Columns("voting_id", "residual_balance", "gas_limit", "creator_address", "proposal_info_with_config").
		Values(value.VotingId, value.Balance.String(), value.GasLimit, value.CreatorAddress, jsonDataProposalInfo)

	err = q.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to insert voting info to db")
	}
	return nil
}

func (q *votingQ) UpdateVotingInfo(value *data.VotingInfo) error {
	jsonDataProposalInfo, err := json.Marshal(value.ProposalInfoWithConfig)
	if err != nil {
		return errors.Wrap(err, "failed to marshal ProposalInfoWithConfig")
	}
	var query sq.UpdateBuilder
	query = sq.Update("voting_contract_accounts")
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
