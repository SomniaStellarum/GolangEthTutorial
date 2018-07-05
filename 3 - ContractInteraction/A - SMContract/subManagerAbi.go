// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/TxHash/go-ethereum"
	"github.com/TxHash/go-ethereum/accounts/abi"
	"github.com/TxHash/go-ethereum/accounts/abi/bind"
	"github.com/TxHash/go-ethereum/common"
	"github.com/TxHash/go-ethereum/core/types"
	"github.com/TxHash/go-ethereum/event"
)

// SubscriptionManagerABI is the input ABI used to generate the binding from.
const SubscriptionManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"callBill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"billerAddress\",\"type\":\"address\"},{\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"setBillRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"billers\",\"outputs\":[{\"name\":\"rate\",\"type\":\"uint256\"},{\"name\":\"lastBillTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_Owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"billerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"BillRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"billerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountBilled\",\"type\":\"uint256\"}],\"name\":\"BillProcessed\",\"type\":\"event\"}]"

// SubscriptionManagerBin is the compiled bytecode used for deploying new contracts.
const SubscriptionManagerBin = `0x608060405234801561001057600080fd5b50604051602080610502833981016040525160008054600160a060020a03909216600160a060020a03199092169190911790556104b0806100526000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632e1a7d4d81146100a85780634f016a2e146100c2578063911f5f98146100d75780639dbbdc3c146100fb578063b4a99a4e14610135575b6040805133815234602082015281517fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c929181900390910190a1005b3480156100b457600080fd5b506100c0600435610166565b005b3480156100ce57600080fd5b506100c0610283565b3480156100e357600080fd5b506100c0600160a060020a036004351660243561030b565b34801561010757600080fd5b5061011c600160a060020a0360043516610384565b6040805192835260208301919091528051918290030190f35b34801561014157600080fd5b5061014a61039d565b60408051600160a060020a039092168252519081900360200190f35b600054600160a060020a0316331461017d57600080fd5b303181111561021357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603d60248201527f52657175657374656420776974686472617720616d6f756e742065786365656460448201527f7320537562736372697074696f6e4d616e616765722062616c616e6365000000606482015290519081900360840190fd5b60008054604051600160a060020a039091169183156108fc02918491818181858888f1935050505015801561024c573d6000803e3d6000fd5b506040805182815290517f5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d9181900360200190a150565b33600081815260016020526040812054116102ff57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f596f7520617265206e6f7420617070726f76656420746f2062696c6c2e000000604482015290519081900360640190fd5b610308816103ac565b50565b600054600160a060020a0316331461032257600080fd5b61032b826103ac565b600160a060020a0382166000818152600160209081526040918290208490558151928352820183905280517fe9171f6a4635bd828cbca9708580c62ac52d6532512942c94e3dd9b557a293e89281900390910190a15050565b6001602081905260009182526040909120805491015482565b600054600160a060020a031681565b600160a060020a038116600090815260016020526040812054819015610426575050600160a060020a0381166000818152600160208190526040808320918201549154905142929092039390840292909183156108fc0291849190818181858888f19350505050158015610424573d6000803e3d6000fd5b505b600160a060020a038316600081815260016020818152604092839020429201919091558151928352820183905280517f251ed24b2f8db3cbf768d6cc0512d3811977ad8589afec9f4cae2a7f24e7ea139281900390910190a15050505600a165627a7a72305820f656a693d32bc0090840a9b6ed3db51461dbd95ef6e3c3439361063f60a283f20029`

// DeploySubscriptionManager deploys a new Ethereum contract, binding an instance of SubscriptionManager to it.
func DeploySubscriptionManager(auth *bind.TransactOpts, backend bind.ContractBackend, _Owner common.Address) (common.Address, *types.Transaction, *SubscriptionManager, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscriptionManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SubscriptionManagerBin), backend, _Owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubscriptionManager{SubscriptionManagerCaller: SubscriptionManagerCaller{contract: contract}, SubscriptionManagerTransactor: SubscriptionManagerTransactor{contract: contract}, SubscriptionManagerFilterer: SubscriptionManagerFilterer{contract: contract}}, nil
}

