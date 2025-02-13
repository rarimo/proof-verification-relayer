/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VotingPredictAttributes struct {
	// The amount of tokens available for covering transactions
	Amount *string `json:"amount,omitempty"`
	// The number of transactions that need to be covered
	CountTx *string `json:"count_tx,omitempty"`
	// Address of the contract to which the vote
	VotingId string `json:"voting_id"`
}
