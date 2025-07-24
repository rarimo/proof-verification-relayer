// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package noirvoting

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

// NoirVotingMetaData contains all meta data concerning the NoirVoting contract.
var NoirVotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"FailedToCallVerifyProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pubSignals\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structAQueryProofExecutor.ProofPoints\",\"name\":\"zkPoints\",\"type\":\"tuple\"}],\"name\":\"InvalidCircomProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"parsedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTimestamp\",\"type\":\"uint256\"}],\"name\":\"InvalidDate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"pubSignals\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"zkPoints\",\"type\":\"bytes\"}],\"name\":\"InvalidNoirProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registrationSMT\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"registrationRoot\",\"type\":\"bytes32\"}],\"name\":\"InvalidRegistrationRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pubSignals_\",\"type\":\"uint256[]\"}],\"name\":\"InvalidZKProof\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"IDENTITY_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registrationSMT_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposalsState_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"votingVerifier_\",\"type\":\"address\"}],\"name\":\"__NoirIDVoting_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"}],\"name\":\"_afterVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"}],\"name\":\"_beforeVerify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"}],\"name\":\"_buildPublicSignals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structAQueryProofExecutor.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"zkPoints_\",\"type\":\"bytes\"}],\"name\":\"executeNoir\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"}],\"name\":\"getProposalRules\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"selector\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"citizenshipWhitelist\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"identityCreationTimestampUpperBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identityCounterUpperBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"birthDateLowerbound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"birthDateUpperbound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDateLowerBound\",\"type\":\"uint256\"}],\"internalType\":\"structBaseVoting.ProposalRules\",\"name\":\"proposalRules_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userPayload_\",\"type\":\"bytes\"}],\"name\":\"getPublicSignals\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"publicSignals\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistrationSMT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalsState\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NoirVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use NoirVotingMetaData.ABI instead.
var NoirVotingABI = NoirVotingMetaData.ABI

// NoirVoting is an auto generated Go binding around an Ethereum contract.
type NoirVoting struct {
	NoirVotingCaller     // Read-only binding to the contract
	NoirVotingTransactor // Write-only binding to the contract
	NoirVotingFilterer   // Log filterer for contract events
}

// NoirVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type NoirVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoirVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NoirVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoirVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NoirVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoirVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NoirVotingSession struct {
	Contract     *NoirVoting       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NoirVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NoirVotingCallerSession struct {
	Contract *NoirVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NoirVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NoirVotingTransactorSession struct {
	Contract     *NoirVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NoirVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type NoirVotingRaw struct {
	Contract *NoirVoting // Generic contract binding to access the raw methods on
}

// NoirVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NoirVotingCallerRaw struct {
	Contract *NoirVotingCaller // Generic read-only contract binding to access the raw methods on
}

// NoirVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NoirVotingTransactorRaw struct {
	Contract *NoirVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNoirVoting creates a new instance of NoirVoting, bound to a specific deployed contract.
func NewNoirVoting(address common.Address, backend bind.ContractBackend) (*NoirVoting, error) {
	contract, err := bindNoirVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoirVoting{NoirVotingCaller: NoirVotingCaller{contract: contract}, NoirVotingTransactor: NoirVotingTransactor{contract: contract}, NoirVotingFilterer: NoirVotingFilterer{contract: contract}}, nil
}

// NewNoirVotingCaller creates a new read-only instance of NoirVoting, bound to a specific deployed contract.
func NewNoirVotingCaller(address common.Address, caller bind.ContractCaller) (*NoirVotingCaller, error) {
	contract, err := bindNoirVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoirVotingCaller{contract: contract}, nil
}

// NewNoirVotingTransactor creates a new write-only instance of NoirVoting, bound to a specific deployed contract.
func NewNoirVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*NoirVotingTransactor, error) {
	contract, err := bindNoirVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoirVotingTransactor{contract: contract}, nil
}

// NewNoirVotingFilterer creates a new log filterer instance of NoirVoting, bound to a specific deployed contract.
func NewNoirVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*NoirVotingFilterer, error) {
	contract, err := bindNoirVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoirVotingFilterer{contract: contract}, nil
}

