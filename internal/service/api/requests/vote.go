package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type VoteRequestData struct {
	TxData       string `json:"tx_data"`
	Voting       string `json:"voting"`
	Registration string `json:"registration"`
}

type VoteRequest struct {
	Data VoteRequestData `json:"data"`
}

func NewVoteRequest(r *http.Request) (VoteRequest, error) {
	var request VoteRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, validateVoteRequest(request)
}

func validateVoteRequest(r VoteRequest) error {
	return validation.Errors{
		"/data/voting": validation.Validate(
			r.Data.Voting, validation.Required, validation.By(isAddressRule)),
		"/data/registration": validation.Validate(
			r.Data.Registration, validation.Required, validation.By(isAddressRule)),
	}.Filter()
}
