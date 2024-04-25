// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xentorrent

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

// XentorrentMetaData contains all meta data concerning the Xentorrent contract.
var XentorrentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"xenCrypto_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"burnRates_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenLimits_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlockNumber_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"forwarder_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"royaltyReceiver_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"EndTorrent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"xenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"xenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"StartTorrent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTHORS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLACKOUT_TERM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMON_CATEGORY_COUNTER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIMITED_CATEGORY_TIME_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_FILTER_REGISTRY\",\"outputs\":[{\"internalType\":\"contractIOperatorFilterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POWER_GROUP_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROYALTY_BP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SPECIAL_CATEGORIES_VMU_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mintInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"specialClassesBurnRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"specialClassesCounters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"specialClassesTokenLimits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vmuCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"xenBurned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xenCrypto\",\"outputs\":[{\"internalType\":\"contractXENCrypto\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"callClaimRank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"callClaimMintReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"powerDown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"burned\",\"type\":\"uint256\"}],\"name\":\"onTokenBurned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedForwarder\",\"type\":\"address\"}],\"name\":\"addForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownedTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isApex\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"apex\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"bulkClaimRank\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"burning\",\"type\":\"uint256\"}],\"name\":\"bulkClaimRankLimited\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"bulkClaimMintReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// XentorrentABI is the input ABI used to generate the binding from.
// Deprecated: Use XentorrentMetaData.ABI instead.
var XentorrentABI = XentorrentMetaData.ABI

// Xentorrent is an auto generated Go binding around an Ethereum contract.
type Xentorrent struct {
	XentorrentCaller     // Read-only binding to the contract
	XentorrentTransactor // Write-only binding to the contract
	XentorrentFilterer   // Log filterer for contract events
}

// XentorrentCaller is an auto generated read-only Go binding around an Ethereum contract.
type XentorrentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XentorrentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XentorrentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XentorrentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type XentorrentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XentorrentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XentorrentSession struct {
	Contract     *Xentorrent       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XentorrentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XentorrentCallerSession struct {
	Contract *XentorrentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// XentorrentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XentorrentTransactorSession struct {
	Contract     *XentorrentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// XentorrentRaw is an auto generated low-level Go binding around an Ethereum contract.
type XentorrentRaw struct {
	Contract *Xentorrent // Generic contract binding to access the raw methods on
}

// XentorrentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XentorrentCallerRaw struct {
	Contract *XentorrentCaller // Generic read-only contract binding to access the raw methods on
}

// XentorrentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XentorrentTransactorRaw struct {
	Contract *XentorrentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXentorrent creates a new instance of Xentorrent, bound to a specific deployed contract.
func NewXentorrent(address common.Address, backend bind.ContractBackend) (*Xentorrent, error) {
	contract, err := bindXentorrent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Xentorrent{XentorrentCaller: XentorrentCaller{contract: contract}, XentorrentTransactor: XentorrentTransactor{contract: contract}, XentorrentFilterer: XentorrentFilterer{contract: contract}}, nil
}

// NewXentorrentCaller creates a new read-only instance of Xentorrent, bound to a specific deployed contract.
func NewXentorrentCaller(address common.Address, caller bind.ContractCaller) (*XentorrentCaller, error) {
	contract, err := bindXentorrent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &XentorrentCaller{contract: contract}, nil
}

// NewXentorrentTransactor creates a new write-only instance of Xentorrent, bound to a specific deployed contract.
func NewXentorrentTransactor(address common.Address, transactor bind.ContractTransactor) (*XentorrentTransactor, error) {
	contract, err := bindXentorrent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &XentorrentTransactor{contract: contract}, nil
}

// NewXentorrentFilterer creates a new log filterer instance of Xentorrent, bound to a specific deployed contract.
func NewXentorrentFilterer(address common.Address, filterer bind.ContractFilterer) (*XentorrentFilterer, error) {
	contract, err := bindXentorrent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &XentorrentFilterer{contract: contract}, nil
}

// bindXentorrent binds a generic wrapper to an already deployed contract.
func bindXentorrent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := XentorrentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Xentorrent *XentorrentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Xentorrent.Contract.XentorrentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Xentorrent *XentorrentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xentorrent.Contract.XentorrentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Xentorrent *XentorrentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Xentorrent.Contract.XentorrentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Xentorrent *XentorrentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Xentorrent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Xentorrent *XentorrentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xentorrent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Xentorrent *XentorrentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Xentorrent.Contract.contract.Transact(opts, method, params...)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Xentorrent *XentorrentCaller) AUTHORS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "AUTHORS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Xentorrent *XentorrentSession) AUTHORS() (string, error) {
	return _Xentorrent.Contract.AUTHORS(&_Xentorrent.CallOpts)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Xentorrent *XentorrentCallerSession) AUTHORS() (string, error) {
	return _Xentorrent.Contract.AUTHORS(&_Xentorrent.CallOpts)
}

// BLACKOUTTERM is a free data retrieval call binding the contract method 0xa126ad1e.
//
// Solidity: function BLACKOUT_TERM() view returns(uint256)
func (_Xentorrent *XentorrentCaller) BLACKOUTTERM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "BLACKOUT_TERM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLACKOUTTERM is a free data retrieval call binding the contract method 0xa126ad1e.
//
// Solidity: function BLACKOUT_TERM() view returns(uint256)
func (_Xentorrent *XentorrentSession) BLACKOUTTERM() (*big.Int, error) {
	return _Xentorrent.Contract.BLACKOUTTERM(&_Xentorrent.CallOpts)
}

