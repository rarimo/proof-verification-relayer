package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type ProposalInfoFilter struct {
	CreatorAddress  []string `filter:"creator"`
	CitizenshipList []string `filter:"citizenship"`
	MinAge          []int64  `filter:"age"`
	ProposalId      []int64  `filter:"proposal_id"`
}

func GetVotesRequest(r *http.Request) (ProposalInfoFilter, error) {
	var request ProposalInfoFilter
	err := urlval.DecodeSilently(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
