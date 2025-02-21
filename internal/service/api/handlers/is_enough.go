package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rarimo/proof-verification-relayer/internal/checker"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func IsEnoughHandler(w http.ResponseWriter, r *http.Request) {
	votingIdStr := chi.URLParam(r, "voting_id")

	votingId, err := strconv.ParseInt(votingIdStr, 10, 64)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	isEnough, err := checker.IsEnough(Config(r), votingId)
	if err != nil {
		if err != sql.ErrNoRows {
			Log(r).WithField("voting_id", votingId).Errorf("Failed check is enough balance: %v", err)
			ape.RenderErr(w, problems.InternalError())
			return
		}
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, resources.VotingAvailabilityResponse{
		Data: resources.VotingAvailability{
			Key: resources.Key{
				ID:   votingIdStr,
				Type: resources.IS_ENOUGH,
			},
			Attributes: resources.VotingAvailabilityAttributes{
				IsEnough: &isEnough,
			},
		},
	})
}
