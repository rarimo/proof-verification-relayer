/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VotingPredict struct {
	Key
	Attributes VotingPredictAttributes `json:"attributes"`
}
type VotingPredictRequest struct {
	Data     VotingPredict `json:"data"`
	Included Included      `json:"included"`
}

type VotingPredictListRequest struct {
	Data     []VotingPredict `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustVotingPredict - returns VotingPredict from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVotingPredict(key Key) *VotingPredict {
	var votingPredict VotingPredict
	if c.tryFindEntry(key, &votingPredict) {
		return &votingPredict
	}
	return nil
}
