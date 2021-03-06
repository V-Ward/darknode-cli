// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
const BasicTokenBin = `0x608060405234801561001057600080fd5b5061027c806100206000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005b57806370a0823114610082578063a9059cbb146100b0575b600080fd5b34801561006757600080fd5b506100706100f5565b60408051918252519081900360200190f35b34801561008e57600080fd5b5061007073ffffffffffffffffffffffffffffffffffffffff600435166100fb565b3480156100bc57600080fd5b506100e173ffffffffffffffffffffffffffffffffffffffff60043516602435610123565b604080519115158252519081900360200190f35b60015490565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b600073ffffffffffffffffffffffffffffffffffffffff8316151561014757600080fd5b3360009081526020819052604090205482111561016357600080fd5b33600090815260208190526040902054610183908363ffffffff61022b16565b336000908152602081905260408082209290925573ffffffffffffffffffffffffffffffffffffffff8516815220546101c2908363ffffffff61023d16565b73ffffffffffffffffffffffffffffffffffffffff8416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b60008282111561023757fe5b50900390565b8181018281101561024a57fe5b929150505600a165627a7a72305820526574b8d11bdfc7c6e031d92b9187a155fb314df8ab95d9b0b5654deff96e8e0029`

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
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
		it.Event = new(BasicTokenTransfer)
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
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// BurnableTokenABI is the input ABI used to generate the binding from.
const BurnableTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BurnableTokenBin is the compiled bytecode used for deploying new contracts.
const BurnableTokenBin = `0x608060405234801561001057600080fd5b50610361806100206000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461006657806342966c681461008d57806370a08231146100a7578063a9059cbb146100c8575b600080fd5b34801561007257600080fd5b5061007b610100565b60408051918252519081900360200190f35b34801561009957600080fd5b506100a5600435610106565b005b3480156100b357600080fd5b5061007b600160a060020a0360043516610113565b3480156100d457600080fd5b506100ec600160a060020a036004351660243561012e565b604080519115158252519081900360200190f35b60015490565b610110338261020f565b50565b600160a060020a031660009081526020819052604090205490565b6000600160a060020a038316151561014557600080fd5b3360009081526020819052604090205482111561016157600080fd5b33600090815260208190526040902054610181908363ffffffff61031016565b3360009081526020819052604080822092909255600160a060020a038516815220546101b3908363ffffffff61032216565b600160a060020a038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b600160a060020a03821660009081526020819052604090205481111561023457600080fd5b600160a060020a03821660009081526020819052604090205461025d908263ffffffff61031016565b600160a060020a038316600090815260208190526040902055600154610289908263ffffffff61031016565b600155604080518281529051600160a060020a038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518281529051600091600160a060020a038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b60008282111561031c57fe5b50900390565b8181018281101561032f57fe5b929150505600a165627a7a723058200f5a2c87f9eff7f70d0c84db3c0a0839ee36e52b479209cc5980d2f5fe6fbd160029`

// DeployBurnableToken deploys a new Ethereum contract, binding an instance of BurnableToken to it.
func DeployBurnableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BurnableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnableToken{BurnableTokenCaller: BurnableTokenCaller{contract: contract}, BurnableTokenTransactor: BurnableTokenTransactor{contract: contract}, BurnableTokenFilterer: BurnableTokenFilterer{contract: contract}}, nil
}

// BurnableToken is an auto generated Go binding around an Ethereum contract.
type BurnableToken struct {
	BurnableTokenCaller     // Read-only binding to the contract
	BurnableTokenTransactor // Write-only binding to the contract
	BurnableTokenFilterer   // Log filterer for contract events
}

// BurnableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnableTokenSession struct {
	Contract     *BurnableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnableTokenCallerSession struct {
	Contract *BurnableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BurnableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnableTokenTransactorSession struct {
	Contract     *BurnableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BurnableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnableTokenRaw struct {
	Contract *BurnableToken // Generic contract binding to access the raw methods on
}

// BurnableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnableTokenCallerRaw struct {
	Contract *BurnableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BurnableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnableTokenTransactorRaw struct {
	Contract *BurnableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnableToken creates a new instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableToken(address common.Address, backend bind.ContractBackend) (*BurnableToken, error) {
	contract, err := bindBurnableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnableToken{BurnableTokenCaller: BurnableTokenCaller{contract: contract}, BurnableTokenTransactor: BurnableTokenTransactor{contract: contract}, BurnableTokenFilterer: BurnableTokenFilterer{contract: contract}}, nil
}

// NewBurnableTokenCaller creates a new read-only instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenCaller(address common.Address, caller bind.ContractCaller) (*BurnableTokenCaller, error) {
	contract, err := bindBurnableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenCaller{contract: contract}, nil
}

// NewBurnableTokenTransactor creates a new write-only instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnableTokenTransactor, error) {
	contract, err := bindBurnableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenTransactor{contract: contract}, nil
}

// NewBurnableTokenFilterer creates a new log filterer instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnableTokenFilterer, error) {
	contract, err := bindBurnableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenFilterer{contract: contract}, nil
}

// bindBurnableToken binds a generic wrapper to an already deployed contract.
func bindBurnableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableToken *BurnableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableToken.Contract.BurnableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableToken *BurnableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableToken *BurnableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableToken *BurnableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableToken *BurnableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableToken *BurnableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BurnableToken *BurnableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BurnableToken *BurnableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.BalanceOf(&_BurnableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BurnableToken *BurnableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.BalanceOf(&_BurnableToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenSession) TotalSupply() (*big.Int, error) {
	return _BurnableToken.Contract.TotalSupply(&_BurnableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnableToken.Contract.TotalSupply(&_BurnableToken.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableToken *BurnableTokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableToken *BurnableTokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Burn(&_BurnableToken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableToken *BurnableTokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Burn(&_BurnableToken.TransactOpts, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Transfer(&_BurnableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Transfer(&_BurnableToken.TransactOpts, _to, _value)
}

// BurnableTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BurnableToken contract.
type BurnableTokenBurnIterator struct {
	Event *BurnableTokenBurn // Event containing the contract specifics and raw log

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
func (it *BurnableTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenBurn)
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
		it.Event = new(BurnableTokenBurn)
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
func (it *BurnableTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenBurn represents a Burn event raised by the BurnableToken contract.
type BurnableTokenBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*BurnableTokenBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenBurnIterator{contract: _BurnableToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BurnableTokenBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenBurn)
				if err := _BurnableToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// BurnableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BurnableToken contract.
type BurnableTokenTransferIterator struct {
	Event *BurnableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BurnableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenTransfer)
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
		it.Event = new(BurnableTokenTransfer)
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
func (it *BurnableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenTransfer represents a Transfer event raised by the BurnableToken contract.
type BurnableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenTransferIterator{contract: _BurnableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenTransfer)
				if err := _BurnableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// DarknodeRegistryABI is the input ABI used to generate the binding from.
const DarknodeRegistryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"numDarknodesNextEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nextMinimumBond\",\"type\":\"uint256\"}],\"name\":\"updateMinimumBond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numDarknodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"bytes20\"}],\"name\":\"isDeregistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"bytes20\"}],\"name\":\"canDeregister\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"bytes20\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"bytes20\"},{\"name\":\"_publicKey\",\"type\":\"bytes\"},{\"name\":\"_bond\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"bytes20\"}],\"name\":\"getDarknodeOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextMinimumEpochInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"bytes20\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumEpochInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"bytes20\"}],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextMinimumBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"bytes20\"}],\"name\":\"getBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextMinimumPodSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"name\":\"epochhash\",\"type\":\"uint256\"},{\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nextMinimumPodSize\",\"type\":\"uint256\"}],\"name\":\"updateMinimumPodSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nextMinimumEpochInterval\",\"type\":\"uint256\"}],\"name\":\"updateEpochInterval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"epoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumPodSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDarknodes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes20[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"bytes20\"}],\"name\":\"isUnregistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_darknodeID\",\"type\":\"bytes20\"}],\"name\":\"deregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_renAddress\",\"type\":\"address\"},{\"name\":\"_minimumBond\",\"type\":\"uint256\"},{\"name\":\"_minimumPodSize\",\"type\":\"uint256\"},{\"name\":\"_minimumEpochInterval\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_darknodeID\",\"type\":\"bytes20\"},{\"indexed\":false,\"name\":\"_bond\",\"type\":\"uint256\"}],\"name\":\"DarknodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_darknodeID\",\"type\":\"bytes20\"}],\"name\":\"DarknodeDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DarknodeOwnerRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousMinimumBond\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nextMinimumBond\",\"type\":\"uint256\"}],\"name\":\"MinimumBondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousMinimumPodSize\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nextMinimumPodSize\",\"type\":\"uint256\"}],\"name\":\"MinimumPodSizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousMinimumEpochInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nextMinimumEpochInterval\",\"type\":\"uint256\"}],\"name\":\"MinimumEpochIntervalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DarknodeRegistryBin is the compiled bytecode used for deploying new contracts.
const DarknodeRegistryBin = `0x608060405234801561001057600080fd5b5060405160808061142283398101604081815282516020808501518386015160609096015160008054600160a060020a03199081163317825560018054600160a060020a0390971696909116959095179094556006919091556007959095556008949094558183019091524360001981014080845293909201829052600c92909255600d556004819055600555611376806100ac6000396000f30060806040526004361061015e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630847e9fa81146101635780630ff9aafe1461018a5780631460e603146101a4578063171f6ea8146101b957806327c4b327146101ef57806332ccd52f14610211578063375a8be3146102a8578063424fe57b14610312578063455dc46d146103505780634f5550fc1461036557806355cacda5146103875780635a8f9b811461039c57806360a22fe4146103be57806368f209eb146103d3578063702c25ee146103f5578063715018a61461040a578063766718081461041f57806380a0c4611461044d57806383907348146104655780638da5cb5b1461047d578063900cf0cf14610492578063aa7517e1146104a7578063c7dbc2be146104bc578063c8a8349b146104d1578063d3841c2514610536578063e08b4c8a14610558578063f2fde38b1461057a575b600080fd5b34801561016f57600080fd5b5061017861059b565b60408051918252519081900360200190f35b34801561019657600080fd5b506101a26004356105a1565b005b3480156101b057600080fd5b506101786105bd565b3480156101c557600080fd5b506101db6001606060020a0319600435166105c3565b604080519115158252519081900360200190f35b3480156101fb57600080fd5b506101db6001606060020a031960043516610612565b34801561021d57600080fd5b506102336001606060020a031960043516610645565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561026d578181015183820152602001610255565b50505050905090810190601f16801561029a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102b457600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101a29583356001606060020a03191695369560449491939091019190819084018382808284375094975050933594506106f29350505050565b34801561031e57600080fd5b506103346001606060020a0319600435166108c1565b60408051600160a060020a039092168252519081900360200190f35b34801561035c57600080fd5b506101786108e6565b34801561037157600080fd5b506101db6001606060020a0319600435166108ec565b34801561039357600080fd5b5061017861094d565b3480156103a857600080fd5b506101a26001606060020a031960043516610953565b3480156103ca57600080fd5b50610178610b4b565b3480156103df57600080fd5b506101786001606060020a031960043516610b51565b34801561040157600080fd5b50610178610b70565b34801561041657600080fd5b506101a2610b76565b34801561042b57600080fd5b50610434610be2565b6040805192835260208301919091528051918290030190f35b34801561045957600080fd5b506101a2600435610beb565b34801561047157600080fd5b506101a2600435610c07565b34801561048957600080fd5b50610334610c23565b34801561049e57600080fd5b506101a2610c32565b3480156104b357600080fd5b50610178610d9e565b3480156104c857600080fd5b50610178610da4565b3480156104dd57600080fd5b506104e6610daa565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561052257818101518382015260200161050a565b505050509050019250505060405180910390f35b34801561054257600080fd5b506101db6001606060020a031960043516610e62565b34801561056457600080fd5b506101a26001606060020a031960043516610e83565b34801561058657600080fd5b506101a2600160a060020a0360043516610f39565b60055481565b600054600160a060020a031633146105b857600080fd5b600955565b60045481565b6001606060020a031981166000908152600260205260408120600301541580159061060c5750600d546001606060020a0319831660009081526002602052604090206003015411155b92915050565b600061061d826108ec565b801561060c5750506001606060020a0319166000908152600260205260409020600301541590565b6001606060020a03198116600090815260026020818152604092839020600401805484516001821615610100026000190190911693909304601f810183900483028401830190945283835260609390918301828280156106e65780601f106106bb576101008083540402835291602001916106e6565b820191906000526020600020905b8154815290600101906020018083116106c957829003601f168201915b50505050509050919050565b826106fc81610e62565b151561070757600080fd5b60065482101561071657600080fd5b600154604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b15801561078957600080fd5b505af115801561079d573d6000803e3d6000fd5b505050506040513d60208110156107b357600080fd5b505115156107c057600080fd5b6040805160a0810182523381526020808201858152600854600d5401838501908152600060608501818152608086018a81526001606060020a03198c1683526002808752979092208651815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178155935160018501559151958301959095555160038201559251805192939261085e92600485019201906112af565b5090505061086d600385610f5c565b600580546001019055604080516001606060020a0319861681526020810184905281517f964f77c16f29e1a7c974315f1fdc990a01866a66c8a3db959112bdfa14cb2d9d929181900390910190a150505050565b6001606060020a031916600090815260026020526040902054600160a060020a031690565b600b5481565b6001606060020a03198116600090815260026020819052604082200154158015906109365750600d546001606060020a031983166000908152600260208190526040909120015411155b801561060c5750610946826105c3565b1592915050565b60085481565b6001606060020a031981166000908152600260205260408120548290600160a060020a0316331461098357600080fd5b8261098d816105c3565b151561099857600080fd5b6001606060020a0319841660009081526002602052604090206001015492506109c2600385610f87565b6040805160a081018252600080825260208083018281528385018381526060850184815286518085018852858152608087019081526001606060020a03198c1686526002808652979095208651815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178155925160018401559051958201959095559351600385015590518051929392610a6492600485019201906112af565b5050600154604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018790529051600160a060020a03909216925063a9059cbb9160448083019260209291908290030181600087803b158015610ad357600080fd5b505af1158015610ae7573d6000803e3d6000fd5b505050506040513d6020811015610afd57600080fd5b50511515610b0a57600080fd5b604080513381526020810185905281517fd52d1526010a1ff0e8591c1d6162705753bfcfafc3d8b243e5f84ce90e48d919929181900390910190a150505050565b60095481565b6001606060020a03191660009081526002602052604090206001015490565b600a5481565b600054600160a060020a03163314610b8d57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600c54600d5482565b600054600160a060020a03163314610c0257600080fd5b600a55565b600054600160a060020a03163314610c1e57600080fd5b600b55565b600054600160a060020a031681565b600854600d5460009101431015610c4857600080fd5b5060408051808201909152436000198101408083526020909201819052600c829055600d5560055460045560065460095414610cc8577fab6324bbd00ba7fc9332772b8e930e40f74d073f902ee642a673877512a5a900600654600954604051808381526020018281526020019250505060405180910390a16009546006555b600754600a5414610d1d577ffec92e92eb77d037492dac6ca1b93ef1ba11d4d30478d05fe718a273fb69fb3d600754600a54604051808381526020018281526020019250505060405180910390a1600a546007555b600854600b5414610d72577f156b9bdf0755726922a37be8541321f59f36e7c8d072428382fad452eb56cb80600854600b54604051808381526020018281526020019250505060405180910390a1600b546008555b6040517fe358419ca0dd7928a310d787a606dfae5d869f5071249efa6107105e7afc40bc90600090a150565b60065481565b60075481565b606080600080600454604051908082528060200260200182016040528015610ddc578160200160208202803883390190505b50925060009150610ded6003611080565b90505b600454821015610e5a57610e03816108ec565b1515610e1b57610e146003826110a5565b9050610df0565b808383815181101515610e2a57fe5b6001606060020a0319909216602092830290910190910152610e4d6003826110a5565b6001909201919050610df0565b509092915050565b6001606060020a031916600090815260026020819052604090912001541590565b80610e8d81610612565b1515610e9857600080fd5b6001606060020a031982166000908152600260205260409020548290600160a060020a03163314610ec857600080fd5b600854600d546001606060020a031985166000818152600260209081526040918290209390940160039093019290925560058054600019019055815190815290517fd261e3f9e22d65cdbecf9c4c79c684a7d4225282f1c80dcbfa6fec5c38a151d4929181900390910190a1505050565b600054600160a060020a03163314610f5057600080fd5b610f59816110ec565b50565b610f668282611169565b15610f7057600080fd5b610f8382610f7d84611189565b836111b1565b5050565b600080610f948484611169565b1515610f9f57600080fd5b6001606060020a031983161515610fb55761107a565b50506001606060020a03198082166000818152602085905260408082208054600180830180546c01000000000000000000000000610100948590048102808b168952878920909401805492820282810473ffffffffffffffffffffffffffffffffffffffff1994851617909155998a168852958720805496840490940274ffffffffffffffffffffffffffffffffffffffff00199096169590951790925594909352805474ffffffffffffffffffffffffffffffffffffffffff191690558154169055905b50505050565b600080805260209190915260409020600101546c010000000000000000000000000290565b60006110b18383611169565b15156110bc57600080fd5b506001606060020a031916600090815260209190915260409020600101546c010000000000000000000000000290565b600160a060020a038116151561110157600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6001606060020a0319166000908152602091909152604090205460ff1690565b60008080526020829052604090205461010090046c0100000000000000000000000002919050565b60006111bd8483611169565b156111c757600080fd5b6111d18484611169565b806111e457506001606060020a03198316155b15156111ef57600080fd5b506001606060020a0319808316600090815260209490945260408085206001908101805485851680895284892080546c01000000000000000000000000998a900461010090810274ffffffffffffffffffffffffffffffffffffffff00199283161783558287018054958c028c810473ffffffffffffffffffffffffffffffffffffffff199788161790915586549b909a049a9094168a179094559690951688529287208054969093029516949094179055909252815460ff1916179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106112f057805160ff191683800117855561131d565b8280016001018555821561131d579182015b8281111561131d578251825591602001919060010190611302565b5061132992915061132d565b5090565b61134791905b808211156113295760008155600101611333565b905600a165627a7a723058209e37691f7c72c57f7ce90e9707aee856a8e62bc67850eaaedcd33602ce009acd0029`

// DeployDarknodeRegistry deploys a new Ethereum contract, binding an instance of DarknodeRegistry to it.
func DeployDarknodeRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _renAddress common.Address, _minimumBond *big.Int, _minimumPodSize *big.Int, _minimumEpochInterval *big.Int) (common.Address, *types.Transaction, *DarknodeRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(DarknodeRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DarknodeRegistryBin), backend, _renAddress, _minimumBond, _minimumPodSize, _minimumEpochInterval)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DarknodeRegistry{DarknodeRegistryCaller: DarknodeRegistryCaller{contract: contract}, DarknodeRegistryTransactor: DarknodeRegistryTransactor{contract: contract}, DarknodeRegistryFilterer: DarknodeRegistryFilterer{contract: contract}}, nil
}

// DarknodeRegistry is an auto generated Go binding around an Ethereum contract.
type DarknodeRegistry struct {
	DarknodeRegistryCaller     // Read-only binding to the contract
	DarknodeRegistryTransactor // Write-only binding to the contract
	DarknodeRegistryFilterer   // Log filterer for contract events
}

// DarknodeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DarknodeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DarknodeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DarknodeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DarknodeRegistrySession struct {
	Contract     *DarknodeRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DarknodeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DarknodeRegistryCallerSession struct {
	Contract *DarknodeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DarknodeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DarknodeRegistryTransactorSession struct {
	Contract     *DarknodeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DarknodeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DarknodeRegistryRaw struct {
	Contract *DarknodeRegistry // Generic contract binding to access the raw methods on
}

// DarknodeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DarknodeRegistryCallerRaw struct {
	Contract *DarknodeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DarknodeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DarknodeRegistryTransactorRaw struct {
	Contract *DarknodeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDarknodeRegistry creates a new instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistry(address common.Address, backend bind.ContractBackend) (*DarknodeRegistry, error) {
	contract, err := bindDarknodeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistry{DarknodeRegistryCaller: DarknodeRegistryCaller{contract: contract}, DarknodeRegistryTransactor: DarknodeRegistryTransactor{contract: contract}, DarknodeRegistryFilterer: DarknodeRegistryFilterer{contract: contract}}, nil
}

// NewDarknodeRegistryCaller creates a new read-only instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistryCaller(address common.Address, caller bind.ContractCaller) (*DarknodeRegistryCaller, error) {
	contract, err := bindDarknodeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryCaller{contract: contract}, nil
}

// NewDarknodeRegistryTransactor creates a new write-only instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DarknodeRegistryTransactor, error) {
	contract, err := bindDarknodeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryTransactor{contract: contract}, nil
}

// NewDarknodeRegistryFilterer creates a new log filterer instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DarknodeRegistryFilterer, error) {
	contract, err := bindDarknodeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryFilterer{contract: contract}, nil
}

// bindDarknodeRegistry binds a generic wrapper to an already deployed contract.
func bindDarknodeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DarknodeRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DarknodeRegistry *DarknodeRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DarknodeRegistry.Contract.DarknodeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DarknodeRegistry *DarknodeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.DarknodeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DarknodeRegistry *DarknodeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.DarknodeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DarknodeRegistry *DarknodeRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DarknodeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DarknodeRegistry *DarknodeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DarknodeRegistry *DarknodeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.contract.Transact(opts, method, params...)
}

// CanDeregister is a free data retrieval call binding the contract method 0x27c4b327.
//
// Solidity: function canDeregister(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) CanDeregister(opts *bind.CallOpts, _darknodeID [20]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "canDeregister", _darknodeID)
	return *ret0, err
}

// CanDeregister is a free data retrieval call binding the contract method 0x27c4b327.
//
// Solidity: function canDeregister(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) CanDeregister(_darknodeID [20]byte) (bool, error) {
	return _DarknodeRegistry.Contract.CanDeregister(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// CanDeregister is a free data retrieval call binding the contract method 0x27c4b327.
//
// Solidity: function canDeregister(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) CanDeregister(_darknodeID [20]byte) (bool, error) {
	return _DarknodeRegistry.Contract.CanDeregister(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(epochhash uint256, blocknumber uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) CurrentEpoch(opts *bind.CallOpts) (struct {
	Epochhash   *big.Int
	Blocknumber *big.Int
}, error) {
	ret := new(struct {
		Epochhash   *big.Int
		Blocknumber *big.Int
	})
	out := ret
	err := _DarknodeRegistry.contract.Call(opts, out, "currentEpoch")
	return *ret, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(epochhash uint256, blocknumber uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) CurrentEpoch() (struct {
	Epochhash   *big.Int
	Blocknumber *big.Int
}, error) {
	return _DarknodeRegistry.Contract.CurrentEpoch(&_DarknodeRegistry.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(epochhash uint256, blocknumber uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) CurrentEpoch() (struct {
	Epochhash   *big.Int
	Blocknumber *big.Int
}, error) {
	return _DarknodeRegistry.Contract.CurrentEpoch(&_DarknodeRegistry.CallOpts)
}

// GetBond is a free data retrieval call binding the contract method 0x68f209eb.
//
// Solidity: function getBond(_darknodeID bytes20) constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) GetBond(opts *bind.CallOpts, _darknodeID [20]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getBond", _darknodeID)
	return *ret0, err
}

// GetBond is a free data retrieval call binding the contract method 0x68f209eb.
//
// Solidity: function getBond(_darknodeID bytes20) constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) GetBond(_darknodeID [20]byte) (*big.Int, error) {
	return _DarknodeRegistry.Contract.GetBond(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetBond is a free data retrieval call binding the contract method 0x68f209eb.
//
// Solidity: function getBond(_darknodeID bytes20) constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetBond(_darknodeID [20]byte) (*big.Int, error) {
	return _DarknodeRegistry.Contract.GetBond(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodeOwner is a free data retrieval call binding the contract method 0x424fe57b.
//
// Solidity: function getDarknodeOwner(_darknodeID bytes20) constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) GetDarknodeOwner(opts *bind.CallOpts, _darknodeID [20]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getDarknodeOwner", _darknodeID)
	return *ret0, err
}

// GetDarknodeOwner is a free data retrieval call binding the contract method 0x424fe57b.
//
// Solidity: function getDarknodeOwner(_darknodeID bytes20) constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) GetDarknodeOwner(_darknodeID [20]byte) (common.Address, error) {
	return _DarknodeRegistry.Contract.GetDarknodeOwner(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodeOwner is a free data retrieval call binding the contract method 0x424fe57b.
//
// Solidity: function getDarknodeOwner(_darknodeID bytes20) constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetDarknodeOwner(_darknodeID [20]byte) (common.Address, error) {
	return _DarknodeRegistry.Contract.GetDarknodeOwner(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodes is a free data retrieval call binding the contract method 0xc8a8349b.
//
// Solidity: function getDarknodes() constant returns(bytes20[])
func (_DarknodeRegistry *DarknodeRegistryCaller) GetDarknodes(opts *bind.CallOpts) ([][20]byte, error) {
	var (
		ret0 = new([][20]byte)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getDarknodes")
	return *ret0, err
}

// GetDarknodes is a free data retrieval call binding the contract method 0xc8a8349b.
//
// Solidity: function getDarknodes() constant returns(bytes20[])
func (_DarknodeRegistry *DarknodeRegistrySession) GetDarknodes() ([][20]byte, error) {
	return _DarknodeRegistry.Contract.GetDarknodes(&_DarknodeRegistry.CallOpts)
}

// GetDarknodes is a free data retrieval call binding the contract method 0xc8a8349b.
//
// Solidity: function getDarknodes() constant returns(bytes20[])
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetDarknodes() ([][20]byte, error) {
	return _DarknodeRegistry.Contract.GetDarknodes(&_DarknodeRegistry.CallOpts)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x32ccd52f.
//
// Solidity: function getPublicKey(_darknodeID bytes20) constant returns(bytes)
func (_DarknodeRegistry *DarknodeRegistryCaller) GetPublicKey(opts *bind.CallOpts, _darknodeID [20]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getPublicKey", _darknodeID)
	return *ret0, err
}

// GetPublicKey is a free data retrieval call binding the contract method 0x32ccd52f.
//
// Solidity: function getPublicKey(_darknodeID bytes20) constant returns(bytes)
func (_DarknodeRegistry *DarknodeRegistrySession) GetPublicKey(_darknodeID [20]byte) ([]byte, error) {
	return _DarknodeRegistry.Contract.GetPublicKey(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x32ccd52f.
//
// Solidity: function getPublicKey(_darknodeID bytes20) constant returns(bytes)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetPublicKey(_darknodeID [20]byte) ([]byte, error) {
	return _DarknodeRegistry.Contract.GetPublicKey(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsDeregistered is a free data retrieval call binding the contract method 0x171f6ea8.
//
// Solidity: function isDeregistered(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsDeregistered(opts *bind.CallOpts, _darknodeID [20]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isDeregistered", _darknodeID)
	return *ret0, err
}

// IsDeregistered is a free data retrieval call binding the contract method 0x171f6ea8.
//
// Solidity: function isDeregistered(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsDeregistered(_darknodeID [20]byte) (bool, error) {
	return _DarknodeRegistry.Contract.IsDeregistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsDeregistered is a free data retrieval call binding the contract method 0x171f6ea8.
//
// Solidity: function isDeregistered(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsDeregistered(_darknodeID [20]byte) (bool, error) {
	return _DarknodeRegistry.Contract.IsDeregistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRegistered is a free data retrieval call binding the contract method 0x4f5550fc.
//
// Solidity: function isRegistered(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsRegistered(opts *bind.CallOpts, _darknodeID [20]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isRegistered", _darknodeID)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0x4f5550fc.
//
// Solidity: function isRegistered(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsRegistered(_darknodeID [20]byte) (bool, error) {
	return _DarknodeRegistry.Contract.IsRegistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRegistered is a free data retrieval call binding the contract method 0x4f5550fc.
//
// Solidity: function isRegistered(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsRegistered(_darknodeID [20]byte) (bool, error) {
	return _DarknodeRegistry.Contract.IsRegistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsUnregistered is a free data retrieval call binding the contract method 0xd3841c25.
//
// Solidity: function isUnregistered(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsUnregistered(opts *bind.CallOpts, _darknodeID [20]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isUnregistered", _darknodeID)
	return *ret0, err
}

// IsUnregistered is a free data retrieval call binding the contract method 0xd3841c25.
//
// Solidity: function isUnregistered(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsUnregistered(_darknodeID [20]byte) (bool, error) {
	return _DarknodeRegistry.Contract.IsUnregistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsUnregistered is a free data retrieval call binding the contract method 0xd3841c25.
//
// Solidity: function isUnregistered(_darknodeID bytes20) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsUnregistered(_darknodeID [20]byte) (bool, error) {
	return _DarknodeRegistry.Contract.IsUnregistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// MinimumBond is a free data retrieval call binding the contract method 0xaa7517e1.
//
// Solidity: function minimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) MinimumBond(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "minimumBond")
	return *ret0, err
}

// MinimumBond is a free data retrieval call binding the contract method 0xaa7517e1.
//
// Solidity: function minimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) MinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumBond(&_DarknodeRegistry.CallOpts)
}

// MinimumBond is a free data retrieval call binding the contract method 0xaa7517e1.
//
// Solidity: function minimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) MinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumBond(&_DarknodeRegistry.CallOpts)
}

// MinimumEpochInterval is a free data retrieval call binding the contract method 0x55cacda5.
//
// Solidity: function minimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) MinimumEpochInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "minimumEpochInterval")
	return *ret0, err
}

// MinimumEpochInterval is a free data retrieval call binding the contract method 0x55cacda5.
//
// Solidity: function minimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) MinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// MinimumEpochInterval is a free data retrieval call binding the contract method 0x55cacda5.
//
// Solidity: function minimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) MinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// MinimumPodSize is a free data retrieval call binding the contract method 0xc7dbc2be.
//
// Solidity: function minimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) MinimumPodSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "minimumPodSize")
	return *ret0, err
}

// MinimumPodSize is a free data retrieval call binding the contract method 0xc7dbc2be.
//
// Solidity: function minimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) MinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// MinimumPodSize is a free data retrieval call binding the contract method 0xc7dbc2be.
//
// Solidity: function minimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) MinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// NextMinimumBond is a free data retrieval call binding the contract method 0x60a22fe4.
//
// Solidity: function nextMinimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NextMinimumBond(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "nextMinimumBond")
	return *ret0, err
}

// NextMinimumBond is a free data retrieval call binding the contract method 0x60a22fe4.
//
// Solidity: function nextMinimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NextMinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumBond(&_DarknodeRegistry.CallOpts)
}

// NextMinimumBond is a free data retrieval call binding the contract method 0x60a22fe4.
//
// Solidity: function nextMinimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NextMinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumBond(&_DarknodeRegistry.CallOpts)
}

// NextMinimumEpochInterval is a free data retrieval call binding the contract method 0x455dc46d.
//
// Solidity: function nextMinimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NextMinimumEpochInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "nextMinimumEpochInterval")
	return *ret0, err
}

// NextMinimumEpochInterval is a free data retrieval call binding the contract method 0x455dc46d.
//
// Solidity: function nextMinimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NextMinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// NextMinimumEpochInterval is a free data retrieval call binding the contract method 0x455dc46d.
//
// Solidity: function nextMinimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NextMinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// NextMinimumPodSize is a free data retrieval call binding the contract method 0x702c25ee.
//
// Solidity: function nextMinimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NextMinimumPodSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "nextMinimumPodSize")
	return *ret0, err
}

// NextMinimumPodSize is a free data retrieval call binding the contract method 0x702c25ee.
//
// Solidity: function nextMinimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NextMinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// NextMinimumPodSize is a free data retrieval call binding the contract method 0x702c25ee.
//
// Solidity: function nextMinimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NextMinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// NumDarknodes is a free data retrieval call binding the contract method 0x1460e603.
//
// Solidity: function numDarknodes() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NumDarknodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "numDarknodes")
	return *ret0, err
}

// NumDarknodes is a free data retrieval call binding the contract method 0x1460e603.
//
// Solidity: function numDarknodes() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NumDarknodes() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodes(&_DarknodeRegistry.CallOpts)
}

// NumDarknodes is a free data retrieval call binding the contract method 0x1460e603.
//
// Solidity: function numDarknodes() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NumDarknodes() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodes(&_DarknodeRegistry.CallOpts)
}

// NumDarknodesNextEpoch is a free data retrieval call binding the contract method 0x0847e9fa.
//
// Solidity: function numDarknodesNextEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NumDarknodesNextEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "numDarknodesNextEpoch")
	return *ret0, err
}

// NumDarknodesNextEpoch is a free data retrieval call binding the contract method 0x0847e9fa.
//
// Solidity: function numDarknodesNextEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NumDarknodesNextEpoch() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodesNextEpoch(&_DarknodeRegistry.CallOpts)
}

// NumDarknodesNextEpoch is a free data retrieval call binding the contract method 0x0847e9fa.
//
// Solidity: function numDarknodesNextEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NumDarknodesNextEpoch() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodesNextEpoch(&_DarknodeRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) Owner() (common.Address, error) {
	return _DarknodeRegistry.Contract.Owner(&_DarknodeRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) Owner() (common.Address, error) {
	return _DarknodeRegistry.Contract.Owner(&_DarknodeRegistry.CallOpts)
}

// Deregister is a paid mutator transaction binding the contract method 0xe08b4c8a.
//
// Solidity: function deregister(_darknodeID bytes20) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Deregister(opts *bind.TransactOpts, _darknodeID [20]byte) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "deregister", _darknodeID)
}

