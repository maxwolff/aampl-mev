// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package view

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

// ViewMetaData contains all meta data concerning the View contract.
var ViewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"aampl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCpi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowDurationSecs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowPercentOfPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetRate\",\"type\":\"uint256\"}],\"name\":\"computeSupplyDelta\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"expectedProfit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"profitETH\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmountAMPL\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRebasePercentage\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"supplyDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lendingPool\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondsUntilRebase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lowerOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRebaseTimeIntervalSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nowTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRebaseTimestamp\",\"type\":\"uint256\"}],\"name\":\"secondsUntilRebasePure\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uFrags\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uPolicy\",\"outputs\":[{\"internalType\":\"contractUFragmentsPolicy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ViewABI is the input ABI used to generate the binding from.
// Deprecated: Use ViewMetaData.ABI instead.
var ViewABI = ViewMetaData.ABI

// View is an auto generated Go binding around an Ethereum contract.
type View struct {
	ViewCaller     // Read-only binding to the contract
	ViewTransactor // Write-only binding to the contract
	ViewFilterer   // Log filterer for contract events
}

// ViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type ViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ViewSession struct {
	Contract     *View             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ViewCallerSession struct {
	Contract *ViewCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ViewTransactorSession struct {
	Contract     *ViewTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type ViewRaw struct {
	Contract *View // Generic contract binding to access the raw methods on
}

// ViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ViewCallerRaw struct {
	Contract *ViewCaller // Generic read-only contract binding to access the raw methods on
}

// ViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ViewTransactorRaw struct {
	Contract *ViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewView creates a new instance of View, bound to a specific deployed contract.
func NewView(address common.Address, backend bind.ContractBackend) (*View, error) {
	contract, err := bindView(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &View{ViewCaller: ViewCaller{contract: contract}, ViewTransactor: ViewTransactor{contract: contract}, ViewFilterer: ViewFilterer{contract: contract}}, nil
}

// NewViewCaller creates a new read-only instance of View, bound to a specific deployed contract.
func NewViewCaller(address common.Address, caller bind.ContractCaller) (*ViewCaller, error) {
	contract, err := bindView(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ViewCaller{contract: contract}, nil
}

// NewViewTransactor creates a new write-only instance of View, bound to a specific deployed contract.
func NewViewTransactor(address common.Address, transactor bind.ContractTransactor) (*ViewTransactor, error) {
	contract, err := bindView(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ViewTransactor{contract: contract}, nil
}

// NewViewFilterer creates a new log filterer instance of View, bound to a specific deployed contract.
func NewViewFilterer(address common.Address, filterer bind.ContractFilterer) (*ViewFilterer, error) {
	contract, err := bindView(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ViewFilterer{contract: contract}, nil
}

// bindView binds a generic wrapper to an already deployed contract.
func bindView(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ViewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_View *ViewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _View.Contract.ViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_View *ViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _View.Contract.ViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_View *ViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _View.Contract.ViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_View *ViewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _View.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_View *ViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _View.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_View *ViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _View.Contract.contract.Transact(opts, method, params...)
}

// Aampl is a free data retrieval call binding the contract method 0xbbd1afde.
//
// Solidity: function aampl() view returns(address)
func (_View *ViewCaller) Aampl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "aampl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aampl is a free data retrieval call binding the contract method 0xbbd1afde.
//
// Solidity: function aampl() view returns(address)
func (_View *ViewSession) Aampl() (common.Address, error) {
	return _View.Contract.Aampl(&_View.CallOpts)
}

// Aampl is a free data retrieval call binding the contract method 0xbbd1afde.
//
// Solidity: function aampl() view returns(address)
func (_View *ViewCallerSession) Aampl() (common.Address, error) {
	return _View.Contract.Aampl(&_View.CallOpts)
}

// BaseCpi is a free data retrieval call binding the contract method 0xa37069ee.
//
// Solidity: function baseCpi() view returns(uint256)
func (_View *ViewCaller) BaseCpi(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "baseCpi")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseCpi is a free data retrieval call binding the contract method 0xa37069ee.
//
// Solidity: function baseCpi() view returns(uint256)
func (_View *ViewSession) BaseCpi() (*big.Int, error) {
	return _View.Contract.BaseCpi(&_View.CallOpts)
}

// BaseCpi is a free data retrieval call binding the contract method 0xa37069ee.
//
// Solidity: function baseCpi() view returns(uint256)
func (_View *ViewCallerSession) BaseCpi() (*big.Int, error) {
	return _View.Contract.BaseCpi(&_View.CallOpts)
}

// BorrowDurationSecs is a free data retrieval call binding the contract method 0x9610db37.
//
// Solidity: function borrowDurationSecs() view returns(uint256)
func (_View *ViewCaller) BorrowDurationSecs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "borrowDurationSecs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowDurationSecs is a free data retrieval call binding the contract method 0x9610db37.
//
// Solidity: function borrowDurationSecs() view returns(uint256)
func (_View *ViewSession) BorrowDurationSecs() (*big.Int, error) {
	return _View.Contract.BorrowDurationSecs(&_View.CallOpts)
}

// BorrowDurationSecs is a free data retrieval call binding the contract method 0x9610db37.
//
// Solidity: function borrowDurationSecs() view returns(uint256)
func (_View *ViewCallerSession) BorrowDurationSecs() (*big.Int, error) {
	return _View.Contract.BorrowDurationSecs(&_View.CallOpts)
}

// BorrowPercentOfPower is a free data retrieval call binding the contract method 0x9b829b7b.
//
// Solidity: function borrowPercentOfPower() view returns(uint256)
func (_View *ViewCaller) BorrowPercentOfPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "borrowPercentOfPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowPercentOfPower is a free data retrieval call binding the contract method 0x9b829b7b.
//
// Solidity: function borrowPercentOfPower() view returns(uint256)
func (_View *ViewSession) BorrowPercentOfPower() (*big.Int, error) {
	return _View.Contract.BorrowPercentOfPower(&_View.CallOpts)
}

// BorrowPercentOfPower is a free data retrieval call binding the contract method 0x9b829b7b.
//
// Solidity: function borrowPercentOfPower() view returns(uint256)
func (_View *ViewCallerSession) BorrowPercentOfPower() (*big.Int, error) {
	return _View.Contract.BorrowPercentOfPower(&_View.CallOpts)
}

// ComputeSupplyDelta is a free data retrieval call binding the contract method 0xf2d0bd56.
//
// Solidity: function computeSupplyDelta(uint256 rate, uint256 targetRate) view returns(int256)
func (_View *ViewCaller) ComputeSupplyDelta(opts *bind.CallOpts, rate *big.Int, targetRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "computeSupplyDelta", rate, targetRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeSupplyDelta is a free data retrieval call binding the contract method 0xf2d0bd56.
//
// Solidity: function computeSupplyDelta(uint256 rate, uint256 targetRate) view returns(int256)
func (_View *ViewSession) ComputeSupplyDelta(rate *big.Int, targetRate *big.Int) (*big.Int, error) {
	return _View.Contract.ComputeSupplyDelta(&_View.CallOpts, rate, targetRate)
}

// ComputeSupplyDelta is a free data retrieval call binding the contract method 0xf2d0bd56.
//
// Solidity: function computeSupplyDelta(uint256 rate, uint256 targetRate) view returns(int256)
func (_View *ViewCallerSession) ComputeSupplyDelta(rate *big.Int, targetRate *big.Int) (*big.Int, error) {
	return _View.Contract.ComputeSupplyDelta(&_View.CallOpts, rate, targetRate)
}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_View *ViewCaller) LendingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "lendingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_View *ViewSession) LendingPool() (common.Address, error) {
	return _View.Contract.LendingPool(&_View.CallOpts)
}

// LendingPool is a free data retrieval call binding the contract method 0xa59a9973.
//
// Solidity: function lendingPool() view returns(address)
func (_View *ViewCallerSession) LendingPool() (common.Address, error) {
	return _View.Contract.LendingPool(&_View.CallOpts)
}

// SecondsUntilRebase is a free data retrieval call binding the contract method 0xa622fda8.
//
// Solidity: function secondsUntilRebase() view returns(uint256)
func (_View *ViewCaller) SecondsUntilRebase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "secondsUntilRebase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsUntilRebase is a free data retrieval call binding the contract method 0xa622fda8.
//
// Solidity: function secondsUntilRebase() view returns(uint256)
func (_View *ViewSession) SecondsUntilRebase() (*big.Int, error) {
	return _View.Contract.SecondsUntilRebase(&_View.CallOpts)
}

// SecondsUntilRebase is a free data retrieval call binding the contract method 0xa622fda8.
//
// Solidity: function secondsUntilRebase() view returns(uint256)
func (_View *ViewCallerSession) SecondsUntilRebase() (*big.Int, error) {
	return _View.Contract.SecondsUntilRebase(&_View.CallOpts)
}

// SecondsUntilRebasePure is a free data retrieval call binding the contract method 0x4e18569e.
//
// Solidity: function secondsUntilRebasePure(uint256 lowerOffset, uint256 upperOffset, uint256 minRebaseTimeIntervalSec, uint256 nowTimestamp, uint256 lastRebaseTimestamp) pure returns(uint256)
func (_View *ViewCaller) SecondsUntilRebasePure(opts *bind.CallOpts, lowerOffset *big.Int, upperOffset *big.Int, minRebaseTimeIntervalSec *big.Int, nowTimestamp *big.Int, lastRebaseTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "secondsUntilRebasePure", lowerOffset, upperOffset, minRebaseTimeIntervalSec, nowTimestamp, lastRebaseTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsUntilRebasePure is a free data retrieval call binding the contract method 0x4e18569e.
//
// Solidity: function secondsUntilRebasePure(uint256 lowerOffset, uint256 upperOffset, uint256 minRebaseTimeIntervalSec, uint256 nowTimestamp, uint256 lastRebaseTimestamp) pure returns(uint256)
func (_View *ViewSession) SecondsUntilRebasePure(lowerOffset *big.Int, upperOffset *big.Int, minRebaseTimeIntervalSec *big.Int, nowTimestamp *big.Int, lastRebaseTimestamp *big.Int) (*big.Int, error) {
	return _View.Contract.SecondsUntilRebasePure(&_View.CallOpts, lowerOffset, upperOffset, minRebaseTimeIntervalSec, nowTimestamp, lastRebaseTimestamp)
}

// SecondsUntilRebasePure is a free data retrieval call binding the contract method 0x4e18569e.
//
// Solidity: function secondsUntilRebasePure(uint256 lowerOffset, uint256 upperOffset, uint256 minRebaseTimeIntervalSec, uint256 nowTimestamp, uint256 lastRebaseTimestamp) pure returns(uint256)
func (_View *ViewCallerSession) SecondsUntilRebasePure(lowerOffset *big.Int, upperOffset *big.Int, minRebaseTimeIntervalSec *big.Int, nowTimestamp *big.Int, lastRebaseTimestamp *big.Int) (*big.Int, error) {
	return _View.Contract.SecondsUntilRebasePure(&_View.CallOpts, lowerOffset, upperOffset, minRebaseTimeIntervalSec, nowTimestamp, lastRebaseTimestamp)
}

// UFrags is a free data retrieval call binding the contract method 0xd3d55d51.
//
// Solidity: function uFrags() view returns(address)
func (_View *ViewCaller) UFrags(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "uFrags")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UFrags is a free data retrieval call binding the contract method 0xd3d55d51.
//
// Solidity: function uFrags() view returns(address)
func (_View *ViewSession) UFrags() (common.Address, error) {
	return _View.Contract.UFrags(&_View.CallOpts)
}

// UFrags is a free data retrieval call binding the contract method 0xd3d55d51.
//
// Solidity: function uFrags() view returns(address)
func (_View *ViewCallerSession) UFrags() (common.Address, error) {
	return _View.Contract.UFrags(&_View.CallOpts)
}

// UPolicy is a free data retrieval call binding the contract method 0x5f882cbb.
//
// Solidity: function uPolicy() view returns(address)
func (_View *ViewCaller) UPolicy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "uPolicy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UPolicy is a free data retrieval call binding the contract method 0x5f882cbb.
//
// Solidity: function uPolicy() view returns(address)
func (_View *ViewSession) UPolicy() (common.Address, error) {
	return _View.Contract.UPolicy(&_View.CallOpts)
}

// UPolicy is a free data retrieval call binding the contract method 0x5f882cbb.
//
// Solidity: function uPolicy() view returns(address)
func (_View *ViewCallerSession) UPolicy() (common.Address, error) {
	return _View.Contract.UPolicy(&_View.CallOpts)
}

// UserDebt is a free data retrieval call binding the contract method 0xb69e5c77.
//
// Solidity: function userDebt(address user) view returns(uint256)
func (_View *ViewCaller) UserDebt(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "userDebt", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserDebt is a free data retrieval call binding the contract method 0xb69e5c77.
//
// Solidity: function userDebt(address user) view returns(uint256)
func (_View *ViewSession) UserDebt(user common.Address) (*big.Int, error) {
	return _View.Contract.UserDebt(&_View.CallOpts, user)
}

// UserDebt is a free data retrieval call binding the contract method 0xb69e5c77.
//
// Solidity: function userDebt(address user) view returns(uint256)
func (_View *ViewCallerSession) UserDebt(user common.Address) (*big.Int, error) {
	return _View.Contract.UserDebt(&_View.CallOpts, user)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_View *ViewCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _View.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_View *ViewSession) Weth() (common.Address, error) {
	return _View.Contract.Weth(&_View.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_View *ViewCallerSession) Weth() (common.Address, error) {
	return _View.Contract.Weth(&_View.CallOpts)
}

// ExpectedProfit is a paid mutator transaction binding the contract method 0xea431811.
//
// Solidity: function expectedProfit(address user) returns(int256 profitETH, uint256 borrowAmountAMPL)
func (_View *ViewTransactor) ExpectedProfit(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _View.contract.Transact(opts, "expectedProfit", user)
}

// ExpectedProfit is a paid mutator transaction binding the contract method 0xea431811.
//
// Solidity: function expectedProfit(address user) returns(int256 profitETH, uint256 borrowAmountAMPL)
func (_View *ViewSession) ExpectedProfit(user common.Address) (*types.Transaction, error) {
	return _View.Contract.ExpectedProfit(&_View.TransactOpts, user)
}

// ExpectedProfit is a paid mutator transaction binding the contract method 0xea431811.
//
// Solidity: function expectedProfit(address user) returns(int256 profitETH, uint256 borrowAmountAMPL)
func (_View *ViewTransactorSession) ExpectedProfit(user common.Address) (*types.Transaction, error) {
	return _View.Contract.ExpectedProfit(&_View.TransactOpts, user)
}

// GetRebasePercentage is a paid mutator transaction binding the contract method 0xc4f16c2f.
//
// Solidity: function getRebasePercentage() returns(int256 supplyDelta, uint256 totalSupply)
func (_View *ViewTransactor) GetRebasePercentage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _View.contract.Transact(opts, "getRebasePercentage")
}

// GetRebasePercentage is a paid mutator transaction binding the contract method 0xc4f16c2f.
//
// Solidity: function getRebasePercentage() returns(int256 supplyDelta, uint256 totalSupply)
func (_View *ViewSession) GetRebasePercentage() (*types.Transaction, error) {
	return _View.Contract.GetRebasePercentage(&_View.TransactOpts)
}

// GetRebasePercentage is a paid mutator transaction binding the contract method 0xc4f16c2f.
//
// Solidity: function getRebasePercentage() returns(int256 supplyDelta, uint256 totalSupply)
func (_View *ViewTransactorSession) GetRebasePercentage() (*types.Transaction, error) {
	return _View.Contract.GetRebasePercentage(&_View.TransactOpts)
}
