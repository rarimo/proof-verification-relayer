package handlers

import (
	"fmt"
	"math/big"
	"net/http"
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

	req, value, err := requests.NewVotingPredictRequest(r, Log(r))
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	reqArgument, _ := new(big.Int).SetString(*value, 10)
	resultAns, err := сheckByType[string(req.Data.Type)](Config(r), *req.Data.Attributes.VotingId, reqArgument)
	if err != nil {
		Log(r).Warnf("Failed check is predict: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	timestamp := time.Now().UTC().Format(time.RFC3339)
	resultAnsStr := fmt.Sprintf("%d", resultAns)
	var attribute resources.VotingPredictRespAttributes
	switch req.Data.Type {
	case resources.VOTE_PREDICT_AMOUNT:
		attribute.AmountPredict = &resultAnsStr
	case resources.VOTE_PREDICT_COUNT_TX:
		attribute.CountTxPredict = &resultAnsStr
	}

	ape.Render(w, resources.VotingPredictRespResponse{
		Data: resources.VotingPredictResp{
			Key: resources.Key{
				ID:   fmt.Sprint(*req.Data.Attributes.VotingId) + ":" + *value + ":" + timestamp,
				Type: req.Data.Type,
			},
			Attributes: attribute,
		},
	})
}
