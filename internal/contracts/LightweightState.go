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

// ILightweightStateIdentitiesStatesRootData is an auto generated low-level Go binding around an user-defined struct.
type ILightweightStateIdentitiesStatesRootData struct {
	Root         [32]byte
	SetTimestamp *big.Int
}

// LightweightStateMetaData contains all meta data concerning the LightweightState contract.
var LightweightStateMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGistRoot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newIdentitesStatesRoot\",\"type\":\"bytes32\"}],\"name\":\"SignedStateTransited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceStateContract_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sourceChainName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"}],\"name\":\"__LightweightState_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"}],\"name\":\"__Signers_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newSignerPubKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSourceStateContract_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSourceStateContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"checkSignatureAndIncrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root_\",\"type\":\"uint256\"}],\"name\":\"geGISTRootData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structILightweightState.GistRootData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentGISTRootInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structILightweightState.GistRootData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGISTRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"getIdentitiesStatesRootData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"setTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structILightweightState.IdentitiesStatesRootData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"getSigComponents\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"identitiesStatesRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"isIdentitiesStatesRootExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newIdentitiesStatesRoot_\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structILightweightState.GistRootData\",\"name\":\"gistData_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"signedTransitState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceChainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceStateContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"validateChangeAddressSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"issuerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"issuerState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structILightweightState.StatesMerkleData\",\"name\":\"statesMerkleData_\",\"type\":\"tuple\"}],\"name\":\"verifyStatesMerkleData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LightweightStateABI is the input ABI used to generate the binding from.
// Deprecated: Use LightweightStateMetaData.ABI instead.
var LightweightStateABI = LightweightStateMetaData.ABI

// LightweightState is an auto generated Go binding around an Ethereum contract.
type LightweightState struct {
	LightweightStateCaller     // Read-only binding to the contract
	LightweightStateTransactor // Write-only binding to the contract
	LightweightStateFilterer   // Log filterer for contract events
}

// LightweightStateCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightweightStateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightweightStateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightweightStateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightweightStateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightweightStateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightweightStateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightweightStateSession struct {
	Contract     *LightweightState // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LightweightStateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightweightStateCallerSession struct {
	Contract *LightweightStateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LightweightStateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightweightStateTransactorSession struct {
	Contract     *LightweightStateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LightweightStateRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightweightStateRaw struct {
	Contract *LightweightState // Generic contract binding to access the raw methods on
}

// LightweightStateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightweightStateCallerRaw struct {
	Contract *LightweightStateCaller // Generic read-only contract binding to access the raw methods on
}

// LightweightStateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightweightStateTransactorRaw struct {
	Contract *LightweightStateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightweightState creates a new instance of LightweightState, bound to a specific deployed contract.
func NewLightweightState(address common.Address, backend bind.ContractBackend) (*LightweightState, error) {
	contract, err := bindLightweightState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LightweightState{LightweightStateCaller: LightweightStateCaller{contract: contract}, LightweightStateTransactor: LightweightStateTransactor{contract: contract}, LightweightStateFilterer: LightweightStateFilterer{contract: contract}}, nil
}

// NewLightweightStateCaller creates a new read-only instance of LightweightState, bound to a specific deployed contract.
func NewLightweightStateCaller(address common.Address, caller bind.ContractCaller) (*LightweightStateCaller, error) {
	contract, err := bindLightweightState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightweightStateCaller{contract: contract}, nil
}

// NewLightweightStateTransactor creates a new write-only instance of LightweightState, bound to a specific deployed contract.
func NewLightweightStateTransactor(address common.Address, transactor bind.ContractTransactor) (*LightweightStateTransactor, error) {
	contract, err := bindLightweightState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightweightStateTransactor{contract: contract}, nil
}

// NewLightweightStateFilterer creates a new log filterer instance of LightweightState, bound to a specific deployed contract.
func NewLightweightStateFilterer(address common.Address, filterer bind.ContractFilterer) (*LightweightStateFilterer, error) {
	contract, err := bindLightweightState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightweightStateFilterer{contract: contract}, nil
}