// Deregister is a paid mutator transaction binding the contract method 0xe08b4c8a.
//
// Solidity: function deregister(_darknodeID bytes20) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Deregister(_darknodeID [20]byte) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Deregister(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Deregister is a paid mutator transaction binding the contract method 0xe08b4c8a.
//
// Solidity: function deregister(_darknodeID bytes20) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Deregister(_darknodeID [20]byte) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Deregister(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Epoch is a paid mutator transaction binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Epoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "epoch")
}

// Epoch is a paid mutator transaction binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Epoch() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Epoch(&_DarknodeRegistry.TransactOpts)
}

// Epoch is a paid mutator transaction binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Epoch() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Epoch(&_DarknodeRegistry.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x5a8f9b81.
//
// Solidity: function refund(_darknodeID bytes20) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Refund(opts *bind.TransactOpts, _darknodeID [20]byte) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "refund", _darknodeID)
}

// Refund is a paid mutator transaction binding the contract method 0x5a8f9b81.
//
// Solidity: function refund(_darknodeID bytes20) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Refund(_darknodeID [20]byte) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Refund(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Refund is a paid mutator transaction binding the contract method 0x5a8f9b81.
//
// Solidity: function refund(_darknodeID bytes20) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Refund(_darknodeID [20]byte) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Refund(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Register is a paid mutator transaction binding the contract method 0x375a8be3.
//
// Solidity: function register(_darknodeID bytes20, _publicKey bytes, _bond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Register(opts *bind.TransactOpts, _darknodeID [20]byte, _publicKey []byte, _bond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "register", _darknodeID, _publicKey, _bond)
}

// Register is a paid mutator transaction binding the contract method 0x375a8be3.
//
// Solidity: function register(_darknodeID bytes20, _publicKey bytes, _bond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Register(_darknodeID [20]byte, _publicKey []byte, _bond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Register(&_DarknodeRegistry.TransactOpts, _darknodeID, _publicKey, _bond)
}

// Register is a paid mutator transaction binding the contract method 0x375a8be3.
//
// Solidity: function register(_darknodeID bytes20, _publicKey bytes, _bond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Register(_darknodeID [20]byte, _publicKey []byte, _bond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Register(&_DarknodeRegistry.TransactOpts, _darknodeID, _publicKey, _bond)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.RenounceOwnership(&_DarknodeRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.RenounceOwnership(&_DarknodeRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.TransferOwnership(&_DarknodeRegistry.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.TransferOwnership(&_DarknodeRegistry.TransactOpts, _newOwner)
}

// UpdateEpochInterval is a paid mutator transaction binding the contract method 0x83907348.
//
// Solidity: function updateEpochInterval(_nextMinimumEpochInterval uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateEpochInterval(opts *bind.TransactOpts, _nextMinimumEpochInterval *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateEpochInterval", _nextMinimumEpochInterval)
}

// UpdateEpochInterval is a paid mutator transaction binding the contract method 0x83907348.
//
// Solidity: function updateEpochInterval(_nextMinimumEpochInterval uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateEpochInterval(_nextMinimumEpochInterval *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateEpochInterval(&_DarknodeRegistry.TransactOpts, _nextMinimumEpochInterval)
}

// UpdateEpochInterval is a paid mutator transaction binding the contract method 0x83907348.
//
// Solidity: function updateEpochInterval(_nextMinimumEpochInterval uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateEpochInterval(_nextMinimumEpochInterval *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateEpochInterval(&_DarknodeRegistry.TransactOpts, _nextMinimumEpochInterval)
}

// UpdateMinimumBond is a paid mutator transaction binding the contract method 0x0ff9aafe.
//
// Solidity: function updateMinimumBond(_nextMinimumBond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateMinimumBond(opts *bind.TransactOpts, _nextMinimumBond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateMinimumBond", _nextMinimumBond)
}

// UpdateMinimumBond is a paid mutator transaction binding the contract method 0x0ff9aafe.
//
// Solidity: function updateMinimumBond(_nextMinimumBond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateMinimumBond(_nextMinimumBond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumBond(&_DarknodeRegistry.TransactOpts, _nextMinimumBond)
}

// UpdateMinimumBond is a paid mutator transaction binding the contract method 0x0ff9aafe.
//
// Solidity: function updateMinimumBond(_nextMinimumBond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateMinimumBond(_nextMinimumBond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumBond(&_DarknodeRegistry.TransactOpts, _nextMinimumBond)
}

// UpdateMinimumPodSize is a paid mutator transaction binding the contract method 0x80a0c461.
//
// Solidity: function updateMinimumPodSize(_nextMinimumPodSize uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateMinimumPodSize(opts *bind.TransactOpts, _nextMinimumPodSize *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateMinimumPodSize", _nextMinimumPodSize)
}

// UpdateMinimumPodSize is a paid mutator transaction binding the contract method 0x80a0c461.
//
// Solidity: function updateMinimumPodSize(_nextMinimumPodSize uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateMinimumPodSize(_nextMinimumPodSize *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumPodSize(&_DarknodeRegistry.TransactOpts, _nextMinimumPodSize)
}

// UpdateMinimumPodSize is a paid mutator transaction binding the contract method 0x80a0c461.
//
// Solidity: function updateMinimumPodSize(_nextMinimumPodSize uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateMinimumPodSize(_nextMinimumPodSize *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumPodSize(&_DarknodeRegistry.TransactOpts, _nextMinimumPodSize)
}

