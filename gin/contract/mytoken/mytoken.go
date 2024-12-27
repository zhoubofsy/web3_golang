// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mytoken

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

// MytokenMetaData contains all meta data concerning the Mytoken contract.
var MytokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f5f553480156011575f5ffd5b503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061127c8061005f5f395ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c80636a6278421161006f5780636a6278421461016557806370a082311461019557806395d89b41146101c5578063a36298c7146101e3578063a9059cbb14610201578063dd62ed3e14610231576100a7565b806306fdde03146100ab578063095ea7b3146100c957806318160ddd146100f957806323b872dd14610117578063313ce56714610147575b5f5ffd5b6100b3610261565b6040516100c09190610bba565b60405180910390f35b6100e360048036038101906100de9190610c6b565b61029a565b6040516100f09190610cc3565b60405180910390f35b610101610387565b60405161010e9190610ceb565b60405180910390f35b610131600480360381019061012c9190610d04565b61038c565b60405161013e9190610cc3565b60405180910390f35b61014f6106e0565b60405161015c9190610d6f565b60405180910390f35b61017f600480360381019061017a9190610d88565b6106e5565b60405161018c9190610ceb565b60405180910390f35b6101af60048036038101906101aa9190610d88565b6108b2565b6040516101bc9190610ceb565b60405180910390f35b6101cd6108c7565b6040516101da9190610bba565b60405180910390f35b6101eb610900565b6040516101f89190610ceb565b60405180910390f35b61021b60048036038101906102169190610c6b565b610920565b6040516102289190610cc3565b60405180910390f35b61024b60048036038101906102469190610db3565b610b2a565b6040516102589190610ceb565b60405180910390f35b6040518060400160405280600b81526020017f426f62277320546f6b656e00000000000000000000000000000000000000000081525081565b5f8160035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516103759190610ceb565b60405180910390a36001905092915050565b5f5481565b5f825f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036103fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f390610e3b565b60405180910390fd5b84838060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101561047e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047590610ea3565b60405180910390fd5b8460035f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610539576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053090610f0b565b60405180910390fd5b8460035f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105c09190610f56565b925050819055508460025f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106139190610f56565b925050819055508460025f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106669190610f89565b925050819055508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef876040516106ca9190610ceb565b60405180910390a3600193505050509392505050565b601281565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610775576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076c90611006565b60405180910390fd5b5f6012600a6107849190611153565b6103e861ffff16610795919061119d565b9050805f546107a49190610f89565b6012600a6107b29190611153565b63019bfcc06107c1919061119d565b1015610802576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f990611228565b60405180910390fd5b805f5f8282546108129190610f89565b925050819055508060025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546108659190610f89565b9250508190555060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054915050919050565b6002602052805f5260405f205f915090505481565b6040518060400160405280600381526020017f42424b000000000000000000000000000000000000000000000000000000000081525081565b6012600a61090e9190611153565b63019bfcc061091d919061119d565b81565b5f825f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610990576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098790610e3b565b60405180910390fd5b33838060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610a12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0990610ea3565b60405180910390fd5b8460025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610a5e9190610f56565b925050819055508460025f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ab19190610f89565b925050819055508573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef87604051610b159190610ceb565b60405180910390a36001935050505092915050565b6003602052815f5260405f20602052805f5260405f205f91509150505481565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610b8c82610b4a565b610b968185610b54565b9350610ba6818560208601610b64565b610baf81610b72565b840191505092915050565b5f6020820190508181035f830152610bd28184610b82565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610c0782610bde565b9050919050565b610c1781610bfd565b8114610c21575f5ffd5b50565b5f81359050610c3281610c0e565b92915050565b5f819050919050565b610c4a81610c38565b8114610c54575f5ffd5b50565b5f81359050610c6581610c41565b92915050565b5f5f60408385031215610c8157610c80610bda565b5b5f610c8e85828601610c24565b9250506020610c9f85828601610c57565b9150509250929050565b5f8115159050919050565b610cbd81610ca9565b82525050565b5f602082019050610cd65f830184610cb4565b92915050565b610ce581610c38565b82525050565b5f602082019050610cfe5f830184610cdc565b92915050565b5f5f5f60608486031215610d1b57610d1a610bda565b5b5f610d2886828701610c24565b9350506020610d3986828701610c24565b9250506040610d4a86828701610c57565b9150509250925092565b5f60ff82169050919050565b610d6981610d54565b82525050565b5f602082019050610d825f830184610d60565b92915050565b5f60208284031215610d9d57610d9c610bda565b5b5f610daa84828501610c24565b91505092915050565b5f5f60408385031215610dc957610dc8610bda565b5b5f610dd685828601610c24565b9250506020610de785828601610c24565b9150509250929050565b7f496e76616c696420616464726573732e000000000000000000000000000000005f82015250565b5f610e25601083610b54565b9150610e3082610df1565b602082019050919050565b5f6020820190508181035f830152610e5281610e19565b9050919050565b7f496e73756666696369656e742062616c616e63650000000000000000000000005f82015250565b5f610e8d601483610b54565b9150610e9882610e59565b602082019050919050565b5f6020820190508181035f830152610eba81610e81565b9050919050565b7f4e6f20656e6f75676820617070726f76652076616c75652e00000000000000005f82015250565b5f610ef5601883610b54565b9150610f0082610ec1565b602082019050919050565b5f6020820190508181035f830152610f2281610ee9565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610f6082610c38565b9150610f6b83610c38565b9250828203905081811115610f8357610f82610f29565b5b92915050565b5f610f9382610c38565b9150610f9e83610c38565b9250828201905080821115610fb657610fb5610f29565b5b92915050565b7f4e6f74206f776e65722e000000000000000000000000000000000000000000005f82015250565b5f610ff0600a83610b54565b9150610ffb82610fbc565b602082019050919050565b5f6020820190508181035f83015261101d81610fe4565b9050919050565b5f8160011c9050919050565b5f5f8291508390505b60018511156110795780860481111561105557611054610f29565b5b60018516156110645780820291505b808102905061107285611024565b9450611039565b94509492505050565b5f82611091576001905061114c565b8161109e575f905061114c565b81600181146110b457600281146110be576110ed565b600191505061114c565b60ff8411156110d0576110cf610f29565b5b8360020a9150848211156110e7576110e6610f29565b5b5061114c565b5060208310610133831016604e8410600b84101617156111225782820a90508381111561111d5761111c610f29565b5b61114c565b61112f8484846001611030565b9250905081840481111561114657611145610f29565b5b81810290505b9392505050565b5f61115d82610c38565b915061116883610d54565b92506111957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611082565b905092915050565b5f6111a782610c38565b91506111b283610c38565b92508282026111c081610c38565b915082820484148315176111d7576111d6610f29565b5b5092915050565b7f4f7574206f66206c696d69742e000000000000000000000000000000000000005f82015250565b5f611212600d83610b54565b915061121d826111de565b602082019050919050565b5f6020820190508181035f83015261123f81611206565b905091905056fea2646970667358221220a2101ed1ac4183510f596d4803d6f7332ec0f76c177ab6c61b5cd57c350a643f64736f6c634300081c0033",
}

