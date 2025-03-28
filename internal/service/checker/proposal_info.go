package checker

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/rarimo/proof-verification-relayer/internal/config"
	"github.com/rarimo/proof-verification-relayer/internal/contracts"
	"github.com/rarimo/proof-verification-relayer/internal/data"
	"github.com/rarimo/proof-verification-relayer/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func GetProposalInfo(proposalId int64, cfg config.Config, creatorAddress string) (*resources.VotingInfoAttributes, error) {
	contractAddress := cfg.VotingV2Config().Address
	client := cfg.VotingV2Config().RPC
	ipfsUrl := cfg.VotingV2Config().IpfsUrl
	caller, err := contracts.NewProposalsStateCaller(contractAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed get new ProposalsStateCaller")
	}

	proposalInfo, err := caller.GetProposalInfo(nil, big.NewInt(proposalId))
	if err != nil {
		return nil, errors.Wrap(err, "failed get proposal info from contract")
	}
	startTimeStamp := proposalInfo.Config.StartTimestamp
	endTimestamp := proposalInfo.Config.StartTimestamp + proposalInfo.Config.Duration

	desc, err := getProposalDescFromIpfs(proposalInfo.Config.Description, ipfsUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed get proposal info from ipfs")
	}
	votingConfig := proposalInfo.Config
	var votingWhitelist []string

	for _, addr := range votingConfig.VotingWhitelist {
		votingWhitelist = append(votingWhitelist, addr.Hex())
	}

	var parsedVotingWhitelistData []resources.ParsedVotingWhiteData
	var votingWhitelistData []string
	for _, data := range votingConfig.VotingWhitelistData {
		votingWhitelistData = append(votingWhitelistData, hex.EncodeToString(data))

		whiteData, err := decodeVotingConfig(data)
		if err != nil {
			cfg.Log().WithError(err).
				WithField("proposal_id", proposalId).
				Error("failed decoding voting white list data")
			continue
		}

		var citizenshipWhitelist []string
		for _, citizenship := range whiteData.CitizenshipWhitelist {
			bytes, err := hex.DecodeString(citizenship.Text(16))
			if err != nil {
				cfg.Log().WithError(err).
					WithField("proposal_id", proposalId).
					Error("failed decoding voting white list data")
				continue
			}
			asciiString := string(bytes)
			citizenshipWhitelist = append(citizenshipWhitelist, asciiString)
		}

		minAge, err := getMinAge(whiteData.BirthDateUpperbound.Text(16), whiteData.ExpirationDateLowerBound.Text(16))
		if err != nil {
			cfg.Log().WithError(err).
				WithField("proposal_id", proposalId).
				Error("failed decoding voting white list data")
		}

		parsedVotingWhitelistData = append(parsedVotingWhitelistData, resources.ParsedVotingWhiteData{
			CitizenshipWhitelist:                citizenshipWhitelist,
			IdentityCreationTimestampUpperBound: whiteData.IdentityCreationTimestampUpperBound.String(),
			IdentityCounterUpperBound:           whiteData.IdentityCounterUpperBound.String(),
			BirthDateUpperBound:                 whiteData.BirthDateUpperbound.Text(16),
			ExpirationDateLowerBound:            whiteData.ExpirationDateLowerBound.Text(16),
			MinAge:                              minAge,
		})
	}
	votingResults := make([][]int64, len(proposalInfo.Config.AcceptedOptions))

	for indx, optionResult := range proposalInfo.VotingResults {
		votingResults[indx] = make([]int64, 8)
		for varIndx := 0; varIndx < 8; varIndx += 1 {
			votingResults[indx][varIndx] = optionResult[varIndx].Int64()
		}
	}

	return &resources.VotingInfoAttributes{
		Author:   creatorAddress,
		Metadata: *desc,
		Contract: resources.VotingInfoContractInfo{
			ProposalSMT:   proposalInfo.ProposalSMT.Hex(),
			Status:        proposalInfo.Status,
			VotingResults: votingResults,
			Config: resources.VotingInfoConfig{
				Description:               votingConfig.Description,
				StartTimestamp:            int64(startTimeStamp),
				EndTimestamp:              int64(endTimestamp),
				Multichoice:               votingConfig.Multichoice.Int64(),
				ProposalId:                proposalId,
				VotingWhitelist:           votingWhitelist,
				VotingWhitelistData:       votingWhitelistData,
				ParsedVotingWhitelistData: parsedVotingWhitelistData,
			},
		},
	}, nil
}

