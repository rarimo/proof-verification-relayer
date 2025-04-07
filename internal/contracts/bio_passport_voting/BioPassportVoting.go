// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package biopassportvoting

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

// VerifierHelperProofPoints is an auto generated low-level Go binding around an user-defined struct.
type VerifierHelperProofPoints struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// BioPassportVotingMetaData contains all meta data concerning the BioPassportVoting contract.
var BioPassportVotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pubSignals_\",\"type\":\"uint256[]\"}],\"name\":\"InvalidZKProof\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"IDENTITY_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROOF_SIGNALS_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_DATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registrationSMT_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposalsState_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"votingVerifier_\",\"type\":\"address\"}],\"name\":\"__BioPassportVoting_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalsState\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrationSMT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registrationRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"currentDate_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"vote_\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nullifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"citizenship\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identityCreationTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structBaseVoting.UserData\",\"name\":\"userData_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BioPassportVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use BioPassportVotingMetaData.ABI instead.
var BioPassportVotingABI = BioPassportVotingMetaData.ABI

// BioPassportVoting is an auto generated Go binding around an Ethereum contract.
type BioPassportVoting struct {
	BioPassportVotingCaller     // Read-only binding to the contract
	BioPassportVotingTransactor // Write-only binding to the contract
	BioPassportVotingFilterer   // Log filterer for contract events
}

// BioPassportVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BioPassportVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BioPassportVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BioPassportVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BioPassportVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BioPassportVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BioPassportVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BioPassportVotingSession struct {
	Contract     *BioPassportVoting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BioPassportVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BioPassportVotingCallerSession struct {
	Contract *BioPassportVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BioPassportVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BioPassportVotingTransactorSession struct {
	Contract     *BioPassportVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BioPassportVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BioPassportVotingRaw struct {
	Contract *BioPassportVoting // Generic contract binding to access the raw methods on
}

// BioPassportVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BioPassportVotingCallerRaw struct {
	Contract *BioPassportVotingCaller // Generic read-only contract binding to access the raw methods on
}

// BioPassportVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BioPassportVotingTransactorRaw struct {
	Contract *BioPassportVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBioPassportVoting creates a new instance of BioPassportVoting, bound to a specific deployed contract.
func NewBioPassportVoting(address common.Address, backend bind.ContractBackend) (*BioPassportVoting, error) {
	contract, err := bindBioPassportVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BioPassportVoting{BioPassportVotingCaller: BioPassportVotingCaller{contract: contract}, BioPassportVotingTransactor: BioPassportVotingTransactor{contract: contract}, BioPassportVotingFilterer: BioPassportVotingFilterer{contract: contract}}, nil
}

// NewBioPassportVotingCaller creates a new read-only instance of BioPassportVoting, bound to a specific deployed contract.
func NewBioPassportVotingCaller(address common.Address, caller bind.ContractCaller) (*BioPassportVotingCaller, error) {
	contract, err := bindBioPassportVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BioPassportVotingCaller{contract: contract}, nil
}

// NewBioPassportVotingTransactor creates a new write-only instance of BioPassportVoting, bound to a specific deployed contract.
func NewBioPassportVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*BioPassportVotingTransactor, error) {
	contract, err := bindBioPassportVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BioPassportVotingTransactor{contract: contract}, nil
}

// NewBioPassportVotingFilterer creates a new log filterer instance of BioPassportVoting, bound to a specific deployed contract.
func NewBioPassportVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*BioPassportVotingFilterer, error) {
	contract, err := bindBioPassportVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BioPassportVotingFilterer{contract: contract}, nil
}

// bindBioPassportVoting binds a generic wrapper to an already deployed contract.
func bindBioPassportVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BioPassportVotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BioPassportVoting *BioPassportVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BioPassportVoting.Contract.BioPassportVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BioPassportVoting *BioPassportVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.BioPassportVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BioPassportVoting *BioPassportVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.BioPassportVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BioPassportVoting *BioPassportVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BioPassportVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BioPassportVoting *BioPassportVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BioPassportVoting *BioPassportVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.contract.Transact(opts, method, params...)
}

