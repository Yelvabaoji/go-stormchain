// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chain3go

import (
	"github.com/filestorm/go-filestorm/moac/moac-lib/common"
	"github.com/filestorm/go-filestorm/moac/moac-vnode/accounts/abi"
	"github.com/filestorm/go-filestorm/moac/moac-vnode/accounts/abi/bind"
	"github.com/filestorm/go-filestorm/moac/moac-vnode/core/types"
	"strings"
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"items\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setItem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_version\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"ItemSet\",\"type\":\"event\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
const StoreBin = `608060405234801561001057600080fd5b5060405161033438038061033483398101604052805101805161003a906000906020840190610041565b50506100dc565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061008257805160ff19168380011785556100af565b828001600101855582156100af579182015b828111156100af578251825591602001919060010190610094565b506100bb9291506100bf565b5090565b6100d991905b808211156100bb57600081556001016100c5565b90565b610249806100eb6000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166348f343f3811461005b57806354fd4d5014610085578063f56256c71461010f575b600080fd5b34801561006757600080fd5b5061007360043561012c565b60408051918252519081900360200190f35b34801561009157600080fd5b5061009a61013e565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100d45781810151838201526020016100bc565b50505050905090810190601f1680156101015780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561011b57600080fd5b5061012a6004356024356101cc565b005b60016020526000908152604090205481565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101c45780601f10610199576101008083540402835291602001916101c4565b820191906000526020600020905b8154815290600101906020018083116101a757829003601f168201915b505050505081565b600082815260016020908152604091829020839055815184815290810183905281517fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4929181900390910190a150505600a165627a7a7230582002e2d39d1250dddad50012a6c9904f851a329e8925e8aeb232151ee1963defda0029`

// DeployStore deploys a new MoacNode contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _version string) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}}, nil
}

// Store is an auto generated Go binding around an MoacNode contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
}

// StoreCaller is an auto generated read-only Go binding around an MoacNode contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an MoacNode contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an MoacNode contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an MoacNode contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an MoacNode contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an MoacNode contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an MoacNode contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an MoacNode contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items( bytes32) constant returns(bytes32)
func (_Store *StoreCaller) Items(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "items", arg0)
	return *ret0, err
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items( bytes32) constant returns(bytes32)
func (_Store *StoreSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items( bytes32) constant returns(bytes32)
func (_Store *StoreCallerSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Store *StoreCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Store *StoreSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Store *StoreCallerSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(key bytes32, value bytes32) returns()
func (_Store *StoreTransactor) SetItem(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(key bytes32, value bytes32) returns()
func (_Store *StoreSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(key bytes32, value bytes32) returns()
func (_Store *StoreTransactorSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}