// MytokenABI is the input ABI used to generate the binding from.
// Deprecated: Use MytokenMetaData.ABI instead.
var MytokenABI = MytokenMetaData.ABI

// MytokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MytokenMetaData.Bin instead.
var MytokenBin = MytokenMetaData.Bin

// DeployMytoken deploys a new Ethereum contract, binding an instance of Mytoken to it.
func DeployMytoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mytoken, error) {
	parsed, err := MytokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MytokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mytoken{MytokenCaller: MytokenCaller{contract: contract}, MytokenTransactor: MytokenTransactor{contract: contract}, MytokenFilterer: MytokenFilterer{contract: contract}}, nil
}

// Mytoken is an auto generated Go binding around an Ethereum contract.
type Mytoken struct {
	MytokenCaller     // Read-only binding to the contract
	MytokenTransactor // Write-only binding to the contract
	MytokenFilterer   // Log filterer for contract events
}

// MytokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MytokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MytokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MytokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MytokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MytokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MytokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MytokenSession struct {
	Contract     *Mytoken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MytokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MytokenCallerSession struct {
	Contract *MytokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MytokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MytokenTransactorSession struct {
	Contract     *MytokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MytokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MytokenRaw struct {
	Contract *Mytoken // Generic contract binding to access the raw methods on
}

// MytokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MytokenCallerRaw struct {
	Contract *MytokenCaller // Generic read-only contract binding to access the raw methods on
}

// MytokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MytokenTransactorRaw struct {
	Contract *MytokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMytoken creates a new instance of Mytoken, bound to a specific deployed contract.
func NewMytoken(address common.Address, backend bind.ContractBackend) (*Mytoken, error) {
	contract, err := bindMytoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mytoken{MytokenCaller: MytokenCaller{contract: contract}, MytokenTransactor: MytokenTransactor{contract: contract}, MytokenFilterer: MytokenFilterer{contract: contract}}, nil
}

// NewMytokenCaller creates a new read-only instance of Mytoken, bound to a specific deployed contract.
func NewMytokenCaller(address common.Address, caller bind.ContractCaller) (*MytokenCaller, error) {
	contract, err := bindMytoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MytokenCaller{contract: contract}, nil
}

// NewMytokenTransactor creates a new write-only instance of Mytoken, bound to a specific deployed contract.
func NewMytokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MytokenTransactor, error) {
	contract, err := bindMytoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MytokenTransactor{contract: contract}, nil
}

// NewMytokenFilterer creates a new log filterer instance of Mytoken, bound to a specific deployed contract.
func NewMytokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MytokenFilterer, error) {
	contract, err := bindMytoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MytokenFilterer{contract: contract}, nil
}

