// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Confirm

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

// ConfirmMetaData contains all meta data concerning the Confirm contract.
var ConfirmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddr_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumStates\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"Confirmed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmProfits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"profitInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"profits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"trxHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumStates\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"setClaimed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"trxHash\",\"type\":\"bytes32\"}],\"name\":\"setTrx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ConfirmABI is the input ABI used to generate the binding from.
// Deprecated: Use ConfirmMetaData.ABI instead.
var ConfirmABI = ConfirmMetaData.ABI

// Confirm is an auto generated Go binding around an Ethereum contract.
type Confirm struct {
	ConfirmCaller     // Read-only binding to the contract
	ConfirmTransactor // Write-only binding to the contract
	ConfirmFilterer   // Log filterer for contract events
}

// ConfirmCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConfirmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConfirmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConfirmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfirmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConfirmSession struct {
	Contract     *Confirm          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfirmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConfirmCallerSession struct {
	Contract *ConfirmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ConfirmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConfirmTransactorSession struct {
	Contract     *ConfirmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConfirmRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConfirmRaw struct {
	Contract *Confirm // Generic contract binding to access the raw methods on
}

// ConfirmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConfirmCallerRaw struct {
	Contract *ConfirmCaller // Generic read-only contract binding to access the raw methods on
}

// ConfirmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConfirmTransactorRaw struct {
	Contract *ConfirmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfirm creates a new instance of Confirm, bound to a specific deployed contract.
func NewConfirm(address common.Address, backend bind.ContractBackend) (*Confirm, error) {
	contract, err := bindConfirm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Confirm{ConfirmCaller: ConfirmCaller{contract: contract}, ConfirmTransactor: ConfirmTransactor{contract: contract}, ConfirmFilterer: ConfirmFilterer{contract: contract}}, nil
}

// NewConfirmCaller creates a new read-only instance of Confirm, bound to a specific deployed contract.
func NewConfirmCaller(address common.Address, caller bind.ContractCaller) (*ConfirmCaller, error) {
	contract, err := bindConfirm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmCaller{contract: contract}, nil
}

// NewConfirmTransactor creates a new write-only instance of Confirm, bound to a specific deployed contract.
func NewConfirmTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfirmTransactor, error) {
	contract, err := bindConfirm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfirmTransactor{contract: contract}, nil
}

// NewConfirmFilterer creates a new log filterer instance of Confirm, bound to a specific deployed contract.
func NewConfirmFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfirmFilterer, error) {
	contract, err := bindConfirm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfirmFilterer{contract: contract}, nil
}

// bindConfirm binds a generic wrapper to an already deployed contract.
func bindConfirm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfirmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Confirm *ConfirmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Confirm.Contract.ConfirmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Confirm *ConfirmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Confirm.Contract.ConfirmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Confirm *ConfirmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Confirm.Contract.ConfirmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Confirm *ConfirmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Confirm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Confirm *ConfirmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Confirm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Confirm *ConfirmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Confirm.Contract.contract.Transact(opts, method, params...)
}

