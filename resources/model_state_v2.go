/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type StateV2 struct {
	Key
	Attributes StateV2Attributes `json:"attributes"`
}
type StateV2Response struct {
	Data     StateV2  `json:"data"`
	Included Included `json:"included"`
}

type StateV2ListResponse struct {
	Data     []StateV2 `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustStateV2 - returns StateV2 from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustStateV2(key Key) *StateV2 {
	var stateV2 StateV2
	if c.tryFindEntry(key, &stateV2) {
		return &stateV2
	}
	return nil
}
