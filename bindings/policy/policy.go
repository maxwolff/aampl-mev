// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package policy

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

// PolicyMetaData contains all meta data concerning the Policy contract.
var PolicyMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exchangeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cpi\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"requestedSupplyAdjustment\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestampSec\",\"type\":\"uint256\"}],\"name\":\"LogRebase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"normalizedRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"lower\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"upper\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"growth\",\"type\":\"int256\"}],\"name\":\"computeRebasePercentage\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cpiOracle\",\"outputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deviationThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalAmpleforthEpochAndAMPLSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inRebaseWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"contractIUFragments\",\"name\":\"uFrags_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseCpi_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRebaseTimestampSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketOracle\",\"outputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minRebaseTimeIntervalSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orchestrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebaseFunctionGrowth\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebaseFunctionLowerPercentage\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebaseFunctionUpperPercentage\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebaseLag\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebaseWindowLengthSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebaseWindowOffsetSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"cpiOracle_\",\"type\":\"address\"}],\"name\":\"setCpiOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deviationThreshold_\",\"type\":\"uint256\"}],\"name\":\"setDeviationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"marketOracle_\",\"type\":\"address\"}],\"name\":\"setMarketOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orchestrator_\",\"type\":\"address\"}],\"name\":\"setOrchestrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"rebaseFunctionGrowth_\",\"type\":\"int256\"}],\"name\":\"setRebaseFunctionGrowth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"rebaseFunctionLowerPercentage_\",\"type\":\"int256\"}],\"name\":\"setRebaseFunctionLowerPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"rebaseFunctionUpperPercentage_\",\"type\":\"int256\"}],\"name\":\"setRebaseFunctionUpperPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minRebaseTimeIntervalSec_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rebaseWindowOffsetSec_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rebaseWindowLengthSec_\",\"type\":\"uint256\"}],\"name\":\"setRebaseTimingParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uFrags\",\"outputs\":[{\"internalType\":\"contractIUFragments\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PolicyABI is the input ABI used to generate the binding from.
// Deprecated: Use PolicyMetaData.ABI instead.
var PolicyABI = PolicyMetaData.ABI

// Policy is an auto generated Go binding around an Ethereum contract.
type Policy struct {
	PolicyCaller     // Read-only binding to the contract
	PolicyTransactor // Write-only binding to the contract
	PolicyFilterer   // Log filterer for contract events
}

// PolicyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolicyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolicyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolicyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolicySession struct {
	Contract     *Policy           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolicyCallerSession struct {
	Contract *PolicyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PolicyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolicyTransactorSession struct {
	Contract     *PolicyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolicyRaw struct {
	Contract *Policy // Generic contract binding to access the raw methods on
}

// PolicyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolicyCallerRaw struct {
	Contract *PolicyCaller // Generic read-only contract binding to access the raw methods on
}

// PolicyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolicyTransactorRaw struct {
	Contract *PolicyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolicy creates a new instance of Policy, bound to a specific deployed contract.
func NewPolicy(address common.Address, backend bind.ContractBackend) (*Policy, error) {
	contract, err := bindPolicy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Policy{PolicyCaller: PolicyCaller{contract: contract}, PolicyTransactor: PolicyTransactor{contract: contract}, PolicyFilterer: PolicyFilterer{contract: contract}}, nil
}

// NewPolicyCaller creates a new read-only instance of Policy, bound to a specific deployed contract.
func NewPolicyCaller(address common.Address, caller bind.ContractCaller) (*PolicyCaller, error) {
	contract, err := bindPolicy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyCaller{contract: contract}, nil
}

// NewPolicyTransactor creates a new write-only instance of Policy, bound to a specific deployed contract.
func NewPolicyTransactor(address common.Address, transactor bind.ContractTransactor) (*PolicyTransactor, error) {
	contract, err := bindPolicy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyTransactor{contract: contract}, nil
}

// NewPolicyFilterer creates a new log filterer instance of Policy, bound to a specific deployed contract.
func NewPolicyFilterer(address common.Address, filterer bind.ContractFilterer) (*PolicyFilterer, error) {
	contract, err := bindPolicy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolicyFilterer{contract: contract}, nil
}

// bindPolicy binds a generic wrapper to an already deployed contract.
func bindPolicy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PolicyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.PolicyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.PolicyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Policy *PolicyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Policy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Policy *PolicyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Policy *PolicyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Policy.Contract.contract.Transact(opts, method, params...)
}

