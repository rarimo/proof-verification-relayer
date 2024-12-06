/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type State struct {
	Key
	Attributes StateAttributes `json:"attributes"`
}
type StateResponse struct {
	Data     State    `json:"data"`
	Included Included `json:"included"`
}

type StateListResponse struct {
	Data     []State  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustState - returns State from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustState(key Key) *State {
	var state State
	if c.tryFindEntry(key, &state) {
		return &state
	}
	return nil
}
