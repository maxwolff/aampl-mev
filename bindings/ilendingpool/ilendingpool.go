// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ilendingpool

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

// IlendingpoolMetaData contains all meta data concerning the Ilendingpool contract.
var IlendingpoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IlendingpoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IlendingpoolMetaData.ABI instead.
var IlendingpoolABI = IlendingpoolMetaData.ABI

// Ilendingpool is an auto generated Go binding around an Ethereum contract.
type Ilendingpool struct {
	IlendingpoolCaller     // Read-only binding to the contract
	IlendingpoolTransactor // Write-only binding to the contract
	IlendingpoolFilterer   // Log filterer for contract events
}

// IlendingpoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IlendingpoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlendingpoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IlendingpoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlendingpoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IlendingpoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IlendingpoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IlendingpoolSession struct {
	Contract     *Ilendingpool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IlendingpoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IlendingpoolCallerSession struct {
	Contract *IlendingpoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IlendingpoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IlendingpoolTransactorSession struct {
	Contract     *IlendingpoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IlendingpoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IlendingpoolRaw struct {
	Contract *Ilendingpool // Generic contract binding to access the raw methods on
}

// IlendingpoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IlendingpoolCallerRaw struct {
	Contract *IlendingpoolCaller // Generic read-only contract binding to access the raw methods on
}

// IlendingpoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IlendingpoolTransactorRaw struct {
	Contract *IlendingpoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIlendingpool creates a new instance of Ilendingpool, bound to a specific deployed contract.
func NewIlendingpool(address common.Address, backend bind.ContractBackend) (*Ilendingpool, error) {
	contract, err := bindIlendingpool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ilendingpool{IlendingpoolCaller: IlendingpoolCaller{contract: contract}, IlendingpoolTransactor: IlendingpoolTransactor{contract: contract}, IlendingpoolFilterer: IlendingpoolFilterer{contract: contract}}, nil
}

// NewIlendingpoolCaller creates a new read-only instance of Ilendingpool, bound to a specific deployed contract.
func NewIlendingpoolCaller(address common.Address, caller bind.ContractCaller) (*IlendingpoolCaller, error) {
	contract, err := bindIlendingpool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IlendingpoolCaller{contract: contract}, nil
}

// NewIlendingpoolTransactor creates a new write-only instance of Ilendingpool, bound to a specific deployed contract.
func NewIlendingpoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IlendingpoolTransactor, error) {
	contract, err := bindIlendingpool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IlendingpoolTransactor{contract: contract}, nil
}

// NewIlendingpoolFilterer creates a new log filterer instance of Ilendingpool, bound to a specific deployed contract.
func NewIlendingpoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IlendingpoolFilterer, error) {
	contract, err := bindIlendingpool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IlendingpoolFilterer{contract: contract}, nil
}

// bindIlendingpool binds a generic wrapper to an already deployed contract.
func bindIlendingpool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IlendingpoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ilendingpool *IlendingpoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ilendingpool.Contract.IlendingpoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ilendingpool *IlendingpoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ilendingpool.Contract.IlendingpoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ilendingpool *IlendingpoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ilendingpool.Contract.IlendingpoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ilendingpool *IlendingpoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ilendingpool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ilendingpool *IlendingpoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ilendingpool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ilendingpool *IlendingpoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ilendingpool.Contract.contract.Transact(opts, method, params...)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_Ilendingpool *IlendingpoolTransactor) Borrow(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Ilendingpool.contract.Transact(opts, "borrow", asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_Ilendingpool *IlendingpoolSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Ilendingpool.Contract.Borrow(&_Ilendingpool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_Ilendingpool *IlendingpoolTransactorSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Ilendingpool.Contract.Borrow(&_Ilendingpool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 rateMode, address onBehalfOf) returns(uint256)
func (_Ilendingpool *IlendingpoolTransactor) Repay(opts *bind.TransactOpts, asset common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Ilendingpool.contract.Transact(opts, "repay", asset, amount, rateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 rateMode, address onBehalfOf) returns(uint256)
func (_Ilendingpool *IlendingpoolSession) Repay(asset common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Ilendingpool.Contract.Repay(&_Ilendingpool.TransactOpts, asset, amount, rateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 rateMode, address onBehalfOf) returns(uint256)
func (_Ilendingpool *IlendingpoolTransactorSession) Repay(asset common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _Ilendingpool.Contract.Repay(&_Ilendingpool.TransactOpts, asset, amount, rateMode, onBehalfOf)
}
