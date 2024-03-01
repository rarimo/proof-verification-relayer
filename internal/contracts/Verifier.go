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

// IBaseVerifierProveIdentityParams is an auto generated low-level Go binding around an user-defined struct.
type IBaseVerifierProveIdentityParams struct {
	StatesMerkleData ILightweightStateStatesMerkleData
	Inputs           []*big.Int
	A                [2]*big.Int
	B                [2][2]*big.Int
	C                [2]*big.Int
}

// IBaseVerifierTransitStateParams is an auto generated low-level Go binding around an user-defined struct.
type IBaseVerifierTransitStateParams struct {
	NewIdentitiesStatesRoot [32]byte
	GistData                ILightweightStateGistRootData
	Proof                   []byte
}

// ILightweightStateGistRootData is an auto generated low-level Go binding around an user-defined struct.
type ILightweightStateGistRootData struct {
	Root               *big.Int
	CreatedAtTimestamp *big.Int
}

// ILightweightStateStatesMerkleData is an auto generated low-level Go binding around an user-defined struct.
type ILightweightStateStatesMerkleData struct {
	IssuerId           *big.Int
	IssuerState        *big.Int
	CreatedAtTimestamp *big.Int
	MerkleProof        [][32]byte
}

// IRegisterVerifierRegisterProofParams is an auto generated low-level Go binding around an user-defined struct.
type IRegisterVerifierRegisterProofParams struct {
	IssuingAuthority  *big.Int
	DocumentNullifier *big.Int
	Commitment        [32]byte
}

// IRegistrationRegistrationCounters is an auto generated low-level Go binding around an user-defined struct.
type IRegistrationRegistrationCounters struct {
	TotalRegistrations *big.Int
}

// IRegistrationRegistrationInfo is an auto generated low-level Go binding around an user-defined struct.
type IRegistrationRegistrationInfo struct {
	Remark   string
	Values   IRegistrationRegistrationValues
	Counters IRegistrationRegistrationCounters
}

// IRegistrationRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IRegistrationRegistrationParams struct {
	Remark           string
	CommitmentStart  *big.Int
	CommitmentPeriod *big.Int
}

// IRegistrationRegistrationValues is an auto generated low-level Go binding around an user-defined struct.
type IRegistrationRegistrationValues struct {
	CommitmentStartTime *big.Int
	CommitmentEndTime   *big.Int
}

// SparseMerkleTreeNode is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeNode struct {
	NodeType   uint8
	ChildLeft  uint64
	ChildRight uint64
	NodeHash   [32]byte
	Key        [32]byte
	Value      [32]byte
}

// SparseMerkleTreeProof is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeProof struct {
	Root         [32]byte
	Siblings     [][32]byte
	Existence    bool
	Key          [32]byte
	Value        [32]byte
	AuxExistence bool
	AuxKey       [32]byte
	AuxValue     [32]byte
}

