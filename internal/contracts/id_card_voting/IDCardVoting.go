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

// BaseVotingUserData is an auto generated low-level Go binding around an user-defined struct.
type BaseVotingUserData struct {
	Nullifier                 *big.Int
	Citizenship               *big.Int
	IdentityCreationTimestamp *big.Int
}

// AQueryProofExecutorProofPoints is an auto generated low-level Go binding around an user-defined struct.
type AQueryProofExecutorProofPoints struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// BaseVotingProposalRules is an auto generated low-level Go binding around an user-defined struct.
type BaseVotingProposalRules struct {
	Selector                            *big.Int
	CitizenshipWhitelist                []*big.Int
	IdentityCreationTimestampUpperBound *big.Int
	IdentityCounterUpperBound           *big.Int
	Sex                                 *big.Int
	BirthDateLowerbound                 *big.Int
	BirthDateUpperbound                 *big.Int
	ExpirationDateLowerBound            *big.Int
}

// IDCardVotingMetaData contains all meta data concerning the IDCardVoting contract.
var IDCardVotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"FailedToCallVerifyProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pubSignals\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structAQueryProofExecutor.ProofPoints\",\"name\":\"zkPoints\",\"type\":\"tuple\"}],\"name\":\"InvalidCircomProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"parsedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"}],\"name\":\"InvalidDate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"pubSignals\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"zkPoints\",\"type\":\"bytes\"}],\"name\":\"InvalidNoirProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registrationSMT\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"registrationRoot\",\"type\":\"bytes32\"}],\"name\":\"InvalidRegistrationRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registrationSMT\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"registrationRoot\",\"type\":\"bytes32\"}],\"name\":\"InvalidRegistrationRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pubSignals_\",\"type\":\"uint256[]\"}],\"name\":\"InvalidZKProof\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"IDENTITY_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registrationSMT_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposalsState_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"votingVerifier_\",\"type\":\"address\"}],\"name\":\"__IDCardVoting_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structAQueryProofExecutor.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"zkPoints_\",\"type\":\"bytes\"}],\"name\":\"executeNoir\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structAQueryProofExecutor.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"executeTD1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"zkPoints_\",\"type\":\"bytes\"}],\"name\":\"executeTD1Noir\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"}],\"name\":\"getProposalRules\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"citizenshipWhitelist\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"identityCreationTimestampUpperBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identityCounterUpperBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"birthDateLowerbound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"birthDateUpperbound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateLowerBound\",\"type\":\"uint256\"}],\"internalType\":\"structBaseVoting.ProposalRules\",\"name\":\"proposalRules_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"}],\"name\":\"getPublicSignals\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"publicSignals\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"}],\"name\":\"getPublicSignalsTD1\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"publicSignals\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistrationSMT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalsState\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IDCardVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use IDCardVotingMetaData.ABI instead.
var IDCardVotingABI = IDCardVotingMetaData.ABI

// IDCardVoting is an auto generated Go binding around an Ethereum contract.
type IDCardVoting struct {
	IDCardVotingCaller     // Read-only binding to the contract
	IDCardVotingTransactor // Write-only binding to the contract
	IDCardVotingFilterer   // Log filterer for contract events
}

// IDCardVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDCardVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDCardVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDCardVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDCardVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDCardVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDCardVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDCardVotingSession struct {
	Contract     *IDCardVoting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDCardVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDCardVotingCallerSession struct {
	Contract *IDCardVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IDCardVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDCardVotingTransactorSession struct {
	Contract     *IDCardVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IDCardVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDCardVotingRaw struct {
	Contract *IDCardVoting // Generic contract binding to access the raw methods on
}

// IDCardVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDCardVotingCallerRaw struct {
	Contract *IDCardVotingCaller // Generic read-only contract binding to access the raw methods on
}

// IDCardVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDCardVotingTransactorRaw struct {
	Contract *IDCardVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDCardVoting creates a new instance of IDCardVoting, bound to a specific deployed contract.
func NewIDCardVoting(address common.Address, backend bind.ContractBackend) (*IDCardVoting, error) {
	contract, err := bindIDCardVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDCardVoting{IDCardVotingCaller: IDCardVotingCaller{contract: contract}, IDCardVotingTransactor: IDCardVotingTransactor{contract: contract}, IDCardVotingFilterer: IDCardVotingFilterer{contract: contract}}, nil
}

// NewIDCardVotingCaller creates a new read-only instance of IDCardVoting, bound to a specific deployed contract.
func NewIDCardVotingCaller(address common.Address, caller bind.ContractCaller) (*IDCardVotingCaller, error) {
	contract, err := bindIDCardVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDCardVotingCaller{contract: contract}, nil
}

// NewIDCardVotingTransactor creates a new write-only instance of IDCardVoting, bound to a specific deployed contract.
func NewIDCardVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*IDCardVotingTransactor, error) {
	contract, err := bindIDCardVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDCardVotingTransactor{contract: contract}, nil
}

// NewIDCardVotingFilterer creates a new log filterer instance of IDCardVoting, bound to a specific deployed contract.
func NewIDCardVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*IDCardVotingFilterer, error) {
	contract, err := bindIDCardVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDCardVotingFilterer{contract: contract}, nil
}

// bindIDCardVoting binds a generic wrapper to an already deployed contract.
func bindIDCardVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDCardVotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDCardVoting *IDCardVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDCardVoting.Contract.IDCardVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDCardVoting *IDCardVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDCardVoting.Contract.IDCardVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDCardVoting *IDCardVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDCardVoting.Contract.IDCardVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDCardVoting *IDCardVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDCardVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDCardVoting *IDCardVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDCardVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDCardVoting *IDCardVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDCardVoting.Contract.contract.Transact(opts, method, params...)
}

// IDENTITYLIMIT is a free data retrieval call binding the contract method 0x7995c0f3.
//
// Solidity: function IDENTITY_LIMIT() view returns(uint256)
func (_IDCardVoting *IDCardVotingCaller) IDENTITYLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDCardVoting.contract.Call(opts, &out, "IDENTITY_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IDENTITYLIMIT is a free data retrieval call binding the contract method 0x7995c0f3.
//
// Solidity: function IDENTITY_LIMIT() view returns(uint256)
func (_IDCardVoting *IDCardVotingSession) IDENTITYLIMIT() (*big.Int, error) {
	return _IDCardVoting.Contract.IDENTITYLIMIT(&_IDCardVoting.CallOpts)
}

// IDENTITYLIMIT is a free data retrieval call binding the contract method 0x7995c0f3.
//
// Solidity: function IDENTITY_LIMIT() view returns(uint256)
func (_IDCardVoting *IDCardVotingCallerSession) IDENTITYLIMIT() (*big.Int, error) {
	return _IDCardVoting.Contract.IDENTITYLIMIT(&_IDCardVoting.CallOpts)
}

// GetProposalRules is a free data retrieval call binding the contract method 0xa99f1dca.
//
// Solidity: function getProposalRules(uint256 proposalId_) view returns((uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256) proposalRules_)
func (_IDCardVoting *IDCardVotingCaller) GetProposalRules(opts *bind.CallOpts, proposalId_ *big.Int) (BaseVotingProposalRules, error) {
	var out []interface{}
	err := _IDCardVoting.contract.Call(opts, &out, "getProposalRules", proposalId_)

	if err != nil {
		return *new(BaseVotingProposalRules), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseVotingProposalRules)).(*BaseVotingProposalRules)

	return out0, err

}

