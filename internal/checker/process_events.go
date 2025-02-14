package checker

import (
	"database/sql"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/data"
)

func (ch *checker) processEvent(event *contracts.ProposalsStateProposalCreated) error {
	var processedEvent data.ProcessedEvent
	block := int64(event.Raw.BlockNumber)
	processedEvent.BlockNumber = block
	processedEvent.LogIndex = int64(event.Raw.Index)
	processedEvent.TransactionHash = event.Raw.TxHash[:]

	err := ch.insertProcessedEventLog(processedEvent)
	if err != nil {
		return err
	}
	votingId := event.ProposalId.Int64()

	value := event.FundAmount
	votingInfo, err := ch.checkVoteAndGetBalance(votingId)
	if err != nil {
		return err
	}

	votingInfo.Balance = votingInfo.Balance + value.Uint64()

	err = ch.checkerQ.UpdateVotingInfo(votingInfo)
	if err != nil {
		ch.log.Errorf("Failed Update voting balance: %v", err)
		return err
	}

	return nil
}

func (ch *checker) processLog(vLog types.Log) error {
	var processedEvent data.ProcessedEvent
	parsedABI, err := abi.JSON(strings.NewReader(contracts.ProposalsStateABI))
	if err != nil {
		ch.log.Errorf("Failed to parse contract ABI: %v", err)
		return err
	}

	var transferEvent contracts.ProposalsStateProposalCreated

	err = parsedABI.UnpackIntoInterface(&transferEvent, "ProposalCreated", vLog.Data)
	if err != nil {
		ch.log.Errorf("Failed to unpack log data: %v", err)
		return err
	}
	block := int64(vLog.BlockNumber)
	processedEvent.BlockNumber = block
	processedEvent.LogIndex = int64(vLog.Index)
	processedEvent.TransactionHash = vLog.TxHash[:]

	err = ch.insertProcessedEventLog(processedEvent)
	if err != nil {
		return err
	}
	votingId := transferEvent.ProposalId.Int64()
	value := transferEvent.FundAmount

	votingInfo, err := ch.checkVoteAndGetBalance(votingId)
	if err != nil {
		return err
	}

	votingInfo.Balance = votingInfo.Balance + value.Uint64()

	err = ch.checkerQ.UpdateVotingInfo(votingInfo)
	if err != nil {
		ch.log.Errorf("Failed Update voting balance: %v", err)
		return err
	}
	return nil
}

func (ch *checker) insertProcessedEventLog(processedEvent data.ProcessedEvent) error {
	isExist, err := ch.checkerQ.CheckProcessedEventExist(processedEvent)
	if isExist {
		ch.log.WithField("event_index_log", processedEvent.LogIndex).Debug("Duplicate event in db")
		return nil
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
