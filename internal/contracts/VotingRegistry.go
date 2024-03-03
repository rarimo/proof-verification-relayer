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

// VotingRegistryMetaData contains all meta data concerning the VotingRegistry contract.
var VotingRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolFactory_\",\"type\":\"address\"}],\"name\":\"__VotingRegistry_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"poolType_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proposer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"name\":\"addProxyPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"voting_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"registration_\",\"type\":\"address\"}],\"name\":\"bindVotingToRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"poolType_\",\"type\":\"string\"}],\"name\":\"getPoolImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"registration_\",\"type\":\"address\"}],\"name\":\"getVotingForRegistration\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"name\":\"isPoolExistByProposer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"poolType_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"name\":\"isPoolExistByProposerAndType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"poolType_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"name\":\"isPoolExistByType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"listPoolsByProposer\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pools_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"poolType_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"listPoolsByProposerAndType\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pools_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"poolType_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"listPoolsByType\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pools_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer_\",\"type\":\"address\"}],\"name\":\"poolCountByProposer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"poolType_\",\"type\":\"string\"}],\"name\":\"poolCountByProposerAndType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"poolType_\",\"type\":\"string\"}],\"name\":\"poolCountByType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"poolTypes_\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"newImplementations_\",\"type\":\"address[]\"}],\"name\":\"setNewImplementations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// VotingRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingRegistryMetaData.ABI instead.
var VotingRegistryABI = VotingRegistryMetaData.ABI

// VotingRegistry is an auto generated Go binding around an Ethereum contract.
type VotingRegistry struct {
	VotingRegistryCaller     // Read-only binding to the contract
	VotingRegistryTransactor // Write-only binding to the contract
	VotingRegistryFilterer   // Log filterer for contract events
}

// VotingRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingRegistrySession struct {
	Contract     *VotingRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingRegistryCallerSession struct {
	Contract *VotingRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VotingRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingRegistryTransactorSession struct {
	Contract     *VotingRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VotingRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRegistryRaw struct {
	Contract *VotingRegistry // Generic contract binding to access the raw methods on
}

// VotingRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingRegistryCallerRaw struct {
	Contract *VotingRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// VotingRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingRegistryTransactorRaw struct {
	Contract *VotingRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotingRegistry creates a new instance of VotingRegistry, bound to a specific deployed contract.
func NewVotingRegistry(address common.Address, backend bind.ContractBackend) (*VotingRegistry, error) {
	contract, err := bindVotingRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotingRegistry{VotingRegistryCaller: VotingRegistryCaller{contract: contract}, VotingRegistryTransactor: VotingRegistryTransactor{contract: contract}, VotingRegistryFilterer: VotingRegistryFilterer{contract: contract}}, nil
}

// NewVotingRegistryCaller creates a new read-only instance of VotingRegistry, bound to a specific deployed contract.
func NewVotingRegistryCaller(address common.Address, caller bind.ContractCaller) (*VotingRegistryCaller, error) {
	contract, err := bindVotingRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingRegistryCaller{contract: contract}, nil
}

// NewVotingRegistryTransactor creates a new write-only instance of VotingRegistry, bound to a specific deployed contract.
func NewVotingRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingRegistryTransactor, error) {
	contract, err := bindVotingRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingRegistryTransactor{contract: contract}, nil
}

// NewVotingRegistryFilterer creates a new log filterer instance of VotingRegistry, bound to a specific deployed contract.
func NewVotingRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingRegistryFilterer, error) {
	contract, err := bindVotingRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingRegistryFilterer{contract: contract}, nil
}