// BLACKOUTTERM is a free data retrieval call binding the contract method 0xa126ad1e.
//
// Solidity: function BLACKOUT_TERM() view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) BLACKOUTTERM() (*big.Int, error) {
	return _Xentorrent.Contract.BLACKOUTTERM(&_Xentorrent.CallOpts)
}

// COMMONCATEGORYCOUNTER is a free data retrieval call binding the contract method 0x4d4b2be4.
//
// Solidity: function COMMON_CATEGORY_COUNTER() view returns(uint256)
func (_Xentorrent *XentorrentCaller) COMMONCATEGORYCOUNTER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "COMMON_CATEGORY_COUNTER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMONCATEGORYCOUNTER is a free data retrieval call binding the contract method 0x4d4b2be4.
//
// Solidity: function COMMON_CATEGORY_COUNTER() view returns(uint256)
func (_Xentorrent *XentorrentSession) COMMONCATEGORYCOUNTER() (*big.Int, error) {
	return _Xentorrent.Contract.COMMONCATEGORYCOUNTER(&_Xentorrent.CallOpts)
}

// COMMONCATEGORYCOUNTER is a free data retrieval call binding the contract method 0x4d4b2be4.
//
// Solidity: function COMMON_CATEGORY_COUNTER() view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) COMMONCATEGORYCOUNTER() (*big.Int, error) {
	return _Xentorrent.Contract.COMMONCATEGORYCOUNTER(&_Xentorrent.CallOpts)
}

// LIMITEDCATEGORYTIMETHRESHOLD is a free data retrieval call binding the contract method 0x700107af.
//
// Solidity: function LIMITED_CATEGORY_TIME_THRESHOLD() view returns(uint256)
func (_Xentorrent *XentorrentCaller) LIMITEDCATEGORYTIMETHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "LIMITED_CATEGORY_TIME_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LIMITEDCATEGORYTIMETHRESHOLD is a free data retrieval call binding the contract method 0x700107af.
//
// Solidity: function LIMITED_CATEGORY_TIME_THRESHOLD() view returns(uint256)
func (_Xentorrent *XentorrentSession) LIMITEDCATEGORYTIMETHRESHOLD() (*big.Int, error) {
	return _Xentorrent.Contract.LIMITEDCATEGORYTIMETHRESHOLD(&_Xentorrent.CallOpts)
}

// LIMITEDCATEGORYTIMETHRESHOLD is a free data retrieval call binding the contract method 0x700107af.
//
// Solidity: function LIMITED_CATEGORY_TIME_THRESHOLD() view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) LIMITEDCATEGORYTIMETHRESHOLD() (*big.Int, error) {
	return _Xentorrent.Contract.LIMITEDCATEGORYTIMETHRESHOLD(&_Xentorrent.CallOpts)
}

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Xentorrent *XentorrentCaller) OPERATORFILTERREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "OPERATOR_FILTER_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Xentorrent *XentorrentSession) OPERATORFILTERREGISTRY() (common.Address, error) {
	return _Xentorrent.Contract.OPERATORFILTERREGISTRY(&_Xentorrent.CallOpts)
}

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Xentorrent *XentorrentCallerSession) OPERATORFILTERREGISTRY() (common.Address, error) {
	return _Xentorrent.Contract.OPERATORFILTERREGISTRY(&_Xentorrent.CallOpts)
}

// POWERGROUPSIZE is a free data retrieval call binding the contract method 0x41b169f3.
//
// Solidity: function POWER_GROUP_SIZE() view returns(uint256)
func (_Xentorrent *XentorrentCaller) POWERGROUPSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "POWER_GROUP_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// POWERGROUPSIZE is a free data retrieval call binding the contract method 0x41b169f3.
//
// Solidity: function POWER_GROUP_SIZE() view returns(uint256)
func (_Xentorrent *XentorrentSession) POWERGROUPSIZE() (*big.Int, error) {
	return _Xentorrent.Contract.POWERGROUPSIZE(&_Xentorrent.CallOpts)
}

// POWERGROUPSIZE is a free data retrieval call binding the contract method 0x41b169f3.
//
// Solidity: function POWER_GROUP_SIZE() view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) POWERGROUPSIZE() (*big.Int, error) {
	return _Xentorrent.Contract.POWERGROUPSIZE(&_Xentorrent.CallOpts)
}

// ROYALTYBP is a free data retrieval call binding the contract method 0x044db8ba.
//
// Solidity: function ROYALTY_BP() view returns(uint256)
func (_Xentorrent *XentorrentCaller) ROYALTYBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "ROYALTY_BP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROYALTYBP is a free data retrieval call binding the contract method 0x044db8ba.
//
// Solidity: function ROYALTY_BP() view returns(uint256)
func (_Xentorrent *XentorrentSession) ROYALTYBP() (*big.Int, error) {
	return _Xentorrent.Contract.ROYALTYBP(&_Xentorrent.CallOpts)
}

// ROYALTYBP is a free data retrieval call binding the contract method 0x044db8ba.
//
// Solidity: function ROYALTY_BP() view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) ROYALTYBP() (*big.Int, error) {
	return _Xentorrent.Contract.ROYALTYBP(&_Xentorrent.CallOpts)
}

