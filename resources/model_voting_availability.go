/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type VotingAvailability struct {
	Key
	Attributes VotingAvailabilityAttributes `json:"attributes"`
}
type VotingAvailabilityResponse struct {
	Data     VotingAvailability `json:"data"`
	Included Included           `json:"included"`
}

type VotingAvailabilityListResponse struct {
	Data     []VotingAvailability `json:"data"`
	Included Included             `json:"included"`
	Links    *Links               `json:"links"`
}

// MustVotingAvailability - returns VotingAvailability from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustVotingAvailability(key Key) *VotingAvailability {
	var votingAvailability VotingAvailability
	if c.tryFindEntry(key, &votingAvailability) {
		return &votingAvailability
	}
	return nil
}
