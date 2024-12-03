package handlers

import (
	"encoding/hex"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func GetSignedState(w http.ResponseWriter, r *http.Request) {
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

	ape.Render(w, resources.Signature{
		Key: resources.NewKeyInt64(state.ID, resources.STATE),
		Attributes: resources.SignatureAttributes{
			Signature: hex.EncodeToString(signature),
		},
	})
}

func signState(state data.State, r *http.Request) ([]byte, error) {
	digest, err := hex.DecodeString(state.Root)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode signature digest", logan.F{"root": state.Root})
	}

	signature, err := crypto.Sign(digest, Config(r).NetworkConfig().PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign state")
	}

	return signature, nil
}
