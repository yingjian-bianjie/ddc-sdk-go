// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// DDC721MetaData contains all meta data concerning the DDC721 contract.
var DDC721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"EnterBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"ExitBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"Locklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"MetaTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ddcIds\",\"type\":\"uint256[]\"}],\"name\":\"MetaTransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"SetNameAndSymbol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ddcURI\",\"type\":\"string\"}],\"name\":\"SetURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ddcIds\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"UnLocklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"ddcURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestDDCId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"metaBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ddcURI_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"metaMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"ddcURIs\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"metaMintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ddcURI_\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"metaSafeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"ddcURIs\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"metaSafeMintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"metaSafeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"metaTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ddcURI_\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"ddcURIs\",\"type\":\"string[]\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ddcURI\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"ddcURIs\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeMintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorityProxyAddress\",\"type\":\"address\"}],\"name\":\"setAuthorityProxyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"chargeProxyAddress\",\"type\":\"address\"}],\"name\":\"setChargeProxyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"separator\",\"type\":\"bytes32\"}],\"name\":\"setMetaSeparatorArg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIDDC721.HashType\",\"name\":\"hashType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"hashValue\",\"type\":\"bytes32\"}],\"name\":\"setMetaTypeHashArgs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"setNameAndSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ddcURI_\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"unFreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// DDC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use DDC721MetaData.ABI instead.
var DDC721ABI = DDC721MetaData.ABI

// DDC721 is an auto generated Go binding around an Ethereum contract.
type DDC721 struct {
	DDC721Caller     // Read-only binding to the contract
	DDC721Transactor // Write-only binding to the contract
	DDC721Filterer   // Log filterer for contract events
}

