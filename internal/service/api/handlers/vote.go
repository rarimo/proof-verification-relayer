package handlers

import (
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func Vote(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVoteRequest(r)
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

	registration := common.HexToAddress(req.Data.Registration)
	voting := common.HexToAddress(req.Data.Voting)

	exists, err := VotingRegistry(r).IsPoolExistByProposer(&bind.CallOpts{}, NetworkConfig(r).Proposer, registration)
	if err != nil {
		Log(r).WithError(err).Error("failed to check if registration exists by proposer")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if !exists {
		Log(r).WithField("registration", registration.String()).Error("registration does not exist")
		ape.RenderErr(w, problems.BadRequest(errors.New("registration does not exist"))...)
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

	gas, err := EthClient(r).EstimateGas(r.Context(), ethereum.CallMsg{
		From:     crypto.PubkeyToAddress(NetworkConfig(r).PrivateKey.PublicKey),
		To:       &voting,
		GasPrice: gasPrice,
		Data:     dataBytes,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to estimate gas")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tx, err := types.SignNewTx(
		NetworkConfig(r).PrivateKey,
		types.NewCancunSigner(NetworkConfig(r).ChainID),
		&types.LegacyTx{
			Nonce:    NetworkConfig(r).Nonce(),
			Gas:      uint64(float64(gas) * NetworkConfig(r).GasMultiplier),
			GasPrice: gasPrice,
			To:       &voting,
			Data:     dataBytes,
		},
	)
	if err != nil {
		Log(r).WithError(err).Error("failed to sign new tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if err := EthClient(r).SendTransaction(r.Context(), tx); err != nil {
		if strings.Contains(err.Error(), "nonce") {
			if err := NetworkConfig(r).ResetNonce(EthClient(r)); err != nil {
				Log(r).WithError(err).Error("failed to reset nonce")
				ape.RenderErr(w, problems.InternalError())
				return
			}

			tx, err = types.SignNewTx(
				NetworkConfig(r).PrivateKey,
				types.NewCancunSigner(NetworkConfig(r).ChainID),
				&types.LegacyTx{
					Nonce:    NetworkConfig(r).Nonce(),
					Gas:      gas,
					GasPrice: gasPrice,
					To:       &voting,
					Data:     dataBytes,
				},
			)
			if err != nil {
				Log(r).WithError(err).Error("failed to sign new tx")
				ape.RenderErr(w, problems.InternalError())
				return
			}

			if err := EthClient(r).SendTransaction(r.Context(), tx); err != nil {
				Log(r).WithError(err).Error("failed to send transaction")
				ape.RenderErr(w, problems.InternalError())
				return
			}
		} else {
			Log(r).WithError(err).Error("failed to send transaction")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	NetworkConfig(r).IncrementNonce()

	ape.Render(w, newTxModel(tx))
}

// ONE - One GWEI
var ONE = 1000000000

func multiplyGasPrice(gasPrice *big.Int, multiplier float64) *big.Int {
	mult := big.NewFloat(0).Mul(big.NewFloat(multiplier), big.NewFloat(float64(ONE)))
	gas, _ := big.NewFloat(0).Mul(big.NewFloat(0).SetInt(gasPrice), mult).Int(nil)
	return big.NewInt(0).Div(gas, big.NewInt(int64(ONE)))
}
