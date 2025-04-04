package data

import (
	"math/big"

	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
)

// The statuses that can be in the voting
var (
	StatusWaiting = "waiting"
	StatusStarted = "started"
	StatusEnded   = "ended"
)

type VotingWhitelistDataBigInt struct {
	Selector                            *big.Int   `json:"selector"`
	CitizenshipWhitelist                []*big.Int `json:"citizenshipWhitelist"`
	IdentityCreationTimestampUpperBound *big.Int   `json:"identityCreationTimestampUpperBound"`
	IdentityCounterUpperBound           *big.Int   `json:"identityCounterUpperBound"`
	Sex                                 *big.Int   `json:"sex"`
	BirthDateLowerbound                 *big.Int   `json:"birthDateLowerbound"`
	BirthDateUpperbound                 *big.Int   `json:"birthDateUpperbound"`
	ExpirationDateLowerBound            *big.Int   `json:"expirationDateLowerBound"`
}

// This structure contains the basic data about the voting (also referred to here as a proposal).
type VotingInfo struct {
	VotingId                        int64                                     `db:"voting_id"`
	Balance                         *big.Int                                  `db:"residual_balance"` // Residual voting balance. It changes both during deposits and vote processing
	GasLimit                        uint64                                    `db:"gas_limit"`
	CreatorAddress                  string                                    `db:"creator_address"`
	ParsedWhiteListDataWithMetadata resources.ParsedWhiteListDataWithMetadata `db:"parsed_whitelist_data_with_metadata"`
	TotalBalance                    *big.Int                                  `db:"total_balance"` // The balance that corresponds to the initial balance with the addition of all deposits
	MinAge                          int64                                     `db:"min_age"`
	MaxAge                          int64                                     `db:"max_age"`
	StartTimestamp                  int64                                     `db:"start_timestamp"`
	EndTimestamp                    int64                                     `db:"end_timestamp"`
	VotesCount                      int64                                     `db:"votes_count"`
	Status                          string                                    `db:"status"`
}

type VotingQ interface {
	New() VotingQ

	Update(fields map[string]any) error
	InsertVotingInfo(value *VotingInfo) error
	UpdateVotingInfo(value *VotingInfo) error

	Get(column string, dest interface{}) error
	GetVotingBalance() (*big.Int, *big.Int, error)
	GetVotingInfo(votingId int64) (*VotingInfo, error)
	SelectVotingInfo() ([]*VotingInfo, error)

	FilterByVotingId(votingIds ...int64) VotingQ
	FilterByCreator(creators ...string) VotingQ
	FilterBySex(sexList ...string) VotingQ
	FilterByMinAge(minAgeList ...int64) VotingQ
	FilterByMaxAge(maxAgeList ...int64) VotingQ
	FilterBySatus(statusList ...string) VotingQ
	WithStatus() VotingQ

	FilterByCitizenship(сitizenshipList ...string) VotingQ
	Page(page *pgdb.OffsetPageParams) VotingQ
}
