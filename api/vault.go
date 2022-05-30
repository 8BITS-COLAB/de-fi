// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// VaultContractMetaData contains all meta data concerning the VaultContract contract.
var VaultContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"70a08231": "balanceOf(address)",
		"27e235e3": "balances(address)",
		"9859387b": "createAccount(address)",
		"47e7ef24": "deposit(address,uint256)",
		"c4e41b22": "getTotalSupply()",
		"f3fef3a3": "withdraw(address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060006001556102e1806100256000396000f3fe6080604052600436106100555760003560e01c806327e235e31461005a57806347e7ef241461009957806370a08231146100ae5780639859387b146100e4578063c4e41b2214610119578063f3fef3a31461012e575b600080fd5b34801561006657600080fd5b5061008761007536600461021a565b60006020819052908152604090205481565b60405190815260200160405180910390f35b6100ac6100a736600461023c565b61014e565b005b3480156100ba57600080fd5b506100876100c936600461021a565b6001600160a01b031660009081526020819052604090205490565b3480156100f057600080fd5b506100ac6100ff36600461021a565b6001600160a01b0316600090815260208190526040812055565b34801561012557600080fd5b50600154610087565b34801561013a57600080fd5b506100ac61014936600461023c565b610198565b6001600160a01b0382166000908152602081905260408120805483929061017690849061027c565b92505081905550806001600082825461018f919061027c565b90915550505050565b6001600160a01b0382166000908152602081905260409020548111156101bd57600080fd5b6001600160a01b038216600090815260208190526040812080548392906101e5908490610294565b92505081905550806001600082825461018f9190610294565b80356001600160a01b038116811461021557600080fd5b919050565b60006020828403121561022c57600080fd5b610235826101fe565b9392505050565b6000806040838503121561024f57600080fd5b610258836101fe565b946020939093013593505050565b634e487b7160e01b600052601160045260246000fd5b6000821982111561028f5761028f610266565b500190565b6000828210156102a6576102a6610266565b50039056fea264697066735822122013763048573edd4bbe03fe86fac23cd4feb83a0b0b77633957600043871b65d664736f6c634300080d0033",
}

// VaultContractABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultContractMetaData.ABI instead.
var VaultContractABI = VaultContractMetaData.ABI

// Deprecated: Use VaultContractMetaData.Sigs instead.
// VaultContractFuncSigs maps the 4-byte function signature to its string representation.
var VaultContractFuncSigs = VaultContractMetaData.Sigs

// VaultContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VaultContractMetaData.Bin instead.
var VaultContractBin = VaultContractMetaData.Bin

// DeployVaultContract deploys a new Ethereum contract, binding an instance of VaultContract to it.
func DeployVaultContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VaultContract, error) {
	parsed, err := VaultContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VaultContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaultContract{VaultContractCaller: VaultContractCaller{contract: contract}, VaultContractTransactor: VaultContractTransactor{contract: contract}, VaultContractFilterer: VaultContractFilterer{contract: contract}}, nil
}

// VaultContract is an auto generated Go binding around an Ethereum contract.
type VaultContract struct {
	VaultContractCaller     // Read-only binding to the contract
	VaultContractTransactor // Write-only binding to the contract
	VaultContractFilterer   // Log filterer for contract events
}

// VaultContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultContractSession struct {
	Contract     *VaultContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultContractCallerSession struct {
	Contract *VaultContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VaultContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultContractTransactorSession struct {
	Contract     *VaultContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VaultContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultContractRaw struct {
	Contract *VaultContract // Generic contract binding to access the raw methods on
}

// VaultContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultContractCallerRaw struct {
	Contract *VaultContractCaller // Generic read-only contract binding to access the raw methods on
}

// VaultContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultContractTransactorRaw struct {
	Contract *VaultContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultContract creates a new instance of VaultContract, bound to a specific deployed contract.
func NewVaultContract(address common.Address, backend bind.ContractBackend) (*VaultContract, error) {
	contract, err := bindVaultContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultContract{VaultContractCaller: VaultContractCaller{contract: contract}, VaultContractTransactor: VaultContractTransactor{contract: contract}, VaultContractFilterer: VaultContractFilterer{contract: contract}}, nil
}

// NewVaultContractCaller creates a new read-only instance of VaultContract, bound to a specific deployed contract.
func NewVaultContractCaller(address common.Address, caller bind.ContractCaller) (*VaultContractCaller, error) {
	contract, err := bindVaultContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultContractCaller{contract: contract}, nil
}

// NewVaultContractTransactor creates a new write-only instance of VaultContract, bound to a specific deployed contract.
func NewVaultContractTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultContractTransactor, error) {
	contract, err := bindVaultContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultContractTransactor{contract: contract}, nil
}

