package handlers

import (
	"database/sql"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	biopassportvoting "github.com/rarimo/proof-verification-relayer/internal/contracts/bio_passport_voting"
	noirvoting "github.com/rarimo/proof-verification-relayer/internal/contracts/noir_voting"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type NoirVoteCalldata struct {
	RegistrationRoot [32]byte
	CurrentDate      *big.Int
	ProposalID       *big.Int
	Vote             []*big.Int
	UserData         biopassportvoting.BaseVotingUserData
	ZkPoints         noirvoting.AQueryProofExecutorProofPoints
}

func VoteV3(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVotingRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var (
		destination = req.Data.Attributes.Destination
		calldata    = req.Data.Attributes.TxData
	)

	var txd txData
	txd.dataBytes, err = hexutil.Decode(calldata)
	if err != nil {
		Log(r).WithError(err).Error("Failed to decode data")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	calldataInfo, err := parseNoirCallData(txd.dataBytes)
	if err != nil {
		Log(r).WithError(err).Error("Failed parsed calldata")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"tx_data": errors.New("invalid noir transaction data format"),
		}.Filter())...)
		return
	}

	proposalID := calldataInfo.ProposalID.Int64()
	log := Log(r).WithFields(logan.F{
		"user-agent":  r.Header.Get("User-Agent"),
		"calldata":    calldata,
		"destination": destination,
		"proposal_id": proposalID,
	})
	log.Debug("noir voting request")

	// destination is valid hex address because of request validation
	votingAddress := common.HexToAddress(destination)
	proposalBigID := big.NewInt(proposalID)
	session, err := contracts.NewProposalsStateCaller(VotingV2Config(r).Address, VotingV2Config(r).RPC)
	if err != nil {
		log.WithError(err).Error("Failed to get proposal state caller")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	proposalConfig, err := session.GetProposalConfig(nil, proposalBigID)
	if err != nil {
		log.WithError(err).Error("Failed to get proposal config")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if !isAddressInWhitelist(votingAddress, proposalConfig.VotingWhitelist) {
		log.Error("Address not in voting whitelist")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	defer VotingV2Config(r).UnlockNonce()
	VotingV2Config(r).LockNonce()

	err = confGas(r, &txd, &votingAddress)
	if err != nil {
		log.WithError(err).Error("Failed to configure gas and gasPrice")
		if strings.Contains(strings.ToLower(err.Error()), strings.ToLower(vm.ErrExecutionReverted.Error())) {
			errParts := strings.Split(err.Error(), ":")
			contractName := strings.TrimSpace(errParts[len(errParts)-2])
			errMsg := errors.New(strings.TrimSpace(errParts[len(errParts)-1]))
			ape.RenderErr(w, problems.BadRequest(validation.Errors{contractName: errMsg}.Filter())...)
			return
		}
		ape.RenderErr(w, problems.InternalError())
		return
	}

	err = checkUpdateGasLimit(txd.gas, Config(r), proposalBigID.Int64())
	if err != nil {
		log.WithError(err).Error("failed to check and update gas limit")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	checkIsEnough, err := isEnough(Config(r), proposalID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.WithError(err).Warn("Insufficient balance check failed")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if !checkIsEnough {
		log.Info("Insufficient funds in the voting account")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	err = updateVotingBalanceAndVotesCount(Config(r), txd.gasPrice, txd.gas, proposalID)
	if err != nil {
		log.WithError(err).Error("failed update voting balance")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tx, err := sendTx(r, &txd, &votingAddress)
	if err != nil {
		log.WithError(err).Error("failed to send tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	VotingV2Config(r).IncrementNonce()

	ape.Render(w, resources.Relation{
		Data: &resources.Key{
			ID:   tx.Hash().String(),
			Type: TRANSACTION,
		},
	})
}

func parseNoirCallData(data []byte) (NoirVoteCalldata, error) {
	var config NoirVoteCalldata
	if len(data) < 4 {
		return config, fmt.Errorf("calldata is too short")
	}

	parsedABI, err := abi.JSON(strings.NewReader(noirvoting.NoirVotingABI))
	if err != nil {
		return config, fmt.Errorf("failed to parse noir ABI: %v", err)
	}

	method, ok := parsedABI.Methods["executeNoir"]
	if !ok {
		return config, fmt.Errorf("method 'executeNoir' not found in noir ABI")
	}

	decoded, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return config, fmt.Errorf("failed to unpack noir calldata: %v", err)
	}

	if len(decoded) != 6 {
		return config, fmt.Errorf("unexpected argument count, expected 6 got %d", len(decoded))
	}

	config.RegistrationRoot = decoded[0].([32]byte)
	config.CurrentDate = decoded[1].(*big.Int)
	config.ProposalID = decoded[2].(*big.Int)
	config.Vote = decoded[3].([]*big.Int)
	userDataRaw := decoded[4]

	userDataStruct, ok := userDataRaw.(struct {
		Nullifier                 *big.Int `json:"nullifier"`
		Citizenship               *big.Int `json:"citizenship"`
		IdentityCreationTimestamp *big.Int `json:"identityCreationTimestamp"`
	})
	if !ok {
		return config, fmt.Errorf("failed to cast userData_ to expected struct, got %T", userDataRaw)
	}

	config.UserData = biopassportvoting.BaseVotingUserData{
		Nullifier:                 userDataStruct.Nullifier,
		Citizenship:               userDataStruct.Citizenship,
		IdentityCreationTimestamp: userDataStruct.IdentityCreationTimestamp,
	}

	zkPointsRaw := decoded[5]
	zkPointsStruct, ok := zkPointsRaw.(struct {
		A [2]*big.Int    `json:"a"`
		B [2][2]*big.Int `json:"b"`
		C [2]*big.Int    `json:"c"`
	})
	if !ok {
		return config, fmt.Errorf("failed to cast zkPoints_ to expected struct, got %T", zkPointsRaw)
	}

	config.ZkPoints = noirvoting.AQueryProofExecutorProofPoints{
		A: zkPointsStruct.A,
		B: zkPointsStruct.B,
		C: zkPointsStruct.C,
	}

	return config, nil
}