// bindLightweightState binds a generic wrapper to an already deployed contract.
func bindLightweightState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LightweightStateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightweightState *LightweightStateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightweightState.Contract.LightweightStateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightweightState *LightweightStateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightweightState.Contract.LightweightStateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightweightState *LightweightStateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightweightState.Contract.LightweightStateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightweightState *LightweightStateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightweightState.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightweightState *LightweightStateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightweightState.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightweightState *LightweightStateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightweightState.Contract.contract.Transact(opts, method, params...)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_LightweightState *LightweightStateCaller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "P")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_LightweightState *LightweightStateSession) P() (*big.Int, error) {
	return _LightweightState.Contract.P(&_LightweightState.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_LightweightState *LightweightStateCallerSession) P() (*big.Int, error) {
	return _LightweightState.Contract.P(&_LightweightState.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_LightweightState *LightweightStateCaller) ChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "chainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_LightweightState *LightweightStateSession) ChainName() (string, error) {
	return _LightweightState.Contract.ChainName(&_LightweightState.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_LightweightState *LightweightStateCallerSession) ChainName() (string, error) {
	return _LightweightState.Contract.ChainName(&_LightweightState.CallOpts)
}

// GeGISTRootData is a free data retrieval call binding the contract method 0x0dd93b5d.
//
// Solidity: function geGISTRootData(uint256 root_) view returns((uint256,uint256))
func (_LightweightState *LightweightStateCaller) GeGISTRootData(opts *bind.CallOpts, root_ *big.Int) (ILightweightStateGistRootData, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "geGISTRootData", root_)

	if err != nil {
		return *new(ILightweightStateGistRootData), err
	}

	out0 := *abi.ConvertType(out[0], new(ILightweightStateGistRootData)).(*ILightweightStateGistRootData)

	return out0, err

}

// GeGISTRootData is a free data retrieval call binding the contract method 0x0dd93b5d.
//
// Solidity: function geGISTRootData(uint256 root_) view returns((uint256,uint256))
func (_LightweightState *LightweightStateSession) GeGISTRootData(root_ *big.Int) (ILightweightStateGistRootData, error) {
	return _LightweightState.Contract.GeGISTRootData(&_LightweightState.CallOpts, root_)
}

// GeGISTRootData is a free data retrieval call binding the contract method 0x0dd93b5d.
//
// Solidity: function geGISTRootData(uint256 root_) view returns((uint256,uint256))
func (_LightweightState *LightweightStateCallerSession) GeGISTRootData(root_ *big.Int) (ILightweightStateGistRootData, error) {
	return _LightweightState.Contract.GeGISTRootData(&_LightweightState.CallOpts, root_)
}

// GetCurrentGISTRootInfo is a free data retrieval call binding the contract method 0xaf7a3f59.
//
// Solidity: function getCurrentGISTRootInfo() view returns((uint256,uint256))
func (_LightweightState *LightweightStateCaller) GetCurrentGISTRootInfo(opts *bind.CallOpts) (ILightweightStateGistRootData, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "getCurrentGISTRootInfo")

	if err != nil {
		return *new(ILightweightStateGistRootData), err
	}

	out0 := *abi.ConvertType(out[0], new(ILightweightStateGistRootData)).(*ILightweightStateGistRootData)

	return out0, err

}

// GetCurrentGISTRootInfo is a free data retrieval call binding the contract method 0xaf7a3f59.
//
// Solidity: function getCurrentGISTRootInfo() view returns((uint256,uint256))
func (_LightweightState *LightweightStateSession) GetCurrentGISTRootInfo() (ILightweightStateGistRootData, error) {
	return _LightweightState.Contract.GetCurrentGISTRootInfo(&_LightweightState.CallOpts)
}

// GetCurrentGISTRootInfo is a free data retrieval call binding the contract method 0xaf7a3f59.
//
// Solidity: function getCurrentGISTRootInfo() view returns((uint256,uint256))
func (_LightweightState *LightweightStateCallerSession) GetCurrentGISTRootInfo() (ILightweightStateGistRootData, error) {
	return _LightweightState.Contract.GetCurrentGISTRootInfo(&_LightweightState.CallOpts)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_LightweightState *LightweightStateCaller) GetGISTRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "getGISTRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_LightweightState *LightweightStateSession) GetGISTRoot() (*big.Int, error) {
	return _LightweightState.Contract.GetGISTRoot(&_LightweightState.CallOpts)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_LightweightState *LightweightStateCallerSession) GetGISTRoot() (*big.Int, error) {
	return _LightweightState.Contract.GetGISTRoot(&_LightweightState.CallOpts)
}

