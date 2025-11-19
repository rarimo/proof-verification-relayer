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
	idcardvoting "github.com/rarimo/proof-verification-relayer/internal/contracts/id_card_voting"
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
	UserPayload      []byte
	ProofBytes       []byte
}

type UserPayload struct {
	ProposalID *big.Int
	Vote       []*big.Int
	UserData   idcardvoting.BaseVotingUserData
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

	userPayload, err := parseUserPayload(calldataInfo.UserPayload)
	if err != nil {
		Log(r).WithError(err).Error("Failed User Payload")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"tx_data": errors.New("invalid user payload data format"),
		}.Filter())...)
		return
	}

	proposalID := userPayload.ProposalID.Int64()
	log := Log(r).WithFields(logan.F{
		"user-agent":  r.Header.Get("User-Agent"),
		"calldata":    calldata,
		"destination": destination,
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
		return config, fmt.Errorf("calldata is too short") //maybe !=4 and "calldata isn`t correct"
	}

	parsedABI, err := abi.JSON(strings.NewReader(idcardvoting.IDCardVotingABI))
	if err != nil {
		return config, fmt.Errorf("failed to parse contract ABI: %v", err)
	}

	method, ok := parsedABI.Methods["executeTD1Noir"]
	if !ok {
		return config, fmt.Errorf("method 'executeTD1Noir' not found in noir ABI")
	}

	decoded, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return config, fmt.Errorf("failed to unpack noir calldata: %v", err)
	}

	if len(decoded) != 4 {
		return config, fmt.Errorf("unexpected argument count, expected 4 got %d", len(decoded))
	}

	config.RegistrationRoot = decoded[0].([32]byte)
	config.CurrentDate = decoded[1].(*big.Int)
	config.UserPayload = decoded[2].([]byte)
	config.ProofBytes = decoded[3].([]byte)

	return config, nil
}

func parseUserPayload(data []byte) (UserPayload, error) {
	uint256Type, _ := abi.NewType("uint256", "", nil)
	uint256Array, _ := abi.NewType("uint256[]", "", nil)
	tupleType, _ := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "nullifier", Type: "uint256"},
		{Name: "citizenship", Type: "uint256"},
		{Name: "timestampUpperbound", Type: "uint256"},
	})

	arguments := abi.Arguments{
		{Type: uint256Type},
		{Type: uint256Array},
		{Type: tupleType},
	}

	decoded, err := arguments.Unpack(data)
	if err != nil {
		return UserPayload{}, err
	}

	if len(decoded) != 3 {
		return UserPayload{}, fmt.Errorf("invalid userDataEncoded structure")
	}

	proposalID := decoded[0].(*big.Int)
	vote := decoded[1].([]*big.Int)

	userDataRaw := decoded[2]

	userDataStruct, ok := userDataRaw.(struct {
		Nullifier           *big.Int `json:"nullifier"`
		Citizenship         *big.Int `json:"citizenship"`
		TimestampUpperbound *big.Int `json:"timestampUpperbound"`
	})
	if !ok {
		return UserPayload{}, fmt.Errorf("failed to cast userData to expected struct, got %T", userDataRaw)
	}

	userData := idcardvoting.BaseVotingUserData{
		Nullifier:                 userDataStruct.Nullifier,
		Citizenship:               userDataStruct.Citizenship,
		IdentityCreationTimestamp: userDataStruct.TimestampUpperbound,
	}

	userPayload := UserPayload{
		ProposalID: proposalID,
		Vote:       vote,
		UserData:   userData,
	}

	return userPayload, nil
}
