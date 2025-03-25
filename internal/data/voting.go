package data

import (
	"math/big"

	"github.com/rarimo/proof-verification-relayer/resources"
)

type VotingWhitelistDataBigInt struct {
	CitizenshipWhitelist                []*big.Int `json:"citizenshipWhitelist"`
	IdentityCreationTimestampUpperBound *big.Int   `json:"identityCreationTimestampUpperBound"`
	IdentityCounterUpperBound           *big.Int   `json:"identityCounterUpperBound"`
	BirthDateUpperbound                 *big.Int   `json:"birthDateUpperbound"`
	ExpirationDateLowerBound            *big.Int   `json:"expirationDateLowerBound"`
}

type VotingInfo struct {
	VotingId               int64                          `db:"voting_id"`
	Balance                *big.Int                       `db:"residual_balance"`
	GasLimit               uint64                         `db:"gas_limit"`
	CreatorAddress         string                         `db:"creator_address"`
	ProposalInfoWithConfig resources.VotingInfoAttributes `db:"proposal_info_with_config"`
}

type VotingQ interface {
	New() VotingQ

	InsertVotingInfo(value *VotingInfo) error
	UpdateVotingInfo(value *VotingInfo) error

	GetVotingInfo(votingId int64) (*VotingInfo, error)
	SelectVotingInfo() ([]*VotingInfo, error)

	FilterByVotingId(votingIds ...int64) VotingQ
	FilterByCreator(creators ...string) VotingQ
	FilterByMinAge(minAgeList ...int64) VotingQ
	FilterByCitizenship(сitizenshipList ...string) VotingQ
}
