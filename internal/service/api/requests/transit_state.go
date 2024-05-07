package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type TransitStateData struct {
	TxData string `json:"tx_data"`
}

type TransitStateDataRequest struct {
	Data TransitStateData `json:"data"`
}

func TransitStateDataRequestRequest(r *http.Request) (TransitStateDataRequest, error) {
	var request TransitStateDataRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, validation.Errors{
			"body": errors.Wrap(err, "failed to unmarshal"),
		}.Filter()
	}

	return request, validateTransitStateDataRequest(request)
}

func validateTransitStateDataRequest(r TransitStateDataRequest) error {
	return validation.Errors{
		"/data/tx_data": validation.Validate(r.Data.TxData, validation.Required),
	}.Filter()
}
