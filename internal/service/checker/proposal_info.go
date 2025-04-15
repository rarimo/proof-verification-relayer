package checker

import (
	"context"
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
	"gitlab.com/distributed_lab/running"
)

type ProposalInfoFromContract struct {
	// Timespamp when the voting ends
	EndTimestamp int64
	// Timespamp when voting starts
	StartTimestamp int64
	// Proposal data received from IPFS according to the CID received from the contract
	Metadata resources.ProposalInfoAttributesMetadata
	// Rules for the proposal
	ParsedVotingWhitelistData []resources.ParsedVotingWhiteData
}

// GetProposalInfo gets information from the contract for the corresponding proposal_id.
// Parses the received information and generates it in the required format.
func GetProposalInfo(proposalId int64, cfg config.Config, creatorAddress string) (*ProposalInfoFromContract, error) {
	contractAddress := cfg.VotingV2Config().Address
	client := cfg.VotingV2Config().RPC
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

	// Description in proposalInfo.Config contains the CID for the resource that corresponds to this proposal in the IPFS repository
	proposalMetadata := &resources.ProposalInfoAttributesMetadata{}
	var lastErr error
	running.WithThreshold(context.Background(), cfg.Log(), "getProposalDescFromIpfs", func(ctx context.Context) (bool, error) {
		desc, err := getProposalDescFromIpfs(proposalInfo.Config.Description, cfg.Ipfs().Url)
		if err != nil {
			lastErr = err
			return false, errors.Wrap(err, "failed get proposal info from ipfs")
		}

		proposalMetadata = desc
		return true, nil
	}, cfg.Ipfs().MinAbnormalPeriod, cfg.Ipfs().MaxAbnormalPeriod, cfg.Ipfs().MaxRetries)

	if lastErr != nil {
		return nil, errors.Wrap(lastErr, "failed get proposal info from ipfs")
	}
	votingConfig := proposalInfo.Config

	// proposalInfo.Config.VotingWhitelistData is a list of rules for this proposal according to the VotingWhitelist,
	// which is a list of Voting contracts that are allowed for this proposal.
	//
	// For more information, see ProposalsState.sol:
	// https://github.com/rarimo/passport-voting-contracts
	var parsedVotingWhitelistData []resources.ParsedVotingWhiteData
	for _, data := range votingConfig.VotingWhitelistData {

		whiteData, err := decodeWhitelistData(data)
		if err != nil {
			cfg.Log().WithError(err).
				WithField("proposal_id", proposalId).
				Error("failed decoding voting white list data")
			continue
		}

		// Convert citizenship whitelist to understandable ASCII format
		var citizenshipWhitelist []string
		for _, citizenship := range whiteData.CitizenshipWhitelist {
			bytes, err := hex.DecodeString(citizenship.Text(16))
			if err != nil {
				cfg.Log().WithError(err).
					WithField("proposal_id", proposalId).
					Error("failed decoding voting citizenship list")
				continue
			}
			asciiString := string(bytes)
			citizenshipWhitelist = append(citizenshipWhitelist, asciiString)
		}

		minAge, err := getBoundaryInYears(whiteData.BirthDateUpperbound.Text(16), startTimeStamp)
		if err != nil {
			cfg.Log().WithError(err).
				WithField("proposal_id", proposalId).
				Error("failed decoding voting min age")
			continue
		}

		maxAge, err := getBoundaryInYears(whiteData.BirthDateLowerbound.Text(16), startTimeStamp)
		if err != nil {
			cfg.Log().WithError(err).
				WithField("proposal_id", proposalId).
				Error("failed decoding voting max age")
			continue
		}

		parsedVotingWhitelistData = append(parsedVotingWhitelistData, resources.ParsedVotingWhiteData{
			Selector:                            whiteData.Selector.Text(16),
			BirthDateLowerbound:                 whiteData.BirthDateLowerbound.Text(16),
			CitizenshipWhitelist:                citizenshipWhitelist,
			IdentityCreationTimestampUpperBound: whiteData.IdentityCreationTimestampUpperBound.String(),
			IdentityCounterUpperBound:           whiteData.IdentityCounterUpperBound.String(),
			BirthDateUpperBound:                 whiteData.BirthDateUpperbound.Text(16),
			ExpirationDateLowerBound:            whiteData.ExpirationDateLowerBound.Text(16),
			MinAge:                              minAge,
			MaxAge:                              maxAge,
			Sex:                                 whiteData.Sex.Text(16),
		})
	}

	return &ProposalInfoFromContract{
		StartTimestamp:            int64(startTimeStamp),
		EndTimestamp:              int64(endTimestamp),
		Metadata:                  *proposalMetadata,
		ParsedVotingWhitelistData: parsedVotingWhitelistData,
	}, nil
}

