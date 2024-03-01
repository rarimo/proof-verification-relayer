/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Tx struct {
	Key
	Attributes TxAttributes `json:"attributes"`
}
type TxResponse struct {
	Data     Tx       `json:"data"`
	Included Included `json:"included"`
}

type TxListResponse struct {
	Data     []Tx     `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustTx - returns Tx from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustTx(key Key) *Tx {
	var tx Tx
	if c.tryFindEntry(key, &tx) {
		return &tx
	}
	return nil
}
