package data

import (
	"math/big"
	"time"

	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
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
	SelectVotingInfo(req requests.ProposalInfoFilter) ([]*VotingInfo, error)

	InsertProcessedEvent(value ProcessedEvent) error
	GetLastBlock() (uint64, error)
	CheckProcessedEventExist(value ProcessedEvent) (bool, error)
}