// GetProposalRules is a free data retrieval call binding the contract method 0xa99f1dca.
//
// Solidity: function getProposalRules(uint256 proposalId_) view returns((uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256) proposalRules_)
func (_IDCardVoting *IDCardVotingSession) GetProposalRules(proposalId_ *big.Int) (BaseVotingProposalRules, error) {
	return _IDCardVoting.Contract.GetProposalRules(&_IDCardVoting.CallOpts, proposalId_)
}

// GetProposalRules is a free data retrieval call binding the contract method 0xa99f1dca.
//
// Solidity: function getProposalRules(uint256 proposalId_) view returns((uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256) proposalRules_)
func (_IDCardVoting *IDCardVotingCallerSession) GetProposalRules(proposalId_ *big.Int) (BaseVotingProposalRules, error) {
	return _IDCardVoting.Contract.GetProposalRules(&_IDCardVoting.CallOpts, proposalId_)
}

// GetPublicSignals is a free data retrieval call binding the contract method 0x5e192e3d.
//
// Solidity: function getPublicSignals(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(bytes32[] publicSignals)
func (_IDCardVoting *IDCardVotingCaller) GetPublicSignals(opts *bind.CallOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) ([][32]byte, error) {
	var out []interface{}
	err := _IDCardVoting.contract.Call(opts, &out, "getPublicSignals", registrationRoot_, currentDate_, userPayload_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPublicSignals is a free data retrieval call binding the contract method 0x5e192e3d.
//
// Solidity: function getPublicSignals(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(bytes32[] publicSignals)
func (_IDCardVoting *IDCardVotingSession) GetPublicSignals(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) ([][32]byte, error) {
	return _IDCardVoting.Contract.GetPublicSignals(&_IDCardVoting.CallOpts, registrationRoot_, currentDate_, userPayload_)
}

// GetPublicSignals is a free data retrieval call binding the contract method 0x5e192e3d.
//
// Solidity: function getPublicSignals(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(bytes32[] publicSignals)
func (_IDCardVoting *IDCardVotingCallerSession) GetPublicSignals(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) ([][32]byte, error) {
	return _IDCardVoting.Contract.GetPublicSignals(&_IDCardVoting.CallOpts, registrationRoot_, currentDate_, userPayload_)
}

// GetPublicSignalsTD1 is a free data retrieval call binding the contract method 0x8529e377.
//
// Solidity: function getPublicSignalsTD1(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(bytes32[] publicSignals)
func (_IDCardVoting *IDCardVotingCaller) GetPublicSignalsTD1(opts *bind.CallOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) ([][32]byte, error) {
	var out []interface{}
	err := _IDCardVoting.contract.Call(opts, &out, "getPublicSignalsTD1", registrationRoot_, currentDate_, userPayload_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPublicSignalsTD1 is a free data retrieval call binding the contract method 0x8529e377.
//
// Solidity: function getPublicSignalsTD1(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(bytes32[] publicSignals)
func (_IDCardVoting *IDCardVotingSession) GetPublicSignalsTD1(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) ([][32]byte, error) {
	return _IDCardVoting.Contract.GetPublicSignalsTD1(&_IDCardVoting.CallOpts, registrationRoot_, currentDate_, userPayload_)
}

// GetPublicSignalsTD1 is a free data retrieval call binding the contract method 0x8529e377.
//
// Solidity: function getPublicSignalsTD1(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(bytes32[] publicSignals)
func (_IDCardVoting *IDCardVotingCallerSession) GetPublicSignalsTD1(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) ([][32]byte, error) {
	return _IDCardVoting.Contract.GetPublicSignalsTD1(&_IDCardVoting.CallOpts, registrationRoot_, currentDate_, userPayload_)
}

// GetRegistrationSMT is a free data retrieval call binding the contract method 0x796164c0.
//
// Solidity: function getRegistrationSMT() view returns(address)
func (_IDCardVoting *IDCardVotingCaller) GetRegistrationSMT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDCardVoting.contract.Call(opts, &out, "getRegistrationSMT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistrationSMT is a free data retrieval call binding the contract method 0x796164c0.
//
// Solidity: function getRegistrationSMT() view returns(address)
func (_IDCardVoting *IDCardVotingSession) GetRegistrationSMT() (common.Address, error) {
	return _IDCardVoting.Contract.GetRegistrationSMT(&_IDCardVoting.CallOpts)
}

// GetRegistrationSMT is a free data retrieval call binding the contract method 0x796164c0.
//
// Solidity: function getRegistrationSMT() view returns(address)
func (_IDCardVoting *IDCardVotingCallerSession) GetRegistrationSMT() (common.Address, error) {
	return _IDCardVoting.Contract.GetRegistrationSMT(&_IDCardVoting.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IDCardVoting *IDCardVotingCaller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDCardVoting.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IDCardVoting *IDCardVotingSession) GetVerifier() (common.Address, error) {
	return _IDCardVoting.Contract.GetVerifier(&_IDCardVoting.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IDCardVoting *IDCardVotingCallerSession) GetVerifier() (common.Address, error) {
	return _IDCardVoting.Contract.GetVerifier(&_IDCardVoting.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IDCardVoting *IDCardVotingCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDCardVoting.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IDCardVoting *IDCardVotingSession) Implementation() (common.Address, error) {
	return _IDCardVoting.Contract.Implementation(&_IDCardVoting.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IDCardVoting *IDCardVotingCallerSession) Implementation() (common.Address, error) {
	return _IDCardVoting.Contract.Implementation(&_IDCardVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IDCardVoting *IDCardVotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDCardVoting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IDCardVoting *IDCardVotingSession) Owner() (common.Address, error) {
	return _IDCardVoting.Contract.Owner(&_IDCardVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IDCardVoting *IDCardVotingCallerSession) Owner() (common.Address, error) {
	return _IDCardVoting.Contract.Owner(&_IDCardVoting.CallOpts)
}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_IDCardVoting *IDCardVotingCaller) ProposalsState(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDCardVoting.contract.Call(opts, &out, "proposalsState")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_IDCardVoting *IDCardVotingSession) ProposalsState() (common.Address, error) {
	return _IDCardVoting.Contract.ProposalsState(&_IDCardVoting.CallOpts)
}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_IDCardVoting *IDCardVotingCallerSession) ProposalsState() (common.Address, error) {
	return _IDCardVoting.Contract.ProposalsState(&_IDCardVoting.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IDCardVoting *IDCardVotingCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IDCardVoting.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IDCardVoting *IDCardVotingSession) ProxiableUUID() ([32]byte, error) {
	return _IDCardVoting.Contract.ProxiableUUID(&_IDCardVoting.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IDCardVoting *IDCardVotingCallerSession) ProxiableUUID() ([32]byte, error) {
	return _IDCardVoting.Contract.ProxiableUUID(&_IDCardVoting.CallOpts)
}

// IDCardVotingInit is a paid mutator transaction binding the contract method 0xfc3a3fa9.
//
// Solidity: function __IDCardVoting_init(address registrationSMT_, address proposalsState_, address votingVerifier_) returns()
func (_IDCardVoting *IDCardVotingTransactor) IDCardVotingInit(opts *bind.TransactOpts, registrationSMT_ common.Address, proposalsState_ common.Address, votingVerifier_ common.Address) (*types.Transaction, error) {
	return _IDCardVoting.contract.Transact(opts, "__IDCardVoting_init", registrationSMT_, proposalsState_, votingVerifier_)
}

// IDCardVotingInit is a paid mutator transaction binding the contract method 0xfc3a3fa9.
//
// Solidity: function __IDCardVoting_init(address registrationSMT_, address proposalsState_, address votingVerifier_) returns()
func (_IDCardVoting *IDCardVotingSession) IDCardVotingInit(registrationSMT_ common.Address, proposalsState_ common.Address, votingVerifier_ common.Address) (*types.Transaction, error) {
	return _IDCardVoting.Contract.IDCardVotingInit(&_IDCardVoting.TransactOpts, registrationSMT_, proposalsState_, votingVerifier_)
}

// IDCardVotingInit is a paid mutator transaction binding the contract method 0xfc3a3fa9.
//
// Solidity: function __IDCardVoting_init(address registrationSMT_, address proposalsState_, address votingVerifier_) returns()
func (_IDCardVoting *IDCardVotingTransactorSession) IDCardVotingInit(registrationSMT_ common.Address, proposalsState_ common.Address, votingVerifier_ common.Address) (*types.Transaction, error) {
	return _IDCardVoting.Contract.IDCardVotingInit(&_IDCardVoting.TransactOpts, registrationSMT_, proposalsState_, votingVerifier_)
}

// Execute is a paid mutator transaction binding the contract method 0xe4ab0833.
//
// Solidity: function execute(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_IDCardVoting *IDCardVotingTransactor) Execute(opts *bind.TransactOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ AQueryProofExecutorProofPoints) (*types.Transaction, error) {
	return _IDCardVoting.contract.Transact(opts, "execute", registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// Execute is a paid mutator transaction binding the contract method 0xe4ab0833.
//
// Solidity: function execute(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_IDCardVoting *IDCardVotingSession) Execute(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ AQueryProofExecutorProofPoints) (*types.Transaction, error) {
	return _IDCardVoting.Contract.Execute(&_IDCardVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// Execute is a paid mutator transaction binding the contract method 0xe4ab0833.
//
// Solidity: function execute(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_IDCardVoting *IDCardVotingTransactorSession) Execute(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ AQueryProofExecutorProofPoints) (*types.Transaction, error) {
	return _IDCardVoting.Contract.Execute(&_IDCardVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteNoir is a paid mutator transaction binding the contract method 0x6effdd9f.
//
// Solidity: function executeNoir(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, bytes zkPoints_) returns()
func (_IDCardVoting *IDCardVotingTransactor) ExecuteNoir(opts *bind.TransactOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _IDCardVoting.contract.Transact(opts, "executeNoir", registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteNoir is a paid mutator transaction binding the contract method 0x6effdd9f.
//
// Solidity: function executeNoir(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, bytes zkPoints_) returns()
func (_IDCardVoting *IDCardVotingSession) ExecuteNoir(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _IDCardVoting.Contract.ExecuteNoir(&_IDCardVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteNoir is a paid mutator transaction binding the contract method 0x6effdd9f.
//
// Solidity: function executeNoir(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, bytes zkPoints_) returns()
func (_IDCardVoting *IDCardVotingTransactorSession) ExecuteNoir(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _IDCardVoting.Contract.ExecuteNoir(&_IDCardVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteTD1 is a paid mutator transaction binding the contract method 0x8208f9af.
//
// Solidity: function executeTD1(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_IDCardVoting *IDCardVotingTransactor) ExecuteTD1(opts *bind.TransactOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ AQueryProofExecutorProofPoints) (*types.Transaction, error) {
	return _IDCardVoting.contract.Transact(opts, "executeTD1", registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteTD1 is a paid mutator transaction binding the contract method 0x8208f9af.
//
// Solidity: function executeTD1(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_IDCardVoting *IDCardVotingSession) ExecuteTD1(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ AQueryProofExecutorProofPoints) (*types.Transaction, error) {
	return _IDCardVoting.Contract.ExecuteTD1(&_IDCardVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteTD1 is a paid mutator transaction binding the contract method 0x8208f9af.
//
// Solidity: function executeTD1(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_IDCardVoting *IDCardVotingTransactorSession) ExecuteTD1(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ AQueryProofExecutorProofPoints) (*types.Transaction, error) {
	return _IDCardVoting.Contract.ExecuteTD1(&_IDCardVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteTD1Noir is a paid mutator transaction binding the contract method 0x3c559c59.
//
// Solidity: function executeTD1Noir(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, bytes zkPoints_) returns()
func (_IDCardVoting *IDCardVotingTransactor) ExecuteTD1Noir(opts *bind.TransactOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _IDCardVoting.contract.Transact(opts, "executeTD1Noir", registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteTD1Noir is a paid mutator transaction binding the contract method 0x3c559c59.
//
// Solidity: function executeTD1Noir(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, bytes zkPoints_) returns()
func (_IDCardVoting *IDCardVotingSession) ExecuteTD1Noir(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _IDCardVoting.Contract.ExecuteTD1Noir(&_IDCardVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteTD1Noir is a paid mutator transaction binding the contract method 0x3c559c59.
//
// Solidity: function executeTD1Noir(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, bytes zkPoints_) returns()
func (_IDCardVoting *IDCardVotingTransactorSession) ExecuteTD1Noir(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _IDCardVoting.Contract.ExecuteTD1Noir(&_IDCardVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IDCardVoting *IDCardVotingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDCardVoting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IDCardVoting *IDCardVotingSession) RenounceOwnership() (*types.Transaction, error) {
	return _IDCardVoting.Contract.RenounceOwnership(&_IDCardVoting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IDCardVoting *IDCardVotingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IDCardVoting.Contract.RenounceOwnership(&_IDCardVoting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IDCardVoting *IDCardVotingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IDCardVoting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IDCardVoting *IDCardVotingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IDCardVoting.Contract.TransferOwnership(&_IDCardVoting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IDCardVoting *IDCardVotingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IDCardVoting.Contract.TransferOwnership(&_IDCardVoting.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_IDCardVoting *IDCardVotingTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _IDCardVoting.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_IDCardVoting *IDCardVotingSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _IDCardVoting.Contract.UpgradeTo(&_IDCardVoting.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_IDCardVoting *IDCardVotingTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _IDCardVoting.Contract.UpgradeTo(&_IDCardVoting.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IDCardVoting *IDCardVotingTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IDCardVoting.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IDCardVoting *IDCardVotingSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IDCardVoting.Contract.UpgradeToAndCall(&_IDCardVoting.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IDCardVoting *IDCardVotingTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IDCardVoting.Contract.UpgradeToAndCall(&_IDCardVoting.TransactOpts, newImplementation, data)
}

// IDCardVotingAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the IDCardVoting contract.
type IDCardVotingAdminChangedIterator struct {
	Event *IDCardVotingAdminChanged // Event containing the contract specifics and raw log

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
func (it *IDCardVotingAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDCardVotingAdminChanged)
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
		it.Event = new(IDCardVotingAdminChanged)
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
func (it *IDCardVotingAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDCardVotingAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDCardVotingAdminChanged represents a AdminChanged event raised by the IDCardVoting contract.
type IDCardVotingAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_IDCardVoting *IDCardVotingFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*IDCardVotingAdminChangedIterator, error) {

	logs, sub, err := _IDCardVoting.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &IDCardVotingAdminChangedIterator{contract: _IDCardVoting.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_IDCardVoting *IDCardVotingFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *IDCardVotingAdminChanged) (event.Subscription, error) {

	logs, sub, err := _IDCardVoting.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDCardVotingAdminChanged)
				if err := _IDCardVoting.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_IDCardVoting *IDCardVotingFilterer) ParseAdminChanged(log types.Log) (*IDCardVotingAdminChanged, error) {
	event := new(IDCardVotingAdminChanged)
	if err := _IDCardVoting.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDCardVotingBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the IDCardVoting contract.
type IDCardVotingBeaconUpgradedIterator struct {
	Event *IDCardVotingBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *IDCardVotingBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDCardVotingBeaconUpgraded)
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
		it.Event = new(IDCardVotingBeaconUpgraded)
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
func (it *IDCardVotingBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDCardVotingBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDCardVotingBeaconUpgraded represents a BeaconUpgraded event raised by the IDCardVoting contract.
type IDCardVotingBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_IDCardVoting *IDCardVotingFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*IDCardVotingBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _IDCardVoting.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &IDCardVotingBeaconUpgradedIterator{contract: _IDCardVoting.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_IDCardVoting *IDCardVotingFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *IDCardVotingBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _IDCardVoting.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDCardVotingBeaconUpgraded)
				if err := _IDCardVoting.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_IDCardVoting *IDCardVotingFilterer) ParseBeaconUpgraded(log types.Log) (*IDCardVotingBeaconUpgraded, error) {
	event := new(IDCardVotingBeaconUpgraded)
	if err := _IDCardVoting.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDCardVotingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IDCardVoting contract.
type IDCardVotingInitializedIterator struct {
	Event *IDCardVotingInitialized // Event containing the contract specifics and raw log

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
func (it *IDCardVotingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDCardVotingInitialized)
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
		it.Event = new(IDCardVotingInitialized)
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
func (it *IDCardVotingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDCardVotingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDCardVotingInitialized represents a Initialized event raised by the IDCardVoting contract.
type IDCardVotingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IDCardVoting *IDCardVotingFilterer) FilterInitialized(opts *bind.FilterOpts) (*IDCardVotingInitializedIterator, error) {

	logs, sub, err := _IDCardVoting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IDCardVotingInitializedIterator{contract: _IDCardVoting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IDCardVoting *IDCardVotingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IDCardVotingInitialized) (event.Subscription, error) {

	logs, sub, err := _IDCardVoting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDCardVotingInitialized)
				if err := _IDCardVoting.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_IDCardVoting *IDCardVotingFilterer) ParseInitialized(log types.Log) (*IDCardVotingInitialized, error) {
	event := new(IDCardVotingInitialized)
	if err := _IDCardVoting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDCardVotingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IDCardVoting contract.
type IDCardVotingOwnershipTransferredIterator struct {
	Event *IDCardVotingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IDCardVotingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDCardVotingOwnershipTransferred)
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
		it.Event = new(IDCardVotingOwnershipTransferred)
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
func (it *IDCardVotingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDCardVotingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDCardVotingOwnershipTransferred represents a OwnershipTransferred event raised by the IDCardVoting contract.
type IDCardVotingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IDCardVoting *IDCardVotingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IDCardVotingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IDCardVoting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IDCardVotingOwnershipTransferredIterator{contract: _IDCardVoting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IDCardVoting *IDCardVotingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IDCardVotingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IDCardVoting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDCardVotingOwnershipTransferred)
				if err := _IDCardVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IDCardVoting *IDCardVotingFilterer) ParseOwnershipTransferred(log types.Log) (*IDCardVotingOwnershipTransferred, error) {
	event := new(IDCardVotingOwnershipTransferred)
	if err := _IDCardVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDCardVotingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the IDCardVoting contract.
type IDCardVotingUpgradedIterator struct {
	Event *IDCardVotingUpgraded // Event containing the contract specifics and raw log

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
func (it *IDCardVotingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDCardVotingUpgraded)
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
		it.Event = new(IDCardVotingUpgraded)
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
func (it *IDCardVotingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDCardVotingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDCardVotingUpgraded represents a Upgraded event raised by the IDCardVoting contract.
type IDCardVotingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IDCardVoting *IDCardVotingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*IDCardVotingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IDCardVoting.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &IDCardVotingUpgradedIterator{contract: _IDCardVoting.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IDCardVoting *IDCardVotingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *IDCardVotingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IDCardVoting.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDCardVotingUpgraded)
				if err := _IDCardVoting.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_IDCardVoting *IDCardVotingFilterer) ParseUpgraded(log types.Log) (*IDCardVotingUpgraded, error) {
	event := new(IDCardVotingUpgraded)
	if err := _IDCardVoting.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