// DarknodeRegistryDarknodeDeregisteredIterator is returned from FilterDarknodeDeregistered and is used to iterate over the raw logs and unpacked data for DarknodeDeregistered events raised by the DarknodeRegistry contract.
type DarknodeRegistryDarknodeDeregisteredIterator struct {
	Event *DarknodeRegistryDarknodeDeregistered // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryDarknodeDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryDarknodeDeregistered)
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
		it.Event = new(DarknodeRegistryDarknodeDeregistered)
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
func (it *DarknodeRegistryDarknodeDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryDarknodeDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryDarknodeDeregistered represents a DarknodeDeregistered event raised by the DarknodeRegistry contract.
type DarknodeRegistryDarknodeDeregistered struct {
	DarknodeID [20]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDarknodeDeregistered is a free log retrieval operation binding the contract event 0xd261e3f9e22d65cdbecf9c4c79c684a7d4225282f1c80dcbfa6fec5c38a151d4.
//
// Solidity: e DarknodeDeregistered(_darknodeID bytes20)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterDarknodeDeregistered(opts *bind.FilterOpts) (*DarknodeRegistryDarknodeDeregisteredIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "DarknodeDeregistered")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryDarknodeDeregisteredIterator{contract: _DarknodeRegistry.contract, event: "DarknodeDeregistered", logs: logs, sub: sub}, nil
}

// WatchDarknodeDeregistered is a free log subscription operation binding the contract event 0xd261e3f9e22d65cdbecf9c4c79c684a7d4225282f1c80dcbfa6fec5c38a151d4.
//
// Solidity: e DarknodeDeregistered(_darknodeID bytes20)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchDarknodeDeregistered(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryDarknodeDeregistered) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "DarknodeDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryDarknodeDeregistered)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "DarknodeDeregistered", log); err != nil {
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

// DarknodeRegistryDarknodeOwnerRefundedIterator is returned from FilterDarknodeOwnerRefunded and is used to iterate over the raw logs and unpacked data for DarknodeOwnerRefunded events raised by the DarknodeRegistry contract.
type DarknodeRegistryDarknodeOwnerRefundedIterator struct {
	Event *DarknodeRegistryDarknodeOwnerRefunded // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryDarknodeOwnerRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryDarknodeOwnerRefunded)
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
		it.Event = new(DarknodeRegistryDarknodeOwnerRefunded)
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
func (it *DarknodeRegistryDarknodeOwnerRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryDarknodeOwnerRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryDarknodeOwnerRefunded represents a DarknodeOwnerRefunded event raised by the DarknodeRegistry contract.
type DarknodeRegistryDarknodeOwnerRefunded struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDarknodeOwnerRefunded is a free log retrieval operation binding the contract event 0xd52d1526010a1ff0e8591c1d6162705753bfcfafc3d8b243e5f84ce90e48d919.
//
// Solidity: e DarknodeOwnerRefunded(_owner address, _amount uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterDarknodeOwnerRefunded(opts *bind.FilterOpts) (*DarknodeRegistryDarknodeOwnerRefundedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "DarknodeOwnerRefunded")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryDarknodeOwnerRefundedIterator{contract: _DarknodeRegistry.contract, event: "DarknodeOwnerRefunded", logs: logs, sub: sub}, nil
}

// WatchDarknodeOwnerRefunded is a free log subscription operation binding the contract event 0xd52d1526010a1ff0e8591c1d6162705753bfcfafc3d8b243e5f84ce90e48d919.
//
// Solidity: e DarknodeOwnerRefunded(_owner address, _amount uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchDarknodeOwnerRefunded(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryDarknodeOwnerRefunded) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "DarknodeOwnerRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryDarknodeOwnerRefunded)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "DarknodeOwnerRefunded", log); err != nil {
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

// DarknodeRegistryDarknodeRegisteredIterator is returned from FilterDarknodeRegistered and is used to iterate over the raw logs and unpacked data for DarknodeRegistered events raised by the DarknodeRegistry contract.
type DarknodeRegistryDarknodeRegisteredIterator struct {
	Event *DarknodeRegistryDarknodeRegistered // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryDarknodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryDarknodeRegistered)
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
		it.Event = new(DarknodeRegistryDarknodeRegistered)
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
func (it *DarknodeRegistryDarknodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryDarknodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryDarknodeRegistered represents a DarknodeRegistered event raised by the DarknodeRegistry contract.
type DarknodeRegistryDarknodeRegistered struct {
	DarknodeID [20]byte
	Bond       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDarknodeRegistered is a free log retrieval operation binding the contract event 0x964f77c16f29e1a7c974315f1fdc990a01866a66c8a3db959112bdfa14cb2d9d.
//
// Solidity: e DarknodeRegistered(_darknodeID bytes20, _bond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterDarknodeRegistered(opts *bind.FilterOpts) (*DarknodeRegistryDarknodeRegisteredIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "DarknodeRegistered")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryDarknodeRegisteredIterator{contract: _DarknodeRegistry.contract, event: "DarknodeRegistered", logs: logs, sub: sub}, nil
}

// WatchDarknodeRegistered is a free log subscription operation binding the contract event 0x964f77c16f29e1a7c974315f1fdc990a01866a66c8a3db959112bdfa14cb2d9d.
//
// Solidity: e DarknodeRegistered(_darknodeID bytes20, _bond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchDarknodeRegistered(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryDarknodeRegistered) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "DarknodeRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryDarknodeRegistered)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "DarknodeRegistered", log); err != nil {
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

// DarknodeRegistryMinimumBondUpdatedIterator is returned from FilterMinimumBondUpdated and is used to iterate over the raw logs and unpacked data for MinimumBondUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryMinimumBondUpdatedIterator struct {
	Event *DarknodeRegistryMinimumBondUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryMinimumBondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryMinimumBondUpdated)
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
		it.Event = new(DarknodeRegistryMinimumBondUpdated)
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
func (it *DarknodeRegistryMinimumBondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryMinimumBondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryMinimumBondUpdated represents a MinimumBondUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryMinimumBondUpdated struct {
	PreviousMinimumBond *big.Int
	NextMinimumBond     *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMinimumBondUpdated is a free log retrieval operation binding the contract event 0xab6324bbd00ba7fc9332772b8e930e40f74d073f902ee642a673877512a5a900.
//
// Solidity: e MinimumBondUpdated(previousMinimumBond uint256, nextMinimumBond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterMinimumBondUpdated(opts *bind.FilterOpts) (*DarknodeRegistryMinimumBondUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "MinimumBondUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryMinimumBondUpdatedIterator{contract: _DarknodeRegistry.contract, event: "MinimumBondUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumBondUpdated is a free log subscription operation binding the contract event 0xab6324bbd00ba7fc9332772b8e930e40f74d073f902ee642a673877512a5a900.
//
// Solidity: e MinimumBondUpdated(previousMinimumBond uint256, nextMinimumBond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchMinimumBondUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryMinimumBondUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "MinimumBondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryMinimumBondUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "MinimumBondUpdated", log); err != nil {
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

// DarknodeRegistryMinimumEpochIntervalUpdatedIterator is returned from FilterMinimumEpochIntervalUpdated and is used to iterate over the raw logs and unpacked data for MinimumEpochIntervalUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryMinimumEpochIntervalUpdatedIterator struct {
	Event *DarknodeRegistryMinimumEpochIntervalUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryMinimumEpochIntervalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryMinimumEpochIntervalUpdated)
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
		it.Event = new(DarknodeRegistryMinimumEpochIntervalUpdated)
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
func (it *DarknodeRegistryMinimumEpochIntervalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryMinimumEpochIntervalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryMinimumEpochIntervalUpdated represents a MinimumEpochIntervalUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryMinimumEpochIntervalUpdated struct {
	PreviousMinimumEpochInterval *big.Int
	NextMinimumEpochInterval     *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterMinimumEpochIntervalUpdated is a free log retrieval operation binding the contract event 0x156b9bdf0755726922a37be8541321f59f36e7c8d072428382fad452eb56cb80.
//
// Solidity: e MinimumEpochIntervalUpdated(previousMinimumEpochInterval uint256, nextMinimumEpochInterval uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterMinimumEpochIntervalUpdated(opts *bind.FilterOpts) (*DarknodeRegistryMinimumEpochIntervalUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "MinimumEpochIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryMinimumEpochIntervalUpdatedIterator{contract: _DarknodeRegistry.contract, event: "MinimumEpochIntervalUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumEpochIntervalUpdated is a free log subscription operation binding the contract event 0x156b9bdf0755726922a37be8541321f59f36e7c8d072428382fad452eb56cb80.
//
// Solidity: e MinimumEpochIntervalUpdated(previousMinimumEpochInterval uint256, nextMinimumEpochInterval uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchMinimumEpochIntervalUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryMinimumEpochIntervalUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "MinimumEpochIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryMinimumEpochIntervalUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "MinimumEpochIntervalUpdated", log); err != nil {
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

// DarknodeRegistryMinimumPodSizeUpdatedIterator is returned from FilterMinimumPodSizeUpdated and is used to iterate over the raw logs and unpacked data for MinimumPodSizeUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryMinimumPodSizeUpdatedIterator struct {
	Event *DarknodeRegistryMinimumPodSizeUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryMinimumPodSizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryMinimumPodSizeUpdated)
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
		it.Event = new(DarknodeRegistryMinimumPodSizeUpdated)
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
func (it *DarknodeRegistryMinimumPodSizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryMinimumPodSizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryMinimumPodSizeUpdated represents a MinimumPodSizeUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryMinimumPodSizeUpdated struct {
	PreviousMinimumPodSize *big.Int
	NextMinimumPodSize     *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterMinimumPodSizeUpdated is a free log retrieval operation binding the contract event 0xfec92e92eb77d037492dac6ca1b93ef1ba11d4d30478d05fe718a273fb69fb3d.
//
// Solidity: e MinimumPodSizeUpdated(previousMinimumPodSize uint256, nextMinimumPodSize uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterMinimumPodSizeUpdated(opts *bind.FilterOpts) (*DarknodeRegistryMinimumPodSizeUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "MinimumPodSizeUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryMinimumPodSizeUpdatedIterator{contract: _DarknodeRegistry.contract, event: "MinimumPodSizeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumPodSizeUpdated is a free log subscription operation binding the contract event 0xfec92e92eb77d037492dac6ca1b93ef1ba11d4d30478d05fe718a273fb69fb3d.
//
// Solidity: e MinimumPodSizeUpdated(previousMinimumPodSize uint256, nextMinimumPodSize uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchMinimumPodSizeUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryMinimumPodSizeUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "MinimumPodSizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryMinimumPodSizeUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "MinimumPodSizeUpdated", log); err != nil {
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

// DarknodeRegistryNewEpochIterator is returned from FilterNewEpoch and is used to iterate over the raw logs and unpacked data for NewEpoch events raised by the DarknodeRegistry contract.
type DarknodeRegistryNewEpochIterator struct {
	Event *DarknodeRegistryNewEpoch // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryNewEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryNewEpoch)
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
		it.Event = new(DarknodeRegistryNewEpoch)
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
func (it *DarknodeRegistryNewEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryNewEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryNewEpoch represents a NewEpoch event raised by the DarknodeRegistry contract.
type DarknodeRegistryNewEpoch struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewEpoch is a free log retrieval operation binding the contract event 0xe358419ca0dd7928a310d787a606dfae5d869f5071249efa6107105e7afc40bc.
//
// Solidity: e NewEpoch()
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterNewEpoch(opts *bind.FilterOpts) (*DarknodeRegistryNewEpochIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "NewEpoch")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryNewEpochIterator{contract: _DarknodeRegistry.contract, event: "NewEpoch", logs: logs, sub: sub}, nil
}

// WatchNewEpoch is a free log subscription operation binding the contract event 0xe358419ca0dd7928a310d787a606dfae5d869f5071249efa6107105e7afc40bc.
//
// Solidity: e NewEpoch()
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchNewEpoch(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryNewEpoch) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "NewEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryNewEpoch)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "NewEpoch", log); err != nil {
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

// DarknodeRegistryOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the DarknodeRegistry contract.
type DarknodeRegistryOwnershipRenouncedIterator struct {
	Event *DarknodeRegistryOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryOwnershipRenounced)
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
		it.Event = new(DarknodeRegistryOwnershipRenounced)
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
func (it *DarknodeRegistryOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryOwnershipRenounced represents a OwnershipRenounced event raised by the DarknodeRegistry contract.
type DarknodeRegistryOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DarknodeRegistryOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryOwnershipRenouncedIterator{contract: _DarknodeRegistry.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryOwnershipRenounced)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// DarknodeRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DarknodeRegistry contract.
type DarknodeRegistryOwnershipTransferredIterator struct {
	Event *DarknodeRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryOwnershipTransferred)
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
		it.Event = new(DarknodeRegistryOwnershipTransferred)
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
func (it *DarknodeRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the DarknodeRegistry contract.
type DarknodeRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DarknodeRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryOwnershipTransferredIterator{contract: _DarknodeRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryOwnershipTransferred)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ECDSAABI is the input ABI used to generate the binding from.
const ECDSAABI = "[]"

// ECDSABin is the compiled bytecode used for deploying new contracts.
const ECDSABin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058204e4f01a9cef363cf253fc7073304ee5df1bb98523fbed68d292a2a443959c7050029`

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// LinkedListABI is the input ABI used to generate the binding from.
const LinkedListABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes20\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LinkedListBin is the compiled bytecode used for deploying new contracts.
const LinkedListBin = `0x60b361002f600b82828239805160001a6073146000811461001f57610021565bfe5b5030600052607381538281f300730000000000000000000000000000000000000000301460806040526004361060555763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663f26be3fc8114605a575b600080fd5b60606082565b604080516bffffffffffffffffffffffff199092168252519081900360200190f35b6000815600a165627a7a72305820885c1238232f79aaff9a0050465a2649a5bdcd102a04a71ba8514158da6d11f00029`

// DeployLinkedList deploys a new Ethereum contract, binding an instance of LinkedList to it.
func DeployLinkedList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LinkedList, error) {
	parsed, err := abi.JSON(strings.NewReader(LinkedListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LinkedListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LinkedList{LinkedListCaller: LinkedListCaller{contract: contract}, LinkedListTransactor: LinkedListTransactor{contract: contract}, LinkedListFilterer: LinkedListFilterer{contract: contract}}, nil
}

// LinkedList is an auto generated Go binding around an Ethereum contract.
type LinkedList struct {
	LinkedListCaller     // Read-only binding to the contract
	LinkedListTransactor // Write-only binding to the contract
	LinkedListFilterer   // Log filterer for contract events
}

// LinkedListCaller is an auto generated read-only Go binding around an Ethereum contract.
type LinkedListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkedListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LinkedListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkedListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LinkedListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkedListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LinkedListSession struct {
	Contract     *LinkedList       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LinkedListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LinkedListCallerSession struct {
	Contract *LinkedListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LinkedListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LinkedListTransactorSession struct {
	Contract     *LinkedListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LinkedListRaw is an auto generated low-level Go binding around an Ethereum contract.
type LinkedListRaw struct {
	Contract *LinkedList // Generic contract binding to access the raw methods on
}

// LinkedListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LinkedListCallerRaw struct {
	Contract *LinkedListCaller // Generic read-only contract binding to access the raw methods on
}

// LinkedListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LinkedListTransactorRaw struct {
	Contract *LinkedListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLinkedList creates a new instance of LinkedList, bound to a specific deployed contract.
func NewLinkedList(address common.Address, backend bind.ContractBackend) (*LinkedList, error) {
	contract, err := bindLinkedList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LinkedList{LinkedListCaller: LinkedListCaller{contract: contract}, LinkedListTransactor: LinkedListTransactor{contract: contract}, LinkedListFilterer: LinkedListFilterer{contract: contract}}, nil
}

// NewLinkedListCaller creates a new read-only instance of LinkedList, bound to a specific deployed contract.
func NewLinkedListCaller(address common.Address, caller bind.ContractCaller) (*LinkedListCaller, error) {
	contract, err := bindLinkedList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LinkedListCaller{contract: contract}, nil
}

// NewLinkedListTransactor creates a new write-only instance of LinkedList, bound to a specific deployed contract.
func NewLinkedListTransactor(address common.Address, transactor bind.ContractTransactor) (*LinkedListTransactor, error) {
	contract, err := bindLinkedList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LinkedListTransactor{contract: contract}, nil
}

// NewLinkedListFilterer creates a new log filterer instance of LinkedList, bound to a specific deployed contract.
func NewLinkedListFilterer(address common.Address, filterer bind.ContractFilterer) (*LinkedListFilterer, error) {
	contract, err := bindLinkedList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LinkedListFilterer{contract: contract}, nil
}

// bindLinkedList binds a generic wrapper to an already deployed contract.
func bindLinkedList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LinkedListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinkedList *LinkedListRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LinkedList.Contract.LinkedListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinkedList *LinkedListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkedList.Contract.LinkedListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinkedList *LinkedListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkedList.Contract.LinkedListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinkedList *LinkedListCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LinkedList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinkedList *LinkedListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkedList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinkedList *LinkedListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkedList.Contract.contract.Transact(opts, method, params...)
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() constant returns(bytes20)
func (_LinkedList *LinkedListCaller) NULL(opts *bind.CallOpts) ([20]byte, error) {
	var (
		ret0 = new([20]byte)
	)
	out := ret0
	err := _LinkedList.contract.Call(opts, out, "NULL")
	return *ret0, err
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() constant returns(bytes20)
func (_LinkedList *LinkedListSession) NULL() ([20]byte, error) {
	return _LinkedList.Contract.NULL(&_LinkedList.CallOpts)
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() constant returns(bytes20)
func (_LinkedList *LinkedListCallerSession) NULL() ([20]byte, error) {
	return _LinkedList.Contract.NULL(&_LinkedList.CallOpts)
}

// OrderbookABI is the input ABI used to generate the binding from.
const OrderbookABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"orderConfirmer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"buyOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"buyOrders\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"openSellOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sellOrders\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"openBuyOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"name\":\"_orderMatches\",\"type\":\"bytes32[]\"}],\"name\":\"confirmOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"orderBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ren\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_offset\",\"type\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getOrders\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint8[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"sellOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"darknodeRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"orderDepth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"orderBroker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"orderState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDarknodeRegistry\",\"type\":\"address\"}],\"name\":\"updateDarknodeRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"orderMatch\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"orderTrader\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"orderPriority\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrdersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_renAddress\",\"type\":\"address\"},{\"name\":\"_darknodeRegistry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nextFee\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousDarknodeRegistry\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextDarknodeRegistry\",\"type\":\"address\"}],\"name\":\"DarknodeRegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OrderbookBin is the compiled bytecode used for deploying new contracts.
const OrderbookBin = `0x608060405234801561001057600080fd5b5060405160608061163c83398101604090815281516020830151919092015160008054600160a060020a0319908116331790915560059390935560068054600160a060020a039384169085161790556007805492909116919092161790556115bf8061007d6000396000f30060806040526004361061015e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631107c3f7811461016357806322f85eaa1461019757806335cea288146101c857806339b0d677146101f25780634a8393f31461024f5780635060340b14610267578063574ed6c1146102c25780637008b9961461031d578063715018a61461037757806389895d531461038c5780638a9b4067146103a45780638da5cb5b146103b95780638f72fc77146103ce5780639012c4a8146104c757806397514d90146104df5780639e45e0d0146104f7578063a188fcb81461050c578063a518166114610524578063aab14d041461053c578063aaff096d1461056a578063af3e8a401461058b578063b1a08010146105f3578063b248e4e11461060b578063b5b3b05114610623578063d09ef24114610638578063ddca3f4314610650578063f2fde38b14610665575b600080fd5b34801561016f57600080fd5b5061017b600435610686565b60408051600160a060020a039092168252519081900360200190f35b3480156101a357600080fd5b506101af6004356106a7565b6040805192835290151560208301528051918290030190f35b3480156101d457600080fd5b506101e06004356106e5565b60408051918252519081900360200190f35b3480156101fe57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261024d94369492936024939284019190819084018382808284375094975050933594506107049350505050565b005b34801561025b57600080fd5b506101e0600435610754565b34801561027357600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261024d94369492936024939284019190819084018382808284375094975050933594506107629350505050565b3480156102ce57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261024d94369492936024939284019190819084018382808284375094975050933594506107b19350505050565b34801561032957600080fd5b5060408051602060046024803582810135848102808701860190975280865261024d968435963696604495919490910192918291850190849080828437509497506109079650505050505050565b34801561038357600080fd5b5061024d610b93565b34801561039857600080fd5b506101e0600435610bff565b3480156103b057600080fd5b5061017b610c15565b3480156103c557600080fd5b5061017b610c24565b3480156103da57600080fd5b506103e9600435602435610c33565b60405180806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b83811015610431578181015183820152602001610419565b50505050905001848103835286818151815260200191508051906020019060200280838360005b83811015610470578181015183820152602001610458565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156104af578181015183820152602001610497565b50505050905001965050505050505060405180910390f35b3480156104d357600080fd5b5061024d600435610ddd565b3480156104eb57600080fd5b506101af600435610e36565b34801561050357600080fd5b5061017b610e5d565b34801561051857600080fd5b506101e0600435610e6c565b34801561053057600080fd5b5061017b600435610ea5565b34801561054857600080fd5b50610554600435610ec3565b6040805160ff9092168252519081900360200190f35b34801561057657600080fd5b5061024d600160a060020a0360043516610ee6565b34801561059757600080fd5b506105a3600435610f74565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156105df5781810151838201526020016105c7565b505050509050019250505060405180910390f35b3480156105ff57600080fd5b5061017b600435610fda565b34801561061757600080fd5b506101e0600435610ffa565b34801561062f57600080fd5b506101e061100f565b34801561064457600080fd5b506101af60043561101a565b34801561065c57600080fd5b506101e0611041565b34801561067157600080fd5b5061024d600160a060020a0360043516611047565b600081815260046020526040902060020154600160a060020a03165b919050565b600154600090819083106106c0575060009050806106e0565b60018054849081106106ce57fe5b90600052602060002001546001915091505b915091565b60018054829081106106f357fe5b600091825260209091200154905081565b61070e828261106a565b600280546001810182557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01829055546000918252600460205260409091206003015550565b60028054829081106106f357fe5b61076c828261106a565b6001805480820182557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601829055546000918252600460205260409091206003015550565b600080600160008481526004602052604090205460ff1660038111156107d357fe5b14156108bb57604080517f52657075626c69632050726f746f636f6c3a2063616e63656c3a200000000000602080830191909152603b80830187905283518084039091018152605b90920192839052815191929182918401908083835b6020831061084f5780518252601f199092019160209182019101610830565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902091506108888285611292565b600084815260046020526040902054909150600160a060020a0380831661010090920416146108b657600080fd5b6108e2565b60008381526004602052604081205460ff1660038111156108d857fe5b146108e257600080fd5b50506000908152600460208190526040909120805460ff191660031781554391015550565b600754604080517f4f5550fc000000000000000000000000000000000000000000000000000000008152336c0100000000000000000000000081026bffffffffffffffffffffffff191660048301529151600093600160a060020a031691634f5550fc91602480830192602092919082900301818887803b15801561098b57600080fd5b505af115801561099f573d6000803e3d6000fd5b505050506040513d60208110156109b557600080fd5b505115156109c257600080fd5b600160008581526004602052604090205460ff1660038111156109e157fe5b146109eb57600080fd5b600091505b8251821015610a4a576001600460008585815181101515610a0d57fe5b602090810290910181015182528101919091526040016000205460ff166003811115610a3557fe5b14610a3f57600080fd5b6001909101906109f0565b600091505b8251821015610b26576002600460008585815181101515610a6c57fe5b60209081029091018101518252810191909152604001600020805460ff19166001836003811115610a9957fe5b021790555060408051602081019091528481528351600490600090869086908110610ac057fe5b60209081029091018101518252810191909152604001600020610aea91600590910190600161152c565b5043600460008585815181101515610afe57fe5b6020908102909101810151825281019190915260400160002060040155600190910190610a4f565b60008481526004602090815260409091208054600260ff19909116811782558101805473ffffffffffffffffffffffffffffffffffffffff1916331790558451610b789260059092019186019061152c565b50505060009182525060046020819052604090912043910155565b600054600160a060020a03163314610baa57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6000908152600460208190526040909120015490565b600654600160a060020a031681565b600054600160a060020a031681565b6060806060600060608060606000806003805490508b101515610c5557610dd0565b6003548a96508b87011115610c6d576003548b900395505b85604051908082528060200260200182016040528015610c97578160200160208202803883390190505b50945085604051908082528060200260200182016040528015610cc4578160200160208202803883390190505b50935085604051908082528060200260200182016040528015610cf1578160200160208202803883390190505b509250600091505b85821015610dc65760038054838d01908110610d1157fe5b90600052602060002001549050808583815181101515610d2d57fe5b6020908102919091018101919091526000828152600490915260409020548451610100909104600160a060020a031690859084908110610d6957fe5b600160a060020a03909216602092830290910182015260008281526004909152604090205460ff166003811115610d9c57fe5b8383815181101515610daa57fe5b60ff909216602092830290910190910152600190910190610cf9565b8484849850985098505b5050505050509250925092565b600054600160a060020a03163314610df457600080fd5b600554604080519182526020820183905280517f528d9479e9f9889a87a3c30c7f7ba537e5e59c4c85a37733b16e57c62df613029281900390910190a1600555565b60025460009081908310610e4f575060009050806106e0565b60028054849081106106ce57fe5b600754600160a060020a031681565b6000818152600460208190526040822001541515610e8c575060006106a2565b5060009081526004602081905260409091200154430390565b600090815260046020526040902060010154600160a060020a031690565b60008181526004602052604081205460ff166003811115610ee057fe5b92915050565b600054600160a060020a03163314610efd57600080fd5b60075460408051600160a060020a039283168152918316602083015280517fd4443dde7b09a100132dba56b0175378b6ad81b64398aef77626357a1d8daf029281900390910190a16007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600081815260046020908152604091829020600501805483518184028101840190945280845260609392830182828015610fce57602002820191906000526020600020905b81548152600190910190602001808311610fb9575b50505050509050919050565b6000908152600460205260409020546101009004600160a060020a031690565b60009081526004602052604090206003015490565b600254600154015b90565b60035460009081908310611033575060009050806106e0565b60038054849081106106ce57fe5b60055481565b600054600160a060020a0316331461105e57600080fd5b61106781611452565b50565b600654600554604080517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810192909252516000928392600160a060020a03909116916323b872dd9160648082019260209290919082900301818787803b1580156110e457600080fd5b505af11580156110f8573d6000803e3d6000fd5b505050506040513d602081101561110e57600080fd5b5051151561111b57600080fd5b60008381526004602052604081205460ff16600381111561113857fe5b1461114257600080fd5b604080517f52657075626c69632050726f746f636f6c3a206f70656e3a2000000000000000602080830191909152603980830187905283518084039091018152605990920192839052815191929182918401908083835b602083106111b85780518252601f199092019160209182019101611199565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902091506111f18285611292565b600084815260046020819052604082208054600160a060020a039490941661010002600160ff19909516851774ffffffffffffffffffffffffffffffffffffffff001916178155838101805473ffffffffffffffffffffffffffffffffffffffff1916331790554391015560038054928301815590527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0192909255505050565b604080518082018252601c8082527f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080840191825293516000948593849386938a9301918291908083835b602083106112fe5780518252601f1990920191602091820191016112df565b51815160209384036101000a600019018019909216911617905292019384525060408051808503815293820190819052835193945092839250908401908083835b6020831061135e5780518252601f19909201916020918201910161133f565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020915084604081518110151561139c57fe5b016020015160f860020a90819004810204905060ff811615806113c257508060ff166001145b156113cb57601b015b600182826113da8860006114cf565b6113e58960206114cf565b60408051600080825260208083018085529790975260ff90951681830152606081019390935260808301919091525160a08083019493601f198301938390039091019190865af115801561143d573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b600160a060020a038116151561146757600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600080805b60208110156115245780600802858286601f01038151811015156114f457fe5b90602001015160f860020a900460f860020a0260f860020a90049060020a028201915080806001019150506114d4565b509392505050565b828054828255906000526020600020908101928215611569579160200282015b82811115611569578251825560209092019160019091019061154c565b50611575929150611579565b5090565b61101791905b80821115611575576000815560010161157f5600a165627a7a72305820f7b24d08eb7a76f33abb8feba851d1f93291a94d20b0b9a48fee0fa7ce138fca0029`

// DeployOrderbook deploys a new Ethereum contract, binding an instance of Orderbook to it.
func DeployOrderbook(auth *bind.TransactOpts, backend bind.ContractBackend, _fee *big.Int, _renAddress common.Address, _darknodeRegistry common.Address) (common.Address, *types.Transaction, *Orderbook, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OrderbookBin), backend, _fee, _renAddress, _darknodeRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// Orderbook is an auto generated Go binding around an Ethereum contract.
type Orderbook struct {
	OrderbookCaller     // Read-only binding to the contract
	OrderbookTransactor // Write-only binding to the contract
	OrderbookFilterer   // Log filterer for contract events
}

// OrderbookCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderbookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderbookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderbookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderbookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderbookSession struct {
	Contract     *Orderbook        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderbookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderbookCallerSession struct {
	Contract *OrderbookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrderbookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderbookTransactorSession struct {
	Contract     *OrderbookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderbookRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderbookRaw struct {
	Contract *Orderbook // Generic contract binding to access the raw methods on
}

// OrderbookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderbookCallerRaw struct {
	Contract *OrderbookCaller // Generic read-only contract binding to access the raw methods on
}

// OrderbookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderbookTransactorRaw struct {
	Contract *OrderbookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderbook creates a new instance of Orderbook, bound to a specific deployed contract.
func NewOrderbook(address common.Address, backend bind.ContractBackend) (*Orderbook, error) {
	contract, err := bindOrderbook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Orderbook{OrderbookCaller: OrderbookCaller{contract: contract}, OrderbookTransactor: OrderbookTransactor{contract: contract}, OrderbookFilterer: OrderbookFilterer{contract: contract}}, nil
}

// NewOrderbookCaller creates a new read-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookCaller(address common.Address, caller bind.ContractCaller) (*OrderbookCaller, error) {
	contract, err := bindOrderbook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookCaller{contract: contract}, nil
}

// NewOrderbookTransactor creates a new write-only instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderbookTransactor, error) {
	contract, err := bindOrderbook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderbookTransactor{contract: contract}, nil
}

// NewOrderbookFilterer creates a new log filterer instance of Orderbook, bound to a specific deployed contract.
func NewOrderbookFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderbookFilterer, error) {
	contract, err := bindOrderbook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderbookFilterer{contract: contract}, nil
}

// bindOrderbook binds a generic wrapper to an already deployed contract.
func bindOrderbook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderbookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.OrderbookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.OrderbookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orderbook *OrderbookCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orderbook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orderbook *OrderbookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orderbook *OrderbookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orderbook.Contract.contract.Transact(opts, method, params...)
}

// BuyOrder is a free data retrieval call binding the contract method 0x22f85eaa.
//
// Solidity: function buyOrder(_index uint256) constant returns(bytes32, bool)
func (_Orderbook *OrderbookCaller) BuyOrder(opts *bind.CallOpts, _index *big.Int) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Orderbook.contract.Call(opts, out, "buyOrder", _index)
	return *ret0, *ret1, err
}

// BuyOrder is a free data retrieval call binding the contract method 0x22f85eaa.
//
// Solidity: function buyOrder(_index uint256) constant returns(bytes32, bool)
func (_Orderbook *OrderbookSession) BuyOrder(_index *big.Int) ([32]byte, bool, error) {
	return _Orderbook.Contract.BuyOrder(&_Orderbook.CallOpts, _index)
}

// BuyOrder is a free data retrieval call binding the contract method 0x22f85eaa.
//
// Solidity: function buyOrder(_index uint256) constant returns(bytes32, bool)
func (_Orderbook *OrderbookCallerSession) BuyOrder(_index *big.Int) ([32]byte, bool, error) {
	return _Orderbook.Contract.BuyOrder(&_Orderbook.CallOpts, _index)
}

// BuyOrders is a free data retrieval call binding the contract method 0x35cea288.
//
// Solidity: function buyOrders( uint256) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) BuyOrders(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "buyOrders", arg0)
	return *ret0, err
}

// BuyOrders is a free data retrieval call binding the contract method 0x35cea288.
//
// Solidity: function buyOrders( uint256) constant returns(bytes32)
func (_Orderbook *OrderbookSession) BuyOrders(arg0 *big.Int) ([32]byte, error) {
	return _Orderbook.Contract.BuyOrders(&_Orderbook.CallOpts, arg0)
}

// BuyOrders is a free data retrieval call binding the contract method 0x35cea288.
//
// Solidity: function buyOrders( uint256) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) BuyOrders(arg0 *big.Int) ([32]byte, error) {
	return _Orderbook.Contract.BuyOrders(&_Orderbook.CallOpts, arg0)
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_Orderbook *OrderbookCaller) DarknodeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "darknodeRegistry")
	return *ret0, err
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_Orderbook *OrderbookSession) DarknodeRegistry() (common.Address, error) {
	return _Orderbook.Contract.DarknodeRegistry(&_Orderbook.CallOpts)
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_Orderbook *OrderbookCallerSession) DarknodeRegistry() (common.Address, error) {
	return _Orderbook.Contract.DarknodeRegistry(&_Orderbook.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Orderbook *OrderbookCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Orderbook *OrderbookSession) Fee() (*big.Int, error) {
	return _Orderbook.Contract.Fee(&_Orderbook.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Orderbook *OrderbookCallerSession) Fee() (*big.Int, error) {
	return _Orderbook.Contract.Fee(&_Orderbook.CallOpts)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(_index uint256) constant returns(bytes32, bool)
func (_Orderbook *OrderbookCaller) GetOrder(opts *bind.CallOpts, _index *big.Int) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Orderbook.contract.Call(opts, out, "getOrder", _index)
	return *ret0, *ret1, err
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(_index uint256) constant returns(bytes32, bool)
func (_Orderbook *OrderbookSession) GetOrder(_index *big.Int) ([32]byte, bool, error) {
	return _Orderbook.Contract.GetOrder(&_Orderbook.CallOpts, _index)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(_index uint256) constant returns(bytes32, bool)
func (_Orderbook *OrderbookCallerSession) GetOrder(_index *big.Int) ([32]byte, bool, error) {
	return _Orderbook.Contract.GetOrder(&_Orderbook.CallOpts, _index)
}

// GetOrders is a free data retrieval call binding the contract method 0x8f72fc77.
//
// Solidity: function getOrders(_offset uint256, _limit uint256) constant returns(bytes32[], address[], uint8[])
func (_Orderbook *OrderbookCaller) GetOrders(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([][32]byte, []common.Address, []uint8, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([]common.Address)
		ret2 = new([]uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Orderbook.contract.Call(opts, out, "getOrders", _offset, _limit)
	return *ret0, *ret1, *ret2, err
}

// GetOrders is a free data retrieval call binding the contract method 0x8f72fc77.
//
// Solidity: function getOrders(_offset uint256, _limit uint256) constant returns(bytes32[], address[], uint8[])
func (_Orderbook *OrderbookSession) GetOrders(_offset *big.Int, _limit *big.Int) ([][32]byte, []common.Address, []uint8, error) {
	return _Orderbook.Contract.GetOrders(&_Orderbook.CallOpts, _offset, _limit)
}

// GetOrders is a free data retrieval call binding the contract method 0x8f72fc77.
//
// Solidity: function getOrders(_offset uint256, _limit uint256) constant returns(bytes32[], address[], uint8[])
func (_Orderbook *OrderbookCallerSession) GetOrders(_offset *big.Int, _limit *big.Int) ([][32]byte, []common.Address, []uint8, error) {
	return _Orderbook.Contract.GetOrders(&_Orderbook.CallOpts, _offset, _limit)
}

// GetOrdersCount is a free data retrieval call binding the contract method 0xb5b3b051.
//
// Solidity: function getOrdersCount() constant returns(uint256)
func (_Orderbook *OrderbookCaller) GetOrdersCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "getOrdersCount")
	return *ret0, err
}

// GetOrdersCount is a free data retrieval call binding the contract method 0xb5b3b051.
//
// Solidity: function getOrdersCount() constant returns(uint256)
func (_Orderbook *OrderbookSession) GetOrdersCount() (*big.Int, error) {
	return _Orderbook.Contract.GetOrdersCount(&_Orderbook.CallOpts)
}

// GetOrdersCount is a free data retrieval call binding the contract method 0xb5b3b051.
//
// Solidity: function getOrdersCount() constant returns(uint256)
func (_Orderbook *OrderbookCallerSession) GetOrdersCount() (*big.Int, error) {
	return _Orderbook.Contract.GetOrdersCount(&_Orderbook.CallOpts)
}

// OrderBlockNumber is a free data retrieval call binding the contract method 0x89895d53.
//
// Solidity: function orderBlockNumber(_orderId bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCaller) OrderBlockNumber(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderBlockNumber", _orderId)
	return *ret0, err
}

// OrderBlockNumber is a free data retrieval call binding the contract method 0x89895d53.
//
// Solidity: function orderBlockNumber(_orderId bytes32) constant returns(uint256)
func (_Orderbook *OrderbookSession) OrderBlockNumber(_orderId [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderBlockNumber(&_Orderbook.CallOpts, _orderId)
}

// OrderBlockNumber is a free data retrieval call binding the contract method 0x89895d53.
//
// Solidity: function orderBlockNumber(_orderId bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCallerSession) OrderBlockNumber(_orderId [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderBlockNumber(&_Orderbook.CallOpts, _orderId)
}

// OrderBroker is a free data retrieval call binding the contract method 0xa5181661.
//
// Solidity: function orderBroker(_orderId bytes32) constant returns(address)
func (_Orderbook *OrderbookCaller) OrderBroker(opts *bind.CallOpts, _orderId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderBroker", _orderId)
	return *ret0, err
}

// OrderBroker is a free data retrieval call binding the contract method 0xa5181661.
//
// Solidity: function orderBroker(_orderId bytes32) constant returns(address)
func (_Orderbook *OrderbookSession) OrderBroker(_orderId [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderBroker(&_Orderbook.CallOpts, _orderId)
}

// OrderBroker is a free data retrieval call binding the contract method 0xa5181661.
//
// Solidity: function orderBroker(_orderId bytes32) constant returns(address)
func (_Orderbook *OrderbookCallerSession) OrderBroker(_orderId [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderBroker(&_Orderbook.CallOpts, _orderId)
}

// OrderConfirmer is a free data retrieval call binding the contract method 0x1107c3f7.
//
// Solidity: function orderConfirmer(_orderId bytes32) constant returns(address)
func (_Orderbook *OrderbookCaller) OrderConfirmer(opts *bind.CallOpts, _orderId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderConfirmer", _orderId)
	return *ret0, err
}

// OrderConfirmer is a free data retrieval call binding the contract method 0x1107c3f7.
//
// Solidity: function orderConfirmer(_orderId bytes32) constant returns(address)
func (_Orderbook *OrderbookSession) OrderConfirmer(_orderId [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderConfirmer(&_Orderbook.CallOpts, _orderId)
}

// OrderConfirmer is a free data retrieval call binding the contract method 0x1107c3f7.
//
// Solidity: function orderConfirmer(_orderId bytes32) constant returns(address)
func (_Orderbook *OrderbookCallerSession) OrderConfirmer(_orderId [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderConfirmer(&_Orderbook.CallOpts, _orderId)
}

// OrderDepth is a free data retrieval call binding the contract method 0xa188fcb8.
//
// Solidity: function orderDepth(_orderId bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCaller) OrderDepth(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderDepth", _orderId)
	return *ret0, err
}

// OrderDepth is a free data retrieval call binding the contract method 0xa188fcb8.
//
// Solidity: function orderDepth(_orderId bytes32) constant returns(uint256)
func (_Orderbook *OrderbookSession) OrderDepth(_orderId [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderDepth(&_Orderbook.CallOpts, _orderId)
}

// OrderDepth is a free data retrieval call binding the contract method 0xa188fcb8.
//
// Solidity: function orderDepth(_orderId bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCallerSession) OrderDepth(_orderId [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderDepth(&_Orderbook.CallOpts, _orderId)
}

// OrderMatch is a free data retrieval call binding the contract method 0xaf3e8a40.
//
// Solidity: function orderMatch(_orderId bytes32) constant returns(bytes32[])
func (_Orderbook *OrderbookCaller) OrderMatch(opts *bind.CallOpts, _orderId [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderMatch", _orderId)
	return *ret0, err
}

// OrderMatch is a free data retrieval call binding the contract method 0xaf3e8a40.
//
// Solidity: function orderMatch(_orderId bytes32) constant returns(bytes32[])
func (_Orderbook *OrderbookSession) OrderMatch(_orderId [32]byte) ([][32]byte, error) {
	return _Orderbook.Contract.OrderMatch(&_Orderbook.CallOpts, _orderId)
}

// OrderMatch is a free data retrieval call binding the contract method 0xaf3e8a40.
//
// Solidity: function orderMatch(_orderId bytes32) constant returns(bytes32[])
func (_Orderbook *OrderbookCallerSession) OrderMatch(_orderId [32]byte) ([][32]byte, error) {
	return _Orderbook.Contract.OrderMatch(&_Orderbook.CallOpts, _orderId)
}

// OrderPriority is a free data retrieval call binding the contract method 0xb248e4e1.
//
// Solidity: function orderPriority(_orderId bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCaller) OrderPriority(opts *bind.CallOpts, _orderId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderPriority", _orderId)
	return *ret0, err
}

// OrderPriority is a free data retrieval call binding the contract method 0xb248e4e1.
//
// Solidity: function orderPriority(_orderId bytes32) constant returns(uint256)
func (_Orderbook *OrderbookSession) OrderPriority(_orderId [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderPriority(&_Orderbook.CallOpts, _orderId)
}

// OrderPriority is a free data retrieval call binding the contract method 0xb248e4e1.
//
// Solidity: function orderPriority(_orderId bytes32) constant returns(uint256)
func (_Orderbook *OrderbookCallerSession) OrderPriority(_orderId [32]byte) (*big.Int, error) {
	return _Orderbook.Contract.OrderPriority(&_Orderbook.CallOpts, _orderId)
}

// OrderState is a free data retrieval call binding the contract method 0xaab14d04.
//
// Solidity: function orderState(_orderId bytes32) constant returns(uint8)
func (_Orderbook *OrderbookCaller) OrderState(opts *bind.CallOpts, _orderId [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderState", _orderId)
	return *ret0, err
}

// OrderState is a free data retrieval call binding the contract method 0xaab14d04.
//
// Solidity: function orderState(_orderId bytes32) constant returns(uint8)
func (_Orderbook *OrderbookSession) OrderState(_orderId [32]byte) (uint8, error) {
	return _Orderbook.Contract.OrderState(&_Orderbook.CallOpts, _orderId)
}

// OrderState is a free data retrieval call binding the contract method 0xaab14d04.
//
// Solidity: function orderState(_orderId bytes32) constant returns(uint8)
func (_Orderbook *OrderbookCallerSession) OrderState(_orderId [32]byte) (uint8, error) {
	return _Orderbook.Contract.OrderState(&_Orderbook.CallOpts, _orderId)
}

// OrderTrader is a free data retrieval call binding the contract method 0xb1a08010.
//
// Solidity: function orderTrader(_orderId bytes32) constant returns(address)
func (_Orderbook *OrderbookCaller) OrderTrader(opts *bind.CallOpts, _orderId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "orderTrader", _orderId)
	return *ret0, err
}

// OrderTrader is a free data retrieval call binding the contract method 0xb1a08010.
//
// Solidity: function orderTrader(_orderId bytes32) constant returns(address)
func (_Orderbook *OrderbookSession) OrderTrader(_orderId [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderTrader(&_Orderbook.CallOpts, _orderId)
}

// OrderTrader is a free data retrieval call binding the contract method 0xb1a08010.
//
// Solidity: function orderTrader(_orderId bytes32) constant returns(address)
func (_Orderbook *OrderbookCallerSession) OrderTrader(_orderId [32]byte) (common.Address, error) {
	return _Orderbook.Contract.OrderTrader(&_Orderbook.CallOpts, _orderId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Orderbook *OrderbookCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Orderbook *OrderbookSession) Owner() (common.Address, error) {
	return _Orderbook.Contract.Owner(&_Orderbook.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Orderbook *OrderbookCallerSession) Owner() (common.Address, error) {
	return _Orderbook.Contract.Owner(&_Orderbook.CallOpts)
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_Orderbook *OrderbookCaller) Ren(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "ren")
	return *ret0, err
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_Orderbook *OrderbookSession) Ren() (common.Address, error) {
	return _Orderbook.Contract.Ren(&_Orderbook.CallOpts)
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_Orderbook *OrderbookCallerSession) Ren() (common.Address, error) {
	return _Orderbook.Contract.Ren(&_Orderbook.CallOpts)
}

// SellOrder is a free data retrieval call binding the contract method 0x97514d90.
//
// Solidity: function sellOrder(_index uint256) constant returns(bytes32, bool)
func (_Orderbook *OrderbookCaller) SellOrder(opts *bind.CallOpts, _index *big.Int) ([32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Orderbook.contract.Call(opts, out, "sellOrder", _index)
	return *ret0, *ret1, err
}

// SellOrder is a free data retrieval call binding the contract method 0x97514d90.
//
// Solidity: function sellOrder(_index uint256) constant returns(bytes32, bool)
func (_Orderbook *OrderbookSession) SellOrder(_index *big.Int) ([32]byte, bool, error) {
	return _Orderbook.Contract.SellOrder(&_Orderbook.CallOpts, _index)
}

// SellOrder is a free data retrieval call binding the contract method 0x97514d90.
//
// Solidity: function sellOrder(_index uint256) constant returns(bytes32, bool)
func (_Orderbook *OrderbookCallerSession) SellOrder(_index *big.Int) ([32]byte, bool, error) {
	return _Orderbook.Contract.SellOrder(&_Orderbook.CallOpts, _index)
}

// SellOrders is a free data retrieval call binding the contract method 0x4a8393f3.
//
// Solidity: function sellOrders( uint256) constant returns(bytes32)
func (_Orderbook *OrderbookCaller) SellOrders(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orderbook.contract.Call(opts, out, "sellOrders", arg0)
	return *ret0, err
}

// SellOrders is a free data retrieval call binding the contract method 0x4a8393f3.
//
// Solidity: function sellOrders( uint256) constant returns(bytes32)
func (_Orderbook *OrderbookSession) SellOrders(arg0 *big.Int) ([32]byte, error) {
	return _Orderbook.Contract.SellOrders(&_Orderbook.CallOpts, arg0)
}

// SellOrders is a free data retrieval call binding the contract method 0x4a8393f3.
//
// Solidity: function sellOrders( uint256) constant returns(bytes32)
func (_Orderbook *OrderbookCallerSession) SellOrders(arg0 *big.Int) ([32]byte, error) {
	return _Orderbook.Contract.SellOrders(&_Orderbook.CallOpts, arg0)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x574ed6c1.
//
// Solidity: function cancelOrder(_signature bytes, _orderId bytes32) returns()
func (_Orderbook *OrderbookTransactor) CancelOrder(opts *bind.TransactOpts, _signature []byte, _orderId [32]byte) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "cancelOrder", _signature, _orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x574ed6c1.
//
// Solidity: function cancelOrder(_signature bytes, _orderId bytes32) returns()
func (_Orderbook *OrderbookSession) CancelOrder(_signature []byte, _orderId [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.CancelOrder(&_Orderbook.TransactOpts, _signature, _orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x574ed6c1.
//
// Solidity: function cancelOrder(_signature bytes, _orderId bytes32) returns()
func (_Orderbook *OrderbookTransactorSession) CancelOrder(_signature []byte, _orderId [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.CancelOrder(&_Orderbook.TransactOpts, _signature, _orderId)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x7008b996.
//
// Solidity: function confirmOrder(_orderId bytes32, _orderMatches bytes32[]) returns()
func (_Orderbook *OrderbookTransactor) ConfirmOrder(opts *bind.TransactOpts, _orderId [32]byte, _orderMatches [][32]byte) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "confirmOrder", _orderId, _orderMatches)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x7008b996.
//
// Solidity: function confirmOrder(_orderId bytes32, _orderMatches bytes32[]) returns()
func (_Orderbook *OrderbookSession) ConfirmOrder(_orderId [32]byte, _orderMatches [][32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.ConfirmOrder(&_Orderbook.TransactOpts, _orderId, _orderMatches)
}

// ConfirmOrder is a paid mutator transaction binding the contract method 0x7008b996.
//
// Solidity: function confirmOrder(_orderId bytes32, _orderMatches bytes32[]) returns()
func (_Orderbook *OrderbookTransactorSession) ConfirmOrder(_orderId [32]byte, _orderMatches [][32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.ConfirmOrder(&_Orderbook.TransactOpts, _orderId, _orderMatches)
}

// OpenBuyOrder is a paid mutator transaction binding the contract method 0x5060340b.
//
// Solidity: function openBuyOrder(_signature bytes, _orderId bytes32) returns()
func (_Orderbook *OrderbookTransactor) OpenBuyOrder(opts *bind.TransactOpts, _signature []byte, _orderId [32]byte) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "openBuyOrder", _signature, _orderId)
}

// OpenBuyOrder is a paid mutator transaction binding the contract method 0x5060340b.
//
// Solidity: function openBuyOrder(_signature bytes, _orderId bytes32) returns()
func (_Orderbook *OrderbookSession) OpenBuyOrder(_signature []byte, _orderId [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.OpenBuyOrder(&_Orderbook.TransactOpts, _signature, _orderId)
}

// OpenBuyOrder is a paid mutator transaction binding the contract method 0x5060340b.
//
// Solidity: function openBuyOrder(_signature bytes, _orderId bytes32) returns()
func (_Orderbook *OrderbookTransactorSession) OpenBuyOrder(_signature []byte, _orderId [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.OpenBuyOrder(&_Orderbook.TransactOpts, _signature, _orderId)
}

// OpenSellOrder is a paid mutator transaction binding the contract method 0x39b0d677.
//
// Solidity: function openSellOrder(_signature bytes, _orderId bytes32) returns()
func (_Orderbook *OrderbookTransactor) OpenSellOrder(opts *bind.TransactOpts, _signature []byte, _orderId [32]byte) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "openSellOrder", _signature, _orderId)
}

// OpenSellOrder is a paid mutator transaction binding the contract method 0x39b0d677.
//
// Solidity: function openSellOrder(_signature bytes, _orderId bytes32) returns()
func (_Orderbook *OrderbookSession) OpenSellOrder(_signature []byte, _orderId [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.OpenSellOrder(&_Orderbook.TransactOpts, _signature, _orderId)
}

// OpenSellOrder is a paid mutator transaction binding the contract method 0x39b0d677.
//
// Solidity: function openSellOrder(_signature bytes, _orderId bytes32) returns()
func (_Orderbook *OrderbookTransactorSession) OpenSellOrder(_signature []byte, _orderId [32]byte) (*types.Transaction, error) {
	return _Orderbook.Contract.OpenSellOrder(&_Orderbook.TransactOpts, _signature, _orderId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Orderbook *OrderbookTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Orderbook *OrderbookSession) RenounceOwnership() (*types.Transaction, error) {
	return _Orderbook.Contract.RenounceOwnership(&_Orderbook.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Orderbook *OrderbookTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Orderbook.Contract.RenounceOwnership(&_Orderbook.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Orderbook *OrderbookTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Orderbook *OrderbookSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.TransferOwnership(&_Orderbook.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Orderbook *OrderbookTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.TransferOwnership(&_Orderbook.TransactOpts, _newOwner)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_Orderbook *OrderbookTransactor) UpdateDarknodeRegistry(opts *bind.TransactOpts, _newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "updateDarknodeRegistry", _newDarknodeRegistry)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_Orderbook *OrderbookSession) UpdateDarknodeRegistry(_newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.UpdateDarknodeRegistry(&_Orderbook.TransactOpts, _newDarknodeRegistry)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_Orderbook *OrderbookTransactorSession) UpdateDarknodeRegistry(_newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _Orderbook.Contract.UpdateDarknodeRegistry(&_Orderbook.TransactOpts, _newDarknodeRegistry)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_newFee uint256) returns()
func (_Orderbook *OrderbookTransactor) UpdateFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _Orderbook.contract.Transact(opts, "updateFee", _newFee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_newFee uint256) returns()
func (_Orderbook *OrderbookSession) UpdateFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Orderbook.Contract.UpdateFee(&_Orderbook.TransactOpts, _newFee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_newFee uint256) returns()
func (_Orderbook *OrderbookTransactorSession) UpdateFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Orderbook.Contract.UpdateFee(&_Orderbook.TransactOpts, _newFee)
}

// OrderbookDarknodeRegistryUpdatedIterator is returned from FilterDarknodeRegistryUpdated and is used to iterate over the raw logs and unpacked data for DarknodeRegistryUpdated events raised by the Orderbook contract.
type OrderbookDarknodeRegistryUpdatedIterator struct {
	Event *OrderbookDarknodeRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *OrderbookDarknodeRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderbookDarknodeRegistryUpdated)
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
		it.Event = new(OrderbookDarknodeRegistryUpdated)
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
func (it *OrderbookDarknodeRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderbookDarknodeRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderbookDarknodeRegistryUpdated represents a DarknodeRegistryUpdated event raised by the Orderbook contract.
type OrderbookDarknodeRegistryUpdated struct {
	PreviousDarknodeRegistry common.Address
	NextDarknodeRegistry     common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterDarknodeRegistryUpdated is a free log retrieval operation binding the contract event 0xd4443dde7b09a100132dba56b0175378b6ad81b64398aef77626357a1d8daf02.
//
// Solidity: e DarknodeRegistryUpdated(previousDarknodeRegistry address, nextDarknodeRegistry address)
func (_Orderbook *OrderbookFilterer) FilterDarknodeRegistryUpdated(opts *bind.FilterOpts) (*OrderbookDarknodeRegistryUpdatedIterator, error) {

	logs, sub, err := _Orderbook.contract.FilterLogs(opts, "DarknodeRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &OrderbookDarknodeRegistryUpdatedIterator{contract: _Orderbook.contract, event: "DarknodeRegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchDarknodeRegistryUpdated is a free log subscription operation binding the contract event 0xd4443dde7b09a100132dba56b0175378b6ad81b64398aef77626357a1d8daf02.
//
// Solidity: e DarknodeRegistryUpdated(previousDarknodeRegistry address, nextDarknodeRegistry address)
func (_Orderbook *OrderbookFilterer) WatchDarknodeRegistryUpdated(opts *bind.WatchOpts, sink chan<- *OrderbookDarknodeRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _Orderbook.contract.WatchLogs(opts, "DarknodeRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderbookDarknodeRegistryUpdated)
				if err := _Orderbook.contract.UnpackLog(event, "DarknodeRegistryUpdated", log); err != nil {
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

// OrderbookFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the Orderbook contract.
type OrderbookFeeUpdatedIterator struct {
	Event *OrderbookFeeUpdated // Event containing the contract specifics and raw log

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
func (it *OrderbookFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderbookFeeUpdated)
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
		it.Event = new(OrderbookFeeUpdated)
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
func (it *OrderbookFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderbookFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderbookFeeUpdated represents a FeeUpdated event raised by the Orderbook contract.
type OrderbookFeeUpdated struct {
	PreviousFee *big.Int
	NextFee     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x528d9479e9f9889a87a3c30c7f7ba537e5e59c4c85a37733b16e57c62df61302.
//
// Solidity: e FeeUpdated(previousFee uint256, nextFee uint256)
func (_Orderbook *OrderbookFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*OrderbookFeeUpdatedIterator, error) {

	logs, sub, err := _Orderbook.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &OrderbookFeeUpdatedIterator{contract: _Orderbook.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x528d9479e9f9889a87a3c30c7f7ba537e5e59c4c85a37733b16e57c62df61302.
//
// Solidity: e FeeUpdated(previousFee uint256, nextFee uint256)
func (_Orderbook *OrderbookFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *OrderbookFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Orderbook.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderbookFeeUpdated)
				if err := _Orderbook.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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

// OrderbookOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Orderbook contract.
type OrderbookOwnershipRenouncedIterator struct {
	Event *OrderbookOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OrderbookOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderbookOwnershipRenounced)
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
		it.Event = new(OrderbookOwnershipRenounced)
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
func (it *OrderbookOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderbookOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderbookOwnershipRenounced represents a OwnershipRenounced event raised by the Orderbook contract.
type OrderbookOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Orderbook *OrderbookFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OrderbookOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Orderbook.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OrderbookOwnershipRenouncedIterator{contract: _Orderbook.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Orderbook *OrderbookFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OrderbookOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Orderbook.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderbookOwnershipRenounced)
				if err := _Orderbook.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// OrderbookOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Orderbook contract.
type OrderbookOwnershipTransferredIterator struct {
	Event *OrderbookOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OrderbookOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderbookOwnershipTransferred)
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
		it.Event = new(OrderbookOwnershipTransferred)
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
func (it *OrderbookOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderbookOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderbookOwnershipTransferred represents a OwnershipTransferred event raised by the Orderbook contract.
type OrderbookOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Orderbook *OrderbookFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OrderbookOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Orderbook.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OrderbookOwnershipTransferredIterator{contract: _Orderbook.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Orderbook *OrderbookFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OrderbookOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Orderbook.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderbookOwnershipTransferred)
				if err := _Orderbook.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319163317905561020b806100326000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a6811461005b5780638da5cb5b14610072578063f2fde38b146100a3575b600080fd5b34801561006757600080fd5b506100706100c4565b005b34801561007e57600080fd5b50610087610130565b60408051600160a060020a039092168252519081900360200190f35b3480156100af57600080fd5b50610070600160a060020a036004351661013f565b600054600160a060020a031633146100db57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a0316331461015657600080fd5b61015f81610162565b50565b600160a060020a038116151561017757600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820249fc14a52d1f495a8d6f604ce65945720d6bbc4d103e7768b1b35e0021ea7450029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
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
		it.Event = new(OwnableOwnershipRenounced)
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
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PausableBin is the compiled bytecode used for deploying new contracts.
const PausableBin = `0x608060405260008054600160a860020a031916331790556103c4806100256000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a811461007c5780635c975abb14610093578063715018a6146100bc5780638456cb59146100d15780638da5cb5b146100e6578063f2fde38b14610117575b600080fd5b34801561008857600080fd5b50610091610138565b005b34801561009f57600080fd5b506100a86101bf565b604080519115158252519081900360200190f35b3480156100c857600080fd5b506100916101e0565b3480156100dd57600080fd5b5061009161024c565b3480156100f257600080fd5b506100fb6102e9565b60408051600160a060020a039092168252519081900360200190f35b34801561012357600080fd5b50610091600160a060020a03600435166102f8565b600054600160a060020a0316331461014f57600080fd5b60005474010000000000000000000000000000000000000000900460ff16151561017857600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b60005474010000000000000000000000000000000000000000900460ff1681565b600054600160a060020a031633146101f757600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a0316331461026357600080fd5b60005474010000000000000000000000000000000000000000900460ff161561028b57600080fd5b6000805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b600054600160a060020a0316331461030f57600080fd5b6103188161031b565b50565b600160a060020a038116151561033057600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820121c33da62d8c7585f47939a982827ef0939aff70bacefc63bc4ca9115512e4c0029`

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.RenounceOwnership(&_Pausable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pausable *PausableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.RenounceOwnership(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Pausable *PausableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Pausable contract.
type PausableOwnershipRenouncedIterator struct {
	Event *PausableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipRenounced)
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
		it.Event = new(PausableOwnershipRenounced)
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
func (it *PausableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipRenounced represents a OwnershipRenounced event raised by the Pausable contract.
type PausableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Pausable *PausableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PausableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipRenouncedIterator{contract: _Pausable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Pausable *PausableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PausableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipRenounced)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PausableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pausable contract.
type PausableOwnershipTransferredIterator struct {
	Event *PausableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipTransferred)
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
		it.Event = new(PausableOwnershipTransferred)
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
func (it *PausableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipTransferred represents a OwnershipTransferred event raised by the Pausable contract.
type PausableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Pausable *PausableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipTransferredIterator{contract: _Pausable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Pausable *PausableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipTransferred)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

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
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
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
		it.Event = new(PausablePause)
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
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

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
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
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
		it.Event = new(PausableUnpause)
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
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// PausableTokenABI is the input ABI used to generate the binding from.
const PausableTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// PausableTokenBin is the compiled bytecode used for deploying new contracts.
const PausableTokenBin = `0x608060405260038054600160a860020a03191633179055610a84806100256000396000f3006080604052600436106100cf5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100d457806318160ddd1461010c57806323b872dd146101335780633f4ba83a1461015d5780635c975abb14610174578063661884631461018957806370a08231146101ad578063715018a6146101ce5780638456cb59146101e35780638da5cb5b146101f8578063a9059cbb14610229578063d73dd6231461024d578063dd62ed3e14610271578063f2fde38b14610298575b600080fd5b3480156100e057600080fd5b506100f8600160a060020a03600435166024356102b9565b604080519115158252519081900360200190f35b34801561011857600080fd5b506101216102e4565b60408051918252519081900360200190f35b34801561013f57600080fd5b506100f8600160a060020a03600435811690602435166044356102ea565b34801561016957600080fd5b50610172610317565b005b34801561018057600080fd5b506100f861038f565b34801561019557600080fd5b506100f8600160a060020a036004351660243561039f565b3480156101b957600080fd5b50610121600160a060020a03600435166103c3565b3480156101da57600080fd5b506101726103de565b3480156101ef57600080fd5b5061017261044c565b34801561020457600080fd5b5061020d6104c9565b60408051600160a060020a039092168252519081900360200190f35b34801561023557600080fd5b506100f8600160a060020a03600435166024356104d8565b34801561025957600080fd5b506100f8600160a060020a03600435166024356104fc565b34801561027d57600080fd5b50610121600160a060020a0360043581169060243516610520565b3480156102a457600080fd5b50610172600160a060020a036004351661054b565b60035460009060a060020a900460ff16156102d357600080fd5b6102dd838361056e565b9392505050565b60015490565b60035460009060a060020a900460ff161561030457600080fd5b61030f8484846105d4565b949350505050565b600354600160a060020a0316331461032e57600080fd5b60035460a060020a900460ff16151561034657600080fd5b6003805474ff0000000000000000000000000000000000000000191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b60035460a060020a900460ff1681565b60035460009060a060020a900460ff16156103b957600080fd5b6102dd838361074b565b600160a060020a031660009081526020819052604090205490565b600354600160a060020a031633146103f557600080fd5b600354604051600160a060020a03909116907ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482090600090a26003805473ffffffffffffffffffffffffffffffffffffffff19169055565b600354600160a060020a0316331461046357600080fd5b60035460a060020a900460ff161561047a57600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b600354600160a060020a031681565b60035460009060a060020a900460ff16156104f257600080fd5b6102dd838361083b565b60035460009060a060020a900460ff161561051657600080fd5b6102dd838361091c565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600160a060020a0316331461056257600080fd5b61056b816109b5565b50565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b6000600160a060020a03831615156105eb57600080fd5b600160a060020a03841660009081526020819052604090205482111561061057600080fd5b600160a060020a038416600090815260026020908152604080832033845290915290205482111561064057600080fd5b600160a060020a038416600090815260208190526040902054610669908363ffffffff610a3316565b600160a060020a03808616600090815260208190526040808220939093559085168152205461069e908363ffffffff610a4516565b600160a060020a038085166000908152602081815260408083209490945591871681526002825282812033825290915220546106e0908363ffffffff610a3316565b600160a060020a03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b336000908152600260209081526040808320600160a060020a0386168452909152812054808311156107a057336000908152600260209081526040808320600160a060020a03881684529091528120556107d5565b6107b0818463ffffffff610a3316565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b6000600160a060020a038316151561085257600080fd5b3360009081526020819052604090205482111561086e57600080fd5b3360009081526020819052604090205461088e908363ffffffff610a3316565b3360009081526020819052604080822092909255600160a060020a038516815220546108c0908363ffffffff610a4516565b600160a060020a038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600260209081526040808320600160a060020a0386168452909152812054610950908363ffffffff610a4516565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03811615156109ca57600080fd5b600354604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610a3f57fe5b50900390565b81810182811015610a5257fe5b929150505600a165627a7a7230582004967fa2580597ff6ac4646288111c2e5ae8307fc3e205544589a904e3c70e5f0029`

// DeployPausableToken deploys a new Ethereum contract, binding an instance of PausableToken to it.
func DeployPausableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PausableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PausableToken{PausableTokenCaller: PausableTokenCaller{contract: contract}, PausableTokenTransactor: PausableTokenTransactor{contract: contract}, PausableTokenFilterer: PausableTokenFilterer{contract: contract}}, nil
}

// PausableToken is an auto generated Go binding around an Ethereum contract.
type PausableToken struct {
	PausableTokenCaller     // Read-only binding to the contract
	PausableTokenTransactor // Write-only binding to the contract
	PausableTokenFilterer   // Log filterer for contract events
}

// PausableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableTokenSession struct {
	Contract     *PausableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableTokenCallerSession struct {
	Contract *PausableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PausableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTokenTransactorSession struct {
	Contract     *PausableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PausableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableTokenRaw struct {
	Contract *PausableToken // Generic contract binding to access the raw methods on
}

// PausableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableTokenCallerRaw struct {
	Contract *PausableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTokenTransactorRaw struct {
	Contract *PausableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableToken creates a new instance of PausableToken, bound to a specific deployed contract.
func NewPausableToken(address common.Address, backend bind.ContractBackend) (*PausableToken, error) {
	contract, err := bindPausableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableToken{PausableTokenCaller: PausableTokenCaller{contract: contract}, PausableTokenTransactor: PausableTokenTransactor{contract: contract}, PausableTokenFilterer: PausableTokenFilterer{contract: contract}}, nil
}

// NewPausableTokenCaller creates a new read-only instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenCaller(address common.Address, caller bind.ContractCaller) (*PausableTokenCaller, error) {
	contract, err := bindPausableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTokenCaller{contract: contract}, nil
}

// NewPausableTokenTransactor creates a new write-only instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTokenTransactor, error) {
	contract, err := bindPausableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTokenTransactor{contract: contract}, nil
}

// NewPausableTokenFilterer creates a new log filterer instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableTokenFilterer, error) {
	contract, err := bindPausableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableTokenFilterer{contract: contract}, nil
}

// bindPausableToken binds a generic wrapper to an already deployed contract.
func bindPausableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableToken *PausableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableToken.Contract.PausableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableToken *PausableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.Contract.PausableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableToken *PausableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableToken.Contract.PausableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableToken *PausableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableToken *PausableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableToken *PausableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PausableToken.Contract.Allowance(&_PausableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PausableToken.Contract.Allowance(&_PausableToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_PausableToken *PausableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_PausableToken *PausableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PausableToken.Contract.BalanceOf(&_PausableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_PausableToken *PausableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PausableToken.Contract.BalanceOf(&_PausableToken.CallOpts, _owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenSession) Owner() (common.Address, error) {
	return _PausableToken.Contract.Owner(&_PausableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenCallerSession) Owner() (common.Address, error) {
	return _PausableToken.Contract.Owner(&_PausableToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenSession) Paused() (bool, error) {
	return _PausableToken.Contract.Paused(&_PausableToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenCallerSession) Paused() (bool, error) {
	return _PausableToken.Contract.Paused(&_PausableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenSession) TotalSupply() (*big.Int, error) {
	return _PausableToken.Contract.TotalSupply(&_PausableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PausableToken.Contract.TotalSupply(&_PausableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Approve(&_PausableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Approve(&_PausableToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.DecreaseApproval(&_PausableToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.DecreaseApproval(&_PausableToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.IncreaseApproval(&_PausableToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.IncreaseApproval(&_PausableToken.TransactOpts, _spender, _addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenSession) Pause() (*types.Transaction, error) {
	return _PausableToken.Contract.Pause(&_PausableToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _PausableToken.Contract.Pause(&_PausableToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PausableToken *PausableTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PausableToken *PausableTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _PausableToken.Contract.RenounceOwnership(&_PausableToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PausableToken *PausableTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PausableToken.Contract.RenounceOwnership(&_PausableToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Transfer(&_PausableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Transfer(&_PausableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferFrom(&_PausableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferFrom(&_PausableToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PausableToken *PausableTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PausableToken *PausableTokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferOwnership(&_PausableToken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_PausableToken *PausableTokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferOwnership(&_PausableToken.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenSession) Unpause() (*types.Transaction, error) {
	return _PausableToken.Contract.Unpause(&_PausableToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _PausableToken.Contract.Unpause(&_PausableToken.TransactOpts)
}

// PausableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PausableToken contract.
type PausableTokenApprovalIterator struct {
	Event *PausableTokenApproval // Event containing the contract specifics and raw log

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
func (it *PausableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenApproval)
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
		it.Event = new(PausableTokenApproval)
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
func (it *PausableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenApproval represents a Approval event raised by the PausableToken contract.
type PausableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PausableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenApprovalIterator{contract: _PausableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PausableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenApproval)
				if err := _PausableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PausableTokenOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PausableToken contract.
type PausableTokenOwnershipRenouncedIterator struct {
	Event *PausableTokenOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PausableTokenOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenOwnershipRenounced)
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
		it.Event = new(PausableTokenOwnershipRenounced)
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
func (it *PausableTokenOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenOwnershipRenounced represents a OwnershipRenounced event raised by the PausableToken contract.
type PausableTokenOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PausableToken *PausableTokenFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PausableTokenOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenOwnershipRenouncedIterator{contract: _PausableToken.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PausableToken *PausableTokenFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PausableTokenOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenOwnershipRenounced)
				if err := _PausableToken.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PausableTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PausableToken contract.
type PausableTokenOwnershipTransferredIterator struct {
	Event *PausableTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenOwnershipTransferred)
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
		it.Event = new(PausableTokenOwnershipTransferred)
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
func (it *PausableTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenOwnershipTransferred represents a OwnershipTransferred event raised by the PausableToken contract.
type PausableTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PausableToken *PausableTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenOwnershipTransferredIterator{contract: _PausableToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PausableToken *PausableTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenOwnershipTransferred)
				if err := _PausableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausableTokenPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PausableToken contract.
type PausableTokenPauseIterator struct {
	Event *PausableTokenPause // Event containing the contract specifics and raw log

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
func (it *PausableTokenPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenPause)
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
		it.Event = new(PausableTokenPause)
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
func (it *PausableTokenPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenPause represents a Pause event raised by the PausableToken contract.
type PausableTokenPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PausableToken *PausableTokenFilterer) FilterPause(opts *bind.FilterOpts) (*PausableTokenPauseIterator, error) {

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausableTokenPauseIterator{contract: _PausableToken.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PausableToken *PausableTokenFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausableTokenPause) (event.Subscription, error) {

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenPause)
				if err := _PausableToken.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PausableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PausableToken contract.
type PausableTokenTransferIterator struct {
	Event *PausableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PausableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenTransfer)
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
		it.Event = new(PausableTokenTransfer)
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
func (it *PausableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenTransfer represents a Transfer event raised by the PausableToken contract.
type PausableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PausableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenTransferIterator{contract: _PausableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PausableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenTransfer)
				if err := _PausableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// PausableTokenUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PausableToken contract.
type PausableTokenUnpauseIterator struct {
	Event *PausableTokenUnpause // Event containing the contract specifics and raw log

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
func (it *PausableTokenUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenUnpause)
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
		it.Event = new(PausableTokenUnpause)
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
func (it *PausableTokenUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenUnpause represents a Unpause event raised by the PausableToken contract.
type PausableTokenUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PausableToken *PausableTokenFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableTokenUnpauseIterator, error) {

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableTokenUnpauseIterator{contract: _PausableToken.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PausableToken *PausableTokenFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableTokenUnpause) (event.Subscription, error) {

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenUnpause)
				if err := _PausableToken.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// RenExBalancesABI is the input ABI used to generate the binding from.
const RenExBalancesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"traderTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSettlementContract\",\"type\":\"address\"}],\"name\":\"updateRenExSettlementContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardVaultContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_trader\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"incrementBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_trader\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"feePayee\",\"type\":\"address\"}],\"name\":\"decrementBalanceWithFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"traderBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_trader\",\"type\":\"address\"}],\"name\":\"getBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRewardVaultContract\",\"type\":\"address\"}],\"name\":\"updateRewardVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"settlementContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETHEREUM\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rewardVaultContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BalanceDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BalanceIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newRenExSettlementContract\",\"type\":\"address\"}],\"name\":\"RenExSettlementContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newRewardVaultContract\",\"type\":\"address\"}],\"name\":\"RewardVaultContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RenExBalancesBin is the compiled bytecode used for deploying new contracts.
const RenExBalancesBin = `0x608060405234801561001057600080fd5b50604051602080610f0b833981016040525160008054600160a060020a0319908116331790915560028054600160a060020a0390931692909116919091179055610eac8061005f6000396000f3006080604052600436106100da5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631a6947ce81146100df5780631d65551d1461011f57806323017a3a146101425780632b15e8571461015757806347e7ef2414610181578063715018a6146101985780638da5cb5b146101ad57806392dd79e6146101c2578063c43c633b146101f7578063c84aae1714610230578063de10f619146102ea578063ea42418b1461030b578063f2fde38b14610320578063f3fef3a314610341578063f7cdf47c14610365575b600080fd5b3480156100eb57600080fd5b50610103600160a060020a036004351660243561037a565b60408051600160a060020a039092168252519081900360200190f35b34801561012b57600080fd5b50610140600160a060020a03600435166103b1565b005b34801561014e57600080fd5b5061010361042b565b34801561016357600080fd5b50610140600160a060020a036004358116906024351660443561043a565b610140600160a060020a0360043516602435610461565b3480156101a457600080fd5b5061014061054c565b3480156101b957600080fd5b506101036105b8565b3480156101ce57600080fd5b50610140600160a060020a036004358116906024358116906044359060643590608435166105c7565b34801561020357600080fd5b5061021e600160a060020a03600435811690602435166107d8565b60408051918252519081900360200190f35b34801561023c57600080fd5b50610251600160a060020a03600435166107f5565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b8381101561029557818101518382015260200161027d565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156102d45781810151838201526020016102bc565b5050505090500194505050505060405180910390f35b3480156102f657600080fd5b50610140600160a060020a0360043516610938565b34801561031757600080fd5b506101036109b2565b34801561032c57600080fd5b50610140600160a060020a03600435166109c1565b34801561034d57600080fd5b50610140600160a060020a03600435166024356109e4565b34801561037157600080fd5b50610103610be1565b60036020528160005260406000208181548110151561039557fe5b600091825260209091200154600160a060020a03169150829050565b600054600160a060020a031633146103c857600080fd5b604051600160a060020a038216907f13ea4b9a37263b8dd3ea2fbd9c274a8845489805d06e98f12e40d216fa02507190600090a26001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031681565b600154600160a060020a0316331461045157600080fd5b61045c838383610bf9565b505050565b33600160a060020a03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14156104985734821461049357600080fd5b610541565b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a038381166004830152306024830152604482018590529151918516916323b872dd916064808201926020929091908290030181600087803b15801561050a57600080fd5b505af115801561051e573d6000803e3d6000fd5b505050506040513d602081101561053457600080fd5b5051151561054157600080fd5b61045c818484610bf9565b600054600160a060020a0316331461056357600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600154600160a060020a031633146105de57600080fd5b600160a060020a03841673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee141561069a57600254604080517f8340f549000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015287811660248301526044820186905291519190921691638340f54991859160648082019260009290919082900301818588803b15801561067c57600080fd5b505af1158015610690573d6000803e3d6000fd5b50505050506107c4565b600254604080517f095ea7b3000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201526024810185905290519186169163095ea7b3916044808201926020929091908290030181600087803b15801561070957600080fd5b505af115801561071d573d6000803e3d6000fd5b505050506040513d602081101561073357600080fd5b5050600254604080517f8340f549000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015287811660248301526044820186905291519190921691638340f54991606480830192600092919082900301818387803b1580156107ab57600080fd5b505af11580156107bf573d6000803e3d6000fd5b505050505b6107d18585848601610d38565b5050505050565b600460209081526000928352604080842090915290825290205481565b60608060608060006003600087600160a060020a0316600160a060020a0316815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561087657602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610858575b5050505050925082516040519080825280602002602001820160405280156108a8578160200160208202803883390190505b509150600090505b825181101561092d57600160a060020a038616600090815260046020526040812084519091908590849081106108e257fe5b90602001906020020151600160a060020a0316600160a060020a0316815260200190815260200160002054828281518110151561091b57fe5b602090810290910101526001016108b0565b509094909350915050565b600054600160a060020a0316331461094f57600080fd5b604051600160a060020a038216907fddbdde9ff1bc1b3e2779975798b74bbbdf2b31c326d86cfa200e381843afdd1d90600090a26002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b600054600160a060020a031633146109d857600080fd5b6109e181610dde565b50565b336000818152600460209081526040808320600160a060020a0387168452909152902054821115610a1457600080fd5b600154604080517fa3bdaedc000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301528681166024830152604482018690529151919092169163a3bdaedc9160648083019260209291908290030181600087803b158015610a8b57600080fd5b505af1158015610a9f573d6000803e3d6000fd5b505050506040513d6020811015610ab557600080fd5b50511515610ac257600080fd5b610acd818484610d38565b600160a060020a03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1415610b2e57604051600160a060020a0382169083156108fc029084906000818181858888f19350505050158015610b28573d6000803e3d6000fd5b5061045c565b82600160a060020a031663a9059cbb82846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b158015610baa57600080fd5b505af1158015610bbe573d6000803e3d6000fd5b505050506040513d6020811015610bd457600080fd5b5051151561045c57600080fd5b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b600160a060020a0380841660009081526005602090815260408083209386168352929052205460ff161515610c9257600160a060020a038084166000818152600560209081526040808320948716808452948252808320805460ff19166001908117909155938352600382528220805493840181558252902001805473ffffffffffffffffffffffffffffffffffffffff191690911790555b600160a060020a03808416600090815260046020908152604080832093861683529290522054610cc8908263ffffffff610e5b16565b600160a060020a038085166000818152600460209081526040808320948816808452948252918290209490945580519182529281019190915280820183905290517f0d66f59c9991adc17dd3339490c5058d2d6fe20395e7b55ceb6ca8019a31667d9181900360600190a1505050565b600160a060020a03808416600090815260046020908152604080832093861683529290522054610d6e908263ffffffff610e6e16565b600160a060020a038085166000818152600460209081526040808320948816808452948252918290209490945580519182529281019190915280820183905290517f6ac2cd906088d873624fa62ca95170d967629e7d964651df19a3aa2e49b44aa19181900360600190a1505050565b600160a060020a0381161515610df357600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b81810182811015610e6857fe5b92915050565b600082821115610e7a57fe5b509003905600a165627a7a72305820bf184a10ad087d35f24affa538ecf121681baedc9167ff4827ec5d8839c980010029`

// DeployRenExBalances deploys a new Ethereum contract, binding an instance of RenExBalances to it.
func DeployRenExBalances(auth *bind.TransactOpts, backend bind.ContractBackend, _rewardVaultContract common.Address) (common.Address, *types.Transaction, *RenExBalances, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExBalancesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RenExBalancesBin), backend, _rewardVaultContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RenExBalances{RenExBalancesCaller: RenExBalancesCaller{contract: contract}, RenExBalancesTransactor: RenExBalancesTransactor{contract: contract}, RenExBalancesFilterer: RenExBalancesFilterer{contract: contract}}, nil
}

// RenExBalances is an auto generated Go binding around an Ethereum contract.
type RenExBalances struct {
	RenExBalancesCaller     // Read-only binding to the contract
	RenExBalancesTransactor // Write-only binding to the contract
	RenExBalancesFilterer   // Log filterer for contract events
}

// RenExBalancesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenExBalancesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBalancesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenExBalancesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBalancesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenExBalancesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExBalancesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenExBalancesSession struct {
	Contract     *RenExBalances    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RenExBalancesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenExBalancesCallerSession struct {
	Contract *RenExBalancesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RenExBalancesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenExBalancesTransactorSession struct {
	Contract     *RenExBalancesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RenExBalancesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenExBalancesRaw struct {
	Contract *RenExBalances // Generic contract binding to access the raw methods on
}

// RenExBalancesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenExBalancesCallerRaw struct {
	Contract *RenExBalancesCaller // Generic read-only contract binding to access the raw methods on
}

// RenExBalancesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenExBalancesTransactorRaw struct {
	Contract *RenExBalancesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRenExBalances creates a new instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalances(address common.Address, backend bind.ContractBackend) (*RenExBalances, error) {
	contract, err := bindRenExBalances(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RenExBalances{RenExBalancesCaller: RenExBalancesCaller{contract: contract}, RenExBalancesTransactor: RenExBalancesTransactor{contract: contract}, RenExBalancesFilterer: RenExBalancesFilterer{contract: contract}}, nil
}

// NewRenExBalancesCaller creates a new read-only instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalancesCaller(address common.Address, caller bind.ContractCaller) (*RenExBalancesCaller, error) {
	contract, err := bindRenExBalances(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesCaller{contract: contract}, nil
}

// NewRenExBalancesTransactor creates a new write-only instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalancesTransactor(address common.Address, transactor bind.ContractTransactor) (*RenExBalancesTransactor, error) {
	contract, err := bindRenExBalances(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesTransactor{contract: contract}, nil
}

// NewRenExBalancesFilterer creates a new log filterer instance of RenExBalances, bound to a specific deployed contract.
func NewRenExBalancesFilterer(address common.Address, filterer bind.ContractFilterer) (*RenExBalancesFilterer, error) {
	contract, err := bindRenExBalances(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesFilterer{contract: contract}, nil
}

// bindRenExBalances binds a generic wrapper to an already deployed contract.
func bindRenExBalances(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExBalancesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExBalances *RenExBalancesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExBalances.Contract.RenExBalancesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExBalances *RenExBalancesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBalances.Contract.RenExBalancesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExBalances *RenExBalancesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExBalances.Contract.RenExBalancesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExBalances *RenExBalancesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExBalances.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExBalances *RenExBalancesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBalances.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExBalances *RenExBalancesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExBalances.Contract.contract.Transact(opts, method, params...)
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_RenExBalances *RenExBalancesCaller) ETHEREUM(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "ETHEREUM")
	return *ret0, err
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_RenExBalances *RenExBalancesSession) ETHEREUM() (common.Address, error) {
	return _RenExBalances.Contract.ETHEREUM(&_RenExBalances.CallOpts)
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_RenExBalances *RenExBalancesCallerSession) ETHEREUM() (common.Address, error) {
	return _RenExBalances.Contract.ETHEREUM(&_RenExBalances.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0xc84aae17.
//
// Solidity: function getBalances(_trader address) constant returns(address[], uint256[])
func (_RenExBalances *RenExBalancesCaller) GetBalances(opts *bind.CallOpts, _trader common.Address) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RenExBalances.contract.Call(opts, out, "getBalances", _trader)
	return *ret0, *ret1, err
}

// GetBalances is a free data retrieval call binding the contract method 0xc84aae17.
//
// Solidity: function getBalances(_trader address) constant returns(address[], uint256[])
func (_RenExBalances *RenExBalancesSession) GetBalances(_trader common.Address) ([]common.Address, []*big.Int, error) {
	return _RenExBalances.Contract.GetBalances(&_RenExBalances.CallOpts, _trader)
}

// GetBalances is a free data retrieval call binding the contract method 0xc84aae17.
//
// Solidity: function getBalances(_trader address) constant returns(address[], uint256[])
func (_RenExBalances *RenExBalancesCallerSession) GetBalances(_trader common.Address) ([]common.Address, []*big.Int, error) {
	return _RenExBalances.Contract.GetBalances(&_RenExBalances.CallOpts, _trader)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBalances *RenExBalancesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBalances *RenExBalancesSession) Owner() (common.Address, error) {
	return _RenExBalances.Contract.Owner(&_RenExBalances.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExBalances *RenExBalancesCallerSession) Owner() (common.Address, error) {
	return _RenExBalances.Contract.Owner(&_RenExBalances.CallOpts)
}

// RewardVaultContract is a free data retrieval call binding the contract method 0x23017a3a.
//
// Solidity: function rewardVaultContract() constant returns(address)
func (_RenExBalances *RenExBalancesCaller) RewardVaultContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "rewardVaultContract")
	return *ret0, err
}

// RewardVaultContract is a free data retrieval call binding the contract method 0x23017a3a.
//
// Solidity: function rewardVaultContract() constant returns(address)
func (_RenExBalances *RenExBalancesSession) RewardVaultContract() (common.Address, error) {
	return _RenExBalances.Contract.RewardVaultContract(&_RenExBalances.CallOpts)
}

// RewardVaultContract is a free data retrieval call binding the contract method 0x23017a3a.
//
// Solidity: function rewardVaultContract() constant returns(address)
func (_RenExBalances *RenExBalancesCallerSession) RewardVaultContract() (common.Address, error) {
	return _RenExBalances.Contract.RewardVaultContract(&_RenExBalances.CallOpts)
}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() constant returns(address)
func (_RenExBalances *RenExBalancesCaller) SettlementContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "settlementContract")
	return *ret0, err
}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() constant returns(address)
func (_RenExBalances *RenExBalancesSession) SettlementContract() (common.Address, error) {
	return _RenExBalances.Contract.SettlementContract(&_RenExBalances.CallOpts)
}

// SettlementContract is a free data retrieval call binding the contract method 0xea42418b.
//
// Solidity: function settlementContract() constant returns(address)
func (_RenExBalances *RenExBalancesCallerSession) SettlementContract() (common.Address, error) {
	return _RenExBalances.Contract.SettlementContract(&_RenExBalances.CallOpts)
}

// TraderBalances is a free data retrieval call binding the contract method 0xc43c633b.
//
// Solidity: function traderBalances( address,  address) constant returns(uint256)
func (_RenExBalances *RenExBalancesCaller) TraderBalances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "traderBalances", arg0, arg1)
	return *ret0, err
}

// TraderBalances is a free data retrieval call binding the contract method 0xc43c633b.
//
// Solidity: function traderBalances( address,  address) constant returns(uint256)
func (_RenExBalances *RenExBalancesSession) TraderBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RenExBalances.Contract.TraderBalances(&_RenExBalances.CallOpts, arg0, arg1)
}

// TraderBalances is a free data retrieval call binding the contract method 0xc43c633b.
//
// Solidity: function traderBalances( address,  address) constant returns(uint256)
func (_RenExBalances *RenExBalancesCallerSession) TraderBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RenExBalances.Contract.TraderBalances(&_RenExBalances.CallOpts, arg0, arg1)
}

// TraderTokens is a free data retrieval call binding the contract method 0x1a6947ce.
//
// Solidity: function traderTokens( address,  uint256) constant returns(address)
func (_RenExBalances *RenExBalancesCaller) TraderTokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExBalances.contract.Call(opts, out, "traderTokens", arg0, arg1)
	return *ret0, err
}

// TraderTokens is a free data retrieval call binding the contract method 0x1a6947ce.
//
// Solidity: function traderTokens( address,  uint256) constant returns(address)
func (_RenExBalances *RenExBalancesSession) TraderTokens(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _RenExBalances.Contract.TraderTokens(&_RenExBalances.CallOpts, arg0, arg1)
}

// TraderTokens is a free data retrieval call binding the contract method 0x1a6947ce.
//
// Solidity: function traderTokens( address,  uint256) constant returns(address)
func (_RenExBalances *RenExBalancesCallerSession) TraderTokens(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _RenExBalances.Contract.TraderTokens(&_RenExBalances.CallOpts, arg0, arg1)
}

// DecrementBalanceWithFee is a paid mutator transaction binding the contract method 0x92dd79e6.
//
// Solidity: function decrementBalanceWithFee(_trader address, _token address, _value uint256, _fee uint256, feePayee address) returns()
func (_RenExBalances *RenExBalancesTransactor) DecrementBalanceWithFee(opts *bind.TransactOpts, _trader common.Address, _token common.Address, _value *big.Int, _fee *big.Int, feePayee common.Address) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "decrementBalanceWithFee", _trader, _token, _value, _fee, feePayee)
}

// DecrementBalanceWithFee is a paid mutator transaction binding the contract method 0x92dd79e6.
//
// Solidity: function decrementBalanceWithFee(_trader address, _token address, _value uint256, _fee uint256, feePayee address) returns()
func (_RenExBalances *RenExBalancesSession) DecrementBalanceWithFee(_trader common.Address, _token common.Address, _value *big.Int, _fee *big.Int, feePayee common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.DecrementBalanceWithFee(&_RenExBalances.TransactOpts, _trader, _token, _value, _fee, feePayee)
}

// DecrementBalanceWithFee is a paid mutator transaction binding the contract method 0x92dd79e6.
//
// Solidity: function decrementBalanceWithFee(_trader address, _token address, _value uint256, _fee uint256, feePayee address) returns()
func (_RenExBalances *RenExBalancesTransactorSession) DecrementBalanceWithFee(_trader common.Address, _token common.Address, _value *big.Int, _fee *big.Int, feePayee common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.DecrementBalanceWithFee(&_RenExBalances.TransactOpts, _trader, _token, _value, _fee, feePayee)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "deposit", _token, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesSession) Deposit(_token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.Contract.Deposit(&_RenExBalances.TransactOpts, _token, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(_token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesTransactorSession) Deposit(_token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.Contract.Deposit(&_RenExBalances.TransactOpts, _token, _value)
}

// IncrementBalance is a paid mutator transaction binding the contract method 0x2b15e857.
//
// Solidity: function incrementBalance(_trader address, _token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesTransactor) IncrementBalance(opts *bind.TransactOpts, _trader common.Address, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "incrementBalance", _trader, _token, _value)
}

// IncrementBalance is a paid mutator transaction binding the contract method 0x2b15e857.
//
// Solidity: function incrementBalance(_trader address, _token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesSession) IncrementBalance(_trader common.Address, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.Contract.IncrementBalance(&_RenExBalances.TransactOpts, _trader, _token, _value)
}

// IncrementBalance is a paid mutator transaction binding the contract method 0x2b15e857.
//
// Solidity: function incrementBalance(_trader address, _token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesTransactorSession) IncrementBalance(_trader common.Address, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.Contract.IncrementBalance(&_RenExBalances.TransactOpts, _trader, _token, _value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBalances *RenExBalancesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBalances *RenExBalancesSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExBalances.Contract.RenounceOwnership(&_RenExBalances.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExBalances *RenExBalancesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExBalances.Contract.RenounceOwnership(&_RenExBalances.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBalances *RenExBalancesTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBalances *RenExBalancesSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.TransferOwnership(&_RenExBalances.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExBalances *RenExBalancesTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.TransferOwnership(&_RenExBalances.TransactOpts, _newOwner)
}

// UpdateRenExSettlementContract is a paid mutator transaction binding the contract method 0x1d65551d.
//
// Solidity: function updateRenExSettlementContract(_newSettlementContract address) returns()
func (_RenExBalances *RenExBalancesTransactor) UpdateRenExSettlementContract(opts *bind.TransactOpts, _newSettlementContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "updateRenExSettlementContract", _newSettlementContract)
}

// UpdateRenExSettlementContract is a paid mutator transaction binding the contract method 0x1d65551d.
//
// Solidity: function updateRenExSettlementContract(_newSettlementContract address) returns()
func (_RenExBalances *RenExBalancesSession) UpdateRenExSettlementContract(_newSettlementContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.UpdateRenExSettlementContract(&_RenExBalances.TransactOpts, _newSettlementContract)
}

// UpdateRenExSettlementContract is a paid mutator transaction binding the contract method 0x1d65551d.
//
// Solidity: function updateRenExSettlementContract(_newSettlementContract address) returns()
func (_RenExBalances *RenExBalancesTransactorSession) UpdateRenExSettlementContract(_newSettlementContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.UpdateRenExSettlementContract(&_RenExBalances.TransactOpts, _newSettlementContract)
}

// UpdateRewardVault is a paid mutator transaction binding the contract method 0xde10f619.
//
// Solidity: function updateRewardVault(_newRewardVaultContract address) returns()
func (_RenExBalances *RenExBalancesTransactor) UpdateRewardVault(opts *bind.TransactOpts, _newRewardVaultContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "updateRewardVault", _newRewardVaultContract)
}

// UpdateRewardVault is a paid mutator transaction binding the contract method 0xde10f619.
//
// Solidity: function updateRewardVault(_newRewardVaultContract address) returns()
func (_RenExBalances *RenExBalancesSession) UpdateRewardVault(_newRewardVaultContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.UpdateRewardVault(&_RenExBalances.TransactOpts, _newRewardVaultContract)
}

// UpdateRewardVault is a paid mutator transaction binding the contract method 0xde10f619.
//
// Solidity: function updateRewardVault(_newRewardVaultContract address) returns()
func (_RenExBalances *RenExBalancesTransactorSession) UpdateRewardVault(_newRewardVaultContract common.Address) (*types.Transaction, error) {
	return _RenExBalances.Contract.UpdateRewardVault(&_RenExBalances.TransactOpts, _newRewardVaultContract)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(_token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.contract.Transact(opts, "withdraw", _token, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(_token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesSession) Withdraw(_token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.Contract.Withdraw(&_RenExBalances.TransactOpts, _token, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(_token address, _value uint256) returns()
func (_RenExBalances *RenExBalancesTransactorSession) Withdraw(_token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RenExBalances.Contract.Withdraw(&_RenExBalances.TransactOpts, _token, _value)
}

// RenExBalancesBalanceDecreasedIterator is returned from FilterBalanceDecreased and is used to iterate over the raw logs and unpacked data for BalanceDecreased events raised by the RenExBalances contract.
type RenExBalancesBalanceDecreasedIterator struct {
	Event *RenExBalancesBalanceDecreased // Event containing the contract specifics and raw log

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
func (it *RenExBalancesBalanceDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesBalanceDecreased)
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
		it.Event = new(RenExBalancesBalanceDecreased)
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
func (it *RenExBalancesBalanceDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesBalanceDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesBalanceDecreased represents a BalanceDecreased event raised by the RenExBalances contract.
type RenExBalancesBalanceDecreased struct {
	Trader common.Address
	Token  common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBalanceDecreased is a free log retrieval operation binding the contract event 0x6ac2cd906088d873624fa62ca95170d967629e7d964651df19a3aa2e49b44aa1.
//
// Solidity: e BalanceDecreased(trader address, token address, value uint256)
func (_RenExBalances *RenExBalancesFilterer) FilterBalanceDecreased(opts *bind.FilterOpts) (*RenExBalancesBalanceDecreasedIterator, error) {

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "BalanceDecreased")
	if err != nil {
		return nil, err
	}
	return &RenExBalancesBalanceDecreasedIterator{contract: _RenExBalances.contract, event: "BalanceDecreased", logs: logs, sub: sub}, nil
}

// WatchBalanceDecreased is a free log subscription operation binding the contract event 0x6ac2cd906088d873624fa62ca95170d967629e7d964651df19a3aa2e49b44aa1.
//
// Solidity: e BalanceDecreased(trader address, token address, value uint256)
func (_RenExBalances *RenExBalancesFilterer) WatchBalanceDecreased(opts *bind.WatchOpts, sink chan<- *RenExBalancesBalanceDecreased) (event.Subscription, error) {

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "BalanceDecreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesBalanceDecreased)
				if err := _RenExBalances.contract.UnpackLog(event, "BalanceDecreased", log); err != nil {
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

// RenExBalancesBalanceIncreasedIterator is returned from FilterBalanceIncreased and is used to iterate over the raw logs and unpacked data for BalanceIncreased events raised by the RenExBalances contract.
type RenExBalancesBalanceIncreasedIterator struct {
	Event *RenExBalancesBalanceIncreased // Event containing the contract specifics and raw log

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
func (it *RenExBalancesBalanceIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesBalanceIncreased)
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
		it.Event = new(RenExBalancesBalanceIncreased)
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
func (it *RenExBalancesBalanceIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesBalanceIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesBalanceIncreased represents a BalanceIncreased event raised by the RenExBalances contract.
type RenExBalancesBalanceIncreased struct {
	Trader common.Address
	Token  common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBalanceIncreased is a free log retrieval operation binding the contract event 0x0d66f59c9991adc17dd3339490c5058d2d6fe20395e7b55ceb6ca8019a31667d.
//
// Solidity: e BalanceIncreased(trader address, token address, value uint256)
func (_RenExBalances *RenExBalancesFilterer) FilterBalanceIncreased(opts *bind.FilterOpts) (*RenExBalancesBalanceIncreasedIterator, error) {

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "BalanceIncreased")
	if err != nil {
		return nil, err
	}
	return &RenExBalancesBalanceIncreasedIterator{contract: _RenExBalances.contract, event: "BalanceIncreased", logs: logs, sub: sub}, nil
}

// WatchBalanceIncreased is a free log subscription operation binding the contract event 0x0d66f59c9991adc17dd3339490c5058d2d6fe20395e7b55ceb6ca8019a31667d.
//
// Solidity: e BalanceIncreased(trader address, token address, value uint256)
func (_RenExBalances *RenExBalancesFilterer) WatchBalanceIncreased(opts *bind.WatchOpts, sink chan<- *RenExBalancesBalanceIncreased) (event.Subscription, error) {

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "BalanceIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesBalanceIncreased)
				if err := _RenExBalances.contract.UnpackLog(event, "BalanceIncreased", log); err != nil {
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

// RenExBalancesOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RenExBalances contract.
type RenExBalancesOwnershipRenouncedIterator struct {
	Event *RenExBalancesOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RenExBalancesOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesOwnershipRenounced)
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
		it.Event = new(RenExBalancesOwnershipRenounced)
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
func (it *RenExBalancesOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesOwnershipRenounced represents a OwnershipRenounced event raised by the RenExBalances contract.
type RenExBalancesOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExBalances *RenExBalancesFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RenExBalancesOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesOwnershipRenouncedIterator{contract: _RenExBalances.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExBalances *RenExBalancesFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RenExBalancesOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesOwnershipRenounced)
				if err := _RenExBalances.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RenExBalancesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RenExBalances contract.
type RenExBalancesOwnershipTransferredIterator struct {
	Event *RenExBalancesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RenExBalancesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesOwnershipTransferred)
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
		it.Event = new(RenExBalancesOwnershipTransferred)
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
func (it *RenExBalancesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesOwnershipTransferred represents a OwnershipTransferred event raised by the RenExBalances contract.
type RenExBalancesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExBalances *RenExBalancesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RenExBalancesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesOwnershipTransferredIterator{contract: _RenExBalances.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExBalances *RenExBalancesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RenExBalancesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesOwnershipTransferred)
				if err := _RenExBalances.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RenExBalancesRenExSettlementContractUpdatedIterator is returned from FilterRenExSettlementContractUpdated and is used to iterate over the raw logs and unpacked data for RenExSettlementContractUpdated events raised by the RenExBalances contract.
type RenExBalancesRenExSettlementContractUpdatedIterator struct {
	Event *RenExBalancesRenExSettlementContractUpdated // Event containing the contract specifics and raw log

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
func (it *RenExBalancesRenExSettlementContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesRenExSettlementContractUpdated)
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
		it.Event = new(RenExBalancesRenExSettlementContractUpdated)
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
func (it *RenExBalancesRenExSettlementContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesRenExSettlementContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesRenExSettlementContractUpdated represents a RenExSettlementContractUpdated event raised by the RenExBalances contract.
type RenExBalancesRenExSettlementContractUpdated struct {
	NewRenExSettlementContract common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterRenExSettlementContractUpdated is a free log retrieval operation binding the contract event 0x13ea4b9a37263b8dd3ea2fbd9c274a8845489805d06e98f12e40d216fa025071.
//
// Solidity: e RenExSettlementContractUpdated(newRenExSettlementContract indexed address)
func (_RenExBalances *RenExBalancesFilterer) FilterRenExSettlementContractUpdated(opts *bind.FilterOpts, newRenExSettlementContract []common.Address) (*RenExBalancesRenExSettlementContractUpdatedIterator, error) {

	var newRenExSettlementContractRule []interface{}
	for _, newRenExSettlementContractItem := range newRenExSettlementContract {
		newRenExSettlementContractRule = append(newRenExSettlementContractRule, newRenExSettlementContractItem)
	}

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "RenExSettlementContractUpdated", newRenExSettlementContractRule)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesRenExSettlementContractUpdatedIterator{contract: _RenExBalances.contract, event: "RenExSettlementContractUpdated", logs: logs, sub: sub}, nil
}

// WatchRenExSettlementContractUpdated is a free log subscription operation binding the contract event 0x13ea4b9a37263b8dd3ea2fbd9c274a8845489805d06e98f12e40d216fa025071.
//
// Solidity: e RenExSettlementContractUpdated(newRenExSettlementContract indexed address)
func (_RenExBalances *RenExBalancesFilterer) WatchRenExSettlementContractUpdated(opts *bind.WatchOpts, sink chan<- *RenExBalancesRenExSettlementContractUpdated, newRenExSettlementContract []common.Address) (event.Subscription, error) {

	var newRenExSettlementContractRule []interface{}
	for _, newRenExSettlementContractItem := range newRenExSettlementContract {
		newRenExSettlementContractRule = append(newRenExSettlementContractRule, newRenExSettlementContractItem)
	}

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "RenExSettlementContractUpdated", newRenExSettlementContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesRenExSettlementContractUpdated)
				if err := _RenExBalances.contract.UnpackLog(event, "RenExSettlementContractUpdated", log); err != nil {
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

// RenExBalancesRewardVaultContractUpdatedIterator is returned from FilterRewardVaultContractUpdated and is used to iterate over the raw logs and unpacked data for RewardVaultContractUpdated events raised by the RenExBalances contract.
type RenExBalancesRewardVaultContractUpdatedIterator struct {
	Event *RenExBalancesRewardVaultContractUpdated // Event containing the contract specifics and raw log

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
func (it *RenExBalancesRewardVaultContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExBalancesRewardVaultContractUpdated)
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
		it.Event = new(RenExBalancesRewardVaultContractUpdated)
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
func (it *RenExBalancesRewardVaultContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExBalancesRewardVaultContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExBalancesRewardVaultContractUpdated represents a RewardVaultContractUpdated event raised by the RenExBalances contract.
type RenExBalancesRewardVaultContractUpdated struct {
	NewRewardVaultContract common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRewardVaultContractUpdated is a free log retrieval operation binding the contract event 0xddbdde9ff1bc1b3e2779975798b74bbbdf2b31c326d86cfa200e381843afdd1d.
//
// Solidity: e RewardVaultContractUpdated(newRewardVaultContract indexed address)
func (_RenExBalances *RenExBalancesFilterer) FilterRewardVaultContractUpdated(opts *bind.FilterOpts, newRewardVaultContract []common.Address) (*RenExBalancesRewardVaultContractUpdatedIterator, error) {

	var newRewardVaultContractRule []interface{}
	for _, newRewardVaultContractItem := range newRewardVaultContract {
		newRewardVaultContractRule = append(newRewardVaultContractRule, newRewardVaultContractItem)
	}

	logs, sub, err := _RenExBalances.contract.FilterLogs(opts, "RewardVaultContractUpdated", newRewardVaultContractRule)
	if err != nil {
		return nil, err
	}
	return &RenExBalancesRewardVaultContractUpdatedIterator{contract: _RenExBalances.contract, event: "RewardVaultContractUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardVaultContractUpdated is a free log subscription operation binding the contract event 0xddbdde9ff1bc1b3e2779975798b74bbbdf2b31c326d86cfa200e381843afdd1d.
//
// Solidity: e RewardVaultContractUpdated(newRewardVaultContract indexed address)
func (_RenExBalances *RenExBalancesFilterer) WatchRewardVaultContractUpdated(opts *bind.WatchOpts, sink chan<- *RenExBalancesRewardVaultContractUpdated, newRewardVaultContract []common.Address) (event.Subscription, error) {

	var newRewardVaultContractRule []interface{}
	for _, newRewardVaultContractItem := range newRewardVaultContract {
		newRewardVaultContractRule = append(newRewardVaultContractRule, newRewardVaultContractItem)
	}

	logs, sub, err := _RenExBalances.contract.WatchLogs(opts, "RewardVaultContractUpdated", newRewardVaultContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExBalancesRewardVaultContractUpdated)
				if err := _RenExBalances.contract.UnpackLog(event, "RewardVaultContractUpdated", log); err != nil {
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

// RenExSettlementABI is the input ABI used to generate the binding from.
const RenExSettlementABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"SETTLEMENT_IDENTIFIER\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orderID\",\"type\":\"bytes32\"}],\"name\":\"verifyOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_buyID\",\"type\":\"bytes32\"},{\"name\":\"_sellID\",\"type\":\"bytes32\"}],\"name\":\"verifyMatch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_buyID\",\"type\":\"bytes32\"},{\"name\":\"_sellID\",\"type\":\"bytes32\"}],\"name\":\"submitMatch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orderType\",\"type\":\"uint8\"},{\"name\":\"_parity\",\"type\":\"uint8\"},{\"name\":\"_expiry\",\"type\":\"uint64\"},{\"name\":\"_tokens\",\"type\":\"uint64\"},{\"name\":\"_priceC\",\"type\":\"uint16\"},{\"name\":\"_priceQ\",\"type\":\"uint16\"},{\"name\":\"_volumeC\",\"type\":\"uint16\"},{\"name\":\"_volumeQ\",\"type\":\"uint16\"},{\"name\":\"_minimumVolumeC\",\"type\":\"uint16\"},{\"name\":\"_minimumVolumeQ\",\"type\":\"uint16\"},{\"name\":\"_nonceHash\",\"type\":\"uint256\"}],\"name\":\"submitOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"renExTokensContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submissionGasPriceLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOrderbookContract\",\"type\":\"address\"}],\"name\":\"updateOrderbook\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSubmissionGasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"updateSubmissionGasPriceLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEES_NUMERATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_buyID\",\"type\":\"bytes32\"},{\"name\":\"_sellID\",\"type\":\"bytes32\"}],\"name\":\"getMidPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orders\",\"outputs\":[{\"name\":\"parity\",\"type\":\"uint8\"},{\"name\":\"orderType\",\"type\":\"uint8\"},{\"name\":\"expiry\",\"type\":\"uint64\"},{\"name\":\"tokens\",\"type\":\"uint64\"},{\"name\":\"priceC\",\"type\":\"uint64\"},{\"name\":\"priceQ\",\"type\":\"uint64\"},{\"name\":\"volumeC\",\"type\":\"uint64\"},{\"name\":\"volumeQ\",\"type\":\"uint64\"},{\"name\":\"minimumVolumeC\",\"type\":\"uint64\"},{\"name\":\"minimumVolumeQ\",\"type\":\"uint64\"},{\"name\":\"nonceHash\",\"type\":\"uint256\"},{\"name\":\"trader\",\"type\":\"address\"},{\"name\":\"submitter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_trader\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"traderCanWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEES_DENOMINATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"orderbookContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"renExBalancesContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_buyID\",\"type\":\"bytes32\"},{\"name\":\"_sellID\",\"type\":\"bytes32\"}],\"name\":\"getSettlementDetails\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRenExBalancesContract\",\"type\":\"address\"}],\"name\":\"updateRenExBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_orderbookContract\",\"type\":\"address\"},{\"name\":\"_renExTokensContract\",\"type\":\"address\"},{\"name\":\"_renExBalancesContract\",\"type\":\"address\"},{\"name\":\"_submissionGasPriceLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousOrderbook\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextOrderbook\",\"type\":\"address\"}],\"name\":\"OrderbookUpdates\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousRenExBalances\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextRenExBalances\",\"type\":\"address\"}],\"name\":\"RenExBalancesUpdates\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousSubmissionGasPriceLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nextSubmissionGasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"SubmissionGasPriceLimitUpdates\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RenExSettlementBin is the compiled bytecode used for deploying new contracts.
const RenExSettlementBin = `0x608060405234801561001057600080fd5b506040516080806120e1833981016040908152815160208301519183015160609093015160008054600160a060020a0319908116331790915560018054600160a060020a03948516908316179055600280549484169482169490941790935560038054929094169190921617909155600455612050806100916000396000f3006080604052600436106101035763ffffffff60e060020a60003504166307067ef08114610108578063169fb328146101365780631f88d9c7146101505780632a337d301461018e578063367ba52c146101a95780634015e83b146102085780634c3052de146102395780636074b80614610260578063675df16f14610281578063715018a6146102995780637a089c26146102ae5780638da5cb5b146102c35780638fac5ff5146102d85780639c3f1e90146102f3578063a3bdaedc14610396578063bae119ef146103d4578063bbe7221e146103e9578063ddf25ce9146103fe578063eb534ea014610413578063ee0715ed14610459578063f2fde38b1461047a575b600080fd5b34801561011457600080fd5b5061011d61049b565b6040805163ffffffff9092168252519081900360200190f35b34801561014257600080fd5b5061014e6004356104a0565b005b34801561015c57600080fd5b5061016b6004356024356105f8565b6040805163ffffffff938416815291909216602082015281519081900390910190f35b34801561019a57600080fd5b5061014e600435602435610a05565b3480156101b557600080fd5b5061014e60ff6004358116906024351667ffffffffffffffff6044358116906064351661ffff60843581169060a43581169060c43581169060e43581169061010435811690610124351661014435610a7c565b34801561021457600080fd5b5061021d610f74565b60408051600160a060020a039092168252519081900360200190f35b34801561024557600080fd5b5061024e610f83565b60408051918252519081900360200190f35b34801561026c57600080fd5b5061014e600160a060020a0360043516610f89565b34801561028d57600080fd5b5061014e600435611017565b3480156102a557600080fd5b5061014e611070565b3480156102ba57600080fd5b5061024e6110dc565b3480156102cf57600080fd5b5061021d6110e1565b3480156102e457600080fd5b5061024e6004356024356110f0565b3480156102ff57600080fd5b5061030b6004356111ec565b6040805160ff9e8f1681529c909d1660208d015267ffffffffffffffff9a8b168c8e0152988a1660608c015296891660808b015294881660a08a015292871660c089015290861660e08801528516610100870152909316610120850152610140840192909252600160a060020a03918216610160840152166101808201529051908190036101a00190f35b3480156103a257600080fd5b506103c0600160a060020a03600435811690602435166044356112a1565b604080519115158252519081900360200190f35b3480156103e057600080fd5b5061024e6112aa565b3480156103f557600080fd5b5061021d6112b0565b34801561040a57600080fd5b5061021d6112bf565b34801561041f57600080fd5b5061042e6004356024356112ce565b6040805195865260208601949094528484019290925260608401526080830152519081900360a00190f35b34801561046557600080fd5b5061014e600160a060020a0360043516611356565b34801561048657600080fd5b5061014e600160a060020a03600435166113e4565b600181565b600160008281526006602052604090205460ff1660028111156104bf57fe5b146104c957600080fd5b6000818152600560205260409020546107cf720100000000000000000000000000000000000090910467ffffffffffffffff16111561050757600080fd5b600081815260056020526040902060010154603467ffffffffffffffff909116111561053257600080fd5b60008181526005602052604090206001015460316801000000000000000090910467ffffffffffffffff16111561056857600080fd5b6000818152600560205260409020600101546034608060020a90910467ffffffffffffffff16111561059957600080fd5b600081815260056020526040902060010154603160c060020a90910467ffffffffffffffff1611156105ca57600080fd5b600081815260056020526040902060020154603467ffffffffffffffff90911611156105f557600080fd5b50565b60008281526005602052604081205481908190819060ff161561061a57600080fd5b60008581526005602052604090205460ff1660011461063857600080fd5b600154604080517faf3e8a400000000000000000000000000000000000000000000000000000000081526004810189905290518792600160a060020a03169163af3e8a4091602480830192600092919082900301818387803b15801561069d57600080fd5b505af11580156106b1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156106da57600080fd5b8101908080516401000000008111156106f257600080fd5b8201602081018481111561070557600080fd5b815185602082028301116401000000008211171561072257600080fd5b50508051909350600092508210905061073757fe5b602090810290910101511461074b57600080fd5b60008681526005602052604080822080546001918201548985529290932080549101546107a69367ffffffffffffffff7201000000000000000000000000000000000000918290048116948116939190910481169116611404565b15156107b157600080fd5b600086815260056020526040808220600190810154888452919092209182015460029092015461080e9267ffffffffffffffff680100000000000000008404811693608060020a900481169260c060020a90920481169116611404565b151561081957600080fd5b60008581526005602052604080822060019081015489845291909220918201546002909201546108769267ffffffffffffffff680100000000000000008404811693608060020a900481169260c060020a90920481169116611404565b151561088157600080fd5b505060008381526005602090815260408083205460025482517f329ed3810000000000000000000000000000000000000000000000000000000081526a010000000000000000000090920463ffffffff81166004840152925167ffffffffffffffff90931694640100000000860494600160a060020a039092169363329ed3819360248082019493918390030190829087803b15801561092057600080fd5b505af1158015610934573d6000803e3d6000fd5b505050506040513d602081101561094a57600080fd5b5051151561095757600080fd5b600254604080517f329ed38100000000000000000000000000000000000000000000000000000000815263ffffffff841660048201529051600160a060020a039092169163329ed381916024808201926020929091908290030181600087803b1580156109c357600080fd5b505af11580156109d7573d6000803e3d6000fd5b505050506040513d60208110156109ed57600080fd5b505115156109fa57600080fd5b909590945092505050565b600080600080610a14866104a0565b610a1d856104a0565b610a2786866105f8565b93509350610a3786868686611464565b91509150610a49868686868686611618565b505050600092835250600660205260408083208054600260ff199182168117909255928452922080549091169091179055565b610a84611fb8565b6000600454803a11151515610a9857600080fd5b6101a0604051908101604052808e60ff1681526020018f60ff1681526020018d67ffffffffffffffff1681526020018c67ffffffffffffffff1681526020018b61ffff1667ffffffffffffffff1681526020018a61ffff1667ffffffffffffffff1681526020018961ffff1667ffffffffffffffff1681526020018861ffff1667ffffffffffffffff1681526020018761ffff1667ffffffffffffffff1681526020018661ffff1667ffffffffffffffff1681526020018581526020016000600160a060020a0316815260200133600160a060020a03168152509250610b7d83611a45565b91506000808381526006602052604090205460ff166002811115610b9d57fe5b14610ba757600080fd5b6000828152600660209081526040808320805460ff191660019081179091555481517faab14d04000000000000000000000000000000000000000000000000000000008152600481018790529151600160a060020a039091169363aab14d0493602480850194919392918390030190829087803b158015610c2757600080fd5b505af1158015610c3b573d6000803e3d6000fd5b505050506040513d6020811015610c5157600080fd5b505160ff16600214610c6257600080fd5b600154604080517fb1a08010000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a039092169163b1a08010916024808201926020929091908290030181600087803b158015610cc957600080fd5b505af1158015610cdd573d6000803e3d6000fd5b505050506040513d6020811015610cf357600080fd5b8101908080519060200190929190505050836101600190600160a060020a03169081600160a060020a0316815250508260056000846000191660001916815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548160ff021916908360ff16021790555060408201518160000160026101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550606082015181600001600a6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060808201518160000160126101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060a08201518160010160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060c08201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060e08201518160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101008201518160010160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101208201518160020160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061014082015181600301556101608201518160040160006101000a815481600160a060020a030219169083600160a060020a031602179055506101808201518160050160006101000a815481600160a060020a030219169083600160a060020a031602179055509050505050505050505050505050505050565b600254600160a060020a031681565b60045481565b600054600160a060020a03163314610fa057600080fd5b60015460408051600160a060020a039283168152918316602083015280517f05565154275c40146888f6e69d9557d712c9384838de11284a6742df83e0e01b9281900390910190a16001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461102e57600080fd5b600454604080519182526020820183905280517fccf294e3e783bf1a9ff94493761c96e0ccbc809312bcb5376beaa9baa31de3249281900390910190a1600455565b600054600160a060020a0316331461108757600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600281565b600054600160a060020a031681565b60008060008060006111028787611c82565b60008881526005602090815260408083205460025482517fc865992b0000000000000000000000000000000000000000000000000000000081526401000000006a010000000000000000000090930467ffffffffffffffff169290920463ffffffff811660048401529251969a50949850909650600160a060020a039093169363c865992b93602480820194918390030190829087803b1580156111a557600080fd5b505af11580156111b9573d6000803e3d6000fd5b505050506040513d60208110156111cf57600080fd5b505160ff1690506111e1848483611d12565b979650505050505050565b600560208190526000918252604090912080546001820154600283015460038401546004850154949095015460ff8085169661010086049091169567ffffffffffffffff6201000087048116966a01000000000000000000008104821696720100000000000000000000000000000000000090910482169581831695680100000000000000008304841695608060020a840485169560c060020a9094048516949290921692600160a060020a0390811691168d565b60019392505050565b6103e881565b600154600160a060020a031681565b600354600160a060020a031681565b60008181526005602052604081205481908190819081906a0100000000000000000000900467ffffffffffffffff16640100000000810482808080806113168e8e8989611464565b90955093506103e86103e686020492506103e86103e6850204915061133b8e8e6110f0565b9e929d50909b5050918b900398508990039650945050505050565b600054600160a060020a0316331461136d57600080fd5b60035460408051600160a060020a039283168152918316602083015280517f1370a598dc4f33911a2368195ac8e693faec43f2ce450cd909dcd64d590cd2229281900390910190a16003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031633146113fb57600080fd5b6105f581611d53565b6000808267ffffffffffffffff168567ffffffffffffffff16101561142c576000915061145b565b82850367ffffffffffffffff16600a0a8667ffffffffffffffff160290508367ffffffffffffffff1681101591505b50949350505050565b6000806000806000806000806000806000600260009054906101000a9004600160a060020a0316600160a060020a031663c865992b8e6040518263ffffffff1660e060020a028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b1580156114df57600080fd5b505af11580156114f3573d6000803e3d6000fd5b505050506040513d602081101561150957600080fd5b810190808051906020019092919050505060ff169850600260009054906101000a9004600160a060020a0316600160a060020a031663c865992b8d6040518263ffffffff1660e060020a028152600401808263ffffffff1663ffffffff168152602001915050602060405180830381600087803b15801561158957600080fd5b505af115801561159d573d6000803e3d6000fd5b505050506040513d60208110156115b357600080fd5b505160ff1697506115c48f8f611c82565b965096506115d48f8f8989611dd0565b9450945094506115ee85858989878d63ffffffff16611efd565b91506116028585858c63ffffffff16611f52565b919f919e50909c50505050505050505050505050565b600254604080517f29debc9f00000000000000000000000000000000000000000000000000000000815263ffffffff87166004820152905160009283928392839283928392600160a060020a03909216916329debc9f9160248082019260209290919082900301818787803b15801561169057600080fd5b505af11580156116a4573d6000803e3d6000fd5b505050506040513d60208110156116ba57600080fd5b5051600254604080517f29debc9f00000000000000000000000000000000000000000000000000000000815263ffffffff8d1660048201529051929850600160a060020a03909116916329debc9f916024808201926020929091908290030181600087803b15801561172b57600080fd5b505af115801561173f573d6000803e3d6000fd5b505050506040513d602081101561175557600080fd5b505160008d81526005602081905260408083208201548f845292200154919650600160a060020a0390811695501692506103e86103e689020491506103e86103e6880260035460008f81526005602052604080822060049081015482517f92dd79e6000000000000000000000000000000000000000000000000000000008152600160a060020a03918216928101929092528b8116602483015260448201899052888f0360648301528a811660848301529151959094049550909116926392dd79e69260a480820193929182900301818387803b15801561183557600080fd5b505af1158015611849573d6000803e3d6000fd5b505060035460008e81526005602052604080822060049081015482517f92dd79e6000000000000000000000000000000000000000000000000000000008152600160a060020a03918216928101929092528c8116602483015260448201889052878e036064830152898116608483015291519190931694506392dd79e6935060a48084019382900301818387803b1580156118e357600080fd5b505af11580156118f7573d6000803e3d6000fd5b505060035460008e81526005602052604080822060049081015482517f2b15e857000000000000000000000000000000000000000000000000000000008152600160a060020a03918216928101929092528b81166024830152604482018990529151919093169450632b15e857935060648084019382900301818387803b15801561198157600080fd5b505af1158015611995573d6000803e3d6000fd5b505060035460008f81526005602052604080822060049081015482517f2b15e857000000000000000000000000000000000000000000000000000000008152600160a060020a03918216928101929092528c81166024830152604482018890529151919093169450632b15e857935060648084019382900301818387803b158015611a1f57600080fd5b505af1158015611a33573d6000803e3d6000fd5b50505050505050505050505050505050565b60008160200151826000015160018460400151856060015186608001518760a001518860c001518960e001518a61010001518b61012001518c6101400151604051602001808d60ff1660ff167f01000000000000000000000000000000000000000000000000000000000000000281526001018c60ff1660ff167f01000000000000000000000000000000000000000000000000000000000000000281526001018b63ffffffff1663ffffffff1660e060020a0281526004018a67ffffffffffffffff1667ffffffffffffffff1660c060020a0281526008018967ffffffffffffffff1667ffffffffffffffff1660c060020a0281526008018867ffffffffffffffff1667ffffffffffffffff1660c060020a0281526008018767ffffffffffffffff1667ffffffffffffffff1660c060020a0281526008018667ffffffffffffffff1667ffffffffffffffff1660c060020a0281526008018567ffffffffffffffff1667ffffffffffffffff1660c060020a0281526008018467ffffffffffffffff1667ffffffffffffffff1660c060020a0281526008018367ffffffffffffffff1667ffffffffffffffff1660c060020a0281526008018281526020019c505050505050505050505050506040516020818303038152906040526040518082805190602001908083835b60208310611c505780518252601f199092019160209182019101611c31565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050565b6000818152600560205260408082206001808201548685529284209081015490548585529154849367ffffffffffffffff9081169281168390038116600a0a72010000000000000000000000000000000000009485900482160293909104168201600281061515611cfb57600281048294509450611d08565b8060050260018303945094505b5050509250929050565b6000600584026028198385010182808212611d345781600a0a83029050611d49565b81600003600a0a83811515611d4557fe5b0490505b9695505050505050565b600160a060020a0381161515611d6857600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008481526005602052604081206001908101548291829182918291611e189167ffffffffffffffff680100000000000000008304811692608060020a90041690600c611f52565b6000898152600560205260409020600190810154919350611e609167ffffffffffffffff680100000000000000008204811692608060020a90920416908a908a90600c611efd565b905080821015611eb65760008981526005602052604090206001015468010000000000000000810467ffffffffffffffff90811660c80281169650608060020a9091048116602601168690039350869250611ef1565b600088815260056020526040902060019081015467ffffffffffffffff68010000000000000000820481169750608060020a90910416945092505b50509450945094915050565b6000600a8786020260351983880186010182808212611f235781600a0a83029050611f38565b81600003600a0a83811515611f3457fe5b0490505b8581811515611f4357fe5b049a9950505050505050505050565b60006002850282600d8380881215611f71578760000382019150611f76565b918701915b818310611f8b5750808203600a0a8302611f9f565b828203600a0a84811515611f9b57fe5b0490505b8681811515611faa57fe5b049998505050505050505050565b604080516101a081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081018290526101608101829052610180810191909152905600a165627a7a7230582086ee1a530c81c2e1f17c66c0609e0401a88e1c923a574e0fe74d143347c18db40029`

// DeployRenExSettlement deploys a new Ethereum contract, binding an instance of RenExSettlement to it.
func DeployRenExSettlement(auth *bind.TransactOpts, backend bind.ContractBackend, _orderbookContract common.Address, _renExTokensContract common.Address, _renExBalancesContract common.Address, _submissionGasPriceLimit *big.Int) (common.Address, *types.Transaction, *RenExSettlement, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExSettlementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RenExSettlementBin), backend, _orderbookContract, _renExTokensContract, _renExBalancesContract, _submissionGasPriceLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RenExSettlement{RenExSettlementCaller: RenExSettlementCaller{contract: contract}, RenExSettlementTransactor: RenExSettlementTransactor{contract: contract}, RenExSettlementFilterer: RenExSettlementFilterer{contract: contract}}, nil
}

// RenExSettlement is an auto generated Go binding around an Ethereum contract.
type RenExSettlement struct {
	RenExSettlementCaller     // Read-only binding to the contract
	RenExSettlementTransactor // Write-only binding to the contract
	RenExSettlementFilterer   // Log filterer for contract events
}

// RenExSettlementCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenExSettlementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExSettlementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenExSettlementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExSettlementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenExSettlementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExSettlementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenExSettlementSession struct {
	Contract     *RenExSettlement  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RenExSettlementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenExSettlementCallerSession struct {
	Contract *RenExSettlementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RenExSettlementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenExSettlementTransactorSession struct {
	Contract     *RenExSettlementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RenExSettlementRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenExSettlementRaw struct {
	Contract *RenExSettlement // Generic contract binding to access the raw methods on
}

// RenExSettlementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenExSettlementCallerRaw struct {
	Contract *RenExSettlementCaller // Generic read-only contract binding to access the raw methods on
}

// RenExSettlementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenExSettlementTransactorRaw struct {
	Contract *RenExSettlementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRenExSettlement creates a new instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlement(address common.Address, backend bind.ContractBackend) (*RenExSettlement, error) {
	contract, err := bindRenExSettlement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RenExSettlement{RenExSettlementCaller: RenExSettlementCaller{contract: contract}, RenExSettlementTransactor: RenExSettlementTransactor{contract: contract}, RenExSettlementFilterer: RenExSettlementFilterer{contract: contract}}, nil
}

// NewRenExSettlementCaller creates a new read-only instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlementCaller(address common.Address, caller bind.ContractCaller) (*RenExSettlementCaller, error) {
	contract, err := bindRenExSettlement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementCaller{contract: contract}, nil
}

// NewRenExSettlementTransactor creates a new write-only instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlementTransactor(address common.Address, transactor bind.ContractTransactor) (*RenExSettlementTransactor, error) {
	contract, err := bindRenExSettlement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementTransactor{contract: contract}, nil
}

// NewRenExSettlementFilterer creates a new log filterer instance of RenExSettlement, bound to a specific deployed contract.
func NewRenExSettlementFilterer(address common.Address, filterer bind.ContractFilterer) (*RenExSettlementFilterer, error) {
	contract, err := bindRenExSettlement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementFilterer{contract: contract}, nil
}

// bindRenExSettlement binds a generic wrapper to an already deployed contract.
func bindRenExSettlement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExSettlementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExSettlement *RenExSettlementRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExSettlement.Contract.RenExSettlementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExSettlement *RenExSettlementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenExSettlementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExSettlement *RenExSettlementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenExSettlementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExSettlement *RenExSettlementCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExSettlement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExSettlement *RenExSettlementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExSettlement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExSettlement *RenExSettlementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExSettlement.Contract.contract.Transact(opts, method, params...)
}

// FEESDENOMINATOR is a free data retrieval call binding the contract method 0xbae119ef.
//
// Solidity: function FEES_DENOMINATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) FEESDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "FEES_DENOMINATOR")
	return *ret0, err
}

// FEESDENOMINATOR is a free data retrieval call binding the contract method 0xbae119ef.
//
// Solidity: function FEES_DENOMINATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) FEESDENOMINATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.FEESDENOMINATOR(&_RenExSettlement.CallOpts)
}

// FEESDENOMINATOR is a free data retrieval call binding the contract method 0xbae119ef.
//
// Solidity: function FEES_DENOMINATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) FEESDENOMINATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.FEESDENOMINATOR(&_RenExSettlement.CallOpts)
}

// FEESNUMERATOR is a free data retrieval call binding the contract method 0x7a089c26.
//
// Solidity: function FEES_NUMERATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) FEESNUMERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "FEES_NUMERATOR")
	return *ret0, err
}

// FEESNUMERATOR is a free data retrieval call binding the contract method 0x7a089c26.
//
// Solidity: function FEES_NUMERATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) FEESNUMERATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.FEESNUMERATOR(&_RenExSettlement.CallOpts)
}

// FEESNUMERATOR is a free data retrieval call binding the contract method 0x7a089c26.
//
// Solidity: function FEES_NUMERATOR() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) FEESNUMERATOR() (*big.Int, error) {
	return _RenExSettlement.Contract.FEESNUMERATOR(&_RenExSettlement.CallOpts)
}

// SETTLEMENTIDENTIFIER is a free data retrieval call binding the contract method 0x07067ef0.
//
// Solidity: function SETTLEMENT_IDENTIFIER() constant returns(uint32)
func (_RenExSettlement *RenExSettlementCaller) SETTLEMENTIDENTIFIER(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "SETTLEMENT_IDENTIFIER")
	return *ret0, err
}

// SETTLEMENTIDENTIFIER is a free data retrieval call binding the contract method 0x07067ef0.
//
// Solidity: function SETTLEMENT_IDENTIFIER() constant returns(uint32)
func (_RenExSettlement *RenExSettlementSession) SETTLEMENTIDENTIFIER() (uint32, error) {
	return _RenExSettlement.Contract.SETTLEMENTIDENTIFIER(&_RenExSettlement.CallOpts)
}

// SETTLEMENTIDENTIFIER is a free data retrieval call binding the contract method 0x07067ef0.
//
// Solidity: function SETTLEMENT_IDENTIFIER() constant returns(uint32)
func (_RenExSettlement *RenExSettlementCallerSession) SETTLEMENTIDENTIFIER() (uint32, error) {
	return _RenExSettlement.Contract.SETTLEMENTIDENTIFIER(&_RenExSettlement.CallOpts)
}

// GetMidPrice is a free data retrieval call binding the contract method 0x8fac5ff5.
//
// Solidity: function getMidPrice(_buyID bytes32, _sellID bytes32) constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) GetMidPrice(opts *bind.CallOpts, _buyID [32]byte, _sellID [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "getMidPrice", _buyID, _sellID)
	return *ret0, err
}

// GetMidPrice is a free data retrieval call binding the contract method 0x8fac5ff5.
//
// Solidity: function getMidPrice(_buyID bytes32, _sellID bytes32) constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) GetMidPrice(_buyID [32]byte, _sellID [32]byte) (*big.Int, error) {
	return _RenExSettlement.Contract.GetMidPrice(&_RenExSettlement.CallOpts, _buyID, _sellID)
}

// GetMidPrice is a free data retrieval call binding the contract method 0x8fac5ff5.
//
// Solidity: function getMidPrice(_buyID bytes32, _sellID bytes32) constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) GetMidPrice(_buyID [32]byte, _sellID [32]byte) (*big.Int, error) {
	return _RenExSettlement.Contract.GetMidPrice(&_RenExSettlement.CallOpts, _buyID, _sellID)
}

// GetSettlementDetails is a free data retrieval call binding the contract method 0xeb534ea0.
//
// Solidity: function getSettlementDetails(_buyID bytes32, _sellID bytes32) constant returns(uint256, uint256, uint256, uint256, uint256)
func (_RenExSettlement *RenExSettlementCaller) GetSettlementDetails(opts *bind.CallOpts, _buyID [32]byte, _sellID [32]byte) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _RenExSettlement.contract.Call(opts, out, "getSettlementDetails", _buyID, _sellID)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetSettlementDetails is a free data retrieval call binding the contract method 0xeb534ea0.
//
// Solidity: function getSettlementDetails(_buyID bytes32, _sellID bytes32) constant returns(uint256, uint256, uint256, uint256, uint256)
func (_RenExSettlement *RenExSettlementSession) GetSettlementDetails(_buyID [32]byte, _sellID [32]byte) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _RenExSettlement.Contract.GetSettlementDetails(&_RenExSettlement.CallOpts, _buyID, _sellID)
}

// GetSettlementDetails is a free data retrieval call binding the contract method 0xeb534ea0.
//
// Solidity: function getSettlementDetails(_buyID bytes32, _sellID bytes32) constant returns(uint256, uint256, uint256, uint256, uint256)
func (_RenExSettlement *RenExSettlementCallerSession) GetSettlementDetails(_buyID [32]byte, _sellID [32]byte) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _RenExSettlement.Contract.GetSettlementDetails(&_RenExSettlement.CallOpts, _buyID, _sellID)
}

// OrderbookContract is a free data retrieval call binding the contract method 0xbbe7221e.
//
// Solidity: function orderbookContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) OrderbookContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "orderbookContract")
	return *ret0, err
}

// OrderbookContract is a free data retrieval call binding the contract method 0xbbe7221e.
//
// Solidity: function orderbookContract() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) OrderbookContract() (common.Address, error) {
	return _RenExSettlement.Contract.OrderbookContract(&_RenExSettlement.CallOpts)
}

// OrderbookContract is a free data retrieval call binding the contract method 0xbbe7221e.
//
// Solidity: function orderbookContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) OrderbookContract() (common.Address, error) {
	return _RenExSettlement.Contract.OrderbookContract(&_RenExSettlement.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders( bytes32) constant returns(parity uint8, orderType uint8, expiry uint64, tokens uint64, priceC uint64, priceQ uint64, volumeC uint64, volumeQ uint64, minimumVolumeC uint64, minimumVolumeQ uint64, nonceHash uint256, trader address, submitter address)
func (_RenExSettlement *RenExSettlementCaller) Orders(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Parity         uint8
	OrderType      uint8
	Expiry         uint64
	Tokens         uint64
	PriceC         uint64
	PriceQ         uint64
	VolumeC        uint64
	VolumeQ        uint64
	MinimumVolumeC uint64
	MinimumVolumeQ uint64
	NonceHash      *big.Int
	Trader         common.Address
	Submitter      common.Address
}, error) {
	ret := new(struct {
		Parity         uint8
		OrderType      uint8
		Expiry         uint64
		Tokens         uint64
		PriceC         uint64
		PriceQ         uint64
		VolumeC        uint64
		VolumeQ        uint64
		MinimumVolumeC uint64
		MinimumVolumeQ uint64
		NonceHash      *big.Int
		Trader         common.Address
		Submitter      common.Address
	})
	out := ret
	err := _RenExSettlement.contract.Call(opts, out, "orders", arg0)
	return *ret, err
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders( bytes32) constant returns(parity uint8, orderType uint8, expiry uint64, tokens uint64, priceC uint64, priceQ uint64, volumeC uint64, volumeQ uint64, minimumVolumeC uint64, minimumVolumeQ uint64, nonceHash uint256, trader address, submitter address)
func (_RenExSettlement *RenExSettlementSession) Orders(arg0 [32]byte) (struct {
	Parity         uint8
	OrderType      uint8
	Expiry         uint64
	Tokens         uint64
	PriceC         uint64
	PriceQ         uint64
	VolumeC        uint64
	VolumeQ        uint64
	MinimumVolumeC uint64
	MinimumVolumeQ uint64
	NonceHash      *big.Int
	Trader         common.Address
	Submitter      common.Address
}, error) {
	return _RenExSettlement.Contract.Orders(&_RenExSettlement.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders( bytes32) constant returns(parity uint8, orderType uint8, expiry uint64, tokens uint64, priceC uint64, priceQ uint64, volumeC uint64, volumeQ uint64, minimumVolumeC uint64, minimumVolumeQ uint64, nonceHash uint256, trader address, submitter address)
func (_RenExSettlement *RenExSettlementCallerSession) Orders(arg0 [32]byte) (struct {
	Parity         uint8
	OrderType      uint8
	Expiry         uint64
	Tokens         uint64
	PriceC         uint64
	PriceQ         uint64
	VolumeC        uint64
	VolumeQ        uint64
	MinimumVolumeC uint64
	MinimumVolumeQ uint64
	NonceHash      *big.Int
	Trader         common.Address
	Submitter      common.Address
}, error) {
	return _RenExSettlement.Contract.Orders(&_RenExSettlement.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) Owner() (common.Address, error) {
	return _RenExSettlement.Contract.Owner(&_RenExSettlement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) Owner() (common.Address, error) {
	return _RenExSettlement.Contract.Owner(&_RenExSettlement.CallOpts)
}

// RenExBalancesContract is a free data retrieval call binding the contract method 0xddf25ce9.
//
// Solidity: function renExBalancesContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) RenExBalancesContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "renExBalancesContract")
	return *ret0, err
}

// RenExBalancesContract is a free data retrieval call binding the contract method 0xddf25ce9.
//
// Solidity: function renExBalancesContract() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) RenExBalancesContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExBalancesContract(&_RenExSettlement.CallOpts)
}

// RenExBalancesContract is a free data retrieval call binding the contract method 0xddf25ce9.
//
// Solidity: function renExBalancesContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) RenExBalancesContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExBalancesContract(&_RenExSettlement.CallOpts)
}

// RenExTokensContract is a free data retrieval call binding the contract method 0x4015e83b.
//
// Solidity: function renExTokensContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCaller) RenExTokensContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "renExTokensContract")
	return *ret0, err
}

// RenExTokensContract is a free data retrieval call binding the contract method 0x4015e83b.
//
// Solidity: function renExTokensContract() constant returns(address)
func (_RenExSettlement *RenExSettlementSession) RenExTokensContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExTokensContract(&_RenExSettlement.CallOpts)
}

// RenExTokensContract is a free data retrieval call binding the contract method 0x4015e83b.
//
// Solidity: function renExTokensContract() constant returns(address)
func (_RenExSettlement *RenExSettlementCallerSession) RenExTokensContract() (common.Address, error) {
	return _RenExSettlement.Contract.RenExTokensContract(&_RenExSettlement.CallOpts)
}

// SubmissionGasPriceLimit is a free data retrieval call binding the contract method 0x4c3052de.
//
// Solidity: function submissionGasPriceLimit() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCaller) SubmissionGasPriceLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RenExSettlement.contract.Call(opts, out, "submissionGasPriceLimit")
	return *ret0, err
}

// SubmissionGasPriceLimit is a free data retrieval call binding the contract method 0x4c3052de.
//
// Solidity: function submissionGasPriceLimit() constant returns(uint256)
func (_RenExSettlement *RenExSettlementSession) SubmissionGasPriceLimit() (*big.Int, error) {
	return _RenExSettlement.Contract.SubmissionGasPriceLimit(&_RenExSettlement.CallOpts)
}

// SubmissionGasPriceLimit is a free data retrieval call binding the contract method 0x4c3052de.
//
// Solidity: function submissionGasPriceLimit() constant returns(uint256)
func (_RenExSettlement *RenExSettlementCallerSession) SubmissionGasPriceLimit() (*big.Int, error) {
	return _RenExSettlement.Contract.SubmissionGasPriceLimit(&_RenExSettlement.CallOpts)
}

// VerifyMatch is a free data retrieval call binding the contract method 0x1f88d9c7.
//
// Solidity: function verifyMatch(_buyID bytes32, _sellID bytes32) constant returns(uint32, uint32)
func (_RenExSettlement *RenExSettlementCaller) VerifyMatch(opts *bind.CallOpts, _buyID [32]byte, _sellID [32]byte) (uint32, uint32, error) {
	var (
		ret0 = new(uint32)
		ret1 = new(uint32)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RenExSettlement.contract.Call(opts, out, "verifyMatch", _buyID, _sellID)
	return *ret0, *ret1, err
}

// VerifyMatch is a free data retrieval call binding the contract method 0x1f88d9c7.
//
// Solidity: function verifyMatch(_buyID bytes32, _sellID bytes32) constant returns(uint32, uint32)
func (_RenExSettlement *RenExSettlementSession) VerifyMatch(_buyID [32]byte, _sellID [32]byte) (uint32, uint32, error) {
	return _RenExSettlement.Contract.VerifyMatch(&_RenExSettlement.CallOpts, _buyID, _sellID)
}

// VerifyMatch is a free data retrieval call binding the contract method 0x1f88d9c7.
//
// Solidity: function verifyMatch(_buyID bytes32, _sellID bytes32) constant returns(uint32, uint32)
func (_RenExSettlement *RenExSettlementCallerSession) VerifyMatch(_buyID [32]byte, _sellID [32]byte) (uint32, uint32, error) {
	return _RenExSettlement.Contract.VerifyMatch(&_RenExSettlement.CallOpts, _buyID, _sellID)
}

// VerifyOrder is a free data retrieval call binding the contract method 0x169fb328.
//
// Solidity: function verifyOrder(_orderID bytes32) constant returns()
func (_RenExSettlement *RenExSettlementCaller) VerifyOrder(opts *bind.CallOpts, _orderID [32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _RenExSettlement.contract.Call(opts, out, "verifyOrder", _orderID)
	return err
}

// VerifyOrder is a free data retrieval call binding the contract method 0x169fb328.
//
// Solidity: function verifyOrder(_orderID bytes32) constant returns()
func (_RenExSettlement *RenExSettlementSession) VerifyOrder(_orderID [32]byte) error {
	return _RenExSettlement.Contract.VerifyOrder(&_RenExSettlement.CallOpts, _orderID)
}

// VerifyOrder is a free data retrieval call binding the contract method 0x169fb328.
//
// Solidity: function verifyOrder(_orderID bytes32) constant returns()
func (_RenExSettlement *RenExSettlementCallerSession) VerifyOrder(_orderID [32]byte) error {
	return _RenExSettlement.Contract.VerifyOrder(&_RenExSettlement.CallOpts, _orderID)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExSettlement *RenExSettlementTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExSettlement *RenExSettlementSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenounceOwnership(&_RenExSettlement.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExSettlement *RenExSettlementTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExSettlement.Contract.RenounceOwnership(&_RenExSettlement.TransactOpts)
}

// SubmitMatch is a paid mutator transaction binding the contract method 0x2a337d30.
//
// Solidity: function submitMatch(_buyID bytes32, _sellID bytes32) returns()
func (_RenExSettlement *RenExSettlementTransactor) SubmitMatch(opts *bind.TransactOpts, _buyID [32]byte, _sellID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "submitMatch", _buyID, _sellID)
}

// SubmitMatch is a paid mutator transaction binding the contract method 0x2a337d30.
//
// Solidity: function submitMatch(_buyID bytes32, _sellID bytes32) returns()
func (_RenExSettlement *RenExSettlementSession) SubmitMatch(_buyID [32]byte, _sellID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.Contract.SubmitMatch(&_RenExSettlement.TransactOpts, _buyID, _sellID)
}

// SubmitMatch is a paid mutator transaction binding the contract method 0x2a337d30.
//
// Solidity: function submitMatch(_buyID bytes32, _sellID bytes32) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) SubmitMatch(_buyID [32]byte, _sellID [32]byte) (*types.Transaction, error) {
	return _RenExSettlement.Contract.SubmitMatch(&_RenExSettlement.TransactOpts, _buyID, _sellID)
}

// SubmitOrder is a paid mutator transaction binding the contract method 0x367ba52c.
//
// Solidity: function submitOrder(_orderType uint8, _parity uint8, _expiry uint64, _tokens uint64, _priceC uint16, _priceQ uint16, _volumeC uint16, _volumeQ uint16, _minimumVolumeC uint16, _minimumVolumeQ uint16, _nonceHash uint256) returns()
func (_RenExSettlement *RenExSettlementTransactor) SubmitOrder(opts *bind.TransactOpts, _orderType uint8, _parity uint8, _expiry uint64, _tokens uint64, _priceC uint16, _priceQ uint16, _volumeC uint16, _volumeQ uint16, _minimumVolumeC uint16, _minimumVolumeQ uint16, _nonceHash *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "submitOrder", _orderType, _parity, _expiry, _tokens, _priceC, _priceQ, _volumeC, _volumeQ, _minimumVolumeC, _minimumVolumeQ, _nonceHash)
}

// SubmitOrder is a paid mutator transaction binding the contract method 0x367ba52c.
//
// Solidity: function submitOrder(_orderType uint8, _parity uint8, _expiry uint64, _tokens uint64, _priceC uint16, _priceQ uint16, _volumeC uint16, _volumeQ uint16, _minimumVolumeC uint16, _minimumVolumeQ uint16, _nonceHash uint256) returns()
func (_RenExSettlement *RenExSettlementSession) SubmitOrder(_orderType uint8, _parity uint8, _expiry uint64, _tokens uint64, _priceC uint16, _priceQ uint16, _volumeC uint16, _volumeQ uint16, _minimumVolumeC uint16, _minimumVolumeQ uint16, _nonceHash *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.SubmitOrder(&_RenExSettlement.TransactOpts, _orderType, _parity, _expiry, _tokens, _priceC, _priceQ, _volumeC, _volumeQ, _minimumVolumeC, _minimumVolumeQ, _nonceHash)
}

// SubmitOrder is a paid mutator transaction binding the contract method 0x367ba52c.
//
// Solidity: function submitOrder(_orderType uint8, _parity uint8, _expiry uint64, _tokens uint64, _priceC uint16, _priceQ uint16, _volumeC uint16, _volumeQ uint16, _minimumVolumeC uint16, _minimumVolumeQ uint16, _nonceHash uint256) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) SubmitOrder(_orderType uint8, _parity uint8, _expiry uint64, _tokens uint64, _priceC uint16, _priceQ uint16, _volumeC uint16, _volumeQ uint16, _minimumVolumeC uint16, _minimumVolumeQ uint16, _nonceHash *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.SubmitOrder(&_RenExSettlement.TransactOpts, _orderType, _parity, _expiry, _tokens, _priceC, _priceQ, _volumeC, _volumeQ, _minimumVolumeC, _minimumVolumeQ, _nonceHash)
}

// TraderCanWithdraw is a paid mutator transaction binding the contract method 0xa3bdaedc.
//
// Solidity: function traderCanWithdraw(_trader address, _token address, amount uint256) returns(bool)
func (_RenExSettlement *RenExSettlementTransactor) TraderCanWithdraw(opts *bind.TransactOpts, _trader common.Address, _token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "traderCanWithdraw", _trader, _token, amount)
}

// TraderCanWithdraw is a paid mutator transaction binding the contract method 0xa3bdaedc.
//
// Solidity: function traderCanWithdraw(_trader address, _token address, amount uint256) returns(bool)
func (_RenExSettlement *RenExSettlementSession) TraderCanWithdraw(_trader common.Address, _token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.TraderCanWithdraw(&_RenExSettlement.TransactOpts, _trader, _token, amount)
}

// TraderCanWithdraw is a paid mutator transaction binding the contract method 0xa3bdaedc.
//
// Solidity: function traderCanWithdraw(_trader address, _token address, amount uint256) returns(bool)
func (_RenExSettlement *RenExSettlementTransactorSession) TraderCanWithdraw(_trader common.Address, _token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.TraderCanWithdraw(&_RenExSettlement.TransactOpts, _trader, _token, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExSettlement *RenExSettlementTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExSettlement *RenExSettlementSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.TransferOwnership(&_RenExSettlement.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.TransferOwnership(&_RenExSettlement.TransactOpts, _newOwner)
}

// UpdateOrderbook is a paid mutator transaction binding the contract method 0x6074b806.
//
// Solidity: function updateOrderbook(_newOrderbookContract address) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateOrderbook(opts *bind.TransactOpts, _newOrderbookContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateOrderbook", _newOrderbookContract)
}

// UpdateOrderbook is a paid mutator transaction binding the contract method 0x6074b806.
//
// Solidity: function updateOrderbook(_newOrderbookContract address) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateOrderbook(_newOrderbookContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateOrderbook(&_RenExSettlement.TransactOpts, _newOrderbookContract)
}

// UpdateOrderbook is a paid mutator transaction binding the contract method 0x6074b806.
//
// Solidity: function updateOrderbook(_newOrderbookContract address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateOrderbook(_newOrderbookContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateOrderbook(&_RenExSettlement.TransactOpts, _newOrderbookContract)
}

// UpdateRenExBalances is a paid mutator transaction binding the contract method 0xee0715ed.
//
// Solidity: function updateRenExBalances(_newRenExBalancesContract address) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateRenExBalances(opts *bind.TransactOpts, _newRenExBalancesContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateRenExBalances", _newRenExBalancesContract)
}

// UpdateRenExBalances is a paid mutator transaction binding the contract method 0xee0715ed.
//
// Solidity: function updateRenExBalances(_newRenExBalancesContract address) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateRenExBalances(_newRenExBalancesContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateRenExBalances(&_RenExSettlement.TransactOpts, _newRenExBalancesContract)
}

// UpdateRenExBalances is a paid mutator transaction binding the contract method 0xee0715ed.
//
// Solidity: function updateRenExBalances(_newRenExBalancesContract address) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateRenExBalances(_newRenExBalancesContract common.Address) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateRenExBalances(&_RenExSettlement.TransactOpts, _newRenExBalancesContract)
}

// UpdateSubmissionGasPriceLimit is a paid mutator transaction binding the contract method 0x675df16f.
//
// Solidity: function updateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit uint256) returns()
func (_RenExSettlement *RenExSettlementTransactor) UpdateSubmissionGasPriceLimit(opts *bind.TransactOpts, _newSubmissionGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.contract.Transact(opts, "updateSubmissionGasPriceLimit", _newSubmissionGasPriceLimit)
}

// UpdateSubmissionGasPriceLimit is a paid mutator transaction binding the contract method 0x675df16f.
//
// Solidity: function updateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit uint256) returns()
func (_RenExSettlement *RenExSettlementSession) UpdateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateSubmissionGasPriceLimit(&_RenExSettlement.TransactOpts, _newSubmissionGasPriceLimit)
}

// UpdateSubmissionGasPriceLimit is a paid mutator transaction binding the contract method 0x675df16f.
//
// Solidity: function updateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit uint256) returns()
func (_RenExSettlement *RenExSettlementTransactorSession) UpdateSubmissionGasPriceLimit(_newSubmissionGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _RenExSettlement.Contract.UpdateSubmissionGasPriceLimit(&_RenExSettlement.TransactOpts, _newSubmissionGasPriceLimit)
}

// RenExSettlementOrderbookUpdatesIterator is returned from FilterOrderbookUpdates and is used to iterate over the raw logs and unpacked data for OrderbookUpdates events raised by the RenExSettlement contract.
type RenExSettlementOrderbookUpdatesIterator struct {
	Event *RenExSettlementOrderbookUpdates // Event containing the contract specifics and raw log

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
func (it *RenExSettlementOrderbookUpdatesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementOrderbookUpdates)
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
		it.Event = new(RenExSettlementOrderbookUpdates)
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
func (it *RenExSettlementOrderbookUpdatesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementOrderbookUpdatesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementOrderbookUpdates represents a OrderbookUpdates event raised by the RenExSettlement contract.
type RenExSettlementOrderbookUpdates struct {
	PreviousOrderbook common.Address
	NextOrderbook     common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrderbookUpdates is a free log retrieval operation binding the contract event 0x05565154275c40146888f6e69d9557d712c9384838de11284a6742df83e0e01b.
//
// Solidity: e OrderbookUpdates(previousOrderbook address, nextOrderbook address)
func (_RenExSettlement *RenExSettlementFilterer) FilterOrderbookUpdates(opts *bind.FilterOpts) (*RenExSettlementOrderbookUpdatesIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "OrderbookUpdates")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementOrderbookUpdatesIterator{contract: _RenExSettlement.contract, event: "OrderbookUpdates", logs: logs, sub: sub}, nil
}

// WatchOrderbookUpdates is a free log subscription operation binding the contract event 0x05565154275c40146888f6e69d9557d712c9384838de11284a6742df83e0e01b.
//
// Solidity: e OrderbookUpdates(previousOrderbook address, nextOrderbook address)
func (_RenExSettlement *RenExSettlementFilterer) WatchOrderbookUpdates(opts *bind.WatchOpts, sink chan<- *RenExSettlementOrderbookUpdates) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "OrderbookUpdates")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementOrderbookUpdates)
				if err := _RenExSettlement.contract.UnpackLog(event, "OrderbookUpdates", log); err != nil {
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

// RenExSettlementOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RenExSettlement contract.
type RenExSettlementOwnershipRenouncedIterator struct {
	Event *RenExSettlementOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RenExSettlementOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementOwnershipRenounced)
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
		it.Event = new(RenExSettlementOwnershipRenounced)
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
func (it *RenExSettlementOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementOwnershipRenounced represents a OwnershipRenounced event raised by the RenExSettlement contract.
type RenExSettlementOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RenExSettlementOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementOwnershipRenouncedIterator{contract: _RenExSettlement.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RenExSettlementOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementOwnershipRenounced)
				if err := _RenExSettlement.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RenExSettlementOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RenExSettlement contract.
type RenExSettlementOwnershipTransferredIterator struct {
	Event *RenExSettlementOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RenExSettlementOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementOwnershipTransferred)
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
		it.Event = new(RenExSettlementOwnershipTransferred)
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
func (it *RenExSettlementOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementOwnershipTransferred represents a OwnershipTransferred event raised by the RenExSettlement contract.
type RenExSettlementOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RenExSettlementOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExSettlementOwnershipTransferredIterator{contract: _RenExSettlement.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExSettlement *RenExSettlementFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RenExSettlementOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementOwnershipTransferred)
				if err := _RenExSettlement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RenExSettlementRenExBalancesUpdatesIterator is returned from FilterRenExBalancesUpdates and is used to iterate over the raw logs and unpacked data for RenExBalancesUpdates events raised by the RenExSettlement contract.
type RenExSettlementRenExBalancesUpdatesIterator struct {
	Event *RenExSettlementRenExBalancesUpdates // Event containing the contract specifics and raw log

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
func (it *RenExSettlementRenExBalancesUpdatesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementRenExBalancesUpdates)
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
		it.Event = new(RenExSettlementRenExBalancesUpdates)
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
func (it *RenExSettlementRenExBalancesUpdatesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementRenExBalancesUpdatesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementRenExBalancesUpdates represents a RenExBalancesUpdates event raised by the RenExSettlement contract.
type RenExSettlementRenExBalancesUpdates struct {
	PreviousRenExBalances common.Address
	NextRenExBalances     common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRenExBalancesUpdates is a free log retrieval operation binding the contract event 0x1370a598dc4f33911a2368195ac8e693faec43f2ce450cd909dcd64d590cd222.
//
// Solidity: e RenExBalancesUpdates(previousRenExBalances address, nextRenExBalances address)
func (_RenExSettlement *RenExSettlementFilterer) FilterRenExBalancesUpdates(opts *bind.FilterOpts) (*RenExSettlementRenExBalancesUpdatesIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "RenExBalancesUpdates")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementRenExBalancesUpdatesIterator{contract: _RenExSettlement.contract, event: "RenExBalancesUpdates", logs: logs, sub: sub}, nil
}

// WatchRenExBalancesUpdates is a free log subscription operation binding the contract event 0x1370a598dc4f33911a2368195ac8e693faec43f2ce450cd909dcd64d590cd222.
//
// Solidity: e RenExBalancesUpdates(previousRenExBalances address, nextRenExBalances address)
func (_RenExSettlement *RenExSettlementFilterer) WatchRenExBalancesUpdates(opts *bind.WatchOpts, sink chan<- *RenExSettlementRenExBalancesUpdates) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "RenExBalancesUpdates")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementRenExBalancesUpdates)
				if err := _RenExSettlement.contract.UnpackLog(event, "RenExBalancesUpdates", log); err != nil {
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

// RenExSettlementSubmissionGasPriceLimitUpdatesIterator is returned from FilterSubmissionGasPriceLimitUpdates and is used to iterate over the raw logs and unpacked data for SubmissionGasPriceLimitUpdates events raised by the RenExSettlement contract.
type RenExSettlementSubmissionGasPriceLimitUpdatesIterator struct {
	Event *RenExSettlementSubmissionGasPriceLimitUpdates // Event containing the contract specifics and raw log

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
func (it *RenExSettlementSubmissionGasPriceLimitUpdatesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementSubmissionGasPriceLimitUpdates)
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
		it.Event = new(RenExSettlementSubmissionGasPriceLimitUpdates)
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
func (it *RenExSettlementSubmissionGasPriceLimitUpdatesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementSubmissionGasPriceLimitUpdatesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementSubmissionGasPriceLimitUpdates represents a SubmissionGasPriceLimitUpdates event raised by the RenExSettlement contract.
type RenExSettlementSubmissionGasPriceLimitUpdates struct {
	PreviousSubmissionGasPriceLimit *big.Int
	NextSubmissionGasPriceLimit     *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterSubmissionGasPriceLimitUpdates is a free log retrieval operation binding the contract event 0xccf294e3e783bf1a9ff94493761c96e0ccbc809312bcb5376beaa9baa31de324.
//
// Solidity: e SubmissionGasPriceLimitUpdates(previousSubmissionGasPriceLimit uint256, nextSubmissionGasPriceLimit uint256)
func (_RenExSettlement *RenExSettlementFilterer) FilterSubmissionGasPriceLimitUpdates(opts *bind.FilterOpts) (*RenExSettlementSubmissionGasPriceLimitUpdatesIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "SubmissionGasPriceLimitUpdates")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementSubmissionGasPriceLimitUpdatesIterator{contract: _RenExSettlement.contract, event: "SubmissionGasPriceLimitUpdates", logs: logs, sub: sub}, nil
}

// WatchSubmissionGasPriceLimitUpdates is a free log subscription operation binding the contract event 0xccf294e3e783bf1a9ff94493761c96e0ccbc809312bcb5376beaa9baa31de324.
//
// Solidity: e SubmissionGasPriceLimitUpdates(previousSubmissionGasPriceLimit uint256, nextSubmissionGasPriceLimit uint256)
func (_RenExSettlement *RenExSettlementFilterer) WatchSubmissionGasPriceLimitUpdates(opts *bind.WatchOpts, sink chan<- *RenExSettlementSubmissionGasPriceLimitUpdates) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "SubmissionGasPriceLimitUpdates")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementSubmissionGasPriceLimitUpdates)
				if err := _RenExSettlement.contract.UnpackLog(event, "SubmissionGasPriceLimitUpdates", log); err != nil {
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

// RenExSettlementTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RenExSettlement contract.
type RenExSettlementTransferIterator struct {
	Event *RenExSettlementTransfer // Event containing the contract specifics and raw log

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
func (it *RenExSettlementTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExSettlementTransfer)
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
		it.Event = new(RenExSettlementTransfer)
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
func (it *RenExSettlementTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExSettlementTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExSettlementTransfer represents a Transfer event raised by the RenExSettlement contract.
type RenExSettlementTransfer struct {
	From  common.Address
	To    common.Address
	Token uint32
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xe6d64cac65c05a9715c20bfc2ed8851bfe1c155b6cd1ee42d313d0880fc48b9f.
//
// Solidity: e Transfer(from address, to address, token uint32, value uint256)
func (_RenExSettlement *RenExSettlementFilterer) FilterTransfer(opts *bind.FilterOpts) (*RenExSettlementTransferIterator, error) {

	logs, sub, err := _RenExSettlement.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &RenExSettlementTransferIterator{contract: _RenExSettlement.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xe6d64cac65c05a9715c20bfc2ed8851bfe1c155b6cd1ee42d313d0880fc48b9f.
//
// Solidity: e Transfer(from address, to address, token uint32, value uint256)
func (_RenExSettlement *RenExSettlementFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RenExSettlementTransfer) (event.Subscription, error) {

	logs, sub, err := _RenExSettlement.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExSettlementTransfer)
				if err := _RenExSettlement.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// RenExTokensABI is the input ABI used to generate the binding from.
const RenExTokensABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenCode\",\"type\":\"uint32\"}],\"name\":\"deregisterToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"tokenAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"tokenIsRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenCode\",\"type\":\"uint32\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"registerToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"tokenDecimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenCode\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenCode\",\"type\":\"uint32\"}],\"name\":\"TokenDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RenExTokensBin is the compiled bytecode used for deploying new contracts.
const RenExTokensBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610489806100326000396000f30060806040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166325fc575a811461009257806329debc9f146100b2578063329ed381146100ec578063715018a61461011e5780638da5cb5b146101335780639e20a9a014610148578063c865992b14610178578063f2fde38b146101ac575b600080fd5b34801561009e57600080fd5b506100b063ffffffff600435166101cd565b005b3480156100be57600080fd5b506100d063ffffffff60043516610239565b60408051600160a060020a039092168252519081900360200190f35b3480156100f857600080fd5b5061010a63ffffffff60043516610254565b604080519115158252519081900360200190f35b34801561012a57600080fd5b506100b0610269565b34801561013f57600080fd5b506100d06102d5565b34801561015457600080fd5b506100b063ffffffff60043516600160a060020a036024351660ff604435166102e4565b34801561018457600080fd5b5061019663ffffffff600435166103a8565b6040805160ff9092168252519081900360200190f35b3480156101b857600080fd5b506100b0600160a060020a03600435166103bd565b600054600160a060020a031633146101e457600080fd5b63ffffffff8116600081815260036020908152604091829020805460ff19169055815192835290517f370a83728af25aff893eba93c92cf098d1f4ebd31ecce8487f73edfa7db8a13b9281900390910190a150565b600160205260009081526040902054600160a060020a031681565b60036020526000908152604090205460ff1681565b600054600160a060020a0316331461028057600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600054600160a060020a031633146102fb57600080fd5b63ffffffff831660008181526001602081815260408084208054600160a060020a03891673ffffffffffffffffffffffffffffffffffffffff19909116811790915560028352818520805460ff891660ff1991821681179092556003855295839020805490961690941790945580519485529084019290925282820152517fd6a4fb09bbf95930f0ad37c272f735e3453496dcdfc294f49d11d8d1fc3718189181900360600190a1505050565b60026020526000908152604090205460ff1681565b600054600160a060020a031633146103d457600080fd5b6103dd816103e0565b50565b600160a060020a03811615156103f557600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820361178c298483e7e0fd915692f38a9a9ee23a3a80d290c2f00b949a51973a9880029`

// DeployRenExTokens deploys a new Ethereum contract, binding an instance of RenExTokens to it.
func DeployRenExTokens(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RenExTokens, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExTokensABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RenExTokensBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RenExTokens{RenExTokensCaller: RenExTokensCaller{contract: contract}, RenExTokensTransactor: RenExTokensTransactor{contract: contract}, RenExTokensFilterer: RenExTokensFilterer{contract: contract}}, nil
}

// RenExTokens is an auto generated Go binding around an Ethereum contract.
type RenExTokens struct {
	RenExTokensCaller     // Read-only binding to the contract
	RenExTokensTransactor // Write-only binding to the contract
	RenExTokensFilterer   // Log filterer for contract events
}

// RenExTokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type RenExTokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExTokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RenExTokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExTokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RenExTokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RenExTokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RenExTokensSession struct {
	Contract     *RenExTokens      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RenExTokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RenExTokensCallerSession struct {
	Contract *RenExTokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RenExTokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RenExTokensTransactorSession struct {
	Contract     *RenExTokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RenExTokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type RenExTokensRaw struct {
	Contract *RenExTokens // Generic contract binding to access the raw methods on
}

// RenExTokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RenExTokensCallerRaw struct {
	Contract *RenExTokensCaller // Generic read-only contract binding to access the raw methods on
}

// RenExTokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RenExTokensTransactorRaw struct {
	Contract *RenExTokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRenExTokens creates a new instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokens(address common.Address, backend bind.ContractBackend) (*RenExTokens, error) {
	contract, err := bindRenExTokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RenExTokens{RenExTokensCaller: RenExTokensCaller{contract: contract}, RenExTokensTransactor: RenExTokensTransactor{contract: contract}, RenExTokensFilterer: RenExTokensFilterer{contract: contract}}, nil
}

// NewRenExTokensCaller creates a new read-only instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokensCaller(address common.Address, caller bind.ContractCaller) (*RenExTokensCaller, error) {
	contract, err := bindRenExTokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RenExTokensCaller{contract: contract}, nil
}

// NewRenExTokensTransactor creates a new write-only instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*RenExTokensTransactor, error) {
	contract, err := bindRenExTokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RenExTokensTransactor{contract: contract}, nil
}

// NewRenExTokensFilterer creates a new log filterer instance of RenExTokens, bound to a specific deployed contract.
func NewRenExTokensFilterer(address common.Address, filterer bind.ContractFilterer) (*RenExTokensFilterer, error) {
	contract, err := bindRenExTokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RenExTokensFilterer{contract: contract}, nil
}

// bindRenExTokens binds a generic wrapper to an already deployed contract.
func bindRenExTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RenExTokensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExTokens *RenExTokensRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExTokens.Contract.RenExTokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExTokens *RenExTokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExTokens.Contract.RenExTokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExTokens *RenExTokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExTokens.Contract.RenExTokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RenExTokens *RenExTokensCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RenExTokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RenExTokens *RenExTokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExTokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RenExTokens *RenExTokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RenExTokens.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExTokens *RenExTokensCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExTokens.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExTokens *RenExTokensSession) Owner() (common.Address, error) {
	return _RenExTokens.Contract.Owner(&_RenExTokens.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RenExTokens *RenExTokensCallerSession) Owner() (common.Address, error) {
	return _RenExTokens.Contract.Owner(&_RenExTokens.CallOpts)
}

// TokenAddresses is a free data retrieval call binding the contract method 0x29debc9f.
//
// Solidity: function tokenAddresses( uint32) constant returns(address)
func (_RenExTokens *RenExTokensCaller) TokenAddresses(opts *bind.CallOpts, arg0 uint32) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RenExTokens.contract.Call(opts, out, "tokenAddresses", arg0)
	return *ret0, err
}

// TokenAddresses is a free data retrieval call binding the contract method 0x29debc9f.
//
// Solidity: function tokenAddresses( uint32) constant returns(address)
func (_RenExTokens *RenExTokensSession) TokenAddresses(arg0 uint32) (common.Address, error) {
	return _RenExTokens.Contract.TokenAddresses(&_RenExTokens.CallOpts, arg0)
}

// TokenAddresses is a free data retrieval call binding the contract method 0x29debc9f.
//
// Solidity: function tokenAddresses( uint32) constant returns(address)
func (_RenExTokens *RenExTokensCallerSession) TokenAddresses(arg0 uint32) (common.Address, error) {
	return _RenExTokens.Contract.TokenAddresses(&_RenExTokens.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0xc865992b.
//
// Solidity: function tokenDecimals( uint32) constant returns(uint8)
func (_RenExTokens *RenExTokensCaller) TokenDecimals(opts *bind.CallOpts, arg0 uint32) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RenExTokens.contract.Call(opts, out, "tokenDecimals", arg0)
	return *ret0, err
}

// TokenDecimals is a free data retrieval call binding the contract method 0xc865992b.
//
// Solidity: function tokenDecimals( uint32) constant returns(uint8)
func (_RenExTokens *RenExTokensSession) TokenDecimals(arg0 uint32) (uint8, error) {
	return _RenExTokens.Contract.TokenDecimals(&_RenExTokens.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0xc865992b.
//
// Solidity: function tokenDecimals( uint32) constant returns(uint8)
func (_RenExTokens *RenExTokensCallerSession) TokenDecimals(arg0 uint32) (uint8, error) {
	return _RenExTokens.Contract.TokenDecimals(&_RenExTokens.CallOpts, arg0)
}

// TokenIsRegistered is a free data retrieval call binding the contract method 0x329ed381.
//
// Solidity: function tokenIsRegistered( uint32) constant returns(bool)
func (_RenExTokens *RenExTokensCaller) TokenIsRegistered(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RenExTokens.contract.Call(opts, out, "tokenIsRegistered", arg0)
	return *ret0, err
}

// TokenIsRegistered is a free data retrieval call binding the contract method 0x329ed381.
//
// Solidity: function tokenIsRegistered( uint32) constant returns(bool)
func (_RenExTokens *RenExTokensSession) TokenIsRegistered(arg0 uint32) (bool, error) {
	return _RenExTokens.Contract.TokenIsRegistered(&_RenExTokens.CallOpts, arg0)
}

// TokenIsRegistered is a free data retrieval call binding the contract method 0x329ed381.
//
// Solidity: function tokenIsRegistered( uint32) constant returns(bool)
func (_RenExTokens *RenExTokensCallerSession) TokenIsRegistered(arg0 uint32) (bool, error) {
	return _RenExTokens.Contract.TokenIsRegistered(&_RenExTokens.CallOpts, arg0)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0x25fc575a.
//
// Solidity: function deregisterToken(_tokenCode uint32) returns()
func (_RenExTokens *RenExTokensTransactor) DeregisterToken(opts *bind.TransactOpts, _tokenCode uint32) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "deregisterToken", _tokenCode)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0x25fc575a.
//
// Solidity: function deregisterToken(_tokenCode uint32) returns()
func (_RenExTokens *RenExTokensSession) DeregisterToken(_tokenCode uint32) (*types.Transaction, error) {
	return _RenExTokens.Contract.DeregisterToken(&_RenExTokens.TransactOpts, _tokenCode)
}

// DeregisterToken is a paid mutator transaction binding the contract method 0x25fc575a.
//
// Solidity: function deregisterToken(_tokenCode uint32) returns()
func (_RenExTokens *RenExTokensTransactorSession) DeregisterToken(_tokenCode uint32) (*types.Transaction, error) {
	return _RenExTokens.Contract.DeregisterToken(&_RenExTokens.TransactOpts, _tokenCode)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x9e20a9a0.
//
// Solidity: function registerToken(_tokenCode uint32, _tokenAddress address, _tokenDecimals uint8) returns()
func (_RenExTokens *RenExTokensTransactor) RegisterToken(opts *bind.TransactOpts, _tokenCode uint32, _tokenAddress common.Address, _tokenDecimals uint8) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "registerToken", _tokenCode, _tokenAddress, _tokenDecimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x9e20a9a0.
//
// Solidity: function registerToken(_tokenCode uint32, _tokenAddress address, _tokenDecimals uint8) returns()
func (_RenExTokens *RenExTokensSession) RegisterToken(_tokenCode uint32, _tokenAddress common.Address, _tokenDecimals uint8) (*types.Transaction, error) {
	return _RenExTokens.Contract.RegisterToken(&_RenExTokens.TransactOpts, _tokenCode, _tokenAddress, _tokenDecimals)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x9e20a9a0.
//
// Solidity: function registerToken(_tokenCode uint32, _tokenAddress address, _tokenDecimals uint8) returns()
func (_RenExTokens *RenExTokensTransactorSession) RegisterToken(_tokenCode uint32, _tokenAddress common.Address, _tokenDecimals uint8) (*types.Transaction, error) {
	return _RenExTokens.Contract.RegisterToken(&_RenExTokens.TransactOpts, _tokenCode, _tokenAddress, _tokenDecimals)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExTokens *RenExTokensTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExTokens *RenExTokensSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExTokens.Contract.RenounceOwnership(&_RenExTokens.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RenExTokens *RenExTokensTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RenExTokens.Contract.RenounceOwnership(&_RenExTokens.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExTokens *RenExTokensTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RenExTokens.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExTokens *RenExTokensSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExTokens.Contract.TransferOwnership(&_RenExTokens.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RenExTokens *RenExTokensTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RenExTokens.Contract.TransferOwnership(&_RenExTokens.TransactOpts, _newOwner)
}

// RenExTokensOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RenExTokens contract.
type RenExTokensOwnershipRenouncedIterator struct {
	Event *RenExTokensOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RenExTokensOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExTokensOwnershipRenounced)
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
		it.Event = new(RenExTokensOwnershipRenounced)
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
func (it *RenExTokensOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExTokensOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExTokensOwnershipRenounced represents a OwnershipRenounced event raised by the RenExTokens contract.
type RenExTokensOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RenExTokensOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExTokensOwnershipRenouncedIterator{contract: _RenExTokens.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RenExTokensOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExTokensOwnershipRenounced)
				if err := _RenExTokens.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RenExTokensOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RenExTokens contract.
type RenExTokensOwnershipTransferredIterator struct {
	Event *RenExTokensOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RenExTokensOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExTokensOwnershipTransferred)
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
		it.Event = new(RenExTokensOwnershipTransferred)
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
func (it *RenExTokensOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExTokensOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExTokensOwnershipTransferred represents a OwnershipTransferred event raised by the RenExTokens contract.
type RenExTokensOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RenExTokensOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RenExTokensOwnershipTransferredIterator{contract: _RenExTokens.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RenExTokens *RenExTokensFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RenExTokensOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RenExTokens.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExTokensOwnershipTransferred)
				if err := _RenExTokens.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RenExTokensTokenDeregisteredIterator is returned from FilterTokenDeregistered and is used to iterate over the raw logs and unpacked data for TokenDeregistered events raised by the RenExTokens contract.
type RenExTokensTokenDeregisteredIterator struct {
	Event *RenExTokensTokenDeregistered // Event containing the contract specifics and raw log

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
func (it *RenExTokensTokenDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExTokensTokenDeregistered)
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
		it.Event = new(RenExTokensTokenDeregistered)
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
func (it *RenExTokensTokenDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExTokensTokenDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExTokensTokenDeregistered represents a TokenDeregistered event raised by the RenExTokens contract.
type RenExTokensTokenDeregistered struct {
	TokenCode uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenDeregistered is a free log retrieval operation binding the contract event 0x370a83728af25aff893eba93c92cf098d1f4ebd31ecce8487f73edfa7db8a13b.
//
// Solidity: e TokenDeregistered(tokenCode uint32)
func (_RenExTokens *RenExTokensFilterer) FilterTokenDeregistered(opts *bind.FilterOpts) (*RenExTokensTokenDeregisteredIterator, error) {

	logs, sub, err := _RenExTokens.contract.FilterLogs(opts, "TokenDeregistered")
	if err != nil {
		return nil, err
	}
	return &RenExTokensTokenDeregisteredIterator{contract: _RenExTokens.contract, event: "TokenDeregistered", logs: logs, sub: sub}, nil
}

// WatchTokenDeregistered is a free log subscription operation binding the contract event 0x370a83728af25aff893eba93c92cf098d1f4ebd31ecce8487f73edfa7db8a13b.
//
// Solidity: e TokenDeregistered(tokenCode uint32)
func (_RenExTokens *RenExTokensFilterer) WatchTokenDeregistered(opts *bind.WatchOpts, sink chan<- *RenExTokensTokenDeregistered) (event.Subscription, error) {

	logs, sub, err := _RenExTokens.contract.WatchLogs(opts, "TokenDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExTokensTokenDeregistered)
				if err := _RenExTokens.contract.UnpackLog(event, "TokenDeregistered", log); err != nil {
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

// RenExTokensTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the RenExTokens contract.
type RenExTokensTokenRegisteredIterator struct {
	Event *RenExTokensTokenRegistered // Event containing the contract specifics and raw log

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
func (it *RenExTokensTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RenExTokensTokenRegistered)
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
		it.Event = new(RenExTokensTokenRegistered)
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
func (it *RenExTokensTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RenExTokensTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RenExTokensTokenRegistered represents a TokenRegistered event raised by the RenExTokens contract.
type RenExTokensTokenRegistered struct {
	TokenCode     uint32
	TokenAddress  common.Address
	TokenDecimals uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0xd6a4fb09bbf95930f0ad37c272f735e3453496dcdfc294f49d11d8d1fc371818.
//
// Solidity: e TokenRegistered(tokenCode uint32, tokenAddress address, tokenDecimals uint8)
func (_RenExTokens *RenExTokensFilterer) FilterTokenRegistered(opts *bind.FilterOpts) (*RenExTokensTokenRegisteredIterator, error) {

	logs, sub, err := _RenExTokens.contract.FilterLogs(opts, "TokenRegistered")
	if err != nil {
		return nil, err
	}
	return &RenExTokensTokenRegisteredIterator{contract: _RenExTokens.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0xd6a4fb09bbf95930f0ad37c272f735e3453496dcdfc294f49d11d8d1fc371818.
//
// Solidity: e TokenRegistered(tokenCode uint32, tokenAddress address, tokenDecimals uint8)
func (_RenExTokens *RenExTokensFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *RenExTokensTokenRegistered) (event.Subscription, error) {

	logs, sub, err := _RenExTokens.contract.WatchLogs(opts, "TokenRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RenExTokensTokenRegistered)
				if err := _RenExTokens.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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

// RepublicTokenABI is the input ABI used to generate the binding from.
const RepublicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// RepublicTokenBin is the compiled bytecode used for deploying new contracts.
const RepublicTokenBin = `0x60806040526003805460a060020a60ff021916905534801561002057600080fd5b5060038054600160a060020a031916339081179091556b033b2e3c9fd0803ce8000000600181905560009182526020829052604090912055610ea6806100676000396000f3006080604052600436106101115763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610116578063095ea7b3146101a057806318160ddd146101d857806323b872dd146101ff5780632ff2e9dc14610229578063313ce5671461023e5780633f4ba83a1461026957806342966c68146102805780635c975abb1461029857806366188463146102ad57806370a08231146102d1578063715018a6146102f25780638456cb59146103075780638da5cb5b1461031c57806395d89b411461034d578063a9059cbb14610362578063bec3fa1714610386578063d73dd623146103aa578063dd62ed3e146103ce578063f2fde38b146103f5575b600080fd5b34801561012257600080fd5b5061012b610416565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016557818101518382015260200161014d565b50505050905090810190601f1680156101925780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ac57600080fd5b506101c4600160a060020a036004351660243561044d565b604080519115158252519081900360200190f35b3480156101e457600080fd5b506101ed610478565b60408051918252519081900360200190f35b34801561020b57600080fd5b506101c4600160a060020a036004358116906024351660443561047e565b34801561023557600080fd5b506101ed6104c1565b34801561024a57600080fd5b506102536104d1565b6040805160ff9092168252519081900360200190f35b34801561027557600080fd5b5061027e6104d6565b005b34801561028c57600080fd5b5061027e60043561054e565b3480156102a457600080fd5b506101c461055b565b3480156102b957600080fd5b506101c4600160a060020a036004351660243561056b565b3480156102dd57600080fd5b506101ed600160a060020a036004351661058f565b3480156102fe57600080fd5b5061027e6105aa565b34801561031357600080fd5b5061027e610618565b34801561032857600080fd5b50610331610695565b60408051600160a060020a039092168252519081900360200190f35b34801561035957600080fd5b5061012b6106a4565b34801561036e57600080fd5b506101c4600160a060020a03600435166024356106db565b34801561039257600080fd5b506101c4600160a060020a0360043516602435610715565b3480156103b657600080fd5b506101c4600160a060020a03600435166024356107ed565b3480156103da57600080fd5b506101ed600160a060020a0360043581169060243516610811565b34801561040157600080fd5b5061027e600160a060020a036004351661083c565b60408051808201909152600e81527f52657075626c696320546f6b656e000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff161561046757600080fd5b610471838361085c565b9392505050565b60015490565b60035460009060a060020a900460ff161561049857600080fd5b600160a060020a0383163014156104ae57600080fd5b6104b98484846108c2565b949350505050565b6b033b2e3c9fd0803ce800000081565b601281565b600354600160a060020a031633146104ed57600080fd5b60035460a060020a900460ff16151561050557600080fd5b6003805474ff0000000000000000000000000000000000000000191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b61055833826108e7565b50565b60035460a060020a900460ff1681565b60035460009060a060020a900460ff161561058557600080fd5b61047183836109d6565b600160a060020a031660009081526020819052604090205490565b600354600160a060020a031633146105c157600080fd5b600354604051600160a060020a03909116907ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482090600090a26003805473ffffffffffffffffffffffffffffffffffffffff19169055565b600354600160a060020a0316331461062f57600080fd5b60035460a060020a900460ff161561064657600080fd5b6003805474ff0000000000000000000000000000000000000000191660a060020a1790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b600354600160a060020a031681565b60408051808201909152600381527f52454e0000000000000000000000000000000000000000000000000000000000602082015281565b60035460009060a060020a900460ff16156106f557600080fd5b600160a060020a03831630141561070b57600080fd5b6104718383610ac6565b600354600090600160a060020a0316331461072f57600080fd5b6000821161073c57600080fd5b600354600160a060020a0316600090815260208190526040902054610767908363ffffffff610aea16565b600354600160a060020a03908116600090815260208190526040808220939093559085168152205461079f908363ffffffff610afc16565b600160a060020a038085166000818152602081815260409182902094909455600354815187815291519294931692600080516020610e5b83398151915292918290030190a350600192915050565b60035460009060a060020a900460ff161561080757600080fd5b6104718383610b0f565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600160a060020a0316331461085357600080fd5b61055881610ba8565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60035460009060a060020a900460ff16156108dc57600080fd5b6104b9848484610c26565b600160a060020a03821660009081526020819052604090205481111561090c57600080fd5b600160a060020a038216600090815260208190526040902054610935908263ffffffff610aea16565b600160a060020a038316600090815260208190526040902055600154610961908263ffffffff610aea16565b600155604080518281529051600160a060020a038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518281529051600091600160a060020a03851691600080516020610e5b8339815191529181900360200190a35050565b336000908152600260209081526040808320600160a060020a038616845290915281205480831115610a2b57336000908152600260209081526040808320600160a060020a0388168452909152812055610a60565b610a3b818463ffffffff610aea16565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60035460009060a060020a900460ff1615610ae057600080fd5b6104718383610d8b565b600082821115610af657fe5b50900390565b81810182811015610b0957fe5b92915050565b336000908152600260209081526040808320600160a060020a0386168452909152812054610b43908363ffffffff610afc16565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a0381161515610bbd57600080fd5b600354604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000600160a060020a0383161515610c3d57600080fd5b600160a060020a038416600090815260208190526040902054821115610c6257600080fd5b600160a060020a0384166000908152600260209081526040808320338452909152902054821115610c9257600080fd5b600160a060020a038416600090815260208190526040902054610cbb908363ffffffff610aea16565b600160a060020a038086166000908152602081905260408082209390935590851681522054610cf0908363ffffffff610afc16565b600160a060020a03808516600090815260208181526040808320949094559187168152600282528281203382529091522054610d32908363ffffffff610aea16565b600160a060020a0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020610e5b833981519152929181900390910190a35060019392505050565b6000600160a060020a0383161515610da257600080fd5b33600090815260208190526040902054821115610dbe57600080fd5b33600090815260208190526040902054610dde908363ffffffff610aea16565b3360009081526020819052604080822092909255600160a060020a03851681522054610e10908363ffffffff610afc16565b600160a060020a03841660008181526020818152604091829020939093558051858152905191923392600080516020610e5b8339815191529281900390910190a3506001929150505600ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa165627a7a723058200c28a001c7b16a4f33d91c07654eb9f75f4ff8a0abf4807057596763060247fa0029`

// DeployRepublicToken deploys a new Ethereum contract, binding an instance of RepublicToken to it.
func DeployRepublicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RepublicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(RepublicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RepublicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RepublicToken{RepublicTokenCaller: RepublicTokenCaller{contract: contract}, RepublicTokenTransactor: RepublicTokenTransactor{contract: contract}, RepublicTokenFilterer: RepublicTokenFilterer{contract: contract}}, nil
}

// RepublicToken is an auto generated Go binding around an Ethereum contract.
type RepublicToken struct {
	RepublicTokenCaller     // Read-only binding to the contract
	RepublicTokenTransactor // Write-only binding to the contract
	RepublicTokenFilterer   // Log filterer for contract events
}

// RepublicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RepublicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepublicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RepublicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepublicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RepublicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RepublicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RepublicTokenSession struct {
	Contract     *RepublicToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RepublicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RepublicTokenCallerSession struct {
	Contract *RepublicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RepublicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RepublicTokenTransactorSession struct {
	Contract     *RepublicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RepublicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RepublicTokenRaw struct {
	Contract *RepublicToken // Generic contract binding to access the raw methods on
}

// RepublicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RepublicTokenCallerRaw struct {
	Contract *RepublicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RepublicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RepublicTokenTransactorRaw struct {
	Contract *RepublicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRepublicToken creates a new instance of RepublicToken, bound to a specific deployed contract.
func NewRepublicToken(address common.Address, backend bind.ContractBackend) (*RepublicToken, error) {
	contract, err := bindRepublicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RepublicToken{RepublicTokenCaller: RepublicTokenCaller{contract: contract}, RepublicTokenTransactor: RepublicTokenTransactor{contract: contract}, RepublicTokenFilterer: RepublicTokenFilterer{contract: contract}}, nil
}

// NewRepublicTokenCaller creates a new read-only instance of RepublicToken, bound to a specific deployed contract.
func NewRepublicTokenCaller(address common.Address, caller bind.ContractCaller) (*RepublicTokenCaller, error) {
	contract, err := bindRepublicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenCaller{contract: contract}, nil
}

// NewRepublicTokenTransactor creates a new write-only instance of RepublicToken, bound to a specific deployed contract.
func NewRepublicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RepublicTokenTransactor, error) {
	contract, err := bindRepublicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenTransactor{contract: contract}, nil
}

// NewRepublicTokenFilterer creates a new log filterer instance of RepublicToken, bound to a specific deployed contract.
func NewRepublicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RepublicTokenFilterer, error) {
	contract, err := bindRepublicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenFilterer{contract: contract}, nil
}

// bindRepublicToken binds a generic wrapper to an already deployed contract.
func bindRepublicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RepublicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepublicToken *RepublicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RepublicToken.Contract.RepublicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepublicToken *RepublicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepublicToken.Contract.RepublicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepublicToken *RepublicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepublicToken.Contract.RepublicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RepublicToken *RepublicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RepublicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RepublicToken *RepublicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepublicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RepublicToken *RepublicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RepublicToken.Contract.contract.Transact(opts, method, params...)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_RepublicToken *RepublicTokenCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_RepublicToken *RepublicTokenSession) INITIALSUPPLY() (*big.Int, error) {
	return _RepublicToken.Contract.INITIALSUPPLY(&_RepublicToken.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_RepublicToken *RepublicTokenCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _RepublicToken.Contract.INITIALSUPPLY(&_RepublicToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RepublicToken *RepublicTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RepublicToken *RepublicTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RepublicToken.Contract.Allowance(&_RepublicToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RepublicToken *RepublicTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RepublicToken.Contract.Allowance(&_RepublicToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RepublicToken *RepublicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RepublicToken *RepublicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _RepublicToken.Contract.BalanceOf(&_RepublicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RepublicToken *RepublicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _RepublicToken.Contract.BalanceOf(&_RepublicToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RepublicToken *RepublicTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RepublicToken *RepublicTokenSession) Decimals() (uint8, error) {
	return _RepublicToken.Contract.Decimals(&_RepublicToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RepublicToken *RepublicTokenCallerSession) Decimals() (uint8, error) {
	return _RepublicToken.Contract.Decimals(&_RepublicToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RepublicToken *RepublicTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RepublicToken *RepublicTokenSession) Name() (string, error) {
	return _RepublicToken.Contract.Name(&_RepublicToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RepublicToken *RepublicTokenCallerSession) Name() (string, error) {
	return _RepublicToken.Contract.Name(&_RepublicToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RepublicToken *RepublicTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RepublicToken *RepublicTokenSession) Owner() (common.Address, error) {
	return _RepublicToken.Contract.Owner(&_RepublicToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RepublicToken *RepublicTokenCallerSession) Owner() (common.Address, error) {
	return _RepublicToken.Contract.Owner(&_RepublicToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RepublicToken *RepublicTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RepublicToken *RepublicTokenSession) Paused() (bool, error) {
	return _RepublicToken.Contract.Paused(&_RepublicToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RepublicToken *RepublicTokenCallerSession) Paused() (bool, error) {
	return _RepublicToken.Contract.Paused(&_RepublicToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RepublicToken *RepublicTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RepublicToken *RepublicTokenSession) Symbol() (string, error) {
	return _RepublicToken.Contract.Symbol(&_RepublicToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RepublicToken *RepublicTokenCallerSession) Symbol() (string, error) {
	return _RepublicToken.Contract.Symbol(&_RepublicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RepublicToken *RepublicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RepublicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RepublicToken *RepublicTokenSession) TotalSupply() (*big.Int, error) {
	return _RepublicToken.Contract.TotalSupply(&_RepublicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RepublicToken *RepublicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _RepublicToken.Contract.TotalSupply(&_RepublicToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Approve(&_RepublicToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Approve(&_RepublicToken.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_RepublicToken *RepublicTokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_RepublicToken *RepublicTokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Burn(&_RepublicToken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_RepublicToken *RepublicTokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Burn(&_RepublicToken.TransactOpts, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.DecreaseApproval(&_RepublicToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.DecreaseApproval(&_RepublicToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.IncreaseApproval(&_RepublicToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_RepublicToken *RepublicTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.IncreaseApproval(&_RepublicToken.TransactOpts, _spender, _addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RepublicToken *RepublicTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RepublicToken *RepublicTokenSession) Pause() (*types.Transaction, error) {
	return _RepublicToken.Contract.Pause(&_RepublicToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RepublicToken *RepublicTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _RepublicToken.Contract.Pause(&_RepublicToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RepublicToken *RepublicTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RepublicToken *RepublicTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _RepublicToken.Contract.RenounceOwnership(&_RepublicToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RepublicToken *RepublicTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RepublicToken.Contract.RenounceOwnership(&_RepublicToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Transfer(&_RepublicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.Transfer(&_RepublicToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferFrom(&_RepublicToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferFrom(&_RepublicToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RepublicToken *RepublicTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RepublicToken *RepublicTokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferOwnership(&_RepublicToken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RepublicToken *RepublicTokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferOwnership(&_RepublicToken.TransactOpts, _newOwner)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xbec3fa17.
//
// Solidity: function transferTokens(beneficiary address, amount uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactor) TransferTokens(opts *bind.TransactOpts, beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "transferTokens", beneficiary, amount)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xbec3fa17.
//
// Solidity: function transferTokens(beneficiary address, amount uint256) returns(bool)
func (_RepublicToken *RepublicTokenSession) TransferTokens(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferTokens(&_RepublicToken.TransactOpts, beneficiary, amount)
}

// TransferTokens is a paid mutator transaction binding the contract method 0xbec3fa17.
//
// Solidity: function transferTokens(beneficiary address, amount uint256) returns(bool)
func (_RepublicToken *RepublicTokenTransactorSession) TransferTokens(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RepublicToken.Contract.TransferTokens(&_RepublicToken.TransactOpts, beneficiary, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RepublicToken *RepublicTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RepublicToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RepublicToken *RepublicTokenSession) Unpause() (*types.Transaction, error) {
	return _RepublicToken.Contract.Unpause(&_RepublicToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RepublicToken *RepublicTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _RepublicToken.Contract.Unpause(&_RepublicToken.TransactOpts)
}

// RepublicTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RepublicToken contract.
type RepublicTokenApprovalIterator struct {
	Event *RepublicTokenApproval // Event containing the contract specifics and raw log

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
func (it *RepublicTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenApproval)
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
		it.Event = new(RepublicTokenApproval)
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
func (it *RepublicTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenApproval represents a Approval event raised by the RepublicToken contract.
type RepublicTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RepublicTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenApprovalIterator{contract: _RepublicToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RepublicTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenApproval)
				if err := _RepublicToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// RepublicTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the RepublicToken contract.
type RepublicTokenBurnIterator struct {
	Event *RepublicTokenBurn // Event containing the contract specifics and raw log

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
func (it *RepublicTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenBurn)
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
		it.Event = new(RepublicTokenBurn)
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
func (it *RepublicTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenBurn represents a Burn event raised by the RepublicToken contract.
type RepublicTokenBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*RepublicTokenBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenBurnIterator{contract: _RepublicToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *RepublicTokenBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenBurn)
				if err := _RepublicToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// RepublicTokenOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RepublicToken contract.
type RepublicTokenOwnershipRenouncedIterator struct {
	Event *RepublicTokenOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RepublicTokenOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenOwnershipRenounced)
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
		it.Event = new(RepublicTokenOwnershipRenounced)
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
func (it *RepublicTokenOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenOwnershipRenounced represents a OwnershipRenounced event raised by the RepublicToken contract.
type RepublicTokenOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RepublicToken *RepublicTokenFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RepublicTokenOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenOwnershipRenouncedIterator{contract: _RepublicToken.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RepublicToken *RepublicTokenFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RepublicTokenOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenOwnershipRenounced)
				if err := _RepublicToken.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RepublicTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RepublicToken contract.
type RepublicTokenOwnershipTransferredIterator struct {
	Event *RepublicTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RepublicTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenOwnershipTransferred)
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
		it.Event = new(RepublicTokenOwnershipTransferred)
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
func (it *RepublicTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenOwnershipTransferred represents a OwnershipTransferred event raised by the RepublicToken contract.
type RepublicTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RepublicToken *RepublicTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RepublicTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenOwnershipTransferredIterator{contract: _RepublicToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RepublicToken *RepublicTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RepublicTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenOwnershipTransferred)
				if err := _RepublicToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RepublicTokenPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the RepublicToken contract.
type RepublicTokenPauseIterator struct {
	Event *RepublicTokenPause // Event containing the contract specifics and raw log

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
func (it *RepublicTokenPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenPause)
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
		it.Event = new(RepublicTokenPause)
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
func (it *RepublicTokenPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenPause represents a Pause event raised by the RepublicToken contract.
type RepublicTokenPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_RepublicToken *RepublicTokenFilterer) FilterPause(opts *bind.FilterOpts) (*RepublicTokenPauseIterator, error) {

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &RepublicTokenPauseIterator{contract: _RepublicToken.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_RepublicToken *RepublicTokenFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *RepublicTokenPause) (event.Subscription, error) {

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenPause)
				if err := _RepublicToken.contract.UnpackLog(event, "Pause", log); err != nil {
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

// RepublicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RepublicToken contract.
type RepublicTokenTransferIterator struct {
	Event *RepublicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RepublicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenTransfer)
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
		it.Event = new(RepublicTokenTransfer)
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
func (it *RepublicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenTransfer represents a Transfer event raised by the RepublicToken contract.
type RepublicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RepublicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RepublicTokenTransferIterator{contract: _RepublicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_RepublicToken *RepublicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RepublicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenTransfer)
				if err := _RepublicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// RepublicTokenUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the RepublicToken contract.
type RepublicTokenUnpauseIterator struct {
	Event *RepublicTokenUnpause // Event containing the contract specifics and raw log

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
func (it *RepublicTokenUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RepublicTokenUnpause)
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
		it.Event = new(RepublicTokenUnpause)
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
func (it *RepublicTokenUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RepublicTokenUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RepublicTokenUnpause represents a Unpause event raised by the RepublicToken contract.
type RepublicTokenUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_RepublicToken *RepublicTokenFilterer) FilterUnpause(opts *bind.FilterOpts) (*RepublicTokenUnpauseIterator, error) {

	logs, sub, err := _RepublicToken.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &RepublicTokenUnpauseIterator{contract: _RepublicToken.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_RepublicToken *RepublicTokenFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *RepublicTokenUnpause) (event.Subscription, error) {

	logs, sub, err := _RepublicToken.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RepublicTokenUnpause)
				if err := _RepublicToken.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// RewardVaultABI is the input ABI used to generate the binding from.
const RewardVaultABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"darknodeBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_darknode\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"darknodeRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDarknodeRegistry\",\"type\":\"address\"}],\"name\":\"updateDarknodeRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETHEREUM\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_darknode\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_darknodeRegistry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousDarknodeRegistry\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"nextDarknodeRegistry\",\"type\":\"address\"}],\"name\":\"DarknodeRegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RewardVaultBin is the compiled bytecode used for deploying new contracts.
const RewardVaultBin = `0x608060405234801561001057600080fd5b506040516020806107a5833981016040525160008054600160a060020a0319908116331790915560018054600160a060020a03909316929091169190911790556107468061005f6000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166370324b77811461009d578063715018a6146100d65780638340f549146100ed5780638da5cb5b1461010a5780639e45e0d01461013b578063aaff096d14610150578063f2fde38b14610171578063f7cdf47c14610192578063f940e385146101a7575b600080fd5b3480156100a957600080fd5b506100c4600160a060020a03600435811690602435166101ce565b60408051918252519081900360200190f35b3480156100e257600080fd5b506100eb6101eb565b005b6100eb600160a060020a0360043581169060243516604435610257565b34801561011657600080fd5b5061011f610399565b60408051600160a060020a039092168252519081900360200190f35b34801561014757600080fd5b5061011f6103a8565b34801561015c57600080fd5b506100eb600160a060020a03600435166103b7565b34801561017d57600080fd5b506100eb600160a060020a0360043516610445565b34801561019e57600080fd5b5061011f610468565b3480156101b357600080fd5b506100eb600160a060020a0360043581169060243516610480565b600260209081526000928352604080842090915290825290205481565b600054600160a060020a0316331461020257600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600160a060020a03821673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee141561028d5734811461028857600080fd5b610332565b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018390529051600160a060020a038416916323b872dd9160648083019260209291908290030181600087803b1580156102fb57600080fd5b505af115801561030f573d6000803e3d6000fd5b505050506040513d602081101561032557600080fd5b5051151561033257600080fd5b600160a060020a03808416600090815260026020908152604080832093861683529290522054610368908263ffffffff61068a16565b600160a060020a03938416600090815260026020908152604080832095909616825293909352929091209190915550565b600054600160a060020a031681565b600154600160a060020a031681565b600054600160a060020a031633146103ce57600080fd5b60015460408051600160a060020a039283168152918316602083015280517fd4443dde7b09a100132dba56b0175378b6ad81b64398aef77626357a1d8daf029281900390910190a16001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a0316331461045c57600080fd5b6104658161069d565b50565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b600154604080517f424fe57b0000000000000000000000000000000000000000000000000000000081526bffffffffffffffffffffffff196c01000000000000000000000000860216600482015290516000928392600160a060020a039091169163424fe57b9160248082019260209290919082900301818787803b15801561050857600080fd5b505af115801561051c573d6000803e3d6000fd5b505050506040513d602081101561053257600080fd5b50519150600160a060020a038216151561054b57600080fd5b50600160a060020a03838116600090815260026020908152604080832093861680845293909152812080549190559073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee14156105d157604051600160a060020a0383169082156108fc029083906000818181858888f193505050501580156105cb573d6000803e3d6000fd5b50610684565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561064d57600080fd5b505af1158015610661573d6000803e3d6000fd5b505050506040513d602081101561067757600080fd5b5051151561068457600080fd5b50505050565b8181018281101561069757fe5b92915050565b600160a060020a03811615156106b257600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820d8196a8eecf9b46ee85e5524a540a03a3fa6323babba1c7780cf160136fa345f0029`

// DeployRewardVault deploys a new Ethereum contract, binding an instance of RewardVault to it.
func DeployRewardVault(auth *bind.TransactOpts, backend bind.ContractBackend, _darknodeRegistry common.Address) (common.Address, *types.Transaction, *RewardVault, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardVaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RewardVaultBin), backend, _darknodeRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RewardVault{RewardVaultCaller: RewardVaultCaller{contract: contract}, RewardVaultTransactor: RewardVaultTransactor{contract: contract}, RewardVaultFilterer: RewardVaultFilterer{contract: contract}}, nil
}

// RewardVault is an auto generated Go binding around an Ethereum contract.
type RewardVault struct {
	RewardVaultCaller     // Read-only binding to the contract
	RewardVaultTransactor // Write-only binding to the contract
	RewardVaultFilterer   // Log filterer for contract events
}

// RewardVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardVaultSession struct {
	Contract     *RewardVault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardVaultCallerSession struct {
	Contract *RewardVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RewardVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardVaultTransactorSession struct {
	Contract     *RewardVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RewardVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardVaultRaw struct {
	Contract *RewardVault // Generic contract binding to access the raw methods on
}

// RewardVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardVaultCallerRaw struct {
	Contract *RewardVaultCaller // Generic read-only contract binding to access the raw methods on
}

// RewardVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardVaultTransactorRaw struct {
	Contract *RewardVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardVault creates a new instance of RewardVault, bound to a specific deployed contract.
func NewRewardVault(address common.Address, backend bind.ContractBackend) (*RewardVault, error) {
	contract, err := bindRewardVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardVault{RewardVaultCaller: RewardVaultCaller{contract: contract}, RewardVaultTransactor: RewardVaultTransactor{contract: contract}, RewardVaultFilterer: RewardVaultFilterer{contract: contract}}, nil
}

// NewRewardVaultCaller creates a new read-only instance of RewardVault, bound to a specific deployed contract.
func NewRewardVaultCaller(address common.Address, caller bind.ContractCaller) (*RewardVaultCaller, error) {
	contract, err := bindRewardVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardVaultCaller{contract: contract}, nil
}

// NewRewardVaultTransactor creates a new write-only instance of RewardVault, bound to a specific deployed contract.
func NewRewardVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardVaultTransactor, error) {
	contract, err := bindRewardVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardVaultTransactor{contract: contract}, nil
}

// NewRewardVaultFilterer creates a new log filterer instance of RewardVault, bound to a specific deployed contract.
func NewRewardVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardVaultFilterer, error) {
	contract, err := bindRewardVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardVaultFilterer{contract: contract}, nil
}

// bindRewardVault binds a generic wrapper to an already deployed contract.
func bindRewardVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardVault *RewardVaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RewardVault.Contract.RewardVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardVault *RewardVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardVault.Contract.RewardVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardVault *RewardVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardVault.Contract.RewardVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardVault *RewardVaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RewardVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardVault *RewardVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardVault *RewardVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardVault.Contract.contract.Transact(opts, method, params...)
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_RewardVault *RewardVaultCaller) ETHEREUM(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RewardVault.contract.Call(opts, out, "ETHEREUM")
	return *ret0, err
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_RewardVault *RewardVaultSession) ETHEREUM() (common.Address, error) {
	return _RewardVault.Contract.ETHEREUM(&_RewardVault.CallOpts)
}

// ETHEREUM is a free data retrieval call binding the contract method 0xf7cdf47c.
//
// Solidity: function ETHEREUM() constant returns(address)
func (_RewardVault *RewardVaultCallerSession) ETHEREUM() (common.Address, error) {
	return _RewardVault.Contract.ETHEREUM(&_RewardVault.CallOpts)
}

// DarknodeBalances is a free data retrieval call binding the contract method 0x70324b77.
//
// Solidity: function darknodeBalances( address,  address) constant returns(uint256)
func (_RewardVault *RewardVaultCaller) DarknodeBalances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RewardVault.contract.Call(opts, out, "darknodeBalances", arg0, arg1)
	return *ret0, err
}

// DarknodeBalances is a free data retrieval call binding the contract method 0x70324b77.
//
// Solidity: function darknodeBalances( address,  address) constant returns(uint256)
func (_RewardVault *RewardVaultSession) DarknodeBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RewardVault.Contract.DarknodeBalances(&_RewardVault.CallOpts, arg0, arg1)
}

// DarknodeBalances is a free data retrieval call binding the contract method 0x70324b77.
//
// Solidity: function darknodeBalances( address,  address) constant returns(uint256)
func (_RewardVault *RewardVaultCallerSession) DarknodeBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RewardVault.Contract.DarknodeBalances(&_RewardVault.CallOpts, arg0, arg1)
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_RewardVault *RewardVaultCaller) DarknodeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RewardVault.contract.Call(opts, out, "darknodeRegistry")
	return *ret0, err
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_RewardVault *RewardVaultSession) DarknodeRegistry() (common.Address, error) {
	return _RewardVault.Contract.DarknodeRegistry(&_RewardVault.CallOpts)
}

// DarknodeRegistry is a free data retrieval call binding the contract method 0x9e45e0d0.
//
// Solidity: function darknodeRegistry() constant returns(address)
func (_RewardVault *RewardVaultCallerSession) DarknodeRegistry() (common.Address, error) {
	return _RewardVault.Contract.DarknodeRegistry(&_RewardVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RewardVault *RewardVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RewardVault.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RewardVault *RewardVaultSession) Owner() (common.Address, error) {
	return _RewardVault.Contract.Owner(&_RewardVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RewardVault *RewardVaultCallerSession) Owner() (common.Address, error) {
	return _RewardVault.Contract.Owner(&_RewardVault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(_darknode address, _token address, _value uint256) returns()
func (_RewardVault *RewardVaultTransactor) Deposit(opts *bind.TransactOpts, _darknode common.Address, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RewardVault.contract.Transact(opts, "deposit", _darknode, _token, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(_darknode address, _token address, _value uint256) returns()
func (_RewardVault *RewardVaultSession) Deposit(_darknode common.Address, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RewardVault.Contract.Deposit(&_RewardVault.TransactOpts, _darknode, _token, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(_darknode address, _token address, _value uint256) returns()
func (_RewardVault *RewardVaultTransactorSession) Deposit(_darknode common.Address, _token common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RewardVault.Contract.Deposit(&_RewardVault.TransactOpts, _darknode, _token, _value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardVault *RewardVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardVault *RewardVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _RewardVault.Contract.RenounceOwnership(&_RewardVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardVault *RewardVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RewardVault.Contract.RenounceOwnership(&_RewardVault.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RewardVault *RewardVaultTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RewardVault.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RewardVault *RewardVaultSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RewardVault.Contract.TransferOwnership(&_RewardVault.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RewardVault *RewardVaultTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RewardVault.Contract.TransferOwnership(&_RewardVault.TransactOpts, _newOwner)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_RewardVault *RewardVaultTransactor) UpdateDarknodeRegistry(opts *bind.TransactOpts, _newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _RewardVault.contract.Transact(opts, "updateDarknodeRegistry", _newDarknodeRegistry)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_RewardVault *RewardVaultSession) UpdateDarknodeRegistry(_newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _RewardVault.Contract.UpdateDarknodeRegistry(&_RewardVault.TransactOpts, _newDarknodeRegistry)
}

// UpdateDarknodeRegistry is a paid mutator transaction binding the contract method 0xaaff096d.
//
// Solidity: function updateDarknodeRegistry(_newDarknodeRegistry address) returns()
func (_RewardVault *RewardVaultTransactorSession) UpdateDarknodeRegistry(_newDarknodeRegistry common.Address) (*types.Transaction, error) {
	return _RewardVault.Contract.UpdateDarknodeRegistry(&_RewardVault.TransactOpts, _newDarknodeRegistry)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(_darknode address, _token address) returns()
func (_RewardVault *RewardVaultTransactor) Withdraw(opts *bind.TransactOpts, _darknode common.Address, _token common.Address) (*types.Transaction, error) {
	return _RewardVault.contract.Transact(opts, "withdraw", _darknode, _token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(_darknode address, _token address) returns()
func (_RewardVault *RewardVaultSession) Withdraw(_darknode common.Address, _token common.Address) (*types.Transaction, error) {
	return _RewardVault.Contract.Withdraw(&_RewardVault.TransactOpts, _darknode, _token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf940e385.
//
// Solidity: function withdraw(_darknode address, _token address) returns()
func (_RewardVault *RewardVaultTransactorSession) Withdraw(_darknode common.Address, _token common.Address) (*types.Transaction, error) {
	return _RewardVault.Contract.Withdraw(&_RewardVault.TransactOpts, _darknode, _token)
}

// RewardVaultDarknodeRegistryUpdatedIterator is returned from FilterDarknodeRegistryUpdated and is used to iterate over the raw logs and unpacked data for DarknodeRegistryUpdated events raised by the RewardVault contract.
type RewardVaultDarknodeRegistryUpdatedIterator struct {
	Event *RewardVaultDarknodeRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *RewardVaultDarknodeRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardVaultDarknodeRegistryUpdated)
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
		it.Event = new(RewardVaultDarknodeRegistryUpdated)
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
func (it *RewardVaultDarknodeRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardVaultDarknodeRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardVaultDarknodeRegistryUpdated represents a DarknodeRegistryUpdated event raised by the RewardVault contract.
type RewardVaultDarknodeRegistryUpdated struct {
	PreviousDarknodeRegistry common.Address
	NextDarknodeRegistry     common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterDarknodeRegistryUpdated is a free log retrieval operation binding the contract event 0xd4443dde7b09a100132dba56b0175378b6ad81b64398aef77626357a1d8daf02.
//
// Solidity: e DarknodeRegistryUpdated(previousDarknodeRegistry address, nextDarknodeRegistry address)
func (_RewardVault *RewardVaultFilterer) FilterDarknodeRegistryUpdated(opts *bind.FilterOpts) (*RewardVaultDarknodeRegistryUpdatedIterator, error) {

	logs, sub, err := _RewardVault.contract.FilterLogs(opts, "DarknodeRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &RewardVaultDarknodeRegistryUpdatedIterator{contract: _RewardVault.contract, event: "DarknodeRegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchDarknodeRegistryUpdated is a free log subscription operation binding the contract event 0xd4443dde7b09a100132dba56b0175378b6ad81b64398aef77626357a1d8daf02.
//
// Solidity: e DarknodeRegistryUpdated(previousDarknodeRegistry address, nextDarknodeRegistry address)
func (_RewardVault *RewardVaultFilterer) WatchDarknodeRegistryUpdated(opts *bind.WatchOpts, sink chan<- *RewardVaultDarknodeRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _RewardVault.contract.WatchLogs(opts, "DarknodeRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardVaultDarknodeRegistryUpdated)
				if err := _RewardVault.contract.UnpackLog(event, "DarknodeRegistryUpdated", log); err != nil {
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

// RewardVaultOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RewardVault contract.
type RewardVaultOwnershipRenouncedIterator struct {
	Event *RewardVaultOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RewardVaultOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardVaultOwnershipRenounced)
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
		it.Event = new(RewardVaultOwnershipRenounced)
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
func (it *RewardVaultOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardVaultOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardVaultOwnershipRenounced represents a OwnershipRenounced event raised by the RewardVault contract.
type RewardVaultOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RewardVault *RewardVaultFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RewardVaultOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RewardVault.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewardVaultOwnershipRenouncedIterator{contract: _RewardVault.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RewardVault *RewardVaultFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RewardVaultOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RewardVault.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardVaultOwnershipRenounced)
				if err := _RewardVault.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RewardVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RewardVault contract.
type RewardVaultOwnershipTransferredIterator struct {
	Event *RewardVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewardVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardVaultOwnershipTransferred)
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
		it.Event = new(RewardVaultOwnershipTransferred)
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
func (it *RewardVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardVaultOwnershipTransferred represents a OwnershipTransferred event raised by the RewardVault contract.
type RewardVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RewardVault *RewardVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewardVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewardVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewardVaultOwnershipTransferredIterator{contract: _RewardVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RewardVault *RewardVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewardVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewardVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardVaultOwnershipTransferred)
				if err := _RewardVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582079a38b9223e0b1ede7697d25979b7764c9392cc051d1444b83223686a70dd9ef0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x608060405234801561001057600080fd5b506106b3806100206000396000f30060806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009257806318160ddd146100ca57806323b872dd146100f1578063661884631461011b57806370a082311461013f578063a9059cbb14610160578063d73dd62314610184578063dd62ed3e146101a8575b600080fd5b34801561009e57600080fd5b506100b6600160a060020a03600435166024356101cf565b604080519115158252519081900360200190f35b3480156100d657600080fd5b506100df610235565b60408051918252519081900360200190f35b3480156100fd57600080fd5b506100b6600160a060020a036004358116906024351660443561023b565b34801561012757600080fd5b506100b6600160a060020a03600435166024356103b2565b34801561014b57600080fd5b506100df600160a060020a03600435166104a2565b34801561016c57600080fd5b506100b6600160a060020a03600435166024356104bd565b34801561019057600080fd5b506100b6600160a060020a036004351660243561059e565b3480156101b457600080fd5b506100df600160a060020a0360043581169060243516610637565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60015490565b6000600160a060020a038316151561025257600080fd5b600160a060020a03841660009081526020819052604090205482111561027757600080fd5b600160a060020a03841660009081526002602090815260408083203384529091529020548211156102a757600080fd5b600160a060020a0384166000908152602081905260409020546102d0908363ffffffff61066216565b600160a060020a038086166000908152602081905260408082209390935590851681522054610305908363ffffffff61067416565b600160a060020a03808516600090815260208181526040808320949094559187168152600282528281203382529091522054610347908363ffffffff61066216565b600160a060020a03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b336000908152600260209081526040808320600160a060020a03861684529091528120548083111561040757336000908152600260209081526040808320600160a060020a038816845290915281205561043c565b610417818463ffffffff61066216565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526020819052604090205490565b6000600160a060020a03831615156104d457600080fd5b336000908152602081905260409020548211156104f057600080fd5b33600090815260208190526040902054610510908363ffffffff61066216565b3360009081526020819052604080822092909255600160a060020a03851681522054610542908363ffffffff61067416565b600160a060020a038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120546105d2908363ffffffff61067416565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561066e57fe5b50900390565b8181018281101561068157fe5b929150505600a165627a7a72305820a848098d8d6b2b51b437f56383eea6d4c66fdbc499bd6cc6bdf762be4d7d21c20029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

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
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
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
		it.Event = new(StandardTokenApproval)
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
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
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
		it.Event = new(StandardTokenTransfer)
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
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
const UtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058209541f35ac6e52b1b5e65f28389eabdae95eb1a344a2af45c314ed06a221d3b160029`

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}
