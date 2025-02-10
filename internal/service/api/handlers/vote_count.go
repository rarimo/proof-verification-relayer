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

func VoteCountHandlers(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	address = strings.ToLower(address)

	countTx, err := checker.GetCountTx(Config(r), address)
	if err == sql.ErrNoRows {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if err != nil {
		Log(r).Errorf("Failed get count tx for wallet: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, &resources.VotePrediction{
		Id:         address,
		Type:       "vote_count",
		Attributes: resources.VotePredictionAttributes{VoteCount: &countTx}})
}
