/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VotingInfoContractInfo struct {
	Config VotingInfoConfig `json:"config"`
	// Address of the proposal SMT
	ProposalSMT string `json:"proposalSMT"`
	// Status of the proposal
	Status        uint8     `json:"status"`
	VotingResults [][]int64 `json:"voting_results"`
}
