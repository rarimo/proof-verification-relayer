/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type StateV2Attributes struct {
	// Root hash of the state tree
	Root string `json:"root"`
	// Signature of root state signed by relayer private key, with the last byte of the signature set to 27/28.
	Signature string `json:"signature"`
	// Time indicates when the event was caught, a.k.a state transition timestamp
	Timestamp int64 `json:"timestamp"`
}