// GetIdentitiesStatesRootData is a free data retrieval call binding the contract method 0xa055a692.
//
// Solidity: function getIdentitiesStatesRootData(bytes32 root_) view returns((bytes32,uint256))
func (_LightweightState *LightweightStateCaller) GetIdentitiesStatesRootData(opts *bind.CallOpts, root_ [32]byte) (ILightweightStateIdentitiesStatesRootData, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "getIdentitiesStatesRootData", root_)

	if err != nil {
		return *new(ILightweightStateIdentitiesStatesRootData), err
	}

	out0 := *abi.ConvertType(out[0], new(ILightweightStateIdentitiesStatesRootData)).(*ILightweightStateIdentitiesStatesRootData)

	return out0, err

}

// GetIdentitiesStatesRootData is a free data retrieval call binding the contract method 0xa055a692.
//
// Solidity: function getIdentitiesStatesRootData(bytes32 root_) view returns((bytes32,uint256))
func (_LightweightState *LightweightStateSession) GetIdentitiesStatesRootData(root_ [32]byte) (ILightweightStateIdentitiesStatesRootData, error) {
	return _LightweightState.Contract.GetIdentitiesStatesRootData(&_LightweightState.CallOpts, root_)
}

// GetIdentitiesStatesRootData is a free data retrieval call binding the contract method 0xa055a692.
//
// Solidity: function getIdentitiesStatesRootData(bytes32 root_) view returns((bytes32,uint256))
func (_LightweightState *LightweightStateCallerSession) GetIdentitiesStatesRootData(root_ [32]byte) (ILightweightStateIdentitiesStatesRootData, error) {
	return _LightweightState.Contract.GetIdentitiesStatesRootData(&_LightweightState.CallOpts, root_)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_LightweightState *LightweightStateCaller) GetSigComponents(opts *bind.CallOpts, methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "getSigComponents", methodId_, contractAddress_)

	outstruct := new(struct {
		ChainName string
		Nonce     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Nonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_LightweightState *LightweightStateSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _LightweightState.Contract.GetSigComponents(&_LightweightState.CallOpts, methodId_, contractAddress_)
}

// GetSigComponents is a free data retrieval call binding the contract method 0x827e099e.
//
// Solidity: function getSigComponents(uint8 methodId_, address contractAddress_) view returns(string chainName_, uint256 nonce_)
func (_LightweightState *LightweightStateCallerSession) GetSigComponents(methodId_ uint8, contractAddress_ common.Address) (struct {
	ChainName string
	Nonce     *big.Int
}, error) {
	return _LightweightState.Contract.GetSigComponents(&_LightweightState.CallOpts, methodId_, contractAddress_)
}

// IdentitiesStatesRoot is a free data retrieval call binding the contract method 0xe08e70bb.
//
// Solidity: function identitiesStatesRoot() view returns(bytes32)
func (_LightweightState *LightweightStateCaller) IdentitiesStatesRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "identitiesStatesRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IdentitiesStatesRoot is a free data retrieval call binding the contract method 0xe08e70bb.
//
// Solidity: function identitiesStatesRoot() view returns(bytes32)
func (_LightweightState *LightweightStateSession) IdentitiesStatesRoot() ([32]byte, error) {
	return _LightweightState.Contract.IdentitiesStatesRoot(&_LightweightState.CallOpts)
}

// IdentitiesStatesRoot is a free data retrieval call binding the contract method 0xe08e70bb.
//
// Solidity: function identitiesStatesRoot() view returns(bytes32)
func (_LightweightState *LightweightStateCallerSession) IdentitiesStatesRoot() ([32]byte, error) {
	return _LightweightState.Contract.IdentitiesStatesRoot(&_LightweightState.CallOpts)
}

// IsIdentitiesStatesRootExists is a free data retrieval call binding the contract method 0xbfd73455.
//
// Solidity: function isIdentitiesStatesRootExists(bytes32 root_) view returns(bool)
func (_LightweightState *LightweightStateCaller) IsIdentitiesStatesRootExists(opts *bind.CallOpts, root_ [32]byte) (bool, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "isIdentitiesStatesRootExists", root_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdentitiesStatesRootExists is a free data retrieval call binding the contract method 0xbfd73455.
//
// Solidity: function isIdentitiesStatesRootExists(bytes32 root_) view returns(bool)
func (_LightweightState *LightweightStateSession) IsIdentitiesStatesRootExists(root_ [32]byte) (bool, error) {
	return _LightweightState.Contract.IsIdentitiesStatesRootExists(&_LightweightState.CallOpts, root_)
}

// IsIdentitiesStatesRootExists is a free data retrieval call binding the contract method 0xbfd73455.
//
// Solidity: function isIdentitiesStatesRootExists(bytes32 root_) view returns(bool)
func (_LightweightState *LightweightStateCallerSession) IsIdentitiesStatesRootExists(root_ [32]byte) (bool, error) {
	return _LightweightState.Contract.IsIdentitiesStatesRootExists(&_LightweightState.CallOpts, root_)
}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_LightweightState *LightweightStateCaller) Nonces(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (*big.Int, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "nonces", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_LightweightState *LightweightStateSession) Nonces(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _LightweightState.Contract.Nonces(&_LightweightState.CallOpts, arg0, arg1)
}

// Nonces is a free data retrieval call binding the contract method 0xed3218a2.
//
// Solidity: function nonces(address , uint8 ) view returns(uint256)
func (_LightweightState *LightweightStateCallerSession) Nonces(arg0 common.Address, arg1 uint8) (*big.Int, error) {
	return _LightweightState.Contract.Nonces(&_LightweightState.CallOpts, arg0, arg1)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LightweightState *LightweightStateCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LightweightState *LightweightStateSession) ProxiableUUID() ([32]byte, error) {
	return _LightweightState.Contract.ProxiableUUID(&_LightweightState.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_LightweightState *LightweightStateCallerSession) ProxiableUUID() ([32]byte, error) {
	return _LightweightState.Contract.ProxiableUUID(&_LightweightState.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_LightweightState *LightweightStateCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_LightweightState *LightweightStateSession) Signer() (common.Address, error) {
	return _LightweightState.Contract.Signer(&_LightweightState.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_LightweightState *LightweightStateCallerSession) Signer() (common.Address, error) {
	return _LightweightState.Contract.Signer(&_LightweightState.CallOpts)
}

// SourceChainName is a free data retrieval call binding the contract method 0xe4ffd04a.
//
// Solidity: function sourceChainName() view returns(string)
func (_LightweightState *LightweightStateCaller) SourceChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "sourceChainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SourceChainName is a free data retrieval call binding the contract method 0xe4ffd04a.
//
// Solidity: function sourceChainName() view returns(string)
func (_LightweightState *LightweightStateSession) SourceChainName() (string, error) {
	return _LightweightState.Contract.SourceChainName(&_LightweightState.CallOpts)
}

// SourceChainName is a free data retrieval call binding the contract method 0xe4ffd04a.
//
// Solidity: function sourceChainName() view returns(string)
func (_LightweightState *LightweightStateCallerSession) SourceChainName() (string, error) {
	return _LightweightState.Contract.SourceChainName(&_LightweightState.CallOpts)
}

// SourceStateContract is a free data retrieval call binding the contract method 0xfc228319.
//
// Solidity: function sourceStateContract() view returns(address)
func (_LightweightState *LightweightStateCaller) SourceStateContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "sourceStateContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SourceStateContract is a free data retrieval call binding the contract method 0xfc228319.
//
// Solidity: function sourceStateContract() view returns(address)
func (_LightweightState *LightweightStateSession) SourceStateContract() (common.Address, error) {
	return _LightweightState.Contract.SourceStateContract(&_LightweightState.CallOpts)
}

// SourceStateContract is a free data retrieval call binding the contract method 0xfc228319.
//
// Solidity: function sourceStateContract() view returns(address)
func (_LightweightState *LightweightStateCallerSession) SourceStateContract() (common.Address, error) {
	return _LightweightState.Contract.SourceStateContract(&_LightweightState.CallOpts)
}

// VerifyStatesMerkleData is a free data retrieval call binding the contract method 0xb0d46f2c.
//
// Solidity: function verifyStatesMerkleData((uint256,uint256,uint256,bytes32[]) statesMerkleData_) view returns(bool, bytes32)
func (_LightweightState *LightweightStateCaller) VerifyStatesMerkleData(opts *bind.CallOpts, statesMerkleData_ ILightweightStateStatesMerkleData) (bool, [32]byte, error) {
	var out []interface{}
	err := _LightweightState.contract.Call(opts, &out, "verifyStatesMerkleData", statesMerkleData_)

	if err != nil {
		return *new(bool), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// VerifyStatesMerkleData is a free data retrieval call binding the contract method 0xb0d46f2c.
//
// Solidity: function verifyStatesMerkleData((uint256,uint256,uint256,bytes32[]) statesMerkleData_) view returns(bool, bytes32)
func (_LightweightState *LightweightStateSession) VerifyStatesMerkleData(statesMerkleData_ ILightweightStateStatesMerkleData) (bool, [32]byte, error) {
	return _LightweightState.Contract.VerifyStatesMerkleData(&_LightweightState.CallOpts, statesMerkleData_)
}

// VerifyStatesMerkleData is a free data retrieval call binding the contract method 0xb0d46f2c.
//
// Solidity: function verifyStatesMerkleData((uint256,uint256,uint256,bytes32[]) statesMerkleData_) view returns(bool, bytes32)
func (_LightweightState *LightweightStateCallerSession) VerifyStatesMerkleData(statesMerkleData_ ILightweightStateStatesMerkleData) (bool, [32]byte, error) {
	return _LightweightState.Contract.VerifyStatesMerkleData(&_LightweightState.CallOpts, statesMerkleData_)
}

// LightweightStateInit is a paid mutator transaction binding the contract method 0x27e25ecf.
//
// Solidity: function __LightweightState_init(address signer_, address sourceStateContract_, string sourceChainName_, string chainName_) returns()
func (_LightweightState *LightweightStateTransactor) LightweightStateInit(opts *bind.TransactOpts, signer_ common.Address, sourceStateContract_ common.Address, sourceChainName_ string, chainName_ string) (*types.Transaction, error) {
	return _LightweightState.contract.Transact(opts, "__LightweightState_init", signer_, sourceStateContract_, sourceChainName_, chainName_)
}

// LightweightStateInit is a paid mutator transaction binding the contract method 0x27e25ecf.
//
// Solidity: function __LightweightState_init(address signer_, address sourceStateContract_, string sourceChainName_, string chainName_) returns()
func (_LightweightState *LightweightStateSession) LightweightStateInit(signer_ common.Address, sourceStateContract_ common.Address, sourceChainName_ string, chainName_ string) (*types.Transaction, error) {
	return _LightweightState.Contract.LightweightStateInit(&_LightweightState.TransactOpts, signer_, sourceStateContract_, sourceChainName_, chainName_)
}

// LightweightStateInit is a paid mutator transaction binding the contract method 0x27e25ecf.
//
// Solidity: function __LightweightState_init(address signer_, address sourceStateContract_, string sourceChainName_, string chainName_) returns()
func (_LightweightState *LightweightStateTransactorSession) LightweightStateInit(signer_ common.Address, sourceStateContract_ common.Address, sourceChainName_ string, chainName_ string) (*types.Transaction, error) {
	return _LightweightState.Contract.LightweightStateInit(&_LightweightState.TransactOpts, signer_, sourceStateContract_, sourceChainName_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_LightweightState *LightweightStateTransactor) SignersInit(opts *bind.TransactOpts, signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _LightweightState.contract.Transact(opts, "__Signers_init", signer_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_LightweightState *LightweightStateSession) SignersInit(signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _LightweightState.Contract.SignersInit(&_LightweightState.TransactOpts, signer_, chainName_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x509ace95.
//
// Solidity: function __Signers_init(address signer_, string chainName_) returns()
func (_LightweightState *LightweightStateTransactorSession) SignersInit(signer_ common.Address, chainName_ string) (*types.Transaction, error) {
	return _LightweightState.Contract.SignersInit(&_LightweightState.TransactOpts, signer_, chainName_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_LightweightState *LightweightStateTransactor) ChangeSigner(opts *bind.TransactOpts, newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.contract.Transact(opts, "changeSigner", newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_LightweightState *LightweightStateSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.ChangeSigner(&_LightweightState.TransactOpts, newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_LightweightState *LightweightStateTransactorSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.ChangeSigner(&_LightweightState.TransactOpts, newSignerPubKey_, signature_)
}

// ChangeSourceStateContract is a paid mutator transaction binding the contract method 0x89aeb0f5.
//
// Solidity: function changeSourceStateContract(address newSourceStateContract_, bytes signature_) returns()
func (_LightweightState *LightweightStateTransactor) ChangeSourceStateContract(opts *bind.TransactOpts, newSourceStateContract_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.contract.Transact(opts, "changeSourceStateContract", newSourceStateContract_, signature_)
}

// ChangeSourceStateContract is a paid mutator transaction binding the contract method 0x89aeb0f5.
//
// Solidity: function changeSourceStateContract(address newSourceStateContract_, bytes signature_) returns()
func (_LightweightState *LightweightStateSession) ChangeSourceStateContract(newSourceStateContract_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.ChangeSourceStateContract(&_LightweightState.TransactOpts, newSourceStateContract_, signature_)
}

// ChangeSourceStateContract is a paid mutator transaction binding the contract method 0x89aeb0f5.
//
// Solidity: function changeSourceStateContract(address newSourceStateContract_, bytes signature_) returns()
func (_LightweightState *LightweightStateTransactorSession) ChangeSourceStateContract(newSourceStateContract_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.ChangeSourceStateContract(&_LightweightState.TransactOpts, newSourceStateContract_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_LightweightState *LightweightStateTransactor) CheckSignatureAndIncrementNonce(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.contract.Transact(opts, "checkSignatureAndIncrementNonce", methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_LightweightState *LightweightStateSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.CheckSignatureAndIncrementNonce(&_LightweightState.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// CheckSignatureAndIncrementNonce is a paid mutator transaction binding the contract method 0xe3754f90.
//
// Solidity: function checkSignatureAndIncrementNonce(uint8 methodId_, address contractAddress_, bytes32 signHash_, bytes signature_) returns()
func (_LightweightState *LightweightStateTransactorSession) CheckSignatureAndIncrementNonce(methodId_ uint8, contractAddress_ common.Address, signHash_ [32]byte, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.CheckSignatureAndIncrementNonce(&_LightweightState.TransactOpts, methodId_, contractAddress_, signHash_, signature_)
}

// SignedTransitState is a paid mutator transaction binding the contract method 0x189a5073.
//
// Solidity: function signedTransitState(bytes32 newIdentitiesStatesRoot_, (uint256,uint256) gistData_, bytes proof_) returns()
func (_LightweightState *LightweightStateTransactor) SignedTransitState(opts *bind.TransactOpts, newIdentitiesStatesRoot_ [32]byte, gistData_ ILightweightStateGistRootData, proof_ []byte) (*types.Transaction, error) {
	return _LightweightState.contract.Transact(opts, "signedTransitState", newIdentitiesStatesRoot_, gistData_, proof_)
}

// SignedTransitState is a paid mutator transaction binding the contract method 0x189a5073.
//
// Solidity: function signedTransitState(bytes32 newIdentitiesStatesRoot_, (uint256,uint256) gistData_, bytes proof_) returns()
func (_LightweightState *LightweightStateSession) SignedTransitState(newIdentitiesStatesRoot_ [32]byte, gistData_ ILightweightStateGistRootData, proof_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.SignedTransitState(&_LightweightState.TransactOpts, newIdentitiesStatesRoot_, gistData_, proof_)
}

// SignedTransitState is a paid mutator transaction binding the contract method 0x189a5073.
//
// Solidity: function signedTransitState(bytes32 newIdentitiesStatesRoot_, (uint256,uint256) gistData_, bytes proof_) returns()
func (_LightweightState *LightweightStateTransactorSession) SignedTransitState(newIdentitiesStatesRoot_ [32]byte, gistData_ ILightweightStateGistRootData, proof_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.SignedTransitState(&_LightweightState.TransactOpts, newIdentitiesStatesRoot_, gistData_, proof_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LightweightState *LightweightStateTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _LightweightState.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LightweightState *LightweightStateSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LightweightState.Contract.UpgradeTo(&_LightweightState.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_LightweightState *LightweightStateTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _LightweightState.Contract.UpgradeTo(&_LightweightState.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LightweightState *LightweightStateTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LightweightState.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LightweightState *LightweightStateSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.UpgradeToAndCall(&_LightweightState.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_LightweightState *LightweightStateTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.UpgradeToAndCall(&_LightweightState.TransactOpts, newImplementation, data)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_LightweightState *LightweightStateTransactor) UpgradeToWithSig(opts *bind.TransactOpts, newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.contract.Transact(opts, "upgradeToWithSig", newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_LightweightState *LightweightStateSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.UpgradeToWithSig(&_LightweightState.TransactOpts, newImplementation_, signature_)
}

// UpgradeToWithSig is a paid mutator transaction binding the contract method 0x52d04661.
//
// Solidity: function upgradeToWithSig(address newImplementation_, bytes signature_) returns()
func (_LightweightState *LightweightStateTransactorSession) UpgradeToWithSig(newImplementation_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.UpgradeToWithSig(&_LightweightState.TransactOpts, newImplementation_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_LightweightState *LightweightStateTransactor) ValidateChangeAddressSignature(opts *bind.TransactOpts, methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.contract.Transact(opts, "validateChangeAddressSignature", methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_LightweightState *LightweightStateSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.ValidateChangeAddressSignature(&_LightweightState.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// ValidateChangeAddressSignature is a paid mutator transaction binding the contract method 0x7d1e764b.
//
// Solidity: function validateChangeAddressSignature(uint8 methodId_, address contractAddress_, address newAddress_, bytes signature_) returns()
func (_LightweightState *LightweightStateTransactorSession) ValidateChangeAddressSignature(methodId_ uint8, contractAddress_ common.Address, newAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _LightweightState.Contract.ValidateChangeAddressSignature(&_LightweightState.TransactOpts, methodId_, contractAddress_, newAddress_, signature_)
}

// LightweightStateAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the LightweightState contract.
type LightweightStateAdminChangedIterator struct {
	Event *LightweightStateAdminChanged // Event containing the contract specifics and raw log

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
func (it *LightweightStateAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightweightStateAdminChanged)
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
		it.Event = new(LightweightStateAdminChanged)
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
func (it *LightweightStateAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightweightStateAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightweightStateAdminChanged represents a AdminChanged event raised by the LightweightState contract.
type LightweightStateAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LightweightState *LightweightStateFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LightweightStateAdminChangedIterator, error) {

	logs, sub, err := _LightweightState.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LightweightStateAdminChangedIterator{contract: _LightweightState.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_LightweightState *LightweightStateFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LightweightStateAdminChanged) (event.Subscription, error) {

	logs, sub, err := _LightweightState.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightweightStateAdminChanged)
				if err := _LightweightState.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_LightweightState *LightweightStateFilterer) ParseAdminChanged(log types.Log) (*LightweightStateAdminChanged, error) {
	event := new(LightweightStateAdminChanged)
	if err := _LightweightState.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightweightStateBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the LightweightState contract.
type LightweightStateBeaconUpgradedIterator struct {
	Event *LightweightStateBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *LightweightStateBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightweightStateBeaconUpgraded)
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
		it.Event = new(LightweightStateBeaconUpgraded)
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
func (it *LightweightStateBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightweightStateBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightweightStateBeaconUpgraded represents a BeaconUpgraded event raised by the LightweightState contract.
type LightweightStateBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LightweightState *LightweightStateFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LightweightStateBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LightweightState.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LightweightStateBeaconUpgradedIterator{contract: _LightweightState.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_LightweightState *LightweightStateFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LightweightStateBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _LightweightState.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightweightStateBeaconUpgraded)
				if err := _LightweightState.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_LightweightState *LightweightStateFilterer) ParseBeaconUpgraded(log types.Log) (*LightweightStateBeaconUpgraded, error) {
	event := new(LightweightStateBeaconUpgraded)
	if err := _LightweightState.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightweightStateInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LightweightState contract.
type LightweightStateInitializedIterator struct {
	Event *LightweightStateInitialized // Event containing the contract specifics and raw log

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
func (it *LightweightStateInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightweightStateInitialized)
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
		it.Event = new(LightweightStateInitialized)
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
func (it *LightweightStateInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightweightStateInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightweightStateInitialized represents a Initialized event raised by the LightweightState contract.
type LightweightStateInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightweightState *LightweightStateFilterer) FilterInitialized(opts *bind.FilterOpts) (*LightweightStateInitializedIterator, error) {

	logs, sub, err := _LightweightState.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LightweightStateInitializedIterator{contract: _LightweightState.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightweightState *LightweightStateFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LightweightStateInitialized) (event.Subscription, error) {

	logs, sub, err := _LightweightState.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightweightStateInitialized)
				if err := _LightweightState.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LightweightState *LightweightStateFilterer) ParseInitialized(log types.Log) (*LightweightStateInitialized, error) {
	event := new(LightweightStateInitialized)
	if err := _LightweightState.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightweightStateSignedStateTransitedIterator is returned from FilterSignedStateTransited and is used to iterate over the raw logs and unpacked data for SignedStateTransited events raised by the LightweightState contract.
type LightweightStateSignedStateTransitedIterator struct {
	Event *LightweightStateSignedStateTransited // Event containing the contract specifics and raw log

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
func (it *LightweightStateSignedStateTransitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightweightStateSignedStateTransited)
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
		it.Event = new(LightweightStateSignedStateTransited)
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
func (it *LightweightStateSignedStateTransitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightweightStateSignedStateTransitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightweightStateSignedStateTransited represents a SignedStateTransited event raised by the LightweightState contract.
type LightweightStateSignedStateTransited struct {
	NewGistRoot            *big.Int
	NewIdentitesStatesRoot [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSignedStateTransited is a free log retrieval operation binding the contract event 0x8e8ff16553fbf4a457c08a5e07cc27c8aac14b9a1a8e1f546a6c1b9366304a56.
//
// Solidity: event SignedStateTransited(uint256 newGistRoot, bytes32 newIdentitesStatesRoot)
func (_LightweightState *LightweightStateFilterer) FilterSignedStateTransited(opts *bind.FilterOpts) (*LightweightStateSignedStateTransitedIterator, error) {

	logs, sub, err := _LightweightState.contract.FilterLogs(opts, "SignedStateTransited")
	if err != nil {
		return nil, err
	}
	return &LightweightStateSignedStateTransitedIterator{contract: _LightweightState.contract, event: "SignedStateTransited", logs: logs, sub: sub}, nil
}

// WatchSignedStateTransited is a free log subscription operation binding the contract event 0x8e8ff16553fbf4a457c08a5e07cc27c8aac14b9a1a8e1f546a6c1b9366304a56.
//
// Solidity: event SignedStateTransited(uint256 newGistRoot, bytes32 newIdentitesStatesRoot)
func (_LightweightState *LightweightStateFilterer) WatchSignedStateTransited(opts *bind.WatchOpts, sink chan<- *LightweightStateSignedStateTransited) (event.Subscription, error) {

	logs, sub, err := _LightweightState.contract.WatchLogs(opts, "SignedStateTransited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightweightStateSignedStateTransited)
				if err := _LightweightState.contract.UnpackLog(event, "SignedStateTransited", log); err != nil {
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

// ParseSignedStateTransited is a log parse operation binding the contract event 0x8e8ff16553fbf4a457c08a5e07cc27c8aac14b9a1a8e1f546a6c1b9366304a56.
//
// Solidity: event SignedStateTransited(uint256 newGistRoot, bytes32 newIdentitesStatesRoot)
func (_LightweightState *LightweightStateFilterer) ParseSignedStateTransited(log types.Log) (*LightweightStateSignedStateTransited, error) {
	event := new(LightweightStateSignedStateTransited)
	if err := _LightweightState.contract.UnpackLog(event, "SignedStateTransited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightweightStateUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the LightweightState contract.
type LightweightStateUpgradedIterator struct {
	Event *LightweightStateUpgraded // Event containing the contract specifics and raw log

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
func (it *LightweightStateUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightweightStateUpgraded)
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
		it.Event = new(LightweightStateUpgraded)
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
func (it *LightweightStateUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightweightStateUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightweightStateUpgraded represents a Upgraded event raised by the LightweightState contract.
type LightweightStateUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LightweightState *LightweightStateFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LightweightStateUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LightweightState.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LightweightStateUpgradedIterator{contract: _LightweightState.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_LightweightState *LightweightStateFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LightweightStateUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _LightweightState.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightweightStateUpgraded)
				if err := _LightweightState.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_LightweightState *LightweightStateFilterer) ParseUpgraded(log types.Log) (*LightweightStateUpgraded, error) {
	event := new(LightweightStateUpgraded)
	if err := _LightweightState.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