// DDC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type DDC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DDC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DDC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DDC721Session struct {
	Contract     *DDC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DDC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DDC721CallerSession struct {
	Contract *DDC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DDC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DDC721TransactorSession struct {
	Contract     *DDC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DDC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type DDC721Raw struct {
	Contract *DDC721 // Generic contract binding to access the raw methods on
}

// DDC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DDC721CallerRaw struct {
	Contract *DDC721Caller // Generic read-only contract binding to access the raw methods on
}

// DDC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DDC721TransactorRaw struct {
	Contract *DDC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDDC721 creates a new instance of DDC721, bound to a specific deployed contract.
func NewDDC721(address common.Address, backend bind.ContractBackend) (*DDC721, error) {
	contract, err := bindDDC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DDC721{DDC721Caller: DDC721Caller{contract: contract}, DDC721Transactor: DDC721Transactor{contract: contract}, DDC721Filterer: DDC721Filterer{contract: contract}}, nil
}

// NewDDC721Caller creates a new read-only instance of DDC721, bound to a specific deployed contract.
func NewDDC721Caller(address common.Address, caller bind.ContractCaller) (*DDC721Caller, error) {
	contract, err := bindDDC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DDC721Caller{contract: contract}, nil
}

// NewDDC721Transactor creates a new write-only instance of DDC721, bound to a specific deployed contract.
func NewDDC721Transactor(address common.Address, transactor bind.ContractTransactor) (*DDC721Transactor, error) {
	contract, err := bindDDC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DDC721Transactor{contract: contract}, nil
}

// NewDDC721Filterer creates a new log filterer instance of DDC721, bound to a specific deployed contract.
func NewDDC721Filterer(address common.Address, filterer bind.ContractFilterer) (*DDC721Filterer, error) {
	contract, err := bindDDC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DDC721Filterer{contract: contract}, nil
}

// bindDDC721 binds a generic wrapper to an already deployed contract.
func bindDDC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DDC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DDC721 *DDC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DDC721.Contract.DDC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DDC721 *DDC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC721.Contract.DDC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DDC721 *DDC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DDC721.Contract.DDC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DDC721 *DDC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DDC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DDC721 *DDC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DDC721 *DDC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DDC721.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DDC721 *DDC721Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DDC721 *DDC721Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _DDC721.Contract.DOMAINSEPARATOR(&_DDC721.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DDC721 *DDC721CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _DDC721.Contract.DOMAINSEPARATOR(&_DDC721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DDC721 *DDC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DDC721 *DDC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DDC721.Contract.BalanceOf(&_DDC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DDC721 *DDC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DDC721.Contract.BalanceOf(&_DDC721.CallOpts, owner)
}

// DdcURI is a free data retrieval call binding the contract method 0x293ec97c.
//
// Solidity: function ddcURI(uint256 ddcId) view returns(string)
func (_DDC721 *DDC721Caller) DdcURI(opts *bind.CallOpts, ddcId *big.Int) (string, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "ddcURI", ddcId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DdcURI is a free data retrieval call binding the contract method 0x293ec97c.
//
// Solidity: function ddcURI(uint256 ddcId) view returns(string)
func (_DDC721 *DDC721Session) DdcURI(ddcId *big.Int) (string, error) {
	return _DDC721.Contract.DdcURI(&_DDC721.CallOpts, ddcId)
}

// DdcURI is a free data retrieval call binding the contract method 0x293ec97c.
//
// Solidity: function ddcURI(uint256 ddcId) view returns(string)
func (_DDC721 *DDC721CallerSession) DdcURI(ddcId *big.Int) (string, error) {
	return _DDC721.Contract.DdcURI(&_DDC721.CallOpts, ddcId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ddcId) view returns(address)
func (_DDC721 *DDC721Caller) GetApproved(opts *bind.CallOpts, ddcId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "getApproved", ddcId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ddcId) view returns(address)
func (_DDC721 *DDC721Session) GetApproved(ddcId *big.Int) (common.Address, error) {
	return _DDC721.Contract.GetApproved(&_DDC721.CallOpts, ddcId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ddcId) view returns(address)
func (_DDC721 *DDC721CallerSession) GetApproved(ddcId *big.Int) (common.Address, error) {
	return _DDC721.Contract.GetApproved(&_DDC721.CallOpts, ddcId)
}

// GetLatestDDCId is a free data retrieval call binding the contract method 0xb9f7fad9.
//
// Solidity: function getLatestDDCId() view returns(uint256)
func (_DDC721 *DDC721Caller) GetLatestDDCId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "getLatestDDCId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestDDCId is a free data retrieval call binding the contract method 0xb9f7fad9.
//
// Solidity: function getLatestDDCId() view returns(uint256)
func (_DDC721 *DDC721Session) GetLatestDDCId() (*big.Int, error) {
	return _DDC721.Contract.GetLatestDDCId(&_DDC721.CallOpts)
}

// GetLatestDDCId is a free data retrieval call binding the contract method 0xb9f7fad9.
//
// Solidity: function getLatestDDCId() view returns(uint256)
func (_DDC721 *DDC721CallerSession) GetLatestDDCId() (*big.Int, error) {
	return _DDC721.Contract.GetLatestDDCId(&_DDC721.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_DDC721 *DDC721Caller) GetNonce(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "getNonce", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_DDC721 *DDC721Session) GetNonce(from common.Address) (*big.Int, error) {
	return _DDC721.Contract.GetNonce(&_DDC721.CallOpts, from)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address from) view returns(uint256)
func (_DDC721 *DDC721CallerSession) GetNonce(from common.Address) (*big.Int, error) {
	return _DDC721.Contract.GetNonce(&_DDC721.CallOpts, from)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DDC721 *DDC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DDC721 *DDC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DDC721.Contract.IsApprovedForAll(&_DDC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DDC721 *DDC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DDC721.Contract.IsApprovedForAll(&_DDC721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DDC721 *DDC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DDC721 *DDC721Session) Name() (string, error) {
	return _DDC721.Contract.Name(&_DDC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DDC721 *DDC721CallerSession) Name() (string, error) {
	return _DDC721.Contract.Name(&_DDC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDC721 *DDC721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDC721 *DDC721Session) Owner() (common.Address, error) {
	return _DDC721.Contract.Owner(&_DDC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DDC721 *DDC721CallerSession) Owner() (common.Address, error) {
	return _DDC721.Contract.Owner(&_DDC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 ddcId) view returns(address)
func (_DDC721 *DDC721Caller) OwnerOf(opts *bind.CallOpts, ddcId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "ownerOf", ddcId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 ddcId) view returns(address)
func (_DDC721 *DDC721Session) OwnerOf(ddcId *big.Int) (common.Address, error) {
	return _DDC721.Contract.OwnerOf(&_DDC721.CallOpts, ddcId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 ddcId) view returns(address)
func (_DDC721 *DDC721CallerSession) OwnerOf(ddcId *big.Int) (common.Address, error) {
	return _DDC721.Contract.OwnerOf(&_DDC721.CallOpts, ddcId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DDC721 *DDC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DDC721 *DDC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DDC721.Contract.SupportsInterface(&_DDC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DDC721 *DDC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DDC721.Contract.SupportsInterface(&_DDC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DDC721 *DDC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DDC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DDC721 *DDC721Session) Symbol() (string, error) {
	return _DDC721.Contract.Symbol(&_DDC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DDC721 *DDC721CallerSession) Symbol() (string, error) {
	return _DDC721.Contract.Symbol(&_DDC721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 ddcId) returns()
func (_DDC721 *DDC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "approve", to, ddcId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 ddcId) returns()
func (_DDC721 *DDC721Session) Approve(to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.Approve(&_DDC721.TransactOpts, to, ddcId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 ddcId) returns()
func (_DDC721 *DDC721TransactorSession) Approve(to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.Approve(&_DDC721.TransactOpts, to, ddcId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 ddcId) returns()
func (_DDC721 *DDC721Transactor) Burn(opts *bind.TransactOpts, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "burn", ddcId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 ddcId) returns()
func (_DDC721 *DDC721Session) Burn(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.Burn(&_DDC721.TransactOpts, ddcId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 ddcId) returns()
func (_DDC721 *DDC721TransactorSession) Burn(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.Burn(&_DDC721.TransactOpts, ddcId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 ddcId) returns()
func (_DDC721 *DDC721Transactor) Freeze(opts *bind.TransactOpts, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "freeze", ddcId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 ddcId) returns()
func (_DDC721 *DDC721Session) Freeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.Freeze(&_DDC721.TransactOpts, ddcId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 ddcId) returns()
func (_DDC721 *DDC721TransactorSession) Freeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.Freeze(&_DDC721.TransactOpts, ddcId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DDC721 *DDC721Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DDC721 *DDC721Session) Initialize() (*types.Transaction, error) {
	return _DDC721.Contract.Initialize(&_DDC721.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DDC721 *DDC721TransactorSession) Initialize() (*types.Transaction, error) {
	return _DDC721.Contract.Initialize(&_DDC721.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 ddcId) returns()
func (_DDC721 *DDC721Transactor) Lock(opts *bind.TransactOpts, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "lock", ddcId)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 ddcId) returns()
func (_DDC721 *DDC721Session) Lock(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.Lock(&_DDC721.TransactOpts, ddcId)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 ddcId) returns()
func (_DDC721 *DDC721TransactorSession) Lock(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.Lock(&_DDC721.TransactOpts, ddcId)
}

// MetaBurn is a paid mutator transaction binding the contract method 0x60c568a6.
//
// Solidity: function metaBurn(uint256 ddcId, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Transactor) MetaBurn(opts *bind.TransactOpts, ddcId *big.Int, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "metaBurn", ddcId, nonce, deadline, sign)
}

// MetaBurn is a paid mutator transaction binding the contract method 0x60c568a6.
//
// Solidity: function metaBurn(uint256 ddcId, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Session) MetaBurn(ddcId *big.Int, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaBurn(&_DDC721.TransactOpts, ddcId, nonce, deadline, sign)
}

// MetaBurn is a paid mutator transaction binding the contract method 0x60c568a6.
//
// Solidity: function metaBurn(uint256 ddcId, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721TransactorSession) MetaBurn(ddcId *big.Int, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaBurn(&_DDC721.TransactOpts, ddcId, nonce, deadline, sign)
}

// MetaMint is a paid mutator transaction binding the contract method 0xb6852b71.
//
// Solidity: function metaMint(address to, string ddcURI_, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Transactor) MetaMint(opts *bind.TransactOpts, to common.Address, ddcURI_ string, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "metaMint", to, ddcURI_, nonce, deadline, sign)
}

// MetaMint is a paid mutator transaction binding the contract method 0xb6852b71.
//
// Solidity: function metaMint(address to, string ddcURI_, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Session) MetaMint(to common.Address, ddcURI_ string, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaMint(&_DDC721.TransactOpts, to, ddcURI_, nonce, deadline, sign)
}

// MetaMint is a paid mutator transaction binding the contract method 0xb6852b71.
//
// Solidity: function metaMint(address to, string ddcURI_, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721TransactorSession) MetaMint(to common.Address, ddcURI_ string, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaMint(&_DDC721.TransactOpts, to, ddcURI_, nonce, deadline, sign)
}

// MetaMintBatch is a paid mutator transaction binding the contract method 0x25c016da.
//
// Solidity: function metaMintBatch(address to, string[] ddcURIs, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Transactor) MetaMintBatch(opts *bind.TransactOpts, to common.Address, ddcURIs []string, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "metaMintBatch", to, ddcURIs, nonce, deadline, sign)
}

// MetaMintBatch is a paid mutator transaction binding the contract method 0x25c016da.
//
// Solidity: function metaMintBatch(address to, string[] ddcURIs, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Session) MetaMintBatch(to common.Address, ddcURIs []string, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaMintBatch(&_DDC721.TransactOpts, to, ddcURIs, nonce, deadline, sign)
}

// MetaMintBatch is a paid mutator transaction binding the contract method 0x25c016da.
//
// Solidity: function metaMintBatch(address to, string[] ddcURIs, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721TransactorSession) MetaMintBatch(to common.Address, ddcURIs []string, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaMintBatch(&_DDC721.TransactOpts, to, ddcURIs, nonce, deadline, sign)
}

// MetaSafeMint is a paid mutator transaction binding the contract method 0xa3c2bfdc.
//
// Solidity: function metaSafeMint(address to, string ddcURI_, bytes data_, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Transactor) MetaSafeMint(opts *bind.TransactOpts, to common.Address, ddcURI_ string, data_ []byte, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "metaSafeMint", to, ddcURI_, data_, nonce, deadline, sign)
}

// MetaSafeMint is a paid mutator transaction binding the contract method 0xa3c2bfdc.
//
// Solidity: function metaSafeMint(address to, string ddcURI_, bytes data_, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Session) MetaSafeMint(to common.Address, ddcURI_ string, data_ []byte, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaSafeMint(&_DDC721.TransactOpts, to, ddcURI_, data_, nonce, deadline, sign)
}

// MetaSafeMint is a paid mutator transaction binding the contract method 0xa3c2bfdc.
//
// Solidity: function metaSafeMint(address to, string ddcURI_, bytes data_, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721TransactorSession) MetaSafeMint(to common.Address, ddcURI_ string, data_ []byte, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaSafeMint(&_DDC721.TransactOpts, to, ddcURI_, data_, nonce, deadline, sign)
}

// MetaSafeMintBatch is a paid mutator transaction binding the contract method 0x705b6e37.
//
// Solidity: function metaSafeMintBatch(address to, string[] ddcURIs, bytes data, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Transactor) MetaSafeMintBatch(opts *bind.TransactOpts, to common.Address, ddcURIs []string, data []byte, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "metaSafeMintBatch", to, ddcURIs, data, nonce, deadline, sign)
}

// MetaSafeMintBatch is a paid mutator transaction binding the contract method 0x705b6e37.
//
// Solidity: function metaSafeMintBatch(address to, string[] ddcURIs, bytes data, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Session) MetaSafeMintBatch(to common.Address, ddcURIs []string, data []byte, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaSafeMintBatch(&_DDC721.TransactOpts, to, ddcURIs, data, nonce, deadline, sign)
}

// MetaSafeMintBatch is a paid mutator transaction binding the contract method 0x705b6e37.
//
// Solidity: function metaSafeMintBatch(address to, string[] ddcURIs, bytes data, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721TransactorSession) MetaSafeMintBatch(to common.Address, ddcURIs []string, data []byte, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaSafeMintBatch(&_DDC721.TransactOpts, to, ddcURIs, data, nonce, deadline, sign)
}

// MetaSafeTransferFrom is a paid mutator transaction binding the contract method 0x0439a541.
//
// Solidity: function metaSafeTransferFrom(address from, address to, uint256 ddcId, bytes data, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Transactor) MetaSafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ddcId *big.Int, data []byte, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "metaSafeTransferFrom", from, to, ddcId, data, nonce, deadline, sign)
}

// MetaSafeTransferFrom is a paid mutator transaction binding the contract method 0x0439a541.
//
// Solidity: function metaSafeTransferFrom(address from, address to, uint256 ddcId, bytes data, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Session) MetaSafeTransferFrom(from common.Address, to common.Address, ddcId *big.Int, data []byte, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaSafeTransferFrom(&_DDC721.TransactOpts, from, to, ddcId, data, nonce, deadline, sign)
}

// MetaSafeTransferFrom is a paid mutator transaction binding the contract method 0x0439a541.
//
// Solidity: function metaSafeTransferFrom(address from, address to, uint256 ddcId, bytes data, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721TransactorSession) MetaSafeTransferFrom(from common.Address, to common.Address, ddcId *big.Int, data []byte, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaSafeTransferFrom(&_DDC721.TransactOpts, from, to, ddcId, data, nonce, deadline, sign)
}

// MetaTransferFrom is a paid mutator transaction binding the contract method 0x6cce4b24.
//
// Solidity: function metaTransferFrom(address from, address to, uint256 ddcId, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Transactor) MetaTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ddcId *big.Int, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "metaTransferFrom", from, to, ddcId, nonce, deadline, sign)
}

// MetaTransferFrom is a paid mutator transaction binding the contract method 0x6cce4b24.
//
// Solidity: function metaTransferFrom(address from, address to, uint256 ddcId, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721Session) MetaTransferFrom(from common.Address, to common.Address, ddcId *big.Int, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaTransferFrom(&_DDC721.TransactOpts, from, to, ddcId, nonce, deadline, sign)
}

// MetaTransferFrom is a paid mutator transaction binding the contract method 0x6cce4b24.
//
// Solidity: function metaTransferFrom(address from, address to, uint256 ddcId, uint256 nonce, uint256 deadline, bytes sign) returns()
func (_DDC721 *DDC721TransactorSession) MetaTransferFrom(from common.Address, to common.Address, ddcId *big.Int, nonce *big.Int, deadline *big.Int, sign []byte) (*types.Transaction, error) {
	return _DDC721.Contract.MetaTransferFrom(&_DDC721.TransactOpts, from, to, ddcId, nonce, deadline, sign)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string ddcURI_) returns()
func (_DDC721 *DDC721Transactor) Mint(opts *bind.TransactOpts, to common.Address, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "mint", to, ddcURI_)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string ddcURI_) returns()
func (_DDC721 *DDC721Session) Mint(to common.Address, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721.Contract.Mint(&_DDC721.TransactOpts, to, ddcURI_)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string ddcURI_) returns()
func (_DDC721 *DDC721TransactorSession) Mint(to common.Address, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721.Contract.Mint(&_DDC721.TransactOpts, to, ddcURI_)
}

// MintBatch is a paid mutator transaction binding the contract method 0xed0e31de.
//
// Solidity: function mintBatch(address to, string[] ddcURIs) returns()
func (_DDC721 *DDC721Transactor) MintBatch(opts *bind.TransactOpts, to common.Address, ddcURIs []string) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "mintBatch", to, ddcURIs)
}

// MintBatch is a paid mutator transaction binding the contract method 0xed0e31de.
//
// Solidity: function mintBatch(address to, string[] ddcURIs) returns()
func (_DDC721 *DDC721Session) MintBatch(to common.Address, ddcURIs []string) (*types.Transaction, error) {
	return _DDC721.Contract.MintBatch(&_DDC721.TransactOpts, to, ddcURIs)
}

// MintBatch is a paid mutator transaction binding the contract method 0xed0e31de.
//
// Solidity: function mintBatch(address to, string[] ddcURIs) returns()
func (_DDC721 *DDC721TransactorSession) MintBatch(to common.Address, ddcURIs []string) (*types.Transaction, error) {
	return _DDC721.Contract.MintBatch(&_DDC721.TransactOpts, to, ddcURIs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDC721 *DDC721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDC721 *DDC721Session) RenounceOwnership() (*types.Transaction, error) {
	return _DDC721.Contract.RenounceOwnership(&_DDC721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DDC721 *DDC721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DDC721.Contract.RenounceOwnership(&_DDC721.TransactOpts)
}

// SafeMint is a paid mutator transaction binding the contract method 0xf6dda936.
//
// Solidity: function safeMint(address to, string _ddcURI, bytes _data) returns()
func (_DDC721 *DDC721Transactor) SafeMint(opts *bind.TransactOpts, to common.Address, _ddcURI string, _data []byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "safeMint", to, _ddcURI, _data)
}

// SafeMint is a paid mutator transaction binding the contract method 0xf6dda936.
//
// Solidity: function safeMint(address to, string _ddcURI, bytes _data) returns()
func (_DDC721 *DDC721Session) SafeMint(to common.Address, _ddcURI string, _data []byte) (*types.Transaction, error) {
	return _DDC721.Contract.SafeMint(&_DDC721.TransactOpts, to, _ddcURI, _data)
}

// SafeMint is a paid mutator transaction binding the contract method 0xf6dda936.
//
// Solidity: function safeMint(address to, string _ddcURI, bytes _data) returns()
func (_DDC721 *DDC721TransactorSession) SafeMint(to common.Address, _ddcURI string, _data []byte) (*types.Transaction, error) {
	return _DDC721.Contract.SafeMint(&_DDC721.TransactOpts, to, _ddcURI, _data)
}

// SafeMintBatch is a paid mutator transaction binding the contract method 0xf96ab8c4.
//
// Solidity: function safeMintBatch(address to, string[] ddcURIs, bytes data) returns()
func (_DDC721 *DDC721Transactor) SafeMintBatch(opts *bind.TransactOpts, to common.Address, ddcURIs []string, data []byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "safeMintBatch", to, ddcURIs, data)
}

// SafeMintBatch is a paid mutator transaction binding the contract method 0xf96ab8c4.
//
// Solidity: function safeMintBatch(address to, string[] ddcURIs, bytes data) returns()
func (_DDC721 *DDC721Session) SafeMintBatch(to common.Address, ddcURIs []string, data []byte) (*types.Transaction, error) {
	return _DDC721.Contract.SafeMintBatch(&_DDC721.TransactOpts, to, ddcURIs, data)
}

// SafeMintBatch is a paid mutator transaction binding the contract method 0xf96ab8c4.
//
// Solidity: function safeMintBatch(address to, string[] ddcURIs, bytes data) returns()
func (_DDC721 *DDC721TransactorSession) SafeMintBatch(to common.Address, ddcURIs []string, data []byte) (*types.Transaction, error) {
	return _DDC721.Contract.SafeMintBatch(&_DDC721.TransactOpts, to, ddcURIs, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 ddcId, bytes data) returns()
func (_DDC721 *DDC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ddcId *big.Int, data []byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "safeTransferFrom", from, to, ddcId, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 ddcId, bytes data) returns()
func (_DDC721 *DDC721Session) SafeTransferFrom(from common.Address, to common.Address, ddcId *big.Int, data []byte) (*types.Transaction, error) {
	return _DDC721.Contract.SafeTransferFrom(&_DDC721.TransactOpts, from, to, ddcId, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 ddcId, bytes data) returns()
func (_DDC721 *DDC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, ddcId *big.Int, data []byte) (*types.Transaction, error) {
	return _DDC721.Contract.SafeTransferFrom(&_DDC721.TransactOpts, from, to, ddcId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DDC721 *DDC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DDC721 *DDC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DDC721.Contract.SetApprovalForAll(&_DDC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DDC721 *DDC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DDC721.Contract.SetApprovalForAll(&_DDC721.TransactOpts, operator, approved)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_DDC721 *DDC721Transactor) SetAuthorityProxyAddress(opts *bind.TransactOpts, authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "setAuthorityProxyAddress", authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_DDC721 *DDC721Session) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721.Contract.SetAuthorityProxyAddress(&_DDC721.TransactOpts, authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_DDC721 *DDC721TransactorSession) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721.Contract.SetAuthorityProxyAddress(&_DDC721.TransactOpts, authorityProxyAddress)
}

// SetChargeProxyAddress is a paid mutator transaction binding the contract method 0x44d891df.
//
// Solidity: function setChargeProxyAddress(address chargeProxyAddress) returns()
func (_DDC721 *DDC721Transactor) SetChargeProxyAddress(opts *bind.TransactOpts, chargeProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "setChargeProxyAddress", chargeProxyAddress)
}

// SetChargeProxyAddress is a paid mutator transaction binding the contract method 0x44d891df.
//
// Solidity: function setChargeProxyAddress(address chargeProxyAddress) returns()
func (_DDC721 *DDC721Session) SetChargeProxyAddress(chargeProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721.Contract.SetChargeProxyAddress(&_DDC721.TransactOpts, chargeProxyAddress)
}

// SetChargeProxyAddress is a paid mutator transaction binding the contract method 0x44d891df.
//
// Solidity: function setChargeProxyAddress(address chargeProxyAddress) returns()
func (_DDC721 *DDC721TransactorSession) SetChargeProxyAddress(chargeProxyAddress common.Address) (*types.Transaction, error) {
	return _DDC721.Contract.SetChargeProxyAddress(&_DDC721.TransactOpts, chargeProxyAddress)
}

// SetMetaSeparatorArg is a paid mutator transaction binding the contract method 0xac10118b.
//
// Solidity: function setMetaSeparatorArg(bytes32 separator) returns()
func (_DDC721 *DDC721Transactor) SetMetaSeparatorArg(opts *bind.TransactOpts, separator [32]byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "setMetaSeparatorArg", separator)
}

// SetMetaSeparatorArg is a paid mutator transaction binding the contract method 0xac10118b.
//
// Solidity: function setMetaSeparatorArg(bytes32 separator) returns()
func (_DDC721 *DDC721Session) SetMetaSeparatorArg(separator [32]byte) (*types.Transaction, error) {
	return _DDC721.Contract.SetMetaSeparatorArg(&_DDC721.TransactOpts, separator)
}

// SetMetaSeparatorArg is a paid mutator transaction binding the contract method 0xac10118b.
//
// Solidity: function setMetaSeparatorArg(bytes32 separator) returns()
func (_DDC721 *DDC721TransactorSession) SetMetaSeparatorArg(separator [32]byte) (*types.Transaction, error) {
	return _DDC721.Contract.SetMetaSeparatorArg(&_DDC721.TransactOpts, separator)
}

// SetMetaTypeHashArgs is a paid mutator transaction binding the contract method 0x8d643054.
//
// Solidity: function setMetaTypeHashArgs(uint8 hashType, bytes32 hashValue) returns()
func (_DDC721 *DDC721Transactor) SetMetaTypeHashArgs(opts *bind.TransactOpts, hashType uint8, hashValue [32]byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "setMetaTypeHashArgs", hashType, hashValue)
}

// SetMetaTypeHashArgs is a paid mutator transaction binding the contract method 0x8d643054.
//
// Solidity: function setMetaTypeHashArgs(uint8 hashType, bytes32 hashValue) returns()
func (_DDC721 *DDC721Session) SetMetaTypeHashArgs(hashType uint8, hashValue [32]byte) (*types.Transaction, error) {
	return _DDC721.Contract.SetMetaTypeHashArgs(&_DDC721.TransactOpts, hashType, hashValue)
}

// SetMetaTypeHashArgs is a paid mutator transaction binding the contract method 0x8d643054.
//
// Solidity: function setMetaTypeHashArgs(uint8 hashType, bytes32 hashValue) returns()
func (_DDC721 *DDC721TransactorSession) SetMetaTypeHashArgs(hashType uint8, hashValue [32]byte) (*types.Transaction, error) {
	return _DDC721.Contract.SetMetaTypeHashArgs(&_DDC721.TransactOpts, hashType, hashValue)
}

// SetNameAndSymbol is a paid mutator transaction binding the contract method 0x5a446215.
//
// Solidity: function setNameAndSymbol(string name_, string symbol_) returns()
func (_DDC721 *DDC721Transactor) SetNameAndSymbol(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "setNameAndSymbol", name_, symbol_)
}

