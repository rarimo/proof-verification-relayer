/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VotingPredictResp struct {
	Key
	Attributes VotingPredictRespAttributes `json:"attributes"`
}
type VotingPredictRespResponse struct {
	Data     VotingPredictResp `json:"data"`
	Included Included          `json:"included"`
}

type VotingPredictRespListResponse struct {
	Data     []VotingPredictResp `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
}

// MustVotingPredictResp - returns VotingPredictResp from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVotingPredictResp(key Key) *VotingPredictResp {
	var votingPredictResp VotingPredictResp
	if c.tryFindEntry(key, &votingPredictResp) {
		return &votingPredictResp
	}
	return nil
}
