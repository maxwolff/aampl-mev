// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arb

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

// ArbMetaData contains all meta data concerning the Arb contract.
var ArbMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetRate\",\"type\":\"uint256\"}],\"name\":\"computeSupplyDelta\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRebasePercentage\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lending_pool\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ArbABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbMetaData.ABI instead.
var ArbABI = ArbMetaData.ABI

// Arb is an auto generated Go binding around an Ethereum contract.
type Arb struct {
	ArbCaller     // Read-only binding to the contract
	ArbTransactor // Write-only binding to the contract
	ArbFilterer   // Log filterer for contract events
}

// ArbCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbSession struct {
	Contract     *Arb              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbCallerSession struct {
	Contract *ArbCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbTransactorSession struct {
	Contract     *ArbTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbRaw struct {
	Contract *Arb // Generic contract binding to access the raw methods on
}

// ArbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbCallerRaw struct {
	Contract *ArbCaller // Generic read-only contract binding to access the raw methods on
}

// ArbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbTransactorRaw struct {
	Contract *ArbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArb creates a new instance of Arb, bound to a specific deployed contract.
func NewArb(address common.Address, backend bind.ContractBackend) (*Arb, error) {
	contract, err := bindArb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arb{ArbCaller: ArbCaller{contract: contract}, ArbTransactor: ArbTransactor{contract: contract}, ArbFilterer: ArbFilterer{contract: contract}}, nil
}

// NewArbCaller creates a new read-only instance of Arb, bound to a specific deployed contract.
func NewArbCaller(address common.Address, caller bind.ContractCaller) (*ArbCaller, error) {
	contract, err := bindArb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbCaller{contract: contract}, nil
}

// NewArbTransactor creates a new write-only instance of Arb, bound to a specific deployed contract.
func NewArbTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbTransactor, error) {
	contract, err := bindArb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbTransactor{contract: contract}, nil
}

// NewArbFilterer creates a new log filterer instance of Arb, bound to a specific deployed contract.
func NewArbFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbFilterer, error) {
	contract, err := bindArb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbFilterer{contract: contract}, nil
}

// bindArb binds a generic wrapper to an already deployed contract.
func bindArb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arb *ArbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arb.Contract.ArbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arb *ArbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.Contract.ArbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arb *ArbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arb.Contract.ArbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arb *ArbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arb *ArbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arb *ArbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arb.Contract.contract.Transact(opts, method, params...)
}

// ComputeSupplyDelta is a free data retrieval call binding the contract method 0xf2d0bd56.
//
// Solidity: function computeSupplyDelta(uint256 rate, uint256 targetRate) view returns(int256)
func (_Arb *ArbCaller) ComputeSupplyDelta(opts *bind.CallOpts, rate *big.Int, targetRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "computeSupplyDelta", rate, targetRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeSupplyDelta is a free data retrieval call binding the contract method 0xf2d0bd56.
//
// Solidity: function computeSupplyDelta(uint256 rate, uint256 targetRate) view returns(int256)
func (_Arb *ArbSession) ComputeSupplyDelta(rate *big.Int, targetRate *big.Int) (*big.Int, error) {
	return _Arb.Contract.ComputeSupplyDelta(&_Arb.CallOpts, rate, targetRate)
}

// ComputeSupplyDelta is a free data retrieval call binding the contract method 0xf2d0bd56.
//
// Solidity: function computeSupplyDelta(uint256 rate, uint256 targetRate) view returns(int256)
func (_Arb *ArbCallerSession) ComputeSupplyDelta(rate *big.Int, targetRate *big.Int) (*big.Int, error) {
	return _Arb.Contract.ComputeSupplyDelta(&_Arb.CallOpts, rate, targetRate)
}

// LendingPool is a free data retrieval call binding the contract method 0xf984ead8.
//
// Solidity: function lending_pool() view returns(address)
func (_Arb *ArbCaller) LendingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "lending_pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LendingPool is a free data retrieval call binding the contract method 0xf984ead8.
//
// Solidity: function lending_pool() view returns(address)
func (_Arb *ArbSession) LendingPool() (common.Address, error) {
	return _Arb.Contract.LendingPool(&_Arb.CallOpts)
}

// LendingPool is a free data retrieval call binding the contract method 0xf984ead8.
//
// Solidity: function lending_pool() view returns(address)
func (_Arb *ArbCallerSession) LendingPool() (common.Address, error) {
	return _Arb.Contract.LendingPool(&_Arb.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Arb *ArbCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Arb *ArbSession) Weth() (common.Address, error) {
	return _Arb.Contract.Weth(&_Arb.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Arb *ArbCallerSession) Weth() (common.Address, error) {
	return _Arb.Contract.Weth(&_Arb.CallOpts)
}

// GetRebasePercentage is a paid mutator transaction binding the contract method 0xc4f16c2f.
//
// Solidity: function getRebasePercentage() returns(int256, uint256)
func (_Arb *ArbTransactor) GetRebasePercentage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "getRebasePercentage")
}

// GetRebasePercentage is a paid mutator transaction binding the contract method 0xc4f16c2f.
//
// Solidity: function getRebasePercentage() returns(int256, uint256)
func (_Arb *ArbSession) GetRebasePercentage() (*types.Transaction, error) {
	return _Arb.Contract.GetRebasePercentage(&_Arb.TransactOpts)
}

// GetRebasePercentage is a paid mutator transaction binding the contract method 0xc4f16c2f.
//
// Solidity: function getRebasePercentage() returns(int256, uint256)
func (_Arb *ArbTransactorSession) GetRebasePercentage() (*types.Transaction, error) {
	return _Arb.Contract.GetRebasePercentage(&_Arb.TransactOpts)
}
