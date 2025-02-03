package requests

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	calldataRegexp = regexp.MustCompile("^0x[0-9a-fA-F]+$")
	addressRegexp  = regexp.MustCompile("^0x[0-9a-fA-F]{40}")
)

type SendTx struct {
	Key        resources.Key
	Attributes SendTxAttributes `json:"attributes"`
}

type SendTxAttributes struct {
	// Address of the contract to which the transaction data should be sent
	Destination string `json:"destination"`
	ProposalID  int64  `json:"proposal_id"`
	// Serialized transaction data
	TxData string `json:"tx_data"`
}
type SendTxRequest struct {
	Data SendTx `json:"data"`
}

func NewVotingRequest(r *http.Request) (req SendTxRequest, err error) {
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, errors.Wrap(err, "failed to unmarshal")
	}

	req.Data.Attributes.Destination = strings.ToLower(req.Data.Attributes.Destination)

	return req, validation.Errors{
		"data/tx_data":     validation.Validate(req.Data.Attributes.TxData, validation.Required, validation.Match(calldataRegexp)),
		"data/destination": validation.Validate(req.Data.Attributes.Destination, validation.Required, validation.Match(addressRegexp)),
		"data/proposal_id": validation.Validate(req.Data.Attributes.ProposalID, validation.Required),
	}.Filter()
}
