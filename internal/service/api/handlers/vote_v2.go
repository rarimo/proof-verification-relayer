package handlers

import (
	"bytes"
	"database/sql"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	biopassportvoting "github.com/rarimo/proof-verification-relayer/internal/contracts/bio_passport_voting"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	TRANSACTION resources.ResourceType = "transaction"
)

type txData struct {
	dataBytes []byte
	gasPrice  *big.Int
	gas       uint64
}

func isAddressInWhitelist(votingAddress common.Address, whitelist []common.Address) bool {
	votingAddressBytes := votingAddress.Bytes()
	for _, addr := range whitelist {
		if bytes.Equal(votingAddressBytes, addr.Bytes()) {
			return true
		}
	}
	return false
}

func VoteV2(w http.ResponseWriter, r *http.Request) {
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
	calldataInfo, err := parseCallData(txd.dataBytes)
	if err != nil {
		Log(r).WithError(err).Error("Failed parsed calldata")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"tx_data": errors.New("invalid transaction data format"),
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
	log.Debug("voting request")

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
		// `errors.Is` is not working for rpc errors, they passed as a string without additional wrapping
		// because of this we operate with raw strings
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

func confGas(r *http.Request, txd *txData, receiver *common.Address) (err error) {
	txd.gasPrice, err = VotingV2Config(r).RPC.SuggestGasPrice(r.Context())
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %w", err)
	}

	txd.gas, err = VotingV2Config(r).RPC.EstimateGas(r.Context(), ethereum.CallMsg{
		From:     crypto.PubkeyToAddress(VotingV2Config(r).PrivateKey.PublicKey),
		To:       receiver,
		GasPrice: txd.gasPrice,
		Data:     txd.dataBytes,
	})
	if err != nil {
		return fmt.Errorf("failed to estimate gas: %w", err)
	}

	return nil
}

func sendTx(r *http.Request, txd *txData, receiver *common.Address) (tx *types.Transaction, err error) {
	tx, err = signTx(r, txd, receiver)
	if err != nil {
		return nil, fmt.Errorf("failed to sign new tx: %w", err)
	}

	if err = VotingV2Config(r).RPC.SendTransaction(r.Context(), tx); err != nil {
		if strings.Contains(err.Error(), "nonce") {
			if err = VotingV2Config(r).ResetNonce(VotingV2Config(r).RPC); err != nil {
				return nil, fmt.Errorf("failed to reset nonce: %w", err)
			}

			tx, err = signTx(r, txd, receiver)
			if err != nil {
				return nil, fmt.Errorf("failed to sign new tx: %w", err)
			}

			if err := VotingV2Config(r).RPC.SendTransaction(r.Context(), tx); err != nil {
				return nil, fmt.Errorf("failed to send transaction: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to send transaction: %w", err)
		}
	}

	return tx, nil
}

func signTx(r *http.Request, txd *txData, receiver *common.Address) (tx *types.Transaction, err error) {
	tx, err = types.SignNewTx(
		VotingV2Config(r).PrivateKey,
		types.NewCancunSigner(VotingV2Config(r).ChainID),
		&types.LegacyTx{
			Nonce:    VotingV2Config(r).Nonce(),
			Gas:      txd.gas,
			GasPrice: txd.gasPrice,
			To:       receiver,
			Data:     txd.dataBytes,
		},
	)
	return tx, err
}

type VoteCalldata struct {
	RegistrationRoot [32]byte
	CurrentDate      *big.Int
	ProposalID       *big.Int
	Vote             []*big.Int
	UserData         biopassportvoting.BaseVotingUserData
	ZkPoints         biopassportvoting.VerifierHelperProofPoints
}

func parseCallData(data []byte) (VoteCalldata, error) {
	var config VoteCalldata
	if len(data) < 4 {
		return config, fmt.Errorf("calldata is too short")
	}

	parsedABI, err := abi.JSON(strings.NewReader(biopassportvoting.BioPassportVotingABI))
	if err != nil {
		return config, fmt.Errorf("failed to parse ABI: %v", err)
	}

	method, ok := parsedABI.Methods["vote"]
	if !ok {
		return config, fmt.Errorf("method 'vote' not found in ABI")
	}

	decoded, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return config, fmt.Errorf("failed to unpack calldata: %v", err)
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

	config.ZkPoints = biopassportvoting.VerifierHelperProofPoints{
		A: zkPointsStruct.A,
		B: zkPointsStruct.B,
		C: zkPointsStruct.C,
	}

	return config, nil
}

func updateVotingBalanceAndVotesCount(cfg config.Config, gasPrice *big.Int, gas uint64, votingId int64) error {
	pgDB := pg.NewVotingQ(cfg.DB().Clone()).FilterByVotingId(votingId)

	balance, _, err := pgDB.GetVotingBalance()
	if err != nil {
		return fmt.Errorf("failed get voting info from db: %w", err)
	}
	if balance == nil {
		return errors.New("Voting Not Found")
	}

	if gasPrice.Int64() == 0 {
		gasPrice = big.NewInt(1)
	}

	var votes_count int64
	if err := pgDB.Get("votes_count", &votes_count); err != nil {
		if err != sql.ErrNoRows {
			return errors.Wrap(err, "failed to get voting info from db")
		}
		return errors.New("voting Not Found")
	}

	err = pgDB.Update(map[string]any{
		"residual_balance": new(big.Int).Sub(balance, new(big.Int).Mul(big.NewInt(int64(gas)), gasPrice)).String(),
		"votes_count":      votes_count + 1,
	})
	if err != nil {
		return fmt.Errorf("failed update voting balance: %w", err)
	}

	return nil
}

func checkUpdateGasLimit(value uint64, cfg config.Config, votingId int64) error {
	pgDB := pg.NewVotingQ(cfg.DB().Clone()).FilterByVotingId(votingId)

	var one int
	if err := pgDB.Get("1", &one); err != nil {
		if err != sql.ErrNoRows {
			return errors.Wrap(err, "failed to get voting info from db")
		}
		return errors.New("voting Not Found")
	}

	if err := pgDB.Update(map[string]any{"gas_limit": value}); err != nil {
		return errors.Wrap(err, "failed to update voting gas limit from db")
	}
	return nil
}