// IDENTITYLIMIT is a free data retrieval call binding the contract method 0x7995c0f3.
//
// Solidity: function IDENTITY_LIMIT() view returns(uint256)
func (_BioPassportVoting *BioPassportVotingCaller) IDENTITYLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BioPassportVoting.contract.Call(opts, &out, "IDENTITY_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IDENTITYLIMIT is a free data retrieval call binding the contract method 0x7995c0f3.
//
// Solidity: function IDENTITY_LIMIT() view returns(uint256)
func (_BioPassportVoting *BioPassportVotingSession) IDENTITYLIMIT() (*big.Int, error) {
	return _BioPassportVoting.Contract.IDENTITYLIMIT(&_BioPassportVoting.CallOpts)
}

// IDENTITYLIMIT is a free data retrieval call binding the contract method 0x7995c0f3.
//
// Solidity: function IDENTITY_LIMIT() view returns(uint256)
func (_BioPassportVoting *BioPassportVotingCallerSession) IDENTITYLIMIT() (*big.Int, error) {
	return _BioPassportVoting.Contract.IDENTITYLIMIT(&_BioPassportVoting.CallOpts)
}

// PROOFSIGNALSCOUNT is a free data retrieval call binding the contract method 0xbdb4017f.
//
// Solidity: function PROOF_SIGNALS_COUNT() view returns(uint256)
func (_BioPassportVoting *BioPassportVotingCaller) PROOFSIGNALSCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BioPassportVoting.contract.Call(opts, &out, "PROOF_SIGNALS_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROOFSIGNALSCOUNT is a free data retrieval call binding the contract method 0xbdb4017f.
//
// Solidity: function PROOF_SIGNALS_COUNT() view returns(uint256)
func (_BioPassportVoting *BioPassportVotingSession) PROOFSIGNALSCOUNT() (*big.Int, error) {
	return _BioPassportVoting.Contract.PROOFSIGNALSCOUNT(&_BioPassportVoting.CallOpts)
}

// PROOFSIGNALSCOUNT is a free data retrieval call binding the contract method 0xbdb4017f.
//
// Solidity: function PROOF_SIGNALS_COUNT() view returns(uint256)
func (_BioPassportVoting *BioPassportVotingCallerSession) PROOFSIGNALSCOUNT() (*big.Int, error) {
	return _BioPassportVoting.Contract.PROOFSIGNALSCOUNT(&_BioPassportVoting.CallOpts)
}

// ZERODATE is a free data retrieval call binding the contract method 0x2c15be3c.
//
// Solidity: function ZERO_DATE() view returns(uint256)
func (_BioPassportVoting *BioPassportVotingCaller) ZERODATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BioPassportVoting.contract.Call(opts, &out, "ZERO_DATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZERODATE is a free data retrieval call binding the contract method 0x2c15be3c.
//
// Solidity: function ZERO_DATE() view returns(uint256)
func (_BioPassportVoting *BioPassportVotingSession) ZERODATE() (*big.Int, error) {
	return _BioPassportVoting.Contract.ZERODATE(&_BioPassportVoting.CallOpts)
}

// ZERODATE is a free data retrieval call binding the contract method 0x2c15be3c.
//
// Solidity: function ZERO_DATE() view returns(uint256)
func (_BioPassportVoting *BioPassportVotingCallerSession) ZERODATE() (*big.Int, error) {
	return _BioPassportVoting.Contract.ZERODATE(&_BioPassportVoting.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BioPassportVoting *BioPassportVotingCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BioPassportVoting.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BioPassportVoting *BioPassportVotingSession) Implementation() (common.Address, error) {
	return _BioPassportVoting.Contract.Implementation(&_BioPassportVoting.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_BioPassportVoting *BioPassportVotingCallerSession) Implementation() (common.Address, error) {
	return _BioPassportVoting.Contract.Implementation(&_BioPassportVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BioPassportVoting *BioPassportVotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BioPassportVoting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BioPassportVoting *BioPassportVotingSession) Owner() (common.Address, error) {
	return _BioPassportVoting.Contract.Owner(&_BioPassportVoting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BioPassportVoting *BioPassportVotingCallerSession) Owner() (common.Address, error) {
	return _BioPassportVoting.Contract.Owner(&_BioPassportVoting.CallOpts)
}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_BioPassportVoting *BioPassportVotingCaller) ProposalsState(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BioPassportVoting.contract.Call(opts, &out, "proposalsState")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_BioPassportVoting *BioPassportVotingSession) ProposalsState() (common.Address, error) {
	return _BioPassportVoting.Contract.ProposalsState(&_BioPassportVoting.CallOpts)
}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_BioPassportVoting *BioPassportVotingCallerSession) ProposalsState() (common.Address, error) {
	return _BioPassportVoting.Contract.ProposalsState(&_BioPassportVoting.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BioPassportVoting *BioPassportVotingCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BioPassportVoting.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BioPassportVoting *BioPassportVotingSession) ProxiableUUID() ([32]byte, error) {
	return _BioPassportVoting.Contract.ProxiableUUID(&_BioPassportVoting.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BioPassportVoting *BioPassportVotingCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BioPassportVoting.Contract.ProxiableUUID(&_BioPassportVoting.CallOpts)
}

// RegistrationSMT is a free data retrieval call binding the contract method 0x24a518a5.
//
// Solidity: function registrationSMT() view returns(address)
func (_BioPassportVoting *BioPassportVotingCaller) RegistrationSMT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BioPassportVoting.contract.Call(opts, &out, "registrationSMT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistrationSMT is a free data retrieval call binding the contract method 0x24a518a5.
//
// Solidity: function registrationSMT() view returns(address)
func (_BioPassportVoting *BioPassportVotingSession) RegistrationSMT() (common.Address, error) {
	return _BioPassportVoting.Contract.RegistrationSMT(&_BioPassportVoting.CallOpts)
}

// RegistrationSMT is a free data retrieval call binding the contract method 0x24a518a5.
//
// Solidity: function registrationSMT() view returns(address)
func (_BioPassportVoting *BioPassportVotingCallerSession) RegistrationSMT() (common.Address, error) {
	return _BioPassportVoting.Contract.RegistrationSMT(&_BioPassportVoting.CallOpts)
}

// VotingVerifier is a free data retrieval call binding the contract method 0x1f2c317f.
//
// Solidity: function votingVerifier() view returns(address)
func (_BioPassportVoting *BioPassportVotingCaller) VotingVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BioPassportVoting.contract.Call(opts, &out, "votingVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingVerifier is a free data retrieval call binding the contract method 0x1f2c317f.
//
// Solidity: function votingVerifier() view returns(address)
func (_BioPassportVoting *BioPassportVotingSession) VotingVerifier() (common.Address, error) {
	return _BioPassportVoting.Contract.VotingVerifier(&_BioPassportVoting.CallOpts)
}

// VotingVerifier is a free data retrieval call binding the contract method 0x1f2c317f.
//
// Solidity: function votingVerifier() view returns(address)
func (_BioPassportVoting *BioPassportVotingCallerSession) VotingVerifier() (common.Address, error) {
	return _BioPassportVoting.Contract.VotingVerifier(&_BioPassportVoting.CallOpts)
}

// BioPassportVotingInit is a paid mutator transaction binding the contract method 0xb6b8038c.
//
// Solidity: function __BioPassportVoting_init(address registrationSMT_, address proposalsState_, address votingVerifier_) returns()
func (_BioPassportVoting *BioPassportVotingTransactor) BioPassportVotingInit(opts *bind.TransactOpts, registrationSMT_ common.Address, proposalsState_ common.Address, votingVerifier_ common.Address) (*types.Transaction, error) {
	return _BioPassportVoting.contract.Transact(opts, "__BioPassportVoting_init", registrationSMT_, proposalsState_, votingVerifier_)
}

// BioPassportVotingInit is a paid mutator transaction binding the contract method 0xb6b8038c.
//
// Solidity: function __BioPassportVoting_init(address registrationSMT_, address proposalsState_, address votingVerifier_) returns()
func (_BioPassportVoting *BioPassportVotingSession) BioPassportVotingInit(registrationSMT_ common.Address, proposalsState_ common.Address, votingVerifier_ common.Address) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.BioPassportVotingInit(&_BioPassportVoting.TransactOpts, registrationSMT_, proposalsState_, votingVerifier_)
}

// BioPassportVotingInit is a paid mutator transaction binding the contract method 0xb6b8038c.
//
// Solidity: function __BioPassportVoting_init(address registrationSMT_, address proposalsState_, address votingVerifier_) returns()
func (_BioPassportVoting *BioPassportVotingTransactorSession) BioPassportVotingInit(registrationSMT_ common.Address, proposalsState_ common.Address, votingVerifier_ common.Address) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.BioPassportVotingInit(&_BioPassportVoting.TransactOpts, registrationSMT_, proposalsState_, votingVerifier_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BioPassportVoting *BioPassportVotingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BioPassportVoting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BioPassportVoting *BioPassportVotingSession) RenounceOwnership() (*types.Transaction, error) {
	return _BioPassportVoting.Contract.RenounceOwnership(&_BioPassportVoting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BioPassportVoting *BioPassportVotingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BioPassportVoting.Contract.RenounceOwnership(&_BioPassportVoting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BioPassportVoting *BioPassportVotingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BioPassportVoting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BioPassportVoting *BioPassportVotingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.TransferOwnership(&_BioPassportVoting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BioPassportVoting *BioPassportVotingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.TransferOwnership(&_BioPassportVoting.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BioPassportVoting *BioPassportVotingTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _BioPassportVoting.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BioPassportVoting *BioPassportVotingSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.UpgradeTo(&_BioPassportVoting.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BioPassportVoting *BioPassportVotingTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.UpgradeTo(&_BioPassportVoting.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BioPassportVoting *BioPassportVotingTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BioPassportVoting.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BioPassportVoting *BioPassportVotingSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.UpgradeToAndCall(&_BioPassportVoting.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BioPassportVoting *BioPassportVotingTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.UpgradeToAndCall(&_BioPassportVoting.TransactOpts, newImplementation, data)
}

// Vote is a paid mutator transaction binding the contract method 0x11181976.
//
// Solidity: function vote(bytes32 registrationRoot_, uint256 currentDate_, uint256 proposalId_, uint256[] vote_, (uint256,uint256,uint256) userData_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_BioPassportVoting *BioPassportVotingTransactor) Vote(opts *bind.TransactOpts, registrationRoot_ [32]byte, currentDate_ *big.Int, proposalId_ *big.Int, vote_ []*big.Int, userData_ BaseVotingUserData, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _BioPassportVoting.contract.Transact(opts, "vote", registrationRoot_, currentDate_, proposalId_, vote_, userData_, zkPoints_)
}

// Vote is a paid mutator transaction binding the contract method 0x11181976.
//
// Solidity: function vote(bytes32 registrationRoot_, uint256 currentDate_, uint256 proposalId_, uint256[] vote_, (uint256,uint256,uint256) userData_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_BioPassportVoting *BioPassportVotingSession) Vote(registrationRoot_ [32]byte, currentDate_ *big.Int, proposalId_ *big.Int, vote_ []*big.Int, userData_ BaseVotingUserData, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.Vote(&_BioPassportVoting.TransactOpts, registrationRoot_, currentDate_, proposalId_, vote_, userData_, zkPoints_)
}

// Vote is a paid mutator transaction binding the contract method 0x11181976.
//
// Solidity: function vote(bytes32 registrationRoot_, uint256 currentDate_, uint256 proposalId_, uint256[] vote_, (uint256,uint256,uint256) userData_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_BioPassportVoting *BioPassportVotingTransactorSession) Vote(registrationRoot_ [32]byte, currentDate_ *big.Int, proposalId_ *big.Int, vote_ []*big.Int, userData_ BaseVotingUserData, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _BioPassportVoting.Contract.Vote(&_BioPassportVoting.TransactOpts, registrationRoot_, currentDate_, proposalId_, vote_, userData_, zkPoints_)
}

// BioPassportVotingAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the BioPassportVoting contract.
type BioPassportVotingAdminChangedIterator struct {
	Event *BioPassportVotingAdminChanged // Event containing the contract specifics and raw log

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
func (it *BioPassportVotingAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BioPassportVotingAdminChanged)
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
		it.Event = new(BioPassportVotingAdminChanged)
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
func (it *BioPassportVotingAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BioPassportVotingAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BioPassportVotingAdminChanged represents a AdminChanged event raised by the BioPassportVoting contract.
type BioPassportVotingAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BioPassportVoting *BioPassportVotingFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BioPassportVotingAdminChangedIterator, error) {

	logs, sub, err := _BioPassportVoting.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BioPassportVotingAdminChangedIterator{contract: _BioPassportVoting.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BioPassportVoting *BioPassportVotingFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BioPassportVotingAdminChanged) (event.Subscription, error) {

	logs, sub, err := _BioPassportVoting.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BioPassportVotingAdminChanged)
				if err := _BioPassportVoting.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_BioPassportVoting *BioPassportVotingFilterer) ParseAdminChanged(log types.Log) (*BioPassportVotingAdminChanged, error) {
	event := new(BioPassportVotingAdminChanged)
	if err := _BioPassportVoting.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BioPassportVotingBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the BioPassportVoting contract.
type BioPassportVotingBeaconUpgradedIterator struct {
	Event *BioPassportVotingBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BioPassportVotingBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BioPassportVotingBeaconUpgraded)
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
		it.Event = new(BioPassportVotingBeaconUpgraded)
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
func (it *BioPassportVotingBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BioPassportVotingBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BioPassportVotingBeaconUpgraded represents a BeaconUpgraded event raised by the BioPassportVoting contract.
type BioPassportVotingBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BioPassportVoting *BioPassportVotingFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BioPassportVotingBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BioPassportVoting.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BioPassportVotingBeaconUpgradedIterator{contract: _BioPassportVoting.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BioPassportVoting *BioPassportVotingFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BioPassportVotingBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BioPassportVoting.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BioPassportVotingBeaconUpgraded)
				if err := _BioPassportVoting.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_BioPassportVoting *BioPassportVotingFilterer) ParseBeaconUpgraded(log types.Log) (*BioPassportVotingBeaconUpgraded, error) {
	event := new(BioPassportVotingBeaconUpgraded)
	if err := _BioPassportVoting.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BioPassportVotingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BioPassportVoting contract.
type BioPassportVotingInitializedIterator struct {
	Event *BioPassportVotingInitialized // Event containing the contract specifics and raw log

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
func (it *BioPassportVotingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BioPassportVotingInitialized)
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
		it.Event = new(BioPassportVotingInitialized)
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
func (it *BioPassportVotingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BioPassportVotingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BioPassportVotingInitialized represents a Initialized event raised by the BioPassportVoting contract.
type BioPassportVotingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BioPassportVoting *BioPassportVotingFilterer) FilterInitialized(opts *bind.FilterOpts) (*BioPassportVotingInitializedIterator, error) {

	logs, sub, err := _BioPassportVoting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BioPassportVotingInitializedIterator{contract: _BioPassportVoting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BioPassportVoting *BioPassportVotingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BioPassportVotingInitialized) (event.Subscription, error) {

	logs, sub, err := _BioPassportVoting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BioPassportVotingInitialized)
				if err := _BioPassportVoting.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BioPassportVoting *BioPassportVotingFilterer) ParseInitialized(log types.Log) (*BioPassportVotingInitialized, error) {
	event := new(BioPassportVotingInitialized)
	if err := _BioPassportVoting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BioPassportVotingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BioPassportVoting contract.
type BioPassportVotingOwnershipTransferredIterator struct {
	Event *BioPassportVotingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BioPassportVotingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BioPassportVotingOwnershipTransferred)
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
		it.Event = new(BioPassportVotingOwnershipTransferred)
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
func (it *BioPassportVotingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BioPassportVotingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BioPassportVotingOwnershipTransferred represents a OwnershipTransferred event raised by the BioPassportVoting contract.
type BioPassportVotingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BioPassportVoting *BioPassportVotingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BioPassportVotingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BioPassportVoting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BioPassportVotingOwnershipTransferredIterator{contract: _BioPassportVoting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BioPassportVoting *BioPassportVotingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BioPassportVotingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BioPassportVoting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BioPassportVotingOwnershipTransferred)
				if err := _BioPassportVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BioPassportVoting *BioPassportVotingFilterer) ParseOwnershipTransferred(log types.Log) (*BioPassportVotingOwnershipTransferred, error) {
	event := new(BioPassportVotingOwnershipTransferred)
	if err := _BioPassportVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BioPassportVotingUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BioPassportVoting contract.
type BioPassportVotingUpgradedIterator struct {
	Event *BioPassportVotingUpgraded // Event containing the contract specifics and raw log

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
func (it *BioPassportVotingUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BioPassportVotingUpgraded)
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
		it.Event = new(BioPassportVotingUpgraded)
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
func (it *BioPassportVotingUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BioPassportVotingUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BioPassportVotingUpgraded represents a Upgraded event raised by the BioPassportVoting contract.
type BioPassportVotingUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BioPassportVoting *BioPassportVotingFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BioPassportVotingUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BioPassportVoting.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BioPassportVotingUpgradedIterator{contract: _BioPassportVoting.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BioPassportVoting *BioPassportVotingFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BioPassportVotingUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BioPassportVoting.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BioPassportVotingUpgraded)
				if err := _BioPassportVoting.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BioPassportVoting *BioPassportVotingFilterer) ParseUpgraded(log types.Log) (*BioPassportVotingUpgraded, error) {
	event := new(BioPassportVotingUpgraded)
	if err := _BioPassportVoting.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
