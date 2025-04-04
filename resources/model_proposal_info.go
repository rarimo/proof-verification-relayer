/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ProposalInfo struct {
	Key
	Attributes ProposalInfoAttributes `json:"attributes"`
}
type ProposalInfoResponse struct {
	Data     ProposalInfo `json:"data"`
	Included Included     `json:"included"`
}

type ProposalInfoListResponse struct {
	Data     []ProposalInfo `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustProposalInfo - returns ProposalInfo from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustProposalInfo(key Key) *ProposalInfo {
	var proposalInfo ProposalInfo
	if c.tryFindEntry(key, &proposalInfo) {
		return &proposalInfo
	}
	return nil
}