// SubscriptionManager is an auto generated Go binding around an Ethereum contract.
type SubscriptionManager struct {
	SubscriptionManagerCaller     // Read-only binding to the contract
	SubscriptionManagerTransactor // Write-only binding to the contract
	SubscriptionManagerFilterer   // Log filterer for contract events
}

// SubscriptionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubscriptionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubscriptionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubscriptionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubscriptionManagerSession struct {
	Contract     *SubscriptionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SubscriptionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubscriptionManagerCallerSession struct {
	Contract *SubscriptionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SubscriptionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubscriptionManagerTransactorSession struct {
	Contract     *SubscriptionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SubscriptionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubscriptionManagerRaw struct {
	Contract *SubscriptionManager // Generic contract binding to access the raw methods on
}

// SubscriptionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubscriptionManagerCallerRaw struct {
	Contract *SubscriptionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// SubscriptionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubscriptionManagerTransactorRaw struct {
	Contract *SubscriptionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubscriptionManager creates a new instance of SubscriptionManager, bound to a specific deployed contract.
func NewSubscriptionManager(address common.Address, backend bind.ContractBackend) (*SubscriptionManager, error) {
	contract, err := bindSubscriptionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubscriptionManager{SubscriptionManagerCaller: SubscriptionManagerCaller{contract: contract}, SubscriptionManagerTransactor: SubscriptionManagerTransactor{contract: contract}, SubscriptionManagerFilterer: SubscriptionManagerFilterer{contract: contract}}, nil
}

// NewSubscriptionManagerCaller creates a new read-only instance of SubscriptionManager, bound to a specific deployed contract.
func NewSubscriptionManagerCaller(address common.Address, caller bind.ContractCaller) (*SubscriptionManagerCaller, error) {
	contract, err := bindSubscriptionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionManagerCaller{contract: contract}, nil
}

// NewSubscriptionManagerTransactor creates a new write-only instance of SubscriptionManager, bound to a specific deployed contract.
func NewSubscriptionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*SubscriptionManagerTransactor, error) {
	contract, err := bindSubscriptionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionManagerTransactor{contract: contract}, nil
}

// NewSubscriptionManagerFilterer creates a new log filterer instance of SubscriptionManager, bound to a specific deployed contract.
func NewSubscriptionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*SubscriptionManagerFilterer, error) {
	contract, err := bindSubscriptionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubscriptionManagerFilterer{contract: contract}, nil
}