// bindNoirVoting binds a generic wrapper to an already deployed contract.
func bindNoirVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NoirVotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoirVoting *NoirVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoirVoting.Contract.NoirVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoirVoting *NoirVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoirVoting.Contract.NoirVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoirVoting *NoirVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoirVoting.Contract.NoirVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoirVoting *NoirVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoirVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoirVoting *NoirVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoirVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoirVoting *NoirVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoirVoting.Contract.contract.Transact(opts, method, params...)
}

// IDENTITYLIMIT is a free data retrieval call binding the contract method 0x7995c0f3.
//
// Solidity: function IDENTITY_LIMIT() view returns(uint256)
func (_NoirVoting *NoirVotingCaller) IDENTITYLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NoirVoting.contract.Call(opts, &out, "IDENTITY_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IDENTITYLIMIT is a free data retrieval call binding the contract method 0x7995c0f3.
//
// Solidity: function IDENTITY_LIMIT() view returns(uint256)
func (_NoirVoting *NoirVotingSession) IDENTITYLIMIT() (*big.Int, error) {
	return _NoirVoting.Contract.IDENTITYLIMIT(&_NoirVoting.CallOpts)
}

// IDENTITYLIMIT is a free data retrieval call binding the contract method 0x7995c0f3.
//
// Solidity: function IDENTITY_LIMIT() view returns(uint256)
func (_NoirVoting *NoirVotingCallerSession) IDENTITYLIMIT() (*big.Int, error) {
	return _NoirVoting.Contract.IDENTITYLIMIT(&_NoirVoting.CallOpts)
}

// BuildPublicSignals is a free data retrieval call binding the contract method 0xdd29625a.
//
// Solidity: function _buildPublicSignals(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(uint256)
func (_NoirVoting *NoirVotingCaller) BuildPublicSignals(opts *bind.CallOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) (*big.Int, error) {
	var out []interface{}
	err := _NoirVoting.contract.Call(opts, &out, "_buildPublicSignals", registrationRoot_, currentDate_, userPayload_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuildPublicSignals is a free data retrieval call binding the contract method 0xdd29625a.
//
// Solidity: function _buildPublicSignals(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(uint256)
func (_NoirVoting *NoirVotingSession) BuildPublicSignals(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) (*big.Int, error) {
	return _NoirVoting.Contract.BuildPublicSignals(&_NoirVoting.CallOpts, registrationRoot_, currentDate_, userPayload_)
}

// BuildPublicSignals is a free data retrieval call binding the contract method 0xdd29625a.
//
// Solidity: function _buildPublicSignals(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(uint256)
func (_NoirVoting *NoirVotingCallerSession) BuildPublicSignals(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) (*big.Int, error) {
	return _NoirVoting.Contract.BuildPublicSignals(&_NoirVoting.CallOpts, registrationRoot_, currentDate_, userPayload_)
}

// GetProposalRules is a free data retrieval call binding the contract method 0xa99f1dca.
//
// Solidity: function getProposalRules(uint256 proposalId_) view returns((uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256) proposalRules_)
func (_NoirVoting *NoirVotingCaller) GetProposalRules(opts *bind.CallOpts, proposalId_ *big.Int) (BaseVotingProposalRules, error) {
	var out []interface{}
	err := _NoirVoting.contract.Call(opts, &out, "getProposalRules", proposalId_)

	if err != nil {
		return *new(BaseVotingProposalRules), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseVotingProposalRules)).(*BaseVotingProposalRules)

	return out0, err

}

// GetProposalRules is a free data retrieval call binding the contract method 0xa99f1dca.
//
// Solidity: function getProposalRules(uint256 proposalId_) view returns((uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256) proposalRules_)
func (_NoirVoting *NoirVotingSession) GetProposalRules(proposalId_ *big.Int) (BaseVotingProposalRules, error) {
	return _NoirVoting.Contract.GetProposalRules(&_NoirVoting.CallOpts, proposalId_)
}

// GetProposalRules is a free data retrieval call binding the contract method 0xa99f1dca.
//
// Solidity: function getProposalRules(uint256 proposalId_) view returns((uint256,uint256[],uint256,uint256,uint256,uint256,uint256,uint256) proposalRules_)
func (_NoirVoting *NoirVotingCallerSession) GetProposalRules(proposalId_ *big.Int) (BaseVotingProposalRules, error) {
	return _NoirVoting.Contract.GetProposalRules(&_NoirVoting.CallOpts, proposalId_)
}

// GetPublicSignals is a free data retrieval call binding the contract method 0x5e192e3d.
//
// Solidity: function getPublicSignals(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(bytes32[] publicSignals)
func (_NoirVoting *NoirVotingCaller) GetPublicSignals(opts *bind.CallOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) ([][32]byte, error) {
	var out []interface{}
	err := _NoirVoting.contract.Call(opts, &out, "getPublicSignals", registrationRoot_, currentDate_, userPayload_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPublicSignals is a free data retrieval call binding the contract method 0x5e192e3d.
//
// Solidity: function getPublicSignals(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(bytes32[] publicSignals)
func (_NoirVoting *NoirVotingSession) GetPublicSignals(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) ([][32]byte, error) {
	return _NoirVoting.Contract.GetPublicSignals(&_NoirVoting.CallOpts, registrationRoot_, currentDate_, userPayload_)
}

// GetPublicSignals is a free data retrieval call binding the contract method 0x5e192e3d.
//
// Solidity: function getPublicSignals(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) view returns(bytes32[] publicSignals)
func (_NoirVoting *NoirVotingCallerSession) GetPublicSignals(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) ([][32]byte, error) {
	return _NoirVoting.Contract.GetPublicSignals(&_NoirVoting.CallOpts, registrationRoot_, currentDate_, userPayload_)
}

// GetRegistrationSMT is a free data retrieval call binding the contract method 0x796164c0.
//
// Solidity: function getRegistrationSMT() view returns(address)
func (_NoirVoting *NoirVotingCaller) GetRegistrationSMT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoirVoting.contract.Call(opts, &out, "getRegistrationSMT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistrationSMT is a free data retrieval call binding the contract method 0x796164c0.
//
// Solidity: function getRegistrationSMT() view returns(address)
func (_NoirVoting *NoirVotingSession) GetRegistrationSMT() (common.Address, error) {
	return _NoirVoting.Contract.GetRegistrationSMT(&_NoirVoting.CallOpts)
}

// GetRegistrationSMT is a free data retrieval call binding the contract method 0x796164c0.
//
// Solidity: function getRegistrationSMT() view returns(address)
func (_NoirVoting *NoirVotingCallerSession) GetRegistrationSMT() (common.Address, error) {
	return _NoirVoting.Contract.GetRegistrationSMT(&_NoirVoting.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_NoirVoting *NoirVotingCaller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoirVoting.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_NoirVoting *NoirVotingSession) GetVerifier() (common.Address, error) {
	return _NoirVoting.Contract.GetVerifier(&_NoirVoting.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_NoirVoting *NoirVotingCallerSession) GetVerifier() (common.Address, error) {
	return _NoirVoting.Contract.GetVerifier(&_NoirVoting.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NoirVoting *NoirVotingCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoirVoting.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NoirVoting *NoirVotingSession) Implementation() (common.Address, error) {
	return _NoirVoting.Contract.Implementation(&_NoirVoting.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NoirVoting *NoirVotingCallerSession) Implementation() (common.Address, error) {
	return _NoirVoting.Contract.Implementation(&_NoirVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NoirVoting *NoirVotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoirVoting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NoirVoting *NoirVotingSession) Owner() (common.Address, error) {
	return _NoirVoting.Contract.Owner(&_NoirVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NoirVoting *NoirVotingCallerSession) Owner() (common.Address, error) {
	return _NoirVoting.Contract.Owner(&_NoirVoting.CallOpts)
}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_NoirVoting *NoirVotingCaller) ProposalsState(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NoirVoting.contract.Call(opts, &out, "proposalsState")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_NoirVoting *NoirVotingSession) ProposalsState() (common.Address, error) {
	return _NoirVoting.Contract.ProposalsState(&_NoirVoting.CallOpts)
}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_NoirVoting *NoirVotingCallerSession) ProposalsState() (common.Address, error) {
	return _NoirVoting.Contract.ProposalsState(&_NoirVoting.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NoirVoting *NoirVotingCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NoirVoting.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NoirVoting *NoirVotingSession) ProxiableUUID() ([32]byte, error) {
	return _NoirVoting.Contract.ProxiableUUID(&_NoirVoting.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NoirVoting *NoirVotingCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NoirVoting.Contract.ProxiableUUID(&_NoirVoting.CallOpts)
}

// NoirIDVotingInit is a paid mutator transaction binding the contract method 0x6941af14.
//
// Solidity: function __NoirIDVoting_init(address registrationSMT_, address proposalsState_, address votingVerifier_) returns()
func (_NoirVoting *NoirVotingTransactor) NoirIDVotingInit(opts *bind.TransactOpts, registrationSMT_ common.Address, proposalsState_ common.Address, votingVerifier_ common.Address) (*types.Transaction, error) {
	return _NoirVoting.contract.Transact(opts, "__NoirIDVoting_init", registrationSMT_, proposalsState_, votingVerifier_)
}

// NoirIDVotingInit is a paid mutator transaction binding the contract method 0x6941af14.
//
// Solidity: function __NoirIDVoting_init(address registrationSMT_, address proposalsState_, address votingVerifier_) returns()
func (_NoirVoting *NoirVotingSession) NoirIDVotingInit(registrationSMT_ common.Address, proposalsState_ common.Address, votingVerifier_ common.Address) (*types.Transaction, error) {
	return _NoirVoting.Contract.NoirIDVotingInit(&_NoirVoting.TransactOpts, registrationSMT_, proposalsState_, votingVerifier_)
}

// NoirIDVotingInit is a paid mutator transaction binding the contract method 0x6941af14.
//
// Solidity: function __NoirIDVoting_init(address registrationSMT_, address proposalsState_, address votingVerifier_) returns()
func (_NoirVoting *NoirVotingTransactorSession) NoirIDVotingInit(registrationSMT_ common.Address, proposalsState_ common.Address, votingVerifier_ common.Address) (*types.Transaction, error) {
	return _NoirVoting.Contract.NoirIDVotingInit(&_NoirVoting.TransactOpts, registrationSMT_, proposalsState_, votingVerifier_)
}

// AfterVerify is a paid mutator transaction binding the contract method 0x687fbf10.
//
// Solidity: function _afterVerify(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) returns()
func (_NoirVoting *NoirVotingTransactor) AfterVerify(opts *bind.TransactOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) (*types.Transaction, error) {
	return _NoirVoting.contract.Transact(opts, "_afterVerify", registrationRoot_, currentDate_, userPayload_)
}

// AfterVerify is a paid mutator transaction binding the contract method 0x687fbf10.
//
// Solidity: function _afterVerify(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) returns()
func (_NoirVoting *NoirVotingSession) AfterVerify(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) (*types.Transaction, error) {
	return _NoirVoting.Contract.AfterVerify(&_NoirVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_)
}

// AfterVerify is a paid mutator transaction binding the contract method 0x687fbf10.
//
// Solidity: function _afterVerify(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) returns()
func (_NoirVoting *NoirVotingTransactorSession) AfterVerify(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) (*types.Transaction, error) {
	return _NoirVoting.Contract.AfterVerify(&_NoirVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_)
}

// BeforeVerify is a paid mutator transaction binding the contract method 0xe73f3f11.
//
// Solidity: function _beforeVerify(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) returns()
func (_NoirVoting *NoirVotingTransactor) BeforeVerify(opts *bind.TransactOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) (*types.Transaction, error) {
	return _NoirVoting.contract.Transact(opts, "_beforeVerify", registrationRoot_, currentDate_, userPayload_)
}

// BeforeVerify is a paid mutator transaction binding the contract method 0xe73f3f11.
//
// Solidity: function _beforeVerify(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) returns()
func (_NoirVoting *NoirVotingSession) BeforeVerify(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) (*types.Transaction, error) {
	return _NoirVoting.Contract.BeforeVerify(&_NoirVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_)
}

// BeforeVerify is a paid mutator transaction binding the contract method 0xe73f3f11.
//
// Solidity: function _beforeVerify(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_) returns()
func (_NoirVoting *NoirVotingTransactorSession) BeforeVerify(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte) (*types.Transaction, error) {
	return _NoirVoting.Contract.BeforeVerify(&_NoirVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_)
}

// Execute is a paid mutator transaction binding the contract method 0xe4ab0833.
//
// Solidity: function execute(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_NoirVoting *NoirVotingTransactor) Execute(opts *bind.TransactOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ AQueryProofExecutorProofPoints) (*types.Transaction, error) {
	return _NoirVoting.contract.Transact(opts, "execute", registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// Execute is a paid mutator transaction binding the contract method 0xe4ab0833.
//
// Solidity: function execute(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_NoirVoting *NoirVotingSession) Execute(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ AQueryProofExecutorProofPoints) (*types.Transaction, error) {
	return _NoirVoting.Contract.Execute(&_NoirVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// Execute is a paid mutator transaction binding the contract method 0xe4ab0833.
//
// Solidity: function execute(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_NoirVoting *NoirVotingTransactorSession) Execute(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ AQueryProofExecutorProofPoints) (*types.Transaction, error) {
	return _NoirVoting.Contract.Execute(&_NoirVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteNoir is a paid mutator transaction binding the contract method 0x6effdd9f.
//
// Solidity: function executeNoir(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, bytes zkPoints_) returns()
func (_NoirVoting *NoirVotingTransactor) ExecuteNoir(opts *bind.TransactOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoirVoting.contract.Transact(opts, "executeNoir", registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteNoir is a paid mutator transaction binding the contract method 0x6effdd9f.
//
// Solidity: function executeNoir(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, bytes zkPoints_) returns()
func (_NoirVoting *NoirVotingSession) ExecuteNoir(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoirVoting.Contract.ExecuteNoir(&_NoirVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// ExecuteNoir is a paid mutator transaction binding the contract method 0x6effdd9f.
//
// Solidity: function executeNoir(bytes32 registrationRoot_, uint256 currentDate_, bytes userPayload_, bytes zkPoints_) returns()
func (_NoirVoting *NoirVotingTransactorSession) ExecuteNoir(registrationRoot_ [32]byte, currentDate_ *big.Int, userPayload_ []byte, zkPoints_ []byte) (*types.Transaction, error) {
	return _NoirVoting.Contract.ExecuteNoir(&_NoirVoting.TransactOpts, registrationRoot_, currentDate_, userPayload_, zkPoints_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NoirVoting *NoirVotingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoirVoting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NoirVoting *NoirVotingSession) RenounceOwnership() (*types.Transaction, error) {
	return _NoirVoting.Contract.RenounceOwnership(&_NoirVoting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NoirVoting *NoirVotingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NoirVoting.Contract.RenounceOwnership(&_NoirVoting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NoirVoting *NoirVotingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NoirVoting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NoirVoting *NoirVotingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NoirVoting.Contract.TransferOwnership(&_NoirVoting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NoirVoting *NoirVotingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NoirVoting.Contract.TransferOwnership(&_NoirVoting.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NoirVoting *NoirVotingTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _NoirVoting.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NoirVoting *NoirVotingSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NoirVoting.Contract.UpgradeTo(&_NoirVoting.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NoirVoting *NoirVotingTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NoirVoting.Contract.UpgradeTo(&_NoirVoting.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NoirVoting *NoirVotingTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NoirVoting.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NoirVoting *NoirVotingSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NoirVoting.Contract.UpgradeToAndCall(&_NoirVoting.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NoirVoting *NoirVotingTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NoirVoting.Contract.UpgradeToAndCall(&_NoirVoting.TransactOpts, newImplementation, data)
}

// NoirVotingAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the NoirVoting contract.
type NoirVotingAdminChangedIterator struct {
	Event *NoirVotingAdminChanged // Event containing the contract specifics and raw log

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
func (it *NoirVotingAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoirVotingAdminChanged)
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
		it.Event = new(NoirVotingAdminChanged)
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
func (it *NoirVotingAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoirVotingAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoirVotingAdminChanged represents a AdminChanged event raised by the NoirVoting contract.
type NoirVotingAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NoirVoting *NoirVotingFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NoirVotingAdminChangedIterator, error) {

	logs, sub, err := _NoirVoting.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NoirVotingAdminChangedIterator{contract: _NoirVoting.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NoirVoting *NoirVotingFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NoirVotingAdminChanged) (event.Subscription, error) {

	logs, sub, err := _NoirVoting.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoirVotingAdminChanged)
				if err := _NoirVoting.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_NoirVoting *NoirVotingFilterer) ParseAdminChanged(log types.Log) (*NoirVotingAdminChanged, error) {
	event := new(NoirVotingAdminChanged)
	if err := _NoirVoting.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoirVotingBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the NoirVoting contract.
type NoirVotingBeaconUpgradedIterator struct {
	Event *NoirVotingBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NoirVotingBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoirVotingBeaconUpgraded)
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
		it.Event = new(NoirVotingBeaconUpgraded)
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
func (it *NoirVotingBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoirVotingBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoirVotingBeaconUpgraded represents a BeaconUpgraded event raised by the NoirVoting contract.
type NoirVotingBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NoirVoting *NoirVotingFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NoirVotingBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NoirVoting.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NoirVotingBeaconUpgradedIterator{contract: _NoirVoting.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NoirVoting *NoirVotingFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NoirVotingBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NoirVoting.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoirVotingBeaconUpgraded)
				if err := _NoirVoting.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_NoirVoting *NoirVotingFilterer) ParseBeaconUpgraded(log types.Log) (*NoirVotingBeaconUpgraded, error) {
	event := new(NoirVotingBeaconUpgraded)
	if err := _NoirVoting.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoirVotingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NoirVoting contract.
type NoirVotingInitializedIterator struct {
	Event *NoirVotingInitialized // Event containing the contract specifics and raw log

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
func (it *NoirVotingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoirVotingInitialized)
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
		it.Event = new(NoirVotingInitialized)
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
func (it *NoirVotingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoirVotingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoirVotingInitialized represents a Initialized event raised by the NoirVoting contract.
type NoirVotingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NoirVoting *NoirVotingFilterer) FilterInitialized(opts *bind.FilterOpts) (*NoirVotingInitializedIterator, error) {

	logs, sub, err := _NoirVoting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NoirVotingInitializedIterator{contract: _NoirVoting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NoirVoting *NoirVotingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NoirVotingInitialized) (event.Subscription, error) {

	logs, sub, err := _NoirVoting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoirVotingInitialized)
				if err := _NoirVoting.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NoirVoting *NoirVotingFilterer) ParseInitialized(log types.Log) (*NoirVotingInitialized, error) {
	event := new(NoirVotingInitialized)
	if err := _NoirVoting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoirVotingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NoirVoting contract.
type NoirVotingOwnershipTransferredIterator struct {
	Event *NoirVotingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NoirVotingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoirVotingOwnershipTransferred)
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
		it.Event = new(NoirVotingOwnershipTransferred)
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
func (it *NoirVotingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoirVotingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoirVotingOwnershipTransferred represents a OwnershipTransferred event raised by the NoirVoting contract.
type NoirVotingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NoirVoting *NoirVotingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NoirVotingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NoirVoting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NoirVotingOwnershipTransferredIterator{contract: _NoirVoting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NoirVoting *NoirVotingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NoirVotingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NoirVoting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoirVotingOwnershipTransferred)
				if err := _NoirVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NoirVoting *NoirVotingFilterer) ParseOwnershipTransferred(log types.Log) (*NoirVotingOwnershipTransferred, error) {
	event := new(NoirVotingOwnershipTransferred)
	if err := _NoirVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NoirVotingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NoirVoting contract.
type NoirVotingUpgradedIterator struct {
	Event *NoirVotingUpgraded // Event containing the contract specifics and raw log

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
func (it *NoirVotingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoirVotingUpgraded)
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
		it.Event = new(NoirVotingUpgraded)
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
func (it *NoirVotingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoirVotingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoirVotingUpgraded represents a Upgraded event raised by the NoirVoting contract.
type NoirVotingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NoirVoting *NoirVotingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NoirVotingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NoirVoting.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NoirVotingUpgradedIterator{contract: _NoirVoting.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NoirVoting *NoirVotingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NoirVotingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NoirVoting.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoirVotingUpgraded)
				if err := _NoirVoting.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_NoirVoting *NoirVotingFilterer) ParseUpgraded(log types.Log) (*NoirVotingUpgraded, error) {
	event := new(NoirVotingUpgraded)
	if err := _NoirVoting.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
