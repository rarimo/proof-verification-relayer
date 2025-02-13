package handlers

import (
	"bytes"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/rarimo/proof-verification-relayer/internal/checker"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"

	// Instead of "proposalsstate", use the package
	// that will be generated for the required contract.
	"github.com/rarimo/voting-relayer/pkg/contracts/proposalsstate"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	OPERATION   resources.ResourceType = "operation"
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

func Voting(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewVotingRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to get request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var (
		destination = req.Data.Attributes.Destination
		calldata    = req.Data.Attributes.TxData
		proposalID  = req.Data.Attributes.ProposalId
	)

	log := Log(r).WithFields(logan.F{
		"user-agent":  r.Header.Get("User-Agent"),
		"calldata":    calldata,
		"destination": destination,
		"proposal_id": proposalID,
	})
	log.Debug("voting request")

	// destination is valid hex address because of request validation
	votingAddress := common.HexToAddress(destination)

	var txd txData
	txd.dataBytes, err = hexutil.Decode(calldata)
	if err != nil {
		log.WithError(err).Error("Failed to decode data")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	proposalBigID := big.NewInt(proposalID)

	// Instead of "proposalsstate", use the package
	// that will be generated for the required contract.
	session, err := proposalsstate.NewProposalsStateCaller(RelayerConfig(r).Address, RelayerConfig(r).RPC)

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

	defer RelayerConfig(r).UnlockNonce()
	RelayerConfig(r).LockNonce()

	err = confGas(r, &txd, &votingAddress)

	if err != nil {
		log.WithError(err).Error("Failed to configure gas and gasPrice")
		// `errors.Is` is not working for rpc errors, they passed as a string without additional wrapping
		// because of this we operate with raw strings
		if strings.Contains(err.Error(), vm.ErrExecutionReverted.Error()) {
			errParts := strings.Split(err.Error(), ":")
			contractName := strings.TrimSpace(errParts[len(errParts)-2])
			errMsg := errors.New(strings.TrimSpace(errParts[len(errParts)-1]))
			ape.RenderErr(w, problems.BadRequest(validation.Errors{contractName: errMsg}.Filter())...)
			return
		}
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tx, err := sendTx(r, &txd, &votingAddress)
	if err != nil {
		log.WithError(err).Error("failed to send tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	RelayerConfig(r).IncrementNonce()

	ape.Render(w, resources.Relation{
		Data: &resources.Key{
			ID:   tx.Hash().String(),
			Type: TRANSACTION,
		},
	})
}

func confGas(r *http.Request, txd *txData, receiver *common.Address) (err error) {
	txd.gasPrice, err = RelayerConfig(r).RPC.SuggestGasPrice(r.Context())
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %w", err)
	}

	txd.gas, err = RelayerConfig(r).RPC.EstimateGas(r.Context(), ethereum.CallMsg{
		From:     crypto.PubkeyToAddress(RelayerConfig(r).PrivateKey.PublicKey),
		To:       receiver,
		GasPrice: txd.gasPrice,
		Data:     txd.dataBytes,
	})
	if err != nil {
		return fmt.Errorf("failed to estimate gas: %w", err)
	}

	err = checker.CheckUpdateGasLimit(txd.gas, Config(r), receiver)

	return nil
}

func sendTx(r *http.Request, txd *txData, receiver *common.Address) (tx *types.Transaction, err error) {
	tx, err = signTx(r, txd, receiver)
	if err != nil {
		return nil, fmt.Errorf("failed to sign new tx: %w", err)
	}

	if err = RelayerConfig(r).RPC.SendTransaction(r.Context(), tx); err != nil {
		if strings.Contains(err.Error(), "nonce") {
			if err = RelayerConfig(r).ResetNonce(RelayerConfig(r).RPC); err != nil {
				return nil, fmt.Errorf("failed to reset nonce: %w", err)
			}

			tx, err = signTx(r, txd, receiver)
			if err != nil {
				return nil, fmt.Errorf("failed to sign new tx: %w", err)
			}

			if err := RelayerConfig(r).RPC.SendTransaction(r.Context(), tx); err != nil {
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
		RelayerConfig(r).PrivateKey,
		types.NewCancunSigner(RelayerConfig(r).ChainID),
		&types.LegacyTx{
			Nonce:    RelayerConfig(r).Nonce(),
			Gas:      txd.gas,
			GasPrice: txd.gasPrice,
			To:       receiver,
			Data:     txd.dataBytes,
		},
	)
	return tx, err
}
