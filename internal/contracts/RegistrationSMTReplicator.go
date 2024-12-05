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

// RegistrationSMTReplicatorMetaData contains all meta data concerning the RegistrationSMTReplicator contract.
var RegistrationSMTReplicatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transitionTimestamp\",\"type\":\"uint256\"}],\"name\":\"RootTransitioned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAGIC_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTRATION_ROOT_PREFIX\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_VALIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceSMT_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"}],\"name\":\"__RegistrationSMTReplicator_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newSignerPubKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"isRootLatest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"isRootValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"simpleSigner_\",\"type\":\"address\"}],\"name\":\"setSimpleSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"simpleSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceSMT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"transitionTimestamp_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"transitionRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"transitionTimestamp_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"transitionRootSimple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCallWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistrationSMTReplicatorABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrationSMTReplicatorMetaData.ABI instead.
var RegistrationSMTReplicatorABI = RegistrationSMTReplicatorMetaData.ABI

// RegistrationSMTReplicator is an auto generated Go binding around an Ethereum contract.
type RegistrationSMTReplicator struct {
	RegistrationSMTReplicatorCaller     // Read-only binding to the contract
	RegistrationSMTReplicatorTransactor // Write-only binding to the contract
	RegistrationSMTReplicatorFilterer   // Log filterer for contract events
}

// RegistrationSMTReplicatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrationSMTReplicatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationSMTReplicatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrationSMTReplicatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationSMTReplicatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrationSMTReplicatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationSMTReplicatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrationSMTReplicatorSession struct {
	Contract     *RegistrationSMTReplicator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RegistrationSMTReplicatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrationSMTReplicatorCallerSession struct {
	Contract *RegistrationSMTReplicatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// RegistrationSMTReplicatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrationSMTReplicatorTransactorSession struct {
	Contract     *RegistrationSMTReplicatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// RegistrationSMTReplicatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrationSMTReplicatorRaw struct {
	Contract *RegistrationSMTReplicator // Generic contract binding to access the raw methods on
}

// RegistrationSMTReplicatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrationSMTReplicatorCallerRaw struct {
	Contract *RegistrationSMTReplicatorCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrationSMTReplicatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrationSMTReplicatorTransactorRaw struct {
	Contract *RegistrationSMTReplicatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrationSMTReplicator creates a new instance of RegistrationSMTReplicator, bound to a specific deployed contract.
func NewRegistrationSMTReplicator(address common.Address, backend bind.ContractBackend) (*RegistrationSMTReplicator, error) {
	contract, err := bindRegistrationSMTReplicator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistrationSMTReplicator{RegistrationSMTReplicatorCaller: RegistrationSMTReplicatorCaller{contract: contract}, RegistrationSMTReplicatorTransactor: RegistrationSMTReplicatorTransactor{contract: contract}, RegistrationSMTReplicatorFilterer: RegistrationSMTReplicatorFilterer{contract: contract}}, nil
}

// NewRegistrationSMTReplicatorCaller creates a new read-only instance of RegistrationSMTReplicator, bound to a specific deployed contract.
func NewRegistrationSMTReplicatorCaller(address common.Address, caller bind.ContractCaller) (*RegistrationSMTReplicatorCaller, error) {
	contract, err := bindRegistrationSMTReplicator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationSMTReplicatorCaller{contract: contract}, nil
}

// NewRegistrationSMTReplicatorTransactor creates a new write-only instance of RegistrationSMTReplicator, bound to a specific deployed contract.
func NewRegistrationSMTReplicatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrationSMTReplicatorTransactor, error) {
	contract, err := bindRegistrationSMTReplicator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationSMTReplicatorTransactor{contract: contract}, nil
}

// NewRegistrationSMTReplicatorFilterer creates a new log filterer instance of RegistrationSMTReplicator, bound to a specific deployed contract.
func NewRegistrationSMTReplicatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrationSMTReplicatorFilterer, error) {
	contract, err := bindRegistrationSMTReplicator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrationSMTReplicatorFilterer{contract: contract}, nil
}

