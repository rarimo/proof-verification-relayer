package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/proof-verification-relayer/internal/config"
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

	voteVerifier, err := contracts.NewVoteVerifier(registration, Config(r).NetworkConfig().Client)
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

	Config(r).NetworkConfig().LockNonce()
	defer Config(r).NetworkConfig().UnlockNonce()

	tx, err := processLegacyTx(r, registration, dataBytes)
	if err != nil {
		Log(r).WithError(err).Error("failed to process legacy transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	Config(r).NetworkConfig().IncrementNonce()

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
