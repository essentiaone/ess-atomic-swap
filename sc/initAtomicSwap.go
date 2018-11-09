// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sc

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addressOfInitiator\",\"type\":\"address\"},{\"name\":\"_password\",\"type\":\"string\"},{\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"confirmInit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"inits\",\"outputs\":[{\"name\":\"addressFrom\",\"type\":\"address\"},{\"name\":\"addressTo\",\"type\":\"address\"},{\"name\":\"isShow\",\"type\":\"bool\"},{\"name\":\"isRedeem\",\"type\":\"bool\"},{\"name\":\"isInit\",\"type\":\"bool\"},{\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"hashSecret\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addressFrom\",\"type\":\"address\"},{\"name\":\"_addressTo\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_password\",\"type\":\"string\"}],\"name\":\"addInit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"confirmedInits\",\"outputs\":[{\"name\":\"addressFrom\",\"type\":\"address\"},{\"name\":\"addressTo\",\"type\":\"address\"},{\"name\":\"isShow\",\"type\":\"bool\"},{\"name\":\"isRedeem\",\"type\":\"bool\"},{\"name\":\"isInit\",\"type\":\"bool\"},{\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"hashSecret\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addressOfInitiator\",\"type\":\"address\"}],\"name\":\"getInit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
const StoreBin = `608060405234801561001057600080fd5b50610877806100206000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631800a4438114610071578063361ad7d8146100f557806387e3a20f14610163578063a28062f1146101e4578063c9f73b5214610208575b600080fd5b34801561007d57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100e1958335600160a060020a03169536956044949193909101919081908401838280828437509497505084359550505060209092013591506102609050565b604080519115158252519081900360200190f35b34801561010157600080fd5b50610116600160a060020a0360043516610533565b60408051600160a060020a03998a168152979098166020880152941515868801529215156060860152901515608085015260a084015260c083015260e08201529051908190036101000190f35b34801561016f57600080fd5b50604080516020601f6064356004818101359283018490048402850184019095528184526101d294600160a060020a0381358116956024803590921695604435953695608494019181908401838280828437509497506105b39650505050505050565b60408051918252519081900360200190f35b3480156101f057600080fd5b50610116600160a060020a0360043516602435610784565b34801561021457600080fd5b50610229600160a060020a036004351661080d565b60408051600160a060020a039687168152949095166020850152838501929092526060830152608082015290519081900360a00190f35b600160a060020a0384811660009081526020818152604080832080546001820154600383015460029093015493519187166c0100000000000000000000000081810284880190815292909816978802603484015260488301849052606883018590528a5196978b978d97929691959491938993909260880191908401908083835b602083106103005780518252601f1990920191602091820191016102e1565b6001836020036101000a038019825116818451168082178552505050505050905001955050505050506040516020818303038152906040526040518082805190602001908083835b602083106103675780518252601f199092019160209182019101610348565b51815160209384036101000a60001901801990921691161790526040805192909401829003909120600160a060020a0387166000908152918290529290206004015490911492506103ba91505057600080fd5b600160a060020a038716600090815260208190526040902060020154879085908111156103e657600080fd5b50505050600160a060020a0394851660009081526020818152604080832080546001808552838620988652979093529220805491881673ffffffffffffffffffffffffffffffffffffffff1992831617815585830180549682018054979099169690921695909517808855815474ff000000000000000000000000000000000000000019909116740100000000000000000000000000000000000000009182900460ff908116151590920217808955825475ff000000000000000000000000000000000000000000199091167501000000000000000000000000000000000000000000918290048316151590910217808955915476ff000000000000000000000000000000000000000000001990921660b060020a928390049091161515909102179095556003808601549084015560028086015490840155600494850154949092019390935592915050565b60006020819052908152604090208054600182015460028301546003840154600490940154600160a060020a03938416949383169360ff7401000000000000000000000000000000000000000085048116947501000000000000000000000000000000000000000000810482169460b060020a9091049091169290919088565b600160a060020a038416600090815260208190526040812060010154859060b060020a900460ff16156105e557600080fd5b600160a060020a0380871660008181526020818152604091829020805473ffffffffffffffffffffffffffffffffffffffff199081168517825560018201805476ffffff000000000000000000000000000000000000000019978d16921682179690961660b060020a179095554260028201819055600390910189905591516c01000000000000000000000000938402818301908152939094026034850152604884018890526068840182905286518a948a948a94938a93919260880191908401908083835b602083106106ca5780518252601f1990920191602091820191016106ab565b6001836020036101000a038019825116818451168082178552505050505050905001955050505050506040516020818303038152906040526040518082805190602001908083835b602083106107315780518252601f199092019160209182019101610712565b51815160001960209485036101000a019081169019919091161790526040805194909201849003909320600160a060020a038c166000908152938490529220600401829055509350505050949350505050565b6001602081815260009384526040808520909152918352912080549181015460028201546003830154600490930154600160a060020a039485169483169360ff7401000000000000000000000000000000000000000085048116947501000000000000000000000000000000000000000000810482169460b060020a9091049091169290919088565b600160a060020a03908116600090815260208190526040902080546001820154600383015460028401546004909401549285169591909416939291905600a165627a7a7230582004bb0914342d3d08fad9bc3c77e2a8d7d88188ab24bd4292403fc2423c22e1dd0029`

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// ConfirmedInits is a free data retrieval call binding the contract method 0xa28062f1.
//
// Solidity: function confirmedInits( address,  bytes32) constant returns(addressFrom address, addressTo address, isShow bool, isRedeem bool, isInit bool, blockTimestamp uint256, amount uint256, hashSecret bytes32)
func (_Store *StoreCaller) ConfirmedInits(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	AddressFrom    common.Address
	AddressTo      common.Address
	IsShow         bool
	IsRedeem       bool
	IsInit         bool
	BlockTimestamp *big.Int
	Amount         *big.Int
	HashSecret     [32]byte
}, error) {
	ret := new(struct {
		AddressFrom    common.Address
		AddressTo      common.Address
		IsShow         bool
		IsRedeem       bool
		IsInit         bool
		BlockTimestamp *big.Int
		Amount         *big.Int
		HashSecret     [32]byte
	})
	out := ret
	err := _Store.contract.Call(opts, out, "confirmedInits", arg0, arg1)
	return *ret, err
}