// bindVotingRegistry binds a generic wrapper to an already deployed contract.
func bindVotingRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingRegistry *VotingRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotingRegistry.Contract.VotingRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingRegistry *VotingRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingRegistry.Contract.VotingRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingRegistry *VotingRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingRegistry.Contract.VotingRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingRegistry *VotingRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotingRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingRegistry *VotingRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingRegistry *VotingRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetPoolImplementation is a free data retrieval call binding the contract method 0x5fd2b459.
//
// Solidity: function getPoolImplementation(string poolType_) view returns(address)
func (_VotingRegistry *VotingRegistryCaller) GetPoolImplementation(opts *bind.CallOpts, poolType_ string) (common.Address, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "getPoolImplementation", poolType_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolImplementation is a free data retrieval call binding the contract method 0x5fd2b459.
//
// Solidity: function getPoolImplementation(string poolType_) view returns(address)
func (_VotingRegistry *VotingRegistrySession) GetPoolImplementation(poolType_ string) (common.Address, error) {
	return _VotingRegistry.Contract.GetPoolImplementation(&_VotingRegistry.CallOpts, poolType_)
}

// GetPoolImplementation is a free data retrieval call binding the contract method 0x5fd2b459.
//
// Solidity: function getPoolImplementation(string poolType_) view returns(address)
func (_VotingRegistry *VotingRegistryCallerSession) GetPoolImplementation(poolType_ string) (common.Address, error) {
	return _VotingRegistry.Contract.GetPoolImplementation(&_VotingRegistry.CallOpts, poolType_)
}

// GetVotingForRegistration is a free data retrieval call binding the contract method 0xa2daa62a.
//
// Solidity: function getVotingForRegistration(address proposer_, address registration_) view returns(address)
func (_VotingRegistry *VotingRegistryCaller) GetVotingForRegistration(opts *bind.CallOpts, proposer_ common.Address, registration_ common.Address) (common.Address, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "getVotingForRegistration", proposer_, registration_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVotingForRegistration is a free data retrieval call binding the contract method 0xa2daa62a.
//
// Solidity: function getVotingForRegistration(address proposer_, address registration_) view returns(address)
func (_VotingRegistry *VotingRegistrySession) GetVotingForRegistration(proposer_ common.Address, registration_ common.Address) (common.Address, error) {
	return _VotingRegistry.Contract.GetVotingForRegistration(&_VotingRegistry.CallOpts, proposer_, registration_)
}

// GetVotingForRegistration is a free data retrieval call binding the contract method 0xa2daa62a.
//
// Solidity: function getVotingForRegistration(address proposer_, address registration_) view returns(address)
func (_VotingRegistry *VotingRegistryCallerSession) GetVotingForRegistration(proposer_ common.Address, registration_ common.Address) (common.Address, error) {
	return _VotingRegistry.Contract.GetVotingForRegistration(&_VotingRegistry.CallOpts, proposer_, registration_)
}

// IsPoolExistByProposer is a free data retrieval call binding the contract method 0xd4ba209a.
//
// Solidity: function isPoolExistByProposer(address proposer_, address pool_) view returns(bool)
func (_VotingRegistry *VotingRegistryCaller) IsPoolExistByProposer(opts *bind.CallOpts, proposer_ common.Address, pool_ common.Address) (bool, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "isPoolExistByProposer", proposer_, pool_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolExistByProposer is a free data retrieval call binding the contract method 0xd4ba209a.
//
// Solidity: function isPoolExistByProposer(address proposer_, address pool_) view returns(bool)
func (_VotingRegistry *VotingRegistrySession) IsPoolExistByProposer(proposer_ common.Address, pool_ common.Address) (bool, error) {
	return _VotingRegistry.Contract.IsPoolExistByProposer(&_VotingRegistry.CallOpts, proposer_, pool_)
}

// IsPoolExistByProposer is a free data retrieval call binding the contract method 0xd4ba209a.
//
// Solidity: function isPoolExistByProposer(address proposer_, address pool_) view returns(bool)
func (_VotingRegistry *VotingRegistryCallerSession) IsPoolExistByProposer(proposer_ common.Address, pool_ common.Address) (bool, error) {
	return _VotingRegistry.Contract.IsPoolExistByProposer(&_VotingRegistry.CallOpts, proposer_, pool_)
}

// IsPoolExistByProposerAndType is a free data retrieval call binding the contract method 0x581a9a47.
//
// Solidity: function isPoolExistByProposerAndType(address proposer_, string poolType_, address pool_) view returns(bool)
func (_VotingRegistry *VotingRegistryCaller) IsPoolExistByProposerAndType(opts *bind.CallOpts, proposer_ common.Address, poolType_ string, pool_ common.Address) (bool, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "isPoolExistByProposerAndType", proposer_, poolType_, pool_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolExistByProposerAndType is a free data retrieval call binding the contract method 0x581a9a47.
//
// Solidity: function isPoolExistByProposerAndType(address proposer_, string poolType_, address pool_) view returns(bool)
func (_VotingRegistry *VotingRegistrySession) IsPoolExistByProposerAndType(proposer_ common.Address, poolType_ string, pool_ common.Address) (bool, error) {
	return _VotingRegistry.Contract.IsPoolExistByProposerAndType(&_VotingRegistry.CallOpts, proposer_, poolType_, pool_)
}

// IsPoolExistByProposerAndType is a free data retrieval call binding the contract method 0x581a9a47.
//
// Solidity: function isPoolExistByProposerAndType(address proposer_, string poolType_, address pool_) view returns(bool)
func (_VotingRegistry *VotingRegistryCallerSession) IsPoolExistByProposerAndType(proposer_ common.Address, poolType_ string, pool_ common.Address) (bool, error) {
	return _VotingRegistry.Contract.IsPoolExistByProposerAndType(&_VotingRegistry.CallOpts, proposer_, poolType_, pool_)
}

// IsPoolExistByType is a free data retrieval call binding the contract method 0xd9134bb6.
//
// Solidity: function isPoolExistByType(string poolType_, address pool_) view returns(bool)
func (_VotingRegistry *VotingRegistryCaller) IsPoolExistByType(opts *bind.CallOpts, poolType_ string, pool_ common.Address) (bool, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "isPoolExistByType", poolType_, pool_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolExistByType is a free data retrieval call binding the contract method 0xd9134bb6.
//
// Solidity: function isPoolExistByType(string poolType_, address pool_) view returns(bool)
func (_VotingRegistry *VotingRegistrySession) IsPoolExistByType(poolType_ string, pool_ common.Address) (bool, error) {
	return _VotingRegistry.Contract.IsPoolExistByType(&_VotingRegistry.CallOpts, poolType_, pool_)
}

// IsPoolExistByType is a free data retrieval call binding the contract method 0xd9134bb6.
//
// Solidity: function isPoolExistByType(string poolType_, address pool_) view returns(bool)
func (_VotingRegistry *VotingRegistryCallerSession) IsPoolExistByType(poolType_ string, pool_ common.Address) (bool, error) {
	return _VotingRegistry.Contract.IsPoolExistByType(&_VotingRegistry.CallOpts, poolType_, pool_)
}

// ListPoolsByProposer is a free data retrieval call binding the contract method 0x95c81529.
//
// Solidity: function listPoolsByProposer(address proposer_, uint256 offset_, uint256 limit_) view returns(address[] pools_)
func (_VotingRegistry *VotingRegistryCaller) ListPoolsByProposer(opts *bind.CallOpts, proposer_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "listPoolsByProposer", proposer_, offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListPoolsByProposer is a free data retrieval call binding the contract method 0x95c81529.
//
// Solidity: function listPoolsByProposer(address proposer_, uint256 offset_, uint256 limit_) view returns(address[] pools_)
func (_VotingRegistry *VotingRegistrySession) ListPoolsByProposer(proposer_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _VotingRegistry.Contract.ListPoolsByProposer(&_VotingRegistry.CallOpts, proposer_, offset_, limit_)
}

// ListPoolsByProposer is a free data retrieval call binding the contract method 0x95c81529.
//
// Solidity: function listPoolsByProposer(address proposer_, uint256 offset_, uint256 limit_) view returns(address[] pools_)
func (_VotingRegistry *VotingRegistryCallerSession) ListPoolsByProposer(proposer_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _VotingRegistry.Contract.ListPoolsByProposer(&_VotingRegistry.CallOpts, proposer_, offset_, limit_)
}

// ListPoolsByProposerAndType is a free data retrieval call binding the contract method 0xc4e09b63.
//
// Solidity: function listPoolsByProposerAndType(address proposer_, string poolType_, uint256 offset_, uint256 limit_) view returns(address[] pools_)
func (_VotingRegistry *VotingRegistryCaller) ListPoolsByProposerAndType(opts *bind.CallOpts, proposer_ common.Address, poolType_ string, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "listPoolsByProposerAndType", proposer_, poolType_, offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListPoolsByProposerAndType is a free data retrieval call binding the contract method 0xc4e09b63.
//
// Solidity: function listPoolsByProposerAndType(address proposer_, string poolType_, uint256 offset_, uint256 limit_) view returns(address[] pools_)
func (_VotingRegistry *VotingRegistrySession) ListPoolsByProposerAndType(proposer_ common.Address, poolType_ string, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _VotingRegistry.Contract.ListPoolsByProposerAndType(&_VotingRegistry.CallOpts, proposer_, poolType_, offset_, limit_)
}

// ListPoolsByProposerAndType is a free data retrieval call binding the contract method 0xc4e09b63.
//
// Solidity: function listPoolsByProposerAndType(address proposer_, string poolType_, uint256 offset_, uint256 limit_) view returns(address[] pools_)
func (_VotingRegistry *VotingRegistryCallerSession) ListPoolsByProposerAndType(proposer_ common.Address, poolType_ string, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _VotingRegistry.Contract.ListPoolsByProposerAndType(&_VotingRegistry.CallOpts, proposer_, poolType_, offset_, limit_)
}

// ListPoolsByType is a free data retrieval call binding the contract method 0x0eb64966.
//
// Solidity: function listPoolsByType(string poolType_, uint256 offset_, uint256 limit_) view returns(address[] pools_)
func (_VotingRegistry *VotingRegistryCaller) ListPoolsByType(opts *bind.CallOpts, poolType_ string, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "listPoolsByType", poolType_, offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListPoolsByType is a free data retrieval call binding the contract method 0x0eb64966.
//
// Solidity: function listPoolsByType(string poolType_, uint256 offset_, uint256 limit_) view returns(address[] pools_)
func (_VotingRegistry *VotingRegistrySession) ListPoolsByType(poolType_ string, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _VotingRegistry.Contract.ListPoolsByType(&_VotingRegistry.CallOpts, poolType_, offset_, limit_)
}

// ListPoolsByType is a free data retrieval call binding the contract method 0x0eb64966.
//
// Solidity: function listPoolsByType(string poolType_, uint256 offset_, uint256 limit_) view returns(address[] pools_)
func (_VotingRegistry *VotingRegistryCallerSession) ListPoolsByType(poolType_ string, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _VotingRegistry.Contract.ListPoolsByType(&_VotingRegistry.CallOpts, poolType_, offset_, limit_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotingRegistry *VotingRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotingRegistry *VotingRegistrySession) Owner() (common.Address, error) {
	return _VotingRegistry.Contract.Owner(&_VotingRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VotingRegistry *VotingRegistryCallerSession) Owner() (common.Address, error) {
	return _VotingRegistry.Contract.Owner(&_VotingRegistry.CallOpts)
}

// PoolCountByProposer is a free data retrieval call binding the contract method 0x74d3d2f7.
//
// Solidity: function poolCountByProposer(address proposer_) view returns(uint256)
func (_VotingRegistry *VotingRegistryCaller) PoolCountByProposer(opts *bind.CallOpts, proposer_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "poolCountByProposer", proposer_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCountByProposer is a free data retrieval call binding the contract method 0x74d3d2f7.
//
// Solidity: function poolCountByProposer(address proposer_) view returns(uint256)
func (_VotingRegistry *VotingRegistrySession) PoolCountByProposer(proposer_ common.Address) (*big.Int, error) {
	return _VotingRegistry.Contract.PoolCountByProposer(&_VotingRegistry.CallOpts, proposer_)
}

// PoolCountByProposer is a free data retrieval call binding the contract method 0x74d3d2f7.
//
// Solidity: function poolCountByProposer(address proposer_) view returns(uint256)
func (_VotingRegistry *VotingRegistryCallerSession) PoolCountByProposer(proposer_ common.Address) (*big.Int, error) {
	return _VotingRegistry.Contract.PoolCountByProposer(&_VotingRegistry.CallOpts, proposer_)
}

// PoolCountByProposerAndType is a free data retrieval call binding the contract method 0xd978e638.
//
// Solidity: function poolCountByProposerAndType(address proposer_, string poolType_) view returns(uint256)
func (_VotingRegistry *VotingRegistryCaller) PoolCountByProposerAndType(opts *bind.CallOpts, proposer_ common.Address, poolType_ string) (*big.Int, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "poolCountByProposerAndType", proposer_, poolType_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCountByProposerAndType is a free data retrieval call binding the contract method 0xd978e638.
//
// Solidity: function poolCountByProposerAndType(address proposer_, string poolType_) view returns(uint256)
func (_VotingRegistry *VotingRegistrySession) PoolCountByProposerAndType(proposer_ common.Address, poolType_ string) (*big.Int, error) {
	return _VotingRegistry.Contract.PoolCountByProposerAndType(&_VotingRegistry.CallOpts, proposer_, poolType_)
}

// PoolCountByProposerAndType is a free data retrieval call binding the contract method 0xd978e638.
//
// Solidity: function poolCountByProposerAndType(address proposer_, string poolType_) view returns(uint256)
func (_VotingRegistry *VotingRegistryCallerSession) PoolCountByProposerAndType(proposer_ common.Address, poolType_ string) (*big.Int, error) {
	return _VotingRegistry.Contract.PoolCountByProposerAndType(&_VotingRegistry.CallOpts, proposer_, poolType_)
}

// PoolCountByType is a free data retrieval call binding the contract method 0x2b1d3489.
//
// Solidity: function poolCountByType(string poolType_) view returns(uint256)
func (_VotingRegistry *VotingRegistryCaller) PoolCountByType(opts *bind.CallOpts, poolType_ string) (*big.Int, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "poolCountByType", poolType_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCountByType is a free data retrieval call binding the contract method 0x2b1d3489.
//
// Solidity: function poolCountByType(string poolType_) view returns(uint256)
func (_VotingRegistry *VotingRegistrySession) PoolCountByType(poolType_ string) (*big.Int, error) {
	return _VotingRegistry.Contract.PoolCountByType(&_VotingRegistry.CallOpts, poolType_)
}

// PoolCountByType is a free data retrieval call binding the contract method 0x2b1d3489.
//
// Solidity: function poolCountByType(string poolType_) view returns(uint256)
func (_VotingRegistry *VotingRegistryCallerSession) PoolCountByType(poolType_ string) (*big.Int, error) {
	return _VotingRegistry.Contract.PoolCountByType(&_VotingRegistry.CallOpts, poolType_)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_VotingRegistry *VotingRegistryCaller) PoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "poolFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_VotingRegistry *VotingRegistrySession) PoolFactory() (common.Address, error) {
	return _VotingRegistry.Contract.PoolFactory(&_VotingRegistry.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_VotingRegistry *VotingRegistryCallerSession) PoolFactory() (common.Address, error) {
	return _VotingRegistry.Contract.PoolFactory(&_VotingRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VotingRegistry *VotingRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VotingRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VotingRegistry *VotingRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _VotingRegistry.Contract.ProxiableUUID(&_VotingRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VotingRegistry *VotingRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VotingRegistry.Contract.ProxiableUUID(&_VotingRegistry.CallOpts)
}

// VotingRegistryInit is a paid mutator transaction binding the contract method 0x49f74dae.
//
// Solidity: function __VotingRegistry_init(address poolFactory_) returns()
func (_VotingRegistry *VotingRegistryTransactor) VotingRegistryInit(opts *bind.TransactOpts, poolFactory_ common.Address) (*types.Transaction, error) {
	return _VotingRegistry.contract.Transact(opts, "__VotingRegistry_init", poolFactory_)
}

// VotingRegistryInit is a paid mutator transaction binding the contract method 0x49f74dae.
//
// Solidity: function __VotingRegistry_init(address poolFactory_) returns()
func (_VotingRegistry *VotingRegistrySession) VotingRegistryInit(poolFactory_ common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.VotingRegistryInit(&_VotingRegistry.TransactOpts, poolFactory_)
}

// VotingRegistryInit is a paid mutator transaction binding the contract method 0x49f74dae.
//
// Solidity: function __VotingRegistry_init(address poolFactory_) returns()
func (_VotingRegistry *VotingRegistryTransactorSession) VotingRegistryInit(poolFactory_ common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.VotingRegistryInit(&_VotingRegistry.TransactOpts, poolFactory_)
}

// AddProxyPool is a paid mutator transaction binding the contract method 0xce41a736.
//
// Solidity: function addProxyPool(string poolType_, address proposer_, address pool_) returns()
func (_VotingRegistry *VotingRegistryTransactor) AddProxyPool(opts *bind.TransactOpts, poolType_ string, proposer_ common.Address, pool_ common.Address) (*types.Transaction, error) {
	return _VotingRegistry.contract.Transact(opts, "addProxyPool", poolType_, proposer_, pool_)
}

// AddProxyPool is a paid mutator transaction binding the contract method 0xce41a736.
//
// Solidity: function addProxyPool(string poolType_, address proposer_, address pool_) returns()
func (_VotingRegistry *VotingRegistrySession) AddProxyPool(poolType_ string, proposer_ common.Address, pool_ common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.AddProxyPool(&_VotingRegistry.TransactOpts, poolType_, proposer_, pool_)
}

// AddProxyPool is a paid mutator transaction binding the contract method 0xce41a736.
//
// Solidity: function addProxyPool(string poolType_, address proposer_, address pool_) returns()
func (_VotingRegistry *VotingRegistryTransactorSession) AddProxyPool(poolType_ string, proposer_ common.Address, pool_ common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.AddProxyPool(&_VotingRegistry.TransactOpts, poolType_, proposer_, pool_)
}

// BindVotingToRegistration is a paid mutator transaction binding the contract method 0x5f4999b8.
//
// Solidity: function bindVotingToRegistration(address proposer_, address voting_, address registration_) returns()
func (_VotingRegistry *VotingRegistryTransactor) BindVotingToRegistration(opts *bind.TransactOpts, proposer_ common.Address, voting_ common.Address, registration_ common.Address) (*types.Transaction, error) {
	return _VotingRegistry.contract.Transact(opts, "bindVotingToRegistration", proposer_, voting_, registration_)
}

// BindVotingToRegistration is a paid mutator transaction binding the contract method 0x5f4999b8.
//
// Solidity: function bindVotingToRegistration(address proposer_, address voting_, address registration_) returns()
func (_VotingRegistry *VotingRegistrySession) BindVotingToRegistration(proposer_ common.Address, voting_ common.Address, registration_ common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.BindVotingToRegistration(&_VotingRegistry.TransactOpts, proposer_, voting_, registration_)
}

// BindVotingToRegistration is a paid mutator transaction binding the contract method 0x5f4999b8.
//
// Solidity: function bindVotingToRegistration(address proposer_, address voting_, address registration_) returns()
func (_VotingRegistry *VotingRegistryTransactorSession) BindVotingToRegistration(proposer_ common.Address, voting_ common.Address, registration_ common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.BindVotingToRegistration(&_VotingRegistry.TransactOpts, proposer_, voting_, registration_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VotingRegistry *VotingRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VotingRegistry *VotingRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _VotingRegistry.Contract.RenounceOwnership(&_VotingRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VotingRegistry *VotingRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VotingRegistry.Contract.RenounceOwnership(&_VotingRegistry.TransactOpts)
}

// SetNewImplementations is a paid mutator transaction binding the contract method 0x05c05408.
//
// Solidity: function setNewImplementations(string[] poolTypes_, address[] newImplementations_) returns()
func (_VotingRegistry *VotingRegistryTransactor) SetNewImplementations(opts *bind.TransactOpts, poolTypes_ []string, newImplementations_ []common.Address) (*types.Transaction, error) {
	return _VotingRegistry.contract.Transact(opts, "setNewImplementations", poolTypes_, newImplementations_)
}

// SetNewImplementations is a paid mutator transaction binding the contract method 0x05c05408.
//
// Solidity: function setNewImplementations(string[] poolTypes_, address[] newImplementations_) returns()
func (_VotingRegistry *VotingRegistrySession) SetNewImplementations(poolTypes_ []string, newImplementations_ []common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.SetNewImplementations(&_VotingRegistry.TransactOpts, poolTypes_, newImplementations_)
}

// SetNewImplementations is a paid mutator transaction binding the contract method 0x05c05408.
//
// Solidity: function setNewImplementations(string[] poolTypes_, address[] newImplementations_) returns()
func (_VotingRegistry *VotingRegistryTransactorSession) SetNewImplementations(poolTypes_ []string, newImplementations_ []common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.SetNewImplementations(&_VotingRegistry.TransactOpts, poolTypes_, newImplementations_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotingRegistry *VotingRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VotingRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotingRegistry *VotingRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.TransferOwnership(&_VotingRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VotingRegistry *VotingRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.TransferOwnership(&_VotingRegistry.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VotingRegistry *VotingRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _VotingRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VotingRegistry *VotingRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.UpgradeTo(&_VotingRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VotingRegistry *VotingRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VotingRegistry.Contract.UpgradeTo(&_VotingRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VotingRegistry *VotingRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VotingRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VotingRegistry *VotingRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VotingRegistry.Contract.UpgradeToAndCall(&_VotingRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VotingRegistry *VotingRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VotingRegistry.Contract.UpgradeToAndCall(&_VotingRegistry.TransactOpts, newImplementation, data)
}

// VotingRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the VotingRegistry contract.
type VotingRegistryAdminChangedIterator struct {
	Event *VotingRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *VotingRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingRegistryAdminChanged)
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
		it.Event = new(VotingRegistryAdminChanged)
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
func (it *VotingRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingRegistryAdminChanged represents a AdminChanged event raised by the VotingRegistry contract.
type VotingRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VotingRegistry *VotingRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VotingRegistryAdminChangedIterator, error) {

	logs, sub, err := _VotingRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VotingRegistryAdminChangedIterator{contract: _VotingRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VotingRegistry *VotingRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VotingRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _VotingRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingRegistryAdminChanged)
				if err := _VotingRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_VotingRegistry *VotingRegistryFilterer) ParseAdminChanged(log types.Log) (*VotingRegistryAdminChanged, error) {
	event := new(VotingRegistryAdminChanged)
	if err := _VotingRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the VotingRegistry contract.
type VotingRegistryBeaconUpgradedIterator struct {
	Event *VotingRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VotingRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingRegistryBeaconUpgraded)
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
		it.Event = new(VotingRegistryBeaconUpgraded)
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
func (it *VotingRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the VotingRegistry contract.
type VotingRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VotingRegistry *VotingRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VotingRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VotingRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VotingRegistryBeaconUpgradedIterator{contract: _VotingRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VotingRegistry *VotingRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VotingRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VotingRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingRegistryBeaconUpgraded)
				if err := _VotingRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_VotingRegistry *VotingRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*VotingRegistryBeaconUpgraded, error) {
	event := new(VotingRegistryBeaconUpgraded)
	if err := _VotingRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VotingRegistry contract.
type VotingRegistryInitializedIterator struct {
	Event *VotingRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *VotingRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingRegistryInitialized)
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
		it.Event = new(VotingRegistryInitialized)
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
func (it *VotingRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingRegistryInitialized represents a Initialized event raised by the VotingRegistry contract.
type VotingRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VotingRegistry *VotingRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*VotingRegistryInitializedIterator, error) {

	logs, sub, err := _VotingRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VotingRegistryInitializedIterator{contract: _VotingRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VotingRegistry *VotingRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VotingRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _VotingRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingRegistryInitialized)
				if err := _VotingRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VotingRegistry *VotingRegistryFilterer) ParseInitialized(log types.Log) (*VotingRegistryInitialized, error) {
	event := new(VotingRegistryInitialized)
	if err := _VotingRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VotingRegistry contract.
type VotingRegistryOwnershipTransferredIterator struct {
	Event *VotingRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VotingRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingRegistryOwnershipTransferred)
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
		it.Event = new(VotingRegistryOwnershipTransferred)
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
func (it *VotingRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the VotingRegistry contract.
type VotingRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotingRegistry *VotingRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VotingRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotingRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VotingRegistryOwnershipTransferredIterator{contract: _VotingRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VotingRegistry *VotingRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VotingRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VotingRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingRegistryOwnershipTransferred)
				if err := _VotingRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VotingRegistry *VotingRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*VotingRegistryOwnershipTransferred, error) {
	event := new(VotingRegistryOwnershipTransferred)
	if err := _VotingRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VotingRegistry contract.
type VotingRegistryUpgradedIterator struct {
	Event *VotingRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *VotingRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingRegistryUpgraded)
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
		it.Event = new(VotingRegistryUpgraded)
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
func (it *VotingRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingRegistryUpgraded represents a Upgraded event raised by the VotingRegistry contract.
type VotingRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VotingRegistry *VotingRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VotingRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VotingRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VotingRegistryUpgradedIterator{contract: _VotingRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VotingRegistry *VotingRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VotingRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VotingRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingRegistryUpgraded)
				if err := _VotingRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_VotingRegistry *VotingRegistryFilterer) ParseUpgraded(log types.Log) (*VotingRegistryUpgraded, error) {
	event := new(VotingRegistryUpgraded)
	if err := _VotingRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
