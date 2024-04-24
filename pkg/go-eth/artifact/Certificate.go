// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goeth

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

// CertificateOfCompletionFileMetaData is an auto generated low-level Go binding around an user-defined struct.
type CertificateOfCompletionFileMetaData struct {
	Md5 string
}

// CertificateOfCompletionPerson is an auto generated low-level Go binding around an user-defined struct.
type CertificateOfCompletionPerson struct {
	Name  string
	Email string
}

// CertificateMetaData contains all meta data concerning the Certificate contract.
var CertificateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_recipientName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_recipientEmail\",\"type\":\"string\"}],\"name\":\"createCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_newMd5\",\"type\":\"string\"}],\"name\":\"updateMd5Certificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"certificates\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"md5\",\"type\":\"string\"}],\"internalType\":\"structCertificateOfCompletion.FileMetaData\",\"name\":\"fileMetaData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"}],\"internalType\":\"structCertificateOfCompletion.Person\",\"name\":\"person\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"fileMd5toAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_recipientName\",\"type\":\"string\"}],\"name\":\"getAddressByName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"getCertificate\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_md5\",\"type\":\"string\"}],\"name\":\"getCertificateByFileName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nameToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CertificateABI is the input ABI used to generate the binding from.
// Deprecated: Use CertificateMetaData.ABI instead.
var CertificateABI = CertificateMetaData.ABI

// Certificate is an auto generated Go binding around an Ethereum contract.
type Certificate struct {
	CertificateCaller     // Read-only binding to the contract
	CertificateTransactor // Write-only binding to the contract
	CertificateFilterer   // Log filterer for contract events
}

// CertificateCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertificateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertificateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertificateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertificateSession struct {
	Contract     *Certificate      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertificateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertificateCallerSession struct {
	Contract *CertificateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CertificateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertificateTransactorSession struct {
	Contract     *CertificateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CertificateRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertificateRaw struct {
	Contract *Certificate // Generic contract binding to access the raw methods on
}

// CertificateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertificateCallerRaw struct {
	Contract *CertificateCaller // Generic read-only contract binding to access the raw methods on
}

// CertificateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertificateTransactorRaw struct {
	Contract *CertificateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertificate creates a new instance of Certificate, bound to a specific deployed contract.
func NewCertificate(address common.Address, backend bind.ContractBackend) (*Certificate, error) {
	contract, err := bindCertificate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Certificate{CertificateCaller: CertificateCaller{contract: contract}, CertificateTransactor: CertificateTransactor{contract: contract}, CertificateFilterer: CertificateFilterer{contract: contract}}, nil
}

// NewCertificateCaller creates a new read-only instance of Certificate, bound to a specific deployed contract.
func NewCertificateCaller(address common.Address, caller bind.ContractCaller) (*CertificateCaller, error) {
	contract, err := bindCertificate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateCaller{contract: contract}, nil
}

// NewCertificateTransactor creates a new write-only instance of Certificate, bound to a specific deployed contract.
func NewCertificateTransactor(address common.Address, transactor bind.ContractTransactor) (*CertificateTransactor, error) {
	contract, err := bindCertificate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateTransactor{contract: contract}, nil
}

// NewCertificateFilterer creates a new log filterer instance of Certificate, bound to a specific deployed contract.
func NewCertificateFilterer(address common.Address, filterer bind.ContractFilterer) (*CertificateFilterer, error) {
	contract, err := bindCertificate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertificateFilterer{contract: contract}, nil
}

// bindCertificate binds a generic wrapper to an already deployed contract.
func bindCertificate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CertificateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Certificate *CertificateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Certificate.Contract.CertificateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Certificate *CertificateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Certificate.Contract.CertificateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Certificate *CertificateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Certificate.Contract.CertificateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Certificate *CertificateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Certificate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Certificate *CertificateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Certificate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Certificate *CertificateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Certificate.Contract.contract.Transact(opts, method, params...)
}

