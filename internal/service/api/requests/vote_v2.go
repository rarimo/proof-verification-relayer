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

func NewVotingRequest(r *http.Request) (req resources.SendTxRequest, err error) {
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, errors.Wrap(err, "failed to unmarshal")
	}

	req.Data.Attributes.Destination = strings.ToLower(req.Data.Attributes.Destination)

	return req, validation.Errors{
		"data/tx_data":     validation.Validate(req.Data.Attributes.TxData, validation.Required, validation.Match(calldataRegexp)),
		"data/destination": validation.Validate(req.Data.Attributes.Destination, validation.Required, validation.Match(addressRegexp)),
	}.Filter()
}