// bindMytoken binds a generic wrapper to an already deployed contract.
func bindMytoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MytokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mytoken *MytokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mytoken.Contract.MytokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mytoken *MytokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mytoken.Contract.MytokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mytoken *MytokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mytoken.Contract.MytokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mytoken *MytokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mytoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mytoken *MytokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mytoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mytoken *MytokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mytoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Mytoken *MytokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Mytoken *MytokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Mytoken.Contract.Allowance(&_Mytoken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Mytoken *MytokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Mytoken.Contract.Allowance(&_Mytoken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Mytoken *MytokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Mytoken *MytokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Mytoken.Contract.BalanceOf(&_Mytoken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Mytoken *MytokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Mytoken.Contract.BalanceOf(&_Mytoken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mytoken *MytokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mytoken *MytokenSession) Decimals() (uint8, error) {
	return _Mytoken.Contract.Decimals(&_Mytoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mytoken *MytokenCallerSession) Decimals() (uint8, error) {
	return _Mytoken.Contract.Decimals(&_Mytoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mytoken *MytokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mytoken *MytokenSession) Name() (string, error) {
	return _Mytoken.Contract.Name(&_Mytoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mytoken *MytokenCallerSession) Name() (string, error) {
	return _Mytoken.Contract.Name(&_Mytoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mytoken *MytokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mytoken *MytokenSession) Symbol() (string, error) {
	return _Mytoken.Contract.Symbol(&_Mytoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mytoken *MytokenCallerSession) Symbol() (string, error) {
	return _Mytoken.Contract.Symbol(&_Mytoken.CallOpts)
}

// TotalLimit is a free data retrieval call binding the contract method 0xa36298c7.
//
// Solidity: function totalLimit() view returns(uint256)
func (_Mytoken *MytokenCaller) TotalLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "totalLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLimit is a free data retrieval call binding the contract method 0xa36298c7.
//
// Solidity: function totalLimit() view returns(uint256)
func (_Mytoken *MytokenSession) TotalLimit() (*big.Int, error) {
	return _Mytoken.Contract.TotalLimit(&_Mytoken.CallOpts)
}

// TotalLimit is a free data retrieval call binding the contract method 0xa36298c7.
//
// Solidity: function totalLimit() view returns(uint256)
func (_Mytoken *MytokenCallerSession) TotalLimit() (*big.Int, error) {
	return _Mytoken.Contract.TotalLimit(&_Mytoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mytoken *MytokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mytoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mytoken *MytokenSession) TotalSupply() (*big.Int, error) {
	return _Mytoken.Contract.TotalSupply(&_Mytoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mytoken *MytokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Mytoken.Contract.TotalSupply(&_Mytoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Mytoken *MytokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Mytoken *MytokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Approve(&_Mytoken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Mytoken *MytokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Approve(&_Mytoken.TransactOpts, _spender, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256)
func (_Mytoken *MytokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "mint", _to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256)
func (_Mytoken *MytokenSession) Mint(_to common.Address) (*types.Transaction, error) {
	return _Mytoken.Contract.Mint(&_Mytoken.TransactOpts, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address _to) returns(uint256)
func (_Mytoken *MytokenTransactorSession) Mint(_to common.Address) (*types.Transaction, error) {
	return _Mytoken.Contract.Mint(&_Mytoken.TransactOpts, _to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Mytoken *MytokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Mytoken *MytokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Transfer(&_Mytoken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Mytoken *MytokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.Transfer(&_Mytoken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Mytoken *MytokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Mytoken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Mytoken *MytokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.TransferFrom(&_Mytoken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Mytoken *MytokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Mytoken.Contract.TransferFrom(&_Mytoken.TransactOpts, _from, _to, _value)
}

// MytokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mytoken contract.
type MytokenApprovalIterator struct {
	Event *MytokenApproval // Event containing the contract specifics and raw log

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
func (it *MytokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MytokenApproval)
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
		it.Event = new(MytokenApproval)
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
func (it *MytokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MytokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MytokenApproval represents a Approval event raised by the Mytoken contract.
type MytokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Mytoken *MytokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*MytokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Mytoken.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &MytokenApprovalIterator{contract: _Mytoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Mytoken *MytokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MytokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Mytoken.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MytokenApproval)
				if err := _Mytoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Mytoken *MytokenFilterer) ParseApproval(log types.Log) (*MytokenApproval, error) {
	event := new(MytokenApproval)
	if err := _Mytoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MytokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mytoken contract.
type MytokenTransferIterator struct {
	Event *MytokenTransfer // Event containing the contract specifics and raw log

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
func (it *MytokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MytokenTransfer)
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
		it.Event = new(MytokenTransfer)
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
func (it *MytokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MytokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MytokenTransfer represents a Transfer event raised by the Mytoken contract.
type MytokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Mytoken *MytokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*MytokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Mytoken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &MytokenTransferIterator{contract: _Mytoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Mytoken *MytokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MytokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Mytoken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MytokenTransfer)
				if err := _Mytoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Mytoken *MytokenFilterer) ParseTransfer(log types.Log) (*MytokenTransfer, error) {
	event := new(MytokenTransfer)
	if err := _Mytoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
