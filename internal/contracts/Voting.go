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
)

// IVotingVotingCounters is an auto generated low-level Go binding around an user-defined struct.
type IVotingVotingCounters struct {
	VotesCount *big.Int
}

// IVotingVotingParams is an auto generated low-level Go binding around an user-defined struct.
type IVotingVotingParams struct {
	Registration common.Address
	Remark       string
	VotingStart  *big.Int
	VotingPeriod *big.Int
	Candidates   [][32]byte
}

// IVotingVotingValues is an auto generated low-level Go binding around an user-defined struct.
type IVotingVotingValues struct {
	VotingStartTime *big.Int
	VotingEndTime   *big.Int
	Candidates      [][32]byte
}

// VerifierHelperProofPoints is an auto generated low-level Go binding around an user-defined struct.
type VerifierHelperProofPoints struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voteVerifier_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nullifierHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"candidate\",\"type\":\"bytes32\"}],\"name\":\"UserVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIRegistration\",\"name\":\"registration\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votingStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"candidates\",\"type\":\"bytes32[]\"}],\"indexed\":false,\"internalType\":\"structIVoting.VotingParams\",\"name\":\"votingParams\",\"type\":\"tuple\"}],\"name\":\"VotingInitialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_CANDIDATES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIRegistration\",\"name\":\"registration\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votingStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"candidates\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"votingParams_\",\"type\":\"tuple\"}],\"name\":\"__Voting_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposalStatus\",\"outputs\":[{\"internalType\":\"enumIVoting.VotingStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistrationAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nullifiers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registration\",\"outputs\":[{\"internalType\":\"contractIRegistration\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nullifierHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"candidate_\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"proof_\",\"type\":\"tuple\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"votesPerCandidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"candidates\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIVoting.VotingValues\",\"name\":\"values\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votesCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// MAXCANDIDATES is a free data retrieval call binding the contract method 0xf0786096.
//
// Solidity: function MAX_CANDIDATES() view returns(uint256)
func (_Voting *VotingCaller) MAXCANDIDATES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "MAX_CANDIDATES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCANDIDATES is a free data retrieval call binding the contract method 0xf0786096.
//
// Solidity: function MAX_CANDIDATES() view returns(uint256)
func (_Voting *VotingSession) MAXCANDIDATES() (*big.Int, error) {
	return _Voting.Contract.MAXCANDIDATES(&_Voting.CallOpts)
}

