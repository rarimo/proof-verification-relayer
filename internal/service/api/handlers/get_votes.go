package handlers

import (
	"net/http"
	"strconv"

	"github.com/rarimo/proof-verification-relayer/internal/checker"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetVotes(w http.ResponseWriter, r *http.Request) {
	cfg := Config(r)

	req, err := requests.GetVotesRequest(r)
	if err != nil {
		Log(r).Errorf("failed get filters params: %v", err)
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	voteList, err := checker.GetProposalList(cfg, req)
	if err != nil {
		Log(r).Errorf("failed get voting info list: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if len(voteList) == 0 {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	var resp resources.VotingInfoListResponse
	for _, vote := range voteList {
		id := strconv.FormatInt(vote.Contract.Config.ProposalId, 10)

		resp.Data = append(resp.Data, resources.VotingInfo{
			Key: resources.Key{
				ID:   id,
				Type: resources.VOTE_INFO,
			},
			Attributes: *vote,
		})
	}
	ape.Render(w, resp)
}
