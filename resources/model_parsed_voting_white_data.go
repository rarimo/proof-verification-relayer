/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ParsedVotingWhiteData struct {
	// birth_date_lowerbound in hex format without first 0x
	BirthDateLowerbound string `json:"birth_date_lowerbound"`
	// birth_date_upper_bound in hex format without first 0x
	BirthDateUpperBound string `json:"birth_date_upper_bound"`
	// Citizenship whitelist for the voting
	CitizenshipWhitelist []string `json:"citizenship_whitelist"`
	// expiration_date_lower_bound in hex format without first 0x
	ExpirationDateLowerBound string `json:"expiration_date_lower_bound"`
	// uniqueness 0 | 1
	IdentityCounterUpperBound           string `json:"identity_counter_upper_bound"`
	IdentityCreationTimestampUpperBound string `json:"identity_creation_timestamp_upper_bound"`
	MaxAge                              int64  `json:"max_age"`
	MinAge                              int64  `json:"min_age"`
	// Proposal selector without first 0x
	Selector string `json:"selector"`
	// toHex('M') // male, toHex('F') // female, 0 // any
	Sex string `json:"sex"`
}