// Certificates is a free data retrieval call binding the contract method 0x0016e526.
//
// Solidity: function certificates(address ) view returns((string) fileMetaData, address recipient, (string,string) person, uint256 timestamp)
func (_Certificate *CertificateCaller) Certificates(opts *bind.CallOpts, arg0 common.Address) (struct {
	FileMetaData CertificateOfCompletionFileMetaData
	Recipient    common.Address
	Person       CertificateOfCompletionPerson
	Timestamp    *big.Int
}, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "certificates", arg0)

	outstruct := new(struct {
		FileMetaData CertificateOfCompletionFileMetaData
		Recipient    common.Address
		Person       CertificateOfCompletionPerson
		Timestamp    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FileMetaData = *abi.ConvertType(out[0], new(CertificateOfCompletionFileMetaData)).(*CertificateOfCompletionFileMetaData)
	outstruct.Recipient = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Person = *abi.ConvertType(out[2], new(CertificateOfCompletionPerson)).(*CertificateOfCompletionPerson)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Certificates is a free data retrieval call binding the contract method 0x0016e526.
//
// Solidity: function certificates(address ) view returns((string) fileMetaData, address recipient, (string,string) person, uint256 timestamp)
func (_Certificate *CertificateSession) Certificates(arg0 common.Address) (struct {
	FileMetaData CertificateOfCompletionFileMetaData
	Recipient    common.Address
	Person       CertificateOfCompletionPerson
	Timestamp    *big.Int
}, error) {
	return _Certificate.Contract.Certificates(&_Certificate.CallOpts, arg0)
}

// Certificates is a free data retrieval call binding the contract method 0x0016e526.
//
// Solidity: function certificates(address ) view returns((string) fileMetaData, address recipient, (string,string) person, uint256 timestamp)
func (_Certificate *CertificateCallerSession) Certificates(arg0 common.Address) (struct {
	FileMetaData CertificateOfCompletionFileMetaData
	Recipient    common.Address
	Person       CertificateOfCompletionPerson
	Timestamp    *big.Int
}, error) {
	return _Certificate.Contract.Certificates(&_Certificate.CallOpts, arg0)
}

// FileMd5toAddress is a free data retrieval call binding the contract method 0x99ccd007.
//
// Solidity: function fileMd5toAddress(string ) view returns(address)
func (_Certificate *CertificateCaller) FileMd5toAddress(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "fileMd5toAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FileMd5toAddress is a free data retrieval call binding the contract method 0x99ccd007.
//
// Solidity: function fileMd5toAddress(string ) view returns(address)
func (_Certificate *CertificateSession) FileMd5toAddress(arg0 string) (common.Address, error) {
	return _Certificate.Contract.FileMd5toAddress(&_Certificate.CallOpts, arg0)
}

// FileMd5toAddress is a free data retrieval call binding the contract method 0x99ccd007.
//
// Solidity: function fileMd5toAddress(string ) view returns(address)
func (_Certificate *CertificateCallerSession) FileMd5toAddress(arg0 string) (common.Address, error) {
	return _Certificate.Contract.FileMd5toAddress(&_Certificate.CallOpts, arg0)
}

// GetAddressByName is a free data retrieval call binding the contract method 0x9a65ddec.
//
// Solidity: function getAddressByName(string _recipientName) view returns(address)
func (_Certificate *CertificateCaller) GetAddressByName(opts *bind.CallOpts, _recipientName string) (common.Address, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "getAddressByName", _recipientName)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressByName is a free data retrieval call binding the contract method 0x9a65ddec.
//
// Solidity: function getAddressByName(string _recipientName) view returns(address)
func (_Certificate *CertificateSession) GetAddressByName(_recipientName string) (common.Address, error) {
	return _Certificate.Contract.GetAddressByName(&_Certificate.CallOpts, _recipientName)
}

// GetAddressByName is a free data retrieval call binding the contract method 0x9a65ddec.
//
// Solidity: function getAddressByName(string _recipientName) view returns(address)
func (_Certificate *CertificateCallerSession) GetAddressByName(_recipientName string) (common.Address, error) {
	return _Certificate.Contract.GetAddressByName(&_Certificate.CallOpts, _recipientName)
}

// GetCertificate is a free data retrieval call binding the contract method 0xfd531e93.
//
// Solidity: function getCertificate(address _recipient) view returns(string, string, string, address, uint256)
func (_Certificate *CertificateCaller) GetCertificate(opts *bind.CallOpts, _recipient common.Address) (string, string, string, common.Address, *big.Int, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "getCertificate", _recipient)

	if err != nil {
		return *new(string), *new(string), *new(string), *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetCertificate is a free data retrieval call binding the contract method 0xfd531e93.
//
// Solidity: function getCertificate(address _recipient) view returns(string, string, string, address, uint256)
func (_Certificate *CertificateSession) GetCertificate(_recipient common.Address) (string, string, string, common.Address, *big.Int, error) {
	return _Certificate.Contract.GetCertificate(&_Certificate.CallOpts, _recipient)
}

// GetCertificate is a free data retrieval call binding the contract method 0xfd531e93.
//
// Solidity: function getCertificate(address _recipient) view returns(string, string, string, address, uint256)
func (_Certificate *CertificateCallerSession) GetCertificate(_recipient common.Address) (string, string, string, common.Address, *big.Int, error) {
	return _Certificate.Contract.GetCertificate(&_Certificate.CallOpts, _recipient)
}

// GetCertificateByFileName is a free data retrieval call binding the contract method 0x10d04a5c.
//
// Solidity: function getCertificateByFileName(string _md5) view returns(address)
func (_Certificate *CertificateCaller) GetCertificateByFileName(opts *bind.CallOpts, _md5 string) (common.Address, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "getCertificateByFileName", _md5)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCertificateByFileName is a free data retrieval call binding the contract method 0x10d04a5c.
//
// Solidity: function getCertificateByFileName(string _md5) view returns(address)
func (_Certificate *CertificateSession) GetCertificateByFileName(_md5 string) (common.Address, error) {
	return _Certificate.Contract.GetCertificateByFileName(&_Certificate.CallOpts, _md5)
}

// GetCertificateByFileName is a free data retrieval call binding the contract method 0x10d04a5c.
//
// Solidity: function getCertificateByFileName(string _md5) view returns(address)
func (_Certificate *CertificateCallerSession) GetCertificateByFileName(_md5 string) (common.Address, error) {
	return _Certificate.Contract.GetCertificateByFileName(&_Certificate.CallOpts, _md5)
}

// NameToAddress is a free data retrieval call binding the contract method 0x08d88aad.
//
// Solidity: function nameToAddress(string ) view returns(address)
func (_Certificate *CertificateCaller) NameToAddress(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "nameToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NameToAddress is a free data retrieval call binding the contract method 0x08d88aad.
//
// Solidity: function nameToAddress(string ) view returns(address)
func (_Certificate *CertificateSession) NameToAddress(arg0 string) (common.Address, error) {
	return _Certificate.Contract.NameToAddress(&_Certificate.CallOpts, arg0)
}

// NameToAddress is a free data retrieval call binding the contract method 0x08d88aad.
//
// Solidity: function nameToAddress(string ) view returns(address)
func (_Certificate *CertificateCallerSession) NameToAddress(arg0 string) (common.Address, error) {
	return _Certificate.Contract.NameToAddress(&_Certificate.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Certificate *CertificateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Certificate *CertificateSession) Owner() (common.Address, error) {
	return _Certificate.Contract.Owner(&_Certificate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Certificate *CertificateCallerSession) Owner() (common.Address, error) {
	return _Certificate.Contract.Owner(&_Certificate.CallOpts)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xf02c5a20.
//
// Solidity: function createCertificate(string _recipientName, string _recipientEmail) returns()
func (_Certificate *CertificateTransactor) CreateCertificate(opts *bind.TransactOpts, _recipientName string, _recipientEmail string) (*types.Transaction, error) {
	return _Certificate.contract.Transact(opts, "createCertificate", _recipientName, _recipientEmail)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xf02c5a20.
//
// Solidity: function createCertificate(string _recipientName, string _recipientEmail) returns()
func (_Certificate *CertificateSession) CreateCertificate(_recipientName string, _recipientEmail string) (*types.Transaction, error) {
	return _Certificate.Contract.CreateCertificate(&_Certificate.TransactOpts, _recipientName, _recipientEmail)
}

// CreateCertificate is a paid mutator transaction binding the contract method 0xf02c5a20.
//
// Solidity: function createCertificate(string _recipientName, string _recipientEmail) returns()
func (_Certificate *CertificateTransactorSession) CreateCertificate(_recipientName string, _recipientEmail string) (*types.Transaction, error) {
	return _Certificate.Contract.CreateCertificate(&_Certificate.TransactOpts, _recipientName, _recipientEmail)
}

// UpdateMd5Certificate is a paid mutator transaction binding the contract method 0xe2dab62d.
//
// Solidity: function updateMd5Certificate(address _recipient, string _newMd5) returns()
func (_Certificate *CertificateTransactor) UpdateMd5Certificate(opts *bind.TransactOpts, _recipient common.Address, _newMd5 string) (*types.Transaction, error) {
	return _Certificate.contract.Transact(opts, "updateMd5Certificate", _recipient, _newMd5)
}

// UpdateMd5Certificate is a paid mutator transaction binding the contract method 0xe2dab62d.
//
// Solidity: function updateMd5Certificate(address _recipient, string _newMd5) returns()
func (_Certificate *CertificateSession) UpdateMd5Certificate(_recipient common.Address, _newMd5 string) (*types.Transaction, error) {
	return _Certificate.Contract.UpdateMd5Certificate(&_Certificate.TransactOpts, _recipient, _newMd5)
}

// UpdateMd5Certificate is a paid mutator transaction binding the contract method 0xe2dab62d.
//
// Solidity: function updateMd5Certificate(address _recipient, string _newMd5) returns()
func (_Certificate *CertificateTransactorSession) UpdateMd5Certificate(_recipient common.Address, _newMd5 string) (*types.Transaction, error) {
	return _Certificate.Contract.UpdateMd5Certificate(&_Certificate.TransactOpts, _recipient, _newMd5)
}