// bindRegistrationSMTReplicator binds a generic wrapper to an already deployed contract.
func bindRegistrationSMTReplicator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistrationSMTReplicatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistrationSMTReplicator.Contract.RegistrationSMTReplicatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.RegistrationSMTReplicatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.RegistrationSMTReplicatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistrationSMTReplicator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.contract.Transact(opts, method, params...)
}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) MAGICID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "MAGIC_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) MAGICID() (uint8, error) {
	return _RegistrationSMTReplicator.Contract.MAGICID(&_RegistrationSMTReplicator.CallOpts)
}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) MAGICID() (uint8, error) {
	return _RegistrationSMTReplicator.Contract.MAGICID(&_RegistrationSMTReplicator.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "P")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) P() (*big.Int, error) {
	return _RegistrationSMTReplicator.Contract.P(&_RegistrationSMTReplicator.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) P() (*big.Int, error) {
	return _RegistrationSMTReplicator.Contract.P(&_RegistrationSMTReplicator.CallOpts)
}

// REGISTRATIONROOTPREFIX is a free data retrieval call binding the contract method 0x790bdd4c.
//
// Solidity: function REGISTRATION_ROOT_PREFIX() view returns(string)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) REGISTRATIONROOTPREFIX(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "REGISTRATION_ROOT_PREFIX")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REGISTRATIONROOTPREFIX is a free data retrieval call binding the contract method 0x790bdd4c.
//
// Solidity: function REGISTRATION_ROOT_PREFIX() view returns(string)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) REGISTRATIONROOTPREFIX() (string, error) {
	return _RegistrationSMTReplicator.Contract.REGISTRATIONROOTPREFIX(&_RegistrationSMTReplicator.CallOpts)
}

// REGISTRATIONROOTPREFIX is a free data retrieval call binding the contract method 0x790bdd4c.
//
// Solidity: function REGISTRATION_ROOT_PREFIX() view returns(string)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) REGISTRATIONROOTPREFIX() (string, error) {
	return _RegistrationSMTReplicator.Contract.REGISTRATIONROOTPREFIX(&_RegistrationSMTReplicator.CallOpts)
}

// ROOTVALIDITY is a free data retrieval call binding the contract method 0xcffe9676.
//
// Solidity: function ROOT_VALIDITY() view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) ROOTVALIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "ROOT_VALIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROOTVALIDITY is a free data retrieval call binding the contract method 0xcffe9676.
//
// Solidity: function ROOT_VALIDITY() view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) ROOTVALIDITY() (*big.Int, error) {
	return _RegistrationSMTReplicator.Contract.ROOTVALIDITY(&_RegistrationSMTReplicator.CallOpts)
}

// ROOTVALIDITY is a free data retrieval call binding the contract method 0xcffe9676.
//
// Solidity: function ROOT_VALIDITY() view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) ROOTVALIDITY() (*big.Int, error) {
	return _RegistrationSMTReplicator.Contract.ROOTVALIDITY(&_RegistrationSMTReplicator.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) ChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "chainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) ChainName() (string, error) {
	return _RegistrationSMTReplicator.Contract.ChainName(&_RegistrationSMTReplicator.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) ChainName() (string, error) {
	return _RegistrationSMTReplicator.Contract.ChainName(&_RegistrationSMTReplicator.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) GetNonce(opts *bind.CallOpts, methodId_ uint8) (*big.Int, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "getNonce", methodId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) GetNonce(methodId_ uint8) (*big.Int, error) {
	return _RegistrationSMTReplicator.Contract.GetNonce(&_RegistrationSMTReplicator.CallOpts, methodId_)
}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) GetNonce(methodId_ uint8) (*big.Int, error) {
	return _RegistrationSMTReplicator.Contract.GetNonce(&_RegistrationSMTReplicator.CallOpts, methodId_)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) Implementation() (common.Address, error) {
	return _RegistrationSMTReplicator.Contract.Implementation(&_RegistrationSMTReplicator.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) Implementation() (common.Address, error) {
	return _RegistrationSMTReplicator.Contract.Implementation(&_RegistrationSMTReplicator.CallOpts)
}

// IsRootLatest is a free data retrieval call binding the contract method 0x8492307f.
//
// Solidity: function isRootLatest(bytes32 root_) view returns(bool)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) IsRootLatest(opts *bind.CallOpts, root_ [32]byte) (bool, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "isRootLatest", root_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRootLatest is a free data retrieval call binding the contract method 0x8492307f.
//
// Solidity: function isRootLatest(bytes32 root_) view returns(bool)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) IsRootLatest(root_ [32]byte) (bool, error) {
	return _RegistrationSMTReplicator.Contract.IsRootLatest(&_RegistrationSMTReplicator.CallOpts, root_)
}

