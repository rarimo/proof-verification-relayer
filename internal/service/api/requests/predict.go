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
	amountRegexp = regexp.MustCompile("[0-9]{1,}")
)

func NewVotingPredictRequest(r *http.Request) (req resources.VotingPredictRequest, err error) {
	// log.Fatalf("Request")
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, errors.Wrap(err, "failed to unmarshal")
	}

	req.Data.Attributes.VoteAddress = strings.ToLower(req.Data.Attributes.VoteAddress)

	return req, validation.Errors{
		"data/amount":  validation.Validate(req.Data.Attributes.Amount, validation.Required, validation.Match(amountRegexp)),
		"data/address": validation.Validate(req.Data.Attributes.VoteAddress, validation.Required, validation.Match(AddressRegexp)),
	}.Filter()
}
