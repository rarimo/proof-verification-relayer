package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/rarimo/proof-verification-relayer/internal/checker"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func PredictHandlers(w http.ResponseWriter, r *http.Request) {
	сheckByType := map[string]func(cfg config.Config, addr string, countTx uint64) (uint64, error){
		string(resources.VOTE_PREDICT_AMOUNT):   checker.AmountForCountTx,
		string(resources.VOTE_PREDICT_COUNT_TX): checker.GetPredictCount,
	}

	req, value, err := requests.NewVotingPredictRequest(r)
	if err != nil || value == nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	_, exists := сheckByType[string(req.Data.Type)]
	if !exists {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	Log(r).Infof("New predict request for address: %v", req.Data.Attributes.VoteAddress)

	reqArgument, err := strconv.ParseUint(*value, 10, 64)
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	votingAddress := req.Data.Attributes.VoteAddress
	resultAns, err := сheckByType[string(req.Data.Type)](Config(r), votingAddress, reqArgument)
	if err == sql.ErrNoRows {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if err != nil {
		Log(r).Warnf("Failed check is predict: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	timestamp := time.Now().UTC().Format(time.RFC3339)
	resultAnsStr := fmt.Sprintf("%d", resultAns)
	var attribut resources.VotingPredictRespAttributes
	switch req.Data.Type {
	case resources.VOTE_PREDICT_AMOUNT:
		attribut.AmountPredict = &resultAnsStr
	case resources.VOTE_PREDICT_COUNT_TX:
		attribut.CountTxPredict = &resultAnsStr
	}

	ape.Render(w, resources.VotingPredictRespResponse{
		Data: resources.VotingPredictResp{
			Key: resources.Key{
				ID:   votingAddress + ":" + *value + ":" + timestamp,
				Type: req.Data.Type,
			},
			Attributes: attribut,
		},
	})

}