// SPECIALCATEGORIESVMUTHRESHOLD is a free data retrieval call binding the contract method 0x55ee08ba.
//
// Solidity: function SPECIAL_CATEGORIES_VMU_THRESHOLD() view returns(uint256)
func (_Xentorrent *XentorrentCaller) SPECIALCATEGORIESVMUTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "SPECIAL_CATEGORIES_VMU_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SPECIALCATEGORIESVMUTHRESHOLD is a free data retrieval call binding the contract method 0x55ee08ba.
//
// Solidity: function SPECIAL_CATEGORIES_VMU_THRESHOLD() view returns(uint256)
func (_Xentorrent *XentorrentSession) SPECIALCATEGORIESVMUTHRESHOLD() (*big.Int, error) {
	return _Xentorrent.Contract.SPECIALCATEGORIESVMUTHRESHOLD(&_Xentorrent.CallOpts)
}

// SPECIALCATEGORIESVMUTHRESHOLD is a free data retrieval call binding the contract method 0x55ee08ba.
//
// Solidity: function SPECIAL_CATEGORIES_VMU_THRESHOLD() view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) SPECIALCATEGORIESVMUTHRESHOLD() (*big.Int, error) {
	return _Xentorrent.Contract.SPECIALCATEGORIESVMUTHRESHOLD(&_Xentorrent.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Xentorrent *XentorrentCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Xentorrent *XentorrentSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Xentorrent.Contract.BalanceOf(&_Xentorrent.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Xentorrent.Contract.BalanceOf(&_Xentorrent.CallOpts, owner)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_Xentorrent *XentorrentCaller) GenesisTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "genesisTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_Xentorrent *XentorrentSession) GenesisTs() (*big.Int, error) {
	return _Xentorrent.Contract.GenesisTs(&_Xentorrent.CallOpts)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) GenesisTs() (*big.Int, error) {
	return _Xentorrent.Contract.GenesisTs(&_Xentorrent.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Xentorrent *XentorrentCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Xentorrent *XentorrentSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Xentorrent.Contract.GetApproved(&_Xentorrent.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Xentorrent *XentorrentCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Xentorrent.Contract.GetApproved(&_Xentorrent.CallOpts, tokenId)
}

// IsApex is a free data retrieval call binding the contract method 0xee8743d7.
//
// Solidity: function isApex(uint256 tokenId) pure returns(bool apex)
func (_Xentorrent *XentorrentCaller) IsApex(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "isApex", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApex is a free data retrieval call binding the contract method 0xee8743d7.
//
// Solidity: function isApex(uint256 tokenId) pure returns(bool apex)
func (_Xentorrent *XentorrentSession) IsApex(tokenId *big.Int) (bool, error) {
	return _Xentorrent.Contract.IsApex(&_Xentorrent.CallOpts, tokenId)
}

// IsApex is a free data retrieval call binding the contract method 0xee8743d7.
//
// Solidity: function isApex(uint256 tokenId) pure returns(bool apex)
func (_Xentorrent *XentorrentCallerSession) IsApex(tokenId *big.Int) (bool, error) {
	return _Xentorrent.Contract.IsApex(&_Xentorrent.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Xentorrent *XentorrentCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Xentorrent *XentorrentSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Xentorrent.Contract.IsApprovedForAll(&_Xentorrent.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Xentorrent *XentorrentCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Xentorrent.Contract.IsApprovedForAll(&_Xentorrent.CallOpts, owner, operator)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Xentorrent *XentorrentCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Xentorrent *XentorrentSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Xentorrent.Contract.IsTrustedForwarder(&_Xentorrent.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Xentorrent *XentorrentCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Xentorrent.Contract.IsTrustedForwarder(&_Xentorrent.CallOpts, forwarder)
}

// MintInfo is a free data retrieval call binding the contract method 0x443aa533.
//
// Solidity: function mintInfo(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCaller) MintInfo(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "mintInfo", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintInfo is a free data retrieval call binding the contract method 0x443aa533.
//
// Solidity: function mintInfo(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentSession) MintInfo(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.MintInfo(&_Xentorrent.CallOpts, arg0)
}

// MintInfo is a free data retrieval call binding the contract method 0x443aa533.
//
// Solidity: function mintInfo(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) MintInfo(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.MintInfo(&_Xentorrent.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Xentorrent *XentorrentCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Xentorrent *XentorrentSession) Name() (string, error) {
	return _Xentorrent.Contract.Name(&_Xentorrent.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Xentorrent *XentorrentCallerSession) Name() (string, error) {
	return _Xentorrent.Contract.Name(&_Xentorrent.CallOpts)
}

// OwnedTokens is a free data retrieval call binding the contract method 0x19cba6b4.
//
// Solidity: function ownedTokens() view returns(uint256[])
func (_Xentorrent *XentorrentCaller) OwnedTokens(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "ownedTokens")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// OwnedTokens is a free data retrieval call binding the contract method 0x19cba6b4.
//
// Solidity: function ownedTokens() view returns(uint256[])
func (_Xentorrent *XentorrentSession) OwnedTokens() ([]*big.Int, error) {
	return _Xentorrent.Contract.OwnedTokens(&_Xentorrent.CallOpts)
}

// OwnedTokens is a free data retrieval call binding the contract method 0x19cba6b4.
//
// Solidity: function ownedTokens() view returns(uint256[])
func (_Xentorrent *XentorrentCallerSession) OwnedTokens() ([]*big.Int, error) {
	return _Xentorrent.Contract.OwnedTokens(&_Xentorrent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Xentorrent *XentorrentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Xentorrent *XentorrentSession) Owner() (common.Address, error) {
	return _Xentorrent.Contract.Owner(&_Xentorrent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Xentorrent *XentorrentCallerSession) Owner() (common.Address, error) {
	return _Xentorrent.Contract.Owner(&_Xentorrent.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Xentorrent *XentorrentCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Xentorrent *XentorrentSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Xentorrent.Contract.OwnerOf(&_Xentorrent.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Xentorrent *XentorrentCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Xentorrent.Contract.OwnerOf(&_Xentorrent.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Xentorrent *XentorrentCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "royaltyInfo", arg0, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Xentorrent *XentorrentSession) RoyaltyInfo(arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Xentorrent.Contract.RoyaltyInfo(&_Xentorrent.CallOpts, arg0, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Xentorrent *XentorrentCallerSession) RoyaltyInfo(arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Xentorrent.Contract.RoyaltyInfo(&_Xentorrent.CallOpts, arg0, salePrice)
}

// SpecialClassesBurnRates is a free data retrieval call binding the contract method 0xd0d5f5b4.
//
// Solidity: function specialClassesBurnRates(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCaller) SpecialClassesBurnRates(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "specialClassesBurnRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpecialClassesBurnRates is a free data retrieval call binding the contract method 0xd0d5f5b4.
//
// Solidity: function specialClassesBurnRates(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentSession) SpecialClassesBurnRates(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.SpecialClassesBurnRates(&_Xentorrent.CallOpts, arg0)
}

// SpecialClassesBurnRates is a free data retrieval call binding the contract method 0xd0d5f5b4.
//
// Solidity: function specialClassesBurnRates(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) SpecialClassesBurnRates(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.SpecialClassesBurnRates(&_Xentorrent.CallOpts, arg0)
}

// SpecialClassesCounters is a free data retrieval call binding the contract method 0x89776eb0.
//
// Solidity: function specialClassesCounters(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCaller) SpecialClassesCounters(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "specialClassesCounters", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpecialClassesCounters is a free data retrieval call binding the contract method 0x89776eb0.
//
// Solidity: function specialClassesCounters(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentSession) SpecialClassesCounters(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.SpecialClassesCounters(&_Xentorrent.CallOpts, arg0)
}

// SpecialClassesCounters is a free data retrieval call binding the contract method 0x89776eb0.
//
// Solidity: function specialClassesCounters(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) SpecialClassesCounters(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.SpecialClassesCounters(&_Xentorrent.CallOpts, arg0)
}

// SpecialClassesTokenLimits is a free data retrieval call binding the contract method 0x74a1dff2.
//
// Solidity: function specialClassesTokenLimits(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCaller) SpecialClassesTokenLimits(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "specialClassesTokenLimits", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpecialClassesTokenLimits is a free data retrieval call binding the contract method 0x74a1dff2.
//
// Solidity: function specialClassesTokenLimits(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentSession) SpecialClassesTokenLimits(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.SpecialClassesTokenLimits(&_Xentorrent.CallOpts, arg0)
}

// SpecialClassesTokenLimits is a free data retrieval call binding the contract method 0x74a1dff2.
//
// Solidity: function specialClassesTokenLimits(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) SpecialClassesTokenLimits(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.SpecialClassesTokenLimits(&_Xentorrent.CallOpts, arg0)
}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_Xentorrent *XentorrentCaller) StartBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "startBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_Xentorrent *XentorrentSession) StartBlockNumber() (*big.Int, error) {
	return _Xentorrent.Contract.StartBlockNumber(&_Xentorrent.CallOpts)
}

// StartBlockNumber is a free data retrieval call binding the contract method 0x498a4c2d.
//
// Solidity: function startBlockNumber() view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) StartBlockNumber() (*big.Int, error) {
	return _Xentorrent.Contract.StartBlockNumber(&_Xentorrent.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Xentorrent *XentorrentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Xentorrent *XentorrentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Xentorrent.Contract.SupportsInterface(&_Xentorrent.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Xentorrent *XentorrentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Xentorrent.Contract.SupportsInterface(&_Xentorrent.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Xentorrent *XentorrentCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Xentorrent *XentorrentSession) Symbol() (string, error) {
	return _Xentorrent.Contract.Symbol(&_Xentorrent.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Xentorrent *XentorrentCallerSession) Symbol() (string, error) {
	return _Xentorrent.Contract.Symbol(&_Xentorrent.CallOpts)
}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_Xentorrent *XentorrentCaller) TokenIdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "tokenIdCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_Xentorrent *XentorrentSession) TokenIdCounter() (*big.Int, error) {
	return _Xentorrent.Contract.TokenIdCounter(&_Xentorrent.CallOpts)
}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) TokenIdCounter() (*big.Int, error) {
	return _Xentorrent.Contract.TokenIdCounter(&_Xentorrent.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Xentorrent *XentorrentCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Xentorrent *XentorrentSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Xentorrent.Contract.TokenURI(&_Xentorrent.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Xentorrent *XentorrentCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Xentorrent.Contract.TokenURI(&_Xentorrent.CallOpts, tokenId)
}

// VmuCount is a free data retrieval call binding the contract method 0xa1a53fa1.
//
// Solidity: function vmuCount(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCaller) VmuCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "vmuCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VmuCount is a free data retrieval call binding the contract method 0xa1a53fa1.
//
// Solidity: function vmuCount(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentSession) VmuCount(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.VmuCount(&_Xentorrent.CallOpts, arg0)
}

// VmuCount is a free data retrieval call binding the contract method 0xa1a53fa1.
//
// Solidity: function vmuCount(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) VmuCount(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.VmuCount(&_Xentorrent.CallOpts, arg0)
}

// XenBurned is a free data retrieval call binding the contract method 0xbd333033.
//
// Solidity: function xenBurned(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCaller) XenBurned(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "xenBurned", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XenBurned is a free data retrieval call binding the contract method 0xbd333033.
//
// Solidity: function xenBurned(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentSession) XenBurned(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.XenBurned(&_Xentorrent.CallOpts, arg0)
}

// XenBurned is a free data retrieval call binding the contract method 0xbd333033.
//
// Solidity: function xenBurned(uint256 ) view returns(uint256)
func (_Xentorrent *XentorrentCallerSession) XenBurned(arg0 *big.Int) (*big.Int, error) {
	return _Xentorrent.Contract.XenBurned(&_Xentorrent.CallOpts, arg0)
}

// XenCrypto is a free data retrieval call binding the contract method 0x71141a58.
//
// Solidity: function xenCrypto() view returns(address)
func (_Xentorrent *XentorrentCaller) XenCrypto(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Xentorrent.contract.Call(opts, &out, "xenCrypto")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XenCrypto is a free data retrieval call binding the contract method 0x71141a58.
//
// Solidity: function xenCrypto() view returns(address)
func (_Xentorrent *XentorrentSession) XenCrypto() (common.Address, error) {
	return _Xentorrent.Contract.XenCrypto(&_Xentorrent.CallOpts)
}

// XenCrypto is a free data retrieval call binding the contract method 0x71141a58.
//
// Solidity: function xenCrypto() view returns(address)
func (_Xentorrent *XentorrentCallerSession) XenCrypto() (common.Address, error) {
	return _Xentorrent.Contract.XenCrypto(&_Xentorrent.CallOpts)
}

// AddForwarder is a paid mutator transaction binding the contract method 0x5c41d2fe.
//
// Solidity: function addForwarder(address trustedForwarder) returns()
func (_Xentorrent *XentorrentTransactor) AddForwarder(opts *bind.TransactOpts, trustedForwarder common.Address) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "addForwarder", trustedForwarder)
}

// AddForwarder is a paid mutator transaction binding the contract method 0x5c41d2fe.
//
// Solidity: function addForwarder(address trustedForwarder) returns()
func (_Xentorrent *XentorrentSession) AddForwarder(trustedForwarder common.Address) (*types.Transaction, error) {
	return _Xentorrent.Contract.AddForwarder(&_Xentorrent.TransactOpts, trustedForwarder)
}

// AddForwarder is a paid mutator transaction binding the contract method 0x5c41d2fe.
//
// Solidity: function addForwarder(address trustedForwarder) returns()
func (_Xentorrent *XentorrentTransactorSession) AddForwarder(trustedForwarder common.Address) (*types.Transaction, error) {
	return _Xentorrent.Contract.AddForwarder(&_Xentorrent.TransactOpts, trustedForwarder)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Xentorrent *XentorrentTransactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Xentorrent *XentorrentSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.Approve(&_Xentorrent.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) returns()
func (_Xentorrent *XentorrentTransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.Approve(&_Xentorrent.TransactOpts, operator, tokenId)
}

// BulkClaimMintReward is a paid mutator transaction binding the contract method 0xf5878b9b.
//
// Solidity: function bulkClaimMintReward(uint256 tokenId, address to) returns()
func (_Xentorrent *XentorrentTransactor) BulkClaimMintReward(opts *bind.TransactOpts, tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "bulkClaimMintReward", tokenId, to)
}

// BulkClaimMintReward is a paid mutator transaction binding the contract method 0xf5878b9b.
//
// Solidity: function bulkClaimMintReward(uint256 tokenId, address to) returns()
func (_Xentorrent *XentorrentSession) BulkClaimMintReward(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Xentorrent.Contract.BulkClaimMintReward(&_Xentorrent.TransactOpts, tokenId, to)
}

// BulkClaimMintReward is a paid mutator transaction binding the contract method 0xf5878b9b.
//
// Solidity: function bulkClaimMintReward(uint256 tokenId, address to) returns()
func (_Xentorrent *XentorrentTransactorSession) BulkClaimMintReward(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Xentorrent.Contract.BulkClaimMintReward(&_Xentorrent.TransactOpts, tokenId, to)
}

// BulkClaimRank is a paid mutator transaction binding the contract method 0xecef9201.
//
// Solidity: function bulkClaimRank(uint256 count, uint256 term) returns(uint256 tokenId)
func (_Xentorrent *XentorrentTransactor) BulkClaimRank(opts *bind.TransactOpts, count *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "bulkClaimRank", count, term)
}

