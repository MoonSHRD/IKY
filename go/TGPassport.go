// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TGPassport

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

// TGPassportMetaData contains all meta data concerning the TGPassport contract.
var TGPassportMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"applyerTg\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet_address\",\"type\":\"address\"}],\"name\":\"passportApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"applyerTg\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"passportApproved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"applyerTg\",\"type\":\"string\"}],\"name\":\"applyForPassport\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"passportToApprove\",\"type\":\"address\"}],\"name\":\"approvePassport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPassportFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tgId_\",\"type\":\"string\"}],\"name\":\"getPassportWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"passports\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tgId\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"passportFee_\",\"type\":\"uint256\"}],\"name\":\"setPassportFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tgIdToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TGPassportABI is the input ABI used to generate the binding from.
// Deprecated: Use TGPassportMetaData.ABI instead.
var TGPassportABI = TGPassportMetaData.ABI

// TGPassport is an auto generated Go binding around an Ethereum contract.
type TGPassport struct {
	TGPassportCaller     // Read-only binding to the contract
	TGPassportTransactor // Write-only binding to the contract
	TGPassportFilterer   // Log filterer for contract events
}

// TGPassportCaller is an auto generated read-only Go binding around an Ethereum contract.
type TGPassportCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TGPassportTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TGPassportTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TGPassportFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TGPassportFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TGPassportSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TGPassportSession struct {
	Contract     *TGPassport       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TGPassportCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TGPassportCallerSession struct {
	Contract *TGPassportCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TGPassportTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TGPassportTransactorSession struct {
	Contract     *TGPassportTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TGPassportRaw is an auto generated low-level Go binding around an Ethereum contract.
type TGPassportRaw struct {
	Contract *TGPassport // Generic contract binding to access the raw methods on
}

// TGPassportCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TGPassportCallerRaw struct {
	Contract *TGPassportCaller // Generic read-only contract binding to access the raw methods on
}

// TGPassportTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TGPassportTransactorRaw struct {
	Contract *TGPassportTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTGPassport creates a new instance of TGPassport, bound to a specific deployed contract.
func NewTGPassport(address common.Address, backend bind.ContractBackend) (*TGPassport, error) {
	contract, err := bindTGPassport(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TGPassport{TGPassportCaller: TGPassportCaller{contract: contract}, TGPassportTransactor: TGPassportTransactor{contract: contract}, TGPassportFilterer: TGPassportFilterer{contract: contract}}, nil
}

// NewTGPassportCaller creates a new read-only instance of TGPassport, bound to a specific deployed contract.
func NewTGPassportCaller(address common.Address, caller bind.ContractCaller) (*TGPassportCaller, error) {
	contract, err := bindTGPassport(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TGPassportCaller{contract: contract}, nil
}

// NewTGPassportTransactor creates a new write-only instance of TGPassport, bound to a specific deployed contract.
func NewTGPassportTransactor(address common.Address, transactor bind.ContractTransactor) (*TGPassportTransactor, error) {
	contract, err := bindTGPassport(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TGPassportTransactor{contract: contract}, nil
}

// NewTGPassportFilterer creates a new log filterer instance of TGPassport, bound to a specific deployed contract.
func NewTGPassportFilterer(address common.Address, filterer bind.ContractFilterer) (*TGPassportFilterer, error) {
	contract, err := bindTGPassport(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TGPassportFilterer{contract: contract}, nil
}

// bindTGPassport binds a generic wrapper to an already deployed contract.
func bindTGPassport(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TGPassportABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TGPassport *TGPassportRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TGPassport.Contract.TGPassportCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TGPassport *TGPassportRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TGPassport.Contract.TGPassportTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TGPassport *TGPassportRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TGPassport.Contract.TGPassportTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TGPassport *TGPassportCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TGPassport.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TGPassport *TGPassportTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TGPassport.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TGPassport *TGPassportTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TGPassport.Contract.contract.Transact(opts, method, params...)
}

// GetPassportFee is a free data retrieval call binding the contract method 0xc627c170.
//
// Solidity: function getPassportFee() view returns(uint256)
func (_TGPassport *TGPassportCaller) GetPassportFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TGPassport.contract.Call(opts, &out, "getPassportFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPassportFee is a free data retrieval call binding the contract method 0xc627c170.
//
// Solidity: function getPassportFee() view returns(uint256)
func (_TGPassport *TGPassportSession) GetPassportFee() (*big.Int, error) {
	return _TGPassport.Contract.GetPassportFee(&_TGPassport.CallOpts)
}

// GetPassportFee is a free data retrieval call binding the contract method 0xc627c170.
//
// Solidity: function getPassportFee() view returns(uint256)
func (_TGPassport *TGPassportCallerSession) GetPassportFee() (*big.Int, error) {
	return _TGPassport.Contract.GetPassportFee(&_TGPassport.CallOpts)
}

// GetPassportWallet is a free data retrieval call binding the contract method 0x95939e71.
//
// Solidity: function getPassportWallet(string tgId_) view returns(address)
func (_TGPassport *TGPassportCaller) GetPassportWallet(opts *bind.CallOpts, tgId_ string) (common.Address, error) {
	var out []interface{}
	err := _TGPassport.contract.Call(opts, &out, "getPassportWallet", tgId_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPassportWallet is a free data retrieval call binding the contract method 0x95939e71.
//
// Solidity: function getPassportWallet(string tgId_) view returns(address)
func (_TGPassport *TGPassportSession) GetPassportWallet(tgId_ string) (common.Address, error) {
	return _TGPassport.Contract.GetPassportWallet(&_TGPassport.CallOpts, tgId_)
}

// GetPassportWallet is a free data retrieval call binding the contract method 0x95939e71.
//
// Solidity: function getPassportWallet(string tgId_) view returns(address)
func (_TGPassport *TGPassportCallerSession) GetPassportWallet(tgId_ string) (common.Address, error) {
	return _TGPassport.Contract.GetPassportWallet(&_TGPassport.CallOpts, tgId_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TGPassport *TGPassportCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TGPassport.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TGPassport *TGPassportSession) Owner() (common.Address, error) {
	return _TGPassport.Contract.Owner(&_TGPassport.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TGPassport *TGPassportCallerSession) Owner() (common.Address, error) {
	return _TGPassport.Contract.Owner(&_TGPassport.CallOpts)
}

// Passports is a free data retrieval call binding the contract method 0xe37c132b.
//
// Solidity: function passports(address ) view returns(address userAddress, string tgId, bool valid, address validatorAddress)
func (_TGPassport *TGPassportCaller) Passports(opts *bind.CallOpts, arg0 common.Address) (struct {
	UserAddress      common.Address
	TgId             string
	Valid            bool
	ValidatorAddress common.Address
}, error) {
	var out []interface{}
	err := _TGPassport.contract.Call(opts, &out, "passports", arg0)

	outstruct := new(struct {
		UserAddress      common.Address
		TgId             string
		Valid            bool
		ValidatorAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TgId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Valid = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Passports is a free data retrieval call binding the contract method 0xe37c132b.
//
// Solidity: function passports(address ) view returns(address userAddress, string tgId, bool valid, address validatorAddress)
func (_TGPassport *TGPassportSession) Passports(arg0 common.Address) (struct {
	UserAddress      common.Address
	TgId             string
	Valid            bool
	ValidatorAddress common.Address
}, error) {
	return _TGPassport.Contract.Passports(&_TGPassport.CallOpts, arg0)
}

// Passports is a free data retrieval call binding the contract method 0xe37c132b.
//
// Solidity: function passports(address ) view returns(address userAddress, string tgId, bool valid, address validatorAddress)
func (_TGPassport *TGPassportCallerSession) Passports(arg0 common.Address) (struct {
	UserAddress      common.Address
	TgId             string
	Valid            bool
	ValidatorAddress common.Address
}, error) {
	return _TGPassport.Contract.Passports(&_TGPassport.CallOpts, arg0)
}

// TgIdToAddress is a free data retrieval call binding the contract method 0x86e81ba5.
//
// Solidity: function tgIdToAddress(string ) view returns(address)
func (_TGPassport *TGPassportCaller) TgIdToAddress(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _TGPassport.contract.Call(opts, &out, "tgIdToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TgIdToAddress is a free data retrieval call binding the contract method 0x86e81ba5.
//
// Solidity: function tgIdToAddress(string ) view returns(address)
func (_TGPassport *TGPassportSession) TgIdToAddress(arg0 string) (common.Address, error) {
	return _TGPassport.Contract.TgIdToAddress(&_TGPassport.CallOpts, arg0)
}

// TgIdToAddress is a free data retrieval call binding the contract method 0x86e81ba5.
//
// Solidity: function tgIdToAddress(string ) view returns(address)
func (_TGPassport *TGPassportCallerSession) TgIdToAddress(arg0 string) (common.Address, error) {
	return _TGPassport.Contract.TgIdToAddress(&_TGPassport.CallOpts, arg0)
}

// ApplyForPassport is a paid mutator transaction binding the contract method 0x1b1a5d64.
//
// Solidity: function applyForPassport(string applyerTg) payable returns()
func (_TGPassport *TGPassportTransactor) ApplyForPassport(opts *bind.TransactOpts, applyerTg string) (*types.Transaction, error) {
	return _TGPassport.contract.Transact(opts, "applyForPassport", applyerTg)
}

// ApplyForPassport is a paid mutator transaction binding the contract method 0x1b1a5d64.
//
// Solidity: function applyForPassport(string applyerTg) payable returns()
func (_TGPassport *TGPassportSession) ApplyForPassport(applyerTg string) (*types.Transaction, error) {
	return _TGPassport.Contract.ApplyForPassport(&_TGPassport.TransactOpts, applyerTg)
}

// ApplyForPassport is a paid mutator transaction binding the contract method 0x1b1a5d64.
//
// Solidity: function applyForPassport(string applyerTg) payable returns()
func (_TGPassport *TGPassportTransactorSession) ApplyForPassport(applyerTg string) (*types.Transaction, error) {
	return _TGPassport.Contract.ApplyForPassport(&_TGPassport.TransactOpts, applyerTg)
}

// ApprovePassport is a paid mutator transaction binding the contract method 0x576cff98.
//
// Solidity: function approvePassport(address passportToApprove) returns()
func (_TGPassport *TGPassportTransactor) ApprovePassport(opts *bind.TransactOpts, passportToApprove common.Address) (*types.Transaction, error) {
	return _TGPassport.contract.Transact(opts, "approvePassport", passportToApprove)
}

// ApprovePassport is a paid mutator transaction binding the contract method 0x576cff98.
//
// Solidity: function approvePassport(address passportToApprove) returns()
func (_TGPassport *TGPassportSession) ApprovePassport(passportToApprove common.Address) (*types.Transaction, error) {
	return _TGPassport.Contract.ApprovePassport(&_TGPassport.TransactOpts, passportToApprove)
}

// ApprovePassport is a paid mutator transaction binding the contract method 0x576cff98.
//
// Solidity: function approvePassport(address passportToApprove) returns()
func (_TGPassport *TGPassportTransactorSession) ApprovePassport(passportToApprove common.Address) (*types.Transaction, error) {
	return _TGPassport.Contract.ApprovePassport(&_TGPassport.TransactOpts, passportToApprove)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TGPassport *TGPassportTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TGPassport.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TGPassport *TGPassportSession) RenounceOwnership() (*types.Transaction, error) {
	return _TGPassport.Contract.RenounceOwnership(&_TGPassport.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TGPassport *TGPassportTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TGPassport.Contract.RenounceOwnership(&_TGPassport.TransactOpts)
}

// SetPassportFee is a paid mutator transaction binding the contract method 0x875823c0.
//
// Solidity: function setPassportFee(uint256 passportFee_) returns()
func (_TGPassport *TGPassportTransactor) SetPassportFee(opts *bind.TransactOpts, passportFee_ *big.Int) (*types.Transaction, error) {
	return _TGPassport.contract.Transact(opts, "setPassportFee", passportFee_)
}

// SetPassportFee is a paid mutator transaction binding the contract method 0x875823c0.
//
// Solidity: function setPassportFee(uint256 passportFee_) returns()
func (_TGPassport *TGPassportSession) SetPassportFee(passportFee_ *big.Int) (*types.Transaction, error) {
	return _TGPassport.Contract.SetPassportFee(&_TGPassport.TransactOpts, passportFee_)
}

// SetPassportFee is a paid mutator transaction binding the contract method 0x875823c0.
//
// Solidity: function setPassportFee(uint256 passportFee_) returns()
func (_TGPassport *TGPassportTransactorSession) SetPassportFee(passportFee_ *big.Int) (*types.Transaction, error) {
	return _TGPassport.Contract.SetPassportFee(&_TGPassport.TransactOpts, passportFee_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TGPassport *TGPassportTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TGPassport.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TGPassport *TGPassportSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TGPassport.Contract.TransferOwnership(&_TGPassport.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TGPassport *TGPassportTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TGPassport.Contract.TransferOwnership(&_TGPassport.TransactOpts, newOwner)
}

// TGPassportOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TGPassport contract.
type TGPassportOwnershipTransferredIterator struct {
	Event *TGPassportOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TGPassportOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TGPassportOwnershipTransferred)
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
		it.Event = new(TGPassportOwnershipTransferred)
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
func (it *TGPassportOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TGPassportOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TGPassportOwnershipTransferred represents a OwnershipTransferred event raised by the TGPassport contract.
type TGPassportOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TGPassport *TGPassportFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TGPassportOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TGPassport.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TGPassportOwnershipTransferredIterator{contract: _TGPassport.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TGPassport *TGPassportFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TGPassportOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TGPassport.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TGPassportOwnershipTransferred)
				if err := _TGPassport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TGPassport *TGPassportFilterer) ParseOwnershipTransferred(log types.Log) (*TGPassportOwnershipTransferred, error) {
	event := new(TGPassportOwnershipTransferred)
	if err := _TGPassport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TGPassportPassportAppliedIterator is returned from FilterPassportApplied and is used to iterate over the raw logs and unpacked data for PassportApplied events raised by the TGPassport contract.
type TGPassportPassportAppliedIterator struct {
	Event *TGPassportPassportApplied // Event containing the contract specifics and raw log

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
func (it *TGPassportPassportAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TGPassportPassportApplied)
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
		it.Event = new(TGPassportPassportApplied)
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
func (it *TGPassportPassportAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TGPassportPassportAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TGPassportPassportApplied represents a PassportApplied event raised by the TGPassport contract.
type TGPassportPassportApplied struct {
	ApplyerTg     string
	WalletAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPassportApplied is a free log retrieval operation binding the contract event 0x2578b8463fa9088621234109e3537776ec83fa33de185759f3a0616bfefe8e30.
//
// Solidity: event passportApplied(string applyerTg, address wallet_address)
func (_TGPassport *TGPassportFilterer) FilterPassportApplied(opts *bind.FilterOpts) (*TGPassportPassportAppliedIterator, error) {

	logs, sub, err := _TGPassport.contract.FilterLogs(opts, "passportApplied")
	if err != nil {
		return nil, err
	}
	return &TGPassportPassportAppliedIterator{contract: _TGPassport.contract, event: "passportApplied", logs: logs, sub: sub}, nil
}

// WatchPassportApplied is a free log subscription operation binding the contract event 0x2578b8463fa9088621234109e3537776ec83fa33de185759f3a0616bfefe8e30.
//
// Solidity: event passportApplied(string applyerTg, address wallet_address)
func (_TGPassport *TGPassportFilterer) WatchPassportApplied(opts *bind.WatchOpts, sink chan<- *TGPassportPassportApplied) (event.Subscription, error) {

	logs, sub, err := _TGPassport.contract.WatchLogs(opts, "passportApplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TGPassportPassportApplied)
				if err := _TGPassport.contract.UnpackLog(event, "passportApplied", log); err != nil {
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

// ParsePassportApplied is a log parse operation binding the contract event 0x2578b8463fa9088621234109e3537776ec83fa33de185759f3a0616bfefe8e30.
//
// Solidity: event passportApplied(string applyerTg, address wallet_address)
func (_TGPassport *TGPassportFilterer) ParsePassportApplied(log types.Log) (*TGPassportPassportApplied, error) {
	event := new(TGPassportPassportApplied)
	if err := _TGPassport.contract.UnpackLog(event, "passportApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TGPassportPassportApprovedIterator is returned from FilterPassportApproved and is used to iterate over the raw logs and unpacked data for PassportApproved events raised by the TGPassport contract.
type TGPassportPassportApprovedIterator struct {
	Event *TGPassportPassportApproved // Event containing the contract specifics and raw log

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
func (it *TGPassportPassportApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TGPassportPassportApproved)
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
		it.Event = new(TGPassportPassportApproved)
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
func (it *TGPassportPassportApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TGPassportPassportApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TGPassportPassportApproved represents a PassportApproved event raised by the TGPassport contract.
type TGPassportPassportApproved struct {
	ApplyerTg     string
	WalletAddress common.Address
	Issuer        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPassportApproved is a free log retrieval operation binding the contract event 0x425c0ff4dfba8976f680edd0bb8cd915c8ca31839207163b7bbcc3b72f4dd42f.
//
// Solidity: event passportApproved(string applyerTg, address wallet_address, address issuer)
func (_TGPassport *TGPassportFilterer) FilterPassportApproved(opts *bind.FilterOpts) (*TGPassportPassportApprovedIterator, error) {

	logs, sub, err := _TGPassport.contract.FilterLogs(opts, "passportApproved")
	if err != nil {
		return nil, err
	}
	return &TGPassportPassportApprovedIterator{contract: _TGPassport.contract, event: "passportApproved", logs: logs, sub: sub}, nil
}

// WatchPassportApproved is a free log subscription operation binding the contract event 0x425c0ff4dfba8976f680edd0bb8cd915c8ca31839207163b7bbcc3b72f4dd42f.
//
// Solidity: event passportApproved(string applyerTg, address wallet_address, address issuer)
func (_TGPassport *TGPassportFilterer) WatchPassportApproved(opts *bind.WatchOpts, sink chan<- *TGPassportPassportApproved) (event.Subscription, error) {

	logs, sub, err := _TGPassport.contract.WatchLogs(opts, "passportApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TGPassportPassportApproved)
				if err := _TGPassport.contract.UnpackLog(event, "passportApproved", log); err != nil {
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

// ParsePassportApproved is a log parse operation binding the contract event 0x425c0ff4dfba8976f680edd0bb8cd915c8ca31839207163b7bbcc3b72f4dd42f.
//
// Solidity: event passportApproved(string applyerTg, address wallet_address, address issuer)
func (_TGPassport *TGPassportFilterer) ParsePassportApproved(log types.Log) (*TGPassportPassportApproved, error) {
	event := new(TGPassportPassportApproved)
	if err := _TGPassport.contract.UnpackLog(event, "passportApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
