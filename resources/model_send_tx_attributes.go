/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SendTxAttributes struct {
	// Address of the contract to which the transaction data should be sent
	Destination string `json:"destination"`
	ProposalId  int64  `json:"proposal_id"`
	// Serialized transaction data
	TxData string `json:"tx_data"`
}