// BulkClaimRank is a paid mutator transaction binding the contract method 0xecef9201.
//
// Solidity: function bulkClaimRank(uint256 count, uint256 term) returns(uint256 tokenId)
func (_Xentorrent *XentorrentSession) BulkClaimRank(count *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.BulkClaimRank(&_Xentorrent.TransactOpts, count, term)
}

// BulkClaimRank is a paid mutator transaction binding the contract method 0xecef9201.
//
// Solidity: function bulkClaimRank(uint256 count, uint256 term) returns(uint256 tokenId)
func (_Xentorrent *XentorrentTransactorSession) BulkClaimRank(count *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.BulkClaimRank(&_Xentorrent.TransactOpts, count, term)
}

// BulkClaimRankLimited is a paid mutator transaction binding the contract method 0x53b18de4.
//
// Solidity: function bulkClaimRankLimited(uint256 count, uint256 term, uint256 burning) returns(uint256)
func (_Xentorrent *XentorrentTransactor) BulkClaimRankLimited(opts *bind.TransactOpts, count *big.Int, term *big.Int, burning *big.Int) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "bulkClaimRankLimited", count, term, burning)
}

// BulkClaimRankLimited is a paid mutator transaction binding the contract method 0x53b18de4.
//
// Solidity: function bulkClaimRankLimited(uint256 count, uint256 term, uint256 burning) returns(uint256)
func (_Xentorrent *XentorrentSession) BulkClaimRankLimited(count *big.Int, term *big.Int, burning *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.BulkClaimRankLimited(&_Xentorrent.TransactOpts, count, term, burning)
}