// SetNameAndSymbol is a paid mutator transaction binding the contract method 0x5a446215.
//
// Solidity: function setNameAndSymbol(string name_, string symbol_) returns()
func (_DDC721 *DDC721Session) SetNameAndSymbol(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DDC721.Contract.SetNameAndSymbol(&_DDC721.TransactOpts, name_, symbol_)
}

// SetNameAndSymbol is a paid mutator transaction binding the contract method 0x5a446215.
//
// Solidity: function setNameAndSymbol(string name_, string symbol_) returns()
func (_DDC721 *DDC721TransactorSession) SetNameAndSymbol(name_ string, symbol_ string) (*types.Transaction, error) {
	return _DDC721.Contract.SetNameAndSymbol(&_DDC721.TransactOpts, name_, symbol_)
}

// SetURI is a paid mutator transaction binding the contract method 0x862440e2.
//
// Solidity: function setURI(uint256 ddcId, string ddcURI_) returns()
func (_DDC721 *DDC721Transactor) SetURI(opts *bind.TransactOpts, ddcId *big.Int, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "setURI", ddcId, ddcURI_)
}

// SetURI is a paid mutator transaction binding the contract method 0x862440e2.
//
// Solidity: function setURI(uint256 ddcId, string ddcURI_) returns()
func (_DDC721 *DDC721Session) SetURI(ddcId *big.Int, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721.Contract.SetURI(&_DDC721.TransactOpts, ddcId, ddcURI_)
}