// ComputeRebasePercentage is a free data retrieval call binding the contract method 0x3a785af1.
//
// Solidity: function computeRebasePercentage(int256 normalizedRate, int256 lower, int256 upper, int256 growth) pure returns(int256)
func (_Policy *PolicyCaller) ComputeRebasePercentage(opts *bind.CallOpts, normalizedRate *big.Int, lower *big.Int, upper *big.Int, growth *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "computeRebasePercentage", normalizedRate, lower, upper, growth)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeRebasePercentage is a free data retrieval call binding the contract method 0x3a785af1.
//
// Solidity: function computeRebasePercentage(int256 normalizedRate, int256 lower, int256 upper, int256 growth) pure returns(int256)
func (_Policy *PolicySession) ComputeRebasePercentage(normalizedRate *big.Int, lower *big.Int, upper *big.Int, growth *big.Int) (*big.Int, error) {
	return _Policy.Contract.ComputeRebasePercentage(&_Policy.CallOpts, normalizedRate, lower, upper, growth)
}

// ComputeRebasePercentage is a free data retrieval call binding the contract method 0x3a785af1.
//
// Solidity: function computeRebasePercentage(int256 normalizedRate, int256 lower, int256 upper, int256 growth) pure returns(int256)
func (_Policy *PolicyCallerSession) ComputeRebasePercentage(normalizedRate *big.Int, lower *big.Int, upper *big.Int, growth *big.Int) (*big.Int, error) {
	return _Policy.Contract.ComputeRebasePercentage(&_Policy.CallOpts, normalizedRate, lower, upper, growth)
}

// CpiOracle is a free data retrieval call binding the contract method 0xab33c5ca.
//
// Solidity: function cpiOracle() view returns(address)
func (_Policy *PolicyCaller) CpiOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "cpiOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CpiOracle is a free data retrieval call binding the contract method 0xab33c5ca.
//
// Solidity: function cpiOracle() view returns(address)
func (_Policy *PolicySession) CpiOracle() (common.Address, error) {
	return _Policy.Contract.CpiOracle(&_Policy.CallOpts)
}

// CpiOracle is a free data retrieval call binding the contract method 0xab33c5ca.
//
// Solidity: function cpiOracle() view returns(address)
func (_Policy *PolicyCallerSession) CpiOracle() (common.Address, error) {
	return _Policy.Contract.CpiOracle(&_Policy.CallOpts)
}

// DeviationThreshold is a free data retrieval call binding the contract method 0xd94ad837.
//
// Solidity: function deviationThreshold() view returns(uint256)
func (_Policy *PolicyCaller) DeviationThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "deviationThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeviationThreshold is a free data retrieval call binding the contract method 0xd94ad837.
//
// Solidity: function deviationThreshold() view returns(uint256)
func (_Policy *PolicySession) DeviationThreshold() (*big.Int, error) {
	return _Policy.Contract.DeviationThreshold(&_Policy.CallOpts)
}

