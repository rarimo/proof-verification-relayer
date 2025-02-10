package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/rarimo/proof-verification-relayer/internal/checker"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func PredictHandlers(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVotingPredictRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Log(r).Infof("New predict request for address: %v", req.Data.Attributes.VoteAddress)
	amount, err := strconv.ParseUint(req.Data.Attributes.Amount, 10, 64)
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	votingAddress := req.Data.Attributes.VoteAddress
	countTx, err := checker.GetPredictCount(Config(r), votingAddress, amount)
	if err == sql.ErrNoRows {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if err != nil {
		Log(r).Warnf("Failed check is predict Tx count: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var resp resources.VotingPredictResponse
	timestamp := time.Now().UTC().Format(time.RFC3339)
	resp.Id = votingAddress + ":" + fmt.Sprintf("%d", amount) + ":" + timestamp
	resp.Type = "vote_predict"
	resp.Attributes.CountTxPredict = fmt.Sprintf("%d", countTx)

	ape.Render(w, resp)
}
