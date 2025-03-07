/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

// Proposal info
type VotingInfoConfig struct {
	// Detailed description of the proposal
	Description string `json:"description"`
	// End time of the proposal
	EndTimestamp int64 `json:"end_timestamp"`
	// Multichoice option for voting
	Multichoice int64 `json:"multichoice"`
	// Unique identifier of the proposal
	ProposalId int64 `json:"proposal_id"`
	// Start time of the proposal
	StartTimestamp int64 `json:"start_timestamp"`
	// List of addresses allowed to vote
	VotingWhitelist []string `json:"votingWhitelist"`
	// Additional whitelist data in bytes
	VotingWhitelistData []string `json:"votingWhitelistData"`
}
