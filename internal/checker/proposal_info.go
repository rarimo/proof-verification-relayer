package checker

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/internal/data/pg"
	"github.com/rarimo/proof-verification-relayer/internal/service/api/requests"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/logan/v3"
)

func decodeVotingConfig(encodedHex string) (data.VotingWhitelistDataBigInt, error) {
	var config data.VotingWhitelistDataBigInt
	dataBytes, err := hexutil.Decode(encodedHex)
	if err != nil {
		return config, err
	}

	rawABI := `
    [
      {
        "type":"function",
        "name":"ProposalRules",
        "inputs":[
          {
            "name":"config",
            "type":"tuple",
            "components":[
              {"name":"citizenshipWhitelist","type":"uint256[]"},
              {"name":"identityCreationTimestampUpperBound","type":"uint256"},
              {"name":"identityCounterUpperBound","type":"uint256"},
              {"name":"birthDateUpperbound","type":"uint256"},
              {"name":"expirationDateLowerBound","type":"uint256"}
            ]
          }
        ],
        "outputs":[]
      }
    ]`
	parsedABI, err := abi.JSON(strings.NewReader(rawABI))
	if err != nil {
		return config, err
	}

	method, ok := parsedABI.Methods["ProposalRules"]
	if !ok {
		return config, fmt.Errorf("method ProposalRules not found in ABI")
	}

	decoded, err := method.Inputs.Unpack(dataBytes)
	if err != nil {
		return config, err
	}

	if len(decoded) != 1 {
		return config, fmt.Errorf("unexpected argument count, want 1 got %d", len(decoded))
	}

	val := reflect.ValueOf(decoded[0])

	arrField := val.FieldByName("CitizenshipWhitelist").Interface()
	citizenshipWhitelist, ok := arrField.([]*big.Int)
	if !ok {
		return config, fmt.Errorf("failed to cast citizenshipWhitelist")
	}

	config = data.VotingWhitelistDataBigInt{
		CitizenshipWhitelist:                citizenshipWhitelist,
		IdentityCreationTimestampUpperBound: val.FieldByName("IdentityCreationTimestampUpperBound").Interface().(*big.Int),
		IdentityCounterUpperBound:           val.FieldByName("IdentityCounterUpperBound").Interface().(*big.Int),
		BirthDateUpperbound:                 val.FieldByName("BirthDateUpperbound").Interface().(*big.Int),
		ExpirationDateLowerBound:            val.FieldByName("ExpirationDateLowerBound").Interface().(*big.Int),
	}

	return config, nil
}

func GetMinAge(birthDateUpperbound string, expirationDateLowerBound string, log *logan.Entry) (int64, error) {
	if birthDateUpperbound == "303030303030" {
		return 0, nil
	}
	bytes, err := hex.DecodeString(birthDateUpperbound)
	if err != nil {
		return 0, err
	}
	birthDateUpperboundTime, err := time.Parse("060102", string(bytes))
	if err != nil {
		return 0, err
	}

	bytes, err = hex.DecodeString(expirationDateLowerBound)
	if err != nil {
		return 0, err
	}

	expirationDateLowerBoundTime, err := time.Parse("060102", string(bytes))
	if err != nil {
		return 0, err
	}

	delta := expirationDateLowerBoundTime.Sub(birthDateUpperboundTime)
	years := int64(delta.Hours() / (24 * 365.25))

	return years, nil
}

func GetProposalList(cfg config.Config, req requests.ProposalInfoFilter) ([]*resources.VotingInfoAttributes, error) {
	var result []*resources.VotingInfoAttributes
	pgDB := pg.NewMaterDB(cfg.DB()).CheckerQ()

	votingInfoList, err := pgDB.SelectVotes(req)
	if err != nil {
		return nil, err
	}
	for _, vote := range votingInfoList {
		result = append(result, &vote.ProposalInfoWithConfig)
	}

	return result, nil
}

