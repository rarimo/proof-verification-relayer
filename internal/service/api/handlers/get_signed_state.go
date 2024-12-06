package handlers

import (
	"encoding/hex"
	"math/big"
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
		root := *req.Root
		if root[:2] == "0x" {
			root = root[2:]
		}
		query = query.FilterByRoot(root)
	}

	state, err := query.SortByBlockHeight(data.DESC).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get state")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if state == nil {
		Log(r).Error("no state found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	signature, err := signState(*state, r)
	if err != nil {
		Log(r).WithError(err).Error("failed to sign state")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.State{
		Key: resources.Key{
			ID:   state.ID,
			Type: resources.STATE,
		},
		Attributes: resources.StateAttributes{
			Signature: hex.EncodeToString(signature),
			Timestamp: int64(state.Timestamp),
		},
	})
}

func signState(state data.State, r *http.Request) ([]byte, error) {
	rootBytes, err := hex.DecodeString(state.Root)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode signature digest", logan.F{"root": state.Root})
	}

	//keccak256(abi.encodePacked(
	//	REGISTRATION_ROOT_PREFIX,
	//	sourceSMT,
	//	address(this),
	//	newRoot_,
	//	transitionTimestamp_
	//));
	digest := crypto.Keccak256(
		[]byte(Config(r).Replicator().RootPrefix),
		Config(r).Replicator().SourceSMT.Bytes(),
		Config(r).Replicator().Address.Bytes(),
		rootBytes,
		new(big.Int).SetUint64(state.Timestamp).Bytes(),
	)

	signature, err := crypto.Sign(digest, Config(r).NetworkConfig().PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign state")
	}

	return signature, nil
}