// ConfirmedInits is a free data retrieval call binding the contract method 0xa28062f1.
//
// Solidity: function confirmedInits( address,  bytes32) constant returns(addressFrom address, addressTo address, isShow bool, isRedeem bool, isInit bool, blockTimestamp uint256, amount uint256, hashSecret bytes32)
func (_Store *StoreSession) ConfirmedInits(arg0 common.Address, arg1 [32]byte) (struct {
	AddressFrom    common.Address
	AddressTo      common.Address
	IsShow         bool
	IsRedeem       bool
	IsInit         bool
	BlockTimestamp *big.Int
	Amount         *big.Int
	HashSecret     [32]byte
}, error) {
	return _Store.Contract.ConfirmedInits(&_Store.CallOpts, arg0, arg1)
}

// ConfirmedInits is a free data retrieval call binding the contract method 0xa28062f1.
//
// Solidity: function confirmedInits( address,  bytes32) constant returns(addressFrom address, addressTo address, isShow bool, isRedeem bool, isInit bool, blockTimestamp uint256, amount uint256, hashSecret bytes32)
func (_Store *StoreCallerSession) ConfirmedInits(arg0 common.Address, arg1 [32]byte) (struct {
	AddressFrom    common.Address
	AddressTo      common.Address
	IsShow         bool
	IsRedeem       bool
	IsInit         bool
	BlockTimestamp *big.Int
	Amount         *big.Int
	HashSecret     [32]byte
}, error) {
	return _Store.Contract.ConfirmedInits(&_Store.CallOpts, arg0, arg1)
}