// NewVaultContractFilterer creates a new log filterer instance of VaultContract, bound to a specific deployed contract.
func NewVaultContractFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultContractFilterer, error) {
	contract, err := bindVaultContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultContractFilterer{contract: contract}, nil
}

// bindVaultContract binds a generic wrapper to an already deployed contract.
func bindVaultContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultContract *VaultContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultContract.Contract.VaultContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultContract *VaultContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultContract.Contract.VaultContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultContract *VaultContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultContract.Contract.VaultContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultContract *VaultContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultContract *VaultContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultContract *VaultContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_VaultContract *VaultContractCaller) BalanceOf(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "balanceOf", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_VaultContract *VaultContractSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _VaultContract.Contract.BalanceOf(&_VaultContract.CallOpts, _addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_VaultContract *VaultContractCallerSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _VaultContract.Contract.BalanceOf(&_VaultContract.CallOpts, _addr)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_VaultContract *VaultContractCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_VaultContract *VaultContractSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _VaultContract.Contract.Balances(&_VaultContract.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_VaultContract *VaultContractCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _VaultContract.Contract.Balances(&_VaultContract.CallOpts, arg0)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_VaultContract *VaultContractCaller) GetTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultContract.contract.Call(opts, &out, "getTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_VaultContract *VaultContractSession) GetTotalSupply() (*big.Int, error) {
	return _VaultContract.Contract.GetTotalSupply(&_VaultContract.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_VaultContract *VaultContractCallerSession) GetTotalSupply() (*big.Int, error) {
	return _VaultContract.Contract.GetTotalSupply(&_VaultContract.CallOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9859387b.
//
// Solidity: function createAccount(address _owner) returns()
func (_VaultContract *VaultContractTransactor) CreateAccount(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "createAccount", _owner)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9859387b.
//
// Solidity: function createAccount(address _owner) returns()
func (_VaultContract *VaultContractSession) CreateAccount(_owner common.Address) (*types.Transaction, error) {
	return _VaultContract.Contract.CreateAccount(&_VaultContract.TransactOpts, _owner)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9859387b.
//
// Solidity: function createAccount(address _owner) returns()
func (_VaultContract *VaultContractTransactorSession) CreateAccount(_owner common.Address) (*types.Transaction, error) {
	return _VaultContract.Contract.CreateAccount(&_VaultContract.TransactOpts, _owner)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _addr, uint256 _amount) payable returns()
func (_VaultContract *VaultContractTransactor) Deposit(opts *bind.TransactOpts, _addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "deposit", _addr, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _addr, uint256 _amount) payable returns()
func (_VaultContract *VaultContractSession) Deposit(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.Deposit(&_VaultContract.TransactOpts, _addr, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _addr, uint256 _amount) payable returns()
func (_VaultContract *VaultContractTransactorSession) Deposit(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.Deposit(&_VaultContract.TransactOpts, _addr, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _addr, uint256 _amount) returns()
func (_VaultContract *VaultContractTransactor) Withdraw(opts *bind.TransactOpts, _addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "withdraw", _addr, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _addr, uint256 _amount) returns()
func (_VaultContract *VaultContractSession) Withdraw(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.Withdraw(&_VaultContract.TransactOpts, _addr, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _addr, uint256 _amount) returns()
func (_VaultContract *VaultContractTransactorSession) Withdraw(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.Withdraw(&_VaultContract.TransactOpts, _addr, _amount)
}
