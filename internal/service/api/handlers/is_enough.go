package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func IsEnoughHandler(w http.ResponseWriter, r *http.Request) {
	votingIdStr := chi.URLParam(r, "voting_id")

	votingId, err := strconv.ParseInt(votingIdStr, 10, 64)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	isEnough, err := isEnough(Config(r), votingId)
	if err != nil {
		if err != sql.ErrNoRows {
			Log(r).WithError(err).WithField("voting_id", votingId).Error("failed check is enough balance")
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

func isEnough(cfg config.Config, votingId int64) (bool, error) {
	countTx, err := getCountTx(cfg, votingId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, err
		}
		return false, errors.Wrap(err, "failed to get the available number of tx for voting")
	}
	return countTx != 0, nil
}
