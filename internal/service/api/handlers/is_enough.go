package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rarimo/proof-verification-relayer/internal/checker"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func IsEnoughHandler(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")

	isEnough, err := checker.IsEnough(Config(r), address)
	if err != nil {
		Log(r).Warnf("Failed check is enough balance: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, &resources.VotingAvailability{
		Id:   address,
		Type: "is_enough",
		Attributes: resources.VotingAvailabilityAttributes{
			IsEnough: &isEnough,
		},
	})

}
