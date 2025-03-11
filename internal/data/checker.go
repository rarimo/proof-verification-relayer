package data

import (
	"math/big"
	"time"

	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
)

type VotingWhitelistDataBigInt struct {
	CitizenshipWhitelist                []*big.Int `abi:"citizenshipWhitelist"`
	IdentityCreationTimestampUpperBound *big.Int   `abi:"identityCreationTimestampUpperBound"`
	IdentityCounterUpperBound           *big.Int   `abi:"identityCounterUpperBound"`
	BirthDateUpperbound                 *big.Int   `abi:"birthDateUpperbound"`
	ExpirationDateLowerBound            *big.Int   `abi:"expirationDateLowerBound"`
}

// type ParsedVotingWhitelistData struct {
// 	CitizenshipWhitelist                []string `json:"citizenship_whitelist"`
// 	IdentityCreationTimestampUpperBound string   `json:"identity_creation_timestamp_upper_bound"`
// 	IdentityCounterUpperBound           string   `json:"identity_counter_upper_bound" `
// 	BirthDateUpperbound                 string   `json:"birth_date_upperbound"`
// 	ExpirationDateLowerBound            string   `json:"expiration_date_lower_bound"`
// 	MinAge                              int64    `json:"min_age"`
// }

type VotingInfo struct {
	VotingId               int64                          `db:"voting_id"`
	Balance                *big.Int                       `db:"residual_balance"`
	GasLimit               uint64                         `db:"gas_limit"`
	CreatorAddress         string                         `db:"creator_address"`
	ProposalInfoWithConfig resources.VotingInfoAttributes `db:"proposal_info_with_config"`
}

type ProcessedEvent struct {
	TransactionHash []byte    `db:"transaction_hash"`
	LogIndex        int64     `db:"log_index"`
	CreatedAt       time.Time `db:"created_at"`
	BlockNumber     int64     `db:"block_number"`
}

type CheckerDB interface {
	New() CheckerDB
	CheckerQ() CheckerQ
}

type CheckerQ interface {
	InsertVotingInfo(value *VotingInfo) error
	GetVotingInfo(votingId int64) (*VotingInfo, error)
	UpdateVotingInfo(value *VotingInfo) error
	SelectVotes(req requests.ProposalInfoFilter) ([]*VotingInfo, error)

	InsertProcessedEvent(value ProcessedEvent) error
	GetLastBlock() (uint64, error)
	CheckProcessedEventExist(value ProcessedEvent) (bool, error)
}