func GetProposalInfo(proposalId int64, cfg config.Config, creatorAddress string) (*resources.VotingInfoAttributes, error) {
	contractAddress := cfg.VotingV2Config().Address
	client := cfg.VotingV2Config().RPC
	ipfsUrl := cfg.VotingV2Config().IpfsUrl
	caller, err := contracts.NewProposalsStateCaller(contractAddress, client)
	if err != nil {
		return nil, err
	}

	proposalInfo, err := caller.GetProposalInfo(nil, big.NewInt(proposalId))
	if err != nil {
		return nil, err
	}
	startTimeStamp := proposalInfo.Config.StartTimestamp
	endTimestamp := proposalInfo.Config.StartTimestamp + proposalInfo.Config.Duration

	desc, err := getProposalDescFromIpfs(proposalInfo.Config.Description, ipfsUrl)
	if err != nil {
		return nil, err
	}
	votingConfig := proposalInfo.Config
	var votingWhitelist []string

	for _, addr := range votingConfig.VotingWhitelist {
		votingWhitelist = append(votingWhitelist, addr.Hex())
	}
	for _, addr := range votingConfig.VotingWhitelist {
		votingWhitelist = append(votingWhitelist, addr.Hex())
	}
	var votingWhitelistData []string
	for _, data := range votingConfig.VotingWhitelistData {
		votingWhitelistData = append(votingWhitelistData, hex.EncodeToString(data))
	}
	votingWhitelistDataHex := fmt.Sprintf("0x%v", votingWhitelistData[0])

	parsedVotingWhitelistData, _ := decodeVotingConfig(votingWhitelistDataHex)
	var citizenshipWhitelist []string
	for _, citizenship := range parsedVotingWhitelistData.CitizenshipWhitelist {
		bytes, err := hex.DecodeString(citizenship.Text(16))
		if err != nil {
			cfg.Log().WithFields(logan.F{
				"error":       err,
				"proposal_id": proposalId,
			}).Debug("failed decoding citizenship")
			continue
		}
		asciiString := string(bytes)
		citizenshipWhitelist = append(citizenshipWhitelist, asciiString)
	}

	minAge, err := GetMinAge(parsedVotingWhitelistData.BirthDateUpperbound.Text(16), parsedVotingWhitelistData.ExpirationDateLowerBound.Text(16), cfg.Log())
	if err != nil {
		cfg.Log().WithFields(logan.F{
			"error":       err,
			"proposal_id": proposalId,
		}).Debug("failed get min age")
	}
	return &resources.VotingInfoAttributes{
		Author:   creatorAddress,
		Metadata: *desc,
		Contract: resources.VotingInfoContractInfo{
			ProposalSMT: proposalInfo.ProposalSMT.Hex(),
			Status:      proposalInfo.Status,
			Config: resources.VotingInfoConfig{
				Description:         votingConfig.Description,
				StartTimestamp:      int64(startTimeStamp),
				EndTimestamp:        int64(endTimestamp),
				Multichoice:         votingConfig.Multichoice.Int64(),
				ProposalId:          proposalId,
				VotingWhitelist:     votingWhitelist,
				VotingWhitelistData: votingWhitelistData,
				ParsedVotingWhitelistData: resources.ParsedVotingWhitelistData{
					CitizenshipWhitelist:                citizenshipWhitelist,
					IdentityCreationTimestampUpperBound: parsedVotingWhitelistData.IdentityCreationTimestampUpperBound.String(),
					IdentityCounterUpperBound:           parsedVotingWhitelistData.IdentityCounterUpperBound.String(),
					BirthDateUpperBound:                 parsedVotingWhitelistData.BirthDateUpperbound.Text(16),
					ExpirationDateLowerBound:            parsedVotingWhitelistData.ExpirationDateLowerBound.Text(16),
					MinAge:                              minAge,
				},
			},
		},
	}, nil
}
