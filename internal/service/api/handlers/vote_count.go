package handlers

import (
	"context"
	"database/sql"
	"math/big"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func VoteCountHandlers(w http.ResponseWriter, r *http.Request) {
	votingIdStr := chi.URLParam(r, "voting_id")

	votingId, err := strconv.ParseInt(votingIdStr, 10, 64)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	countTx, err := getCountTx(Config(r), votingId)
	if err != nil {
		if err != sql.ErrNoRows {
			Log(r).WithError(err).WithField("voting_id", votingId).Error("failed get count tx for vote")
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

func getCountTx(cfg config.Config, votingId int64) (uint64, error) {
	vq := pg.NewVotingQ(cfg.DB().Clone()).FilterByVotingId(votingId)
	txFee, err := getTxFee(cfg, votingId)
	if err != nil {
		return 0, errors.Wrap(err, "failed get fee")
	}

	balance, _, err := vq.GetVotingBalance()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get voting balance from db")
	}
	if balance == nil {
		return 0, sql.ErrNoRows
	}
	if txFee.Int64() == 0 {
		txFee = big.NewInt(int64(cfg.VotingV2Config().GasLimit))
	}
	countTx := new(big.Int).Div(balance, txFee)
	return countTx.Uint64(), nil
}

func getTxFee(cfg config.Config, votingId int64) (*big.Int, error) {
	networkParam := cfg.VotingV2Config()
	client := networkParam.RPC

	feeCap, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed get gas price")
	}
	if feeCap.Uint64() == 0 {
		feeCap = big.NewInt(1)
	}

	var gasLimitFromDB int64
	gasLimit := int64(cfg.VotingV2Config().GasLimit)

	if err := pg.NewVotingQ(cfg.DB()).FilterByVotingId(votingId).Get("gas_limit", &gasLimitFromDB); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "failed to get voting info from db")
		}
	}

	if gasLimitFromDB != 0 {
		gasLimit = gasLimitFromDB
	}

	totalFee := new(big.Int).Mul(big.NewInt(gasLimit), feeCap)
	return totalFee, nil
}
