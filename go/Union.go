// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Union

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

// UnionMetaData contains all meta data concerning the Union contract.
var UnionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"passportContract_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chat_id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"applier_id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"multy_wallet_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUnion.VotingType\",\"name\":\"vote_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voting_token_address\",\"type\":\"address\"}],\"name\":\"ApplicationForJoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"chat_id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"multy_wallet_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUnion.VotingType\",\"name\":\"vote_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voting_token_address\",\"type\":\"address\"}],\"name\":\"ApprovedJoin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"applyerTg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"daoTg\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"dao_\",\"type\":\"address\"},{\"internalType\":\"enumUnion.VotingType\",\"name\":\"votingType_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"votingTokenContract_\",\"type\":\"address\"}],\"name\":\"applyForUnion\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"daoAddress\",\"type\":\"address\"}],\"name\":\"approveJoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"daoAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"daos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"chatOwnerAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tgId\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"multisigAddress\",\"type\":\"address\"},{\"internalType\":\"enumUnion.VotingType\",\"name\":\"votingType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"votingToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tgpassport\",\"outputs\":[{\"internalType\":\"contractTGPassport\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UnionABI is the input ABI used to generate the binding from.
// Deprecated: Use UnionMetaData.ABI instead.
var UnionABI = UnionMetaData.ABI

// Union is an auto generated Go binding around an Ethereum contract.
type Union struct {
	UnionCaller     // Read-only binding to the contract
	UnionTransactor // Write-only binding to the contract
	UnionFilterer   // Log filterer for contract events
}

// UnionCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnionSession struct {
	Contract     *Union            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnionCallerSession struct {
	Contract *UnionCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UnionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnionTransactorSession struct {
	Contract     *UnionTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnionRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnionRaw struct {
	Contract *Union // Generic contract binding to access the raw methods on
}

// UnionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnionCallerRaw struct {
	Contract *UnionCaller // Generic read-only contract binding to access the raw methods on
}

// UnionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnionTransactorRaw struct {
	Contract *UnionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnion creates a new instance of Union, bound to a specific deployed contract.
func NewUnion(address common.Address, backend bind.ContractBackend) (*Union, error) {
	contract, err := bindUnion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Union{UnionCaller: UnionCaller{contract: contract}, UnionTransactor: UnionTransactor{contract: contract}, UnionFilterer: UnionFilterer{contract: contract}}, nil
}

// NewUnionCaller creates a new read-only instance of Union, bound to a specific deployed contract.
func NewUnionCaller(address common.Address, caller bind.ContractCaller) (*UnionCaller, error) {
	contract, err := bindUnion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnionCaller{contract: contract}, nil
}

// NewUnionTransactor creates a new write-only instance of Union, bound to a specific deployed contract.
func NewUnionTransactor(address common.Address, transactor bind.ContractTransactor) (*UnionTransactor, error) {
	contract, err := bindUnion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnionTransactor{contract: contract}, nil
}

// NewUnionFilterer creates a new log filterer instance of Union, bound to a specific deployed contract.
func NewUnionFilterer(address common.Address, filterer bind.ContractFilterer) (*UnionFilterer, error) {
	contract, err := bindUnion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnionFilterer{contract: contract}, nil
}

// bindUnion binds a generic wrapper to an already deployed contract.
func bindUnion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Union *UnionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Union.Contract.UnionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Union *UnionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Union.Contract.UnionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Union *UnionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Union.Contract.UnionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Union *UnionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Union.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Union *UnionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Union.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Union *UnionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Union.Contract.contract.Transact(opts, method, params...)
}

