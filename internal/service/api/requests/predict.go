package requests

import (
	"encoding/json"
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	amountRegexp = regexp.MustCompile("^[0-9]+$")
)

func NewVotingPredictRequest(r *http.Request, logger *logan.Entry) (req resources.VotingPredictRequest, value *string, err error) {
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, nil, errors.Wrap(err, "failed to unmarshal")
	}
	if req.Data.Attributes.VotingId == nil {
		req.Data.Attributes.VotingId = new(int64)
	}

	switch req.Data.Type {
	case resources.VOTE_PREDICT_AMOUNT:
		return req, req.Data.Attributes.CountTx, validation.Errors{
			"data/count_tx": validation.Validate(req.Data.Attributes.CountTx, validation.Required, validation.Match(amountRegexp)),
		}.Filter()
	case resources.VOTE_PREDICT_COUNT_TX:
		return req, req.Data.Attributes.Amount, validation.Errors{
			"data/amount": validation.Validate(req.Data.Attributes.Amount, validation.Required, validation.Match(amountRegexp)),
		}.Filter()
	default:
		return req, nil, errors.New("unknown resource type")
	}
}