func decodeVotingConfig(dataBytes []byte) (data.VotingWhitelistDataBigInt, error) {
	var whiteData data.VotingWhitelistDataBigInt

	rawABI := `[
		{
		  "name": "ProposalRules",
		  "type": "event",
		  "inputs": [	
		  {"name": "whiteData", "type":"tuple", "components": [
		  {"name": "citizenshipWhitelist", "type": "uint256[]"},
			{"name": "identityCreationTimestampUpperBound", "type": "uint256"},
			{"name": "identityCounterUpperBound", "type": "uint256"},
			{"name": "birthDateUpperbound", "type": "uint256"},
			{"name": "expirationDateLowerBound", "type": "uint256"}
		  ]}  
		  ]
		}
	  ]`
	parsedABI, err := abi.JSON(strings.NewReader(rawABI))
	if err != nil {
		return whiteData, errors.Wrap(err, "failed to parse ABI")
	}

	method, ok := parsedABI.Events["ProposalRules"]
	if !ok {
		return whiteData, fmt.Errorf("method ProposalRules not found in ABI")
	}

	decoded, err := method.Inputs.Unpack(dataBytes)
	if err != nil {
		return whiteData, fmt.Errorf("failed to unpack calldata: %v", err)
	}

	decodedWhiteData := decoded[0]
	dc := decodedWhiteData.(struct {
		CitizenshipWhitelist                []*big.Int `json:"citizenshipWhitelist"`
		IdentityCreationTimestampUpperBound *big.Int   `json:"identityCreationTimestampUpperBound"`
		IdentityCounterUpperBound           *big.Int   `json:"identityCounterUpperBound"`
		BirthDateUpperbound                 *big.Int   `json:"birthDateUpperbound"`
		ExpirationDateLowerBound            *big.Int   `json:"expirationDateLowerBound"`
	})

	whiteData = data.VotingWhitelistDataBigInt{
		CitizenshipWhitelist:                dc.CitizenshipWhitelist,
		IdentityCreationTimestampUpperBound: dc.IdentityCreationTimestampUpperBound,
		IdentityCounterUpperBound:           dc.IdentityCounterUpperBound,
		BirthDateUpperbound:                 dc.BirthDateUpperbound,
		ExpirationDateLowerBound:            dc.ExpirationDateLowerBound,
	}

	return whiteData, nil
}

func getMinAge(birthDateUpperbound string, expirationDateLowerBound string) (int64, error) {
	if birthDateUpperbound == "303030303030" {
		return 0, nil
	}
	bytes, err := hex.DecodeString(birthDateUpperbound)
	if err != nil {
		return 0, errors.Wrap(err, "failed decode birthDateUpperbound")
	}
	birthDateUpperboundTime, err := time.Parse("060102", string(bytes))
	if err != nil {
		return 0, errors.Wrap(err, "failed parse birthDateUpperbound")
	}

	bytes, err = hex.DecodeString(expirationDateLowerBound)
	if err != nil {
		return 0, errors.Wrap(err, "failed decode expirationDateLowerBound")
	}

	expirationDateLowerBoundTime, err := time.Parse("060102", string(bytes))
	if err != nil {
		return 0, errors.Wrap(err, "failed parse expirationDateLowerBound")
	}

	years := expirationDateLowerBoundTime.Year() - birthDateUpperboundTime.Year()
	birthDateWithSubYears := birthDateUpperboundTime.AddDate(years, 0, 0)

	if expirationDateLowerBoundTime.Before(birthDateWithSubYears) {
		years--
	}

	return int64(years), nil
}

func getProposalDescFromIpfs(desId string, ipfsUrl string) (*resources.VotingInfoAttributesMetadata, error) {
	requestURL := fmt.Sprintf("%s/%s", ipfsUrl, desId)
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read body response: %v", err)
	}

	var data resources.VotingInfoAttributesMetadata
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed parse JSON: %v", err)
	}

	return &data, nil
}
