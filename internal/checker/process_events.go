package checker

import (
	"database/sql"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (ch *checker) processEvents(event contracts.ProposalEvent) error {
	var processedEvent data.ProcessedEvent
	processedEvent.BlockNumber = event.BlockNumber()
	processedEvent.LogIndex = event.LogIndex()
	processedEvent.TransactionHash = event.TxHash().Bytes()
	if err := ch.insertProcessedEventLog(processedEvent); err != nil {
		ch.log.WithField("Error", err).Debug("failed insert processed event")
		return err
	}
	votingInfo, err := ch.checkVoteAndGetBalance(event.ProposalID())
	if err != nil {
		ch.log.WithField("Error", err).Errorf("failed insert processed event")

		return err
	}
	votingInfo.Balance += event.FundAmountU64()

	if err := ch.checkerQ.UpdateVotingInfo(votingInfo); err != nil {
		ch.log.Errorf("Failed Update voting balance: %v", err)
		return err
	}

	return nil
}

func (ch *checker) processLog(vLog types.Log, eventName string) error {
	var processedEvent data.ProcessedEvent

	block := int64(vLog.BlockNumber)
	processedEvent.BlockNumber = block
	processedEvent.LogIndex = int64(vLog.Index)
	processedEvent.TransactionHash = vLog.TxHash[:]

	err := ch.insertProcessedEventLog(processedEvent)
	if err != nil {
		ch.log.WithFields(logan.F{"Error": err, "eventName": eventName}).Error("failed insert processed event")
		return err
	}

	votingId, value, err := ch.getTransferEvent(eventName, vLog)
	if err != nil {
		return err
	}

	votingInfo, err := ch.checkVoteAndGetBalance(votingId)
	if err != nil {
		ch.log.WithFields(logan.F{"Error": err, "eventName": eventName, "Voting ID": votingId}).Error("failed get voting info")
		return err
	}
	votingInfo.Balance = votingInfo.Balance + value

	err = ch.checkerQ.UpdateVotingInfo(votingInfo)
	if err != nil {
		ch.log.WithFields(logan.F{"Error": err, "eventName": eventName, "Voting ID": votingId}).Error("Failed Update voting balance")
		return err
	}
	return nil
}

func (ch *checker) getTransferEvent(eventName string, vLog types.Log) (votingId int64, value uint64, err error) {
	parsedABI, err := abi.JSON(strings.NewReader(contracts.ProposalsStateABI))
	if err != nil {
		ch.log.Errorf("Failed to parse contract ABI: %v", err)
		return 0, 0, err
	}

	if eventName == "ProposalCreated" {
		var transferEvent contracts.ProposalsStateProposalCreated
		err = parsedABI.UnpackIntoInterface(&transferEvent, eventName, vLog.Data)
		if err != nil {
			ch.log.Errorf("Failed to unpack log data: %v", err)
			return 0, 0, err
		}
		return transferEvent.ProposalId.Int64(), transferEvent.FundAmount.Uint64(), nil
	}

	var transferEvent contracts.ProposalsStateProposalFunded
	err = parsedABI.UnpackIntoInterface(&transferEvent, eventName, vLog.Data)
	if err != nil {
		ch.log.Errorf("Failed to unpack log data: %v", err)
		return 0, 0, err
	}
	return transferEvent.ProposalId.Int64(), transferEvent.FundAmount.Uint64(), nil
}

func (ch *checker) insertProcessedEventLog(processedEvent data.ProcessedEvent) error {
	isExist, err := ch.checkerQ.CheckProcessedEventExist(processedEvent)
	if isExist {
		ch.log.WithField("event_index_log", processedEvent.LogIndex).Debug("Duplicate event in db")
		return errors.New("Duplicate event in db")
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	err = ch.checkerQ.InsertProcessedEvent(processedEvent)
	if err != nil {
		return err
	}
	return nil
}

func (ch *checker) checkVoteAndGetBalance(votingId int64) (data.VotingInfo, error) {
	votingInfo, err := ch.checkerQ.GetVotingInfo(votingId)
	if err == sql.ErrNoRows {
		newVote := &data.VotingInfo{
			VotingId: votingId,
			Balance:  0,
			GasLimit: ch.VotingV2Config.GasLimit,
		}

		err := ch.checkerQ.InsertVotingInfo(*newVote)
		if err != nil {
			ch.log.Errorf("Failed insert new voting info: %v", err)
			return *newVote, err
		}
		return *newVote, nil
	} else if err != nil {
		ch.log.Errorf("Error get voting: %v", err)
		return data.VotingInfo{}, err
	}
	return votingInfo, err
}

func (ch *checker) getStartBlockNumber() (uint64, error) {
	block, err := ch.checkerQ.GetLastBlock()
	if err == sql.ErrNoRows {
		block = ch.VotingV2Config.Block
	} else if err != nil {
		ch.log.Errorf("Failed get block from db: %v", err)
		return ch.VotingV2Config.Block, err
	}

	return block, nil
}
