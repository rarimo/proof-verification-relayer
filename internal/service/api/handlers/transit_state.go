package handlers

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func TransitState(w http.ResponseWriter, r *http.Request) {
	req, err := requests.TransitStateDataRequestRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	dataBytes, err := hexutil.Decode(req.Data.TxData)
	if err != nil {
		Log(r).WithError(err).Error("failed to decode data")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	newStatesRoot, gistData, proof, err := getSignedTransitStateTxDataParams(r, dataBytes)
	if err != nil {
		Log(r).WithError(err).Error("failed to get tx data params")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	gasPrice, err := EthClient(r).SuggestGasPrice(r.Context())
	if err != nil {
		Log(r).WithError(err).Error("failed to suggest gas price")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	gasPrice = multiplyGasPrice(gasPrice, NetworkConfig(r).GasMultiplier)

	NetworkConfig(r).LockNonce()
	defer NetworkConfig(r).UnlockNonce()

	tx, err := LightweightState(r).SignedTransitState(&bind.TransactOpts{
		Nonce:    new(big.Int).SetUint64(NetworkConfig(r).Nonce()),
		GasPrice: gasPrice,
	}, newStatesRoot, gistData, proof)

	NetworkConfig(r).IncrementNonce()

	ape.Render(w, resources.Tx{
		Key: resources.Key{
			ID:   tx.Hash().String(),
			Type: resources.TXS,
		},
		Attributes: resources.TxAttributes{
			TxHash: tx.Hash().String(),
		},
	})
}

func getSignedTransitStateTxDataParams(r *http.Request, data []byte) ([32]byte, contracts.ILightweightStateGistRootData, []byte, error) {
	gistRootData := contracts.ILightweightStateGistRootData{}
	var newStateRoot [32]byte

	unpackResult, err := SignedTransitStateMethod(r).Inputs.Unpack(data[4:])
	if err != nil {
		return newStateRoot, gistRootData, nil, errors.Wrap(err, "failed to unpack tx data")
	}

	if len(unpackResult) != 3 {
		return newStateRoot, gistRootData, nil, errors.Wrap(err, "unpack result is not valid")
	}

	newStateRoot, ok := unpackResult[0].([32]byte)
	if !ok {
		return newStateRoot, gistRootData, nil, errors.New("failed to convert newIdentitiesStatesRoot_ interface to [32]byte")
	}

	gistDataJSON, err := json.Marshal(unpackResult[1])
	if err != nil {
		return newStateRoot, gistRootData, nil, errors.Wrap(err, "failed to marshal JSON")
	}

	if err := json.Unmarshal(gistDataJSON, &gistRootData); err != nil {
		return newStateRoot, gistRootData, nil, errors.Wrap(err, "failed to unmarshal JSON")
	}

	lightweightStateProof, ok := unpackResult[2].([]byte)
	if !ok {
		return newStateRoot, gistRootData, nil, errors.New("failed to convert lightweightStateProof interface to []byte")
	}

	return newStateRoot, gistRootData, lightweightStateProof, nil
}
