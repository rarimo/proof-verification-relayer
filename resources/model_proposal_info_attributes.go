/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ProposalInfoAttributes struct {
	// End time of the proposal
	EndTimestamp int64                          `json:"end_timestamp"`
	Metadata     ProposalInfoAttributesMetadata `json:"metadata"`
	// Address creator of the proposal
	Owner string `json:"owner"`
	// Remaining balance
	RemainingBalance string `json:"remaining_balance"`
	// Total number of available votes for this proposal
	RemainingVotesCount int64 `json:"remaining_votes_count"`
	// Start time of the proposal
	StartTimestamp int64  `json:"start_timestamp"`
	Status         string `json:"status"`
	// Balance what is the starting balance with deposits
	TotalBalance string `json:"total_balance"`
	// Total number of votes already cast for this proposal
	VotesCount int64 `json:"votes_count"`
}
