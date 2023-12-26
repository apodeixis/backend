// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package posts

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

// PostsMetaData contains all meta data concerning the Posts contract.
var PostsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createPost\",\"inputs\":[{\"name\":\"_userID\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"_postHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPosts\",\"inputs\":[{\"name\":\"_userID\",\"type\":\"int64\",\"internalType\":\"int64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"posts\",\"inputs\":[{\"name\":\"\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// PostsABI is the input ABI used to generate the binding from.
// Deprecated: Use PostsMetaData.ABI instead.
var PostsABI = PostsMetaData.ABI

// Posts is an auto generated Go binding around an Ethereum contract.
type Posts struct {
	PostsCaller     // Read-only binding to the contract
	PostsTransactor // Write-only binding to the contract
	PostsFilterer   // Log filterer for contract events
}

// PostsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PostsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PostsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PostsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PostsSession struct {
	Contract     *Posts            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PostsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PostsCallerSession struct {
	Contract *PostsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PostsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PostsTransactorSession struct {
	Contract     *PostsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PostsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PostsRaw struct {
	Contract *Posts // Generic contract binding to access the raw methods on
}

// PostsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PostsCallerRaw struct {
	Contract *PostsCaller // Generic read-only contract binding to access the raw methods on
}

// PostsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PostsTransactorRaw struct {
	Contract *PostsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPosts creates a new instance of Posts, bound to a specific deployed contract.
func NewPosts(address common.Address, backend bind.ContractBackend) (*Posts, error) {
	contract, err := bindPosts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Posts{PostsCaller: PostsCaller{contract: contract}, PostsTransactor: PostsTransactor{contract: contract}, PostsFilterer: PostsFilterer{contract: contract}}, nil
}

// NewPostsCaller creates a new read-only instance of Posts, bound to a specific deployed contract.
func NewPostsCaller(address common.Address, caller bind.ContractCaller) (*PostsCaller, error) {
	contract, err := bindPosts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PostsCaller{contract: contract}, nil
}

// NewPostsTransactor creates a new write-only instance of Posts, bound to a specific deployed contract.
func NewPostsTransactor(address common.Address, transactor bind.ContractTransactor) (*PostsTransactor, error) {
	contract, err := bindPosts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PostsTransactor{contract: contract}, nil
}

// NewPostsFilterer creates a new log filterer instance of Posts, bound to a specific deployed contract.
func NewPostsFilterer(address common.Address, filterer bind.ContractFilterer) (*PostsFilterer, error) {
	contract, err := bindPosts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PostsFilterer{contract: contract}, nil
}

// bindPosts binds a generic wrapper to an already deployed contract.
func bindPosts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PostsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Posts *PostsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Posts.Contract.PostsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Posts *PostsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Posts.Contract.PostsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Posts *PostsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Posts.Contract.PostsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Posts *PostsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Posts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Posts *PostsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Posts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Posts *PostsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Posts.Contract.contract.Transact(opts, method, params...)
}

// GetPosts is a free data retrieval call binding the contract method 0x9150bb55.
//
// Solidity: function getPosts(int64 _userID) view returns(bytes32[])
func (_Posts *PostsCaller) GetPosts(opts *bind.CallOpts, _userID int64) ([][32]byte, error) {
	var out []interface{}
	err := _Posts.contract.Call(opts, &out, "getPosts", _userID)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPosts is a free data retrieval call binding the contract method 0x9150bb55.
//
// Solidity: function getPosts(int64 _userID) view returns(bytes32[])
func (_Posts *PostsSession) GetPosts(_userID int64) ([][32]byte, error) {
	return _Posts.Contract.GetPosts(&_Posts.CallOpts, _userID)
}

// GetPosts is a free data retrieval call binding the contract method 0x9150bb55.
//
// Solidity: function getPosts(int64 _userID) view returns(bytes32[])
func (_Posts *PostsCallerSession) GetPosts(_userID int64) ([][32]byte, error) {
	return _Posts.Contract.GetPosts(&_Posts.CallOpts, _userID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Posts *PostsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Posts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Posts *PostsSession) Owner() (common.Address, error) {
	return _Posts.Contract.Owner(&_Posts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Posts *PostsCallerSession) Owner() (common.Address, error) {
	return _Posts.Contract.Owner(&_Posts.CallOpts)
}

// Posts is a free data retrieval call binding the contract method 0x5b4421da.
//
// Solidity: function posts(int64 , uint256 ) view returns(bytes32)
func (_Posts *PostsCaller) Posts(opts *bind.CallOpts, arg0 int64, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Posts.contract.Call(opts, &out, "posts", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Posts is a free data retrieval call binding the contract method 0x5b4421da.
//
// Solidity: function posts(int64 , uint256 ) view returns(bytes32)
func (_Posts *PostsSession) Posts(arg0 int64, arg1 *big.Int) ([32]byte, error) {
	return _Posts.Contract.Posts(&_Posts.CallOpts, arg0, arg1)
}

// Posts is a free data retrieval call binding the contract method 0x5b4421da.
//
// Solidity: function posts(int64 , uint256 ) view returns(bytes32)
func (_Posts *PostsCallerSession) Posts(arg0 int64, arg1 *big.Int) ([32]byte, error) {
	return _Posts.Contract.Posts(&_Posts.CallOpts, arg0, arg1)
}

// CreatePost is a paid mutator transaction binding the contract method 0x8f4fea57.
//
// Solidity: function createPost(int64 _userID, bytes32 _postHash) returns()
func (_Posts *PostsTransactor) CreatePost(opts *bind.TransactOpts, _userID int64, _postHash [32]byte) (*types.Transaction, error) {
	return _Posts.contract.Transact(opts, "createPost", _userID, _postHash)
}

// CreatePost is a paid mutator transaction binding the contract method 0x8f4fea57.
//
// Solidity: function createPost(int64 _userID, bytes32 _postHash) returns()
func (_Posts *PostsSession) CreatePost(_userID int64, _postHash [32]byte) (*types.Transaction, error) {
	return _Posts.Contract.CreatePost(&_Posts.TransactOpts, _userID, _postHash)
}

// CreatePost is a paid mutator transaction binding the contract method 0x8f4fea57.
//
// Solidity: function createPost(int64 _userID, bytes32 _postHash) returns()
func (_Posts *PostsTransactorSession) CreatePost(_userID int64, _postHash [32]byte) (*types.Transaction, error) {
	return _Posts.Contract.CreatePost(&_Posts.TransactOpts, _userID, _postHash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Posts *PostsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Posts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Posts *PostsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Posts.Contract.RenounceOwnership(&_Posts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Posts *PostsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Posts.Contract.RenounceOwnership(&_Posts.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Posts *PostsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Posts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Posts *PostsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Posts.Contract.TransferOwnership(&_Posts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Posts *PostsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Posts.Contract.TransferOwnership(&_Posts.TransactOpts, newOwner)
}

// PostsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Posts contract.
type PostsOwnershipTransferredIterator struct {
	Event *PostsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PostsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostsOwnershipTransferred)
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
		it.Event = new(PostsOwnershipTransferred)
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
func (it *PostsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostsOwnershipTransferred represents a OwnershipTransferred event raised by the Posts contract.
type PostsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Posts *PostsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PostsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Posts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PostsOwnershipTransferredIterator{contract: _Posts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Posts *PostsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PostsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Posts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostsOwnershipTransferred)
				if err := _Posts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Posts *PostsFilterer) ParseOwnershipTransferred(log types.Log) (*PostsOwnershipTransferred, error) {
	event := new(PostsOwnershipTransferred)
	if err := _Posts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
