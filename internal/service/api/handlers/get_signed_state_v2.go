package handlers

import (
	"encoding/hex"
	"net/http"

	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetSignedStateV2(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetSignedStateRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	query := StateQ(r)
	if req.Block != nil {
		query = query.FilterByBlock(*req.Block)
	}
	if req.Root != nil {
		query = query.FilterByRoot(*req.Root)
	}

	state, err := query.SortByBlockHeight(data.DESC).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get state")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if state == nil {
		Log(r).Error("no block found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	signature, err := signState(*state, r)
	if err != nil {
		Log(r).WithError(err).Error("failed to sign state")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// See more about this here:
	// https://eips.ethereum.org/EIPS/eip-155#:~:text=If%20block.number%20%3E%3D%20FORK_BLKNUM%20and%20v,same%20rules%20as%20it%20did%20previously.
	// https://dev.to/truongpx396/understanding-ethereum-ecdsa-eip-712-and-its-role-in-permit-functionality-26ll
	signature[64] += 27

	ape.Render(w, resources.StateV2Response{
		Data: resources.StateV2{
			Key: resources.Key{
				ID:   state.ID,
				Type: resources.STATE_V2,
			},
			Attributes: resources.StateV2Attributes{
				Root:      state.Root,
				Signature: hex.EncodeToString(signature),
				Timestamp: int64(state.Timestamp),
			},
		},
	})
}
