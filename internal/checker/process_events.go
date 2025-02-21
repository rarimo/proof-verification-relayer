package checker

import (
	"database/sql"
	"math/big"
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
		return errors.Wrap(err, "failed insert processed event")
	}
	votingInfo, err := ch.checkVoteAndGetBalance(event.ProposalID())
	if err != nil {
		return errors.Wrap(err, "failed get voting info")
	}
	votingInfo.Balance = new(big.Int).Add(votingInfo.Balance, event.FundAmount())

	if err := ch.checkerQ.UpdateVotingInfo(votingInfo); err != nil {
		return errors.Wrap(err, "failed Update voting balance")
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
		return errors.Wrap(err, "failed insert processed event")
	}

	votingId, value, err := ch.getTransferEvent(eventName, vLog)
	if err != nil {
		return err
	}

	votingInfo, err := ch.checkVoteAndGetBalance(votingId)
	if err != nil {
		return errors.Wrap(err, "failed get voting info", logan.F{"Voting ID": votingId})
	}
	votingInfo.Balance = new(big.Int).Add(votingInfo.Balance, value)

	err = ch.checkerQ.UpdateVotingInfo(votingInfo)
	if err != nil {
		return errors.Wrap(err, "failed update voting balance", logan.F{"Voting ID": votingId})
	}
	return nil
}

func (ch *checker) getTransferEvent(eventName string, vLog types.Log) (votingId int64, value *big.Int, err error) {
	parsedABI, err := abi.JSON(strings.NewReader(contracts.ProposalsStateABI))
	if err != nil {
		ch.log.Errorf("Failed to parse contract ABI: %v", err)
		return 0, nil, err
	}

	if eventName == "ProposalCreated" {
		var transferEvent contracts.ProposalsStateProposalCreated
		err = parsedABI.UnpackIntoInterface(&transferEvent, eventName, vLog.Data)
		if err != nil {
			ch.log.Errorf("Failed to unpack log data: %v", err)
			return 0, nil, err
		}
		return transferEvent.ProposalId.Int64(), transferEvent.FundAmount, nil
	}

	var transferEvent contracts.ProposalsStateProposalFunded
	err = parsedABI.UnpackIntoInterface(&transferEvent, eventName, vLog.Data)
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to unpack log data")
	}
	return transferEvent.ProposalId.Int64(), transferEvent.FundAmount, nil
}

func (ch *checker) insertProcessedEventLog(processedEvent data.ProcessedEvent) error {
	isExist, err := ch.checkerQ.CheckProcessedEventExist(processedEvent)
	if isExist {
		return errors.New("Duplicate event in db")
	}
	if err != nil {
		return err
	}
	err = ch.checkerQ.InsertProcessedEvent(processedEvent)
	if err != nil {
		return err
	}
	return nil
}

func (ch *checker) checkVoteAndGetBalance(votingId int64) (*data.VotingInfo, error) {
	votingInfo, err := ch.checkerQ.GetVotingInfo(votingId)
	if err != nil {
		return nil, err
	}
	if votingInfo != nil {
		return votingInfo, nil
	}

	votingInfo = &data.VotingInfo{
		VotingId: votingId,
		Balance:  big.NewInt(0),
		GasLimit: ch.VotingV2Config.GasLimit,
	}

	err = ch.checkerQ.InsertVotingInfo(votingInfo)
	if err != nil {
		return votingInfo, errors.Wrap(err, "failed insert new voting info")
	}

	return votingInfo, nil
}

func (ch *checker) getStartBlockNumber() (uint64, error) {
	block, err := ch.checkerQ.GetLastBlock()
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, errors.Wrap(err, "failed get block from db")
		}
		block = ch.VotingV2Config.Block
	}
	return block, nil
}
