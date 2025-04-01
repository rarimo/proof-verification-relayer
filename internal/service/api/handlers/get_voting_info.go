package handlers

import (
	"net/http"
	"strconv"

	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func GetVotingInfo(w http.ResponseWriter, r *http.Request) {
	cfg := Config(r)

	req, err := requests.GetVotingInfoRequest(r)
	if err != nil {
		Log(r).Errorf("failed get filters params: %v", err)
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	voteList, err := GetProposalList(cfg, req)
	if err != nil {
		Log(r).Errorf("failed get voting info list: %v", err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var resp resources.VotingInfoListResponse
	resp.Data = make([]resources.VotingInfo, 0)
	for _, vote := range voteList {
		id := strconv.FormatInt(vote.Contract.Config.ProposalId, 10)

		resp.Data = append(resp.Data, resources.VotingInfo{
			Key: resources.Key{
				ID:   id,
				Type: resources.VOTING_INFO,
			},
			Attributes: *vote,
		})
	}
	ape.Render(w, resp)
}

func GetProposalList(cfg config.Config, req requests.ProposalInfoFilter) ([]*resources.VotingInfoAttributes, error) {
	var result []*resources.VotingInfoAttributes
	pgDB := pg.NewVotingQ(cfg.DB().Clone()).
		FilterByCitizenship(req.CitizenshipList...).
		FilterByCreator(req.CreatorAddress...).
		FilterByMinAge(req.MinAge...).
		FilterByVotingId(req.ProposalId...)

	votingInfoList, err := pgDB.SelectVotingInfo()
	if err != nil {
		return nil, errors.Wrap(err, "failed to select voting info list")
	}
	for _, vote := range votingInfoList {
		result = append(result, &vote.ProposalInfoWithConfig)
	}

	return result, nil
}
