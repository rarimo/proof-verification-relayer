package handlers

import (
	"context"
	"math/big"
	"net/http"
	"strconv"

	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func ProposalsInfo(w http.ResponseWriter, r *http.Request) {
	cfg := Config(r)

	req, err := requests.ProposalsInfoRequest(r)
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get filters params")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	proposalList, err := GetProposalList(cfg, req)
	if err != nil {
		Log(r).WithError(err).Error("failed to get voting info list")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	gasPrice, err := VotingV2Config(r).RPC.SuggestGasPrice(context.Background())
	if err != nil {
		Log(r).WithError(err).Error("failed to get gas price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// if gasPrice is 0, for example, on the testnet, the value is set to 1
	if gasPrice.Uint64() == 0 {
		gasPrice = big.NewInt(1)
	}

	var resp resources.ProposalInfoListResponse
	resp.Data = make([]resources.ProposalInfo, 0, len(proposalList))
	for _, proposal := range proposalList {
		id := strconv.FormatInt(proposal.VotingId, 10)
		totalTXFee := new(big.Int).Mul(big.NewInt(int64(proposal.GasLimit)), gasPrice)
		remainingVotesCount := new(big.Int).Div(proposal.Balance, totalTXFee)

		resp.Data = append(resp.Data, resources.ProposalInfo{
			Key: resources.Key{
				ID:   id,
				Type: resources.PROPOSALS,
			},
			Attributes: resources.ProposalInfoAttributes{
				Metadata:            proposal.ParsedWhiteListDataWithMetadata.Metadata,
				Owner:               proposal.CreatorAddress,
				StartTimestamp:      proposal.StartTimestamp,
				EndTimestamp:        proposal.EndTimestamp,
				RemainingBalance:    proposal.Balance.String(),
				TotalBalance:        proposal.TotalBalance.String(),
				VotesCount:          proposal.VotesCount,
				Status:              proposal.Status,
				RemainingVotesCount: remainingVotesCount.Int64(),
			},
		})
	}
	ape.Render(w, resp)
}

func GetProposalList(cfg config.Config, req requests.ProposalInfoFilter) ([]*data.VotingInfo, error) {
	pgDB := pg.NewVotingQ(cfg.DB().Clone()).
		WithStatus().
		FilterByCitizenship(req.CitizenshipList...).
		FilterByCreator(req.CreatorAddress...).
		FilterByMinAge(req.MinAge...).
		FilterByMaxAge(req.MaxAge...).
		FilterByVotingId(req.ProposalId...).
		FilterBySatus(req.VotingStatus...).
		FilterBySex(req.SexFilter...).
		Page(&req.OffsetPageParams)

	votingInfoList, err := pgDB.SelectVotingInfo()
	if err != nil {
		return nil, errors.Wrap(err, "failed to select voting info list")
	}

	return votingInfoList, nil
}
