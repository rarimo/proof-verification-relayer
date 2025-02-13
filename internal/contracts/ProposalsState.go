// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ProposalsStateProposalConfig is an auto generated low-level Go binding around an user-defined struct.
type ProposalsStateProposalConfig struct {
	StartTimestamp      uint64
	Duration            uint64
	Multichoice         *big.Int
	AcceptedOptions     []*big.Int
	Description         string
	VotingWhitelist     []common.Address
	VotingWhitelistData [][]byte
}

// ProposalsStateProposalInfo is an auto generated low-level Go binding around an user-defined struct.
type ProposalsStateProposalInfo struct {
	ProposalSMT   common.Address
	Status        uint8
	Config        ProposalsStateProposalConfig
	VotingResults [][8]*big.Int
}

// ProposalsStateMetaData contains all meta data concerning the ProposalsState contract.
var ProposalsStateMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposalSMT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundAmount\",\"type\":\"uint256\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"hide\",\"type\":\"bool\"}],\"name\":\"ProposalHidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"userNullifier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"vote\",\"type\":\"uint256[]\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAGIC_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_CHOICES_PER_OPTION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_OPTIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proposalSMTImpl_\",\"type\":\"address\"}],\"name\":\"__ProposalsState_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"votingName_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"votingAddress_\",\"type\":\"address\"}],\"name\":\"addVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"multichoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"acceptedOptions\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"votingWhitelist\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"votingWhitelistData\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposalsState.ProposalConfig\",\"name\":\"newProposalConfig_\",\"type\":\"tuple\"}],\"name\":\"changeProposalConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newSignerPubKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"multichoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"acceptedOptions\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"votingWhitelist\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"votingWhitelistData\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposalsState.ProposalConfig\",\"name\":\"proposalConfig_\",\"type\":\"tuple\"}],\"name\":\"createProposal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"}],\"name\":\"getProposalConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"multichoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"acceptedOptions\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"votingWhitelist\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"votingWhitelistData\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposalsState.ProposalConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"}],\"name\":\"getProposalEventId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"}],\"name\":\"getProposalInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"proposalSMT\",\"type\":\"address\"},{\"internalType\":\"enumProposalsState.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"duration\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"multichoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"acceptedOptions\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"votingWhitelist\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"votingWhitelistData\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposalsState.ProposalConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"internalType\":\"uint256[8][]\",\"name\":\"votingResults\",\"type\":\"uint256[8][]\"}],\"internalType\":\"structProposalsState.ProposalInfo\",\"name\":\"info_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"}],\"name\":\"getProposalStatus\",\"outputs\":[{\"internalType\":\"enumProposalsState.ProposalStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key_\",\"type\":\"string\"}],\"name\":\"getVotingByKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVotings\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys_\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"values_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hide_\",\"type\":\"bool\"}],\"name\":\"hideProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voting_\",\"type\":\"address\"}],\"name\":\"isVoting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalSMTImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"votingName_\",\"type\":\"string\"}],\"name\":\"removeVoting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCallWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userNullifier_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"vote_\",\"type\":\"uint256[]\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProposalsStateABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposalsStateMetaData.ABI instead.
var ProposalsStateABI = ProposalsStateMetaData.ABI

// ProposalsState is an auto generated Go binding around an Ethereum contract.
type ProposalsState struct {
	ProposalsStateCaller     // Read-only binding to the contract
	ProposalsStateTransactor // Write-only binding to the contract
	ProposalsStateFilterer   // Log filterer for contract events
}

// ProposalsStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposalsStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalsStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposalsStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalsStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposalsStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalsStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposalsStateSession struct {
	Contract     *ProposalsState   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProposalsStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposalsStateCallerSession struct {
	Contract *ProposalsStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ProposalsStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposalsStateTransactorSession struct {
	Contract     *ProposalsStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ProposalsStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposalsStateRaw struct {
	Contract *ProposalsState // Generic contract binding to access the raw methods on
}

// ProposalsStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposalsStateCallerRaw struct {
	Contract *ProposalsStateCaller // Generic read-only contract binding to access the raw methods on
}

// ProposalsStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposalsStateTransactorRaw struct {
	Contract *ProposalsStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposalsState creates a new instance of ProposalsState, bound to a specific deployed contract.
func NewProposalsState(address common.Address, backend bind.ContractBackend) (*ProposalsState, error) {
	contract, err := bindProposalsState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProposalsState{ProposalsStateCaller: ProposalsStateCaller{contract: contract}, ProposalsStateTransactor: ProposalsStateTransactor{contract: contract}, ProposalsStateFilterer: ProposalsStateFilterer{contract: contract}}, nil
}

// NewProposalsStateCaller creates a new read-only instance of ProposalsState, bound to a specific deployed contract.
func NewProposalsStateCaller(address common.Address, caller bind.ContractCaller) (*ProposalsStateCaller, error) {
	contract, err := bindProposalsState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalsStateCaller{contract: contract}, nil
}

// NewProposalsStateTransactor creates a new write-only instance of ProposalsState, bound to a specific deployed contract.
func NewProposalsStateTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposalsStateTransactor, error) {
	contract, err := bindProposalsState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalsStateTransactor{contract: contract}, nil
}

// NewProposalsStateFilterer creates a new log filterer instance of ProposalsState, bound to a specific deployed contract.
func NewProposalsStateFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposalsStateFilterer, error) {
	contract, err := bindProposalsState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalsStateFilterer{contract: contract}, nil
}

// bindProposalsState binds a generic wrapper to an already deployed contract.
func bindProposalsState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProposalsStateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalsState *ProposalsStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalsState.Contract.ProposalsStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalsState *ProposalsStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalsState.Contract.ProposalsStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalsState *ProposalsStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalsState.Contract.ProposalsStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalsState *ProposalsStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalsState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalsState *ProposalsStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalsState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalsState *ProposalsStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalsState.Contract.contract.Transact(opts, method, params...)
}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_ProposalsState *ProposalsStateCaller) MAGICID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "MAGIC_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_ProposalsState *ProposalsStateSession) MAGICID() (uint8, error) {
	return _ProposalsState.Contract.MAGICID(&_ProposalsState.CallOpts)
}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_ProposalsState *ProposalsStateCallerSession) MAGICID() (uint8, error) {
	return _ProposalsState.Contract.MAGICID(&_ProposalsState.CallOpts)
}

