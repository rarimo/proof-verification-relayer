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

var DuplicateEventErr = errors.New("Duplicate event in db")

// processLog based on the event name and filtered log, it parses and update the relevant information
func (ch *checker) processLog(vLog types.Log, eventName string) error {
	var processedEvent data.ProcessedEvent

	block := int64(vLog.BlockNumber)
	processedEvent.BlockNumber = block
	processedEvent.LogIndex = int64(vLog.Index)
	processedEvent.TransactionHash = vLog.TxHash[:]
	err := ch.insertProcessedEventLog(processedEvent)
	if err != nil {
		if err == DuplicateEventErr {
			ch.log.WithFields(logan.F{
				"hash_tx":   vLog.TxHash.Hex(),
				"log_index": processedEvent.LogIndex,
			}).Debug("event already processed and stored")
			return nil
		}
		return errors.Wrap(err, "failed insert processed event")
	}
	sender, err := ch.getSender(vLog.TxHash)
	if err != nil {
		return errors.Wrap(err, "failed get processed event creator")
	}
	votingInfo, err := ch.getEventData(eventName, vLog, sender)
	if err != nil {
		return errors.Wrap(err, "failed get event data")
	}
	if err = ch.votingQ.UpdateVotingInfo(votingInfo); err != nil {
		return errors.Wrap(err, "failed update voting balance", logan.F{"Voting ID": votingInfo.VotingId})
	}
	return nil
}

// getEventData According to the event name, the log is unpacked using the ABI contract.
// Based on the decompressed data, voting information is created or updated
//
// event ProposalConfigChanged(uint256 indexed proposalId);
// event ProposalFunded(uint256 indexed proposalId, uint256 fundAmount);
// event ProposalCreated(uint256 indexed proposalId, address proposalSMT, uint256 fundAmount)
func (ch *checker) getEventData(eventName string, vLog types.Log, sender string) (votingInfo *data.VotingInfo, err error) {
	if len(vLog.Topics) < 2 {
		return nil, errors.New("the topic must contain at least two elements")
	}

	// Since the parameter is indexed, it is found in the topics and not in vLog.Data,
	// so it was not passed to the corresponding transferEvent parameter when unpacking
	votingId := new(big.Int).SetBytes(vLog.Topics[1].Bytes()).Int64()

	votingInfo, err = ch.checkVotingAndGetInfo(votingId, sender)
	if err != nil {
		return nil, errors.Wrap(err, "failed get voting info", logan.F{"Voting ID": votingId})
	}

	switch eventName {
	case "ProposalCreated":
		var transferEvent contracts.ProposalsStateProposalCreated
		err = ch.parsedABI.UnpackIntoInterface(&transferEvent, eventName, vLog.Data)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unpack log data")
		}

		votingInfo.Balance = new(big.Int).Add(votingInfo.Balance, transferEvent.FundAmount)
		votingInfo.TotalBalance = new(big.Int).Add(votingInfo.TotalBalance, transferEvent.FundAmount)

		proposalInfoFromContract, err := GetProposalInfo(votingId, ch.cfg, sender)
		if err != nil {
			return nil, errors.Wrap(err, "failed get proposal info")
		}

		votingInfo = unpackInfoFromContractToStruct(proposalInfoFromContract, votingInfo)

		return votingInfo, nil
	case "ProposalFunded":
		var transferEvent contracts.ProposalsStateProposalFunded
		err = ch.parsedABI.UnpackIntoInterface(&transferEvent, eventName, vLog.Data)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unpack log data")
		}

		votingInfo.Balance = new(big.Int).Add(votingInfo.Balance, transferEvent.FundAmount)
		votingInfo.TotalBalance = new(big.Int).Add(votingInfo.TotalBalance, transferEvent.FundAmount)

		return votingInfo, nil
	case "ProposalConfigChanged":
		var transferEvent contracts.ProposalsStateProposalConfigChanged
		err = ch.parsedABI.UnpackIntoInterface(&transferEvent, eventName, vLog.Data)
		if err != nil {
			return nil, errors.Wrap(err, "failed to unpack log data")
		}

		proposalInfoFromContract, err := GetProposalInfo(votingId, ch.cfg, sender)
		if err != nil {
			return nil, errors.Wrap(err, "failed get proposal info")
		}

		votingInfo = unpackInfoFromContractToStruct(proposalInfoFromContract, votingInfo)

		return votingInfo, nil
	default:
		return nil, fmt.Errorf("unknown event: %v", eventName)
	}
}

// unpackInfoFromContractToStruct simply unpack the ProposalInfoFromContract into the required VotingInfo fields
func unpackInfoFromContractToStruct(proposalInfoFromContract *ProposalInfoFromContract, votingInfo *data.VotingInfo) *data.VotingInfo {
	if len(proposalInfoFromContract.ParsedVotingWhitelistData) > 0 {
		votingInfo.MaxAge = proposalInfoFromContract.ParsedVotingWhitelistData[0].MaxAge
		votingInfo.MinAge = proposalInfoFromContract.ParsedVotingWhitelistData[0].MinAge
	}

	votingInfo.StartTimestamp = proposalInfoFromContract.StartTimestamp
	votingInfo.EndTimestamp = proposalInfoFromContract.EndTimestamp
	votingInfo.ParsedWhiteListDataWithMetadata = resources.ParsedWhiteListDataWithMetadata{
		Metadata:                  proposalInfoFromContract.Metadata,
		ParsedVotingWhitelistData: proposalInfoFromContract.ParsedVotingWhitelistData,
	}

	return votingInfo
}

// getSender finds and returns the address of the transaction sender
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

// insertProcessedEventLog checks for the presence of an event in the database and adds information about a new event if it is not there
func (ch *checker) insertProcessedEventLog(processedEvent data.ProcessedEvent) error {
	isExist, err := ch.processedEventQ.CheckProcessedEventExist(processedEvent)
	if isExist {
		return DuplicateEventErr
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

// checkVotingAndGetInfo checks if there is a voting with this ID in the database.
// If there is, it returns information about it in the form of data.VotingInfo,
// otherwise it creates it with the default value, adds it to the database,
// and returns the created object
func (ch *checker) checkVotingAndGetInfo(votingId int64, sender string) (*data.VotingInfo, error) {
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
		TotalBalance:   big.NewInt(0),
		GasLimit:       ch.VotingV2Config.GasLimit,
		CreatorAddress: sender,
	}

	err = ch.votingQ.InsertVotingInfo(votingInfo)
	if err != nil {
		return votingInfo, errors.Wrap(err, "failed insert new voting info")
	}

	return votingInfo, nil
}

// getStartBlockNumber returns the last processed block number from the database.
// If no records are found, uses the default value from config.
// This block number is used as a starting point for scanning events.
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

// getEventHashes search and extract event hashes from ABI to filter logs
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