// IsRootLatest is a free data retrieval call binding the contract method 0x8492307f.
//
// Solidity: function isRootLatest(bytes32 root_) view returns(bool)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) IsRootLatest(root_ [32]byte) (bool, error) {
	return _RegistrationSMTReplicator.Contract.IsRootLatest(&_RegistrationSMTReplicator.CallOpts, root_)
}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 root_) view returns(bool)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) IsRootValid(opts *bind.CallOpts, root_ [32]byte) (bool, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "isRootValid", root_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 root_) view returns(bool)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) IsRootValid(root_ [32]byte) (bool, error) {
	return _RegistrationSMTReplicator.Contract.IsRootValid(&_RegistrationSMTReplicator.CallOpts, root_)
}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 root_) view returns(bool)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) IsRootValid(root_ [32]byte) (bool, error) {
	return _RegistrationSMTReplicator.Contract.IsRootValid(&_RegistrationSMTReplicator.CallOpts, root_)
}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(bytes32)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) LatestRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "latestRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(bytes32)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) LatestRoot() ([32]byte, error) {
	return _RegistrationSMTReplicator.Contract.LatestRoot(&_RegistrationSMTReplicator.CallOpts)
}

// LatestRoot is a free data retrieval call binding the contract method 0xd7b0fef1.
//
// Solidity: function latestRoot() view returns(bytes32)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) LatestRoot() ([32]byte, error) {
	return _RegistrationSMTReplicator.Contract.LatestRoot(&_RegistrationSMTReplicator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) LatestTimestamp() (*big.Int, error) {
	return _RegistrationSMTReplicator.Contract.LatestTimestamp(&_RegistrationSMTReplicator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _RegistrationSMTReplicator.Contract.LatestTimestamp(&_RegistrationSMTReplicator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) Owner() (common.Address, error) {
	return _RegistrationSMTReplicator.Contract.Owner(&_RegistrationSMTReplicator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) Owner() (common.Address, error) {
	return _RegistrationSMTReplicator.Contract.Owner(&_RegistrationSMTReplicator.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) ProxiableUUID() ([32]byte, error) {
	return _RegistrationSMTReplicator.Contract.ProxiableUUID(&_RegistrationSMTReplicator.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) ProxiableUUID() ([32]byte, error) {
	return _RegistrationSMTReplicator.Contract.ProxiableUUID(&_RegistrationSMTReplicator.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) Signer() (common.Address, error) {
	return _RegistrationSMTReplicator.Contract.Signer(&_RegistrationSMTReplicator.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) Signer() (common.Address, error) {
	return _RegistrationSMTReplicator.Contract.Signer(&_RegistrationSMTReplicator.CallOpts)
}

// SimpleSigner is a free data retrieval call binding the contract method 0x1d627425.
//
// Solidity: function simpleSigner() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) SimpleSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "simpleSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SimpleSigner is a free data retrieval call binding the contract method 0x1d627425.
//
// Solidity: function simpleSigner() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) SimpleSigner() (common.Address, error) {
	return _RegistrationSMTReplicator.Contract.SimpleSigner(&_RegistrationSMTReplicator.CallOpts)
}

// SimpleSigner is a free data retrieval call binding the contract method 0x1d627425.
//
// Solidity: function simpleSigner() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) SimpleSigner() (common.Address, error) {
	return _RegistrationSMTReplicator.Contract.SimpleSigner(&_RegistrationSMTReplicator.CallOpts)
}

// SourceSMT is a free data retrieval call binding the contract method 0xc37d7ed0.
//
// Solidity: function sourceSMT() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCaller) SourceSMT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrationSMTReplicator.contract.Call(opts, &out, "sourceSMT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourceSMT is a free data retrieval call binding the contract method 0xc37d7ed0.
//
// Solidity: function sourceSMT() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) SourceSMT() (common.Address, error) {
	return _RegistrationSMTReplicator.Contract.SourceSMT(&_RegistrationSMTReplicator.CallOpts)
}

// SourceSMT is a free data retrieval call binding the contract method 0xc37d7ed0.
//
// Solidity: function sourceSMT() view returns(address)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorCallerSession) SourceSMT() (common.Address, error) {
	return _RegistrationSMTReplicator.Contract.SourceSMT(&_RegistrationSMTReplicator.CallOpts)
}

// RegistrationSMTReplicatorInit is a paid mutator transaction binding the contract method 0x5eb50721.
//
// Solidity: function __RegistrationSMTReplicator_init(address signer_, address sourceSMT_, string chainName_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactor) RegistrationSMTReplicatorInit(opts *bind.TransactOpts, signer_ common.Address, sourceSMT_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.contract.Transact(opts, "__RegistrationSMTReplicator_init", signer_, sourceSMT_, chainName_)
}

// RegistrationSMTReplicatorInit is a paid mutator transaction binding the contract method 0x5eb50721.
//
// Solidity: function __RegistrationSMTReplicator_init(address signer_, address sourceSMT_, string chainName_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) RegistrationSMTReplicatorInit(signer_ common.Address, sourceSMT_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.RegistrationSMTReplicatorInit(&_RegistrationSMTReplicator.TransactOpts, signer_, sourceSMT_, chainName_)
}

// RegistrationSMTReplicatorInit is a paid mutator transaction binding the contract method 0x5eb50721.
//
// Solidity: function __RegistrationSMTReplicator_init(address signer_, address sourceSMT_, string chainName_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorSession) RegistrationSMTReplicatorInit(signer_ common.Address, sourceSMT_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.RegistrationSMTReplicatorInit(&_RegistrationSMTReplicator.TransactOpts, signer_, sourceSMT_, chainName_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactor) ChangeSigner(opts *bind.TransactOpts, newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.contract.Transact(opts, "changeSigner", newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.ChangeSigner(&_RegistrationSMTReplicator.TransactOpts, newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.ChangeSigner(&_RegistrationSMTReplicator.TransactOpts, newSignerPubKey_, signature_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.RenounceOwnership(&_RegistrationSMTReplicator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.RenounceOwnership(&_RegistrationSMTReplicator.TransactOpts)
}

