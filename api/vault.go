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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreateAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"70a08231": "balanceOf(address)",
		"27e235e3": "balances(address)",
		"5fbfb9cf": "createAccount(address,uint256)",
		"47e7ef24": "deposit(address,uint256)",
		"c4e41b22": "getTotalSupply()",
		"beabacc8": "transfer(address,address,uint256)",
		"f3fef3a3": "withdraw(address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506000600155610502806100256000396000f3fe6080604052600436106100705760003560e01c806370a082311161004e57806370a08231146100e9578063beabacc81461011f578063c4e41b221461013f578063f3fef3a31461015457600080fd5b806327e235e31461007557806347e7ef24146100b45780635fbfb9cf146100c9575b600080fd5b34801561008157600080fd5b506100a26100903660046103ff565b60006020819052908152604090205481565b60405190815260200160405180910390f35b6100c76100c2366004610421565b610174565b005b3480156100d557600080fd5b506100c76100e4366004610421565b6101fd565b3480156100f557600080fd5b506100a26101043660046103ff565b6001600160a01b031660009081526020819052604090205490565b34801561012b57600080fd5b506100c761013a36600461044b565b61026a565b34801561014b57600080fd5b506001546100a2565b34801561016057600080fd5b506100c761016f366004610421565b61033d565b6001600160a01b0382166000908152602081905260408120805483929061019c90849061049d565b9250508190555080600160008282546101b5919061049d565b90915550506040518181526001600160a01b038316907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c906020015b60405180910390a25050565b6001600160a01b03821660009081526020819052604081208290556001805483929061022a90849061049d565b90915550506040518181526001600160a01b038316907f85841522199c696c3d4a549fea06732153559ded5db5cf6dfa3bb099827f2c84906020016101f1565b6001600160a01b03831660009081526020819052604090205481111561028f57600080fd5b6001600160a01b038316600090815260208190526040812080548392906102b79084906104b5565b90915550506001600160a01b038216600090815260208190526040812080548392906102e490849061049d565b92505081905550816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161033091815260200190565b60405180910390a3505050565b6001600160a01b03821660009081526020819052604090205481111561036257600080fd5b6001600160a01b0382166000908152602081905260408120805483929061038a9084906104b5565b9250508190555080600160008282546103a391906104b5565b90915550506040518181526001600160a01b038316907f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364906020016101f1565b80356001600160a01b03811681146103fa57600080fd5b919050565b60006020828403121561041157600080fd5b61041a826103e3565b9392505050565b6000806040838503121561043457600080fd5b61043d836103e3565b946020939093013593505050565b60008060006060848603121561046057600080fd5b610469846103e3565b9250610477602085016103e3565b9150604084013590509250925092565b634e487b7160e01b600052601160045260246000fd5b600082198211156104b0576104b0610487565b500190565b6000828210156104c7576104c7610487565b50039056fea26469706673582212201e8e5eb97ffde31535bc0aa3254bbf47f90f452579dc5ec02d3fbabc253490bd64736f6c634300080d0033",
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

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address _addr, uint256 _amount) returns()
func (_VaultContract *VaultContractTransactor) CreateAccount(opts *bind.TransactOpts, _addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "createAccount", _addr, _amount)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address _addr, uint256 _amount) returns()
func (_VaultContract *VaultContractSession) CreateAccount(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.CreateAccount(&_VaultContract.TransactOpts, _addr, _amount)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address _addr, uint256 _amount) returns()
func (_VaultContract *VaultContractTransactorSession) CreateAccount(_addr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.CreateAccount(&_VaultContract.TransactOpts, _addr, _amount)
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

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _from, address _to, uint256 _amount) returns()
func (_VaultContract *VaultContractTransactor) Transfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.contract.Transact(opts, "transfer", _from, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _from, address _to, uint256 _amount) returns()
func (_VaultContract *VaultContractSession) Transfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.Transfer(&_VaultContract.TransactOpts, _from, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address _from, address _to, uint256 _amount) returns()
func (_VaultContract *VaultContractTransactorSession) Transfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultContract.Contract.Transfer(&_VaultContract.TransactOpts, _from, _to, _amount)
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

