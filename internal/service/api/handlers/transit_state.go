package handlers

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
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
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"tx_data": errors.Wrap(err, "failed to decode encoded tx data"),
		})...)
		return
	}

	newStatesRoot, gistData, proof, err := getSignedTransitStateTxDataParams(r, dataBytes)
	if err != nil {
		Log(r).WithError(err).Error("failed to get tx data params")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	NetworkConfig(r).LockNonce()
	defer NetworkConfig(r).UnlockNonce()
	txOpts, err := getTxOpts(r, NetworkConfig(r).LightweightState, dataBytes)
	if err != nil {
		Log(r).WithError(err).Error("failed to get transaction options")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tx, err := LightweightState(r).SignedTransitState(txOpts, newStatesRoot, gistData, proof)
	if err != nil {
		Log(r).WithError(err).Error("failed to make call to signed transit state")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	NetworkConfig(r).IncrementNonce()

	ape.Render(w, newTxModel(tx))
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

func getTxOpts(r *http.Request, receiver common.Address, txData []byte) (*bind.TransactOpts, error) {
	gasPrice, err := EthClient(r).SuggestGasPrice(r.Context())
	if err != nil {
		return nil, errors.Wrap(err, "failed to suggest gas price")
	}
	gasPrice = multiplyGasPrice(gasPrice, NetworkConfig(r).GasMultiplier)

	gasLimit, err := EthClient(r).EstimateGas(r.Context(), ethereum.CallMsg{
		From:     crypto.PubkeyToAddress(NetworkConfig(r).PrivateKey.PublicKey),
		To:       &receiver,
		GasPrice: gasPrice,
		Data:     txData,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to estimate gas limit")
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(NetworkConfig(r).PrivateKey, NetworkConfig(r).ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a transaction signer")
	}

	txOpts.Nonce = new(big.Int).SetUint64(NetworkConfig(r).Nonce())
	txOpts.GasPrice = gasPrice
	txOpts.GasLimit = uint64(float64(gasLimit) * NetworkConfig(r).GasMultiplier)

	return txOpts, nil
}

func newTxModel(tx *types.Transaction) resources.Tx {
	return resources.Tx{
		Key: resources.Key{
			ID:   tx.Hash().String(),
			Type: resources.TXS,
		},
		Attributes: resources.TxAttributes{
			TxHash: tx.Hash().String(),
		},
	}
}