// VerifierMetaData contains all meta data concerning the Verifier contract.
var VerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registerVerifier_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"treeHeight_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"commitmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitmentPeriod\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIRegistration.RegistrationParams\",\"name\":\"registrationParams\",\"type\":\"tuple\"}],\"name\":\"RegistrationInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"issuerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"issuerState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structILightweightState.StatesMerkleData\",\"name\":\"statesMerkleData\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"indexed\":false,\"internalType\":\"structIBaseVerifier.ProveIdentityParams\",\"name\":\"proveIdentityParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"issuingAuthority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"documentNullifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structIRegisterVerifier.RegisterProofParams\",\"name\":\"registerProofParams\",\"type\":\"tuple\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"commitmentStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitmentPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistration.RegistrationParams\",\"name\":\"registrationParams_\",\"type\":\"tuple\"}],\"name\":\"__Registration_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getNodeByKey\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSparseMerkleTree.NodeType\",\"name\":\"nodeType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"childLeft\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"childRight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"auxKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auxValue\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistrationInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"commitmentStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitmentEndTime\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistration.RegistrationValues\",\"name\":\"values\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalRegistrations\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistration.RegistrationCounters\",\"name\":\"counters\",\"type\":\"tuple\"}],\"internalType\":\"structIRegistration.RegistrationInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistrationStatus\",\"outputs\":[{\"internalType\":\"enumIRegistration.RegistrationStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"isRootExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"issuerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"issuerState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structILightweightState.StatesMerkleData\",\"name\":\"statesMerkleData\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"inputs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structIBaseVerifier.ProveIdentityParams\",\"name\":\"proveIdentityParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"issuingAuthority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"documentNullifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIRegisterVerifier.RegisterProofParams\",\"name\":\"registerProofParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"newIdentitiesStatesRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structILightweightState.GistRootData\",\"name\":\"gistData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structIBaseVerifier.TransitStateParams\",\"name\":\"transitStateParams_\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isTransitState_\",\"type\":\"bool\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerVerifier\",\"outputs\":[{\"internalType\":\"contractIRegisterVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrationInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"commitmentStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitmentEndTime\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistration.RegistrationValues\",\"name\":\"values\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalRegistrations\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistration.RegistrationCounters\",\"name\":\"counters\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"rootsHistory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"smtTreeMaxDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierMetaData.ABI instead.
var VerifierABI = VerifierMetaData.ABI

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_Verifier *VerifierCaller) Commitments(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_Verifier *VerifierSession) Commitments(arg0 [32]byte) (bool, error) {
	return _Verifier.Contract.Commitments(&_Verifier.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(bool)
func (_Verifier *VerifierCallerSession) Commitments(arg0 [32]byte) (bool, error) {
	return _Verifier.Contract.Commitments(&_Verifier.CallOpts, arg0)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_Verifier *VerifierCaller) GetNodeByKey(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeNode, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "getNodeByKey", key_)

	if err != nil {
		return *new(SparseMerkleTreeNode), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeNode)).(*SparseMerkleTreeNode)

	return out0, err

}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_Verifier *VerifierSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _Verifier.Contract.GetNodeByKey(&_Verifier.CallOpts, key_)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_Verifier *VerifierCallerSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _Verifier.Contract.GetNodeByKey(&_Verifier.CallOpts, key_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_Verifier *VerifierCaller) GetProof(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeProof, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "getProof", key_)

	if err != nil {
		return *new(SparseMerkleTreeProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeProof)).(*SparseMerkleTreeProof)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_Verifier *VerifierSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _Verifier.Contract.GetProof(&_Verifier.CallOpts, key_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_Verifier *VerifierCallerSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _Verifier.Contract.GetProof(&_Verifier.CallOpts, key_)
}

// GetRegistrationInfo is a free data retrieval call binding the contract method 0x9f2ec56a.
//
// Solidity: function getRegistrationInfo() view returns((string,(uint256,uint256),(uint256)))
func (_Verifier *VerifierCaller) GetRegistrationInfo(opts *bind.CallOpts) (IRegistrationRegistrationInfo, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "getRegistrationInfo")

	if err != nil {
		return *new(IRegistrationRegistrationInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistrationRegistrationInfo)).(*IRegistrationRegistrationInfo)

	return out0, err

}

// GetRegistrationInfo is a free data retrieval call binding the contract method 0x9f2ec56a.
//
// Solidity: function getRegistrationInfo() view returns((string,(uint256,uint256),(uint256)))
func (_Verifier *VerifierSession) GetRegistrationInfo() (IRegistrationRegistrationInfo, error) {
	return _Verifier.Contract.GetRegistrationInfo(&_Verifier.CallOpts)
}

// GetRegistrationInfo is a free data retrieval call binding the contract method 0x9f2ec56a.
//
// Solidity: function getRegistrationInfo() view returns((string,(uint256,uint256),(uint256)))
func (_Verifier *VerifierCallerSession) GetRegistrationInfo() (IRegistrationRegistrationInfo, error) {
	return _Verifier.Contract.GetRegistrationInfo(&_Verifier.CallOpts)
}

// GetRegistrationStatus is a free data retrieval call binding the contract method 0xe35fb241.
//
// Solidity: function getRegistrationStatus() view returns(uint8)
func (_Verifier *VerifierCaller) GetRegistrationStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "getRegistrationStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetRegistrationStatus is a free data retrieval call binding the contract method 0xe35fb241.
//
// Solidity: function getRegistrationStatus() view returns(uint8)
func (_Verifier *VerifierSession) GetRegistrationStatus() (uint8, error) {
	return _Verifier.Contract.GetRegistrationStatus(&_Verifier.CallOpts)
}

