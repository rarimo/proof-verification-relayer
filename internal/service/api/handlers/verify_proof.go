package handlers

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func VerifyProof(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVerifyProofRequest(r)
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

	gasPrice, err := EthClient(r).SuggestGasPrice(r.Context())
	if err != nil {
		Log(r).WithError(err).Error("failed to suggest gas price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	verifierAddress := NetworkConfig(r).VerifierAddress

	gas, err := EthClient(r).EstimateGas(r.Context(), ethereum.CallMsg{
		From:     crypto.PubkeyToAddress(NetworkConfig(r).PrivateKey.PublicKey),
		To:       &verifierAddress,
		GasPrice: gasPrice,
		Data:     dataBytes,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to estimate gas")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	NetworkConfig(r).LockNonce()
	defer NetworkConfig(r).UnlockNonce()

	if _, err := types.SignNewTx(
		NetworkConfig(r).PrivateKey,
		types.NewCancunSigner(NetworkConfig(r).ChainID),
		&types.LegacyTx{
			Nonce:    NetworkConfig(r).Nonce(),
			Gas:      gas,
			GasPrice: gasPrice,
			To:       &verifierAddress,
			Data:     dataBytes,
		},
	); err != nil {
		Log(r).WithError(err).Error("failed to sign new tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if err := MasterQ(r).Proofs().Insert(data.Proof{
		TxData: req.Data.TxData,
	}); err != nil {
		Log(r).WithError(err).Error("failed to insert proof into the database")
		ape.RenderErr(w, problems.InternalError())
		return
	}
}