// MAXIMUMCHOICESPEROPTION is a free data retrieval call binding the contract method 0x299a9b54.
//
// Solidity: function MAXIMUM_CHOICES_PER_OPTION() view returns(uint256)
func (_ProposalsState *ProposalsStateCaller) MAXIMUMCHOICESPEROPTION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "MAXIMUM_CHOICES_PER_OPTION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMCHOICESPEROPTION is a free data retrieval call binding the contract method 0x299a9b54.
//
// Solidity: function MAXIMUM_CHOICES_PER_OPTION() view returns(uint256)
func (_ProposalsState *ProposalsStateSession) MAXIMUMCHOICESPEROPTION() (*big.Int, error) {
	return _ProposalsState.Contract.MAXIMUMCHOICESPEROPTION(&_ProposalsState.CallOpts)
}

// MAXIMUMCHOICESPEROPTION is a free data retrieval call binding the contract method 0x299a9b54.
//
// Solidity: function MAXIMUM_CHOICES_PER_OPTION() view returns(uint256)
func (_ProposalsState *ProposalsStateCallerSession) MAXIMUMCHOICESPEROPTION() (*big.Int, error) {
	return _ProposalsState.Contract.MAXIMUMCHOICESPEROPTION(&_ProposalsState.CallOpts)
}

// MAXIMUMOPTIONS is a free data retrieval call binding the contract method 0x881e3447.
//
// Solidity: function MAXIMUM_OPTIONS() view returns(uint256)
func (_ProposalsState *ProposalsStateCaller) MAXIMUMOPTIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "MAXIMUM_OPTIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMUMOPTIONS is a free data retrieval call binding the contract method 0x881e3447.
//
// Solidity: function MAXIMUM_OPTIONS() view returns(uint256)
func (_ProposalsState *ProposalsStateSession) MAXIMUMOPTIONS() (*big.Int, error) {
	return _ProposalsState.Contract.MAXIMUMOPTIONS(&_ProposalsState.CallOpts)
}