// ContractAddr is a free data retrieval call binding the contract method 0x15a5d9d8.
//
// Solidity: function contractAddr() view returns(address)
func (_Confirm *ConfirmCaller) ContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Confirm.contract.Call(opts, &out, "contractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractAddr is a free data retrieval call binding the contract method 0x15a5d9d8.
//
// Solidity: function contractAddr() view returns(address)
func (_Confirm *ConfirmSession) ContractAddr() (common.Address, error) {
	return _Confirm.Contract.ContractAddr(&_Confirm.CallOpts)
}

// ContractAddr is a free data retrieval call binding the contract method 0x15a5d9d8.
//
// Solidity: function contractAddr() view returns(address)
func (_Confirm *ConfirmCallerSession) ContractAddr() (common.Address, error) {
	return _Confirm.Contract.ContractAddr(&_Confirm.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_Confirm *ConfirmCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Confirm.contract.Call(opts, &out, "getOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_Confirm *ConfirmSession) GetOperator() (common.Address, error) {
	return _Confirm.Contract.GetOperator(&_Confirm.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0xe7f43c68.
//
// Solidity: function getOperator() view returns(address)
func (_Confirm *ConfirmCallerSession) GetOperator() (common.Address, error) {
	return _Confirm.Contract.GetOperator(&_Confirm.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Confirm *ConfirmCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Confirm.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Confirm *ConfirmSession) GetOwner() (common.Address, error) {
	return _Confirm.Contract.GetOwner(&_Confirm.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Confirm *ConfirmCallerSession) GetOwner() (common.Address, error) {
	return _Confirm.Contract.GetOwner(&_Confirm.CallOpts)
}

// ProfitInfo is a free data retrieval call binding the contract method 0x52135551.
//
// Solidity: function profitInfo(address ) view returns(uint256 profits, bytes32 trxHash, uint8 state)
func (_Confirm *ConfirmCaller) ProfitInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Profits *big.Int
	TrxHash [32]byte
	State   uint8
}, error) {
	var out []interface{}
	err := _Confirm.contract.Call(opts, &out, "profitInfo", arg0)

	outstruct := new(struct {
		Profits *big.Int
		TrxHash [32]byte
		State   uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Profits = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TrxHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.State = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// ProfitInfo is a free data retrieval call binding the contract method 0x52135551.
//
// Solidity: function profitInfo(address ) view returns(uint256 profits, bytes32 trxHash, uint8 state)
func (_Confirm *ConfirmSession) ProfitInfo(arg0 common.Address) (struct {
	Profits *big.Int
	TrxHash [32]byte
	State   uint8
}, error) {
	return _Confirm.Contract.ProfitInfo(&_Confirm.CallOpts, arg0)
}

// ProfitInfo is a free data retrieval call binding the contract method 0x52135551.
//
// Solidity: function profitInfo(address ) view returns(uint256 profits, bytes32 trxHash, uint8 state)
func (_Confirm *ConfirmCallerSession) ProfitInfo(arg0 common.Address) (struct {
	Profits *big.Int
	TrxHash [32]byte
	State   uint8
}, error) {
	return _Confirm.Contract.ProfitInfo(&_Confirm.CallOpts, arg0)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address operator_) returns()
func (_Confirm *ConfirmTransactor) ChangeOperator(opts *bind.TransactOpts, operator_ common.Address) (*types.Transaction, error) {
	return _Confirm.contract.Transact(opts, "changeOperator", operator_)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address operator_) returns()
func (_Confirm *ConfirmSession) ChangeOperator(operator_ common.Address) (*types.Transaction, error) {
	return _Confirm.Contract.ChangeOperator(&_Confirm.TransactOpts, operator_)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address operator_) returns()
func (_Confirm *ConfirmTransactorSession) ChangeOperator(operator_ common.Address) (*types.Transaction, error) {
	return _Confirm.Contract.ChangeOperator(&_Confirm.TransactOpts, operator_)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address owner_) returns()
func (_Confirm *ConfirmTransactor) ChangeOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Confirm.contract.Transact(opts, "changeOwner", owner_)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address owner_) returns()
func (_Confirm *ConfirmSession) ChangeOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Confirm.Contract.ChangeOwner(&_Confirm.TransactOpts, owner_)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address owner_) returns()
func (_Confirm *ConfirmTransactorSession) ChangeOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Confirm.Contract.ChangeOwner(&_Confirm.TransactOpts, owner_)
}

// ConfirmProfits is a paid mutator transaction binding the contract method 0x754bb559.
//
// Solidity: function confirmProfits() returns()
func (_Confirm *ConfirmTransactor) ConfirmProfits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Confirm.contract.Transact(opts, "confirmProfits")
}

// ConfirmProfits is a paid mutator transaction binding the contract method 0x754bb559.
//
// Solidity: function confirmProfits() returns()
func (_Confirm *ConfirmSession) ConfirmProfits() (*types.Transaction, error) {
	return _Confirm.Contract.ConfirmProfits(&_Confirm.TransactOpts)
}

// ConfirmProfits is a paid mutator transaction binding the contract method 0x754bb559.
//
// Solidity: function confirmProfits() returns()
func (_Confirm *ConfirmTransactorSession) ConfirmProfits() (*types.Transaction, error) {
	return _Confirm.Contract.ConfirmProfits(&_Confirm.TransactOpts)
}

// SetClaimed is a paid mutator transaction binding the contract method 0x3b16ed2d.
//
// Solidity: function setClaimed(address miner) returns()
func (_Confirm *ConfirmTransactor) SetClaimed(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _Confirm.contract.Transact(opts, "setClaimed", miner)
}

// SetClaimed is a paid mutator transaction binding the contract method 0x3b16ed2d.
//
// Solidity: function setClaimed(address miner) returns()
func (_Confirm *ConfirmSession) SetClaimed(miner common.Address) (*types.Transaction, error) {
	return _Confirm.Contract.SetClaimed(&_Confirm.TransactOpts, miner)
}

// SetClaimed is a paid mutator transaction binding the contract method 0x3b16ed2d.
//
// Solidity: function setClaimed(address miner) returns()
func (_Confirm *ConfirmTransactorSession) SetClaimed(miner common.Address) (*types.Transaction, error) {
	return _Confirm.Contract.SetClaimed(&_Confirm.TransactOpts, miner)
}

// SetTrx is a paid mutator transaction binding the contract method 0xeea85d58.
//
// Solidity: function setTrx(address miner, bytes32 trxHash) returns()
func (_Confirm *ConfirmTransactor) SetTrx(opts *bind.TransactOpts, miner common.Address, trxHash [32]byte) (*types.Transaction, error) {
	return _Confirm.contract.Transact(opts, "setTrx", miner, trxHash)
}

// SetTrx is a paid mutator transaction binding the contract method 0xeea85d58.
//
// Solidity: function setTrx(address miner, bytes32 trxHash) returns()
func (_Confirm *ConfirmSession) SetTrx(miner common.Address, trxHash [32]byte) (*types.Transaction, error) {
	return _Confirm.Contract.SetTrx(&_Confirm.TransactOpts, miner, trxHash)
}

// SetTrx is a paid mutator transaction binding the contract method 0xeea85d58.
//
// Solidity: function setTrx(address miner, bytes32 trxHash) returns()
func (_Confirm *ConfirmTransactorSession) SetTrx(miner common.Address, trxHash [32]byte) (*types.Transaction, error) {
	return _Confirm.Contract.SetTrx(&_Confirm.TransactOpts, miner, trxHash)
}

// ConfirmConfirmedIterator is returned from FilterConfirmed and is used to iterate over the raw logs and unpacked data for Confirmed events raised by the Confirm contract.
type ConfirmConfirmedIterator struct {
	Event *ConfirmConfirmed // Event containing the contract specifics and raw log

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
func (it *ConfirmConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfirmConfirmed)
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
		it.Event = new(ConfirmConfirmed)
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
func (it *ConfirmConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfirmConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfirmConfirmed represents a Confirmed event raised by the Confirm contract.
type ConfirmConfirmed struct {
	Miner   common.Address
	Profits *big.Int
	State   uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConfirmed is a free log retrieval operation binding the contract event 0xb1ab7584dee9fc349f34e594504b437348e4cf9d6c6a4af63690160f1cd484d1.
//
// Solidity: event Confirmed(address miner, uint256 profits, uint8 state)
func (_Confirm *ConfirmFilterer) FilterConfirmed(opts *bind.FilterOpts) (*ConfirmConfirmedIterator, error) {

	logs, sub, err := _Confirm.contract.FilterLogs(opts, "Confirmed")
	if err != nil {
		return nil, err
	}
	return &ConfirmConfirmedIterator{contract: _Confirm.contract, event: "Confirmed", logs: logs, sub: sub}, nil
}

// WatchConfirmed is a free log subscription operation binding the contract event 0xb1ab7584dee9fc349f34e594504b437348e4cf9d6c6a4af63690160f1cd484d1.
//
// Solidity: event Confirmed(address miner, uint256 profits, uint8 state)
func (_Confirm *ConfirmFilterer) WatchConfirmed(opts *bind.WatchOpts, sink chan<- *ConfirmConfirmed) (event.Subscription, error) {

	logs, sub, err := _Confirm.contract.WatchLogs(opts, "Confirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfirmConfirmed)
				if err := _Confirm.contract.UnpackLog(event, "Confirmed", log); err != nil {
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

// ParseConfirmed is a log parse operation binding the contract event 0xb1ab7584dee9fc349f34e594504b437348e4cf9d6c6a4af63690160f1cd484d1.
//
// Solidity: event Confirmed(address miner, uint256 profits, uint8 state)
func (_Confirm *ConfirmFilterer) ParseConfirmed(log types.Log) (*ConfirmConfirmed, error) {
	event := new(ConfirmConfirmed)
	if err := _Confirm.contract.UnpackLog(event, "Confirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
