// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenvault

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"DelayPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"DelayThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DelayedTransferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelayedTransferExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"depositId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"mintChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"EpochLengthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"EpochVolumeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MaxDepositUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinDepositUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"refChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"burnAccount\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delayThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delayedTransfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_mintChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_mintAccount\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_mintChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_mintAccount\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"name\":\"depositNative\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumeCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"executeDelayedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastOpTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeWrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setDelayPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setDelayThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setEpochLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"name\":\"setEpochVolumeCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMaxDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"name\":\"setWrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Contracts *ContractsCaller) DelayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "delayPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Contracts *ContractsSession) DelayPeriod() (*big.Int, error) {
	return _Contracts.Contract.DelayPeriod(&_Contracts.CallOpts)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Contracts *ContractsCallerSession) DelayPeriod() (*big.Int, error) {
	return _Contracts.Contract.DelayPeriod(&_Contracts.CallOpts)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Contracts *ContractsCaller) DelayThresholds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "delayThresholds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Contracts *ContractsSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.DelayThresholds(&_Contracts.CallOpts, arg0)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.DelayThresholds(&_Contracts.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_Contracts *ContractsCaller) DelayedTransfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "delayedTransfers", arg0)

	outstruct := new(struct {
		Receiver  common.Address
		Token     common.Address
		Amount    *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_Contracts *ContractsSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Contracts.Contract.DelayedTransfers(&_Contracts.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_Contracts *ContractsCallerSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Contracts.Contract.DelayedTransfers(&_Contracts.CallOpts, arg0)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Contracts *ContractsCaller) EpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "epochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Contracts *ContractsSession) EpochLength() (*big.Int, error) {
	return _Contracts.Contract.EpochLength(&_Contracts.CallOpts)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Contracts *ContractsCallerSession) EpochLength() (*big.Int, error) {
	return _Contracts.Contract.EpochLength(&_Contracts.CallOpts)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Contracts *ContractsCaller) EpochVolumeCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "epochVolumeCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Contracts *ContractsSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.EpochVolumeCaps(&_Contracts.CallOpts, arg0)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.EpochVolumeCaps(&_Contracts.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Contracts *ContractsCaller) EpochVolumes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "epochVolumes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Contracts *ContractsSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.EpochVolumes(&_Contracts.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.EpochVolumes(&_Contracts.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Contracts *ContractsCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Contracts *ContractsSession) Governors(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Governors(&_Contracts.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Contracts *ContractsCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Governors(&_Contracts.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Contracts *ContractsCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Contracts *ContractsSession) IsGovernor(_account common.Address) (bool, error) {
	return _Contracts.Contract.IsGovernor(&_Contracts.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Contracts *ContractsCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _Contracts.Contract.IsGovernor(&_Contracts.CallOpts, _account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Contracts *ContractsCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Contracts *ContractsSession) IsPauser(account common.Address) (bool, error) {
	return _Contracts.Contract.IsPauser(&_Contracts.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Contracts *ContractsCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Contracts.Contract.IsPauser(&_Contracts.CallOpts, account)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Contracts *ContractsCaller) LastOpTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "lastOpTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Contracts *ContractsSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.LastOpTimestamps(&_Contracts.CallOpts, arg0)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.LastOpTimestamps(&_Contracts.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Contracts *ContractsCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Contracts *ContractsSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.MaxDeposit(&_Contracts.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.MaxDeposit(&_Contracts.CallOpts, arg0)
}

// MinDeposit is a free data retrieval call binding the contract method 0x3c29f839.
//
// Solidity: function minDeposit(address ) view returns(uint256)
func (_Contracts *ContractsCaller) MinDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "minDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x3c29f839.
//
// Solidity: function minDeposit(address ) view returns(uint256)
func (_Contracts *ContractsSession) MinDeposit(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.MinDeposit(&_Contracts.CallOpts, arg0)
}

// MinDeposit is a free data retrieval call binding the contract method 0x3c29f839.
//
// Solidity: function minDeposit(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) MinDeposit(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.MinDeposit(&_Contracts.CallOpts, arg0)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Contracts *ContractsCaller) NativeWrap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "nativeWrap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Contracts *ContractsSession) NativeWrap() (common.Address, error) {
	return _Contracts.Contract.NativeWrap(&_Contracts.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Contracts *ContractsCallerSession) NativeWrap() (common.Address, error) {
	return _Contracts.Contract.NativeWrap(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contracts *ContractsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contracts *ContractsSession) Paused() (bool, error) {
	return _Contracts.Contract.Paused(&_Contracts.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contracts *ContractsCallerSession) Paused() (bool, error) {
	return _Contracts.Contract.Paused(&_Contracts.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Contracts *ContractsCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Contracts *ContractsSession) Pausers(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Pausers(&_Contracts.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Contracts *ContractsCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Pausers(&_Contracts.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(uint256)
func (_Contracts *ContractsCaller) Records(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "records", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(uint256)
func (_Contracts *ContractsSession) Records(arg0 [32]byte) (*big.Int, error) {
	return _Contracts.Contract.Records(&_Contracts.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) view returns(uint256)
func (_Contracts *ContractsCallerSession) Records(arg0 [32]byte) (*big.Int, error) {
	return _Contracts.Contract.Records(&_Contracts.CallOpts, arg0)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_Contracts *ContractsCaller) SigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "sigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_Contracts *ContractsSession) SigsVerifier() (common.Address, error) {
	return _Contracts.Contract.SigsVerifier(&_Contracts.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_Contracts *ContractsCallerSession) SigsVerifier() (common.Address, error) {
	return _Contracts.Contract.SigsVerifier(&_Contracts.CallOpts)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Contracts *ContractsTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Contracts *ContractsSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddGovernor(&_Contracts.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Contracts *ContractsTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddGovernor(&_Contracts.TransactOpts, _account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Contracts *ContractsTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Contracts *ContractsSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddPauser(&_Contracts.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Contracts *ContractsTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddPauser(&_Contracts.TransactOpts, account)
}

// Deposit is a paid mutator transaction binding the contract method 0x23463624.
//
// Solidity: function deposit(address _token, uint256 _amount, uint64 _mintChainId, address _mintAccount, uint64 _nonce) returns(bytes32)
func (_Contracts *ContractsTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _mintChainId uint64, _mintAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deposit", _token, _amount, _mintChainId, _mintAccount, _nonce)
}

// Deposit is a paid mutator transaction binding the contract method 0x23463624.
//
// Solidity: function deposit(address _token, uint256 _amount, uint64 _mintChainId, address _mintAccount, uint64 _nonce) returns(bytes32)
func (_Contracts *ContractsSession) Deposit(_token common.Address, _amount *big.Int, _mintChainId uint64, _mintAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _Contracts.Contract.Deposit(&_Contracts.TransactOpts, _token, _amount, _mintChainId, _mintAccount, _nonce)
}

// Deposit is a paid mutator transaction binding the contract method 0x23463624.
//
// Solidity: function deposit(address _token, uint256 _amount, uint64 _mintChainId, address _mintAccount, uint64 _nonce) returns(bytes32)
func (_Contracts *ContractsTransactorSession) Deposit(_token common.Address, _amount *big.Int, _mintChainId uint64, _mintAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _Contracts.Contract.Deposit(&_Contracts.TransactOpts, _token, _amount, _mintChainId, _mintAccount, _nonce)
}

// DepositNative is a paid mutator transaction binding the contract method 0x00a95fd7.
//
// Solidity: function depositNative(uint256 _amount, uint64 _mintChainId, address _mintAccount, uint64 _nonce) payable returns(bytes32)
func (_Contracts *ContractsTransactor) DepositNative(opts *bind.TransactOpts, _amount *big.Int, _mintChainId uint64, _mintAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "depositNative", _amount, _mintChainId, _mintAccount, _nonce)
}

// DepositNative is a paid mutator transaction binding the contract method 0x00a95fd7.
//
// Solidity: function depositNative(uint256 _amount, uint64 _mintChainId, address _mintAccount, uint64 _nonce) payable returns(bytes32)
func (_Contracts *ContractsSession) DepositNative(_amount *big.Int, _mintChainId uint64, _mintAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _Contracts.Contract.DepositNative(&_Contracts.TransactOpts, _amount, _mintChainId, _mintAccount, _nonce)
}

// DepositNative is a paid mutator transaction binding the contract method 0x00a95fd7.
//
// Solidity: function depositNative(uint256 _amount, uint64 _mintChainId, address _mintAccount, uint64 _nonce) payable returns(bytes32)
func (_Contracts *ContractsTransactorSession) DepositNative(_amount *big.Int, _mintChainId uint64, _mintAccount common.Address, _nonce uint64) (*types.Transaction, error) {
	return _Contracts.Contract.DepositNative(&_Contracts.TransactOpts, _amount, _mintChainId, _mintAccount, _nonce)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Contracts *ContractsTransactor) ExecuteDelayedTransfer(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "executeDelayedTransfer", id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Contracts *ContractsSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.ExecuteDelayedTransfer(&_Contracts.TransactOpts, id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Contracts *ContractsTransactorSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.ExecuteDelayedTransfer(&_Contracts.TransactOpts, id)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Contracts *ContractsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Contracts *ContractsSession) Pause() (*types.Transaction, error) {
	return _Contracts.Contract.Pause(&_Contracts.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Contracts *ContractsTransactorSession) Pause() (*types.Transaction, error) {
	return _Contracts.Contract.Pause(&_Contracts.TransactOpts)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Contracts *ContractsTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Contracts *ContractsSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveGovernor(&_Contracts.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Contracts *ContractsTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveGovernor(&_Contracts.TransactOpts, _account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Contracts *ContractsTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Contracts *ContractsSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemovePauser(&_Contracts.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Contracts *ContractsTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemovePauser(&_Contracts.TransactOpts, account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Contracts *ContractsTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Contracts *ContractsSession) RenounceGovernor() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceGovernor(&_Contracts.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Contracts *ContractsTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceGovernor(&_Contracts.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Contracts *ContractsTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Contracts *ContractsSession) RenouncePauser() (*types.Transaction, error) {
	return _Contracts.Contract.RenouncePauser(&_Contracts.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Contracts *ContractsTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Contracts.Contract.RenouncePauser(&_Contracts.TransactOpts)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Contracts *ContractsTransactor) SetDelayPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setDelayPeriod", _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Contracts *ContractsSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetDelayPeriod(&_Contracts.TransactOpts, _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Contracts *ContractsTransactorSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetDelayPeriod(&_Contracts.TransactOpts, _period)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Contracts *ContractsTransactor) SetDelayThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setDelayThresholds", _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Contracts *ContractsSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetDelayThresholds(&_Contracts.TransactOpts, _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Contracts *ContractsTransactorSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetDelayThresholds(&_Contracts.TransactOpts, _tokens, _thresholds)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Contracts *ContractsTransactor) SetEpochLength(opts *bind.TransactOpts, _length *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setEpochLength", _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Contracts *ContractsSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetEpochLength(&_Contracts.TransactOpts, _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Contracts *ContractsTransactorSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetEpochLength(&_Contracts.TransactOpts, _length)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Contracts *ContractsTransactor) SetEpochVolumeCaps(opts *bind.TransactOpts, _tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setEpochVolumeCaps", _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Contracts *ContractsSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetEpochVolumeCaps(&_Contracts.TransactOpts, _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Contracts *ContractsTransactorSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetEpochVolumeCaps(&_Contracts.TransactOpts, _tokens, _caps)
}

// SetMaxDeposit is a paid mutator transaction binding the contract method 0x303b6442.
//
// Solidity: function setMaxDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_Contracts *ContractsTransactor) SetMaxDeposit(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMaxDeposit", _tokens, _amounts)
}

// SetMaxDeposit is a paid mutator transaction binding the contract method 0x303b6442.
//
// Solidity: function setMaxDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_Contracts *ContractsSession) SetMaxDeposit(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMaxDeposit(&_Contracts.TransactOpts, _tokens, _amounts)
}

// SetMaxDeposit is a paid mutator transaction binding the contract method 0x303b6442.
//
// Solidity: function setMaxDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_Contracts *ContractsTransactorSession) SetMaxDeposit(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMaxDeposit(&_Contracts.TransactOpts, _tokens, _amounts)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x5ec2fa26.
//
// Solidity: function setMinDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_Contracts *ContractsTransactor) SetMinDeposit(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMinDeposit", _tokens, _amounts)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x5ec2fa26.
//
// Solidity: function setMinDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_Contracts *ContractsSession) SetMinDeposit(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinDeposit(&_Contracts.TransactOpts, _tokens, _amounts)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x5ec2fa26.
//
// Solidity: function setMinDeposit(address[] _tokens, uint256[] _amounts) returns()
func (_Contracts *ContractsTransactorSession) SetMinDeposit(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetMinDeposit(&_Contracts.TransactOpts, _tokens, _amounts)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Contracts *ContractsTransactor) SetWrap(opts *bind.TransactOpts, _weth common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setWrap", _weth)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Contracts *ContractsSession) SetWrap(_weth common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetWrap(&_Contracts.TransactOpts, _weth)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Contracts *ContractsTransactorSession) SetWrap(_weth common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetWrap(&_Contracts.TransactOpts, _weth)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Contracts *ContractsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Contracts *ContractsSession) Unpause() (*types.Transaction, error) {
	return _Contracts.Contract.Unpause(&_Contracts.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Contracts *ContractsTransactorSession) Unpause() (*types.Transaction, error) {
	return _Contracts.Contract.Unpause(&_Contracts.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _request, bytes[] _sigs, address[] _signers, uint256[] _powers) returns(bytes32)
func (_Contracts *ContractsTransactor) Withdraw(opts *bind.TransactOpts, _request []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdraw", _request, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _request, bytes[] _sigs, address[] _signers, uint256[] _powers) returns(bytes32)
func (_Contracts *ContractsSession) Withdraw(_request []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, _request, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _request, bytes[] _sigs, address[] _signers, uint256[] _powers) returns(bytes32)
func (_Contracts *ContractsTransactorSession) Withdraw(_request []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, _request, _sigs, _signers, _powers)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactorSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// ContractsDelayPeriodUpdatedIterator is returned from FilterDelayPeriodUpdated and is used to iterate over the raw logs and unpacked data for DelayPeriodUpdated events raised by the Contracts contract.
type ContractsDelayPeriodUpdatedIterator struct {
	Event *ContractsDelayPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsDelayPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDelayPeriodUpdated)
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
		it.Event = new(ContractsDelayPeriodUpdated)
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
func (it *ContractsDelayPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDelayPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDelayPeriodUpdated represents a DelayPeriodUpdated event raised by the Contracts contract.
type ContractsDelayPeriodUpdated struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelayPeriodUpdated is a free log retrieval operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_Contracts *ContractsFilterer) FilterDelayPeriodUpdated(opts *bind.FilterOpts) (*ContractsDelayPeriodUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsDelayPeriodUpdatedIterator{contract: _Contracts.contract, event: "DelayPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayPeriodUpdated is a free log subscription operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_Contracts *ContractsFilterer) WatchDelayPeriodUpdated(opts *bind.WatchOpts, sink chan<- *ContractsDelayPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDelayPeriodUpdated)
				if err := _Contracts.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
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

// ParseDelayPeriodUpdated is a log parse operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_Contracts *ContractsFilterer) ParseDelayPeriodUpdated(log types.Log) (*ContractsDelayPeriodUpdated, error) {
	event := new(ContractsDelayPeriodUpdated)
	if err := _Contracts.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDelayThresholdUpdatedIterator is returned from FilterDelayThresholdUpdated and is used to iterate over the raw logs and unpacked data for DelayThresholdUpdated events raised by the Contracts contract.
type ContractsDelayThresholdUpdatedIterator struct {
	Event *ContractsDelayThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsDelayThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDelayThresholdUpdated)
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
		it.Event = new(ContractsDelayThresholdUpdated)
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
func (it *ContractsDelayThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDelayThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDelayThresholdUpdated represents a DelayThresholdUpdated event raised by the Contracts contract.
type ContractsDelayThresholdUpdated struct {
	Token     common.Address
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayThresholdUpdated is a free log retrieval operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_Contracts *ContractsFilterer) FilterDelayThresholdUpdated(opts *bind.FilterOpts) (*ContractsDelayThresholdUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsDelayThresholdUpdatedIterator{contract: _Contracts.contract, event: "DelayThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayThresholdUpdated is a free log subscription operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_Contracts *ContractsFilterer) WatchDelayThresholdUpdated(opts *bind.WatchOpts, sink chan<- *ContractsDelayThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDelayThresholdUpdated)
				if err := _Contracts.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
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

// ParseDelayThresholdUpdated is a log parse operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_Contracts *ContractsFilterer) ParseDelayThresholdUpdated(log types.Log) (*ContractsDelayThresholdUpdated, error) {
	event := new(ContractsDelayThresholdUpdated)
	if err := _Contracts.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDelayedTransferAddedIterator is returned from FilterDelayedTransferAdded and is used to iterate over the raw logs and unpacked data for DelayedTransferAdded events raised by the Contracts contract.
type ContractsDelayedTransferAddedIterator struct {
	Event *ContractsDelayedTransferAdded // Event containing the contract specifics and raw log

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
func (it *ContractsDelayedTransferAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDelayedTransferAdded)
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
		it.Event = new(ContractsDelayedTransferAdded)
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
func (it *ContractsDelayedTransferAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDelayedTransferAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDelayedTransferAdded represents a DelayedTransferAdded event raised by the Contracts contract.
type ContractsDelayedTransferAdded struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferAdded is a free log retrieval operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_Contracts *ContractsFilterer) FilterDelayedTransferAdded(opts *bind.FilterOpts) (*ContractsDelayedTransferAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsDelayedTransferAddedIterator{contract: _Contracts.contract, event: "DelayedTransferAdded", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferAdded is a free log subscription operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_Contracts *ContractsFilterer) WatchDelayedTransferAdded(opts *bind.WatchOpts, sink chan<- *ContractsDelayedTransferAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDelayedTransferAdded)
				if err := _Contracts.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
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

// ParseDelayedTransferAdded is a log parse operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_Contracts *ContractsFilterer) ParseDelayedTransferAdded(log types.Log) (*ContractsDelayedTransferAdded, error) {
	event := new(ContractsDelayedTransferAdded)
	if err := _Contracts.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDelayedTransferExecutedIterator is returned from FilterDelayedTransferExecuted and is used to iterate over the raw logs and unpacked data for DelayedTransferExecuted events raised by the Contracts contract.
type ContractsDelayedTransferExecutedIterator struct {
	Event *ContractsDelayedTransferExecuted // Event containing the contract specifics and raw log

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
func (it *ContractsDelayedTransferExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDelayedTransferExecuted)
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
		it.Event = new(ContractsDelayedTransferExecuted)
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
func (it *ContractsDelayedTransferExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDelayedTransferExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDelayedTransferExecuted represents a DelayedTransferExecuted event raised by the Contracts contract.
type ContractsDelayedTransferExecuted struct {
	Id       [32]byte
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferExecuted is a free log retrieval operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_Contracts *ContractsFilterer) FilterDelayedTransferExecuted(opts *bind.FilterOpts) (*ContractsDelayedTransferExecutedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return &ContractsDelayedTransferExecutedIterator{contract: _Contracts.contract, event: "DelayedTransferExecuted", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferExecuted is a free log subscription operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_Contracts *ContractsFilterer) WatchDelayedTransferExecuted(opts *bind.WatchOpts, sink chan<- *ContractsDelayedTransferExecuted) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDelayedTransferExecuted)
				if err := _Contracts.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
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

// ParseDelayedTransferExecuted is a log parse operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_Contracts *ContractsFilterer) ParseDelayedTransferExecuted(log types.Log) (*ContractsDelayedTransferExecuted, error) {
	event := new(ContractsDelayedTransferExecuted)
	if err := _Contracts.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Contracts contract.
type ContractsDepositedIterator struct {
	Event *ContractsDeposited // Event containing the contract specifics and raw log

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
func (it *ContractsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDeposited)
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
		it.Event = new(ContractsDeposited)
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
func (it *ContractsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDeposited represents a Deposited event raised by the Contracts contract.
type ContractsDeposited struct {
	DepositId   [32]byte
	Depositor   common.Address
	Token       common.Address
	Amount      *big.Int
	MintChainId uint64
	MintAccount common.Address
	Nonce       uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x28d226819e371600e26624ebc4a9a3947117ee2760209f816c789d3a99bf481b.
//
// Solidity: event Deposited(bytes32 depositId, address depositor, address token, uint256 amount, uint64 mintChainId, address mintAccount, uint64 nonce)
func (_Contracts *ContractsFilterer) FilterDeposited(opts *bind.FilterOpts) (*ContractsDepositedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &ContractsDepositedIterator{contract: _Contracts.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x28d226819e371600e26624ebc4a9a3947117ee2760209f816c789d3a99bf481b.
//
// Solidity: event Deposited(bytes32 depositId, address depositor, address token, uint256 amount, uint64 mintChainId, address mintAccount, uint64 nonce)
func (_Contracts *ContractsFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ContractsDeposited) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDeposited)
				if err := _Contracts.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x28d226819e371600e26624ebc4a9a3947117ee2760209f816c789d3a99bf481b.
//
// Solidity: event Deposited(bytes32 depositId, address depositor, address token, uint256 amount, uint64 mintChainId, address mintAccount, uint64 nonce)
func (_Contracts *ContractsFilterer) ParseDeposited(log types.Log) (*ContractsDeposited, error) {
	event := new(ContractsDeposited)
	if err := _Contracts.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsEpochLengthUpdatedIterator is returned from FilterEpochLengthUpdated and is used to iterate over the raw logs and unpacked data for EpochLengthUpdated events raised by the Contracts contract.
type ContractsEpochLengthUpdatedIterator struct {
	Event *ContractsEpochLengthUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsEpochLengthUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEpochLengthUpdated)
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
		it.Event = new(ContractsEpochLengthUpdated)
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
func (it *ContractsEpochLengthUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEpochLengthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEpochLengthUpdated represents a EpochLengthUpdated event raised by the Contracts contract.
type ContractsEpochLengthUpdated struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEpochLengthUpdated is a free log retrieval operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_Contracts *ContractsFilterer) FilterEpochLengthUpdated(opts *bind.FilterOpts) (*ContractsEpochLengthUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsEpochLengthUpdatedIterator{contract: _Contracts.contract, event: "EpochLengthUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochLengthUpdated is a free log subscription operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_Contracts *ContractsFilterer) WatchEpochLengthUpdated(opts *bind.WatchOpts, sink chan<- *ContractsEpochLengthUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEpochLengthUpdated)
				if err := _Contracts.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
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

// ParseEpochLengthUpdated is a log parse operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_Contracts *ContractsFilterer) ParseEpochLengthUpdated(log types.Log) (*ContractsEpochLengthUpdated, error) {
	event := new(ContractsEpochLengthUpdated)
	if err := _Contracts.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsEpochVolumeUpdatedIterator is returned from FilterEpochVolumeUpdated and is used to iterate over the raw logs and unpacked data for EpochVolumeUpdated events raised by the Contracts contract.
type ContractsEpochVolumeUpdatedIterator struct {
	Event *ContractsEpochVolumeUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsEpochVolumeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEpochVolumeUpdated)
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
		it.Event = new(ContractsEpochVolumeUpdated)
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
func (it *ContractsEpochVolumeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEpochVolumeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEpochVolumeUpdated represents a EpochVolumeUpdated event raised by the Contracts contract.
type ContractsEpochVolumeUpdated struct {
	Token common.Address
	Cap   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEpochVolumeUpdated is a free log retrieval operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_Contracts *ContractsFilterer) FilterEpochVolumeUpdated(opts *bind.FilterOpts) (*ContractsEpochVolumeUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsEpochVolumeUpdatedIterator{contract: _Contracts.contract, event: "EpochVolumeUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochVolumeUpdated is a free log subscription operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_Contracts *ContractsFilterer) WatchEpochVolumeUpdated(opts *bind.WatchOpts, sink chan<- *ContractsEpochVolumeUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEpochVolumeUpdated)
				if err := _Contracts.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
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

// ParseEpochVolumeUpdated is a log parse operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_Contracts *ContractsFilterer) ParseEpochVolumeUpdated(log types.Log) (*ContractsEpochVolumeUpdated, error) {
	event := new(ContractsEpochVolumeUpdated)
	if err := _Contracts.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the Contracts contract.
type ContractsGovernorAddedIterator struct {
	Event *ContractsGovernorAdded // Event containing the contract specifics and raw log

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
func (it *ContractsGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsGovernorAdded)
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
		it.Event = new(ContractsGovernorAdded)
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
func (it *ContractsGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsGovernorAdded represents a GovernorAdded event raised by the Contracts contract.
type ContractsGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Contracts *ContractsFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*ContractsGovernorAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsGovernorAddedIterator{contract: _Contracts.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Contracts *ContractsFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *ContractsGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsGovernorAdded)
				if err := _Contracts.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Contracts *ContractsFilterer) ParseGovernorAdded(log types.Log) (*ContractsGovernorAdded, error) {
	event := new(ContractsGovernorAdded)
	if err := _Contracts.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the Contracts contract.
type ContractsGovernorRemovedIterator struct {
	Event *ContractsGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *ContractsGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsGovernorRemoved)
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
		it.Event = new(ContractsGovernorRemoved)
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
func (it *ContractsGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsGovernorRemoved represents a GovernorRemoved event raised by the Contracts contract.
type ContractsGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Contracts *ContractsFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*ContractsGovernorRemovedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &ContractsGovernorRemovedIterator{contract: _Contracts.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Contracts *ContractsFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *ContractsGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsGovernorRemoved)
				if err := _Contracts.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Contracts *ContractsFilterer) ParseGovernorRemoved(log types.Log) (*ContractsGovernorRemoved, error) {
	event := new(ContractsGovernorRemoved)
	if err := _Contracts.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMaxDepositUpdatedIterator is returned from FilterMaxDepositUpdated and is used to iterate over the raw logs and unpacked data for MaxDepositUpdated events raised by the Contracts contract.
type ContractsMaxDepositUpdatedIterator struct {
	Event *ContractsMaxDepositUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsMaxDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMaxDepositUpdated)
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
		it.Event = new(ContractsMaxDepositUpdated)
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
func (it *ContractsMaxDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMaxDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMaxDepositUpdated represents a MaxDepositUpdated event raised by the Contracts contract.
type ContractsMaxDepositUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxDepositUpdated is a free log retrieval operation binding the contract event 0x0e5d348f9737ccc8b4cf0eea0ccf3670af071af8bea5d64664f10e700c08de72.
//
// Solidity: event MaxDepositUpdated(address token, uint256 amount)
func (_Contracts *ContractsFilterer) FilterMaxDepositUpdated(opts *bind.FilterOpts) (*ContractsMaxDepositUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MaxDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsMaxDepositUpdatedIterator{contract: _Contracts.contract, event: "MaxDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxDepositUpdated is a free log subscription operation binding the contract event 0x0e5d348f9737ccc8b4cf0eea0ccf3670af071af8bea5d64664f10e700c08de72.
//
// Solidity: event MaxDepositUpdated(address token, uint256 amount)
func (_Contracts *ContractsFilterer) WatchMaxDepositUpdated(opts *bind.WatchOpts, sink chan<- *ContractsMaxDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MaxDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMaxDepositUpdated)
				if err := _Contracts.contract.UnpackLog(event, "MaxDepositUpdated", log); err != nil {
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

// ParseMaxDepositUpdated is a log parse operation binding the contract event 0x0e5d348f9737ccc8b4cf0eea0ccf3670af071af8bea5d64664f10e700c08de72.
//
// Solidity: event MaxDepositUpdated(address token, uint256 amount)
func (_Contracts *ContractsFilterer) ParseMaxDepositUpdated(log types.Log) (*ContractsMaxDepositUpdated, error) {
	event := new(ContractsMaxDepositUpdated)
	if err := _Contracts.contract.UnpackLog(event, "MaxDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsMinDepositUpdatedIterator is returned from FilterMinDepositUpdated and is used to iterate over the raw logs and unpacked data for MinDepositUpdated events raised by the Contracts contract.
type ContractsMinDepositUpdatedIterator struct {
	Event *ContractsMinDepositUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsMinDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsMinDepositUpdated)
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
		it.Event = new(ContractsMinDepositUpdated)
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
func (it *ContractsMinDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsMinDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsMinDepositUpdated represents a MinDepositUpdated event raised by the Contracts contract.
type ContractsMinDepositUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinDepositUpdated is a free log retrieval operation binding the contract event 0x0f48d517989455cd80ed52427e80553e66f9b69fd5cee8e26bd1a1f9c364fba6.
//
// Solidity: event MinDepositUpdated(address token, uint256 amount)
func (_Contracts *ContractsFilterer) FilterMinDepositUpdated(opts *bind.FilterOpts) (*ContractsMinDepositUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "MinDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsMinDepositUpdatedIterator{contract: _Contracts.contract, event: "MinDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMinDepositUpdated is a free log subscription operation binding the contract event 0x0f48d517989455cd80ed52427e80553e66f9b69fd5cee8e26bd1a1f9c364fba6.
//
// Solidity: event MinDepositUpdated(address token, uint256 amount)
func (_Contracts *ContractsFilterer) WatchMinDepositUpdated(opts *bind.WatchOpts, sink chan<- *ContractsMinDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "MinDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsMinDepositUpdated)
				if err := _Contracts.contract.UnpackLog(event, "MinDepositUpdated", log); err != nil {
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

// ParseMinDepositUpdated is a log parse operation binding the contract event 0x0f48d517989455cd80ed52427e80553e66f9b69fd5cee8e26bd1a1f9c364fba6.
//
// Solidity: event MinDepositUpdated(address token, uint256 amount)
func (_Contracts *ContractsFilterer) ParseMinDepositUpdated(log types.Log) (*ContractsMinDepositUpdated, error) {
	event := new(ContractsMinDepositUpdated)
	if err := _Contracts.contract.UnpackLog(event, "MinDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Contracts contract.
type ContractsPausedIterator struct {
	Event *ContractsPaused // Event containing the contract specifics and raw log

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
func (it *ContractsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPaused)
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
		it.Event = new(ContractsPaused)
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
func (it *ContractsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPaused represents a Paused event raised by the Contracts contract.
type ContractsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Contracts *ContractsFilterer) FilterPaused(opts *bind.FilterOpts) (*ContractsPausedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractsPausedIterator{contract: _Contracts.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Contracts *ContractsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractsPaused) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPaused)
				if err := _Contracts.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Contracts *ContractsFilterer) ParsePaused(log types.Log) (*ContractsPaused, error) {
	event := new(ContractsPaused)
	if err := _Contracts.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Contracts contract.
type ContractsPauserAddedIterator struct {
	Event *ContractsPauserAdded // Event containing the contract specifics and raw log

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
func (it *ContractsPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPauserAdded)
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
		it.Event = new(ContractsPauserAdded)
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
func (it *ContractsPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPauserAdded represents a PauserAdded event raised by the Contracts contract.
type ContractsPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Contracts *ContractsFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*ContractsPauserAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsPauserAddedIterator{contract: _Contracts.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Contracts *ContractsFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *ContractsPauserAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPauserAdded)
				if err := _Contracts.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Contracts *ContractsFilterer) ParsePauserAdded(log types.Log) (*ContractsPauserAdded, error) {
	event := new(ContractsPauserAdded)
	if err := _Contracts.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Contracts contract.
type ContractsPauserRemovedIterator struct {
	Event *ContractsPauserRemoved // Event containing the contract specifics and raw log

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
func (it *ContractsPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPauserRemoved)
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
		it.Event = new(ContractsPauserRemoved)
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
func (it *ContractsPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPauserRemoved represents a PauserRemoved event raised by the Contracts contract.
type ContractsPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Contracts *ContractsFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*ContractsPauserRemovedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &ContractsPauserRemovedIterator{contract: _Contracts.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Contracts *ContractsFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *ContractsPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPauserRemoved)
				if err := _Contracts.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Contracts *ContractsFilterer) ParsePauserRemoved(log types.Log) (*ContractsPauserRemoved, error) {
	event := new(ContractsPauserRemoved)
	if err := _Contracts.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Contracts contract.
type ContractsUnpausedIterator struct {
	Event *ContractsUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUnpaused)
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
		it.Event = new(ContractsUnpaused)
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
func (it *ContractsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUnpaused represents a Unpaused event raised by the Contracts contract.
type ContractsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Contracts *ContractsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ContractsUnpausedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ContractsUnpausedIterator{contract: _Contracts.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Contracts *ContractsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractsUnpaused) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUnpaused)
				if err := _Contracts.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Contracts *ContractsFilterer) ParseUnpaused(log types.Log) (*ContractsUnpaused, error) {
	event := new(ContractsUnpaused)
	if err := _Contracts.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Contracts contract.
type ContractsWithdrawnIterator struct {
	Event *ContractsWithdrawn // Event containing the contract specifics and raw log

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
func (it *ContractsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsWithdrawn)
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
		it.Event = new(ContractsWithdrawn)
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
func (it *ContractsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsWithdrawn represents a Withdrawn event raised by the Contracts contract.
type ContractsWithdrawn struct {
	WithdrawId  [32]byte
	Receiver    common.Address
	Token       common.Address
	Amount      *big.Int
	RefChainId  uint64
	RefId       [32]byte
	BurnAccount common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x296a629c5265cb4e5319803d016902eb70a9079b89655fe2b7737821ed88beeb.
//
// Solidity: event Withdrawn(bytes32 withdrawId, address receiver, address token, uint256 amount, uint64 refChainId, bytes32 refId, address burnAccount)
func (_Contracts *ContractsFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*ContractsWithdrawnIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &ContractsWithdrawnIterator{contract: _Contracts.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x296a629c5265cb4e5319803d016902eb70a9079b89655fe2b7737821ed88beeb.
//
// Solidity: event Withdrawn(bytes32 withdrawId, address receiver, address token, uint256 amount, uint64 refChainId, bytes32 refId, address burnAccount)
func (_Contracts *ContractsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsWithdrawn)
				if err := _Contracts.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x296a629c5265cb4e5319803d016902eb70a9079b89655fe2b7737821ed88beeb.
//
// Solidity: event Withdrawn(bytes32 withdrawId, address receiver, address token, uint256 amount, uint64 refChainId, bytes32 refId, address burnAccount)
func (_Contracts *ContractsFilterer) ParseWithdrawn(log types.Log) (*ContractsWithdrawn, error) {
	event := new(ContractsWithdrawn)
	if err := _Contracts.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
