/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ParsedVotingWhitelistData struct {
	// birth_date_upperbound in hex format
	BirthDateUpperbound                 string   `json:"birth_date_upperbound"`
	CitizenshipWhitelist                []string `json:"citizenship_whitelist"`
	ExpirationDateLowerBound            string   `json:"expiration_date_lower_bound"`
	IdentityCounterUpperBound           string   `json:"identity_counter_upper_bound"`
	IdentityCreationTimestampUpperBound string   `json:"identity_creation_timestamp_upper_bound"`
	MinAge                              int64    `json:"min_age"`
}
