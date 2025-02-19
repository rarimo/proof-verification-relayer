/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VoteCount struct {
	Key
	Attributes VoteCountAttributes `json:"attributes"`
}
type VoteCountResponse struct {
	Data     VoteCount `json:"data"`
	Included Included  `json:"included"`
}

type VoteCountListResponse struct {
	Data     []VoteCount `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustVoteCount - returns VoteCount from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVoteCount(key Key) *VoteCount {
	var voteCount VoteCount
	if c.tryFindEntry(key, &voteCount) {
		return &voteCount
	}
	return nil
}