// DeviationThreshold is a free data retrieval call binding the contract method 0xd94ad837.
//
// Solidity: function deviationThreshold() view returns(uint256)
func (_Policy *PolicyCallerSession) DeviationThreshold() (*big.Int, error) {
	return _Policy.Contract.DeviationThreshold(&_Policy.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Policy *PolicyCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Policy *PolicySession) Epoch() (*big.Int, error) {
	return _Policy.Contract.Epoch(&_Policy.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Policy *PolicyCallerSession) Epoch() (*big.Int, error) {
	return _Policy.Contract.Epoch(&_Policy.CallOpts)
}

// GlobalAmpleforthEpochAndAMPLSupply is a free data retrieval call binding the contract method 0x105bad4f.
//
// Solidity: function globalAmpleforthEpochAndAMPLSupply() view returns(uint256, uint256)
func (_Policy *PolicyCaller) GlobalAmpleforthEpochAndAMPLSupply(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "globalAmpleforthEpochAndAMPLSupply")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GlobalAmpleforthEpochAndAMPLSupply is a free data retrieval call binding the contract method 0x105bad4f.
//
// Solidity: function globalAmpleforthEpochAndAMPLSupply() view returns(uint256, uint256)
func (_Policy *PolicySession) GlobalAmpleforthEpochAndAMPLSupply() (*big.Int, *big.Int, error) {
	return _Policy.Contract.GlobalAmpleforthEpochAndAMPLSupply(&_Policy.CallOpts)
}

// GlobalAmpleforthEpochAndAMPLSupply is a free data retrieval call binding the contract method 0x105bad4f.
//
// Solidity: function globalAmpleforthEpochAndAMPLSupply() view returns(uint256, uint256)
func (_Policy *PolicyCallerSession) GlobalAmpleforthEpochAndAMPLSupply() (*big.Int, *big.Int, error) {
	return _Policy.Contract.GlobalAmpleforthEpochAndAMPLSupply(&_Policy.CallOpts)
}

// InRebaseWindow is a free data retrieval call binding the contract method 0x111d0498.
//
// Solidity: function inRebaseWindow() view returns(bool)
func (_Policy *PolicyCaller) InRebaseWindow(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "inRebaseWindow")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRebaseWindow is a free data retrieval call binding the contract method 0x111d0498.
//
// Solidity: function inRebaseWindow() view returns(bool)
func (_Policy *PolicySession) InRebaseWindow() (bool, error) {
	return _Policy.Contract.InRebaseWindow(&_Policy.CallOpts)
}

// InRebaseWindow is a free data retrieval call binding the contract method 0x111d0498.
//
// Solidity: function inRebaseWindow() view returns(bool)
func (_Policy *PolicyCallerSession) InRebaseWindow() (bool, error) {
	return _Policy.Contract.InRebaseWindow(&_Policy.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Policy *PolicyCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Policy *PolicySession) IsOwner() (bool, error) {
	return _Policy.Contract.IsOwner(&_Policy.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Policy *PolicyCallerSession) IsOwner() (bool, error) {
	return _Policy.Contract.IsOwner(&_Policy.CallOpts)
}

// LastRebaseTimestampSec is a free data retrieval call binding the contract method 0x3a93069b.
//
// Solidity: function lastRebaseTimestampSec() view returns(uint256)
func (_Policy *PolicyCaller) LastRebaseTimestampSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "lastRebaseTimestampSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRebaseTimestampSec is a free data retrieval call binding the contract method 0x3a93069b.
//
// Solidity: function lastRebaseTimestampSec() view returns(uint256)
func (_Policy *PolicySession) LastRebaseTimestampSec() (*big.Int, error) {
	return _Policy.Contract.LastRebaseTimestampSec(&_Policy.CallOpts)
}

// LastRebaseTimestampSec is a free data retrieval call binding the contract method 0x3a93069b.
//
// Solidity: function lastRebaseTimestampSec() view returns(uint256)
func (_Policy *PolicyCallerSession) LastRebaseTimestampSec() (*big.Int, error) {
	return _Policy.Contract.LastRebaseTimestampSec(&_Policy.CallOpts)
}

// MarketOracle is a free data retrieval call binding the contract method 0x60961528.
//
// Solidity: function marketOracle() view returns(address)
func (_Policy *PolicyCaller) MarketOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "marketOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketOracle is a free data retrieval call binding the contract method 0x60961528.
//
// Solidity: function marketOracle() view returns(address)
func (_Policy *PolicySession) MarketOracle() (common.Address, error) {
	return _Policy.Contract.MarketOracle(&_Policy.CallOpts)
}

// MarketOracle is a free data retrieval call binding the contract method 0x60961528.
//
// Solidity: function marketOracle() view returns(address)
func (_Policy *PolicyCallerSession) MarketOracle() (common.Address, error) {
	return _Policy.Contract.MarketOracle(&_Policy.CallOpts)
}

// MinRebaseTimeIntervalSec is a free data retrieval call binding the contract method 0x02101899.
//
// Solidity: function minRebaseTimeIntervalSec() view returns(uint256)
func (_Policy *PolicyCaller) MinRebaseTimeIntervalSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "minRebaseTimeIntervalSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinRebaseTimeIntervalSec is a free data retrieval call binding the contract method 0x02101899.
//
// Solidity: function minRebaseTimeIntervalSec() view returns(uint256)
func (_Policy *PolicySession) MinRebaseTimeIntervalSec() (*big.Int, error) {
	return _Policy.Contract.MinRebaseTimeIntervalSec(&_Policy.CallOpts)
}

// MinRebaseTimeIntervalSec is a free data retrieval call binding the contract method 0x02101899.
//
// Solidity: function minRebaseTimeIntervalSec() view returns(uint256)
func (_Policy *PolicyCallerSession) MinRebaseTimeIntervalSec() (*big.Int, error) {
	return _Policy.Contract.MinRebaseTimeIntervalSec(&_Policy.CallOpts)
}

// Orchestrator is a free data retrieval call binding the contract method 0xb74795d9.
//
// Solidity: function orchestrator() view returns(address)
func (_Policy *PolicyCaller) Orchestrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "orchestrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Orchestrator is a free data retrieval call binding the contract method 0xb74795d9.
//
// Solidity: function orchestrator() view returns(address)
func (_Policy *PolicySession) Orchestrator() (common.Address, error) {
	return _Policy.Contract.Orchestrator(&_Policy.CallOpts)
}

// Orchestrator is a free data retrieval call binding the contract method 0xb74795d9.
//
// Solidity: function orchestrator() view returns(address)
func (_Policy *PolicyCallerSession) Orchestrator() (common.Address, error) {
	return _Policy.Contract.Orchestrator(&_Policy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Policy *PolicyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Policy *PolicySession) Owner() (common.Address, error) {
	return _Policy.Contract.Owner(&_Policy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Policy *PolicyCallerSession) Owner() (common.Address, error) {
	return _Policy.Contract.Owner(&_Policy.CallOpts)
}

// RebaseFunctionGrowth is a free data retrieval call binding the contract method 0x7486cdea.
//
// Solidity: function rebaseFunctionGrowth() view returns(int256)
func (_Policy *PolicyCaller) RebaseFunctionGrowth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "rebaseFunctionGrowth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseFunctionGrowth is a free data retrieval call binding the contract method 0x7486cdea.
//
// Solidity: function rebaseFunctionGrowth() view returns(int256)
func (_Policy *PolicySession) RebaseFunctionGrowth() (*big.Int, error) {
	return _Policy.Contract.RebaseFunctionGrowth(&_Policy.CallOpts)
}

// RebaseFunctionGrowth is a free data retrieval call binding the contract method 0x7486cdea.
//
// Solidity: function rebaseFunctionGrowth() view returns(int256)
func (_Policy *PolicyCallerSession) RebaseFunctionGrowth() (*big.Int, error) {
	return _Policy.Contract.RebaseFunctionGrowth(&_Policy.CallOpts)
}

// RebaseFunctionLowerPercentage is a free data retrieval call binding the contract method 0x5ee01540.
//
// Solidity: function rebaseFunctionLowerPercentage() view returns(int256)
func (_Policy *PolicyCaller) RebaseFunctionLowerPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "rebaseFunctionLowerPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseFunctionLowerPercentage is a free data retrieval call binding the contract method 0x5ee01540.
//
// Solidity: function rebaseFunctionLowerPercentage() view returns(int256)
func (_Policy *PolicySession) RebaseFunctionLowerPercentage() (*big.Int, error) {
	return _Policy.Contract.RebaseFunctionLowerPercentage(&_Policy.CallOpts)
}

// RebaseFunctionLowerPercentage is a free data retrieval call binding the contract method 0x5ee01540.
//
// Solidity: function rebaseFunctionLowerPercentage() view returns(int256)
func (_Policy *PolicyCallerSession) RebaseFunctionLowerPercentage() (*big.Int, error) {
	return _Policy.Contract.RebaseFunctionLowerPercentage(&_Policy.CallOpts)
}

// RebaseFunctionUpperPercentage is a free data retrieval call binding the contract method 0x9db59b2f.
//
// Solidity: function rebaseFunctionUpperPercentage() view returns(int256)
func (_Policy *PolicyCaller) RebaseFunctionUpperPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "rebaseFunctionUpperPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseFunctionUpperPercentage is a free data retrieval call binding the contract method 0x9db59b2f.
//
// Solidity: function rebaseFunctionUpperPercentage() view returns(int256)
func (_Policy *PolicySession) RebaseFunctionUpperPercentage() (*big.Int, error) {
	return _Policy.Contract.RebaseFunctionUpperPercentage(&_Policy.CallOpts)
}

// RebaseFunctionUpperPercentage is a free data retrieval call binding the contract method 0x9db59b2f.
//
// Solidity: function rebaseFunctionUpperPercentage() view returns(int256)
func (_Policy *PolicyCallerSession) RebaseFunctionUpperPercentage() (*big.Int, error) {
	return _Policy.Contract.RebaseFunctionUpperPercentage(&_Policy.CallOpts)
}

// RebaseLag is a free data retrieval call binding the contract method 0x63f6d4c8.
//
// Solidity: function rebaseLag() pure returns(uint256)
func (_Policy *PolicyCaller) RebaseLag(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "rebaseLag")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseLag is a free data retrieval call binding the contract method 0x63f6d4c8.
//
// Solidity: function rebaseLag() pure returns(uint256)
func (_Policy *PolicySession) RebaseLag() (*big.Int, error) {
	return _Policy.Contract.RebaseLag(&_Policy.CallOpts)
}

// RebaseLag is a free data retrieval call binding the contract method 0x63f6d4c8.
//
// Solidity: function rebaseLag() pure returns(uint256)
func (_Policy *PolicyCallerSession) RebaseLag() (*big.Int, error) {
	return _Policy.Contract.RebaseLag(&_Policy.CallOpts)
}

// RebaseWindowLengthSec is a free data retrieval call binding the contract method 0x9466120f.
//
// Solidity: function rebaseWindowLengthSec() view returns(uint256)
func (_Policy *PolicyCaller) RebaseWindowLengthSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "rebaseWindowLengthSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseWindowLengthSec is a free data retrieval call binding the contract method 0x9466120f.
//
// Solidity: function rebaseWindowLengthSec() view returns(uint256)
func (_Policy *PolicySession) RebaseWindowLengthSec() (*big.Int, error) {
	return _Policy.Contract.RebaseWindowLengthSec(&_Policy.CallOpts)
}

// RebaseWindowLengthSec is a free data retrieval call binding the contract method 0x9466120f.
//
// Solidity: function rebaseWindowLengthSec() view returns(uint256)
func (_Policy *PolicyCallerSession) RebaseWindowLengthSec() (*big.Int, error) {
	return _Policy.Contract.RebaseWindowLengthSec(&_Policy.CallOpts)
}

// RebaseWindowOffsetSec is a free data retrieval call binding the contract method 0x7052b902.
//
// Solidity: function rebaseWindowOffsetSec() view returns(uint256)
func (_Policy *PolicyCaller) RebaseWindowOffsetSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "rebaseWindowOffsetSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseWindowOffsetSec is a free data retrieval call binding the contract method 0x7052b902.
//
// Solidity: function rebaseWindowOffsetSec() view returns(uint256)
func (_Policy *PolicySession) RebaseWindowOffsetSec() (*big.Int, error) {
	return _Policy.Contract.RebaseWindowOffsetSec(&_Policy.CallOpts)
}

// RebaseWindowOffsetSec is a free data retrieval call binding the contract method 0x7052b902.
//
// Solidity: function rebaseWindowOffsetSec() view returns(uint256)
func (_Policy *PolicyCallerSession) RebaseWindowOffsetSec() (*big.Int, error) {
	return _Policy.Contract.RebaseWindowOffsetSec(&_Policy.CallOpts)
}

// UFrags is a free data retrieval call binding the contract method 0xd3d55d51.
//
// Solidity: function uFrags() view returns(address)
func (_Policy *PolicyCaller) UFrags(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Policy.contract.Call(opts, &out, "uFrags")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UFrags is a free data retrieval call binding the contract method 0xd3d55d51.
//
// Solidity: function uFrags() view returns(address)
func (_Policy *PolicySession) UFrags() (common.Address, error) {
	return _Policy.Contract.UFrags(&_Policy.CallOpts)
}

// UFrags is a free data retrieval call binding the contract method 0xd3d55d51.
//
// Solidity: function uFrags() view returns(address)
func (_Policy *PolicyCallerSession) UFrags() (common.Address, error) {
	return _Policy.Contract.UFrags(&_Policy.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address owner_, address uFrags_, uint256 baseCpi_) returns()
func (_Policy *PolicyTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, uFrags_ common.Address, baseCpi_ *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "initialize", owner_, uFrags_, baseCpi_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address owner_, address uFrags_, uint256 baseCpi_) returns()
func (_Policy *PolicySession) Initialize(owner_ common.Address, uFrags_ common.Address, baseCpi_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.Initialize(&_Policy.TransactOpts, owner_, uFrags_, baseCpi_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address owner_, address uFrags_, uint256 baseCpi_) returns()
func (_Policy *PolicyTransactorSession) Initialize(owner_ common.Address, uFrags_ common.Address, baseCpi_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.Initialize(&_Policy.TransactOpts, owner_, uFrags_, baseCpi_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Policy *PolicyTransactor) Initialize0(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "initialize0", sender)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Policy *PolicySession) Initialize0(sender common.Address) (*types.Transaction, error) {
	return _Policy.Contract.Initialize0(&_Policy.TransactOpts, sender)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Policy *PolicyTransactorSession) Initialize0(sender common.Address) (*types.Transaction, error) {
	return _Policy.Contract.Initialize0(&_Policy.TransactOpts, sender)
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns()
func (_Policy *PolicyTransactor) Rebase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "rebase")
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns()
func (_Policy *PolicySession) Rebase() (*types.Transaction, error) {
	return _Policy.Contract.Rebase(&_Policy.TransactOpts)
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns()
func (_Policy *PolicyTransactorSession) Rebase() (*types.Transaction, error) {
	return _Policy.Contract.Rebase(&_Policy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Policy *PolicyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Policy *PolicySession) RenounceOwnership() (*types.Transaction, error) {
	return _Policy.Contract.RenounceOwnership(&_Policy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Policy *PolicyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Policy.Contract.RenounceOwnership(&_Policy.TransactOpts)
}

// SetCpiOracle is a paid mutator transaction binding the contract method 0x3d6a46e5.
//
// Solidity: function setCpiOracle(address cpiOracle_) returns()
func (_Policy *PolicyTransactor) SetCpiOracle(opts *bind.TransactOpts, cpiOracle_ common.Address) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "setCpiOracle", cpiOracle_)
}

// SetCpiOracle is a paid mutator transaction binding the contract method 0x3d6a46e5.
//
// Solidity: function setCpiOracle(address cpiOracle_) returns()
func (_Policy *PolicySession) SetCpiOracle(cpiOracle_ common.Address) (*types.Transaction, error) {
	return _Policy.Contract.SetCpiOracle(&_Policy.TransactOpts, cpiOracle_)
}

// SetCpiOracle is a paid mutator transaction binding the contract method 0x3d6a46e5.
//
// Solidity: function setCpiOracle(address cpiOracle_) returns()
func (_Policy *PolicyTransactorSession) SetCpiOracle(cpiOracle_ common.Address) (*types.Transaction, error) {
	return _Policy.Contract.SetCpiOracle(&_Policy.TransactOpts, cpiOracle_)
}

// SetDeviationThreshold is a paid mutator transaction binding the contract method 0x53a15edc.
//
// Solidity: function setDeviationThreshold(uint256 deviationThreshold_) returns()
func (_Policy *PolicyTransactor) SetDeviationThreshold(opts *bind.TransactOpts, deviationThreshold_ *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "setDeviationThreshold", deviationThreshold_)
}

// SetDeviationThreshold is a paid mutator transaction binding the contract method 0x53a15edc.
//
// Solidity: function setDeviationThreshold(uint256 deviationThreshold_) returns()
func (_Policy *PolicySession) SetDeviationThreshold(deviationThreshold_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.SetDeviationThreshold(&_Policy.TransactOpts, deviationThreshold_)
}

// SetDeviationThreshold is a paid mutator transaction binding the contract method 0x53a15edc.
//
// Solidity: function setDeviationThreshold(uint256 deviationThreshold_) returns()
func (_Policy *PolicyTransactorSession) SetDeviationThreshold(deviationThreshold_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.SetDeviationThreshold(&_Policy.TransactOpts, deviationThreshold_)
}

// SetMarketOracle is a paid mutator transaction binding the contract method 0x9e30bac5.
//
// Solidity: function setMarketOracle(address marketOracle_) returns()
func (_Policy *PolicyTransactor) SetMarketOracle(opts *bind.TransactOpts, marketOracle_ common.Address) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "setMarketOracle", marketOracle_)
}

// SetMarketOracle is a paid mutator transaction binding the contract method 0x9e30bac5.
//
// Solidity: function setMarketOracle(address marketOracle_) returns()
func (_Policy *PolicySession) SetMarketOracle(marketOracle_ common.Address) (*types.Transaction, error) {
	return _Policy.Contract.SetMarketOracle(&_Policy.TransactOpts, marketOracle_)
}

// SetMarketOracle is a paid mutator transaction binding the contract method 0x9e30bac5.
//
// Solidity: function setMarketOracle(address marketOracle_) returns()
func (_Policy *PolicyTransactorSession) SetMarketOracle(marketOracle_ common.Address) (*types.Transaction, error) {
	return _Policy.Contract.SetMarketOracle(&_Policy.TransactOpts, marketOracle_)
}

// SetOrchestrator is a paid mutator transaction binding the contract method 0xcd28ef0d.
//
// Solidity: function setOrchestrator(address orchestrator_) returns()
func (_Policy *PolicyTransactor) SetOrchestrator(opts *bind.TransactOpts, orchestrator_ common.Address) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "setOrchestrator", orchestrator_)
}

// SetOrchestrator is a paid mutator transaction binding the contract method 0xcd28ef0d.
//
// Solidity: function setOrchestrator(address orchestrator_) returns()
func (_Policy *PolicySession) SetOrchestrator(orchestrator_ common.Address) (*types.Transaction, error) {
	return _Policy.Contract.SetOrchestrator(&_Policy.TransactOpts, orchestrator_)
}

// SetOrchestrator is a paid mutator transaction binding the contract method 0xcd28ef0d.
//
// Solidity: function setOrchestrator(address orchestrator_) returns()
func (_Policy *PolicyTransactorSession) SetOrchestrator(orchestrator_ common.Address) (*types.Transaction, error) {
	return _Policy.Contract.SetOrchestrator(&_Policy.TransactOpts, orchestrator_)
}

// SetRebaseFunctionGrowth is a paid mutator transaction binding the contract method 0x8001066d.
//
// Solidity: function setRebaseFunctionGrowth(int256 rebaseFunctionGrowth_) returns()
func (_Policy *PolicyTransactor) SetRebaseFunctionGrowth(opts *bind.TransactOpts, rebaseFunctionGrowth_ *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "setRebaseFunctionGrowth", rebaseFunctionGrowth_)
}

// SetRebaseFunctionGrowth is a paid mutator transaction binding the contract method 0x8001066d.
//
// Solidity: function setRebaseFunctionGrowth(int256 rebaseFunctionGrowth_) returns()
func (_Policy *PolicySession) SetRebaseFunctionGrowth(rebaseFunctionGrowth_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.SetRebaseFunctionGrowth(&_Policy.TransactOpts, rebaseFunctionGrowth_)
}

// SetRebaseFunctionGrowth is a paid mutator transaction binding the contract method 0x8001066d.
//
// Solidity: function setRebaseFunctionGrowth(int256 rebaseFunctionGrowth_) returns()
func (_Policy *PolicyTransactorSession) SetRebaseFunctionGrowth(rebaseFunctionGrowth_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.SetRebaseFunctionGrowth(&_Policy.TransactOpts, rebaseFunctionGrowth_)
}

// SetRebaseFunctionLowerPercentage is a paid mutator transaction binding the contract method 0xe9fa88a4.
//
// Solidity: function setRebaseFunctionLowerPercentage(int256 rebaseFunctionLowerPercentage_) returns()
func (_Policy *PolicyTransactor) SetRebaseFunctionLowerPercentage(opts *bind.TransactOpts, rebaseFunctionLowerPercentage_ *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "setRebaseFunctionLowerPercentage", rebaseFunctionLowerPercentage_)
}

// SetRebaseFunctionLowerPercentage is a paid mutator transaction binding the contract method 0xe9fa88a4.
//
// Solidity: function setRebaseFunctionLowerPercentage(int256 rebaseFunctionLowerPercentage_) returns()
func (_Policy *PolicySession) SetRebaseFunctionLowerPercentage(rebaseFunctionLowerPercentage_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.SetRebaseFunctionLowerPercentage(&_Policy.TransactOpts, rebaseFunctionLowerPercentage_)
}

// SetRebaseFunctionLowerPercentage is a paid mutator transaction binding the contract method 0xe9fa88a4.
//
// Solidity: function setRebaseFunctionLowerPercentage(int256 rebaseFunctionLowerPercentage_) returns()
func (_Policy *PolicyTransactorSession) SetRebaseFunctionLowerPercentage(rebaseFunctionLowerPercentage_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.SetRebaseFunctionLowerPercentage(&_Policy.TransactOpts, rebaseFunctionLowerPercentage_)
}

// SetRebaseFunctionUpperPercentage is a paid mutator transaction binding the contract method 0xf4fefa49.
//
// Solidity: function setRebaseFunctionUpperPercentage(int256 rebaseFunctionUpperPercentage_) returns()
func (_Policy *PolicyTransactor) SetRebaseFunctionUpperPercentage(opts *bind.TransactOpts, rebaseFunctionUpperPercentage_ *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "setRebaseFunctionUpperPercentage", rebaseFunctionUpperPercentage_)
}

// SetRebaseFunctionUpperPercentage is a paid mutator transaction binding the contract method 0xf4fefa49.
//
// Solidity: function setRebaseFunctionUpperPercentage(int256 rebaseFunctionUpperPercentage_) returns()
func (_Policy *PolicySession) SetRebaseFunctionUpperPercentage(rebaseFunctionUpperPercentage_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.SetRebaseFunctionUpperPercentage(&_Policy.TransactOpts, rebaseFunctionUpperPercentage_)
}

// SetRebaseFunctionUpperPercentage is a paid mutator transaction binding the contract method 0xf4fefa49.
//
// Solidity: function setRebaseFunctionUpperPercentage(int256 rebaseFunctionUpperPercentage_) returns()
func (_Policy *PolicyTransactorSession) SetRebaseFunctionUpperPercentage(rebaseFunctionUpperPercentage_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.SetRebaseFunctionUpperPercentage(&_Policy.TransactOpts, rebaseFunctionUpperPercentage_)
}

// SetRebaseTimingParameters is a paid mutator transaction binding the contract method 0x16250fd4.
//
// Solidity: function setRebaseTimingParameters(uint256 minRebaseTimeIntervalSec_, uint256 rebaseWindowOffsetSec_, uint256 rebaseWindowLengthSec_) returns()
func (_Policy *PolicyTransactor) SetRebaseTimingParameters(opts *bind.TransactOpts, minRebaseTimeIntervalSec_ *big.Int, rebaseWindowOffsetSec_ *big.Int, rebaseWindowLengthSec_ *big.Int) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "setRebaseTimingParameters", minRebaseTimeIntervalSec_, rebaseWindowOffsetSec_, rebaseWindowLengthSec_)
}