// SetURI is a paid mutator transaction binding the contract method 0x862440e2.
//
// Solidity: function setURI(uint256 ddcId, string ddcURI_) returns()
func (_DDC721 *DDC721TransactorSession) SetURI(ddcId *big.Int, ddcURI_ string) (*types.Transaction, error) {
	return _DDC721.Contract.SetURI(&_DDC721.TransactOpts, ddcId, ddcURI_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 ddcId) returns()
func (_DDC721 *DDC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "transferFrom", from, to, ddcId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 ddcId) returns()
func (_DDC721 *DDC721Session) TransferFrom(from common.Address, to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.TransferFrom(&_DDC721.TransactOpts, from, to, ddcId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 ddcId) returns()
func (_DDC721 *DDC721TransactorSession) TransferFrom(from common.Address, to common.Address, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.TransferFrom(&_DDC721.TransactOpts, from, to, ddcId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDC721 *DDC721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDC721 *DDC721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DDC721.Contract.TransferOwnership(&_DDC721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DDC721 *DDC721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DDC721.Contract.TransferOwnership(&_DDC721.TransactOpts, newOwner)
}

// UnFreeze is a paid mutator transaction binding the contract method 0xd302b0dc.
//
// Solidity: function unFreeze(uint256 ddcId) returns()
func (_DDC721 *DDC721Transactor) UnFreeze(opts *bind.TransactOpts, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "unFreeze", ddcId)
}

// UnFreeze is a paid mutator transaction binding the contract method 0xd302b0dc.
//
// Solidity: function unFreeze(uint256 ddcId) returns()
func (_DDC721 *DDC721Session) UnFreeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.UnFreeze(&_DDC721.TransactOpts, ddcId)
}

// UnFreeze is a paid mutator transaction binding the contract method 0xd302b0dc.
//
// Solidity: function unFreeze(uint256 ddcId) returns()
func (_DDC721 *DDC721TransactorSession) UnFreeze(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.UnFreeze(&_DDC721.TransactOpts, ddcId)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 ddcId) returns()
func (_DDC721 *DDC721Transactor) Unlock(opts *bind.TransactOpts, ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "unlock", ddcId)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 ddcId) returns()
func (_DDC721 *DDC721Session) Unlock(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.Unlock(&_DDC721.TransactOpts, ddcId)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 ddcId) returns()
func (_DDC721 *DDC721TransactorSession) Unlock(ddcId *big.Int) (*types.Transaction, error) {
	return _DDC721.Contract.Unlock(&_DDC721.TransactOpts, ddcId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DDC721 *DDC721Transactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DDC721 *DDC721Session) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _DDC721.Contract.UpgradeTo(&_DDC721.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_DDC721 *DDC721TransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _DDC721.Contract.UpgradeTo(&_DDC721.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DDC721 *DDC721Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DDC721.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DDC721 *DDC721Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DDC721.Contract.UpgradeToAndCall(&_DDC721.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_DDC721 *DDC721TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _DDC721.Contract.UpgradeToAndCall(&_DDC721.TransactOpts, newImplementation, data)
}

// DDC721AdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the DDC721 contract.
type DDC721AdminChangedIterator struct {
	Event *DDC721AdminChanged // Event containing the contract specifics and raw log

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
func (it *DDC721AdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721AdminChanged)
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
		it.Event = new(DDC721AdminChanged)
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
func (it *DDC721AdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721AdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721AdminChanged represents a AdminChanged event raised by the DDC721 contract.
type DDC721AdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_DDC721 *DDC721Filterer) FilterAdminChanged(opts *bind.FilterOpts) (*DDC721AdminChangedIterator, error) {

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &DDC721AdminChangedIterator{contract: _DDC721.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_DDC721 *DDC721Filterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *DDC721AdminChanged) (event.Subscription, error) {

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721AdminChanged)
				if err := _DDC721.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_DDC721 *DDC721Filterer) ParseAdminChanged(log types.Log) (*DDC721AdminChanged, error) {
	event := new(DDC721AdminChanged)
	if err := _DDC721.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DDC721 contract.
type DDC721ApprovalIterator struct {
	Event *DDC721Approval // Event containing the contract specifics and raw log

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
func (it *DDC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721Approval)
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
		it.Event = new(DDC721Approval)
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
func (it *DDC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721Approval represents a Approval event raised by the DDC721 contract.
type DDC721Approval struct {
	Owner    common.Address
	Approved common.Address
	DdcId    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed ddcId)
func (_DDC721 *DDC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, ddcId []*big.Int) (*DDC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, ddcIdRule)
	if err != nil {
		return nil, err
	}
	return &DDC721ApprovalIterator{contract: _DDC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed ddcId)
func (_DDC721 *DDC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DDC721Approval, owner []common.Address, approved []common.Address, ddcId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, ddcIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721Approval)
				if err := _DDC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed ddcId)
func (_DDC721 *DDC721Filterer) ParseApproval(log types.Log) (*DDC721Approval, error) {
	event := new(DDC721Approval)
	if err := _DDC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DDC721 contract.
type DDC721ApprovalForAllIterator struct {
	Event *DDC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DDC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721ApprovalForAll)
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
		it.Event = new(DDC721ApprovalForAll)
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
func (it *DDC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721ApprovalForAll represents a ApprovalForAll event raised by the DDC721 contract.
type DDC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DDC721 *DDC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DDC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DDC721ApprovalForAllIterator{contract: _DDC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DDC721 *DDC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DDC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721ApprovalForAll)
				if err := _DDC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DDC721 *DDC721Filterer) ParseApprovalForAll(log types.Log) (*DDC721ApprovalForAll, error) {
	event := new(DDC721ApprovalForAll)
	if err := _DDC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721BeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the DDC721 contract.
type DDC721BeaconUpgradedIterator struct {
	Event *DDC721BeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *DDC721BeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721BeaconUpgraded)
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
		it.Event = new(DDC721BeaconUpgraded)
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
func (it *DDC721BeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721BeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721BeaconUpgraded represents a BeaconUpgraded event raised by the DDC721 contract.
type DDC721BeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_DDC721 *DDC721Filterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*DDC721BeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &DDC721BeaconUpgradedIterator{contract: _DDC721.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_DDC721 *DDC721Filterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *DDC721BeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721BeaconUpgraded)
				if err := _DDC721.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_DDC721 *DDC721Filterer) ParseBeaconUpgraded(log types.Log) (*DDC721BeaconUpgraded, error) {
	event := new(DDC721BeaconUpgraded)
	if err := _DDC721.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721EnterBlacklistIterator is returned from FilterEnterBlacklist and is used to iterate over the raw logs and unpacked data for EnterBlacklist events raised by the DDC721 contract.
type DDC721EnterBlacklistIterator struct {
	Event *DDC721EnterBlacklist // Event containing the contract specifics and raw log

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
func (it *DDC721EnterBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721EnterBlacklist)
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
		it.Event = new(DDC721EnterBlacklist)
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
func (it *DDC721EnterBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721EnterBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721EnterBlacklist represents a EnterBlacklist event raised by the DDC721 contract.
type DDC721EnterBlacklist struct {
	Sender common.Address
	DdcId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnterBlacklist is a free log retrieval operation binding the contract event 0x027b0995c9aa454830a50ece99b9507432deb5f7ff0173efc4429282c1710a36.
//
// Solidity: event EnterBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721 *DDC721Filterer) FilterEnterBlacklist(opts *bind.FilterOpts, sender []common.Address) (*DDC721EnterBlacklistIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "EnterBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return &DDC721EnterBlacklistIterator{contract: _DDC721.contract, event: "EnterBlacklist", logs: logs, sub: sub}, nil
}

// WatchEnterBlacklist is a free log subscription operation binding the contract event 0x027b0995c9aa454830a50ece99b9507432deb5f7ff0173efc4429282c1710a36.
//
// Solidity: event EnterBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721 *DDC721Filterer) WatchEnterBlacklist(opts *bind.WatchOpts, sink chan<- *DDC721EnterBlacklist, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "EnterBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721EnterBlacklist)
				if err := _DDC721.contract.UnpackLog(event, "EnterBlacklist", log); err != nil {
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

// ParseEnterBlacklist is a log parse operation binding the contract event 0x027b0995c9aa454830a50ece99b9507432deb5f7ff0173efc4429282c1710a36.
//
// Solidity: event EnterBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721 *DDC721Filterer) ParseEnterBlacklist(log types.Log) (*DDC721EnterBlacklist, error) {
	event := new(DDC721EnterBlacklist)
	if err := _DDC721.contract.UnpackLog(event, "EnterBlacklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721ExitBlacklistIterator is returned from FilterExitBlacklist and is used to iterate over the raw logs and unpacked data for ExitBlacklist events raised by the DDC721 contract.
type DDC721ExitBlacklistIterator struct {
	Event *DDC721ExitBlacklist // Event containing the contract specifics and raw log

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
func (it *DDC721ExitBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721ExitBlacklist)
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
		it.Event = new(DDC721ExitBlacklist)
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
func (it *DDC721ExitBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721ExitBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721ExitBlacklist represents a ExitBlacklist event raised by the DDC721 contract.
type DDC721ExitBlacklist struct {
	Sender common.Address
	DdcId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExitBlacklist is a free log retrieval operation binding the contract event 0xaddb66f781fad31382e12b8ad189f90d41b9590625a6736ef67a2792f094874f.
//
// Solidity: event ExitBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721 *DDC721Filterer) FilterExitBlacklist(opts *bind.FilterOpts, sender []common.Address) (*DDC721ExitBlacklistIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "ExitBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return &DDC721ExitBlacklistIterator{contract: _DDC721.contract, event: "ExitBlacklist", logs: logs, sub: sub}, nil
}

// WatchExitBlacklist is a free log subscription operation binding the contract event 0xaddb66f781fad31382e12b8ad189f90d41b9590625a6736ef67a2792f094874f.
//
// Solidity: event ExitBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721 *DDC721Filterer) WatchExitBlacklist(opts *bind.WatchOpts, sink chan<- *DDC721ExitBlacklist, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "ExitBlacklist", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721ExitBlacklist)
				if err := _DDC721.contract.UnpackLog(event, "ExitBlacklist", log); err != nil {
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

// ParseExitBlacklist is a log parse operation binding the contract event 0xaddb66f781fad31382e12b8ad189f90d41b9590625a6736ef67a2792f094874f.
//
// Solidity: event ExitBlacklist(address indexed sender, uint256 ddcId)
func (_DDC721 *DDC721Filterer) ParseExitBlacklist(log types.Log) (*DDC721ExitBlacklist, error) {
	event := new(DDC721ExitBlacklist)
	if err := _DDC721.contract.UnpackLog(event, "ExitBlacklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721LocklistIterator is returned from FilterLocklist and is used to iterate over the raw logs and unpacked data for Locklist events raised by the DDC721 contract.
type DDC721LocklistIterator struct {
	Event *DDC721Locklist // Event containing the contract specifics and raw log

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
func (it *DDC721LocklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721Locklist)
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
		it.Event = new(DDC721Locklist)
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
func (it *DDC721LocklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721LocklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721Locklist represents a Locklist event raised by the DDC721 contract.
type DDC721Locklist struct {
	Operator common.Address
	DdcId    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLocklist is a free log retrieval operation binding the contract event 0xabcf25e46f2d1f3313c34b3fe668d0d864272dd3fb5cda4624f4b4016dba4d5b.
//
// Solidity: event Locklist(address indexed operator, uint256 ddcId)
func (_DDC721 *DDC721Filterer) FilterLocklist(opts *bind.FilterOpts, operator []common.Address) (*DDC721LocklistIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "Locklist", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DDC721LocklistIterator{contract: _DDC721.contract, event: "Locklist", logs: logs, sub: sub}, nil
}

// WatchLocklist is a free log subscription operation binding the contract event 0xabcf25e46f2d1f3313c34b3fe668d0d864272dd3fb5cda4624f4b4016dba4d5b.
//
// Solidity: event Locklist(address indexed operator, uint256 ddcId)
func (_DDC721 *DDC721Filterer) WatchLocklist(opts *bind.WatchOpts, sink chan<- *DDC721Locklist, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "Locklist", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721Locklist)
				if err := _DDC721.contract.UnpackLog(event, "Locklist", log); err != nil {
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

// ParseLocklist is a log parse operation binding the contract event 0xabcf25e46f2d1f3313c34b3fe668d0d864272dd3fb5cda4624f4b4016dba4d5b.
//
// Solidity: event Locklist(address indexed operator, uint256 ddcId)
func (_DDC721 *DDC721Filterer) ParseLocklist(log types.Log) (*DDC721Locklist, error) {
	event := new(DDC721Locklist)
	if err := _DDC721.contract.UnpackLog(event, "Locklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721MetaTransferIterator is returned from FilterMetaTransfer and is used to iterate over the raw logs and unpacked data for MetaTransfer events raised by the DDC721 contract.
type DDC721MetaTransferIterator struct {
	Event *DDC721MetaTransfer // Event containing the contract specifics and raw log

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
func (it *DDC721MetaTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721MetaTransfer)
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
		it.Event = new(DDC721MetaTransfer)
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
func (it *DDC721MetaTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721MetaTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721MetaTransfer represents a MetaTransfer event raised by the DDC721 contract.
type DDC721MetaTransfer struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	DdcId    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMetaTransfer is a free log retrieval operation binding the contract event 0x60f14005dedd16ad7d205fd9ab957c302584267894f4d6e9e99683670e78fc0a.
//
// Solidity: event MetaTransfer(address indexed operator, address indexed from, address indexed to, uint256 ddcId)
func (_DDC721 *DDC721Filterer) FilterMetaTransfer(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*DDC721MetaTransferIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "MetaTransfer", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DDC721MetaTransferIterator{contract: _DDC721.contract, event: "MetaTransfer", logs: logs, sub: sub}, nil
}

// WatchMetaTransfer is a free log subscription operation binding the contract event 0x60f14005dedd16ad7d205fd9ab957c302584267894f4d6e9e99683670e78fc0a.
//
// Solidity: event MetaTransfer(address indexed operator, address indexed from, address indexed to, uint256 ddcId)
func (_DDC721 *DDC721Filterer) WatchMetaTransfer(opts *bind.WatchOpts, sink chan<- *DDC721MetaTransfer, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "MetaTransfer", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721MetaTransfer)
				if err := _DDC721.contract.UnpackLog(event, "MetaTransfer", log); err != nil {
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

// ParseMetaTransfer is a log parse operation binding the contract event 0x60f14005dedd16ad7d205fd9ab957c302584267894f4d6e9e99683670e78fc0a.
//
// Solidity: event MetaTransfer(address indexed operator, address indexed from, address indexed to, uint256 ddcId)
func (_DDC721 *DDC721Filterer) ParseMetaTransfer(log types.Log) (*DDC721MetaTransfer, error) {
	event := new(DDC721MetaTransfer)
	if err := _DDC721.contract.UnpackLog(event, "MetaTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721MetaTransferBatchIterator is returned from FilterMetaTransferBatch and is used to iterate over the raw logs and unpacked data for MetaTransferBatch events raised by the DDC721 contract.
type DDC721MetaTransferBatchIterator struct {
	Event *DDC721MetaTransferBatch // Event containing the contract specifics and raw log

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
func (it *DDC721MetaTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721MetaTransferBatch)
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
		it.Event = new(DDC721MetaTransferBatch)
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
func (it *DDC721MetaTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721MetaTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721MetaTransferBatch represents a MetaTransferBatch event raised by the DDC721 contract.
type DDC721MetaTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	DdcIds   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMetaTransferBatch is a free log retrieval operation binding the contract event 0x3d13f72b8478fbc21aa0a43f1607fbdec5b86a4b2602087e8b0d292e84132052.
//
// Solidity: event MetaTransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ddcIds)
func (_DDC721 *DDC721Filterer) FilterMetaTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*DDC721MetaTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "MetaTransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DDC721MetaTransferBatchIterator{contract: _DDC721.contract, event: "MetaTransferBatch", logs: logs, sub: sub}, nil
}

// WatchMetaTransferBatch is a free log subscription operation binding the contract event 0x3d13f72b8478fbc21aa0a43f1607fbdec5b86a4b2602087e8b0d292e84132052.
//
// Solidity: event MetaTransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ddcIds)
func (_DDC721 *DDC721Filterer) WatchMetaTransferBatch(opts *bind.WatchOpts, sink chan<- *DDC721MetaTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "MetaTransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721MetaTransferBatch)
				if err := _DDC721.contract.UnpackLog(event, "MetaTransferBatch", log); err != nil {
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

// ParseMetaTransferBatch is a log parse operation binding the contract event 0x3d13f72b8478fbc21aa0a43f1607fbdec5b86a4b2602087e8b0d292e84132052.
//
// Solidity: event MetaTransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ddcIds)
func (_DDC721 *DDC721Filterer) ParseMetaTransferBatch(log types.Log) (*DDC721MetaTransferBatch, error) {
	event := new(DDC721MetaTransferBatch)
	if err := _DDC721.contract.UnpackLog(event, "MetaTransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DDC721 contract.
type DDC721OwnershipTransferredIterator struct {
	Event *DDC721OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DDC721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721OwnershipTransferred)
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
		it.Event = new(DDC721OwnershipTransferred)
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
func (it *DDC721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721OwnershipTransferred represents a OwnershipTransferred event raised by the DDC721 contract.
type DDC721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DDC721 *DDC721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DDC721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DDC721OwnershipTransferredIterator{contract: _DDC721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DDC721 *DDC721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DDC721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721OwnershipTransferred)
				if err := _DDC721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DDC721 *DDC721Filterer) ParseOwnershipTransferred(log types.Log) (*DDC721OwnershipTransferred, error) {
	event := new(DDC721OwnershipTransferred)
	if err := _DDC721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721SetNameAndSymbolIterator is returned from FilterSetNameAndSymbol and is used to iterate over the raw logs and unpacked data for SetNameAndSymbol events raised by the DDC721 contract.
type DDC721SetNameAndSymbolIterator struct {
	Event *DDC721SetNameAndSymbol // Event containing the contract specifics and raw log

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
func (it *DDC721SetNameAndSymbolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721SetNameAndSymbol)
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
		it.Event = new(DDC721SetNameAndSymbol)
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
func (it *DDC721SetNameAndSymbolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721SetNameAndSymbolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721SetNameAndSymbol represents a SetNameAndSymbol event raised by the DDC721 contract.
type DDC721SetNameAndSymbol struct {
	Name   string
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetNameAndSymbol is a free log retrieval operation binding the contract event 0xd150542982eaa18f20ceb10f524c418899e6a331a81ee4a70daa921f976fab3b.
//
// Solidity: event SetNameAndSymbol(string name, string symbol)
func (_DDC721 *DDC721Filterer) FilterSetNameAndSymbol(opts *bind.FilterOpts) (*DDC721SetNameAndSymbolIterator, error) {

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "SetNameAndSymbol")
	if err != nil {
		return nil, err
	}
	return &DDC721SetNameAndSymbolIterator{contract: _DDC721.contract, event: "SetNameAndSymbol", logs: logs, sub: sub}, nil
}

// WatchSetNameAndSymbol is a free log subscription operation binding the contract event 0xd150542982eaa18f20ceb10f524c418899e6a331a81ee4a70daa921f976fab3b.
//
// Solidity: event SetNameAndSymbol(string name, string symbol)
func (_DDC721 *DDC721Filterer) WatchSetNameAndSymbol(opts *bind.WatchOpts, sink chan<- *DDC721SetNameAndSymbol) (event.Subscription, error) {

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "SetNameAndSymbol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721SetNameAndSymbol)
				if err := _DDC721.contract.UnpackLog(event, "SetNameAndSymbol", log); err != nil {
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

// ParseSetNameAndSymbol is a log parse operation binding the contract event 0xd150542982eaa18f20ceb10f524c418899e6a331a81ee4a70daa921f976fab3b.
//
// Solidity: event SetNameAndSymbol(string name, string symbol)
func (_DDC721 *DDC721Filterer) ParseSetNameAndSymbol(log types.Log) (*DDC721SetNameAndSymbol, error) {
	event := new(DDC721SetNameAndSymbol)
	if err := _DDC721.contract.UnpackLog(event, "SetNameAndSymbol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721SetURIIterator is returned from FilterSetURI and is used to iterate over the raw logs and unpacked data for SetURI events raised by the DDC721 contract.
type DDC721SetURIIterator struct {
	Event *DDC721SetURI // Event containing the contract specifics and raw log

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
func (it *DDC721SetURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721SetURI)
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
		it.Event = new(DDC721SetURI)
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
func (it *DDC721SetURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721SetURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721SetURI represents a SetURI event raised by the DDC721 contract.
type DDC721SetURI struct {
	DdcId  *big.Int
	DdcURI string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetURI is a free log retrieval operation binding the contract event 0xee1bb82f380189104b74a7647d26f2f35679780e816626ffcaec7cafb7288e46.
//
// Solidity: event SetURI(uint256 indexed ddcId, string ddcURI)
func (_DDC721 *DDC721Filterer) FilterSetURI(opts *bind.FilterOpts, ddcId []*big.Int) (*DDC721SetURIIterator, error) {

	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "SetURI", ddcIdRule)
	if err != nil {
		return nil, err
	}
	return &DDC721SetURIIterator{contract: _DDC721.contract, event: "SetURI", logs: logs, sub: sub}, nil
}

// WatchSetURI is a free log subscription operation binding the contract event 0xee1bb82f380189104b74a7647d26f2f35679780e816626ffcaec7cafb7288e46.
//
// Solidity: event SetURI(uint256 indexed ddcId, string ddcURI)
func (_DDC721 *DDC721Filterer) WatchSetURI(opts *bind.WatchOpts, sink chan<- *DDC721SetURI, ddcId []*big.Int) (event.Subscription, error) {

	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "SetURI", ddcIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721SetURI)
				if err := _DDC721.contract.UnpackLog(event, "SetURI", log); err != nil {
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

// ParseSetURI is a log parse operation binding the contract event 0xee1bb82f380189104b74a7647d26f2f35679780e816626ffcaec7cafb7288e46.
//
// Solidity: event SetURI(uint256 indexed ddcId, string ddcURI)
func (_DDC721 *DDC721Filterer) ParseSetURI(log types.Log) (*DDC721SetURI, error) {
	event := new(DDC721SetURI)
	if err := _DDC721.contract.UnpackLog(event, "SetURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DDC721 contract.
type DDC721TransferIterator struct {
	Event *DDC721Transfer // Event containing the contract specifics and raw log

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
func (it *DDC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721Transfer)
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
		it.Event = new(DDC721Transfer)
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
func (it *DDC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721Transfer represents a Transfer event raised by the DDC721 contract.
type DDC721Transfer struct {
	From  common.Address
	To    common.Address
	DdcId *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed ddcId)
func (_DDC721 *DDC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, ddcId []*big.Int) (*DDC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, ddcIdRule)
	if err != nil {
		return nil, err
	}
	return &DDC721TransferIterator{contract: _DDC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed ddcId)
func (_DDC721 *DDC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DDC721Transfer, from []common.Address, to []common.Address, ddcId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var ddcIdRule []interface{}
	for _, ddcIdItem := range ddcId {
		ddcIdRule = append(ddcIdRule, ddcIdItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, ddcIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721Transfer)
				if err := _DDC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed ddcId)
func (_DDC721 *DDC721Filterer) ParseTransfer(log types.Log) (*DDC721Transfer, error) {
	event := new(DDC721Transfer)
	if err := _DDC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the DDC721 contract.
type DDC721TransferBatchIterator struct {
	Event *DDC721TransferBatch // Event containing the contract specifics and raw log

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
func (it *DDC721TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721TransferBatch)
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
		it.Event = new(DDC721TransferBatch)
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
func (it *DDC721TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721TransferBatch represents a TransferBatch event raised by the DDC721 contract.
type DDC721TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	DdcIds   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x2e75a6cf483a33fd7e40b01fc5b561361f6e9b2d5a492f866bd66ca430a8c557.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ddcIds)
func (_DDC721 *DDC721Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*DDC721TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DDC721TransferBatchIterator{contract: _DDC721.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x2e75a6cf483a33fd7e40b01fc5b561361f6e9b2d5a492f866bd66ca430a8c557.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ddcIds)
func (_DDC721 *DDC721Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *DDC721TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721TransferBatch)
				if err := _DDC721.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x2e75a6cf483a33fd7e40b01fc5b561361f6e9b2d5a492f866bd66ca430a8c557.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ddcIds)
func (_DDC721 *DDC721Filterer) ParseTransferBatch(log types.Log) (*DDC721TransferBatch, error) {
	event := new(DDC721TransferBatch)
	if err := _DDC721.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721UnLocklistIterator is returned from FilterUnLocklist and is used to iterate over the raw logs and unpacked data for UnLocklist events raised by the DDC721 contract.
type DDC721UnLocklistIterator struct {
	Event *DDC721UnLocklist // Event containing the contract specifics and raw log

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
func (it *DDC721UnLocklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721UnLocklist)
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
		it.Event = new(DDC721UnLocklist)
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
func (it *DDC721UnLocklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721UnLocklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721UnLocklist represents a UnLocklist event raised by the DDC721 contract.
type DDC721UnLocklist struct {
	Operator common.Address
	DdcId    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnLocklist is a free log retrieval operation binding the contract event 0x8af3d40d6d540efc57e85305f7646c5e191c8c85b6759f43ddb0c8faf146eb45.
//
// Solidity: event UnLocklist(address indexed operator, uint256 ddcId)
func (_DDC721 *DDC721Filterer) FilterUnLocklist(opts *bind.FilterOpts, operator []common.Address) (*DDC721UnLocklistIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "UnLocklist", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DDC721UnLocklistIterator{contract: _DDC721.contract, event: "UnLocklist", logs: logs, sub: sub}, nil
}

// WatchUnLocklist is a free log subscription operation binding the contract event 0x8af3d40d6d540efc57e85305f7646c5e191c8c85b6759f43ddb0c8faf146eb45.
//
// Solidity: event UnLocklist(address indexed operator, uint256 ddcId)
func (_DDC721 *DDC721Filterer) WatchUnLocklist(opts *bind.WatchOpts, sink chan<- *DDC721UnLocklist, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "UnLocklist", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721UnLocklist)
				if err := _DDC721.contract.UnpackLog(event, "UnLocklist", log); err != nil {
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

// ParseUnLocklist is a log parse operation binding the contract event 0x8af3d40d6d540efc57e85305f7646c5e191c8c85b6759f43ddb0c8faf146eb45.
//
// Solidity: event UnLocklist(address indexed operator, uint256 ddcId)
func (_DDC721 *DDC721Filterer) ParseUnLocklist(log types.Log) (*DDC721UnLocklist, error) {
	event := new(DDC721UnLocklist)
	if err := _DDC721.contract.UnpackLog(event, "UnLocklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DDC721UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the DDC721 contract.
type DDC721UpgradedIterator struct {
	Event *DDC721Upgraded // Event containing the contract specifics and raw log

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
func (it *DDC721UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DDC721Upgraded)
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
		it.Event = new(DDC721Upgraded)
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
func (it *DDC721UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DDC721UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DDC721Upgraded represents a Upgraded event raised by the DDC721 contract.
type DDC721Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DDC721 *DDC721Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*DDC721UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DDC721.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DDC721UpgradedIterator{contract: _DDC721.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DDC721 *DDC721Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *DDC721Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _DDC721.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DDC721Upgraded)
				if err := _DDC721.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_DDC721 *DDC721Filterer) ParseUpgraded(log types.Log) (*DDC721Upgraded, error) {
	event := new(DDC721Upgraded)
	if err := _DDC721.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