// BulkClaimRankLimited is a paid mutator transaction binding the contract method 0x53b18de4.
//
// Solidity: function bulkClaimRankLimited(uint256 count, uint256 term, uint256 burning) returns(uint256)
func (_Xentorrent *XentorrentTransactorSession) BulkClaimRankLimited(count *big.Int, term *big.Int, burning *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.BulkClaimRankLimited(&_Xentorrent.TransactOpts, count, term, burning)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 tokenId) returns()
func (_Xentorrent *XentorrentTransactor) Burn(opts *bind.TransactOpts, user common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "burn", user, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 tokenId) returns()
func (_Xentorrent *XentorrentSession) Burn(user common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.Burn(&_Xentorrent.TransactOpts, user, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 tokenId) returns()
func (_Xentorrent *XentorrentTransactorSession) Burn(user common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.Burn(&_Xentorrent.TransactOpts, user, tokenId)
}

// CallClaimMintReward is a paid mutator transaction binding the contract method 0xdf0030ef.
//
// Solidity: function callClaimMintReward(address to) returns()
func (_Xentorrent *XentorrentTransactor) CallClaimMintReward(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "callClaimMintReward", to)
}

// CallClaimMintReward is a paid mutator transaction binding the contract method 0xdf0030ef.
//
// Solidity: function callClaimMintReward(address to) returns()
func (_Xentorrent *XentorrentSession) CallClaimMintReward(to common.Address) (*types.Transaction, error) {
	return _Xentorrent.Contract.CallClaimMintReward(&_Xentorrent.TransactOpts, to)
}

