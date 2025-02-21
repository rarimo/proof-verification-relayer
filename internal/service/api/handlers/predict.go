package handlers

import (
	"fmt"
	"math/big"
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
	сheckByType := map[string]func(cfg config.Config, votingId int64, countTx *big.Int) (*big.Int, error){
		string(resources.VOTE_PREDICT_AMOUNT):   checker.GetAmountForCountTx,
		string(resources.VOTE_PREDICT_COUNT_TX): checker.GetCountTxForAmount,
	}

	req, value, err := requests.NewVotingPredictRequest(r)
	if err != nil || value == nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	reqArgument, ok := new(big.Int).SetString(*value, 10)
	if !ok {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var votingIdStr string
	switch req.Data.Attributes.VotingId {
	case nil:
		votingIdStr = "0"
	default:
		votingIdStr = *req.Data.Attributes.VotingId
	}

	votingId, err := strconv.ParseInt(votingIdStr, 10, 64)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	resultAns, err := сheckByType[string(req.Data.Type)](Config(r), votingId, reqArgument)
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
				ID:   votingIdStr + ":" + *value + ":" + timestamp,
				Type: req.Data.Type,
			},
			Attributes: attribut,
		},
	})
}