// DaoAddresses is a free data retrieval call binding the contract method 0x5c34780f.
//
// Solidity: function daoAddresses(string ) view returns(address)
func (_Union *UnionCaller) DaoAddresses(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Union.contract.Call(opts, &out, "daoAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoAddresses is a free data retrieval call binding the contract method 0x5c34780f.
//
// Solidity: function daoAddresses(string ) view returns(address)
func (_Union *UnionSession) DaoAddresses(arg0 string) (common.Address, error) {
	return _Union.Contract.DaoAddresses(&_Union.CallOpts, arg0)
}

// DaoAddresses is a free data retrieval call binding the contract method 0x5c34780f.
//
// Solidity: function daoAddresses(string ) view returns(address)
func (_Union *UnionCallerSession) DaoAddresses(arg0 string) (common.Address, error) {
	return _Union.Contract.DaoAddresses(&_Union.CallOpts, arg0)
}

// Daos is a free data retrieval call binding the contract method 0xc25f3cf6.
//
// Solidity: function daos(address ) view returns(address chatOwnerAddress, string tgId, bool valid, address multisigAddress, uint8 votingType, address votingToken)
func (_Union *UnionCaller) Daos(opts *bind.CallOpts, arg0 common.Address) (struct {
	ChatOwnerAddress common.Address
	TgId             string
	Valid            bool
	MultisigAddress  common.Address
	VotingType       uint8
	VotingToken      common.Address
}, error) {
	var out []interface{}
	err := _Union.contract.Call(opts, &out, "daos", arg0)

	outstruct := new(struct {
		ChatOwnerAddress common.Address
		TgId             string
		Valid            bool
		MultisigAddress  common.Address
		VotingType       uint8
		VotingToken      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChatOwnerAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TgId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Valid = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.MultisigAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.VotingType = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.VotingToken = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Daos is a free data retrieval call binding the contract method 0xc25f3cf6.
//
// Solidity: function daos(address ) view returns(address chatOwnerAddress, string tgId, bool valid, address multisigAddress, uint8 votingType, address votingToken)
func (_Union *UnionSession) Daos(arg0 common.Address) (struct {
	ChatOwnerAddress common.Address
	TgId             string
	Valid            bool
	MultisigAddress  common.Address
	VotingType       uint8
	VotingToken      common.Address
}, error) {
	return _Union.Contract.Daos(&_Union.CallOpts, arg0)
}

// Daos is a free data retrieval call binding the contract method 0xc25f3cf6.
//
// Solidity: function daos(address ) view returns(address chatOwnerAddress, string tgId, bool valid, address multisigAddress, uint8 votingType, address votingToken)
func (_Union *UnionCallerSession) Daos(arg0 common.Address) (struct {
	ChatOwnerAddress common.Address
	TgId             string
	Valid            bool
	MultisigAddress  common.Address
	VotingType       uint8
	VotingToken      common.Address
}, error) {
	return _Union.Contract.Daos(&_Union.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Union *UnionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Union.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Union *UnionSession) Owner() (common.Address, error) {
	return _Union.Contract.Owner(&_Union.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Union *UnionCallerSession) Owner() (common.Address, error) {
	return _Union.Contract.Owner(&_Union.CallOpts)
}

// Tgpassport is a free data retrieval call binding the contract method 0x7527d5a1.
//
// Solidity: function tgpassport() view returns(address)
func (_Union *UnionCaller) Tgpassport(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Union.contract.Call(opts, &out, "tgpassport")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tgpassport is a free data retrieval call binding the contract method 0x7527d5a1.
//
// Solidity: function tgpassport() view returns(address)
func (_Union *UnionSession) Tgpassport() (common.Address, error) {
	return _Union.Contract.Tgpassport(&_Union.CallOpts)
}

// Tgpassport is a free data retrieval call binding the contract method 0x7527d5a1.
//
// Solidity: function tgpassport() view returns(address)
func (_Union *UnionCallerSession) Tgpassport() (common.Address, error) {
	return _Union.Contract.Tgpassport(&_Union.CallOpts)
}

// ApplyForUnion is a paid mutator transaction binding the contract method 0xc9128823.
//
// Solidity: function applyForUnion(string applyerTg, string daoTg, address dao_, uint8 votingType_, address votingTokenContract_) payable returns()
func (_Union *UnionTransactor) ApplyForUnion(opts *bind.TransactOpts, applyerTg string, daoTg string, dao_ common.Address, votingType_ uint8, votingTokenContract_ common.Address) (*types.Transaction, error) {
	return _Union.contract.Transact(opts, "applyForUnion", applyerTg, daoTg, dao_, votingType_, votingTokenContract_)
}

// ApplyForUnion is a paid mutator transaction binding the contract method 0xc9128823.
//
// Solidity: function applyForUnion(string applyerTg, string daoTg, address dao_, uint8 votingType_, address votingTokenContract_) payable returns()
func (_Union *UnionSession) ApplyForUnion(applyerTg string, daoTg string, dao_ common.Address, votingType_ uint8, votingTokenContract_ common.Address) (*types.Transaction, error) {
	return _Union.Contract.ApplyForUnion(&_Union.TransactOpts, applyerTg, daoTg, dao_, votingType_, votingTokenContract_)
}

// ApplyForUnion is a paid mutator transaction binding the contract method 0xc9128823.
//
// Solidity: function applyForUnion(string applyerTg, string daoTg, address dao_, uint8 votingType_, address votingTokenContract_) payable returns()
func (_Union *UnionTransactorSession) ApplyForUnion(applyerTg string, daoTg string, dao_ common.Address, votingType_ uint8, votingTokenContract_ common.Address) (*types.Transaction, error) {
	return _Union.Contract.ApplyForUnion(&_Union.TransactOpts, applyerTg, daoTg, dao_, votingType_, votingTokenContract_)
}

// ApproveJoin is a paid mutator transaction binding the contract method 0x564f1139.
//
// Solidity: function approveJoin(address daoAddress) returns()
func (_Union *UnionTransactor) ApproveJoin(opts *bind.TransactOpts, daoAddress common.Address) (*types.Transaction, error) {
	return _Union.contract.Transact(opts, "approveJoin", daoAddress)
}

// ApproveJoin is a paid mutator transaction binding the contract method 0x564f1139.
//
// Solidity: function approveJoin(address daoAddress) returns()
func (_Union *UnionSession) ApproveJoin(daoAddress common.Address) (*types.Transaction, error) {
	return _Union.Contract.ApproveJoin(&_Union.TransactOpts, daoAddress)
}

// ApproveJoin is a paid mutator transaction binding the contract method 0x564f1139.
//
// Solidity: function approveJoin(address daoAddress) returns()
func (_Union *UnionTransactorSession) ApproveJoin(daoAddress common.Address) (*types.Transaction, error) {
	return _Union.Contract.ApproveJoin(&_Union.TransactOpts, daoAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Union *UnionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Union.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Union *UnionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Union.Contract.RenounceOwnership(&_Union.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Union *UnionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Union.Contract.RenounceOwnership(&_Union.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Union *UnionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Union.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Union *UnionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Union.Contract.TransferOwnership(&_Union.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Union *UnionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Union.Contract.TransferOwnership(&_Union.TransactOpts, newOwner)
}

// UnionApplicationForJoinIterator is returned from FilterApplicationForJoin and is used to iterate over the raw logs and unpacked data for ApplicationForJoin events raised by the Union contract.
type UnionApplicationForJoinIterator struct {
	Event *UnionApplicationForJoin // Event containing the contract specifics and raw log

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
func (it *UnionApplicationForJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnionApplicationForJoin)
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
		it.Event = new(UnionApplicationForJoin)
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
func (it *UnionApplicationForJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnionApplicationForJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnionApplicationForJoin represents a ApplicationForJoin event raised by the Union contract.
type UnionApplicationForJoin struct {
	ChatId             string
	ApplierId          string
	MultyWalletAddress common.Address
	VoteType           uint8
	VotingTokenAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterApplicationForJoin is a free log retrieval operation binding the contract event 0x829e7a1d0527546d71c2b19a98fc782e3f2695ac9f8662ed2a57353a3820332f.
//
// Solidity: event ApplicationForJoin(string chat_id, string applier_id, address multy_wallet_address, uint8 vote_type, address voting_token_address)
func (_Union *UnionFilterer) FilterApplicationForJoin(opts *bind.FilterOpts) (*UnionApplicationForJoinIterator, error) {

	logs, sub, err := _Union.contract.FilterLogs(opts, "ApplicationForJoin")
	if err != nil {
		return nil, err
	}
	return &UnionApplicationForJoinIterator{contract: _Union.contract, event: "ApplicationForJoin", logs: logs, sub: sub}, nil
}

// WatchApplicationForJoin is a free log subscription operation binding the contract event 0x829e7a1d0527546d71c2b19a98fc782e3f2695ac9f8662ed2a57353a3820332f.
//
// Solidity: event ApplicationForJoin(string chat_id, string applier_id, address multy_wallet_address, uint8 vote_type, address voting_token_address)
func (_Union *UnionFilterer) WatchApplicationForJoin(opts *bind.WatchOpts, sink chan<- *UnionApplicationForJoin) (event.Subscription, error) {

	logs, sub, err := _Union.contract.WatchLogs(opts, "ApplicationForJoin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnionApplicationForJoin)
				if err := _Union.contract.UnpackLog(event, "ApplicationForJoin", log); err != nil {
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

// ParseApplicationForJoin is a log parse operation binding the contract event 0x829e7a1d0527546d71c2b19a98fc782e3f2695ac9f8662ed2a57353a3820332f.
//
// Solidity: event ApplicationForJoin(string chat_id, string applier_id, address multy_wallet_address, uint8 vote_type, address voting_token_address)
func (_Union *UnionFilterer) ParseApplicationForJoin(log types.Log) (*UnionApplicationForJoin, error) {
	event := new(UnionApplicationForJoin)
	if err := _Union.contract.UnpackLog(event, "ApplicationForJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnionApprovedJoinIterator is returned from FilterApprovedJoin and is used to iterate over the raw logs and unpacked data for ApprovedJoin events raised by the Union contract.
type UnionApprovedJoinIterator struct {
	Event *UnionApprovedJoin // Event containing the contract specifics and raw log

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
func (it *UnionApprovedJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnionApprovedJoin)
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
		it.Event = new(UnionApprovedJoin)
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
func (it *UnionApprovedJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnionApprovedJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnionApprovedJoin represents a ApprovedJoin event raised by the Union contract.
type UnionApprovedJoin struct {
	ChatId             string
	MultyWalletAddress common.Address
	VoteType           uint8
	VotingTokenAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterApprovedJoin is a free log retrieval operation binding the contract event 0xfebb10aa28a3f60a7bb7ac109fc821ae99fce5a09f60f103e2985fb0438c484b.
//
// Solidity: event ApprovedJoin(string chat_id, address multy_wallet_address, uint8 vote_type, address voting_token_address)
func (_Union *UnionFilterer) FilterApprovedJoin(opts *bind.FilterOpts) (*UnionApprovedJoinIterator, error) {

	logs, sub, err := _Union.contract.FilterLogs(opts, "ApprovedJoin")
	if err != nil {
		return nil, err
	}
	return &UnionApprovedJoinIterator{contract: _Union.contract, event: "ApprovedJoin", logs: logs, sub: sub}, nil
}

// WatchApprovedJoin is a free log subscription operation binding the contract event 0xfebb10aa28a3f60a7bb7ac109fc821ae99fce5a09f60f103e2985fb0438c484b.
//
// Solidity: event ApprovedJoin(string chat_id, address multy_wallet_address, uint8 vote_type, address voting_token_address)
func (_Union *UnionFilterer) WatchApprovedJoin(opts *bind.WatchOpts, sink chan<- *UnionApprovedJoin) (event.Subscription, error) {

	logs, sub, err := _Union.contract.WatchLogs(opts, "ApprovedJoin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnionApprovedJoin)
				if err := _Union.contract.UnpackLog(event, "ApprovedJoin", log); err != nil {
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

// ParseApprovedJoin is a log parse operation binding the contract event 0xfebb10aa28a3f60a7bb7ac109fc821ae99fce5a09f60f103e2985fb0438c484b.
//
// Solidity: event ApprovedJoin(string chat_id, address multy_wallet_address, uint8 vote_type, address voting_token_address)
func (_Union *UnionFilterer) ParseApprovedJoin(log types.Log) (*UnionApprovedJoin, error) {
	event := new(UnionApprovedJoin)
	if err := _Union.contract.UnpackLog(event, "ApprovedJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Union contract.
type UnionOwnershipTransferredIterator struct {
	Event *UnionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UnionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnionOwnershipTransferred)
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
		it.Event = new(UnionOwnershipTransferred)
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
func (it *UnionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnionOwnershipTransferred represents a OwnershipTransferred event raised by the Union contract.
type UnionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Union *UnionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UnionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Union.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UnionOwnershipTransferredIterator{contract: _Union.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Union *UnionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UnionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Union.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnionOwnershipTransferred)
				if err := _Union.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Union *UnionFilterer) ParseOwnershipTransferred(log types.Log) (*UnionOwnershipTransferred, error) {
	event := new(UnionOwnershipTransferred)
	if err := _Union.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