// SetSimpleSigner is a paid mutator transaction binding the contract method 0x05e696c1.
//
// Solidity: function setSimpleSigner(address simpleSigner_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactor) SetSimpleSigner(opts *bind.TransactOpts, simpleSigner_ common.Address) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.contract.Transact(opts, "setSimpleSigner", simpleSigner_)
}

// SetSimpleSigner is a paid mutator transaction binding the contract method 0x05e696c1.
//
// Solidity: function setSimpleSigner(address simpleSigner_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) SetSimpleSigner(simpleSigner_ common.Address) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.SetSimpleSigner(&_RegistrationSMTReplicator.TransactOpts, simpleSigner_)
}

// SetSimpleSigner is a paid mutator transaction binding the contract method 0x05e696c1.
//
// Solidity: function setSimpleSigner(address simpleSigner_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorSession) SetSimpleSigner(simpleSigner_ common.Address) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.SetSimpleSigner(&_RegistrationSMTReplicator.TransactOpts, simpleSigner_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.TransferOwnership(&_RegistrationSMTReplicator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.TransferOwnership(&_RegistrationSMTReplicator.TransactOpts, newOwner)
}

// TransitionRoot is a paid mutator transaction binding the contract method 0x211bf46a.
//
// Solidity: function transitionRoot(bytes32 newRoot_, uint256 transitionTimestamp_, bytes proof_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactor) TransitionRoot(opts *bind.TransactOpts, newRoot_ [32]byte, transitionTimestamp_ *big.Int, proof_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.contract.Transact(opts, "transitionRoot", newRoot_, transitionTimestamp_, proof_)
}