// MAXCANDIDATES is a free data retrieval call binding the contract method 0xf0786096.
//
// Solidity: function MAX_CANDIDATES() view returns(uint256)
func (_Voting *VotingCallerSession) MAXCANDIDATES() (*big.Int, error) {
	return _Voting.Contract.MAXCANDIDATES(&_Voting.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x1a0478d5.
//
// Solidity: function candidates(bytes32 ) view returns(bool)
func (_Voting *VotingCaller) Candidates(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "candidates", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Candidates is a free data retrieval call binding the contract method 0x1a0478d5.
//
// Solidity: function candidates(bytes32 ) view returns(bool)
func (_Voting *VotingSession) Candidates(arg0 [32]byte) (bool, error) {
	return _Voting.Contract.Candidates(&_Voting.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x1a0478d5.
//
// Solidity: function candidates(bytes32 ) view returns(bool)
func (_Voting *VotingCallerSession) Candidates(arg0 [32]byte) (bool, error) {
	return _Voting.Contract.Candidates(&_Voting.CallOpts, arg0)
}

// GetProposalStatus is a free data retrieval call binding the contract method 0xcfeb652e.
//
// Solidity: function getProposalStatus() view returns(uint8)
func (_Voting *VotingCaller) GetProposalStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getProposalStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetProposalStatus is a free data retrieval call binding the contract method 0xcfeb652e.
//
// Solidity: function getProposalStatus() view returns(uint8)
func (_Voting *VotingSession) GetProposalStatus() (uint8, error) {
	return _Voting.Contract.GetProposalStatus(&_Voting.CallOpts)
}

// GetProposalStatus is a free data retrieval call binding the contract method 0xcfeb652e.
//
// Solidity: function getProposalStatus() view returns(uint8)
func (_Voting *VotingCallerSession) GetProposalStatus() (uint8, error) {
	return _Voting.Contract.GetProposalStatus(&_Voting.CallOpts)
}

// GetRegistrationAddresses is a free data retrieval call binding the contract method 0xe8188f97.
//
// Solidity: function getRegistrationAddresses() view returns(address[])
func (_Voting *VotingCaller) GetRegistrationAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getRegistrationAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegistrationAddresses is a free data retrieval call binding the contract method 0xe8188f97.
//
// Solidity: function getRegistrationAddresses() view returns(address[])
func (_Voting *VotingSession) GetRegistrationAddresses() ([]common.Address, error) {
	return _Voting.Contract.GetRegistrationAddresses(&_Voting.CallOpts)
}

// GetRegistrationAddresses is a free data retrieval call binding the contract method 0xe8188f97.
//
// Solidity: function getRegistrationAddresses() view returns(address[])
func (_Voting *VotingCallerSession) GetRegistrationAddresses() ([]common.Address, error) {
	return _Voting.Contract.GetRegistrationAddresses(&_Voting.CallOpts)
}

// Nullifiers is a free data retrieval call binding the contract method 0x2997e86b.
//
// Solidity: function nullifiers(bytes32 ) view returns(bool)
func (_Voting *VotingCaller) Nullifiers(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "nullifiers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Nullifiers is a free data retrieval call binding the contract method 0x2997e86b.
//
// Solidity: function nullifiers(bytes32 ) view returns(bool)
func (_Voting *VotingSession) Nullifiers(arg0 [32]byte) (bool, error) {
	return _Voting.Contract.Nullifiers(&_Voting.CallOpts, arg0)
}

// Nullifiers is a free data retrieval call binding the contract method 0x2997e86b.
//
// Solidity: function nullifiers(bytes32 ) view returns(bool)
func (_Voting *VotingCallerSession) Nullifiers(arg0 [32]byte) (bool, error) {
	return _Voting.Contract.Nullifiers(&_Voting.CallOpts, arg0)
}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(address)
func (_Voting *VotingCaller) Registration(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "registration")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(address)
func (_Voting *VotingSession) Registration() (common.Address, error) {
	return _Voting.Contract.Registration(&_Voting.CallOpts)
}

// Registration is a free data retrieval call binding the contract method 0x443bd1d0.
//
// Solidity: function registration() view returns(address)
func (_Voting *VotingCallerSession) Registration() (common.Address, error) {
	return _Voting.Contract.Registration(&_Voting.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voting *VotingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voting *VotingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Voting.Contract.SupportsInterface(&_Voting.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Voting *VotingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Voting.Contract.SupportsInterface(&_Voting.CallOpts, interfaceId)
}

// VoteVerifier is a free data retrieval call binding the contract method 0x99d74c71.
//
// Solidity: function voteVerifier() view returns(address)
func (_Voting *VotingCaller) VoteVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "voteVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteVerifier is a free data retrieval call binding the contract method 0x99d74c71.
//
// Solidity: function voteVerifier() view returns(address)
func (_Voting *VotingSession) VoteVerifier() (common.Address, error) {
	return _Voting.Contract.VoteVerifier(&_Voting.CallOpts)
}

// VoteVerifier is a free data retrieval call binding the contract method 0x99d74c71.
//
// Solidity: function voteVerifier() view returns(address)
func (_Voting *VotingCallerSession) VoteVerifier() (common.Address, error) {
	return _Voting.Contract.VoteVerifier(&_Voting.CallOpts)
}

// VotesPerCandidate is a free data retrieval call binding the contract method 0xe0d38faf.
//
// Solidity: function votesPerCandidate(bytes32 ) view returns(uint256)
func (_Voting *VotingCaller) VotesPerCandidate(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "votesPerCandidate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotesPerCandidate is a free data retrieval call binding the contract method 0xe0d38faf.
//
// Solidity: function votesPerCandidate(bytes32 ) view returns(uint256)
func (_Voting *VotingSession) VotesPerCandidate(arg0 [32]byte) (*big.Int, error) {
	return _Voting.Contract.VotesPerCandidate(&_Voting.CallOpts, arg0)
}

// VotesPerCandidate is a free data retrieval call binding the contract method 0xe0d38faf.
//
// Solidity: function votesPerCandidate(bytes32 ) view returns(uint256)
func (_Voting *VotingCallerSession) VotesPerCandidate(arg0 [32]byte) (*big.Int, error) {
	return _Voting.Contract.VotesPerCandidate(&_Voting.CallOpts, arg0)
}

// VotingInfo is a free data retrieval call binding the contract method 0x19e7998c.
//
// Solidity: function votingInfo() view returns(string remark, (uint256,uint256,bytes32[]) values, (uint256) counters)
func (_Voting *VotingCaller) VotingInfo(opts *bind.CallOpts) (struct {
	Remark   string
	Values   IVotingVotingValues
	Counters IVotingVotingCounters
}, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "votingInfo")

	outstruct := new(struct {
		Remark   string
		Values   IVotingVotingValues
		Counters IVotingVotingCounters
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Remark = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Values = *abi.ConvertType(out[1], new(IVotingVotingValues)).(*IVotingVotingValues)
	outstruct.Counters = *abi.ConvertType(out[2], new(IVotingVotingCounters)).(*IVotingVotingCounters)

	return *outstruct, err

}

// VotingInfo is a free data retrieval call binding the contract method 0x19e7998c.
//
// Solidity: function votingInfo() view returns(string remark, (uint256,uint256,bytes32[]) values, (uint256) counters)
func (_Voting *VotingSession) VotingInfo() (struct {
	Remark   string
	Values   IVotingVotingValues
	Counters IVotingVotingCounters
}, error) {
	return _Voting.Contract.VotingInfo(&_Voting.CallOpts)
}

// VotingInfo is a free data retrieval call binding the contract method 0x19e7998c.
//
// Solidity: function votingInfo() view returns(string remark, (uint256,uint256,bytes32[]) values, (uint256) counters)
func (_Voting *VotingCallerSession) VotingInfo() (struct {
	Remark   string
	Values   IVotingVotingValues
	Counters IVotingVotingCounters
}, error) {
	return _Voting.Contract.VotingInfo(&_Voting.CallOpts)
}

// VotingInit is a paid mutator transaction binding the contract method 0xf8bd2ace.
//
// Solidity: function __Voting_init((address,string,uint256,uint256,bytes32[]) votingParams_) returns()
func (_Voting *VotingTransactor) VotingInit(opts *bind.TransactOpts, votingParams_ IVotingVotingParams) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "__Voting_init", votingParams_)
}

// VotingInit is a paid mutator transaction binding the contract method 0xf8bd2ace.
//
// Solidity: function __Voting_init((address,string,uint256,uint256,bytes32[]) votingParams_) returns()
func (_Voting *VotingSession) VotingInit(votingParams_ IVotingVotingParams) (*types.Transaction, error) {
	return _Voting.Contract.VotingInit(&_Voting.TransactOpts, votingParams_)
}

// VotingInit is a paid mutator transaction binding the contract method 0xf8bd2ace.
//
// Solidity: function __Voting_init((address,string,uint256,uint256,bytes32[]) votingParams_) returns()
func (_Voting *VotingTransactorSession) VotingInit(votingParams_ IVotingVotingParams) (*types.Transaction, error) {
	return _Voting.Contract.VotingInit(&_Voting.TransactOpts, votingParams_)
}

// Vote is a paid mutator transaction binding the contract method 0x0d9eaa43.
//
// Solidity: function vote(bytes32 root_, bytes32 nullifierHash_, bytes32 candidate_, (uint256[2],uint256[2][2],uint256[2]) proof_) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, root_ [32]byte, nullifierHash_ [32]byte, candidate_ [32]byte, proof_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", root_, nullifierHash_, candidate_, proof_)
}

// Vote is a paid mutator transaction binding the contract method 0x0d9eaa43.
//
// Solidity: function vote(bytes32 root_, bytes32 nullifierHash_, bytes32 candidate_, (uint256[2],uint256[2][2],uint256[2]) proof_) returns()
func (_Voting *VotingSession) Vote(root_ [32]byte, nullifierHash_ [32]byte, candidate_ [32]byte, proof_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, root_, nullifierHash_, candidate_, proof_)
}

// Vote is a paid mutator transaction binding the contract method 0x0d9eaa43.
//
// Solidity: function vote(bytes32 root_, bytes32 nullifierHash_, bytes32 candidate_, (uint256[2],uint256[2][2],uint256[2]) proof_) returns()
func (_Voting *VotingTransactorSession) Vote(root_ [32]byte, nullifierHash_ [32]byte, candidate_ [32]byte, proof_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, root_, nullifierHash_, candidate_, proof_)
}

// VotingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Voting contract.
type VotingInitializedIterator struct {
	Event *VotingInitialized // Event containing the contract specifics and raw log

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
func (it *VotingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingInitialized)
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
		it.Event = new(VotingInitialized)
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
func (it *VotingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingInitialized represents a Initialized event raised by the Voting contract.
type VotingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voting *VotingFilterer) FilterInitialized(opts *bind.FilterOpts) (*VotingInitializedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VotingInitializedIterator{contract: _Voting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Voting *VotingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VotingInitialized) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingInitialized)
				if err := _Voting.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Voting *VotingFilterer) ParseInitialized(log types.Log) (*VotingInitialized, error) {
	event := new(VotingInitialized)
	if err := _Voting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingUserVotedIterator is returned from FilterUserVoted and is used to iterate over the raw logs and unpacked data for UserVoted events raised by the Voting contract.
type VotingUserVotedIterator struct {
	Event *VotingUserVoted // Event containing the contract specifics and raw log

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
func (it *VotingUserVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingUserVoted)
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
		it.Event = new(VotingUserVoted)
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
func (it *VotingUserVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingUserVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingUserVoted represents a UserVoted event raised by the Voting contract.
type VotingUserVoted struct {
	User          common.Address
	Root          [32]byte
	NullifierHash [32]byte
	Candidate     [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserVoted is a free log retrieval operation binding the contract event 0xd7b642ac96388fe3d629dbcf233cde0723967951c334d29a2b8db1c476182f30.
//
// Solidity: event UserVoted(address indexed user, bytes32 root, bytes32 nullifierHash, bytes32 candidate)
func (_Voting *VotingFilterer) FilterUserVoted(opts *bind.FilterOpts, user []common.Address) (*VotingUserVotedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "UserVoted", userRule)
	if err != nil {
		return nil, err
	}
	return &VotingUserVotedIterator{contract: _Voting.contract, event: "UserVoted", logs: logs, sub: sub}, nil
}

// WatchUserVoted is a free log subscription operation binding the contract event 0xd7b642ac96388fe3d629dbcf233cde0723967951c334d29a2b8db1c476182f30.
//
// Solidity: event UserVoted(address indexed user, bytes32 root, bytes32 nullifierHash, bytes32 candidate)
func (_Voting *VotingFilterer) WatchUserVoted(opts *bind.WatchOpts, sink chan<- *VotingUserVoted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "UserVoted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingUserVoted)
				if err := _Voting.contract.UnpackLog(event, "UserVoted", log); err != nil {
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

// ParseUserVoted is a log parse operation binding the contract event 0xd7b642ac96388fe3d629dbcf233cde0723967951c334d29a2b8db1c476182f30.
//
// Solidity: event UserVoted(address indexed user, bytes32 root, bytes32 nullifierHash, bytes32 candidate)
func (_Voting *VotingFilterer) ParseUserVoted(log types.Log) (*VotingUserVoted, error) {
	event := new(VotingUserVoted)
	if err := _Voting.contract.UnpackLog(event, "UserVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingVotingInitializedIterator is returned from FilterVotingInitialized and is used to iterate over the raw logs and unpacked data for VotingInitialized events raised by the Voting contract.
type VotingVotingInitializedIterator struct {
	Event *VotingVotingInitialized // Event containing the contract specifics and raw log

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
func (it *VotingVotingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingVotingInitialized)
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
		it.Event = new(VotingVotingInitialized)
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
func (it *VotingVotingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingVotingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingVotingInitialized represents a VotingInitialized event raised by the Voting contract.
type VotingVotingInitialized struct {
	Proposer     common.Address
	VotingParams IVotingVotingParams
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVotingInitialized is a free log retrieval operation binding the contract event 0xb9bbd4e09a06db4d1f06fe2ebe360738047ec15a1f30d6daeb02b0241ebcd290.
//
// Solidity: event VotingInitialized(address indexed proposer, (address,string,uint256,uint256,bytes32[]) votingParams)
func (_Voting *VotingFilterer) FilterVotingInitialized(opts *bind.FilterOpts, proposer []common.Address) (*VotingVotingInitializedIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Voting.contract.FilterLogs(opts, "VotingInitialized", proposerRule)
	if err != nil {
		return nil, err
	}
	return &VotingVotingInitializedIterator{contract: _Voting.contract, event: "VotingInitialized", logs: logs, sub: sub}, nil
}

// WatchVotingInitialized is a free log subscription operation binding the contract event 0xb9bbd4e09a06db4d1f06fe2ebe360738047ec15a1f30d6daeb02b0241ebcd290.
//
// Solidity: event VotingInitialized(address indexed proposer, (address,string,uint256,uint256,bytes32[]) votingParams)
func (_Voting *VotingFilterer) WatchVotingInitialized(opts *bind.WatchOpts, sink chan<- *VotingVotingInitialized, proposer []common.Address) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Voting.contract.WatchLogs(opts, "VotingInitialized", proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingVotingInitialized)
				if err := _Voting.contract.UnpackLog(event, "VotingInitialized", log); err != nil {
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

// ParseVotingInitialized is a log parse operation binding the contract event 0xb9bbd4e09a06db4d1f06fe2ebe360738047ec15a1f30d6daeb02b0241ebcd290.
//
// Solidity: event VotingInitialized(address indexed proposer, (address,string,uint256,uint256,bytes32[]) votingParams)
func (_Voting *VotingFilterer) ParseVotingInitialized(log types.Log) (*VotingVotingInitialized, error) {
	event := new(VotingVotingInitialized)
	if err := _Voting.contract.UnpackLog(event, "VotingInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
