// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package whitelist

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

// WhitelistMetaData contains all meta data concerning the Whitelist contract.
var WhitelistMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sanctions\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"functionSig\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"PublicCapabilityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"functionSig\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"RoleCapabilityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"UserRoleUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"functionSig\",\"type\":\"bytes4\"}],\"name\":\"canCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"functionSig\",\"type\":\"bytes4\"}],\"name\":\"doesRoleHaveCapability\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"doesUserHaveRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"getRolesWithCapability\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getUserRoles\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"isCapabilityPublic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sanctions\",\"outputs\":[{\"internalType\":\"contractISanctions\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"functionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setPublicCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"functionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setRoleCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setUserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}];",
}

// WhitelistABI is the input ABI used to generate the binding from.
// Deprecated: Use WhitelistMetaData.ABI instead.
var WhitelistABI = WhitelistMetaData.ABI

// Whitelist is an auto generated Go binding around an Ethereum contract.
type Whitelist struct {
	WhitelistCaller     // Read-only binding to the contract
	WhitelistTransactor // Write-only binding to the contract
	WhitelistFilterer   // Log filterer for contract events
}

// WhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistSession struct {
	Contract     *Whitelist        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistCallerSession struct {
	Contract *WhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistTransactorSession struct {
	Contract     *WhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistRaw struct {
	Contract *Whitelist // Generic contract binding to access the raw methods on
}

// WhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistCallerRaw struct {
	Contract *WhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistTransactorRaw struct {
	Contract *WhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelist creates a new instance of Whitelist, bound to a specific deployed contract.
func NewWhitelist(address common.Address, backend bind.ContractBackend) (*Whitelist, error) {
	contract, err := bindWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// NewWhitelistCaller creates a new read-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistCaller(address common.Address, caller bind.ContractCaller) (*WhitelistCaller, error) {
	contract, err := bindWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistCaller{contract: contract}, nil
}

// NewWhitelistTransactor creates a new write-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistTransactor, error) {
	contract, err := bindWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistTransactor{contract: contract}, nil
}

// NewWhitelistFilterer creates a new log filterer instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistFilterer, error) {
	contract, err := bindWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistFilterer{contract: contract}, nil
}

// bindWhitelist binds a generic wrapper to an already deployed contract.
func bindWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WhitelistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.WhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Whitelist *WhitelistCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Whitelist *WhitelistSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Whitelist.Contract.UPGRADEINTERFACEVERSION(&_Whitelist.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Whitelist *WhitelistCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Whitelist.Contract.UPGRADEINTERFACEVERSION(&_Whitelist.CallOpts)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address user, address target, bytes4 functionSig) view returns(bool)
func (_Whitelist *WhitelistCaller) CanCall(opts *bind.CallOpts, user common.Address, target common.Address, functionSig [4]byte) (bool, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "canCall", user, target, functionSig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address user, address target, bytes4 functionSig) view returns(bool)
func (_Whitelist *WhitelistSession) CanCall(user common.Address, target common.Address, functionSig [4]byte) (bool, error) {
	return _Whitelist.Contract.CanCall(&_Whitelist.CallOpts, user, target, functionSig)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address user, address target, bytes4 functionSig) view returns(bool)
func (_Whitelist *WhitelistCallerSession) CanCall(user common.Address, target common.Address, functionSig [4]byte) (bool, error) {
	return _Whitelist.Contract.CanCall(&_Whitelist.CallOpts, user, target, functionSig)
}

// DoesRoleHaveCapability is a free data retrieval call binding the contract method 0xb4bad06a.
//
// Solidity: function doesRoleHaveCapability(uint8 role, address target, bytes4 functionSig) view returns(bool)
func (_Whitelist *WhitelistCaller) DoesRoleHaveCapability(opts *bind.CallOpts, role uint8, target common.Address, functionSig [4]byte) (bool, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "doesRoleHaveCapability", role, target, functionSig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoesRoleHaveCapability is a free data retrieval call binding the contract method 0xb4bad06a.
//
// Solidity: function doesRoleHaveCapability(uint8 role, address target, bytes4 functionSig) view returns(bool)
func (_Whitelist *WhitelistSession) DoesRoleHaveCapability(role uint8, target common.Address, functionSig [4]byte) (bool, error) {
	return _Whitelist.Contract.DoesRoleHaveCapability(&_Whitelist.CallOpts, role, target, functionSig)
}

// DoesRoleHaveCapability is a free data retrieval call binding the contract method 0xb4bad06a.
//
// Solidity: function doesRoleHaveCapability(uint8 role, address target, bytes4 functionSig) view returns(bool)
func (_Whitelist *WhitelistCallerSession) DoesRoleHaveCapability(role uint8, target common.Address, functionSig [4]byte) (bool, error) {
	return _Whitelist.Contract.DoesRoleHaveCapability(&_Whitelist.CallOpts, role, target, functionSig)
}

// DoesUserHaveRole is a free data retrieval call binding the contract method 0xea7ca276.
//
// Solidity: function doesUserHaveRole(address user, uint8 role) view returns(bool)
func (_Whitelist *WhitelistCaller) DoesUserHaveRole(opts *bind.CallOpts, user common.Address, role uint8) (bool, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "doesUserHaveRole", user, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DoesUserHaveRole is a free data retrieval call binding the contract method 0xea7ca276.
//
// Solidity: function doesUserHaveRole(address user, uint8 role) view returns(bool)
func (_Whitelist *WhitelistSession) DoesUserHaveRole(user common.Address, role uint8) (bool, error) {
	return _Whitelist.Contract.DoesUserHaveRole(&_Whitelist.CallOpts, user, role)
}

// DoesUserHaveRole is a free data retrieval call binding the contract method 0xea7ca276.
//
// Solidity: function doesUserHaveRole(address user, uint8 role) view returns(bool)
func (_Whitelist *WhitelistCallerSession) DoesUserHaveRole(user common.Address, role uint8) (bool, error) {
	return _Whitelist.Contract.DoesUserHaveRole(&_Whitelist.CallOpts, user, role)
}

// GetRolesWithCapability is a free data retrieval call binding the contract method 0x7917b794.
//
// Solidity: function getRolesWithCapability(address , bytes4 ) view returns(bytes32)
func (_Whitelist *WhitelistCaller) GetRolesWithCapability(opts *bind.CallOpts, arg0 common.Address, arg1 [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "getRolesWithCapability", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRolesWithCapability is a free data retrieval call binding the contract method 0x7917b794.
//
// Solidity: function getRolesWithCapability(address , bytes4 ) view returns(bytes32)
func (_Whitelist *WhitelistSession) GetRolesWithCapability(arg0 common.Address, arg1 [4]byte) ([32]byte, error) {
	return _Whitelist.Contract.GetRolesWithCapability(&_Whitelist.CallOpts, arg0, arg1)
}

// GetRolesWithCapability is a free data retrieval call binding the contract method 0x7917b794.
//
// Solidity: function getRolesWithCapability(address , bytes4 ) view returns(bytes32)
func (_Whitelist *WhitelistCallerSession) GetRolesWithCapability(arg0 common.Address, arg1 [4]byte) ([32]byte, error) {
	return _Whitelist.Contract.GetRolesWithCapability(&_Whitelist.CallOpts, arg0, arg1)
}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address ) view returns(bytes32)
func (_Whitelist *WhitelistCaller) GetUserRoles(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "getUserRoles", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address ) view returns(bytes32)
func (_Whitelist *WhitelistSession) GetUserRoles(arg0 common.Address) ([32]byte, error) {
	return _Whitelist.Contract.GetUserRoles(&_Whitelist.CallOpts, arg0)
}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address ) view returns(bytes32)
func (_Whitelist *WhitelistCallerSession) GetUserRoles(arg0 common.Address) ([32]byte, error) {
	return _Whitelist.Contract.GetUserRoles(&_Whitelist.CallOpts, arg0)
}

// IsCapabilityPublic is a free data retrieval call binding the contract method 0x2f47571f.
//
// Solidity: function isCapabilityPublic(address , bytes4 ) view returns(bool)
func (_Whitelist *WhitelistCaller) IsCapabilityPublic(opts *bind.CallOpts, arg0 common.Address, arg1 [4]byte) (bool, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "isCapabilityPublic", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCapabilityPublic is a free data retrieval call binding the contract method 0x2f47571f.
//
// Solidity: function isCapabilityPublic(address , bytes4 ) view returns(bool)
func (_Whitelist *WhitelistSession) IsCapabilityPublic(arg0 common.Address, arg1 [4]byte) (bool, error) {
	return _Whitelist.Contract.IsCapabilityPublic(&_Whitelist.CallOpts, arg0, arg1)
}

// IsCapabilityPublic is a free data retrieval call binding the contract method 0x2f47571f.
//
// Solidity: function isCapabilityPublic(address , bytes4 ) view returns(bool)
func (_Whitelist *WhitelistCallerSession) IsCapabilityPublic(arg0 common.Address, arg1 [4]byte) (bool, error) {
	return _Whitelist.Contract.IsCapabilityPublic(&_Whitelist.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Whitelist *WhitelistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Whitelist *WhitelistSession) Owner() (common.Address, error) {
	return _Whitelist.Contract.Owner(&_Whitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Whitelist *WhitelistCallerSession) Owner() (common.Address, error) {
	return _Whitelist.Contract.Owner(&_Whitelist.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Whitelist *WhitelistCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Whitelist *WhitelistSession) ProxiableUUID() ([32]byte, error) {
	return _Whitelist.Contract.ProxiableUUID(&_Whitelist.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Whitelist *WhitelistCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Whitelist.Contract.ProxiableUUID(&_Whitelist.CallOpts)
}

// Sanctions is a free data retrieval call binding the contract method 0x1a12d585.
//
// Solidity: function sanctions() view returns(address)
func (_Whitelist *WhitelistCaller) Sanctions(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "sanctions")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sanctions is a free data retrieval call binding the contract method 0x1a12d585.
//
// Solidity: function sanctions() view returns(address)
func (_Whitelist *WhitelistSession) Sanctions() (common.Address, error) {
	return _Whitelist.Contract.Sanctions(&_Whitelist.CallOpts)
}

// Sanctions is a free data retrieval call binding the contract method 0x1a12d585.
//
// Solidity: function sanctions() view returns(address)
func (_Whitelist *WhitelistCallerSession) Sanctions() (common.Address, error) {
	return _Whitelist.Contract.Sanctions(&_Whitelist.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Whitelist *WhitelistTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Whitelist *WhitelistSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.Initialize(&_Whitelist.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_Whitelist *WhitelistTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.Initialize(&_Whitelist.TransactOpts, _owner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Whitelist *WhitelistTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Whitelist *WhitelistSession) Pause() (*types.Transaction, error) {
	return _Whitelist.Contract.Pause(&_Whitelist.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Whitelist *WhitelistTransactorSession) Pause() (*types.Transaction, error) {
	return _Whitelist.Contract.Pause(&_Whitelist.TransactOpts)
}

// SetPublicCapability is a paid mutator transaction binding the contract method 0xc6b0263e.
//
// Solidity: function setPublicCapability(address target, bytes4 functionSig, bool enabled) returns()
func (_Whitelist *WhitelistTransactor) SetPublicCapability(opts *bind.TransactOpts, target common.Address, functionSig [4]byte, enabled bool) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "setPublicCapability", target, functionSig, enabled)
}

// SetPublicCapability is a paid mutator transaction binding the contract method 0xc6b0263e.
//
// Solidity: function setPublicCapability(address target, bytes4 functionSig, bool enabled) returns()
func (_Whitelist *WhitelistSession) SetPublicCapability(target common.Address, functionSig [4]byte, enabled bool) (*types.Transaction, error) {
	return _Whitelist.Contract.SetPublicCapability(&_Whitelist.TransactOpts, target, functionSig, enabled)
}

// SetPublicCapability is a paid mutator transaction binding the contract method 0xc6b0263e.
//
// Solidity: function setPublicCapability(address target, bytes4 functionSig, bool enabled) returns()
func (_Whitelist *WhitelistTransactorSession) SetPublicCapability(target common.Address, functionSig [4]byte, enabled bool) (*types.Transaction, error) {
	return _Whitelist.Contract.SetPublicCapability(&_Whitelist.TransactOpts, target, functionSig, enabled)
}

// SetRoleCapability is a paid mutator transaction binding the contract method 0x7d40583d.
//
// Solidity: function setRoleCapability(uint8 role, address target, bytes4 functionSig, bool enabled) returns()
func (_Whitelist *WhitelistTransactor) SetRoleCapability(opts *bind.TransactOpts, role uint8, target common.Address, functionSig [4]byte, enabled bool) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "setRoleCapability", role, target, functionSig, enabled)
}

// SetRoleCapability is a paid mutator transaction binding the contract method 0x7d40583d.
//
// Solidity: function setRoleCapability(uint8 role, address target, bytes4 functionSig, bool enabled) returns()
func (_Whitelist *WhitelistSession) SetRoleCapability(role uint8, target common.Address, functionSig [4]byte, enabled bool) (*types.Transaction, error) {
	return _Whitelist.Contract.SetRoleCapability(&_Whitelist.TransactOpts, role, target, functionSig, enabled)
}

// SetRoleCapability is a paid mutator transaction binding the contract method 0x7d40583d.
//
// Solidity: function setRoleCapability(uint8 role, address target, bytes4 functionSig, bool enabled) returns()
func (_Whitelist *WhitelistTransactorSession) SetRoleCapability(role uint8, target common.Address, functionSig [4]byte, enabled bool) (*types.Transaction, error) {
	return _Whitelist.Contract.SetRoleCapability(&_Whitelist.TransactOpts, role, target, functionSig, enabled)
}

// SetUserRole is a paid mutator transaction binding the contract method 0x67aff484.
//
// Solidity: function setUserRole(address user, uint8 role, bool enabled) returns()
func (_Whitelist *WhitelistTransactor) SetUserRole(opts *bind.TransactOpts, user common.Address, role uint8, enabled bool) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "setUserRole", user, role, enabled)
}

// SetUserRole is a paid mutator transaction binding the contract method 0x67aff484.
//
// Solidity: function setUserRole(address user, uint8 role, bool enabled) returns()
func (_Whitelist *WhitelistSession) SetUserRole(user common.Address, role uint8, enabled bool) (*types.Transaction, error) {
	return _Whitelist.Contract.SetUserRole(&_Whitelist.TransactOpts, user, role, enabled)
}

// SetUserRole is a paid mutator transaction binding the contract method 0x67aff484.
//
// Solidity: function setUserRole(address user, uint8 role, bool enabled) returns()
func (_Whitelist *WhitelistTransactorSession) SetUserRole(user common.Address, role uint8, enabled bool) (*types.Transaction, error) {
	return _Whitelist.Contract.SetUserRole(&_Whitelist.TransactOpts, user, role, enabled)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Whitelist *WhitelistTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Whitelist *WhitelistSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.TransferOwnership(&_Whitelist.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Whitelist *WhitelistTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.TransferOwnership(&_Whitelist.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Whitelist *WhitelistTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Whitelist *WhitelistSession) Unpause() (*types.Transaction, error) {
	return _Whitelist.Contract.Unpause(&_Whitelist.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Whitelist *WhitelistTransactorSession) Unpause() (*types.Transaction, error) {
	return _Whitelist.Contract.Unpause(&_Whitelist.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Whitelist *WhitelistTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Whitelist *WhitelistSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Whitelist.Contract.UpgradeToAndCall(&_Whitelist.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Whitelist *WhitelistTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Whitelist.Contract.UpgradeToAndCall(&_Whitelist.TransactOpts, newImplementation, data)
}

// WhitelistInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Whitelist contract.
type WhitelistInitializedIterator struct {
	Event *WhitelistInitialized // Event containing the contract specifics and raw log

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
func (it *WhitelistInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistInitialized)
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
		it.Event = new(WhitelistInitialized)
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
func (it *WhitelistInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistInitialized represents a Initialized event raised by the Whitelist contract.
type WhitelistInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Whitelist *WhitelistFilterer) FilterInitialized(opts *bind.FilterOpts) (*WhitelistInitializedIterator, error) {

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WhitelistInitializedIterator{contract: _Whitelist.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Whitelist *WhitelistFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WhitelistInitialized) (event.Subscription, error) {

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistInitialized)
				if err := _Whitelist.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Whitelist *WhitelistFilterer) ParseInitialized(log types.Log) (*WhitelistInitialized, error) {
	event := new(WhitelistInitialized)
	if err := _Whitelist.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Whitelist contract.
type WhitelistOwnershipTransferredIterator struct {
	Event *WhitelistOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WhitelistOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistOwnershipTransferred)
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
		it.Event = new(WhitelistOwnershipTransferred)
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
func (it *WhitelistOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistOwnershipTransferred represents a OwnershipTransferred event raised by the Whitelist contract.
type WhitelistOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Whitelist *WhitelistFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*WhitelistOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistOwnershipTransferredIterator{contract: _Whitelist.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Whitelist *WhitelistFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WhitelistOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistOwnershipTransferred)
				if err := _Whitelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Whitelist *WhitelistFilterer) ParseOwnershipTransferred(log types.Log) (*WhitelistOwnershipTransferred, error) {
	event := new(WhitelistOwnershipTransferred)
	if err := _Whitelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Whitelist contract.
type WhitelistPausedIterator struct {
	Event *WhitelistPaused // Event containing the contract specifics and raw log

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
func (it *WhitelistPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistPaused)
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
		it.Event = new(WhitelistPaused)
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
func (it *WhitelistPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistPaused represents a Paused event raised by the Whitelist contract.
type WhitelistPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Whitelist *WhitelistFilterer) FilterPaused(opts *bind.FilterOpts) (*WhitelistPausedIterator, error) {

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WhitelistPausedIterator{contract: _Whitelist.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Whitelist *WhitelistFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WhitelistPaused) (event.Subscription, error) {

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistPaused)
				if err := _Whitelist.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Whitelist *WhitelistFilterer) ParsePaused(log types.Log) (*WhitelistPaused, error) {
	event := new(WhitelistPaused)
	if err := _Whitelist.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistPublicCapabilityUpdatedIterator is returned from FilterPublicCapabilityUpdated and is used to iterate over the raw logs and unpacked data for PublicCapabilityUpdated events raised by the Whitelist contract.
type WhitelistPublicCapabilityUpdatedIterator struct {
	Event *WhitelistPublicCapabilityUpdated // Event containing the contract specifics and raw log

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
func (it *WhitelistPublicCapabilityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistPublicCapabilityUpdated)
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
		it.Event = new(WhitelistPublicCapabilityUpdated)
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
func (it *WhitelistPublicCapabilityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistPublicCapabilityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistPublicCapabilityUpdated represents a PublicCapabilityUpdated event raised by the Whitelist contract.
type WhitelistPublicCapabilityUpdated struct {
	Target      common.Address
	FunctionSig [4]byte
	Enabled     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPublicCapabilityUpdated is a free log retrieval operation binding the contract event 0x950a343f5d10445e82a71036d3f4fb3016180a25805141932543b83e2078a93e.
//
// Solidity: event PublicCapabilityUpdated(address indexed target, bytes4 indexed functionSig, bool enabled)
func (_Whitelist *WhitelistFilterer) FilterPublicCapabilityUpdated(opts *bind.FilterOpts, target []common.Address, functionSig [][4]byte) (*WhitelistPublicCapabilityUpdatedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var functionSigRule []interface{}
	for _, functionSigItem := range functionSig {
		functionSigRule = append(functionSigRule, functionSigItem)
	}

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "PublicCapabilityUpdated", targetRule, functionSigRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistPublicCapabilityUpdatedIterator{contract: _Whitelist.contract, event: "PublicCapabilityUpdated", logs: logs, sub: sub}, nil
}

// WatchPublicCapabilityUpdated is a free log subscription operation binding the contract event 0x950a343f5d10445e82a71036d3f4fb3016180a25805141932543b83e2078a93e.
//
// Solidity: event PublicCapabilityUpdated(address indexed target, bytes4 indexed functionSig, bool enabled)
func (_Whitelist *WhitelistFilterer) WatchPublicCapabilityUpdated(opts *bind.WatchOpts, sink chan<- *WhitelistPublicCapabilityUpdated, target []common.Address, functionSig [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var functionSigRule []interface{}
	for _, functionSigItem := range functionSig {
		functionSigRule = append(functionSigRule, functionSigItem)
	}

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "PublicCapabilityUpdated", targetRule, functionSigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistPublicCapabilityUpdated)
				if err := _Whitelist.contract.UnpackLog(event, "PublicCapabilityUpdated", log); err != nil {
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

// ParsePublicCapabilityUpdated is a log parse operation binding the contract event 0x950a343f5d10445e82a71036d3f4fb3016180a25805141932543b83e2078a93e.
//
// Solidity: event PublicCapabilityUpdated(address indexed target, bytes4 indexed functionSig, bool enabled)
func (_Whitelist *WhitelistFilterer) ParsePublicCapabilityUpdated(log types.Log) (*WhitelistPublicCapabilityUpdated, error) {
	event := new(WhitelistPublicCapabilityUpdated)
	if err := _Whitelist.contract.UnpackLog(event, "PublicCapabilityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistRoleCapabilityUpdatedIterator is returned from FilterRoleCapabilityUpdated and is used to iterate over the raw logs and unpacked data for RoleCapabilityUpdated events raised by the Whitelist contract.
type WhitelistRoleCapabilityUpdatedIterator struct {
	Event *WhitelistRoleCapabilityUpdated // Event containing the contract specifics and raw log

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
func (it *WhitelistRoleCapabilityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistRoleCapabilityUpdated)
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
		it.Event = new(WhitelistRoleCapabilityUpdated)
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
func (it *WhitelistRoleCapabilityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistRoleCapabilityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistRoleCapabilityUpdated represents a RoleCapabilityUpdated event raised by the Whitelist contract.
type WhitelistRoleCapabilityUpdated struct {
	Role        uint8
	Target      common.Address
	FunctionSig [4]byte
	Enabled     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRoleCapabilityUpdated is a free log retrieval operation binding the contract event 0xa52ea92e6e955aa8ac66420b86350f7139959adfcc7e6a14eee1bd116d09860e.
//
// Solidity: event RoleCapabilityUpdated(uint8 indexed role, address indexed target, bytes4 indexed functionSig, bool enabled)
func (_Whitelist *WhitelistFilterer) FilterRoleCapabilityUpdated(opts *bind.FilterOpts, role []uint8, target []common.Address, functionSig [][4]byte) (*WhitelistRoleCapabilityUpdatedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var functionSigRule []interface{}
	for _, functionSigItem := range functionSig {
		functionSigRule = append(functionSigRule, functionSigItem)
	}

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "RoleCapabilityUpdated", roleRule, targetRule, functionSigRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistRoleCapabilityUpdatedIterator{contract: _Whitelist.contract, event: "RoleCapabilityUpdated", logs: logs, sub: sub}, nil
}

// WatchRoleCapabilityUpdated is a free log subscription operation binding the contract event 0xa52ea92e6e955aa8ac66420b86350f7139959adfcc7e6a14eee1bd116d09860e.
//
// Solidity: event RoleCapabilityUpdated(uint8 indexed role, address indexed target, bytes4 indexed functionSig, bool enabled)
func (_Whitelist *WhitelistFilterer) WatchRoleCapabilityUpdated(opts *bind.WatchOpts, sink chan<- *WhitelistRoleCapabilityUpdated, role []uint8, target []common.Address, functionSig [][4]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var functionSigRule []interface{}
	for _, functionSigItem := range functionSig {
		functionSigRule = append(functionSigRule, functionSigItem)
	}

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "RoleCapabilityUpdated", roleRule, targetRule, functionSigRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistRoleCapabilityUpdated)
				if err := _Whitelist.contract.UnpackLog(event, "RoleCapabilityUpdated", log); err != nil {
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

// ParseRoleCapabilityUpdated is a log parse operation binding the contract event 0xa52ea92e6e955aa8ac66420b86350f7139959adfcc7e6a14eee1bd116d09860e.
//
// Solidity: event RoleCapabilityUpdated(uint8 indexed role, address indexed target, bytes4 indexed functionSig, bool enabled)
func (_Whitelist *WhitelistFilterer) ParseRoleCapabilityUpdated(log types.Log) (*WhitelistRoleCapabilityUpdated, error) {
	event := new(WhitelistRoleCapabilityUpdated)
	if err := _Whitelist.contract.UnpackLog(event, "RoleCapabilityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Whitelist contract.
type WhitelistUnpausedIterator struct {
	Event *WhitelistUnpaused // Event containing the contract specifics and raw log

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
func (it *WhitelistUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistUnpaused)
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
		it.Event = new(WhitelistUnpaused)
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
func (it *WhitelistUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistUnpaused represents a Unpaused event raised by the Whitelist contract.
type WhitelistUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Whitelist *WhitelistFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WhitelistUnpausedIterator, error) {

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WhitelistUnpausedIterator{contract: _Whitelist.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Whitelist *WhitelistFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WhitelistUnpaused) (event.Subscription, error) {

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistUnpaused)
				if err := _Whitelist.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Whitelist *WhitelistFilterer) ParseUnpaused(log types.Log) (*WhitelistUnpaused, error) {
	event := new(WhitelistUnpaused)
	if err := _Whitelist.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Whitelist contract.
type WhitelistUpgradedIterator struct {
	Event *WhitelistUpgraded // Event containing the contract specifics and raw log

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
func (it *WhitelistUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistUpgraded)
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
		it.Event = new(WhitelistUpgraded)
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
func (it *WhitelistUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistUpgraded represents a Upgraded event raised by the Whitelist contract.
type WhitelistUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Whitelist *WhitelistFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WhitelistUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistUpgradedIterator{contract: _Whitelist.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Whitelist *WhitelistFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WhitelistUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistUpgraded)
				if err := _Whitelist.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Whitelist *WhitelistFilterer) ParseUpgraded(log types.Log) (*WhitelistUpgraded, error) {
	event := new(WhitelistUpgraded)
	if err := _Whitelist.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistUserRoleUpdatedIterator is returned from FilterUserRoleUpdated and is used to iterate over the raw logs and unpacked data for UserRoleUpdated events raised by the Whitelist contract.
type WhitelistUserRoleUpdatedIterator struct {
	Event *WhitelistUserRoleUpdated // Event containing the contract specifics and raw log

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
func (it *WhitelistUserRoleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistUserRoleUpdated)
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
		it.Event = new(WhitelistUserRoleUpdated)
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
func (it *WhitelistUserRoleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistUserRoleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistUserRoleUpdated represents a UserRoleUpdated event raised by the Whitelist contract.
type WhitelistUserRoleUpdated struct {
	User    common.Address
	Role    uint8
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUserRoleUpdated is a free log retrieval operation binding the contract event 0x4c9bdd0c8e073eb5eda2250b18d8e5121ff27b62064fbeeeed4869bb99bc5bf2.
//
// Solidity: event UserRoleUpdated(address indexed user, uint8 indexed role, bool enabled)
func (_Whitelist *WhitelistFilterer) FilterUserRoleUpdated(opts *bind.FilterOpts, user []common.Address, role []uint8) (*WhitelistUserRoleUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "UserRoleUpdated", userRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistUserRoleUpdatedIterator{contract: _Whitelist.contract, event: "UserRoleUpdated", logs: logs, sub: sub}, nil
}

// WatchUserRoleUpdated is a free log subscription operation binding the contract event 0x4c9bdd0c8e073eb5eda2250b18d8e5121ff27b62064fbeeeed4869bb99bc5bf2.
//
// Solidity: event UserRoleUpdated(address indexed user, uint8 indexed role, bool enabled)
func (_Whitelist *WhitelistFilterer) WatchUserRoleUpdated(opts *bind.WatchOpts, sink chan<- *WhitelistUserRoleUpdated, user []common.Address, role []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "UserRoleUpdated", userRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistUserRoleUpdated)
				if err := _Whitelist.contract.UnpackLog(event, "UserRoleUpdated", log); err != nil {
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

// ParseUserRoleUpdated is a log parse operation binding the contract event 0x4c9bdd0c8e073eb5eda2250b18d8e5121ff27b62064fbeeeed4869bb99bc5bf2.
//
// Solidity: event UserRoleUpdated(address indexed user, uint8 indexed role, bool enabled)
func (_Whitelist *WhitelistFilterer) ParseUserRoleUpdated(log types.Log) (*WhitelistUserRoleUpdated, error) {
	event := new(WhitelistUserRoleUpdated)
	if err := _Whitelist.contract.UnpackLog(event, "UserRoleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