// TransitionRoot is a paid mutator transaction binding the contract method 0x211bf46a.
//
// Solidity: function transitionRoot(bytes32 newRoot_, uint256 transitionTimestamp_, bytes proof_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) TransitionRoot(newRoot_ [32]byte, transitionTimestamp_ *big.Int, proof_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.TransitionRoot(&_RegistrationSMTReplicator.TransactOpts, newRoot_, transitionTimestamp_, proof_)
}

// TransitionRoot is a paid mutator transaction binding the contract method 0x211bf46a.
//
// Solidity: function transitionRoot(bytes32 newRoot_, uint256 transitionTimestamp_, bytes proof_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorSession) TransitionRoot(newRoot_ [32]byte, transitionTimestamp_ *big.Int, proof_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.TransitionRoot(&_RegistrationSMTReplicator.TransactOpts, newRoot_, transitionTimestamp_, proof_)
}

// TransitionRootSimple is a paid mutator transaction binding the contract method 0xbde4432c.
//
// Solidity: function transitionRootSimple(bytes32 newRoot_, uint256 transitionTimestamp_, bytes signature_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactor) TransitionRootSimple(opts *bind.TransactOpts, newRoot_ [32]byte, transitionTimestamp_ *big.Int, signature_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.contract.Transact(opts, "transitionRootSimple", newRoot_, transitionTimestamp_, signature_)
}

// TransitionRootSimple is a paid mutator transaction binding the contract method 0xbde4432c.
//
// Solidity: function transitionRootSimple(bytes32 newRoot_, uint256 transitionTimestamp_, bytes signature_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) TransitionRootSimple(newRoot_ [32]byte, transitionTimestamp_ *big.Int, signature_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.TransitionRootSimple(&_RegistrationSMTReplicator.TransactOpts, newRoot_, transitionTimestamp_, signature_)
}

// TransitionRootSimple is a paid mutator transaction binding the contract method 0xbde4432c.
//
// Solidity: function transitionRootSimple(bytes32 newRoot_, uint256 transitionTimestamp_, bytes signature_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorSession) TransitionRootSimple(newRoot_ [32]byte, transitionTimestamp_ *big.Int, signature_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.TransitionRootSimple(&_RegistrationSMTReplicator.TransactOpts, newRoot_, transitionTimestamp_, signature_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.UpgradeTo(&_RegistrationSMTReplicator.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.UpgradeTo(&_RegistrationSMTReplicator.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.UpgradeToAndCall(&_RegistrationSMTReplicator.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.UpgradeToAndCall(&_RegistrationSMTReplicator.TransactOpts, newImplementation, data)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactor) UpgradeToAndCallWithProof(opts *bind.TransactOpts, newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.contract.Transact(opts, "upgradeToAndCallWithProof", newImplementation_, proof_, data_)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) UpgradeToAndCallWithProof(newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.UpgradeToAndCallWithProof(&_RegistrationSMTReplicator.TransactOpts, newImplementation_, proof_, data_)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorSession) UpgradeToAndCallWithProof(newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.UpgradeToAndCallWithProof(&_RegistrationSMTReplicator.TransactOpts, newImplementation_, proof_, data_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactor) UpgradeToWithProof(opts *bind.TransactOpts, newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.contract.Transact(opts, "upgradeToWithProof", newImplementation_, proof_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorSession) UpgradeToWithProof(newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.UpgradeToWithProof(&_RegistrationSMTReplicator.TransactOpts, newImplementation_, proof_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorTransactorSession) UpgradeToWithProof(newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _RegistrationSMTReplicator.Contract.UpgradeToWithProof(&_RegistrationSMTReplicator.TransactOpts, newImplementation_, proof_)
}

// RegistrationSMTReplicatorAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorAdminChangedIterator struct {
	Event *RegistrationSMTReplicatorAdminChanged // Event containing the contract specifics and raw log

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
func (it *RegistrationSMTReplicatorAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationSMTReplicatorAdminChanged)
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
		it.Event = new(RegistrationSMTReplicatorAdminChanged)
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
func (it *RegistrationSMTReplicatorAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationSMTReplicatorAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationSMTReplicatorAdminChanged represents a AdminChanged event raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RegistrationSMTReplicatorAdminChangedIterator, error) {

	logs, sub, err := _RegistrationSMTReplicator.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RegistrationSMTReplicatorAdminChangedIterator{contract: _RegistrationSMTReplicator.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RegistrationSMTReplicatorAdminChanged) (event.Subscription, error) {

	logs, sub, err := _RegistrationSMTReplicator.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationSMTReplicatorAdminChanged)
				if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) ParseAdminChanged(log types.Log) (*RegistrationSMTReplicatorAdminChanged, error) {
	event := new(RegistrationSMTReplicatorAdminChanged)
	if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationSMTReplicatorBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorBeaconUpgradedIterator struct {
	Event *RegistrationSMTReplicatorBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *RegistrationSMTReplicatorBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationSMTReplicatorBeaconUpgraded)
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
		it.Event = new(RegistrationSMTReplicatorBeaconUpgraded)
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
func (it *RegistrationSMTReplicatorBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationSMTReplicatorBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationSMTReplicatorBeaconUpgraded represents a BeaconUpgraded event raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*RegistrationSMTReplicatorBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RegistrationSMTReplicator.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationSMTReplicatorBeaconUpgradedIterator{contract: _RegistrationSMTReplicator.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *RegistrationSMTReplicatorBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RegistrationSMTReplicator.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationSMTReplicatorBeaconUpgraded)
				if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) ParseBeaconUpgraded(log types.Log) (*RegistrationSMTReplicatorBeaconUpgraded, error) {
	event := new(RegistrationSMTReplicatorBeaconUpgraded)
	if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationSMTReplicatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorInitializedIterator struct {
	Event *RegistrationSMTReplicatorInitialized // Event containing the contract specifics and raw log

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
func (it *RegistrationSMTReplicatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationSMTReplicatorInitialized)
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
		it.Event = new(RegistrationSMTReplicatorInitialized)
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
func (it *RegistrationSMTReplicatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationSMTReplicatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationSMTReplicatorInitialized represents a Initialized event raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistrationSMTReplicatorInitializedIterator, error) {

	logs, sub, err := _RegistrationSMTReplicator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistrationSMTReplicatorInitializedIterator{contract: _RegistrationSMTReplicator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistrationSMTReplicatorInitialized) (event.Subscription, error) {

	logs, sub, err := _RegistrationSMTReplicator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationSMTReplicatorInitialized)
				if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) ParseInitialized(log types.Log) (*RegistrationSMTReplicatorInitialized, error) {
	event := new(RegistrationSMTReplicatorInitialized)
	if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationSMTReplicatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorOwnershipTransferredIterator struct {
	Event *RegistrationSMTReplicatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrationSMTReplicatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationSMTReplicatorOwnershipTransferred)
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
		it.Event = new(RegistrationSMTReplicatorOwnershipTransferred)
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
func (it *RegistrationSMTReplicatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationSMTReplicatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationSMTReplicatorOwnershipTransferred represents a OwnershipTransferred event raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistrationSMTReplicatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegistrationSMTReplicator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationSMTReplicatorOwnershipTransferredIterator{contract: _RegistrationSMTReplicator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistrationSMTReplicatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegistrationSMTReplicator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationSMTReplicatorOwnershipTransferred)
				if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) ParseOwnershipTransferred(log types.Log) (*RegistrationSMTReplicatorOwnershipTransferred, error) {
	event := new(RegistrationSMTReplicatorOwnershipTransferred)
	if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationSMTReplicatorRootTransitionedIterator is returned from FilterRootTransitioned and is used to iterate over the raw logs and unpacked data for RootTransitioned events raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorRootTransitionedIterator struct {
	Event *RegistrationSMTReplicatorRootTransitioned // Event containing the contract specifics and raw log

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
func (it *RegistrationSMTReplicatorRootTransitionedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationSMTReplicatorRootTransitioned)
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
		it.Event = new(RegistrationSMTReplicatorRootTransitioned)
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
func (it *RegistrationSMTReplicatorRootTransitionedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationSMTReplicatorRootTransitionedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationSMTReplicatorRootTransitioned represents a RootTransitioned event raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorRootTransitioned struct {
	NewRoot             [32]byte
	TransitionTimestamp *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRootTransitioned is a free log retrieval operation binding the contract event 0x287d7075e3fdd1ee3cb7eef1d33839a4b50939e7bc33a68d8f6031eb3a1a14c6.
//
// Solidity: event RootTransitioned(bytes32 newRoot, uint256 transitionTimestamp)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) FilterRootTransitioned(opts *bind.FilterOpts) (*RegistrationSMTReplicatorRootTransitionedIterator, error) {

	logs, sub, err := _RegistrationSMTReplicator.contract.FilterLogs(opts, "RootTransitioned")
	if err != nil {
		return nil, err
	}
	return &RegistrationSMTReplicatorRootTransitionedIterator{contract: _RegistrationSMTReplicator.contract, event: "RootTransitioned", logs: logs, sub: sub}, nil
}

// WatchRootTransitioned is a free log subscription operation binding the contract event 0x287d7075e3fdd1ee3cb7eef1d33839a4b50939e7bc33a68d8f6031eb3a1a14c6.
//
// Solidity: event RootTransitioned(bytes32 newRoot, uint256 transitionTimestamp)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) WatchRootTransitioned(opts *bind.WatchOpts, sink chan<- *RegistrationSMTReplicatorRootTransitioned) (event.Subscription, error) {

	logs, sub, err := _RegistrationSMTReplicator.contract.WatchLogs(opts, "RootTransitioned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationSMTReplicatorRootTransitioned)
				if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "RootTransitioned", log); err != nil {
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

// ParseRootTransitioned is a log parse operation binding the contract event 0x287d7075e3fdd1ee3cb7eef1d33839a4b50939e7bc33a68d8f6031eb3a1a14c6.
//
// Solidity: event RootTransitioned(bytes32 newRoot, uint256 transitionTimestamp)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) ParseRootTransitioned(log types.Log) (*RegistrationSMTReplicatorRootTransitioned, error) {
	event := new(RegistrationSMTReplicatorRootTransitioned)
	if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "RootTransitioned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationSMTReplicatorUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorUpgradedIterator struct {
	Event *RegistrationSMTReplicatorUpgraded // Event containing the contract specifics and raw log

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
func (it *RegistrationSMTReplicatorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationSMTReplicatorUpgraded)
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
		it.Event = new(RegistrationSMTReplicatorUpgraded)
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
func (it *RegistrationSMTReplicatorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationSMTReplicatorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationSMTReplicatorUpgraded represents a Upgraded event raised by the RegistrationSMTReplicator contract.
type RegistrationSMTReplicatorUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RegistrationSMTReplicatorUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RegistrationSMTReplicator.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationSMTReplicatorUpgradedIterator{contract: _RegistrationSMTReplicator.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RegistrationSMTReplicatorUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RegistrationSMTReplicator.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationSMTReplicatorUpgraded)
				if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_RegistrationSMTReplicator *RegistrationSMTReplicatorFilterer) ParseUpgraded(log types.Log) (*RegistrationSMTReplicatorUpgraded, error) {
	event := new(RegistrationSMTReplicatorUpgraded)
	if err := _RegistrationSMTReplicator.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
