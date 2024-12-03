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
	"github.com/rarimo/proof-verification-relayer/internal/config"
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

	exists, err := VotingRegistry(r).IsPoolExistByProposer(
		&bind.CallOpts{},
		Config(r).ContractsConfig()[config.Proposer].Address,
		registration,
	)
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

	Config(r).NetworkConfig().LockNonce()
	defer Config(r).NetworkConfig().UnlockNonce()

	tx, err := processLegacyTx(r, voting, dataBytes)
	if err != nil {
		Log(r).WithError(err).Error("failed to process legacy transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	Config(r).NetworkConfig().IncrementNonce()

	ape.Render(w, newTxModel(tx))
}

// ONE - One GWEI
var ONE = 1000000000

func multiplyGasPrice(gasPrice *big.Int, multiplier float64) *big.Int {
	mult := big.NewFloat(0).Mul(big.NewFloat(multiplier), big.NewFloat(float64(ONE)))
	gas, _ := big.NewFloat(0).Mul(big.NewFloat(0).SetInt(gasPrice), mult).Int(nil)
	return big.NewInt(0).Div(gas, big.NewInt(int64(ONE)))
}

func processLegacyTx(r *http.Request, receiver common.Address, txData []byte) (*types.Transaction, error) {
	gasPrice, gasLimit, err := getGasCosts(r, receiver, txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get gas costs")
	}

	tx, err := types.SignNewTx(
		Config(r).NetworkConfig().PrivateKey,
		types.NewCancunSigner(Config(r).NetworkConfig().ChainID),
		&types.LegacyTx{
			Nonce:    Config(r).NetworkConfig().Nonce(),
			Gas:      gasLimit,
			GasPrice: gasPrice,
			To:       &receiver,
			Data:     txData,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign new tx")
	}

	if err = sendTransaction(r, tx); err != nil {
		return nil, errors.Wrap(err, "failed to send transaction")
	}

	return tx, nil
}

func getGasCosts(
	r *http.Request,
	receiver common.Address,
	txData []byte,
) (gasPrice *big.Int, gasLimit uint64, err error) {
	gasPrice, err = Config(r).NetworkConfig().Client.SuggestGasPrice(r.Context())
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to suggest gas price")
	}
	gasPrice = multiplyGasPrice(gasPrice, Config(r).NetworkConfig().GasMultiplier)

	gasLimit, err = Config(r).NetworkConfig().Client.EstimateGas(r.Context(), ethereum.CallMsg{
		From:     crypto.PubkeyToAddress(Config(r).NetworkConfig().PrivateKey.PublicKey),
		To:       &receiver,
		GasPrice: gasPrice,
		Data:     txData,
	})
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to estimate gas costs limit")
	}
	gasLimit = uint64(float64(gasLimit) * Config(r).NetworkConfig().GasMultiplier)

	return
}

func sendTransaction(r *http.Request, tx *types.Transaction) error {
	if err := Config(r).NetworkConfig().Client.SendTransaction(r.Context(), tx); err != nil {
		if strings.Contains(err.Error(), "nonce") {
			if err = Config(r).NetworkConfig().ResetNonce(Config(r).NetworkConfig().Client); err != nil {
				return errors.Wrap(err, "failed to reset nonce")
			}

			tx, err = types.SignNewTx(
				Config(r).NetworkConfig().PrivateKey,
				types.NewCancunSigner(Config(r).NetworkConfig().ChainID),
				&types.LegacyTx{
					Nonce:    Config(r).NetworkConfig().Nonce(),
					Gas:      tx.Gas(),
					GasPrice: tx.GasPrice(),
					To:       tx.To(),
					Data:     tx.Data(),
				},
			)
			if err != nil {
				return errors.Wrap(err, "failed to sign new tx")
			}

			if err := Config(r).NetworkConfig().Client.SendTransaction(r.Context(), tx); err != nil {
				return errors.Wrap(err, "failed to resend transaction")
			}
		} else {
			return errors.Wrap(err, "failed to send transaction")
		}
	}

	return nil
}