// CallClaimMintReward is a paid mutator transaction binding the contract method 0xdf0030ef.
//
// Solidity: function callClaimMintReward(address to) returns()
func (_Xentorrent *XentorrentTransactorSession) CallClaimMintReward(to common.Address) (*types.Transaction, error) {
	return _Xentorrent.Contract.CallClaimMintReward(&_Xentorrent.TransactOpts, to)
}

// CallClaimRank is a paid mutator transaction binding the contract method 0x01bb4116.
//
// Solidity: function callClaimRank(uint256 term) returns()
func (_Xentorrent *XentorrentTransactor) CallClaimRank(opts *bind.TransactOpts, term *big.Int) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "callClaimRank", term)
}

// CallClaimRank is a paid mutator transaction binding the contract method 0x01bb4116.
//
// Solidity: function callClaimRank(uint256 term) returns()
func (_Xentorrent *XentorrentSession) CallClaimRank(term *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.CallClaimRank(&_Xentorrent.TransactOpts, term)
}

// CallClaimRank is a paid mutator transaction binding the contract method 0x01bb4116.
//
// Solidity: function callClaimRank(uint256 term) returns()
func (_Xentorrent *XentorrentTransactorSession) CallClaimRank(term *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.CallClaimRank(&_Xentorrent.TransactOpts, term)
}

// OnTokenBurned is a paid mutator transaction binding the contract method 0x543746b1.
//
// Solidity: function onTokenBurned(address user, uint256 burned) returns()
func (_Xentorrent *XentorrentTransactor) OnTokenBurned(opts *bind.TransactOpts, user common.Address, burned *big.Int) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "onTokenBurned", user, burned)
}

// OnTokenBurned is a paid mutator transaction binding the contract method 0x543746b1.
//
// Solidity: function onTokenBurned(address user, uint256 burned) returns()
func (_Xentorrent *XentorrentSession) OnTokenBurned(user common.Address, burned *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.OnTokenBurned(&_Xentorrent.TransactOpts, user, burned)
}

// OnTokenBurned is a paid mutator transaction binding the contract method 0x543746b1.
//
// Solidity: function onTokenBurned(address user, uint256 burned) returns()
func (_Xentorrent *XentorrentTransactorSession) OnTokenBurned(user common.Address, burned *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.OnTokenBurned(&_Xentorrent.TransactOpts, user, burned)
}

// PowerDown is a paid mutator transaction binding the contract method 0x928dd2a7.
//
// Solidity: function powerDown() returns()
func (_Xentorrent *XentorrentTransactor) PowerDown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "powerDown")
}

// PowerDown is a paid mutator transaction binding the contract method 0x928dd2a7.
//
// Solidity: function powerDown() returns()
func (_Xentorrent *XentorrentSession) PowerDown() (*types.Transaction, error) {
	return _Xentorrent.Contract.PowerDown(&_Xentorrent.TransactOpts)
}