// SetRebaseTimingParameters is a paid mutator transaction binding the contract method 0x16250fd4.
//
// Solidity: function setRebaseTimingParameters(uint256 minRebaseTimeIntervalSec_, uint256 rebaseWindowOffsetSec_, uint256 rebaseWindowLengthSec_) returns()
func (_Policy *PolicySession) SetRebaseTimingParameters(minRebaseTimeIntervalSec_ *big.Int, rebaseWindowOffsetSec_ *big.Int, rebaseWindowLengthSec_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.SetRebaseTimingParameters(&_Policy.TransactOpts, minRebaseTimeIntervalSec_, rebaseWindowOffsetSec_, rebaseWindowLengthSec_)
}

// SetRebaseTimingParameters is a paid mutator transaction binding the contract method 0x16250fd4.
//
// Solidity: function setRebaseTimingParameters(uint256 minRebaseTimeIntervalSec_, uint256 rebaseWindowOffsetSec_, uint256 rebaseWindowLengthSec_) returns()
func (_Policy *PolicyTransactorSession) SetRebaseTimingParameters(minRebaseTimeIntervalSec_ *big.Int, rebaseWindowOffsetSec_ *big.Int, rebaseWindowLengthSec_ *big.Int) (*types.Transaction, error) {
	return _Policy.Contract.SetRebaseTimingParameters(&_Policy.TransactOpts, minRebaseTimeIntervalSec_, rebaseWindowOffsetSec_, rebaseWindowLengthSec_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Policy *PolicyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Policy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Policy *PolicySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Policy.Contract.TransferOwnership(&_Policy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Policy *PolicyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Policy.Contract.TransferOwnership(&_Policy.TransactOpts, newOwner)
}

// PolicyLogRebaseIterator is returned from FilterLogRebase and is used to iterate over the raw logs and unpacked data for LogRebase events raised by the Policy contract.
type PolicyLogRebaseIterator struct {
	Event *PolicyLogRebase // Event containing the contract specifics and raw log

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
func (it *PolicyLogRebaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyLogRebase)
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
		it.Event = new(PolicyLogRebase)
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
func (it *PolicyLogRebaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyLogRebaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyLogRebase represents a LogRebase event raised by the Policy contract.
type PolicyLogRebase struct {
	Epoch                     *big.Int
	ExchangeRate              *big.Int
	Cpi                       *big.Int
	RequestedSupplyAdjustment *big.Int
	TimestampSec              *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterLogRebase is a free log retrieval operation binding the contract event 0x41d948a7f29cc695f5d4b3ec147f766bffa165ddd317470fbe05c86d0a9c3e04.
//
// Solidity: event LogRebase(uint256 indexed epoch, uint256 exchangeRate, uint256 cpi, int256 requestedSupplyAdjustment, uint256 timestampSec)
func (_Policy *PolicyFilterer) FilterLogRebase(opts *bind.FilterOpts, epoch []*big.Int) (*PolicyLogRebaseIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "LogRebase", epochRule)
	if err != nil {
		return nil, err
	}
	return &PolicyLogRebaseIterator{contract: _Policy.contract, event: "LogRebase", logs: logs, sub: sub}, nil
}

// WatchLogRebase is a free log subscription operation binding the contract event 0x41d948a7f29cc695f5d4b3ec147f766bffa165ddd317470fbe05c86d0a9c3e04.
//
// Solidity: event LogRebase(uint256 indexed epoch, uint256 exchangeRate, uint256 cpi, int256 requestedSupplyAdjustment, uint256 timestampSec)
func (_Policy *PolicyFilterer) WatchLogRebase(opts *bind.WatchOpts, sink chan<- *PolicyLogRebase, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "LogRebase", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyLogRebase)
				if err := _Policy.contract.UnpackLog(event, "LogRebase", log); err != nil {
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

// ParseLogRebase is a log parse operation binding the contract event 0x41d948a7f29cc695f5d4b3ec147f766bffa165ddd317470fbe05c86d0a9c3e04.
//
// Solidity: event LogRebase(uint256 indexed epoch, uint256 exchangeRate, uint256 cpi, int256 requestedSupplyAdjustment, uint256 timestampSec)
func (_Policy *PolicyFilterer) ParseLogRebase(log types.Log) (*PolicyLogRebase, error) {
	event := new(PolicyLogRebase)
	if err := _Policy.contract.UnpackLog(event, "LogRebase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolicyOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Policy contract.
type PolicyOwnershipRenouncedIterator struct {
	Event *PolicyOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PolicyOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyOwnershipRenounced)
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
		it.Event = new(PolicyOwnershipRenounced)
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
func (it *PolicyOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyOwnershipRenounced represents a OwnershipRenounced event raised by the Policy contract.
type PolicyOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Policy *PolicyFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PolicyOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolicyOwnershipRenouncedIterator{contract: _Policy.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Policy *PolicyFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PolicyOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyOwnershipRenounced)
				if err := _Policy.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Policy *PolicyFilterer) ParseOwnershipRenounced(log types.Log) (*PolicyOwnershipRenounced, error) {
	event := new(PolicyOwnershipRenounced)
	if err := _Policy.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolicyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Policy contract.
type PolicyOwnershipTransferredIterator struct {
	Event *PolicyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolicyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyOwnershipTransferred)
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
		it.Event = new(PolicyOwnershipTransferred)
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
func (it *PolicyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyOwnershipTransferred represents a OwnershipTransferred event raised by the Policy contract.
type PolicyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Policy *PolicyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolicyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Policy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolicyOwnershipTransferredIterator{contract: _Policy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Policy *PolicyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolicyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Policy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyOwnershipTransferred)
				if err := _Policy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Policy *PolicyFilterer) ParseOwnershipTransferred(log types.Log) (*PolicyOwnershipTransferred, error) {
	event := new(PolicyOwnershipTransferred)
	if err := _Policy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