// GetInit is a free data retrieval call binding the contract method 0xc9f73b52.
//
// Solidity: function getInit(_addressOfInitiator address) constant returns(address, address, uint256, uint256, bytes32)
func (_Store *StoreCaller) GetInit(opts *bind.CallOpts, _addressOfInitiator common.Address) (common.Address, common.Address, *big.Int, *big.Int, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Store.contract.Call(opts, out, "getInit", _addressOfInitiator)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetInit is a free data retrieval call binding the contract method 0xc9f73b52.
//
// Solidity: function getInit(_addressOfInitiator address) constant returns(address, address, uint256, uint256, bytes32)
func (_Store *StoreSession) GetInit(_addressOfInitiator common.Address) (common.Address, common.Address, *big.Int, *big.Int, [32]byte, error) {
	return _Store.Contract.GetInit(&_Store.CallOpts, _addressOfInitiator)
}

// GetInit is a free data retrieval call binding the contract method 0xc9f73b52.
//
// Solidity: function getInit(_addressOfInitiator address) constant returns(address, address, uint256, uint256, bytes32)
func (_Store *StoreCallerSession) GetInit(_addressOfInitiator common.Address) (common.Address, common.Address, *big.Int, *big.Int, [32]byte, error) {
	return _Store.Contract.GetInit(&_Store.CallOpts, _addressOfInitiator)
}

// Inits is a free data retrieval call binding the contract method 0x361ad7d8.
//
// Solidity: function inits( address) constant returns(addressFrom address, addressTo address, isShow bool, isRedeem bool, isInit bool, blockTimestamp uint256, amount uint256, hashSecret bytes32)
func (_Store *StoreCaller) Inits(opts *bind.CallOpts, arg0 common.Address) (struct {
	AddressFrom    common.Address
	AddressTo      common.Address
	IsShow         bool
	IsRedeem       bool
	IsInit         bool
	BlockTimestamp *big.Int
	Amount         *big.Int
	HashSecret     [32]byte
}, error) {
	ret := new(struct {
		AddressFrom    common.Address
		AddressTo      common.Address
		IsShow         bool
		IsRedeem       bool
		IsInit         bool
		BlockTimestamp *big.Int
		Amount         *big.Int
		HashSecret     [32]byte
	})
	out := ret
	err := _Store.contract.Call(opts, out, "inits", arg0)
	return *ret, err
}

// Inits is a free data retrieval call binding the contract method 0x361ad7d8.
//
// Solidity: function inits( address) constant returns(addressFrom address, addressTo address, isShow bool, isRedeem bool, isInit bool, blockTimestamp uint256, amount uint256, hashSecret bytes32)
func (_Store *StoreSession) Inits(arg0 common.Address) (struct {
	AddressFrom    common.Address
	AddressTo      common.Address
	IsShow         bool
	IsRedeem       bool
	IsInit         bool
	BlockTimestamp *big.Int
	Amount         *big.Int
	HashSecret     [32]byte
}, error) {
	return _Store.Contract.Inits(&_Store.CallOpts, arg0)
}

// Inits is a free data retrieval call binding the contract method 0x361ad7d8.
//
// Solidity: function inits( address) constant returns(addressFrom address, addressTo address, isShow bool, isRedeem bool, isInit bool, blockTimestamp uint256, amount uint256, hashSecret bytes32)
func (_Store *StoreCallerSession) Inits(arg0 common.Address) (struct {
	AddressFrom    common.Address
	AddressTo      common.Address
	IsShow         bool
	IsRedeem       bool
	IsInit         bool
	BlockTimestamp *big.Int
	Amount         *big.Int
	HashSecret     [32]byte
}, error) {
	return _Store.Contract.Inits(&_Store.CallOpts, arg0)
}

// AddInit is a paid mutator transaction binding the contract method 0x87e3a20f.
//
// Solidity: function addInit(_addressFrom address, _addressTo address, _amount uint256, _password string) returns(bytes32)
func (_Store *StoreTransactor) AddInit(opts *bind.TransactOpts, _addressFrom common.Address, _addressTo common.Address, _amount *big.Int, _password string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addInit", _addressFrom, _addressTo, _amount, _password)
}

// AddInit is a paid mutator transaction binding the contract method 0x87e3a20f.
//
// Solidity: function addInit(_addressFrom address, _addressTo address, _amount uint256, _password string) returns(bytes32)
func (_Store *StoreSession) AddInit(_addressFrom common.Address, _addressTo common.Address, _amount *big.Int, _password string) (*types.Transaction, error) {
	return _Store.Contract.AddInit(&_Store.TransactOpts, _addressFrom, _addressTo, _amount, _password)
}

// AddInit is a paid mutator transaction binding the contract method 0x87e3a20f.
//
// Solidity: function addInit(_addressFrom address, _addressTo address, _amount uint256, _password string) returns(bytes32)
func (_Store *StoreTransactorSession) AddInit(_addressFrom common.Address, _addressTo common.Address, _amount *big.Int, _password string) (*types.Transaction, error) {
	return _Store.Contract.AddInit(&_Store.TransactOpts, _addressFrom, _addressTo, _amount, _password)
}

// ConfirmInit is a paid mutator transaction binding the contract method 0x1800a443.
//
// Solidity: function confirmInit(_addressOfInitiator address, _password string, _txHash bytes32, _blockTimestamp uint256) returns(bool)
func (_Store *StoreTransactor) ConfirmInit(opts *bind.TransactOpts, _addressOfInitiator common.Address, _password string, _txHash [32]byte, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "confirmInit", _addressOfInitiator, _password, _txHash, _blockTimestamp)
}

// ConfirmInit is a paid mutator transaction binding the contract method 0x1800a443.
//
// Solidity: function confirmInit(_addressOfInitiator address, _password string, _txHash bytes32, _blockTimestamp uint256) returns(bool)
func (_Store *StoreSession) ConfirmInit(_addressOfInitiator common.Address, _password string, _txHash [32]byte, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ConfirmInit(&_Store.TransactOpts, _addressOfInitiator, _password, _txHash, _blockTimestamp)
}

// ConfirmInit is a paid mutator transaction binding the contract method 0x1800a443.
//
// Solidity: function confirmInit(_addressOfInitiator address, _password string, _txHash bytes32, _blockTimestamp uint256) returns(bool)
func (_Store *StoreTransactorSession) ConfirmInit(_addressOfInitiator common.Address, _password string, _txHash [32]byte, _blockTimestamp *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ConfirmInit(&_Store.TransactOpts, _addressOfInitiator, _password, _txHash, _blockTimestamp)
}
