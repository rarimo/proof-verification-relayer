/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SendTx struct {
	Key
	Attributes SendTxAttributes `json:"attributes"`
}
type SendTxRequest struct {
	Data     SendTx   `json:"data"`
	Included Included `json:"included"`
}

type SendTxListRequest struct {
	Data     []SendTx `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustSendTx - returns SendTx from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustSendTx(key Key) *SendTx {
	var sendTx SendTx
	if c.tryFindEntry(key, &sendTx) {
		return &sendTx
	}
	return nil
}