// bindSubscriptionManager binds a generic wrapper to an already deployed contract.
func bindSubscriptionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscriptionManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubscriptionManager *SubscriptionManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubscriptionManager.Contract.SubscriptionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubscriptionManager *SubscriptionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionManager.Contract.SubscriptionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubscriptionManager *SubscriptionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscriptionManager.Contract.SubscriptionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubscriptionManager *SubscriptionManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubscriptionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubscriptionManager *SubscriptionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubscriptionManager *SubscriptionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscriptionManager.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() constant returns(address)
func (_SubscriptionManager *SubscriptionManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SubscriptionManager.contract.Call(opts, out, "Owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() constant returns(address)
func (_SubscriptionManager *SubscriptionManagerSession) Owner() (common.Address, error) {
	return _SubscriptionManager.Contract.Owner(&_SubscriptionManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() constant returns(address)
func (_SubscriptionManager *SubscriptionManagerCallerSession) Owner() (common.Address, error) {
	return _SubscriptionManager.Contract.Owner(&_SubscriptionManager.CallOpts)
}

// Billers is a free data retrieval call binding the contract method 0x9dbbdc3c.
//
// Solidity: function billers( address) constant returns(rate uint256, lastBillTime uint256)
func (_SubscriptionManager *SubscriptionManagerCaller) Billers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Rate         *big.Int
	LastBillTime *big.Int
}, error) {
	ret := new(struct {
		Rate         *big.Int
		LastBillTime *big.Int
	})
	out := ret
	err := _SubscriptionManager.contract.Call(opts, out, "billers", arg0)
	return *ret, err
}

// Billers is a free data retrieval call binding the contract method 0x9dbbdc3c.
//
// Solidity: function billers( address) constant returns(rate uint256, lastBillTime uint256)
func (_SubscriptionManager *SubscriptionManagerSession) Billers(arg0 common.Address) (struct {
	Rate         *big.Int
	LastBillTime *big.Int
}, error) {
	return _SubscriptionManager.Contract.Billers(&_SubscriptionManager.CallOpts, arg0)
}

// Billers is a free data retrieval call binding the contract method 0x9dbbdc3c.
//
// Solidity: function billers( address) constant returns(rate uint256, lastBillTime uint256)
func (_SubscriptionManager *SubscriptionManagerCallerSession) Billers(arg0 common.Address) (struct {
	Rate         *big.Int
	LastBillTime *big.Int
}, error) {
	return _SubscriptionManager.Contract.Billers(&_SubscriptionManager.CallOpts, arg0)
}

// CallBill is a paid mutator transaction binding the contract method 0x4f016a2e.
//
// Solidity: function callBill() returns()
func (_SubscriptionManager *SubscriptionManagerTransactor) CallBill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionManager.contract.Transact(opts, "callBill")
}

// CallBill is a paid mutator transaction binding the contract method 0x4f016a2e.
//
// Solidity: function callBill() returns()
func (_SubscriptionManager *SubscriptionManagerSession) CallBill() (*types.Transaction, error) {
	return _SubscriptionManager.Contract.CallBill(&_SubscriptionManager.TransactOpts)
}

// CallBill is a paid mutator transaction binding the contract method 0x4f016a2e.
//
// Solidity: function callBill() returns()
func (_SubscriptionManager *SubscriptionManagerTransactorSession) CallBill() (*types.Transaction, error) {
	return _SubscriptionManager.Contract.CallBill(&_SubscriptionManager.TransactOpts)
}

// SetBillRate is a paid mutator transaction binding the contract method 0x911f5f98.
//
// Solidity: function setBillRate(billerAddress address, newRate uint256) returns()
func (_SubscriptionManager *SubscriptionManagerTransactor) SetBillRate(opts *bind.TransactOpts, billerAddress common.Address, newRate *big.Int) (*types.Transaction, error) {
	return _SubscriptionManager.contract.Transact(opts, "setBillRate", billerAddress, newRate)
}

// SetBillRate is a paid mutator transaction binding the contract method 0x911f5f98.
//
// Solidity: function setBillRate(billerAddress address, newRate uint256) returns()
func (_SubscriptionManager *SubscriptionManagerSession) SetBillRate(billerAddress common.Address, newRate *big.Int) (*types.Transaction, error) {
	return _SubscriptionManager.Contract.SetBillRate(&_SubscriptionManager.TransactOpts, billerAddress, newRate)
}

// SetBillRate is a paid mutator transaction binding the contract method 0x911f5f98.
//
// Solidity: function setBillRate(billerAddress address, newRate uint256) returns()
func (_SubscriptionManager *SubscriptionManagerTransactorSession) SetBillRate(billerAddress common.Address, newRate *big.Int) (*types.Transaction, error) {
	return _SubscriptionManager.Contract.SetBillRate(&_SubscriptionManager.TransactOpts, billerAddress, newRate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_SubscriptionManager *SubscriptionManagerTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SubscriptionManager.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_SubscriptionManager *SubscriptionManagerSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _SubscriptionManager.Contract.Withdraw(&_SubscriptionManager.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_SubscriptionManager *SubscriptionManagerTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _SubscriptionManager.Contract.Withdraw(&_SubscriptionManager.TransactOpts, amount)
}

// SubscriptionManagerBillProcessedIterator is returned from FilterBillProcessed and is used to iterate over the raw logs and unpacked data for BillProcessed events raised by the SubscriptionManager contract.
type SubscriptionManagerBillProcessedIterator struct {
	Event *SubscriptionManagerBillProcessed // Event containing the contract specifics and raw log

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
func (it *SubscriptionManagerBillProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionManagerBillProcessed)
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
		it.Event = new(SubscriptionManagerBillProcessed)
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
func (it *SubscriptionManagerBillProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionManagerBillProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionManagerBillProcessed represents a BillProcessed event raised by the SubscriptionManager contract.
type SubscriptionManagerBillProcessed struct {
	BillerAddress common.Address
	AmountBilled  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBillProcessed is a free log retrieval operation binding the contract event 0x251ed24b2f8db3cbf768d6cc0512d3811977ad8589afec9f4cae2a7f24e7ea13.
//
// Solidity: e BillProcessed(billerAddress address, amountBilled uint256)
func (_SubscriptionManager *SubscriptionManagerFilterer) FilterBillProcessed(opts *bind.FilterOpts) (*SubscriptionManagerBillProcessedIterator, error) {

	logs, sub, err := _SubscriptionManager.contract.FilterLogs(opts, "BillProcessed")
	if err != nil {
		return nil, err
	}
	return &SubscriptionManagerBillProcessedIterator{contract: _SubscriptionManager.contract, event: "BillProcessed", logs: logs, sub: sub}, nil
}

// WatchBillProcessed is a free log subscription operation binding the contract event 0x251ed24b2f8db3cbf768d6cc0512d3811977ad8589afec9f4cae2a7f24e7ea13.
//
// Solidity: e BillProcessed(billerAddress address, amountBilled uint256)
func (_SubscriptionManager *SubscriptionManagerFilterer) WatchBillProcessed(opts *bind.WatchOpts, sink chan<- *SubscriptionManagerBillProcessed) (event.Subscription, error) {

	logs, sub, err := _SubscriptionManager.contract.WatchLogs(opts, "BillProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionManagerBillProcessed)
				if err := _SubscriptionManager.contract.UnpackLog(event, "BillProcessed", log); err != nil {
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

// SubscriptionManagerBillRateSetIterator is returned from FilterBillRateSet and is used to iterate over the raw logs and unpacked data for BillRateSet events raised by the SubscriptionManager contract.
type SubscriptionManagerBillRateSetIterator struct {
	Event *SubscriptionManagerBillRateSet // Event containing the contract specifics and raw log

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
func (it *SubscriptionManagerBillRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionManagerBillRateSet)
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
		it.Event = new(SubscriptionManagerBillRateSet)
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
func (it *SubscriptionManagerBillRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionManagerBillRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionManagerBillRateSet represents a BillRateSet event raised by the SubscriptionManager contract.
type SubscriptionManagerBillRateSet struct {
	BillerAddress common.Address
	NewRate       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBillRateSet is a free log retrieval operation binding the contract event 0xe9171f6a4635bd828cbca9708580c62ac52d6532512942c94e3dd9b557a293e8.
//
// Solidity: e BillRateSet(billerAddress address, newRate uint256)
func (_SubscriptionManager *SubscriptionManagerFilterer) FilterBillRateSet(opts *bind.FilterOpts) (*SubscriptionManagerBillRateSetIterator, error) {

	logs, sub, err := _SubscriptionManager.contract.FilterLogs(opts, "BillRateSet")
	if err != nil {
		return nil, err
	}
	return &SubscriptionManagerBillRateSetIterator{contract: _SubscriptionManager.contract, event: "BillRateSet", logs: logs, sub: sub}, nil
}

// WatchBillRateSet is a free log subscription operation binding the contract event 0xe9171f6a4635bd828cbca9708580c62ac52d6532512942c94e3dd9b557a293e8.
//
// Solidity: e BillRateSet(billerAddress address, newRate uint256)
func (_SubscriptionManager *SubscriptionManagerFilterer) WatchBillRateSet(opts *bind.WatchOpts, sink chan<- *SubscriptionManagerBillRateSet) (event.Subscription, error) {

	logs, sub, err := _SubscriptionManager.contract.WatchLogs(opts, "BillRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionManagerBillRateSet)
				if err := _SubscriptionManager.contract.UnpackLog(event, "BillRateSet", log); err != nil {
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

// SubscriptionManagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SubscriptionManager contract.
type SubscriptionManagerDepositIterator struct {
	Event *SubscriptionManagerDeposit // Event containing the contract specifics and raw log

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
func (it *SubscriptionManagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionManagerDeposit)
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
		it.Event = new(SubscriptionManagerDeposit)
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
func (it *SubscriptionManagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionManagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionManagerDeposit represents a Deposit event raised by the SubscriptionManager contract.
type SubscriptionManagerDeposit struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: e Deposit(from address, amount uint256)
func (_SubscriptionManager *SubscriptionManagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*SubscriptionManagerDepositIterator, error) {

	logs, sub, err := _SubscriptionManager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &SubscriptionManagerDepositIterator{contract: _SubscriptionManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: e Deposit(from address, amount uint256)
func (_SubscriptionManager *SubscriptionManagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SubscriptionManagerDeposit) (event.Subscription, error) {

	logs, sub, err := _SubscriptionManager.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionManagerDeposit)
				if err := _SubscriptionManager.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// SubscriptionManagerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SubscriptionManager contract.
type SubscriptionManagerWithdrawIterator struct {
	Event *SubscriptionManagerWithdraw // Event containing the contract specifics and raw log

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
func (it *SubscriptionManagerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubscriptionManagerWithdraw)
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
		it.Event = new(SubscriptionManagerWithdraw)
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
func (it *SubscriptionManagerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubscriptionManagerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubscriptionManagerWithdraw represents a Withdraw event raised by the SubscriptionManager contract.
type SubscriptionManagerWithdraw struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: e Withdraw(amount uint256)
func (_SubscriptionManager *SubscriptionManagerFilterer) FilterWithdraw(opts *bind.FilterOpts) (*SubscriptionManagerWithdrawIterator, error) {

	logs, sub, err := _SubscriptionManager.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &SubscriptionManagerWithdrawIterator{contract: _SubscriptionManager.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d.
//
// Solidity: e Withdraw(amount uint256)
func (_SubscriptionManager *SubscriptionManagerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SubscriptionManagerWithdraw) (event.Subscription, error) {

	logs, sub, err := _SubscriptionManager.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubscriptionManagerWithdraw)
				if err := _SubscriptionManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// SubscriptionManagerFactoryABI is the input ABI used to generate the binding from.
const SubscriptionManagerFactoryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"newSubscriptionManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"subscriptionManagerAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SubscriptionManagerFactoryBin is the compiled bytecode used for deploying new contracts.
const SubscriptionManagerFactoryBin = `0x608060405234801561001057600080fd5b506106c8806100206000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632301103f81146100505780632f8daa6f1461008e575b600080fd5b34801561005c57600080fd5b506100656100a6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561009a57600080fd5b50610065600435610155565b600080336100b261018a565b73ffffffffffffffffffffffffffffffffffffffff909116815260405190819003602001906000f0801580156100ec573d6000803e3d6000fd5b50600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff831617905592915050565b600080548290811061016357fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b6040516105028061019b833901905600608060405234801561001057600080fd5b50604051602080610502833981016040525160008054600160a060020a03909216600160a060020a03199092169190911790556104b0806100526000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632e1a7d4d81146100a85780634f016a2e146100c2578063911f5f98146100d75780639dbbdc3c146100fb578063b4a99a4e14610135575b6040805133815234602082015281517fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c929181900390910190a1005b3480156100b457600080fd5b506100c0600435610166565b005b3480156100ce57600080fd5b506100c0610283565b3480156100e357600080fd5b506100c0600160a060020a036004351660243561030b565b34801561010757600080fd5b5061011c600160a060020a0360043516610384565b6040805192835260208301919091528051918290030190f35b34801561014157600080fd5b5061014a61039d565b60408051600160a060020a039092168252519081900360200190f35b600054600160a060020a0316331461017d57600080fd5b303181111561021357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603d60248201527f52657175657374656420776974686472617720616d6f756e742065786365656460448201527f7320537562736372697074696f6e4d616e616765722062616c616e6365000000606482015290519081900360840190fd5b60008054604051600160a060020a039091169183156108fc02918491818181858888f1935050505015801561024c573d6000803e3d6000fd5b506040805182815290517f5b6b431d4476a211bb7d41c20d1aab9ae2321deee0d20be3d9fc9b1093fa6e3d9181900360200190a150565b33600081815260016020526040812054116102ff57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f596f7520617265206e6f7420617070726f76656420746f2062696c6c2e000000604482015290519081900360640190fd5b610308816103ac565b50565b600054600160a060020a0316331461032257600080fd5b61032b826103ac565b600160a060020a0382166000818152600160209081526040918290208490558151928352820183905280517fe9171f6a4635bd828cbca9708580c62ac52d6532512942c94e3dd9b557a293e89281900390910190a15050565b6001602081905260009182526040909120805491015482565b600054600160a060020a031681565b600160a060020a038116600090815260016020526040812054819015610426575050600160a060020a0381166000818152600160208190526040808320918201549154905142929092039390840292909183156108fc0291849190818181858888f19350505050158015610424573d6000803e3d6000fd5b505b600160a060020a038316600081815260016020818152604092839020429201919091558151928352820183905280517f251ed24b2f8db3cbf768d6cc0512d3811977ad8589afec9f4cae2a7f24e7ea139281900390910190a15050505600a165627a7a72305820f656a693d32bc0090840a9b6ed3db51461dbd95ef6e3c3439361063f60a283f20029a165627a7a723058200d8854b33e74c27187e1639f42d167db0a183638f73aea1359d154d308f0d8d10029`

// DeploySubscriptionManagerFactory deploys a new Ethereum contract, binding an instance of SubscriptionManagerFactory to it.
func DeploySubscriptionManagerFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SubscriptionManagerFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscriptionManagerFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SubscriptionManagerFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubscriptionManagerFactory{SubscriptionManagerFactoryCaller: SubscriptionManagerFactoryCaller{contract: contract}, SubscriptionManagerFactoryTransactor: SubscriptionManagerFactoryTransactor{contract: contract}, SubscriptionManagerFactoryFilterer: SubscriptionManagerFactoryFilterer{contract: contract}}, nil
}

// SubscriptionManagerFactory is an auto generated Go binding around an Ethereum contract.
type SubscriptionManagerFactory struct {
	SubscriptionManagerFactoryCaller     // Read-only binding to the contract
	SubscriptionManagerFactoryTransactor // Write-only binding to the contract
	SubscriptionManagerFactoryFilterer   // Log filterer for contract events
}

// SubscriptionManagerFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubscriptionManagerFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionManagerFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubscriptionManagerFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionManagerFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubscriptionManagerFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubscriptionManagerFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubscriptionManagerFactorySession struct {
	Contract     *SubscriptionManagerFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SubscriptionManagerFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubscriptionManagerFactoryCallerSession struct {
	Contract *SubscriptionManagerFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// SubscriptionManagerFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubscriptionManagerFactoryTransactorSession struct {
	Contract     *SubscriptionManagerFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// SubscriptionManagerFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubscriptionManagerFactoryRaw struct {
	Contract *SubscriptionManagerFactory // Generic contract binding to access the raw methods on
}

// SubscriptionManagerFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubscriptionManagerFactoryCallerRaw struct {
	Contract *SubscriptionManagerFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SubscriptionManagerFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubscriptionManagerFactoryTransactorRaw struct {
	Contract *SubscriptionManagerFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubscriptionManagerFactory creates a new instance of SubscriptionManagerFactory, bound to a specific deployed contract.
func NewSubscriptionManagerFactory(address common.Address, backend bind.ContractBackend) (*SubscriptionManagerFactory, error) {
	contract, err := bindSubscriptionManagerFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubscriptionManagerFactory{SubscriptionManagerFactoryCaller: SubscriptionManagerFactoryCaller{contract: contract}, SubscriptionManagerFactoryTransactor: SubscriptionManagerFactoryTransactor{contract: contract}, SubscriptionManagerFactoryFilterer: SubscriptionManagerFactoryFilterer{contract: contract}}, nil
}

// NewSubscriptionManagerFactoryCaller creates a new read-only instance of SubscriptionManagerFactory, bound to a specific deployed contract.
func NewSubscriptionManagerFactoryCaller(address common.Address, caller bind.ContractCaller) (*SubscriptionManagerFactoryCaller, error) {
	contract, err := bindSubscriptionManagerFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionManagerFactoryCaller{contract: contract}, nil
}

// NewSubscriptionManagerFactoryTransactor creates a new write-only instance of SubscriptionManagerFactory, bound to a specific deployed contract.
func NewSubscriptionManagerFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SubscriptionManagerFactoryTransactor, error) {
	contract, err := bindSubscriptionManagerFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubscriptionManagerFactoryTransactor{contract: contract}, nil
}

// NewSubscriptionManagerFactoryFilterer creates a new log filterer instance of SubscriptionManagerFactory, bound to a specific deployed contract.
func NewSubscriptionManagerFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SubscriptionManagerFactoryFilterer, error) {
	contract, err := bindSubscriptionManagerFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubscriptionManagerFactoryFilterer{contract: contract}, nil
}

// bindSubscriptionManagerFactory binds a generic wrapper to an already deployed contract.
func bindSubscriptionManagerFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubscriptionManagerFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubscriptionManagerFactory *SubscriptionManagerFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubscriptionManagerFactory.Contract.SubscriptionManagerFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubscriptionManagerFactory *SubscriptionManagerFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionManagerFactory.Contract.SubscriptionManagerFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubscriptionManagerFactory *SubscriptionManagerFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscriptionManagerFactory.Contract.SubscriptionManagerFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubscriptionManagerFactory *SubscriptionManagerFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SubscriptionManagerFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubscriptionManagerFactory *SubscriptionManagerFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionManagerFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubscriptionManagerFactory *SubscriptionManagerFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubscriptionManagerFactory.Contract.contract.Transact(opts, method, params...)
}

// SubscriptionManagerAddresses is a free data retrieval call binding the contract method 0x2f8daa6f.
//
// Solidity: function subscriptionManagerAddresses( uint256) constant returns(address)
func (_SubscriptionManagerFactory *SubscriptionManagerFactoryCaller) SubscriptionManagerAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SubscriptionManagerFactory.contract.Call(opts, out, "subscriptionManagerAddresses", arg0)
	return *ret0, err
}

// SubscriptionManagerAddresses is a free data retrieval call binding the contract method 0x2f8daa6f.
//
// Solidity: function subscriptionManagerAddresses( uint256) constant returns(address)
func (_SubscriptionManagerFactory *SubscriptionManagerFactorySession) SubscriptionManagerAddresses(arg0 *big.Int) (common.Address, error) {
	return _SubscriptionManagerFactory.Contract.SubscriptionManagerAddresses(&_SubscriptionManagerFactory.CallOpts, arg0)
}

// SubscriptionManagerAddresses is a free data retrieval call binding the contract method 0x2f8daa6f.
//
// Solidity: function subscriptionManagerAddresses( uint256) constant returns(address)
func (_SubscriptionManagerFactory *SubscriptionManagerFactoryCallerSession) SubscriptionManagerAddresses(arg0 *big.Int) (common.Address, error) {
	return _SubscriptionManagerFactory.Contract.SubscriptionManagerAddresses(&_SubscriptionManagerFactory.CallOpts, arg0)
}

// NewSubscriptionManager is a paid mutator transaction binding the contract method 0x2301103f.
//
// Solidity: function newSubscriptionManager() returns(address)
func (_SubscriptionManagerFactory *SubscriptionManagerFactoryTransactor) NewSubscriptionManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubscriptionManagerFactory.contract.Transact(opts, "newSubscriptionManager")
}

// NewSubscriptionManager is a paid mutator transaction binding the contract method 0x2301103f.
//
// Solidity: function newSubscriptionManager() returns(address)
func (_SubscriptionManagerFactory *SubscriptionManagerFactorySession) NewSubscriptionManager() (*types.Transaction, error) {
	return _SubscriptionManagerFactory.Contract.NewSubscriptionManager(&_SubscriptionManagerFactory.TransactOpts)
}

// NewSubscriptionManager is a paid mutator transaction binding the contract method 0x2301103f.
//
// Solidity: function newSubscriptionManager() returns(address)
func (_SubscriptionManagerFactory *SubscriptionManagerFactoryTransactorSession) NewSubscriptionManager() (*types.Transaction, error) {
	return _SubscriptionManagerFactory.Contract.NewSubscriptionManager(&_SubscriptionManagerFactory.TransactOpts)
}
