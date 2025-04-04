package handlers

import (
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func PredictHandlers(w http.ResponseWriter, r *http.Request) {
	req, value, err := requests.NewVotingPredictRequest(r, Log(r))
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	reqArgument, _ := new(big.Int).SetString(*value, 10)

	var resultAns *big.Int
	var attribute resources.VotingPredictRespAttributes
	timestamp := time.Now().UTC().Format(time.RFC3339)

	switch req.Data.Type {
	case resources.VOTE_PREDICT_AMOUNT:
		resultAns, err = predictAmountOrCountTx(Config(r), *req.Data.Attributes.VotingId, reqArgument, nil)
		if err != nil {
			Log(r).WithError(err).Error("failed check is predict")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		resultAnsStr := fmt.Sprintf("%d", resultAns)
		attribute.AmountPredict = &resultAnsStr
	case resources.VOTE_PREDICT_COUNT_TX:
		resultAns, err = predictAmountOrCountTx(Config(r), *req.Data.Attributes.VotingId, nil, reqArgument)
		if err != nil {
			Log(r).WithError(err).Error("failed check is predict")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		resultAnsStr := fmt.Sprintf("%d", resultAns)
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

func predictAmountOrCountTx(cfg config.Config, votingId int64, countTx *big.Int, amount *big.Int) (*big.Int, error) {
	txFee, err := getTxFee(cfg, votingId)
	if err != nil {
		return nil, errors.Wrap(err, "failed get fee")
	}
	if txFee.Int64() == 0 {
		big.NewInt(int64(cfg.VotingV2Config().GasLimit))
	}

	if countTx != nil {
		return new(big.Int).Mul(countTx, txFee), nil
	}

	return new(big.Int).Div(amount, txFee), nil
}
