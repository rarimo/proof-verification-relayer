package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type OffsetParams struct {
	pgdb.OffsetPageParams
}

type ProposalInfoFilter struct {
	CreatorAddress  []string `filter:"creator"`
	CitizenshipList []string `filter:"citizenship"`
	MinAge          []int64  `filter:"min_age"`
	MaxAge          []int64  `filter:"max_age"`
	ProposalId      []int64  `filter:"proposal_id"`
	VotingStatus    []string `filter:"status"`
	SexFilter       []string `filter:"sex"`

	OffsetParams
}

func ProposalsInfoRequest(r *http.Request) (ProposalInfoFilter, error) {
	var request ProposalInfoFilter
	err := urlval.DecodeSilently(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
