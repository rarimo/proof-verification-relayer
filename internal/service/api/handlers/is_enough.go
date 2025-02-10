package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/rarimo/proof-verification-relayer/internal/checker"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func IsEnoughHandler(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	address = strings.ToLower(address)
	isEnough, err := checker.IsEnough(Config(r), address)
	if err == sql.ErrNoRows {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if err != nil {
		Log(r).Warnf("Failed check is enough balance: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.VotingAvailabilityResponse{
		Data: resources.VotingAvailability{
			Key: resources.Key{
				ID:   address,
				Type: resources.IS_ENOUGH,
			},
			Attributes: resources.VotingAvailabilityAttributes{
				IsEnough: &isEnough,
			},
		},
	})

}
