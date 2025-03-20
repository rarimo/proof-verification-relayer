/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VotingInfo struct {
	Key
	Attributes VotingInfoAttributes `json:"attributes"`
}
type VotingInfoResponse struct {
	Data     VotingInfo `json:"data"`
	Included Included   `json:"included"`
}

type VotingInfoListResponse struct {
	Data     []VotingInfo `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustVotingInfo - returns VotingInfo from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVotingInfo(key Key) *VotingInfo {
	var votingInfo VotingInfo
	if c.tryFindEntry(key, &votingInfo) {
		return &votingInfo
	}
	return nil
}