// GetRegistrationStatus is a free data retrieval call binding the contract method 0xe35fb241.
//
// Solidity: function getRegistrationStatus() view returns(uint8)
func (_Verifier *VerifierCallerSession) GetRegistrationStatus() (uint8, error) {
	return _Verifier.Contract.GetRegistrationStatus(&_Verifier.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Verifier *VerifierCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Verifier *VerifierSession) GetRoot() ([32]byte, error) {
	return _Verifier.Contract.GetRoot(&_Verifier.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Verifier *VerifierCallerSession) GetRoot() ([32]byte, error) {
	return _Verifier.Contract.GetRoot(&_Verifier.CallOpts)
}

// IsRootExists is a free data retrieval call binding the contract method 0x74197dd5.
//
// Solidity: function isRootExists(bytes32 root) view returns(bool)
func (_Verifier *VerifierCaller) IsRootExists(opts *bind.CallOpts, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "isRootExists", root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRootExists is a free data retrieval call binding the contract method 0x74197dd5.
//
// Solidity: function isRootExists(bytes32 root) view returns(bool)
func (_Verifier *VerifierSession) IsRootExists(root [32]byte) (bool, error) {
	return _Verifier.Contract.IsRootExists(&_Verifier.CallOpts, root)
}

// IsRootExists is a free data retrieval call binding the contract method 0x74197dd5.
//
// Solidity: function isRootExists(bytes32 root) view returns(bool)
func (_Verifier *VerifierCallerSession) IsRootExists(root [32]byte) (bool, error) {
	return _Verifier.Contract.IsRootExists(&_Verifier.CallOpts, root)
}

// RegisterVerifier is a free data retrieval call binding the contract method 0xea968a2d.
//
// Solidity: function registerVerifier() view returns(address)
func (_Verifier *VerifierCaller) RegisterVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "registerVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegisterVerifier is a free data retrieval call binding the contract method 0xea968a2d.
//
// Solidity: function registerVerifier() view returns(address)
func (_Verifier *VerifierSession) RegisterVerifier() (common.Address, error) {
	return _Verifier.Contract.RegisterVerifier(&_Verifier.CallOpts)
}

// RegisterVerifier is a free data retrieval call binding the contract method 0xea968a2d.
//
// Solidity: function registerVerifier() view returns(address)
func (_Verifier *VerifierCallerSession) RegisterVerifier() (common.Address, error) {
	return _Verifier.Contract.RegisterVerifier(&_Verifier.CallOpts)
}

// RegistrationInfo is a free data retrieval call binding the contract method 0xeb9f4937.
//
// Solidity: function registrationInfo() view returns(string remark, (uint256,uint256) values, (uint256) counters)
func (_Verifier *VerifierCaller) RegistrationInfo(opts *bind.CallOpts) (struct {
	Remark   string
	Values   IRegistrationRegistrationValues
	Counters IRegistrationRegistrationCounters
}, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "registrationInfo")

	outstruct := new(struct {
		Remark   string
		Values   IRegistrationRegistrationValues
		Counters IRegistrationRegistrationCounters
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Remark = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Values = *abi.ConvertType(out[1], new(IRegistrationRegistrationValues)).(*IRegistrationRegistrationValues)
	outstruct.Counters = *abi.ConvertType(out[2], new(IRegistrationRegistrationCounters)).(*IRegistrationRegistrationCounters)

	return *outstruct, err

}

// RegistrationInfo is a free data retrieval call binding the contract method 0xeb9f4937.
//
// Solidity: function registrationInfo() view returns(string remark, (uint256,uint256) values, (uint256) counters)
func (_Verifier *VerifierSession) RegistrationInfo() (struct {
	Remark   string
	Values   IRegistrationRegistrationValues
	Counters IRegistrationRegistrationCounters
}, error) {
	return _Verifier.Contract.RegistrationInfo(&_Verifier.CallOpts)
}

// RegistrationInfo is a free data retrieval call binding the contract method 0xeb9f4937.
//
// Solidity: function registrationInfo() view returns(string remark, (uint256,uint256) values, (uint256) counters)
func (_Verifier *VerifierCallerSession) RegistrationInfo() (struct {
	Remark   string
	Values   IRegistrationRegistrationValues
	Counters IRegistrationRegistrationCounters
}, error) {
	return _Verifier.Contract.RegistrationInfo(&_Verifier.CallOpts)
}

// RootsHistory is a free data retrieval call binding the contract method 0x241e2cfe.
//
// Solidity: function rootsHistory(bytes32 ) view returns(bool)
func (_Verifier *VerifierCaller) RootsHistory(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "rootsHistory", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RootsHistory is a free data retrieval call binding the contract method 0x241e2cfe.
//
// Solidity: function rootsHistory(bytes32 ) view returns(bool)
func (_Verifier *VerifierSession) RootsHistory(arg0 [32]byte) (bool, error) {
	return _Verifier.Contract.RootsHistory(&_Verifier.CallOpts, arg0)
}

// RootsHistory is a free data retrieval call binding the contract method 0x241e2cfe.
//
// Solidity: function rootsHistory(bytes32 ) view returns(bool)
func (_Verifier *VerifierCallerSession) RootsHistory(arg0 [32]byte) (bool, error) {
	return _Verifier.Contract.RootsHistory(&_Verifier.CallOpts, arg0)
}

// SmtTreeMaxDepth is a free data retrieval call binding the contract method 0xf975aa00.
//
// Solidity: function smtTreeMaxDepth() view returns(uint256)
func (_Verifier *VerifierCaller) SmtTreeMaxDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "smtTreeMaxDepth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SmtTreeMaxDepth is a free data retrieval call binding the contract method 0xf975aa00.
//
// Solidity: function smtTreeMaxDepth() view returns(uint256)
func (_Verifier *VerifierSession) SmtTreeMaxDepth() (*big.Int, error) {
	return _Verifier.Contract.SmtTreeMaxDepth(&_Verifier.CallOpts)
}

// SmtTreeMaxDepth is a free data retrieval call binding the contract method 0xf975aa00.
//
// Solidity: function smtTreeMaxDepth() view returns(uint256)
func (_Verifier *VerifierCallerSession) SmtTreeMaxDepth() (*big.Int, error) {
	return _Verifier.Contract.SmtTreeMaxDepth(&_Verifier.CallOpts)
}

// RegistrationInit is a paid mutator transaction binding the contract method 0x2a8f67d5.
//
// Solidity: function __Registration_init((string,uint256,uint256) registrationParams_) returns()
func (_Verifier *VerifierTransactor) RegistrationInit(opts *bind.TransactOpts, registrationParams_ IRegistrationRegistrationParams) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "__Registration_init", registrationParams_)
}

// RegistrationInit is a paid mutator transaction binding the contract method 0x2a8f67d5.
//
// Solidity: function __Registration_init((string,uint256,uint256) registrationParams_) returns()
func (_Verifier *VerifierSession) RegistrationInit(registrationParams_ IRegistrationRegistrationParams) (*types.Transaction, error) {
	return _Verifier.Contract.RegistrationInit(&_Verifier.TransactOpts, registrationParams_)
}

// RegistrationInit is a paid mutator transaction binding the contract method 0x2a8f67d5.
//
// Solidity: function __Registration_init((string,uint256,uint256) registrationParams_) returns()
func (_Verifier *VerifierTransactorSession) RegistrationInit(registrationParams_ IRegistrationRegistrationParams) (*types.Transaction, error) {
	return _Verifier.Contract.RegistrationInit(&_Verifier.TransactOpts, registrationParams_)
}

// Register is a paid mutator transaction binding the contract method 0x24a4831c.
//
// Solidity: function register(((uint256,uint256,uint256,bytes32[]),uint256[],uint256[2],uint256[2][2],uint256[2]) proveIdentityParams_, (uint256,uint256,bytes32) registerProofParams_, (bytes32,(uint256,uint256),bytes) transitStateParams_, bool isTransitState_) returns()
func (_Verifier *VerifierTransactor) Register(opts *bind.TransactOpts, proveIdentityParams_ IBaseVerifierProveIdentityParams, registerProofParams_ IRegisterVerifierRegisterProofParams, transitStateParams_ IBaseVerifierTransitStateParams, isTransitState_ bool) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "register", proveIdentityParams_, registerProofParams_, transitStateParams_, isTransitState_)
}

// Register is a paid mutator transaction binding the contract method 0x24a4831c.
//
// Solidity: function register(((uint256,uint256,uint256,bytes32[]),uint256[],uint256[2],uint256[2][2],uint256[2]) proveIdentityParams_, (uint256,uint256,bytes32) registerProofParams_, (bytes32,(uint256,uint256),bytes) transitStateParams_, bool isTransitState_) returns()
func (_Verifier *VerifierSession) Register(proveIdentityParams_ IBaseVerifierProveIdentityParams, registerProofParams_ IRegisterVerifierRegisterProofParams, transitStateParams_ IBaseVerifierTransitStateParams, isTransitState_ bool) (*types.Transaction, error) {
	return _Verifier.Contract.Register(&_Verifier.TransactOpts, proveIdentityParams_, registerProofParams_, transitStateParams_, isTransitState_)
}

// Register is a paid mutator transaction binding the contract method 0x24a4831c.
//
// Solidity: function register(((uint256,uint256,uint256,bytes32[]),uint256[],uint256[2],uint256[2][2],uint256[2]) proveIdentityParams_, (uint256,uint256,bytes32) registerProofParams_, (bytes32,(uint256,uint256),bytes) transitStateParams_, bool isTransitState_) returns()
func (_Verifier *VerifierTransactorSession) Register(proveIdentityParams_ IBaseVerifierProveIdentityParams, registerProofParams_ IRegisterVerifierRegisterProofParams, transitStateParams_ IBaseVerifierTransitStateParams, isTransitState_ bool) (*types.Transaction, error) {
	return _Verifier.Contract.Register(&_Verifier.TransactOpts, proveIdentityParams_, registerProofParams_, transitStateParams_, isTransitState_)
}

// VerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Verifier contract.
type VerifierInitializedIterator struct {
	Event *VerifierInitialized // Event containing the contract specifics and raw log

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
func (it *VerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierInitialized)
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
		it.Event = new(VerifierInitialized)
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
func (it *VerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifierInitialized represents a Initialized event raised by the Verifier contract.
type VerifierInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Verifier *VerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*VerifierInitializedIterator, error) {

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VerifierInitializedIterator{contract: _Verifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Verifier *VerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifierInitialized)
				if err := _Verifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Verifier *VerifierFilterer) ParseInitialized(log types.Log) (*VerifierInitialized, error) {
	event := new(VerifierInitialized)
	if err := _Verifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifierRegistrationInitializedIterator is returned from FilterRegistrationInitialized and is used to iterate over the raw logs and unpacked data for RegistrationInitialized events raised by the Verifier contract.
type VerifierRegistrationInitializedIterator struct {
	Event *VerifierRegistrationInitialized // Event containing the contract specifics and raw log

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
func (it *VerifierRegistrationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierRegistrationInitialized)
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
		it.Event = new(VerifierRegistrationInitialized)
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
func (it *VerifierRegistrationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifierRegistrationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifierRegistrationInitialized represents a RegistrationInitialized event raised by the Verifier contract.
type VerifierRegistrationInitialized struct {
	Proposer           common.Address
	RegistrationParams IRegistrationRegistrationParams
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegistrationInitialized is a free log retrieval operation binding the contract event 0xe5360f24c53a188e31370cfdddca9192dd7cef8a3f7e9e76e1a32f823bc01998.
//
// Solidity: event RegistrationInitialized(address indexed proposer, (string,uint256,uint256) registrationParams)
func (_Verifier *VerifierFilterer) FilterRegistrationInitialized(opts *bind.FilterOpts, proposer []common.Address) (*VerifierRegistrationInitializedIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "RegistrationInitialized", proposerRule)
	if err != nil {
		return nil, err
	}
	return &VerifierRegistrationInitializedIterator{contract: _Verifier.contract, event: "RegistrationInitialized", logs: logs, sub: sub}, nil
}

// WatchRegistrationInitialized is a free log subscription operation binding the contract event 0xe5360f24c53a188e31370cfdddca9192dd7cef8a3f7e9e76e1a32f823bc01998.
//
// Solidity: event RegistrationInitialized(address indexed proposer, (string,uint256,uint256) registrationParams)
func (_Verifier *VerifierFilterer) WatchRegistrationInitialized(opts *bind.WatchOpts, sink chan<- *VerifierRegistrationInitialized, proposer []common.Address) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "RegistrationInitialized", proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifierRegistrationInitialized)
				if err := _Verifier.contract.UnpackLog(event, "RegistrationInitialized", log); err != nil {
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

// ParseRegistrationInitialized is a log parse operation binding the contract event 0xe5360f24c53a188e31370cfdddca9192dd7cef8a3f7e9e76e1a32f823bc01998.
//
// Solidity: event RegistrationInitialized(address indexed proposer, (string,uint256,uint256) registrationParams)
func (_Verifier *VerifierFilterer) ParseRegistrationInitialized(log types.Log) (*VerifierRegistrationInitialized, error) {
	event := new(VerifierRegistrationInitialized)
	if err := _Verifier.contract.UnpackLog(event, "RegistrationInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifierUserRegisteredIterator is returned from FilterUserRegistered and is used to iterate over the raw logs and unpacked data for UserRegistered events raised by the Verifier contract.
type VerifierUserRegisteredIterator struct {
	Event *VerifierUserRegistered // Event containing the contract specifics and raw log

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
func (it *VerifierUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierUserRegistered)
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
		it.Event = new(VerifierUserRegistered)
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
func (it *VerifierUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifierUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifierUserRegistered represents a UserRegistered event raised by the Verifier contract.
type VerifierUserRegistered struct {
	User                common.Address
	ProveIdentityParams IBaseVerifierProveIdentityParams
	RegisterProofParams IRegisterVerifierRegisterProofParams
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUserRegistered is a free log retrieval operation binding the contract event 0xb10c3702eaaf48572daf444c0bd6dd5e6746695f86b6cf3b73130dc94fcbcbba.
//
// Solidity: event UserRegistered(address indexed user, ((uint256,uint256,uint256,bytes32[]),uint256[],uint256[2],uint256[2][2],uint256[2]) proveIdentityParams, (uint256,uint256,bytes32) registerProofParams)
func (_Verifier *VerifierFilterer) FilterUserRegistered(opts *bind.FilterOpts, user []common.Address) (*VerifierUserRegisteredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "UserRegistered", userRule)
	if err != nil {
		return nil, err
	}
	return &VerifierUserRegisteredIterator{contract: _Verifier.contract, event: "UserRegistered", logs: logs, sub: sub}, nil
}

// WatchUserRegistered is a free log subscription operation binding the contract event 0xb10c3702eaaf48572daf444c0bd6dd5e6746695f86b6cf3b73130dc94fcbcbba.
//
// Solidity: event UserRegistered(address indexed user, ((uint256,uint256,uint256,bytes32[]),uint256[],uint256[2],uint256[2][2],uint256[2]) proveIdentityParams, (uint256,uint256,bytes32) registerProofParams)
func (_Verifier *VerifierFilterer) WatchUserRegistered(opts *bind.WatchOpts, sink chan<- *VerifierUserRegistered, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "UserRegistered", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifierUserRegistered)
				if err := _Verifier.contract.UnpackLog(event, "UserRegistered", log); err != nil {
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

// ParseUserRegistered is a log parse operation binding the contract event 0xb10c3702eaaf48572daf444c0bd6dd5e6746695f86b6cf3b73130dc94fcbcbba.
//
// Solidity: event UserRegistered(address indexed user, ((uint256,uint256,uint256,bytes32[]),uint256[],uint256[2],uint256[2][2],uint256[2]) proveIdentityParams, (uint256,uint256,bytes32) registerProofParams)
func (_Verifier *VerifierFilterer) ParseUserRegistered(log types.Log) (*VerifierUserRegistered, error) {
	event := new(VerifierUserRegistered)
	if err := _Verifier.contract.UnpackLog(event, "UserRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