// MAXIMUMOPTIONS is a free data retrieval call binding the contract method 0x881e3447.
//
// Solidity: function MAXIMUM_OPTIONS() view returns(uint256)
func (_ProposalsState *ProposalsStateCallerSession) MAXIMUMOPTIONS() (*big.Int, error) {
	return _ProposalsState.Contract.MAXIMUMOPTIONS(&_ProposalsState.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_ProposalsState *ProposalsStateCaller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "P")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_ProposalsState *ProposalsStateSession) P() (*big.Int, error) {
	return _ProposalsState.Contract.P(&_ProposalsState.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_ProposalsState *ProposalsStateCallerSession) P() (*big.Int, error) {
	return _ProposalsState.Contract.P(&_ProposalsState.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_ProposalsState *ProposalsStateCaller) ChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "chainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_ProposalsState *ProposalsStateSession) ChainName() (string, error) {
	return _ProposalsState.Contract.ChainName(&_ProposalsState.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_ProposalsState *ProposalsStateCallerSession) ChainName() (string, error) {
	return _ProposalsState.Contract.ChainName(&_ProposalsState.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_ProposalsState *ProposalsStateCaller) GetNonce(opts *bind.CallOpts, methodId_ uint8) (*big.Int, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "getNonce", methodId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_ProposalsState *ProposalsStateSession) GetNonce(methodId_ uint8) (*big.Int, error) {
	return _ProposalsState.Contract.GetNonce(&_ProposalsState.CallOpts, methodId_)
}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_ProposalsState *ProposalsStateCallerSession) GetNonce(methodId_ uint8) (*big.Int, error) {
	return _ProposalsState.Contract.GetNonce(&_ProposalsState.CallOpts, methodId_)
}

// GetProposalConfig is a free data retrieval call binding the contract method 0x7d5d687f.
//
// Solidity: function getProposalConfig(uint256 proposalId_) view returns((uint64,uint64,uint256,uint256[],string,address[],bytes[]))
func (_ProposalsState *ProposalsStateCaller) GetProposalConfig(opts *bind.CallOpts, proposalId_ *big.Int) (ProposalsStateProposalConfig, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "getProposalConfig", proposalId_)

	if err != nil {
		return *new(ProposalsStateProposalConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ProposalsStateProposalConfig)).(*ProposalsStateProposalConfig)

	return out0, err

}

// GetProposalConfig is a free data retrieval call binding the contract method 0x7d5d687f.
//
// Solidity: function getProposalConfig(uint256 proposalId_) view returns((uint64,uint64,uint256,uint256[],string,address[],bytes[]))
func (_ProposalsState *ProposalsStateSession) GetProposalConfig(proposalId_ *big.Int) (ProposalsStateProposalConfig, error) {
	return _ProposalsState.Contract.GetProposalConfig(&_ProposalsState.CallOpts, proposalId_)
}

// GetProposalConfig is a free data retrieval call binding the contract method 0x7d5d687f.
//
// Solidity: function getProposalConfig(uint256 proposalId_) view returns((uint64,uint64,uint256,uint256[],string,address[],bytes[]))
func (_ProposalsState *ProposalsStateCallerSession) GetProposalConfig(proposalId_ *big.Int) (ProposalsStateProposalConfig, error) {
	return _ProposalsState.Contract.GetProposalConfig(&_ProposalsState.CallOpts, proposalId_)
}

// GetProposalEventId is a free data retrieval call binding the contract method 0x31e181c5.
//
// Solidity: function getProposalEventId(uint256 proposalId_) view returns(uint256)
func (_ProposalsState *ProposalsStateCaller) GetProposalEventId(opts *bind.CallOpts, proposalId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "getProposalEventId", proposalId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalEventId is a free data retrieval call binding the contract method 0x31e181c5.
//
// Solidity: function getProposalEventId(uint256 proposalId_) view returns(uint256)
func (_ProposalsState *ProposalsStateSession) GetProposalEventId(proposalId_ *big.Int) (*big.Int, error) {
	return _ProposalsState.Contract.GetProposalEventId(&_ProposalsState.CallOpts, proposalId_)
}

// GetProposalEventId is a free data retrieval call binding the contract method 0x31e181c5.
//
// Solidity: function getProposalEventId(uint256 proposalId_) view returns(uint256)
func (_ProposalsState *ProposalsStateCallerSession) GetProposalEventId(proposalId_ *big.Int) (*big.Int, error) {
	return _ProposalsState.Contract.GetProposalEventId(&_ProposalsState.CallOpts, proposalId_)
}

// GetProposalInfo is a free data retrieval call binding the contract method 0xbc903cb8.
//
// Solidity: function getProposalInfo(uint256 proposalId_) view returns((address,uint8,(uint64,uint64,uint256,uint256[],string,address[],bytes[]),uint256[8][]) info_)
func (_ProposalsState *ProposalsStateCaller) GetProposalInfo(opts *bind.CallOpts, proposalId_ *big.Int) (ProposalsStateProposalInfo, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "getProposalInfo", proposalId_)

	if err != nil {
		return *new(ProposalsStateProposalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ProposalsStateProposalInfo)).(*ProposalsStateProposalInfo)

	return out0, err

}

// GetProposalInfo is a free data retrieval call binding the contract method 0xbc903cb8.
//
// Solidity: function getProposalInfo(uint256 proposalId_) view returns((address,uint8,(uint64,uint64,uint256,uint256[],string,address[],bytes[]),uint256[8][]) info_)
func (_ProposalsState *ProposalsStateSession) GetProposalInfo(proposalId_ *big.Int) (ProposalsStateProposalInfo, error) {
	return _ProposalsState.Contract.GetProposalInfo(&_ProposalsState.CallOpts, proposalId_)
}

// GetProposalInfo is a free data retrieval call binding the contract method 0xbc903cb8.
//
// Solidity: function getProposalInfo(uint256 proposalId_) view returns((address,uint8,(uint64,uint64,uint256,uint256[],string,address[],bytes[]),uint256[8][]) info_)
func (_ProposalsState *ProposalsStateCallerSession) GetProposalInfo(proposalId_ *big.Int) (ProposalsStateProposalInfo, error) {
	return _ProposalsState.Contract.GetProposalInfo(&_ProposalsState.CallOpts, proposalId_)
}

// GetProposalStatus is a free data retrieval call binding the contract method 0x401853b7.
//
// Solidity: function getProposalStatus(uint256 proposalId_) view returns(uint8)
func (_ProposalsState *ProposalsStateCaller) GetProposalStatus(opts *bind.CallOpts, proposalId_ *big.Int) (uint8, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "getProposalStatus", proposalId_)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetProposalStatus is a free data retrieval call binding the contract method 0x401853b7.
//
// Solidity: function getProposalStatus(uint256 proposalId_) view returns(uint8)
func (_ProposalsState *ProposalsStateSession) GetProposalStatus(proposalId_ *big.Int) (uint8, error) {
	return _ProposalsState.Contract.GetProposalStatus(&_ProposalsState.CallOpts, proposalId_)
}

// GetProposalStatus is a free data retrieval call binding the contract method 0x401853b7.
//
// Solidity: function getProposalStatus(uint256 proposalId_) view returns(uint8)
func (_ProposalsState *ProposalsStateCallerSession) GetProposalStatus(proposalId_ *big.Int) (uint8, error) {
	return _ProposalsState.Contract.GetProposalStatus(&_ProposalsState.CallOpts, proposalId_)
}

// GetVotingByKey is a free data retrieval call binding the contract method 0xd8720106.
//
// Solidity: function getVotingByKey(string key_) view returns(address)
func (_ProposalsState *ProposalsStateCaller) GetVotingByKey(opts *bind.CallOpts, key_ string) (common.Address, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "getVotingByKey", key_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVotingByKey is a free data retrieval call binding the contract method 0xd8720106.
//
// Solidity: function getVotingByKey(string key_) view returns(address)
func (_ProposalsState *ProposalsStateSession) GetVotingByKey(key_ string) (common.Address, error) {
	return _ProposalsState.Contract.GetVotingByKey(&_ProposalsState.CallOpts, key_)
}

// GetVotingByKey is a free data retrieval call binding the contract method 0xd8720106.
//
// Solidity: function getVotingByKey(string key_) view returns(address)
func (_ProposalsState *ProposalsStateCallerSession) GetVotingByKey(key_ string) (common.Address, error) {
	return _ProposalsState.Contract.GetVotingByKey(&_ProposalsState.CallOpts, key_)
}

// GetVotings is a free data retrieval call binding the contract method 0x05c1112d.
//
// Solidity: function getVotings() view returns(string[] keys_, address[] values_)
func (_ProposalsState *ProposalsStateCaller) GetVotings(opts *bind.CallOpts) (struct {
	Keys   []string
	Values []common.Address
}, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "getVotings")

	outstruct := new(struct {
		Keys   []string
		Values []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Keys = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Values = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetVotings is a free data retrieval call binding the contract method 0x05c1112d.
//
// Solidity: function getVotings() view returns(string[] keys_, address[] values_)
func (_ProposalsState *ProposalsStateSession) GetVotings() (struct {
	Keys   []string
	Values []common.Address
}, error) {
	return _ProposalsState.Contract.GetVotings(&_ProposalsState.CallOpts)
}

// GetVotings is a free data retrieval call binding the contract method 0x05c1112d.
//
// Solidity: function getVotings() view returns(string[] keys_, address[] values_)
func (_ProposalsState *ProposalsStateCallerSession) GetVotings() (struct {
	Keys   []string
	Values []common.Address
}, error) {
	return _ProposalsState.Contract.GetVotings(&_ProposalsState.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ProposalsState *ProposalsStateCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ProposalsState *ProposalsStateSession) Implementation() (common.Address, error) {
	return _ProposalsState.Contract.Implementation(&_ProposalsState.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ProposalsState *ProposalsStateCallerSession) Implementation() (common.Address, error) {
	return _ProposalsState.Contract.Implementation(&_ProposalsState.CallOpts)
}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address voting_) view returns(bool)
func (_ProposalsState *ProposalsStateCaller) IsVoting(opts *bind.CallOpts, voting_ common.Address) (bool, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "isVoting", voting_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address voting_) view returns(bool)
func (_ProposalsState *ProposalsStateSession) IsVoting(voting_ common.Address) (bool, error) {
	return _ProposalsState.Contract.IsVoting(&_ProposalsState.CallOpts, voting_)
}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address voting_) view returns(bool)
func (_ProposalsState *ProposalsStateCallerSession) IsVoting(voting_ common.Address) (bool, error) {
	return _ProposalsState.Contract.IsVoting(&_ProposalsState.CallOpts, voting_)
}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_ProposalsState *ProposalsStateCaller) LastProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "lastProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_ProposalsState *ProposalsStateSession) LastProposalId() (*big.Int, error) {
	return _ProposalsState.Contract.LastProposalId(&_ProposalsState.CallOpts)
}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_ProposalsState *ProposalsStateCallerSession) LastProposalId() (*big.Int, error) {
	return _ProposalsState.Contract.LastProposalId(&_ProposalsState.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProposalsState *ProposalsStateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProposalsState *ProposalsStateSession) Owner() (common.Address, error) {
	return _ProposalsState.Contract.Owner(&_ProposalsState.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProposalsState *ProposalsStateCallerSession) Owner() (common.Address, error) {
	return _ProposalsState.Contract.Owner(&_ProposalsState.CallOpts)
}

// ProposalSMTImpl is a free data retrieval call binding the contract method 0x59917f46.
//
// Solidity: function proposalSMTImpl() view returns(address)
func (_ProposalsState *ProposalsStateCaller) ProposalSMTImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "proposalSMTImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalSMTImpl is a free data retrieval call binding the contract method 0x59917f46.
//
// Solidity: function proposalSMTImpl() view returns(address)
func (_ProposalsState *ProposalsStateSession) ProposalSMTImpl() (common.Address, error) {
	return _ProposalsState.Contract.ProposalSMTImpl(&_ProposalsState.CallOpts)
}

// ProposalSMTImpl is a free data retrieval call binding the contract method 0x59917f46.
//
// Solidity: function proposalSMTImpl() view returns(address)
func (_ProposalsState *ProposalsStateCallerSession) ProposalSMTImpl() (common.Address, error) {
	return _ProposalsState.Contract.ProposalSMTImpl(&_ProposalsState.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProposalsState *ProposalsStateCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProposalsState *ProposalsStateSession) ProxiableUUID() ([32]byte, error) {
	return _ProposalsState.Contract.ProxiableUUID(&_ProposalsState.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProposalsState *ProposalsStateCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ProposalsState.Contract.ProxiableUUID(&_ProposalsState.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_ProposalsState *ProposalsStateCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProposalsState.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_ProposalsState *ProposalsStateSession) Signer() (common.Address, error) {
	return _ProposalsState.Contract.Signer(&_ProposalsState.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_ProposalsState *ProposalsStateCallerSession) Signer() (common.Address, error) {
	return _ProposalsState.Contract.Signer(&_ProposalsState.CallOpts)
}

// ProposalsStateInit is a paid mutator transaction binding the contract method 0xb5697e6a.
//
// Solidity: function __ProposalsState_init(address signer_, string chainName_, address proposalSMTImpl_) returns()
func (_ProposalsState *ProposalsStateTransactor) ProposalsStateInit(opts *bind.TransactOpts, signer_ common.Address, chainName_ string, proposalSMTImpl_ common.Address) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "__ProposalsState_init", signer_, chainName_, proposalSMTImpl_)
}

// ProposalsStateInit is a paid mutator transaction binding the contract method 0xb5697e6a.
//
// Solidity: function __ProposalsState_init(address signer_, string chainName_, address proposalSMTImpl_) returns()
func (_ProposalsState *ProposalsStateSession) ProposalsStateInit(signer_ common.Address, chainName_ string, proposalSMTImpl_ common.Address) (*types.Transaction, error) {
	return _ProposalsState.Contract.ProposalsStateInit(&_ProposalsState.TransactOpts, signer_, chainName_, proposalSMTImpl_)
}

// ProposalsStateInit is a paid mutator transaction binding the contract method 0xb5697e6a.
//
// Solidity: function __ProposalsState_init(address signer_, string chainName_, address proposalSMTImpl_) returns()
func (_ProposalsState *ProposalsStateTransactorSession) ProposalsStateInit(signer_ common.Address, chainName_ string, proposalSMTImpl_ common.Address) (*types.Transaction, error) {
	return _ProposalsState.Contract.ProposalsStateInit(&_ProposalsState.TransactOpts, signer_, chainName_, proposalSMTImpl_)
}

// AddVoting is a paid mutator transaction binding the contract method 0xde947541.
//
// Solidity: function addVoting(string votingName_, address votingAddress_) returns()
func (_ProposalsState *ProposalsStateTransactor) AddVoting(opts *bind.TransactOpts, votingName_ string, votingAddress_ common.Address) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "addVoting", votingName_, votingAddress_)
}

// AddVoting is a paid mutator transaction binding the contract method 0xde947541.
//
// Solidity: function addVoting(string votingName_, address votingAddress_) returns()
func (_ProposalsState *ProposalsStateSession) AddVoting(votingName_ string, votingAddress_ common.Address) (*types.Transaction, error) {
	return _ProposalsState.Contract.AddVoting(&_ProposalsState.TransactOpts, votingName_, votingAddress_)
}

// AddVoting is a paid mutator transaction binding the contract method 0xde947541.
//
// Solidity: function addVoting(string votingName_, address votingAddress_) returns()
func (_ProposalsState *ProposalsStateTransactorSession) AddVoting(votingName_ string, votingAddress_ common.Address) (*types.Transaction, error) {
	return _ProposalsState.Contract.AddVoting(&_ProposalsState.TransactOpts, votingName_, votingAddress_)
}

// ChangeProposalConfig is a paid mutator transaction binding the contract method 0x1cbdbbc0.
//
// Solidity: function changeProposalConfig(uint256 proposalId_, (uint64,uint64,uint256,uint256[],string,address[],bytes[]) newProposalConfig_) returns()
func (_ProposalsState *ProposalsStateTransactor) ChangeProposalConfig(opts *bind.TransactOpts, proposalId_ *big.Int, newProposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "changeProposalConfig", proposalId_, newProposalConfig_)
}

// ChangeProposalConfig is a paid mutator transaction binding the contract method 0x1cbdbbc0.
//
// Solidity: function changeProposalConfig(uint256 proposalId_, (uint64,uint64,uint256,uint256[],string,address[],bytes[]) newProposalConfig_) returns()
func (_ProposalsState *ProposalsStateSession) ChangeProposalConfig(proposalId_ *big.Int, newProposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _ProposalsState.Contract.ChangeProposalConfig(&_ProposalsState.TransactOpts, proposalId_, newProposalConfig_)
}

// ChangeProposalConfig is a paid mutator transaction binding the contract method 0x1cbdbbc0.
//
// Solidity: function changeProposalConfig(uint256 proposalId_, (uint64,uint64,uint256,uint256[],string,address[],bytes[]) newProposalConfig_) returns()
func (_ProposalsState *ProposalsStateTransactorSession) ChangeProposalConfig(proposalId_ *big.Int, newProposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _ProposalsState.Contract.ChangeProposalConfig(&_ProposalsState.TransactOpts, proposalId_, newProposalConfig_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_ProposalsState *ProposalsStateTransactor) ChangeSigner(opts *bind.TransactOpts, newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "changeSigner", newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_ProposalsState *ProposalsStateSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _ProposalsState.Contract.ChangeSigner(&_ProposalsState.TransactOpts, newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_ProposalsState *ProposalsStateTransactorSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _ProposalsState.Contract.ChangeSigner(&_ProposalsState.TransactOpts, newSignerPubKey_, signature_)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x9151b81f.
//
// Solidity: function createProposal((uint64,uint64,uint256,uint256[],string,address[],bytes[]) proposalConfig_) payable returns()
func (_ProposalsState *ProposalsStateTransactor) CreateProposal(opts *bind.TransactOpts, proposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "createProposal", proposalConfig_)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x9151b81f.
//
// Solidity: function createProposal((uint64,uint64,uint256,uint256[],string,address[],bytes[]) proposalConfig_) payable returns()
func (_ProposalsState *ProposalsStateSession) CreateProposal(proposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _ProposalsState.Contract.CreateProposal(&_ProposalsState.TransactOpts, proposalConfig_)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x9151b81f.
//
// Solidity: function createProposal((uint64,uint64,uint256,uint256[],string,address[],bytes[]) proposalConfig_) payable returns()
func (_ProposalsState *ProposalsStateTransactorSession) CreateProposal(proposalConfig_ ProposalsStateProposalConfig) (*types.Transaction, error) {
	return _ProposalsState.Contract.CreateProposal(&_ProposalsState.TransactOpts, proposalConfig_)
}

// HideProposal is a paid mutator transaction binding the contract method 0x50df86a3.
//
// Solidity: function hideProposal(uint256 proposalId_, bool hide_) returns()
func (_ProposalsState *ProposalsStateTransactor) HideProposal(opts *bind.TransactOpts, proposalId_ *big.Int, hide_ bool) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "hideProposal", proposalId_, hide_)
}

// HideProposal is a paid mutator transaction binding the contract method 0x50df86a3.
//
// Solidity: function hideProposal(uint256 proposalId_, bool hide_) returns()
func (_ProposalsState *ProposalsStateSession) HideProposal(proposalId_ *big.Int, hide_ bool) (*types.Transaction, error) {
	return _ProposalsState.Contract.HideProposal(&_ProposalsState.TransactOpts, proposalId_, hide_)
}

// HideProposal is a paid mutator transaction binding the contract method 0x50df86a3.
//
// Solidity: function hideProposal(uint256 proposalId_, bool hide_) returns()
func (_ProposalsState *ProposalsStateTransactorSession) HideProposal(proposalId_ *big.Int, hide_ bool) (*types.Transaction, error) {
	return _ProposalsState.Contract.HideProposal(&_ProposalsState.TransactOpts, proposalId_, hide_)
}

// RemoveVoting is a paid mutator transaction binding the contract method 0x4fcfccd7.
//
// Solidity: function removeVoting(string votingName_) returns()
func (_ProposalsState *ProposalsStateTransactor) RemoveVoting(opts *bind.TransactOpts, votingName_ string) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "removeVoting", votingName_)
}

// RemoveVoting is a paid mutator transaction binding the contract method 0x4fcfccd7.
//
// Solidity: function removeVoting(string votingName_) returns()
func (_ProposalsState *ProposalsStateSession) RemoveVoting(votingName_ string) (*types.Transaction, error) {
	return _ProposalsState.Contract.RemoveVoting(&_ProposalsState.TransactOpts, votingName_)
}

// RemoveVoting is a paid mutator transaction binding the contract method 0x4fcfccd7.
//
// Solidity: function removeVoting(string votingName_) returns()
func (_ProposalsState *ProposalsStateTransactorSession) RemoveVoting(votingName_ string) (*types.Transaction, error) {
	return _ProposalsState.Contract.RemoveVoting(&_ProposalsState.TransactOpts, votingName_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProposalsState *ProposalsStateTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProposalsState *ProposalsStateSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProposalsState.Contract.RenounceOwnership(&_ProposalsState.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProposalsState *ProposalsStateTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProposalsState.Contract.RenounceOwnership(&_ProposalsState.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProposalsState *ProposalsStateTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProposalsState *ProposalsStateSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProposalsState.Contract.TransferOwnership(&_ProposalsState.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProposalsState *ProposalsStateTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProposalsState.Contract.TransferOwnership(&_ProposalsState.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ProposalsState *ProposalsStateTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ProposalsState *ProposalsStateSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ProposalsState.Contract.UpgradeTo(&_ProposalsState.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ProposalsState *ProposalsStateTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ProposalsState.Contract.UpgradeTo(&_ProposalsState.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProposalsState *ProposalsStateTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProposalsState *ProposalsStateSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProposalsState.Contract.UpgradeToAndCall(&_ProposalsState.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProposalsState *ProposalsStateTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProposalsState.Contract.UpgradeToAndCall(&_ProposalsState.TransactOpts, newImplementation, data)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_ProposalsState *ProposalsStateTransactor) UpgradeToAndCallWithProof(opts *bind.TransactOpts, newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "upgradeToAndCallWithProof", newImplementation_, proof_, data_)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_ProposalsState *ProposalsStateSession) UpgradeToAndCallWithProof(newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _ProposalsState.Contract.UpgradeToAndCallWithProof(&_ProposalsState.TransactOpts, newImplementation_, proof_, data_)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_ProposalsState *ProposalsStateTransactorSession) UpgradeToAndCallWithProof(newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _ProposalsState.Contract.UpgradeToAndCallWithProof(&_ProposalsState.TransactOpts, newImplementation_, proof_, data_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_ProposalsState *ProposalsStateTransactor) UpgradeToWithProof(opts *bind.TransactOpts, newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "upgradeToWithProof", newImplementation_, proof_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_ProposalsState *ProposalsStateSession) UpgradeToWithProof(newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _ProposalsState.Contract.UpgradeToWithProof(&_ProposalsState.TransactOpts, newImplementation_, proof_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_ProposalsState *ProposalsStateTransactorSession) UpgradeToWithProof(newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _ProposalsState.Contract.UpgradeToWithProof(&_ProposalsState.TransactOpts, newImplementation_, proof_)
}

// Vote is a paid mutator transaction binding the contract method 0xe1349bd7.
//
// Solidity: function vote(uint256 proposalId_, uint256 userNullifier_, uint256[] vote_) returns()
func (_ProposalsState *ProposalsStateTransactor) Vote(opts *bind.TransactOpts, proposalId_ *big.Int, userNullifier_ *big.Int, vote_ []*big.Int) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "vote", proposalId_, userNullifier_, vote_)
}

// Vote is a paid mutator transaction binding the contract method 0xe1349bd7.
//
// Solidity: function vote(uint256 proposalId_, uint256 userNullifier_, uint256[] vote_) returns()
func (_ProposalsState *ProposalsStateSession) Vote(proposalId_ *big.Int, userNullifier_ *big.Int, vote_ []*big.Int) (*types.Transaction, error) {
	return _ProposalsState.Contract.Vote(&_ProposalsState.TransactOpts, proposalId_, userNullifier_, vote_)
}

// Vote is a paid mutator transaction binding the contract method 0xe1349bd7.
//
// Solidity: function vote(uint256 proposalId_, uint256 userNullifier_, uint256[] vote_) returns()
func (_ProposalsState *ProposalsStateTransactorSession) Vote(proposalId_ *big.Int, userNullifier_ *big.Int, vote_ []*big.Int) (*types.Transaction, error) {
	return _ProposalsState.Contract.Vote(&_ProposalsState.TransactOpts, proposalId_, userNullifier_, vote_)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address recipient_, uint256 amount_) returns()
func (_ProposalsState *ProposalsStateTransactor) WithdrawFunds(opts *bind.TransactOpts, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _ProposalsState.contract.Transact(opts, "withdrawFunds", recipient_, amount_)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address recipient_, uint256 amount_) returns()
func (_ProposalsState *ProposalsStateSession) WithdrawFunds(recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _ProposalsState.Contract.WithdrawFunds(&_ProposalsState.TransactOpts, recipient_, amount_)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address recipient_, uint256 amount_) returns()
func (_ProposalsState *ProposalsStateTransactorSession) WithdrawFunds(recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _ProposalsState.Contract.WithdrawFunds(&_ProposalsState.TransactOpts, recipient_, amount_)
}

// ProposalsStateAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ProposalsState contract.
type ProposalsStateAdminChangedIterator struct {
	Event *ProposalsStateAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProposalsStateAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalsStateAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProposalsStateAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProposalsStateAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalsStateAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalsStateAdminChanged represents a AdminChanged event raised by the ProposalsState contract.
type ProposalsStateAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ProposalsState *ProposalsStateFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ProposalsStateAdminChangedIterator, error) {

	logs, sub, err := _ProposalsState.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ProposalsStateAdminChangedIterator{contract: _ProposalsState.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ProposalsState *ProposalsStateFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ProposalsStateAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ProposalsState.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalsStateAdminChanged)
				if err := _ProposalsState.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ProposalsState *ProposalsStateFilterer) ParseAdminChanged(log types.Log) (*ProposalsStateAdminChanged, error) {
	event := new(ProposalsStateAdminChanged)
	if err := _ProposalsState.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalsStateBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ProposalsState contract.
type ProposalsStateBeaconUpgradedIterator struct {
	Event *ProposalsStateBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProposalsStateBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalsStateBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProposalsStateBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProposalsStateBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalsStateBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalsStateBeaconUpgraded represents a BeaconUpgraded event raised by the ProposalsState contract.
type ProposalsStateBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ProposalsState *ProposalsStateFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ProposalsStateBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ProposalsState.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ProposalsStateBeaconUpgradedIterator{contract: _ProposalsState.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ProposalsState *ProposalsStateFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ProposalsStateBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ProposalsState.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalsStateBeaconUpgraded)
				if err := _ProposalsState.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ProposalsState *ProposalsStateFilterer) ParseBeaconUpgraded(log types.Log) (*ProposalsStateBeaconUpgraded, error) {
	event := new(ProposalsStateBeaconUpgraded)
	if err := _ProposalsState.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalsStateInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ProposalsState contract.
type ProposalsStateInitializedIterator struct {
	Event *ProposalsStateInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProposalsStateInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalsStateInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProposalsStateInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProposalsStateInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalsStateInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalsStateInitialized represents a Initialized event raised by the ProposalsState contract.
type ProposalsStateInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProposalsState *ProposalsStateFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProposalsStateInitializedIterator, error) {

	logs, sub, err := _ProposalsState.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProposalsStateInitializedIterator{contract: _ProposalsState.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProposalsState *ProposalsStateFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProposalsStateInitialized) (event.Subscription, error) {

	logs, sub, err := _ProposalsState.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalsStateInitialized)
				if err := _ProposalsState.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProposalsState *ProposalsStateFilterer) ParseInitialized(log types.Log) (*ProposalsStateInitialized, error) {
	event := new(ProposalsStateInitialized)
	if err := _ProposalsState.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalsStateOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProposalsState contract.
type ProposalsStateOwnershipTransferredIterator struct {
	Event *ProposalsStateOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProposalsStateOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalsStateOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProposalsStateOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProposalsStateOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalsStateOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalsStateOwnershipTransferred represents a OwnershipTransferred event raised by the ProposalsState contract.
type ProposalsStateOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProposalsState *ProposalsStateFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProposalsStateOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProposalsState.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProposalsStateOwnershipTransferredIterator{contract: _ProposalsState.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProposalsState *ProposalsStateFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProposalsStateOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProposalsState.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalsStateOwnershipTransferred)
				if err := _ProposalsState.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProposalsState *ProposalsStateFilterer) ParseOwnershipTransferred(log types.Log) (*ProposalsStateOwnershipTransferred, error) {
	event := new(ProposalsStateOwnershipTransferred)
	if err := _ProposalsState.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalsStateProposalConfigChangedIterator is returned from FilterProposalConfigChanged and is used to iterate over the raw logs and unpacked data for ProposalConfigChanged events raised by the ProposalsState contract.
type ProposalsStateProposalConfigChangedIterator struct {
	Event *ProposalsStateProposalConfigChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProposalsStateProposalConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalsStateProposalConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProposalsStateProposalConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProposalsStateProposalConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalsStateProposalConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalsStateProposalConfigChanged represents a ProposalConfigChanged event raised by the ProposalsState contract.
type ProposalsStateProposalConfigChanged struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalConfigChanged is a free log retrieval operation binding the contract event 0xa9cc646240fc6ba1b4b124e96765839b67cd0e2e698942d5d5948a36c7b998d5.
//
// Solidity: event ProposalConfigChanged(uint256 indexed proposalId)
func (_ProposalsState *ProposalsStateFilterer) FilterProposalConfigChanged(opts *bind.FilterOpts, proposalId []*big.Int) (*ProposalsStateProposalConfigChangedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _ProposalsState.contract.FilterLogs(opts, "ProposalConfigChanged", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &ProposalsStateProposalConfigChangedIterator{contract: _ProposalsState.contract, event: "ProposalConfigChanged", logs: logs, sub: sub}, nil
}

// WatchProposalConfigChanged is a free log subscription operation binding the contract event 0xa9cc646240fc6ba1b4b124e96765839b67cd0e2e698942d5d5948a36c7b998d5.
//
// Solidity: event ProposalConfigChanged(uint256 indexed proposalId)
func (_ProposalsState *ProposalsStateFilterer) WatchProposalConfigChanged(opts *bind.WatchOpts, sink chan<- *ProposalsStateProposalConfigChanged, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _ProposalsState.contract.WatchLogs(opts, "ProposalConfigChanged", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalsStateProposalConfigChanged)
				if err := _ProposalsState.contract.UnpackLog(event, "ProposalConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalConfigChanged is a log parse operation binding the contract event 0xa9cc646240fc6ba1b4b124e96765839b67cd0e2e698942d5d5948a36c7b998d5.
//
// Solidity: event ProposalConfigChanged(uint256 indexed proposalId)
func (_ProposalsState *ProposalsStateFilterer) ParseProposalConfigChanged(log types.Log) (*ProposalsStateProposalConfigChanged, error) {
	event := new(ProposalsStateProposalConfigChanged)
	if err := _ProposalsState.contract.UnpackLog(event, "ProposalConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalsStateProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the ProposalsState contract.
type ProposalsStateProposalCreatedIterator struct {
	Event *ProposalsStateProposalCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProposalsStateProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalsStateProposalCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProposalsStateProposalCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProposalsStateProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalsStateProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalsStateProposalCreated represents a ProposalCreated event raised by the ProposalsState contract.
type ProposalsStateProposalCreated struct {
	ProposalId  *big.Int
	ProposalSMT common.Address
	FundAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x3417b456fad6209c73445d5efd446d686e75e4560f0f50c13b5a5cde976447b4.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, address proposalSMT, uint256 fundAmount)
func (_ProposalsState *ProposalsStateFilterer) FilterProposalCreated(opts *bind.FilterOpts, proposalId []*big.Int) (*ProposalsStateProposalCreatedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _ProposalsState.contract.FilterLogs(opts, "ProposalCreated", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &ProposalsStateProposalCreatedIterator{contract: _ProposalsState.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x3417b456fad6209c73445d5efd446d686e75e4560f0f50c13b5a5cde976447b4.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, address proposalSMT, uint256 fundAmount)
func (_ProposalsState *ProposalsStateFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *ProposalsStateProposalCreated, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _ProposalsState.contract.WatchLogs(opts, "ProposalCreated", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalsStateProposalCreated)
				if err := _ProposalsState.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCreated is a log parse operation binding the contract event 0x3417b456fad6209c73445d5efd446d686e75e4560f0f50c13b5a5cde976447b4.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, address proposalSMT, uint256 fundAmount)
func (_ProposalsState *ProposalsStateFilterer) ParseProposalCreated(log types.Log) (*ProposalsStateProposalCreated, error) {
	event := new(ProposalsStateProposalCreated)
	if err := _ProposalsState.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalsStateProposalHiddenIterator is returned from FilterProposalHidden and is used to iterate over the raw logs and unpacked data for ProposalHidden events raised by the ProposalsState contract.
type ProposalsStateProposalHiddenIterator struct {
	Event *ProposalsStateProposalHidden // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProposalsStateProposalHiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalsStateProposalHidden)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProposalsStateProposalHidden)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProposalsStateProposalHiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalsStateProposalHiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalsStateProposalHidden represents a ProposalHidden event raised by the ProposalsState contract.
type ProposalsStateProposalHidden struct {
	ProposalId *big.Int
	Hide       bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalHidden is a free log retrieval operation binding the contract event 0x19f289534dc0123bc632723043c2eaf220105017950314fd948774c69282dc93.
//
// Solidity: event ProposalHidden(uint256 indexed proposalId, bool hide)
func (_ProposalsState *ProposalsStateFilterer) FilterProposalHidden(opts *bind.FilterOpts, proposalId []*big.Int) (*ProposalsStateProposalHiddenIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _ProposalsState.contract.FilterLogs(opts, "ProposalHidden", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &ProposalsStateProposalHiddenIterator{contract: _ProposalsState.contract, event: "ProposalHidden", logs: logs, sub: sub}, nil
}

// WatchProposalHidden is a free log subscription operation binding the contract event 0x19f289534dc0123bc632723043c2eaf220105017950314fd948774c69282dc93.
//
// Solidity: event ProposalHidden(uint256 indexed proposalId, bool hide)
func (_ProposalsState *ProposalsStateFilterer) WatchProposalHidden(opts *bind.WatchOpts, sink chan<- *ProposalsStateProposalHidden, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _ProposalsState.contract.WatchLogs(opts, "ProposalHidden", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalsStateProposalHidden)
				if err := _ProposalsState.contract.UnpackLog(event, "ProposalHidden", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalHidden is a log parse operation binding the contract event 0x19f289534dc0123bc632723043c2eaf220105017950314fd948774c69282dc93.
//
// Solidity: event ProposalHidden(uint256 indexed proposalId, bool hide)
func (_ProposalsState *ProposalsStateFilterer) ParseProposalHidden(log types.Log) (*ProposalsStateProposalHidden, error) {
	event := new(ProposalsStateProposalHidden)
	if err := _ProposalsState.contract.UnpackLog(event, "ProposalHidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalsStateUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ProposalsState contract.
type ProposalsStateUpgradedIterator struct {
	Event *ProposalsStateUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProposalsStateUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalsStateUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProposalsStateUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProposalsStateUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalsStateUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalsStateUpgraded represents a Upgraded event raised by the ProposalsState contract.
type ProposalsStateUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProposalsState *ProposalsStateFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ProposalsStateUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ProposalsState.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ProposalsStateUpgradedIterator{contract: _ProposalsState.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProposalsState *ProposalsStateFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ProposalsStateUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ProposalsState.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalsStateUpgraded)
				if err := _ProposalsState.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProposalsState *ProposalsStateFilterer) ParseUpgraded(log types.Log) (*ProposalsStateUpgraded, error) {
	event := new(ProposalsStateUpgraded)
	if err := _ProposalsState.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalsStateVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the ProposalsState contract.
type ProposalsStateVoteCastIterator struct {
	Event *ProposalsStateVoteCast // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProposalsStateVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalsStateVoteCast)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProposalsStateVoteCast)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProposalsStateVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalsStateVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalsStateVoteCast represents a VoteCast event raised by the ProposalsState contract.
type ProposalsStateVoteCast struct {
	ProposalId    *big.Int
	UserNullifier *big.Int
	Vote          []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0x82e882ecc8f666d65e8120d1fb3859261652f808e5001ae2f169e5ea1bf5035c.
//
// Solidity: event VoteCast(uint256 indexed proposalId, uint256 indexed userNullifier, uint256[] vote)
func (_ProposalsState *ProposalsStateFilterer) FilterVoteCast(opts *bind.FilterOpts, proposalId []*big.Int, userNullifier []*big.Int) (*ProposalsStateVoteCastIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var userNullifierRule []interface{}
	for _, userNullifierItem := range userNullifier {
		userNullifierRule = append(userNullifierRule, userNullifierItem)
	}

	logs, sub, err := _ProposalsState.contract.FilterLogs(opts, "VoteCast", proposalIdRule, userNullifierRule)
	if err != nil {
		return nil, err
	}
	return &ProposalsStateVoteCastIterator{contract: _ProposalsState.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0x82e882ecc8f666d65e8120d1fb3859261652f808e5001ae2f169e5ea1bf5035c.
//
// Solidity: event VoteCast(uint256 indexed proposalId, uint256 indexed userNullifier, uint256[] vote)
func (_ProposalsState *ProposalsStateFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *ProposalsStateVoteCast, proposalId []*big.Int, userNullifier []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var userNullifierRule []interface{}
	for _, userNullifierItem := range userNullifier {
		userNullifierRule = append(userNullifierRule, userNullifierItem)
	}

	logs, sub, err := _ProposalsState.contract.WatchLogs(opts, "VoteCast", proposalIdRule, userNullifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalsStateVoteCast)
				if err := _ProposalsState.contract.UnpackLog(event, "VoteCast", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCast is a log parse operation binding the contract event 0x82e882ecc8f666d65e8120d1fb3859261652f808e5001ae2f169e5ea1bf5035c.
//
// Solidity: event VoteCast(uint256 indexed proposalId, uint256 indexed userNullifier, uint256[] vote)
func (_ProposalsState *ProposalsStateFilterer) ParseVoteCast(log types.Log) (*ProposalsStateVoteCast, error) {
	event := new(ProposalsStateVoteCast)
	if err := _ProposalsState.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