// PowerDown is a paid mutator transaction binding the contract method 0x928dd2a7.
//
// Solidity: function powerDown() returns()
func (_Xentorrent *XentorrentTransactorSession) PowerDown() (*types.Transaction, error) {
	return _Xentorrent.Contract.PowerDown(&_Xentorrent.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Xentorrent *XentorrentTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Xentorrent *XentorrentSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.SafeTransferFrom(&_Xentorrent.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Xentorrent *XentorrentTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.SafeTransferFrom(&_Xentorrent.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Xentorrent *XentorrentTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Xentorrent *XentorrentSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Xentorrent.Contract.SafeTransferFrom0(&_Xentorrent.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Xentorrent *XentorrentTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Xentorrent.Contract.SafeTransferFrom0(&_Xentorrent.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Xentorrent *XentorrentTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Xentorrent *XentorrentSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Xentorrent.Contract.SetApprovalForAll(&_Xentorrent.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Xentorrent *XentorrentTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Xentorrent.Contract.SetApprovalForAll(&_Xentorrent.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Xentorrent *XentorrentTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Xentorrent *XentorrentSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.TransferFrom(&_Xentorrent.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Xentorrent *XentorrentTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Xentorrent.Contract.TransferFrom(&_Xentorrent.TransactOpts, from, to, tokenId)
}

// XentorrentApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Xentorrent contract.
type XentorrentApprovalIterator struct {
	Event *XentorrentApproval // Event containing the contract specifics and raw log

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
func (it *XentorrentApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XentorrentApproval)
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
		it.Event = new(XentorrentApproval)
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
func (it *XentorrentApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XentorrentApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XentorrentApproval represents a Approval event raised by the Xentorrent contract.
type XentorrentApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Xentorrent *XentorrentFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*XentorrentApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Xentorrent.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &XentorrentApprovalIterator{contract: _Xentorrent.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Xentorrent *XentorrentFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *XentorrentApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Xentorrent.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XentorrentApproval)
				if err := _Xentorrent.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Xentorrent *XentorrentFilterer) ParseApproval(log types.Log) (*XentorrentApproval, error) {
	event := new(XentorrentApproval)
	if err := _Xentorrent.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XentorrentApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Xentorrent contract.
type XentorrentApprovalForAllIterator struct {
	Event *XentorrentApprovalForAll // Event containing the contract specifics and raw log

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
func (it *XentorrentApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XentorrentApprovalForAll)
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
		it.Event = new(XentorrentApprovalForAll)
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
func (it *XentorrentApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XentorrentApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XentorrentApprovalForAll represents a ApprovalForAll event raised by the Xentorrent contract.
type XentorrentApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Xentorrent *XentorrentFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*XentorrentApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Xentorrent.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &XentorrentApprovalForAllIterator{contract: _Xentorrent.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Xentorrent *XentorrentFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *XentorrentApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Xentorrent.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XentorrentApprovalForAll)
				if err := _Xentorrent.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Xentorrent *XentorrentFilterer) ParseApprovalForAll(log types.Log) (*XentorrentApprovalForAll, error) {
	event := new(XentorrentApprovalForAll)
	if err := _Xentorrent.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XentorrentEndTorrentIterator is returned from FilterEndTorrent and is used to iterate over the raw logs and unpacked data for EndTorrent events raised by the Xentorrent contract.
type XentorrentEndTorrentIterator struct {
	Event *XentorrentEndTorrent // Event containing the contract specifics and raw log

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
func (it *XentorrentEndTorrentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XentorrentEndTorrent)
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
		it.Event = new(XentorrentEndTorrent)
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
func (it *XentorrentEndTorrentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XentorrentEndTorrentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XentorrentEndTorrent represents a EndTorrent event raised by the Xentorrent contract.
type XentorrentEndTorrent struct {
	User    common.Address
	TokenId *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEndTorrent is a free log retrieval operation binding the contract event 0x7ae39cb5fb0bebb7775f35a0009e0c94f59c2e40c8967af20842619edac4694d.
//
// Solidity: event EndTorrent(address indexed user, uint256 tokenId, address to)
func (_Xentorrent *XentorrentFilterer) FilterEndTorrent(opts *bind.FilterOpts, user []common.Address) (*XentorrentEndTorrentIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xentorrent.contract.FilterLogs(opts, "EndTorrent", userRule)
	if err != nil {
		return nil, err
	}
	return &XentorrentEndTorrentIterator{contract: _Xentorrent.contract, event: "EndTorrent", logs: logs, sub: sub}, nil
}

// WatchEndTorrent is a free log subscription operation binding the contract event 0x7ae39cb5fb0bebb7775f35a0009e0c94f59c2e40c8967af20842619edac4694d.
//
// Solidity: event EndTorrent(address indexed user, uint256 tokenId, address to)
func (_Xentorrent *XentorrentFilterer) WatchEndTorrent(opts *bind.WatchOpts, sink chan<- *XentorrentEndTorrent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xentorrent.contract.WatchLogs(opts, "EndTorrent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XentorrentEndTorrent)
				if err := _Xentorrent.contract.UnpackLog(event, "EndTorrent", log); err != nil {
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

// ParseEndTorrent is a log parse operation binding the contract event 0x7ae39cb5fb0bebb7775f35a0009e0c94f59c2e40c8967af20842619edac4694d.
//
// Solidity: event EndTorrent(address indexed user, uint256 tokenId, address to)
func (_Xentorrent *XentorrentFilterer) ParseEndTorrent(log types.Log) (*XentorrentEndTorrent, error) {
	event := new(XentorrentEndTorrent)
	if err := _Xentorrent.contract.UnpackLog(event, "EndTorrent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XentorrentRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the Xentorrent contract.
type XentorrentRedeemedIterator struct {
	Event *XentorrentRedeemed // Event containing the contract specifics and raw log

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
func (it *XentorrentRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XentorrentRedeemed)
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
		it.Event = new(XentorrentRedeemed)
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
func (it *XentorrentRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XentorrentRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XentorrentRedeemed represents a Redeemed event raised by the Xentorrent contract.
type XentorrentRedeemed struct {
	User          common.Address
	XenContract   common.Address
	TokenContract common.Address
	XenAmount     *big.Int
	TokenAmount   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0xda4c370e4f539b8e19d66144cad73d70fbf440a544e56ff38e28e38ed47801a8.
//
// Solidity: event Redeemed(address indexed user, address indexed xenContract, address indexed tokenContract, uint256 xenAmount, uint256 tokenAmount)
func (_Xentorrent *XentorrentFilterer) FilterRedeemed(opts *bind.FilterOpts, user []common.Address, xenContract []common.Address, tokenContract []common.Address) (*XentorrentRedeemedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var xenContractRule []interface{}
	for _, xenContractItem := range xenContract {
		xenContractRule = append(xenContractRule, xenContractItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Xentorrent.contract.FilterLogs(opts, "Redeemed", userRule, xenContractRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &XentorrentRedeemedIterator{contract: _Xentorrent.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0xda4c370e4f539b8e19d66144cad73d70fbf440a544e56ff38e28e38ed47801a8.
//
// Solidity: event Redeemed(address indexed user, address indexed xenContract, address indexed tokenContract, uint256 xenAmount, uint256 tokenAmount)
func (_Xentorrent *XentorrentFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *XentorrentRedeemed, user []common.Address, xenContract []common.Address, tokenContract []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var xenContractRule []interface{}
	for _, xenContractItem := range xenContract {
		xenContractRule = append(xenContractRule, xenContractItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _Xentorrent.contract.WatchLogs(opts, "Redeemed", userRule, xenContractRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XentorrentRedeemed)
				if err := _Xentorrent.contract.UnpackLog(event, "Redeemed", log); err != nil {
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

// ParseRedeemed is a log parse operation binding the contract event 0xda4c370e4f539b8e19d66144cad73d70fbf440a544e56ff38e28e38ed47801a8.
//
// Solidity: event Redeemed(address indexed user, address indexed xenContract, address indexed tokenContract, uint256 xenAmount, uint256 tokenAmount)
func (_Xentorrent *XentorrentFilterer) ParseRedeemed(log types.Log) (*XentorrentRedeemed, error) {
	event := new(XentorrentRedeemed)
	if err := _Xentorrent.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XentorrentStartTorrentIterator is returned from FilterStartTorrent and is used to iterate over the raw logs and unpacked data for StartTorrent events raised by the Xentorrent contract.
type XentorrentStartTorrentIterator struct {
	Event *XentorrentStartTorrent // Event containing the contract specifics and raw log

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
func (it *XentorrentStartTorrentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XentorrentStartTorrent)
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
		it.Event = new(XentorrentStartTorrent)
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
func (it *XentorrentStartTorrentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XentorrentStartTorrentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XentorrentStartTorrent represents a StartTorrent event raised by the Xentorrent contract.
type XentorrentStartTorrent struct {
	User  common.Address
	Count *big.Int
	Term  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStartTorrent is a free log retrieval operation binding the contract event 0xfbb2077593b3594fd0ac359a2d898268191a74843aaf1ba3f517b5514a1b0711.
//
// Solidity: event StartTorrent(address indexed user, uint256 count, uint256 term)
func (_Xentorrent *XentorrentFilterer) FilterStartTorrent(opts *bind.FilterOpts, user []common.Address) (*XentorrentStartTorrentIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xentorrent.contract.FilterLogs(opts, "StartTorrent", userRule)
	if err != nil {
		return nil, err
	}
	return &XentorrentStartTorrentIterator{contract: _Xentorrent.contract, event: "StartTorrent", logs: logs, sub: sub}, nil
}

// WatchStartTorrent is a free log subscription operation binding the contract event 0xfbb2077593b3594fd0ac359a2d898268191a74843aaf1ba3f517b5514a1b0711.
//
// Solidity: event StartTorrent(address indexed user, uint256 count, uint256 term)
func (_Xentorrent *XentorrentFilterer) WatchStartTorrent(opts *bind.WatchOpts, sink chan<- *XentorrentStartTorrent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Xentorrent.contract.WatchLogs(opts, "StartTorrent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XentorrentStartTorrent)
				if err := _Xentorrent.contract.UnpackLog(event, "StartTorrent", log); err != nil {
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

// ParseStartTorrent is a log parse operation binding the contract event 0xfbb2077593b3594fd0ac359a2d898268191a74843aaf1ba3f517b5514a1b0711.
//
// Solidity: event StartTorrent(address indexed user, uint256 count, uint256 term)
func (_Xentorrent *XentorrentFilterer) ParseStartTorrent(log types.Log) (*XentorrentStartTorrent, error) {
	event := new(XentorrentStartTorrent)
	if err := _Xentorrent.contract.UnpackLog(event, "StartTorrent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// XentorrentTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Xentorrent contract.
type XentorrentTransferIterator struct {
	Event *XentorrentTransfer // Event containing the contract specifics and raw log

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
func (it *XentorrentTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(XentorrentTransfer)
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
		it.Event = new(XentorrentTransfer)
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
func (it *XentorrentTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *XentorrentTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// XentorrentTransfer represents a Transfer event raised by the Xentorrent contract.
type XentorrentTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Xentorrent *XentorrentFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*XentorrentTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Xentorrent.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &XentorrentTransferIterator{contract: _Xentorrent.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Xentorrent *XentorrentFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *XentorrentTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Xentorrent.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(XentorrentTransfer)
				if err := _Xentorrent.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Xentorrent *XentorrentFilterer) ParseTransfer(log types.Log) (*XentorrentTransfer, error) {
	event := new(XentorrentTransfer)
	if err := _Xentorrent.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
