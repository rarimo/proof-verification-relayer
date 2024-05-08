package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func Register(w http.ResponseWriter, r *http.Request) {
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

	registration, registerProofParams, err := getRegistrationTxDataParams(r, dataBytes)
	if err != nil {
		Log(r).WithError(err).Error("failed to get tx data params")
		ape.RenderErr(w, problems.InternalError())
		return
	}

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

	voteVerifier, err := contracts.NewVoteVerifier(registration, EthClient(r))
	if err != nil {
		Log(r).WithError(err).Error("failed to initialize new vote verifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	isRegistered, err := voteVerifier.IsUserRegistered(&bind.CallOpts{}, registerProofParams.DocumentNullifier)
	if err != nil {
		Log(r).WithError(err).Error("failed to check if user is registered by document nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if isRegistered {
		Log(r).WithFields(logan.F{
			"registration":       registration.Hex(),
			"document_nullifier": registerProofParams.DocumentNullifier.String(),
		}).Error("too many requests for registration and document nullifier")
		ape.RenderErr(w, problems.TooManyRequests())
		return
	}

	gasPrice, err := EthClient(r).SuggestGasPrice(r.Context())
	if err != nil {
		Log(r).WithError(err).Error("failed to suggest gas price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	NetworkConfig(r).LockNonce()
	defer NetworkConfig(r).UnlockNonce()

	gas, err := EthClient(r).EstimateGas(r.Context(), ethereum.CallMsg{
		From:     crypto.PubkeyToAddress(NetworkConfig(r).PrivateKey.PublicKey),
		To:       &registration,
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
			GasPrice: multiplyGasPrice(gasPrice, NetworkConfig(r).GasMultiplier),
			To:       &registration,
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
					To:       &registration,
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

func getRegistrationTxDataParams(r *http.Request, data []byte) (
	common.Address, contracts.IRegisterVerifierRegisterProofParams, error,
) {
	unpackResult, err := VoteVerifierRegisterMethod(r).Inputs.Unpack(data[4:])
	if err != nil {
		return common.Address{}, contracts.IRegisterVerifierRegisterProofParams{}, errors.Wrap(err, "failed to unpack tx data")
	}

	if len(unpackResult) != 4 {
		return common.Address{}, contracts.IRegisterVerifierRegisterProofParams{}, errors.Wrap(err, "unpack result is not valid")
	}

	proveIdentityParamsJSON, err := json.Marshal(unpackResult[0])
	if err != nil {
		return common.Address{}, contracts.IRegisterVerifierRegisterProofParams{}, errors.Wrap(err, "failed to marshal JSON")
	}
	proveIdentityParams := contracts.IBaseVerifierProveIdentityParams{}
	if err := json.Unmarshal(proveIdentityParamsJSON, &proveIdentityParams); err != nil {
		return common.Address{}, contracts.IRegisterVerifierRegisterProofParams{}, errors.Wrap(err, "failed to unmarshal JSON")
	}

	registrationAddressBytes := proveIdentityParams.Inputs[len(proveIdentityParams.Inputs)-2].Bytes()
	registration := common.BytesToAddress(registrationAddressBytes)

	registerProofParamsJSON, err := json.Marshal(unpackResult[1])
	if err != nil {
		return common.Address{}, contracts.IRegisterVerifierRegisterProofParams{}, errors.Wrap(err, "failed to marshal JSON")
	}
	registerProofParams := contracts.IRegisterVerifierRegisterProofParams{}
	if err := json.Unmarshal(registerProofParamsJSON, &registerProofParams); err != nil {
		return common.Address{}, contracts.IRegisterVerifierRegisterProofParams{}, errors.Wrap(err, "failed to unmarshal JSON")
	}

	return registration, registerProofParams, nil
}
