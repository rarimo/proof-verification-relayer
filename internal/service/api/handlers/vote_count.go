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

func VoteCountHandlers(w http.ResponseWriter, r *http.Request) {
	votingIdStr := chi.URLParam(r, "voting_id")

	votingId, err := strconv.ParseInt(votingIdStr, 10, 64)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	countTx, err := checker.GetCountTx(Config(r), votingId)
	if err != nil {
		if err != sql.ErrNoRows {
			Log(r).WithField("voting_id", votingId).Errorf("Failed get count tx for vote: %v", err)
			ape.RenderErr(w, problems.InternalError())
			return
		}
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, resources.VoteCountResponse{
		Data: resources.VoteCount{
			Key: resources.Key{
				ID:   votingIdStr,
				Type: resources.VOTE_COUNT,
			},
			Attributes: resources.VoteCountAttributes{
				VoteCount: &countTx,
			},
		}})
}
