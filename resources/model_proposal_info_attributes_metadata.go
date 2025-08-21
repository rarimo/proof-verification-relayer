/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

// Additional proposal data stored on IPFS
type ProposalInfoAttributesMetadata struct {
	AcceptedOptions []Options `json:"acceptedOptions"`
	// Detailed description of the proposal
	Description string `json:"description"`
	// Proposal image CID
	ImageCid string `json:"imageCid"`
	// Indicates if the proposal is ranking based
	RankingBased *bool `json:"rankingBased,omitempty"`
	// Title of the proposal
	Title string `json:"title"`
}
