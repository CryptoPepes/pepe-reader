// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cozy

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

// CozyABI is the input ABI used to generate the binding from.
const CozyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"affiliateContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pepeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"}],\"name\":\"calculateBid\",\"outputs\":[{\"name\":\"currentBid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"pepeId\",\"type\":\"uint256\"},{\"name\":\"auctionBegin\",\"type\":\"uint64\"},{\"name\":\"auctionEnd\",\"type\":\"uint64\"},{\"name\":\"beginPrice\",\"type\":\"uint256\"},{\"name\":\"endPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_DIVIDER\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"}],\"name\":\"savePepe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_pepeContract\",\"type\":\"address\"},{\"name\":\"_affiliateContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pepe\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"AuctionWon\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pepe\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pepe\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"AuctionFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"},{\"name\":\"_beginPrice\",\"type\":\"uint256\"},{\"name\":\"_endPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint64\"}],\"name\":\"startAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"},{\"name\":\"_beginPrice\",\"type\":\"uint256\"},{\"name\":\"_endPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint64\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"startAuctionDirect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"},{\"name\":\"_cozyCandidate\",\"type\":\"uint256\"},{\"name\":\"_candidateAsFather\",\"type\":\"bool\"},{\"name\":\"_pepeReceiver\",\"type\":\"address\"}],\"name\":\"buyCozy\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"},{\"name\":\"_cozyCandidate\",\"type\":\"uint256\"},{\"name\":\"_candidateAsFather\",\"type\":\"bool\"},{\"name\":\"_pepeReceiver\",\"type\":\"address\"},{\"name\":\"_affiliate\",\"type\":\"address\"}],\"name\":\"buyCozyAffiliated\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// CozyBin is the compiled bytecode used for deploying new contracts.
const CozyBin = `0x608060405261927c60055534801561001657600080fd5b506040516040806112d28339810160405280516020909101516000805433600160a060020a0319918216811783556001805483169091179055600380548216600160a060020a03958616179055600480549091169390921692909217905561124e90819061008490396000f3006080604052600436106100fb5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630117200581146101005780631887a4d0146101315780631c31f710146101465780632d4a170a1461016957806338af3eed14610193578063571a26a0146101a85780636a1db1bf1461020a578063715018a61461022257806374a25d431461023757806378d6ade4146102625780638da5cb5b14610277578063a6da467c1461028c578063a8f5c673146102b1578063d024cd02146102c9578063d3ce71df14610300578063db8d55f11461031f578063ddca3f4314610334578063f2fde38b14610349575b600080fd5b34801561010c57600080fd5b5061011561036a565b60408051600160a060020a039092168252519081900360200190f35b34801561013d57600080fd5b50610115610379565b34801561015257600080fd5b50610167600160a060020a0360043516610388565b005b34801561017557600080fd5b506101816004356103ce565b60408051918252519081900360200190f35b34801561019f57600080fd5b50610115610473565b3480156101b457600080fd5b506101c0600435610482565b60408051600160a060020a039097168752602087019590955267ffffffffffffffff93841686860152919092166060850152608084019190915260a0830152519081900360c00190f35b34801561021657600080fd5b506101676004356104d5565b34801561022e57600080fd5b506101676104ff565b34801561024357600080fd5b5061016760043560243560443567ffffffffffffffff6064351661056b565b34801561026e57600080fd5b50610181610624565b34801561028357600080fd5b5061011561062b565b6101676004356024356044351515600160a060020a036064358116906084351661063a565b3480156102bd57600080fd5b506101676004356106d8565b3480156102d557600080fd5b5061016760043560243560443567ffffffffffffffff60643516600160a060020a036084351661085a565b6101676004356024356044351515600160a060020a036064351661090e565b34801561032b57600080fd5b50610167610d96565b34801561034057600080fd5b50610181610dd3565b34801561035557600080fd5b50610167600160a060020a0360043516610dd9565b600454600160a060020a031681565b600354600160a060020a031681565b600054600160a060020a0316331461039f57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600081815260026020819052604082209081015467ffffffffffffffff8082164290810392859283928392839268010000000000000000909104161161041a5785600401549650610468565b600386015460048701546002880154919003945067ffffffffffffffff80821668010000000000000000909204160392508285850281151561045857fe5b0591508186600301540190508096505b505050505050919050565b600154600160a060020a031681565b600260208190526000918252604090912080546001820154928201546003830154600490930154600160a060020a03909216939267ffffffffffffffff8083169368010000000000000000909304169186565b600054600160a060020a031633146104ec57600080fd5b60055481106104fa57600080fd5b600555565b600054600160a060020a0316331461051657600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600354604080517f62bc63c40000000000000000000000000000000000000000000000000000000081526004810187905290514292600160a060020a0316916362bc63c49160248083019260209291908290030181600087803b1580156105d157600080fd5b505af11580156105e5573d6000803e3d6000fd5b505050506040513d60208110156105fb57600080fd5b505167ffffffffffffffff16111561061257600080fd5b61061e84848484610df9565b50505050565b620f424081565b600054600160a060020a031681565b60048054604080517f7352e4b8000000000000000000000000000000000000000000000000000000008152600160a060020a0386811694820194909452848416602482015290519290911691637352e4b89160448082019260009290919082900301818387803b1580156106ad57600080fd5b505af11580156106c1573d6000803e3d6000fd5b505050506106d18585858561090e565b5050505050565b60008181526002602081905260409091200154426801000000000000000090910467ffffffffffffffff161061070d57600080fd5b60035460008281526002602090815260408083205481517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a03918216600482015260248101879052915194169363a9059cbb93604480840194938390030190829087803b15801561078657600080fd5b505af115801561079a573d6000803e3d6000fd5b505050506040513d60208110156107b057600080fd5b505115156107bd57600080fd5b600081815260026020526040808220549051600160a060020a039091169183917f95b73f79c6d7b09d4dd9a323589aec50a424621f53a70ece1cc21aa75554b5199190a360009081526002602081905260408220805473ffffffffffffffffffffffffffffffffffffffff191681556001810183905590810180546fffffffffffffffffffffffffffffffff191690556003810182905560040155565b600354604080517f62bc63c40000000000000000000000000000000000000000000000000000000081526004810188905290514292600160a060020a0316916362bc63c49160248083019260209291908290030181600087803b1580156108c057600080fd5b505af11580156108d4573d6000803e3d6000fd5b505050506040513d60208110156108ea57600080fd5b505167ffffffffffffffff16111561090157600080fd5b6106d18585858585610ff9565b600354600090819081908190600160a060020a0316331461092e57600080fd5b60008881526002602081905260409091209081015490945068010000000000000000900467ffffffffffffffff16421061096757600080fd5b610970886103ce565b92503483111561097f57600080fd5b600554620f42409084028554604051929091049350600160a060020a03169083850380156108fc02916000818181858888f193505050501580156109c7573d6000803e3d6000fd5b5060048054604080517f873abd75000000000000000000000000000000000000000000000000000000008152600160a060020a03898116948201949094529051929091169163873abd75916024808201926020929091908290030181600087803b158015610a3457600080fd5b505af1158015610a48573d6000803e3d6000fd5b505050506040513d6020811015610a5e57600080fd5b50519050600160a060020a03811615801590610a9d5750604051600160a060020a038216906002840480156108fc02916000818181858888f193505050505b508515610b5b576003546001850154604080517fc29338cf0000000000000000000000000000000000000000000000000000000081526004810192909252602482018a9052600160a060020a038881166044840152905192169163c29338cf916064808201926020929091908290030181600087803b158015610b1f57600080fd5b505af1158015610b33573d6000803e3d6000fd5b505050506040513d6020811015610b4957600080fd5b50511515610b5657600080fd5b610c0d565b6003546001850154604080517fc29338cf000000000000000000000000000000000000000000000000000000008152600481018b90526024810192909252600160a060020a038881166044840152905192169163c29338cf916064808201926020929091908290030181600087803b158015610bd657600080fd5b505af1158015610bea573d6000803e3d6000fd5b505050506040513d6020811015610c0057600080fd5b50511515610c0d57600080fd5b6003548454604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152602481018c90529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015610c7e57600080fd5b505af1158015610c92573d6000803e3d6000fd5b505050506040513d6020811015610ca857600080fd5b50511515610cb557600080fd5b82341115610cf757604051600160a060020a038616903485900380156108fc02916000818181858888f19350505050158015610cf5573d6000803e3d6000fd5b505b8354604051600160a060020a03918216918716908a907f94ffdfa810a5f08da0ec8ea0f74619814453cfc101a90504a3a8f77e0eb6198690600090a4505050600094855250506002602081905260408420805473ffffffffffffffffffffffffffffffffffffffff191681556001810185905590810180546fffffffffffffffffffffffffffffffff1916905560038101849055600401929092555050565b600154604051600160a060020a0390911690303180156108fc02916000818181858888f19350505050158015610dd0573d6000803e3d6000fd5b50565b60055481565b600054600160a060020a03163314610df057600080fd5b610dd081611170565b610e016111ed565b600354604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018890529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b158015610e7457600080fd5b505af1158015610e88573d6000803e3d6000fd5b505050506040513d6020811015610e9e57600080fd5b50511515610eab57600080fd5b6000858152600260208190526040909120015468010000000000000000900467ffffffffffffffff164211610edf57600080fd5b338152602081018590524267ffffffffffffffff81811660408401819052918401166060830181905211610f1257600080fd5b6080810184815260a0820184815260008781526002602081815260408084208751815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178155918701516001830155808701519282018054606089015167ffffffffffffffff1990911667ffffffffffffffff958616176fffffffffffffffff000000000000000019166801000000000000000095909116949094029390931790925593516003850155915160049093019290925551339187917f16da476d7265fc95576888b93de4fa4849d6cea1228235887f569c6530ddfec19190a35050505050565b6110016111ed565b600354600160a060020a0316331461101857600080fd5b6000868152600260208190526040909120015468010000000000000000900467ffffffffffffffff16421161104c57600080fd5b600160a060020a0382168152602081018690524267ffffffffffffffff8181166040840181905291850116606083018190521161108857600080fd5b6080810185815260a0820185815260008881526002602081815260408084208751815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391821617825592880151600182015581880151938101805460608a015167ffffffffffffffff1990911667ffffffffffffffff968716176fffffffffffffffff000000000000000019166801000000000000000096909116959095029490941790935594516003830155925160049091015591519084169188917f16da476d7265fc95576888b93de4fa4849d6cea1228235887f569c6530ddfec19190a3505050505050565b600160a060020a038116151561118557600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152905600a165627a7a723058200ab6d823b80046fd29d168a16d30f2769706fea78bf2768755a369395c3f0a750029`

// DeployCozy deploys a new Ethereum contract, binding an instance of Cozy to it.
func DeployCozy(auth *bind.TransactOpts, backend bind.ContractBackend, _pepeContract common.Address, _affiliateContract common.Address) (common.Address, *types.Transaction, *Cozy, error) {
	parsed, err := abi.JSON(strings.NewReader(CozyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CozyBin), backend, _pepeContract, _affiliateContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cozy{CozyCaller: CozyCaller{contract: contract}, CozyTransactor: CozyTransactor{contract: contract}, CozyFilterer: CozyFilterer{contract: contract}}, nil
}

// Cozy is an auto generated Go binding around an Ethereum contract.
type Cozy struct {
	CozyCaller     // Read-only binding to the contract
	CozyTransactor // Write-only binding to the contract
	CozyFilterer   // Log filterer for contract events
}

// CozyCaller is an auto generated read-only Go binding around an Ethereum contract.
type CozyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CozyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CozyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CozyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CozyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CozySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CozySession struct {
	Contract     *Cozy             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CozyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CozyCallerSession struct {
	Contract *CozyCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CozyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CozyTransactorSession struct {
	Contract     *CozyTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CozyRaw is an auto generated low-level Go binding around an Ethereum contract.
type CozyRaw struct {
	Contract *Cozy // Generic contract binding to access the raw methods on
}

// CozyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CozyCallerRaw struct {
	Contract *CozyCaller // Generic read-only contract binding to access the raw methods on
}

// CozyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CozyTransactorRaw struct {
	Contract *CozyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCozy creates a new instance of Cozy, bound to a specific deployed contract.
func NewCozy(address common.Address, backend bind.ContractBackend) (*Cozy, error) {
	contract, err := bindCozy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cozy{CozyCaller: CozyCaller{contract: contract}, CozyTransactor: CozyTransactor{contract: contract}, CozyFilterer: CozyFilterer{contract: contract}}, nil
}

// NewCozyCaller creates a new read-only instance of Cozy, bound to a specific deployed contract.
func NewCozyCaller(address common.Address, caller bind.ContractCaller) (*CozyCaller, error) {
	contract, err := bindCozy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CozyCaller{contract: contract}, nil
}

// NewCozyTransactor creates a new write-only instance of Cozy, bound to a specific deployed contract.
func NewCozyTransactor(address common.Address, transactor bind.ContractTransactor) (*CozyTransactor, error) {
	contract, err := bindCozy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CozyTransactor{contract: contract}, nil
}

// NewCozyFilterer creates a new log filterer instance of Cozy, bound to a specific deployed contract.
func NewCozyFilterer(address common.Address, filterer bind.ContractFilterer) (*CozyFilterer, error) {
	contract, err := bindCozy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CozyFilterer{contract: contract}, nil
}

// bindCozy binds a generic wrapper to an already deployed contract.
func bindCozy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CozyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cozy *CozyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cozy.Contract.CozyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cozy *CozyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cozy.Contract.CozyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cozy *CozyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cozy.Contract.CozyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cozy *CozyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cozy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cozy *CozyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cozy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cozy *CozyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cozy.Contract.contract.Transact(opts, method, params...)
}

// FEEDIVIDER is a free data retrieval call binding the contract method 0x78d6ade4.
//
// Solidity: function FEE_DIVIDER() constant returns(uint256)
func (_Cozy *CozyCaller) FEEDIVIDER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cozy.contract.Call(opts, out, "FEE_DIVIDER")
	return *ret0, err
}

// FEEDIVIDER is a free data retrieval call binding the contract method 0x78d6ade4.
//
// Solidity: function FEE_DIVIDER() constant returns(uint256)
func (_Cozy *CozySession) FEEDIVIDER() (*big.Int, error) {
	return _Cozy.Contract.FEEDIVIDER(&_Cozy.CallOpts)
}

// FEEDIVIDER is a free data retrieval call binding the contract method 0x78d6ade4.
//
// Solidity: function FEE_DIVIDER() constant returns(uint256)
func (_Cozy *CozyCallerSession) FEEDIVIDER() (*big.Int, error) {
	return _Cozy.Contract.FEEDIVIDER(&_Cozy.CallOpts)
}

// AffiliateContract is a free data retrieval call binding the contract method 0x01172005.
//
// Solidity: function affiliateContract() constant returns(address)
func (_Cozy *CozyCaller) AffiliateContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cozy.contract.Call(opts, out, "affiliateContract")
	return *ret0, err
}

// AffiliateContract is a free data retrieval call binding the contract method 0x01172005.
//
// Solidity: function affiliateContract() constant returns(address)
func (_Cozy *CozySession) AffiliateContract() (common.Address, error) {
	return _Cozy.Contract.AffiliateContract(&_Cozy.CallOpts)
}

// AffiliateContract is a free data retrieval call binding the contract method 0x01172005.
//
// Solidity: function affiliateContract() constant returns(address)
func (_Cozy *CozyCallerSession) AffiliateContract() (common.Address, error) {
	return _Cozy.Contract.AffiliateContract(&_Cozy.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions( uint256) constant returns(seller address, pepeId uint256, auctionBegin uint64, auctionEnd uint64, beginPrice uint256, endPrice uint256)
func (_Cozy *CozyCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller       common.Address
	PepeId       *big.Int
	AuctionBegin uint64
	AuctionEnd   uint64
	BeginPrice   *big.Int
	EndPrice     *big.Int
}, error) {
	ret := new(struct {
		Seller       common.Address
		PepeId       *big.Int
		AuctionBegin uint64
		AuctionEnd   uint64
		BeginPrice   *big.Int
		EndPrice     *big.Int
	})
	out := ret
	err := _Cozy.contract.Call(opts, out, "auctions", arg0)
	return *ret, err
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions( uint256) constant returns(seller address, pepeId uint256, auctionBegin uint64, auctionEnd uint64, beginPrice uint256, endPrice uint256)
func (_Cozy *CozySession) Auctions(arg0 *big.Int) (struct {
	Seller       common.Address
	PepeId       *big.Int
	AuctionBegin uint64
	AuctionEnd   uint64
	BeginPrice   *big.Int
	EndPrice     *big.Int
}, error) {
	return _Cozy.Contract.Auctions(&_Cozy.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions( uint256) constant returns(seller address, pepeId uint256, auctionBegin uint64, auctionEnd uint64, beginPrice uint256, endPrice uint256)
func (_Cozy *CozyCallerSession) Auctions(arg0 *big.Int) (struct {
	Seller       common.Address
	PepeId       *big.Int
	AuctionBegin uint64
	AuctionEnd   uint64
	BeginPrice   *big.Int
	EndPrice     *big.Int
}, error) {
	return _Cozy.Contract.Auctions(&_Cozy.CallOpts, arg0)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_Cozy *CozyCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cozy.contract.Call(opts, out, "beneficiary")
	return *ret0, err
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_Cozy *CozySession) Beneficiary() (common.Address, error) {
	return _Cozy.Contract.Beneficiary(&_Cozy.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_Cozy *CozyCallerSession) Beneficiary() (common.Address, error) {
	return _Cozy.Contract.Beneficiary(&_Cozy.CallOpts)
}

// CalculateBid is a free data retrieval call binding the contract method 0x2d4a170a.
//
// Solidity: function calculateBid(_pepeId uint256) constant returns(currentBid uint256)
func (_Cozy *CozyCaller) CalculateBid(opts *bind.CallOpts, _pepeId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cozy.contract.Call(opts, out, "calculateBid", _pepeId)
	return *ret0, err
}

// CalculateBid is a free data retrieval call binding the contract method 0x2d4a170a.
//
// Solidity: function calculateBid(_pepeId uint256) constant returns(currentBid uint256)
func (_Cozy *CozySession) CalculateBid(_pepeId *big.Int) (*big.Int, error) {
	return _Cozy.Contract.CalculateBid(&_Cozy.CallOpts, _pepeId)
}

// CalculateBid is a free data retrieval call binding the contract method 0x2d4a170a.
//
// Solidity: function calculateBid(_pepeId uint256) constant returns(currentBid uint256)
func (_Cozy *CozyCallerSession) CalculateBid(_pepeId *big.Int) (*big.Int, error) {
	return _Cozy.Contract.CalculateBid(&_Cozy.CallOpts, _pepeId)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Cozy *CozyCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cozy.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Cozy *CozySession) Fee() (*big.Int, error) {
	return _Cozy.Contract.Fee(&_Cozy.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Cozy *CozyCallerSession) Fee() (*big.Int, error) {
	return _Cozy.Contract.Fee(&_Cozy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Cozy *CozyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cozy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Cozy *CozySession) Owner() (common.Address, error) {
	return _Cozy.Contract.Owner(&_Cozy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Cozy *CozyCallerSession) Owner() (common.Address, error) {
	return _Cozy.Contract.Owner(&_Cozy.CallOpts)
}

// PepeContract is a free data retrieval call binding the contract method 0x1887a4d0.
//
// Solidity: function pepeContract() constant returns(address)
func (_Cozy *CozyCaller) PepeContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cozy.contract.Call(opts, out, "pepeContract")
	return *ret0, err
}

// PepeContract is a free data retrieval call binding the contract method 0x1887a4d0.
//
// Solidity: function pepeContract() constant returns(address)
func (_Cozy *CozySession) PepeContract() (common.Address, error) {
	return _Cozy.Contract.PepeContract(&_Cozy.CallOpts)
}

// PepeContract is a free data retrieval call binding the contract method 0x1887a4d0.
//
// Solidity: function pepeContract() constant returns(address)
func (_Cozy *CozyCallerSession) PepeContract() (common.Address, error) {
	return _Cozy.Contract.PepeContract(&_Cozy.CallOpts)
}

// BuyCozy is a paid mutator transaction binding the contract method 0xd3ce71df.
//
// Solidity: function buyCozy(_pepeId uint256, _cozyCandidate uint256, _candidateAsFather bool, _pepeReceiver address) returns()
func (_Cozy *CozyTransactor) BuyCozy(opts *bind.TransactOpts, _pepeId *big.Int, _cozyCandidate *big.Int, _candidateAsFather bool, _pepeReceiver common.Address) (*types.Transaction, error) {
	return _Cozy.contract.Transact(opts, "buyCozy", _pepeId, _cozyCandidate, _candidateAsFather, _pepeReceiver)
}

// BuyCozy is a paid mutator transaction binding the contract method 0xd3ce71df.
//
// Solidity: function buyCozy(_pepeId uint256, _cozyCandidate uint256, _candidateAsFather bool, _pepeReceiver address) returns()
func (_Cozy *CozySession) BuyCozy(_pepeId *big.Int, _cozyCandidate *big.Int, _candidateAsFather bool, _pepeReceiver common.Address) (*types.Transaction, error) {
	return _Cozy.Contract.BuyCozy(&_Cozy.TransactOpts, _pepeId, _cozyCandidate, _candidateAsFather, _pepeReceiver)
}

// BuyCozy is a paid mutator transaction binding the contract method 0xd3ce71df.
//
// Solidity: function buyCozy(_pepeId uint256, _cozyCandidate uint256, _candidateAsFather bool, _pepeReceiver address) returns()
func (_Cozy *CozyTransactorSession) BuyCozy(_pepeId *big.Int, _cozyCandidate *big.Int, _candidateAsFather bool, _pepeReceiver common.Address) (*types.Transaction, error) {
	return _Cozy.Contract.BuyCozy(&_Cozy.TransactOpts, _pepeId, _cozyCandidate, _candidateAsFather, _pepeReceiver)
}

// BuyCozyAffiliated is a paid mutator transaction binding the contract method 0xa6da467c.
//
// Solidity: function buyCozyAffiliated(_pepeId uint256, _cozyCandidate uint256, _candidateAsFather bool, _pepeReceiver address, _affiliate address) returns()
func (_Cozy *CozyTransactor) BuyCozyAffiliated(opts *bind.TransactOpts, _pepeId *big.Int, _cozyCandidate *big.Int, _candidateAsFather bool, _pepeReceiver common.Address, _affiliate common.Address) (*types.Transaction, error) {
	return _Cozy.contract.Transact(opts, "buyCozyAffiliated", _pepeId, _cozyCandidate, _candidateAsFather, _pepeReceiver, _affiliate)
}

// BuyCozyAffiliated is a paid mutator transaction binding the contract method 0xa6da467c.
//
// Solidity: function buyCozyAffiliated(_pepeId uint256, _cozyCandidate uint256, _candidateAsFather bool, _pepeReceiver address, _affiliate address) returns()
func (_Cozy *CozySession) BuyCozyAffiliated(_pepeId *big.Int, _cozyCandidate *big.Int, _candidateAsFather bool, _pepeReceiver common.Address, _affiliate common.Address) (*types.Transaction, error) {
	return _Cozy.Contract.BuyCozyAffiliated(&_Cozy.TransactOpts, _pepeId, _cozyCandidate, _candidateAsFather, _pepeReceiver, _affiliate)
}

// BuyCozyAffiliated is a paid mutator transaction binding the contract method 0xa6da467c.
//
// Solidity: function buyCozyAffiliated(_pepeId uint256, _cozyCandidate uint256, _candidateAsFather bool, _pepeReceiver address, _affiliate address) returns()
func (_Cozy *CozyTransactorSession) BuyCozyAffiliated(_pepeId *big.Int, _cozyCandidate *big.Int, _candidateAsFather bool, _pepeReceiver common.Address, _affiliate common.Address) (*types.Transaction, error) {
	return _Cozy.Contract.BuyCozyAffiliated(&_Cozy.TransactOpts, _pepeId, _cozyCandidate, _candidateAsFather, _pepeReceiver, _affiliate)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(_fee uint256) returns()
func (_Cozy *CozyTransactor) ChangeFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Cozy.contract.Transact(opts, "changeFee", _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(_fee uint256) returns()
func (_Cozy *CozySession) ChangeFee(_fee *big.Int) (*types.Transaction, error) {
	return _Cozy.Contract.ChangeFee(&_Cozy.TransactOpts, _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(_fee uint256) returns()
func (_Cozy *CozyTransactorSession) ChangeFee(_fee *big.Int) (*types.Transaction, error) {
	return _Cozy.Contract.ChangeFee(&_Cozy.TransactOpts, _fee)
}

// GetFees is a paid mutator transaction binding the contract method 0xdb8d55f1.
//
// Solidity: function getFees() returns()
func (_Cozy *CozyTransactor) GetFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cozy.contract.Transact(opts, "getFees")
}

// GetFees is a paid mutator transaction binding the contract method 0xdb8d55f1.
//
// Solidity: function getFees() returns()
func (_Cozy *CozySession) GetFees() (*types.Transaction, error) {
	return _Cozy.Contract.GetFees(&_Cozy.TransactOpts)
}

// GetFees is a paid mutator transaction binding the contract method 0xdb8d55f1.
//
// Solidity: function getFees() returns()
func (_Cozy *CozyTransactorSession) GetFees() (*types.Transaction, error) {
	return _Cozy.Contract.GetFees(&_Cozy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cozy *CozyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cozy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cozy *CozySession) RenounceOwnership() (*types.Transaction, error) {
	return _Cozy.Contract.RenounceOwnership(&_Cozy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cozy *CozyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cozy.Contract.RenounceOwnership(&_Cozy.TransactOpts)
}

// SavePepe is a paid mutator transaction binding the contract method 0xa8f5c673.
//
// Solidity: function savePepe(_pepeId uint256) returns()
func (_Cozy *CozyTransactor) SavePepe(opts *bind.TransactOpts, _pepeId *big.Int) (*types.Transaction, error) {
	return _Cozy.contract.Transact(opts, "savePepe", _pepeId)
}

// SavePepe is a paid mutator transaction binding the contract method 0xa8f5c673.
//
// Solidity: function savePepe(_pepeId uint256) returns()
func (_Cozy *CozySession) SavePepe(_pepeId *big.Int) (*types.Transaction, error) {
	return _Cozy.Contract.SavePepe(&_Cozy.TransactOpts, _pepeId)
}

// SavePepe is a paid mutator transaction binding the contract method 0xa8f5c673.
//
// Solidity: function savePepe(_pepeId uint256) returns()
func (_Cozy *CozyTransactorSession) SavePepe(_pepeId *big.Int) (*types.Transaction, error) {
	return _Cozy.Contract.SavePepe(&_Cozy.TransactOpts, _pepeId)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(_beneficiary address) returns()
func (_Cozy *CozyTransactor) SetBeneficiary(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _Cozy.contract.Transact(opts, "setBeneficiary", _beneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(_beneficiary address) returns()
func (_Cozy *CozySession) SetBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _Cozy.Contract.SetBeneficiary(&_Cozy.TransactOpts, _beneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(_beneficiary address) returns()
func (_Cozy *CozyTransactorSession) SetBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _Cozy.Contract.SetBeneficiary(&_Cozy.TransactOpts, _beneficiary)
}

// StartAuction is a paid mutator transaction binding the contract method 0x74a25d43.
//
// Solidity: function startAuction(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64) returns()
func (_Cozy *CozyTransactor) StartAuction(opts *bind.TransactOpts, _pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64) (*types.Transaction, error) {
	return _Cozy.contract.Transact(opts, "startAuction", _pepeId, _beginPrice, _endPrice, _duration)
}

// StartAuction is a paid mutator transaction binding the contract method 0x74a25d43.
//
// Solidity: function startAuction(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64) returns()
func (_Cozy *CozySession) StartAuction(_pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64) (*types.Transaction, error) {
	return _Cozy.Contract.StartAuction(&_Cozy.TransactOpts, _pepeId, _beginPrice, _endPrice, _duration)
}

// StartAuction is a paid mutator transaction binding the contract method 0x74a25d43.
//
// Solidity: function startAuction(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64) returns()
func (_Cozy *CozyTransactorSession) StartAuction(_pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64) (*types.Transaction, error) {
	return _Cozy.Contract.StartAuction(&_Cozy.TransactOpts, _pepeId, _beginPrice, _endPrice, _duration)
}

// StartAuctionDirect is a paid mutator transaction binding the contract method 0xd024cd02.
//
// Solidity: function startAuctionDirect(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64, _seller address) returns()
func (_Cozy *CozyTransactor) StartAuctionDirect(opts *bind.TransactOpts, _pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64, _seller common.Address) (*types.Transaction, error) {
	return _Cozy.contract.Transact(opts, "startAuctionDirect", _pepeId, _beginPrice, _endPrice, _duration, _seller)
}

// StartAuctionDirect is a paid mutator transaction binding the contract method 0xd024cd02.
//
// Solidity: function startAuctionDirect(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64, _seller address) returns()
func (_Cozy *CozySession) StartAuctionDirect(_pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64, _seller common.Address) (*types.Transaction, error) {
	return _Cozy.Contract.StartAuctionDirect(&_Cozy.TransactOpts, _pepeId, _beginPrice, _endPrice, _duration, _seller)
}

// StartAuctionDirect is a paid mutator transaction binding the contract method 0xd024cd02.
//
// Solidity: function startAuctionDirect(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64, _seller address) returns()
func (_Cozy *CozyTransactorSession) StartAuctionDirect(_pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64, _seller common.Address) (*types.Transaction, error) {
	return _Cozy.Contract.StartAuctionDirect(&_Cozy.TransactOpts, _pepeId, _beginPrice, _endPrice, _duration, _seller)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Cozy *CozyTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Cozy.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Cozy *CozySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Cozy.Contract.TransferOwnership(&_Cozy.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Cozy *CozyTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Cozy.Contract.TransferOwnership(&_Cozy.TransactOpts, _newOwner)
}

// CozyAuctionFinalizedIterator is returned from FilterAuctionFinalized and is used to iterate over the raw logs and unpacked data for AuctionFinalized events raised by the Cozy contract.
type CozyAuctionFinalizedIterator struct {
	Event *CozyAuctionFinalized // Event containing the contract specifics and raw log

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
func (it *CozyAuctionFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CozyAuctionFinalized)
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
		it.Event = new(CozyAuctionFinalized)
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
func (it *CozyAuctionFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CozyAuctionFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CozyAuctionFinalized represents a AuctionFinalized event raised by the Cozy contract.
type CozyAuctionFinalized struct {
	Pepe   *big.Int
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionFinalized is a free log retrieval operation binding the contract event 0x95b73f79c6d7b09d4dd9a323589aec50a424621f53a70ece1cc21aa75554b519.
//
// Solidity: e AuctionFinalized(pepe indexed uint256, seller indexed address)
func (_Cozy *CozyFilterer) FilterAuctionFinalized(opts *bind.FilterOpts, pepe []*big.Int, seller []common.Address) (*CozyAuctionFinalizedIterator, error) {

	var pepeRule []interface{}
	for _, pepeItem := range pepe {
		pepeRule = append(pepeRule, pepeItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Cozy.contract.FilterLogs(opts, "AuctionFinalized", pepeRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &CozyAuctionFinalizedIterator{contract: _Cozy.contract, event: "AuctionFinalized", logs: logs, sub: sub}, nil
}

// WatchAuctionFinalized is a free log subscription operation binding the contract event 0x95b73f79c6d7b09d4dd9a323589aec50a424621f53a70ece1cc21aa75554b519.
//
// Solidity: e AuctionFinalized(pepe indexed uint256, seller indexed address)
func (_Cozy *CozyFilterer) WatchAuctionFinalized(opts *bind.WatchOpts, sink chan<- *CozyAuctionFinalized, pepe []*big.Int, seller []common.Address) (event.Subscription, error) {

	var pepeRule []interface{}
	for _, pepeItem := range pepe {
		pepeRule = append(pepeRule, pepeItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Cozy.contract.WatchLogs(opts, "AuctionFinalized", pepeRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CozyAuctionFinalized)
				if err := _Cozy.contract.UnpackLog(event, "AuctionFinalized", log); err != nil {
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

// CozyAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the Cozy contract.
type CozyAuctionStartedIterator struct {
	Event *CozyAuctionStarted // Event containing the contract specifics and raw log

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
func (it *CozyAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CozyAuctionStarted)
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
		it.Event = new(CozyAuctionStarted)
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
func (it *CozyAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CozyAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CozyAuctionStarted represents a AuctionStarted event raised by the Cozy contract.
type CozyAuctionStarted struct {
	Pepe   *big.Int
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x16da476d7265fc95576888b93de4fa4849d6cea1228235887f569c6530ddfec1.
//
// Solidity: e AuctionStarted(pepe indexed uint256, seller indexed address)
func (_Cozy *CozyFilterer) FilterAuctionStarted(opts *bind.FilterOpts, pepe []*big.Int, seller []common.Address) (*CozyAuctionStartedIterator, error) {

	var pepeRule []interface{}
	for _, pepeItem := range pepe {
		pepeRule = append(pepeRule, pepeItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Cozy.contract.FilterLogs(opts, "AuctionStarted", pepeRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &CozyAuctionStartedIterator{contract: _Cozy.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x16da476d7265fc95576888b93de4fa4849d6cea1228235887f569c6530ddfec1.
//
// Solidity: e AuctionStarted(pepe indexed uint256, seller indexed address)
func (_Cozy *CozyFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *CozyAuctionStarted, pepe []*big.Int, seller []common.Address) (event.Subscription, error) {

	var pepeRule []interface{}
	for _, pepeItem := range pepe {
		pepeRule = append(pepeRule, pepeItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Cozy.contract.WatchLogs(opts, "AuctionStarted", pepeRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CozyAuctionStarted)
				if err := _Cozy.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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

// CozyAuctionWonIterator is returned from FilterAuctionWon and is used to iterate over the raw logs and unpacked data for AuctionWon events raised by the Cozy contract.
type CozyAuctionWonIterator struct {
	Event *CozyAuctionWon // Event containing the contract specifics and raw log

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
func (it *CozyAuctionWonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CozyAuctionWon)
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
		it.Event = new(CozyAuctionWon)
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
func (it *CozyAuctionWonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CozyAuctionWonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CozyAuctionWon represents a AuctionWon event raised by the Cozy contract.
type CozyAuctionWon struct {
	Pepe   *big.Int
	Winner common.Address
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionWon is a free log retrieval operation binding the contract event 0x94ffdfa810a5f08da0ec8ea0f74619814453cfc101a90504a3a8f77e0eb61986.
//
// Solidity: e AuctionWon(pepe indexed uint256, winner indexed address, seller indexed address)
func (_Cozy *CozyFilterer) FilterAuctionWon(opts *bind.FilterOpts, pepe []*big.Int, winner []common.Address, seller []common.Address) (*CozyAuctionWonIterator, error) {

	var pepeRule []interface{}
	for _, pepeItem := range pepe {
		pepeRule = append(pepeRule, pepeItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Cozy.contract.FilterLogs(opts, "AuctionWon", pepeRule, winnerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &CozyAuctionWonIterator{contract: _Cozy.contract, event: "AuctionWon", logs: logs, sub: sub}, nil
}

// WatchAuctionWon is a free log subscription operation binding the contract event 0x94ffdfa810a5f08da0ec8ea0f74619814453cfc101a90504a3a8f77e0eb61986.
//
// Solidity: e AuctionWon(pepe indexed uint256, winner indexed address, seller indexed address)
func (_Cozy *CozyFilterer) WatchAuctionWon(opts *bind.WatchOpts, sink chan<- *CozyAuctionWon, pepe []*big.Int, winner []common.Address, seller []common.Address) (event.Subscription, error) {

	var pepeRule []interface{}
	for _, pepeItem := range pepe {
		pepeRule = append(pepeRule, pepeItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Cozy.contract.WatchLogs(opts, "AuctionWon", pepeRule, winnerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CozyAuctionWon)
				if err := _Cozy.contract.UnpackLog(event, "AuctionWon", log); err != nil {
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

// CozyOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Cozy contract.
type CozyOwnershipRenouncedIterator struct {
	Event *CozyOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *CozyOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CozyOwnershipRenounced)
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
		it.Event = new(CozyOwnershipRenounced)
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
func (it *CozyOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CozyOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CozyOwnershipRenounced represents a OwnershipRenounced event raised by the Cozy contract.
type CozyOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Cozy *CozyFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*CozyOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Cozy.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CozyOwnershipRenouncedIterator{contract: _Cozy.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Cozy *CozyFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *CozyOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Cozy.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CozyOwnershipRenounced)
				if err := _Cozy.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// CozyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cozy contract.
type CozyOwnershipTransferredIterator struct {
	Event *CozyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CozyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CozyOwnershipTransferred)
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
		it.Event = new(CozyOwnershipTransferred)
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
func (it *CozyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CozyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CozyOwnershipTransferred represents a OwnershipTransferred event raised by the Cozy contract.
type CozyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Cozy *CozyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CozyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cozy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CozyOwnershipTransferredIterator{contract: _Cozy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Cozy *CozyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CozyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cozy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CozyOwnershipTransferred)
				if err := _Cozy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