// VaultContractCreateAccountIterator is returned from FilterCreateAccount and is used to iterate over the raw logs and unpacked data for CreateAccount events raised by the VaultContract contract.
type VaultContractCreateAccountIterator struct {
	Event *VaultContractCreateAccount // Event containing the contract specifics and raw log

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
func (it *VaultContractCreateAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultContractCreateAccount)
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
		it.Event = new(VaultContractCreateAccount)
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
func (it *VaultContractCreateAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultContractCreateAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultContractCreateAccount represents a CreateAccount event raised by the VaultContract contract.
type VaultContractCreateAccount struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateAccount is a free log retrieval operation binding the contract event 0x85841522199c696c3d4a549fea06732153559ded5db5cf6dfa3bb099827f2c84.
//
// Solidity: event CreateAccount(address indexed account, uint256 amount)
func (_VaultContract *VaultContractFilterer) FilterCreateAccount(opts *bind.FilterOpts, account []common.Address) (*VaultContractCreateAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VaultContract.contract.FilterLogs(opts, "CreateAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return &VaultContractCreateAccountIterator{contract: _VaultContract.contract, event: "CreateAccount", logs: logs, sub: sub}, nil
}

// WatchCreateAccount is a free log subscription operation binding the contract event 0x85841522199c696c3d4a549fea06732153559ded5db5cf6dfa3bb099827f2c84.
//
// Solidity: event CreateAccount(address indexed account, uint256 amount)
func (_VaultContract *VaultContractFilterer) WatchCreateAccount(opts *bind.WatchOpts, sink chan<- *VaultContractCreateAccount, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VaultContract.contract.WatchLogs(opts, "CreateAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultContractCreateAccount)
				if err := _VaultContract.contract.UnpackLog(event, "CreateAccount", log); err != nil {
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

// ParseCreateAccount is a log parse operation binding the contract event 0x85841522199c696c3d4a549fea06732153559ded5db5cf6dfa3bb099827f2c84.
//
// Solidity: event CreateAccount(address indexed account, uint256 amount)
func (_VaultContract *VaultContractFilterer) ParseCreateAccount(log types.Log) (*VaultContractCreateAccount, error) {
	event := new(VaultContractCreateAccount)
	if err := _VaultContract.contract.UnpackLog(event, "CreateAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the VaultContract contract.
type VaultContractDepositIterator struct {
	Event *VaultContractDeposit // Event containing the contract specifics and raw log

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
func (it *VaultContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultContractDeposit)
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
		it.Event = new(VaultContractDeposit)
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
func (it *VaultContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultContractDeposit represents a Deposit event raised by the VaultContract contract.
type VaultContractDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_VaultContract *VaultContractFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*VaultContractDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VaultContract.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &VaultContractDepositIterator{contract: _VaultContract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_VaultContract *VaultContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VaultContractDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VaultContract.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultContractDeposit)
				if err := _VaultContract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_VaultContract *VaultContractFilterer) ParseDeposit(log types.Log) (*VaultContractDeposit, error) {
	event := new(VaultContractDeposit)
	if err := _VaultContract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VaultContract contract.
type VaultContractTransferIterator struct {
	Event *VaultContractTransfer // Event containing the contract specifics and raw log

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
func (it *VaultContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultContractTransfer)
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
		it.Event = new(VaultContractTransfer)
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
func (it *VaultContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultContractTransfer represents a Transfer event raised by the VaultContract contract.
type VaultContractTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_VaultContract *VaultContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VaultContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VaultContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VaultContractTransferIterator{contract: _VaultContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_VaultContract *VaultContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VaultContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VaultContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultContractTransfer)
				if err := _VaultContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_VaultContract *VaultContractFilterer) ParseTransfer(log types.Log) (*VaultContractTransfer, error) {
	event := new(VaultContractTransfer)
	if err := _VaultContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the VaultContract contract.
type VaultContractWithdrawIterator struct {
	Event *VaultContractWithdraw // Event containing the contract specifics and raw log

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
func (it *VaultContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultContractWithdraw)
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
		it.Event = new(VaultContractWithdraw)
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
func (it *VaultContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultContractWithdraw represents a Withdraw event raised by the VaultContract contract.
type VaultContractWithdraw struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_VaultContract *VaultContractFilterer) FilterWithdraw(opts *bind.FilterOpts, account []common.Address) (*VaultContractWithdrawIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VaultContract.contract.FilterLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return &VaultContractWithdrawIterator{contract: _VaultContract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_VaultContract *VaultContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VaultContractWithdraw, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _VaultContract.contract.WatchLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultContractWithdraw)
				if err := _VaultContract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_VaultContract *VaultContractFilterer) ParseWithdraw(log types.Log) (*VaultContractWithdraw, error) {
	event := new(VaultContractWithdraw)
	if err := _VaultContract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
