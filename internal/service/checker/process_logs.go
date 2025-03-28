package checker

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (ch *checker) processLog(vLog types.Log, eventName string) error {
	var processedEvent data.ProcessedEvent

	block := int64(vLog.BlockNumber)
	processedEvent.BlockNumber = block
	processedEvent.LogIndex = int64(vLog.Index)
	processedEvent.TransactionHash = vLog.TxHash[:]
	err := ch.insertProcessedEventLog(processedEvent)
	if err != nil {
		return errors.Wrap(err, "failed insert processed event")
	}
	sender, err := ch.getSender(vLog.TxHash)
	if err != nil {
		return errors.Wrap(err, "failed get processed event creator")
	}
	votingId, value, proposalInfoWithConfig, err := ch.getEventData(eventName, vLog, sender)
	if err != nil {
		return errors.Wrap(err, "failed get event data")
	}

	votingInfo, err := ch.checkVotingAndGetBalance(votingId, sender)
	if err != nil {
		return errors.Wrap(err, "failed get voting info", logan.F{"Voting ID": votingId})
	}

	if value != nil {
		votingInfo.Balance = new(big.Int).Add(votingInfo.Balance, value)
	}
	if proposalInfoWithConfig != nil {
		votingInfo.ProposalInfoWithConfig = *proposalInfoWithConfig
	}

	err = ch.votingQ.UpdateVotingInfo(votingInfo)
	if err != nil {
		return errors.Wrap(err, "failed update voting balance", logan.F{"Voting ID": votingId})
	}
	return nil
}

func (ch *checker) getEventData(eventName string, vLog types.Log, sender string) (
	votingId int64,
	value *big.Int,
	proposalInfoWithConfig *resources.VotingInfoAttributes,
	err error) {
	parsedABI, err := abi.JSON(strings.NewReader(contracts.ProposalsStateABI))
	if err != nil {
		return 0, nil, nil, errors.Wrap(err, "failed to parse contract ABI")
	}

	switch eventName {
	case "ProposalCreated":
		var transferEvent contracts.ProposalsStateProposalCreated
		err = parsedABI.UnpackIntoInterface(&transferEvent, eventName, vLog.Data)
		if err != nil {
			return 0, nil, nil, errors.Wrap(err, "failed to unpack log data")
		}

		transferEvent.ProposalId = new(big.Int).SetBytes(vLog.Topics[1].Bytes())
		transferEvent.Raw = vLog

		proposalInfoWithConfig, err = GetProposalInfo(transferEvent.ProposalId.Int64(), ch.cfg, sender)
		if err != nil {
			return 0, nil, nil, errors.Wrap(err, "failed get proposal info")
		}

		return transferEvent.ProposalId.Int64(), transferEvent.FundAmount, proposalInfoWithConfig, nil
	case "ProposalFunded":
		var transferEvent contracts.ProposalsStateProposalFunded
		err = parsedABI.UnpackIntoInterface(&transferEvent, eventName, vLog.Data)
		if err != nil {
			return 0, nil, nil, errors.Wrap(err, "failed to unpack log data")
		}

		transferEvent.ProposalId = new(big.Int).SetBytes(vLog.Topics[1].Bytes())
		transferEvent.Raw = vLog

		return transferEvent.ProposalId.Int64(), transferEvent.FundAmount, nil, nil
	case "ProposalConfigChanged":
		var transferEvent contracts.ProposalsStateProposalConfigChanged
		err = parsedABI.UnpackIntoInterface(&transferEvent, eventName, vLog.Data)
		if err != nil {
			return 0, nil, nil, errors.Wrap(err, "failed to unpack log data")
		}

		transferEvent.ProposalId = new(big.Int).SetBytes(vLog.Topics[1].Bytes())
		transferEvent.Raw = vLog

		proposalInfoWithConfig, err = GetProposalInfo(transferEvent.ProposalId.Int64(), ch.cfg, sender)
		if err != nil {
			return 0, nil, nil, errors.Wrap(err, "failed get proposal info")
		}

		return transferEvent.ProposalId.Int64(), nil, proposalInfoWithConfig, nil
	default:
		return 0, nil, nil, errors.New(fmt.Sprintf("unknown event: %v", eventName))
	}
}

func (ch *checker) getSender(txHash common.Hash) (string, error) {
	tx, _, err := ch.client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return "", errors.Wrap(err, "failed to get transaction by hash")
	}

	sender, err := types.Sender(types.NewCancunSigner(tx.ChainId()), tx)
	if err != nil {
		return "", errors.Wrap(err, "failed to get sender from transaction")
	}

	return sender.Hex(), nil
}

func (ch *checker) insertProcessedEventLog(processedEvent data.ProcessedEvent) error {
	isExist, err := ch.processedEventQ.CheckProcessedEventExist(processedEvent)
	if isExist {
		return errors.New("Duplicate event in db")
	}
	if err != nil {
		return errors.Wrap(err, "failed to check processed event exist in db")
	}

	err = ch.processedEventQ.InsertProcessedEvent(processedEvent)
	if err != nil {
		return errors.Wrap(err, "failed to insert processed event")
	}
	return nil
}

func (ch *checker) checkVotingAndGetBalance(votingId int64, sender string) (*data.VotingInfo, error) {
	votingInfo, err := ch.votingQ.GetVotingInfo(votingId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get voting info")
	}
	if votingInfo != nil {
		return votingInfo, nil
	}

	votingInfo = &data.VotingInfo{
		VotingId:       votingId,
		Balance:        big.NewInt(0),
		GasLimit:       ch.VotingV2Config.GasLimit,
		CreatorAddress: sender,
	}

	err = ch.votingQ.InsertVotingInfo(votingInfo)
	if err != nil {
		return votingInfo, errors.Wrap(err, "failed insert new voting info")
	}

	return votingInfo, nil
}

func (ch *checker) getStartBlockNumber() (uint64, error) {
	block, err := ch.processedEventQ.GetLastBlock()
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, errors.Wrap(err, "failed get block from db")
		}
		block = ch.VotingV2Config.Block
	}

	return block, nil
}

func (ch *checker) getEventHashes() (eventHashes []common.Hash, err error) {
	parsedABI, err := abi.JSON(strings.NewReader(contracts.ProposalsStateABI))
	if err != nil {
		return eventHashes, errors.Wrap(err, "failed to parse contract ABI")
	}

	for _, eventName := range eventNames {
		event, exists := parsedABI.Events[eventName]
		if !exists {
			ch.log.WithField("eventName", eventName).Warn("Event not found in ABI")
			continue
		}
		eventHashes = append(eventHashes, event.ID)
	}
	if len(eventHashes) == 0 {
		return eventHashes, errors.New("No valid event hashes found.")
	}

	return eventHashes, nil
}