// decodeWhitelistData decodes WhitelistData to generate proposal rules according to the ProposalRules structure in the BaseVoting contract.
// Read more about contracts:
// https://github.com/rarimo/passport-voting-contracts
func decodeWhitelistData(dataBytes []byte) (data.VotingWhitelistDataBigInt, error) {
	var whiteData data.VotingWhitelistDataBigInt

	rawABI := `[
		{
		  "name": "ProposalRules",
		  "type": "event",
		  "inputs": [	
		  {"name": "whiteData", "type":"tuple", "components": [
			{"name": "selector", "type": "uint256"},
		    {"name": "citizenshipWhitelist", "type": "uint256[]"},
			{"name": "identityCreationTimestampUpperBound", "type": "uint256"},
			{"name": "identityCounterUpperBound", "type": "uint256"},
			{"name": "sex", "type": "uint256"},
			{"name": "birthDateLowerbound", "type": "uint256"},
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

	// gets the first element because it is a tuple
	decodedWhiteData := decoded[0]
	dc := decodedWhiteData.(struct {
		Selector                            *big.Int   `json:"selector"`
		CitizenshipWhitelist                []*big.Int `json:"citizenshipWhitelist"`
		IdentityCreationTimestampUpperBound *big.Int   `json:"identityCreationTimestampUpperBound"`
		IdentityCounterUpperBound           *big.Int   `json:"identityCounterUpperBound"`
		Sex                                 *big.Int   `json:"sex"`
		BirthDateLowerbound                 *big.Int   `json:"birthDateLowerbound"`
		BirthDateUpperbound                 *big.Int   `json:"birthDateUpperbound"`
		ExpirationDateLowerBound            *big.Int   `json:"expirationDateLowerBound"`
	})

	whiteData = data.VotingWhitelistDataBigInt{
		Selector:                            dc.Selector,
		CitizenshipWhitelist:                dc.CitizenshipWhitelist,
		IdentityCreationTimestampUpperBound: dc.IdentityCreationTimestampUpperBound,
		IdentityCounterUpperBound:           dc.IdentityCounterUpperBound,
		Sex:                                 dc.Sex,
		BirthDateLowerbound:                 dc.BirthDateLowerbound,
		BirthDateUpperbound:                 dc.BirthDateUpperbound,
		ExpirationDateLowerBound:            dc.ExpirationDateLowerBound,
	}

	return whiteData, nil
}

// getBoundaryInYears counts the exact number of years from boundaryDate to startTimeStamp
func getBoundaryInYears(boundaryDateHex string, startTimeStamp uint64) (int64, error) {
	// ZERO_DATE 00000 in hex format
	if boundaryDateHex == "303030303030" {
		return 0, nil
	}
	bytes, err := hex.DecodeString(boundaryDateHex)
	if err != nil {
		return 0, errors.Wrap(err, "failed to decode boundary date")
	}
	boundaryDate, err := time.Parse("060102", string(bytes))
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse boundary date")
	}

	startDate := time.Unix(int64(startTimeStamp), 0)

	// If the year of birth is mistakenly interpreted as future (due to the YYMMDD format),
	// add 100 years to the boundary to keep the correct difference between the dates.
	if boundaryDate.Year() > startDate.Year() {
		startDate = startDate.AddDate(100, 0, 0)
	}

	years := startDate.Year() - boundaryDate.Year()

	// Check if the birthday has already occurred in the year startDate.
	// If not, subtract 1 year to accurately calculate the full age.
	// For example:
	// birthDate = 02.01.2000, expirationDate = 01.01.2023 → age = 22 (not 23)
	birthDateWithSubYears := boundaryDate.AddDate(years, 0, 0)
	if startDate.Before(birthDateWithSubYears) {
		years--
	}

	return int64(years), nil
}

// getProposalDescFromIpfs requests and returns proposal information from ipfs by CID
func getProposalDescFromIpfs(desId string, ipfsUrl string) (*resources.ProposalInfoAttributesMetadata, error) {
	var data resources.ProposalInfoAttributesMetadata
	requestURL := fmt.Sprintf("%s/%s", ipfsUrl, desId)

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response from ipfs: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err)
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &data, nil
}
