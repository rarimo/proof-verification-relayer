package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type VerifyProofRequestData struct {
	TxData string `json:"tx_data"`
}

type VerifyProofRequest struct {
	Data VerifyProofRequestData `json:"data"`
}

func NewVerifyProofRequest(r *http.Request) (VerifyProofRequest, error) {
	var request VerifyProofRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, nil
}
