// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iuniswapv2

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

// Iuniswapv2MetaData contains all meta data concerning the Iuniswapv2 contract.
var Iuniswapv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Iuniswapv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Iuniswapv2MetaData.ABI instead.
var Iuniswapv2ABI = Iuniswapv2MetaData.ABI

// Iuniswapv2 is an auto generated Go binding around an Ethereum contract.
type Iuniswapv2 struct {
	Iuniswapv2Caller     // Read-only binding to the contract
	Iuniswapv2Transactor // Write-only binding to the contract
	Iuniswapv2Filterer   // Log filterer for contract events
}

// Iuniswapv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Iuniswapv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iuniswapv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Iuniswapv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iuniswapv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Iuniswapv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iuniswapv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Iuniswapv2Session struct {
	Contract     *Iuniswapv2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Iuniswapv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Iuniswapv2CallerSession struct {
	Contract *Iuniswapv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Iuniswapv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Iuniswapv2TransactorSession struct {
	Contract     *Iuniswapv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Iuniswapv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Iuniswapv2Raw struct {
	Contract *Iuniswapv2 // Generic contract binding to access the raw methods on
}

// Iuniswapv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Iuniswapv2CallerRaw struct {
	Contract *Iuniswapv2Caller // Generic read-only contract binding to access the raw methods on
}

// Iuniswapv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Iuniswapv2TransactorRaw struct {
	Contract *Iuniswapv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIuniswapv2 creates a new instance of Iuniswapv2, bound to a specific deployed contract.
func NewIuniswapv2(address common.Address, backend bind.ContractBackend) (*Iuniswapv2, error) {
	contract, err := bindIuniswapv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2{Iuniswapv2Caller: Iuniswapv2Caller{contract: contract}, Iuniswapv2Transactor: Iuniswapv2Transactor{contract: contract}, Iuniswapv2Filterer: Iuniswapv2Filterer{contract: contract}}, nil
}

// NewIuniswapv2Caller creates a new read-only instance of Iuniswapv2, bound to a specific deployed contract.
func NewIuniswapv2Caller(address common.Address, caller bind.ContractCaller) (*Iuniswapv2Caller, error) {
	contract, err := bindIuniswapv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2Caller{contract: contract}, nil
}

// NewIuniswapv2Transactor creates a new write-only instance of Iuniswapv2, bound to a specific deployed contract.
func NewIuniswapv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Iuniswapv2Transactor, error) {
	contract, err := bindIuniswapv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2Transactor{contract: contract}, nil
}

// NewIuniswapv2Filterer creates a new log filterer instance of Iuniswapv2, bound to a specific deployed contract.
func NewIuniswapv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Iuniswapv2Filterer, error) {
	contract, err := bindIuniswapv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Iuniswapv2Filterer{contract: contract}, nil
}

// bindIuniswapv2 binds a generic wrapper to an already deployed contract.
func bindIuniswapv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Iuniswapv2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iuniswapv2 *Iuniswapv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iuniswapv2.Contract.Iuniswapv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iuniswapv2 *Iuniswapv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iuniswapv2.Contract.Iuniswapv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iuniswapv2 *Iuniswapv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iuniswapv2.Contract.Iuniswapv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iuniswapv2 *Iuniswapv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iuniswapv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iuniswapv2 *Iuniswapv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iuniswapv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iuniswapv2 *Iuniswapv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iuniswapv2.Contract.contract.Transact(opts, method, params...)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Iuniswapv2 *Iuniswapv2Transactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Iuniswapv2 *Iuniswapv2Session) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2.Contract.SwapExactTokensForETH(&_Iuniswapv2.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Iuniswapv2 *Iuniswapv2TransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Iuniswapv2.Contract.SwapExactTokensForETH(&_Iuniswapv2.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}
