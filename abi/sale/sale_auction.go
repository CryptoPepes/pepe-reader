// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sale

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

// SaleABI is the input ABI used to generate the binding from.
const SaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"affiliateContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pepeContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"}],\"name\":\"calculateBid\",\"outputs\":[{\"name\":\"currentBid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"pepeId\",\"type\":\"uint256\"},{\"name\":\"auctionBegin\",\"type\":\"uint64\"},{\"name\":\"auctionEnd\",\"type\":\"uint64\"},{\"name\":\"beginPrice\",\"type\":\"uint256\"},{\"name\":\"endPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"},{\"name\":\"_beginPrice\",\"type\":\"uint256\"},{\"name\":\"_endPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint64\"}],\"name\":\"startAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_DIVIDER\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"}],\"name\":\"savePepe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"},{\"name\":\"_beginPrice\",\"type\":\"uint256\"},{\"name\":\"_endPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint64\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"startAuctionDirect\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_pepeContract\",\"type\":\"address\"},{\"name\":\"_affiliateContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pepe\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"AuctionWon\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pepe\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pepe\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"AuctionFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"}],\"name\":\"buyPepe\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pepeId\",\"type\":\"uint256\"},{\"name\":\"_affiliate\",\"type\":\"address\"}],\"name\":\"buyPepeAffiliated\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// SaleBin is the compiled bytecode used for deploying new contracts.
const SaleBin = `0x608060405261927c60055534801561001657600080fd5b506040516040806110238339810160405280516020909101516000805433600160a060020a0319918216811783556001805483169091179055600380548216600160a060020a039586161790556004805490911693909216929092179055610f9f90819061008490396000f3006080604052600436106100fb5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630117200581146101005780631887a4d0146101315780631c31f710146101465780632d4a170a1461016957806338af3eed14610193578063571a26a0146101a85780636a1db1bf1461020a578063715018a61461022257806374a25d431461023757806378d6ade4146102625780638da5cb5b14610277578063a8f5c6731461028c578063b9a82f80146102a4578063d024cd02146102af578063db8d55f1146102e6578063ddca3f43146102fb578063f2fde38b14610310578063f49dec7014610331575b600080fd5b34801561010c57600080fd5b50610115610348565b60408051600160a060020a039092168252519081900360200190f35b34801561013d57600080fd5b50610115610357565b34801561015257600080fd5b50610167600160a060020a0360043516610366565b005b34801561017557600080fd5b506101816004356103ac565b60408051918252519081900360200190f35b34801561019f57600080fd5b50610115610451565b3480156101b457600080fd5b506101c0600435610460565b60408051600160a060020a039097168752602087019590955267ffffffffffffffff93841686860152919092166060850152608084019190915260a0830152519081900360c00190f35b34801561021657600080fd5b506101676004356104b3565b34801561022e57600080fd5b506101676104dd565b34801561024357600080fd5b5061016760043560243560443567ffffffffffffffff60643516610549565b34801561026e57600080fd5b50610181610749565b34801561028357600080fd5b50610115610750565b34801561029857600080fd5b5061016760043561075f565b6101676004356108e1565b3480156102bb57600080fd5b5061016760043560243560443567ffffffffffffffff60643516600160a060020a0360843516610c54565b3480156102f257600080fd5b50610167610dcb565b34801561030757600080fd5b50610181610e08565b34801561031c57600080fd5b50610167600160a060020a0360043516610e0e565b610167600435600160a060020a0360243516610e2e565b600454600160a060020a031681565b600354600160a060020a031681565b600054600160a060020a0316331461037d57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600081815260026020819052604082209081015467ffffffffffffffff808216429081039285928392839283926801000000000000000090910416116103f85785600401549650610446565b600386015460048701546002880154919003945067ffffffffffffffff80821668010000000000000000909204160392508285850281151561043657fe5b0591508186600301540190508096505b505050505050919050565b600154600160a060020a031681565b600260208190526000918252604090912080546001820154928201546003830154600490930154600160a060020a03909216939267ffffffffffffffff8083169368010000000000000000909304169186565b600054600160a060020a031633146104ca57600080fd5b60055481106104d857600080fd5b600555565b600054600160a060020a031633146104f457600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b610551610f3e565b600354604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018890529051600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b1580156105c457600080fd5b505af11580156105d8573d6000803e3d6000fd5b505050506040513d60208110156105ee57600080fd5b505115156105fb57600080fd5b6000858152600260208190526040909120015468010000000000000000900467ffffffffffffffff16421161062f57600080fd5b338152602081018590524267ffffffffffffffff8181166040840181905291840116606083018190521161066257600080fd5b6080810184815260a0820184815260008781526002602081815260408084208751815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178155918701516001830155808701519282018054606089015167ffffffffffffffff1990911667ffffffffffffffff958616176fffffffffffffffff000000000000000019166801000000000000000095909116949094029390931790925593516003850155915160049093019290925551339187917f16da476d7265fc95576888b93de4fa4849d6cea1228235887f569c6530ddfec19190a35050505050565b620f424081565b600054600160a060020a031681565b60008181526002602081905260409091200154426801000000000000000090910467ffffffffffffffff161061079457600080fd5b60035460008281526002602090815260408083205481517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a03918216600482015260248101879052915194169363a9059cbb93604480840194938390030190829087803b15801561080d57600080fd5b505af1158015610821573d6000803e3d6000fd5b505050506040513d602081101561083757600080fd5b5051151561084457600080fd5b600081815260026020526040808220549051600160a060020a039091169183917f95b73f79c6d7b09d4dd9a323589aec50a424621f53a70ece1cc21aa75554b5199190a360009081526002602081905260408220805473ffffffffffffffffffffffffffffffffffffffff191681556001810183905590810180546fffffffffffffffffffffffffffffffff191690556003810182905560040155565b6000818152600260208190526040822090810154909190819068010000000000000000900467ffffffffffffffff16421061091b57600080fd5b610924846103ac565b91503482111561093357600080fd5b600554620f42409083028454604051929091049250600160a060020a03169082840380156108fc02916000818181858888f1935050505015801561097b573d6000803e3d6000fd5b5060048054604080517f873abd75000000000000000000000000000000000000000000000000000000008152339381019390935251600092600160a060020a039092169163873abd7591602480830192602092919082900301818787803b1580156109e557600080fd5b505af11580156109f9573d6000803e3d6000fd5b505050506040513d6020811015610a0f57600080fd5b5051600160a060020a031614801590610adf575060048054604080517f873abd75000000000000000000000000000000000000000000000000000000008152339381019390935251600160a060020a039091169163873abd759160248083019260209291908290030181600087803b158015610a8a57600080fd5b505af1158015610a9e573d6000803e3d6000fd5b505050506040513d6020811015610ab457600080fd5b5051604051600160a060020a03909116906002830480156108fc02916000818181858888f193505050505b50600354604080517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018790529051600160a060020a039092169163a9059cbb916044808201926020929091908290030181600087803b158015610b4d57600080fd5b505af1158015610b61573d6000803e3d6000fd5b505050506040513d6020811015610b7757600080fd5b50511515610b8457600080fd5b8254604051600160a060020a0390911690339086907f94ffdfa810a5f08da0ec8ea0f74619814453cfc101a90504a3a8f77e0eb6198690600090a481341115610bf85760405133903484900380156108fc02916000818181858888f19350505050158015610bf6573d6000803e3d6000fd5b505b50505060009081526002602081905260408220805473ffffffffffffffffffffffffffffffffffffffff191681556001810183905590810180546fffffffffffffffffffffffffffffffff191690556003810182905560040155565b610c5c610f3e565b600354600160a060020a03163314610c7357600080fd5b6000868152600260208190526040909120015468010000000000000000900467ffffffffffffffff164211610ca757600080fd5b600160a060020a0382168152602081018690524267ffffffffffffffff81811660408401819052918501166060830181905211610ce357600080fd5b6080810185815260a0820185815260008881526002602081815260408084208751815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391821617825592880151600182015581880151938101805460608a015167ffffffffffffffff1990911667ffffffffffffffff968716176fffffffffffffffff000000000000000019166801000000000000000096909116959095029490941790935594516003830155925160049091015591519084169188917f16da476d7265fc95576888b93de4fa4849d6cea1228235887f569c6530ddfec19190a3505050505050565b600154604051600160a060020a0390911690303180156108fc02916000818181858888f19350505050158015610e05573d6000803e3d6000fd5b50565b60055481565b600054600160a060020a03163314610e2557600080fd5b610e0581610ec1565b60048054604080517f7352e4b80000000000000000000000000000000000000000000000000000000081523393810193909352600160a060020a0384811660248501529051911691637352e4b891604480830192600092919082900301818387803b158015610e9c57600080fd5b505af1158015610eb0573d6000803e3d6000fd5b50505050610ebd826108e1565b5050565b600160a060020a0381161515610ed657600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152905600a165627a7a72305820df9ea8ddb5fe166c6751fcfd2ae5cc503ef94034442974353e7b7066e65cc7090029`

// DeploySale deploys a new Ethereum contract, binding an instance of Sale to it.
func DeploySale(auth *bind.TransactOpts, backend bind.ContractBackend, _pepeContract common.Address, _affiliateContract common.Address) (common.Address, *types.Transaction, *Sale, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SaleBin), backend, _pepeContract, _affiliateContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sale{SaleCaller: SaleCaller{contract: contract}, SaleTransactor: SaleTransactor{contract: contract}, SaleFilterer: SaleFilterer{contract: contract}}, nil
}

// Sale is an auto generated Go binding around an Ethereum contract.
type Sale struct {
	SaleCaller     // Read-only binding to the contract
	SaleTransactor // Write-only binding to the contract
	SaleFilterer   // Log filterer for contract events
}

// SaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaleSession struct {
	Contract     *Sale             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaleCallerSession struct {
	Contract *SaleCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaleTransactorSession struct {
	Contract     *SaleTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaleRaw struct {
	Contract *Sale // Generic contract binding to access the raw methods on
}

// SaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaleCallerRaw struct {
	Contract *SaleCaller // Generic read-only contract binding to access the raw methods on
}

// SaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaleTransactorRaw struct {
	Contract *SaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSale creates a new instance of Sale, bound to a specific deployed contract.
func NewSale(address common.Address, backend bind.ContractBackend) (*Sale, error) {
	contract, err := bindSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sale{SaleCaller: SaleCaller{contract: contract}, SaleTransactor: SaleTransactor{contract: contract}, SaleFilterer: SaleFilterer{contract: contract}}, nil
}

// NewSaleCaller creates a new read-only instance of Sale, bound to a specific deployed contract.
func NewSaleCaller(address common.Address, caller bind.ContractCaller) (*SaleCaller, error) {
	contract, err := bindSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaleCaller{contract: contract}, nil
}

// NewSaleTransactor creates a new write-only instance of Sale, bound to a specific deployed contract.
func NewSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*SaleTransactor, error) {
	contract, err := bindSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaleTransactor{contract: contract}, nil
}

// NewSaleFilterer creates a new log filterer instance of Sale, bound to a specific deployed contract.
func NewSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*SaleFilterer, error) {
	contract, err := bindSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaleFilterer{contract: contract}, nil
}

// bindSale binds a generic wrapper to an already deployed contract.
func bindSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale *SaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sale.Contract.SaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale *SaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.Contract.SaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale *SaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale.Contract.SaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale *SaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Sale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale *SaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale *SaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale.Contract.contract.Transact(opts, method, params...)
}

// FEEDIVIDER is a free data retrieval call binding the contract method 0x78d6ade4.
//
// Solidity: function FEE_DIVIDER() constant returns(uint256)
func (_Sale *SaleCaller) FEEDIVIDER(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sale.contract.Call(opts, out, "FEE_DIVIDER")
	return *ret0, err
}

// FEEDIVIDER is a free data retrieval call binding the contract method 0x78d6ade4.
//
// Solidity: function FEE_DIVIDER() constant returns(uint256)
func (_Sale *SaleSession) FEEDIVIDER() (*big.Int, error) {
	return _Sale.Contract.FEEDIVIDER(&_Sale.CallOpts)
}

// FEEDIVIDER is a free data retrieval call binding the contract method 0x78d6ade4.
//
// Solidity: function FEE_DIVIDER() constant returns(uint256)
func (_Sale *SaleCallerSession) FEEDIVIDER() (*big.Int, error) {
	return _Sale.Contract.FEEDIVIDER(&_Sale.CallOpts)
}

// AffiliateContract is a free data retrieval call binding the contract method 0x01172005.
//
// Solidity: function affiliateContract() constant returns(address)
func (_Sale *SaleCaller) AffiliateContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sale.contract.Call(opts, out, "affiliateContract")
	return *ret0, err
}

// AffiliateContract is a free data retrieval call binding the contract method 0x01172005.
//
// Solidity: function affiliateContract() constant returns(address)
func (_Sale *SaleSession) AffiliateContract() (common.Address, error) {
	return _Sale.Contract.AffiliateContract(&_Sale.CallOpts)
}

// AffiliateContract is a free data retrieval call binding the contract method 0x01172005.
//
// Solidity: function affiliateContract() constant returns(address)
func (_Sale *SaleCallerSession) AffiliateContract() (common.Address, error) {
	return _Sale.Contract.AffiliateContract(&_Sale.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions( uint256) constant returns(seller address, pepeId uint256, auctionBegin uint64, auctionEnd uint64, beginPrice uint256, endPrice uint256)
func (_Sale *SaleCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Sale.contract.Call(opts, out, "auctions", arg0)
	return *ret, err
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions( uint256) constant returns(seller address, pepeId uint256, auctionBegin uint64, auctionEnd uint64, beginPrice uint256, endPrice uint256)
func (_Sale *SaleSession) Auctions(arg0 *big.Int) (struct {
	Seller       common.Address
	PepeId       *big.Int
	AuctionBegin uint64
	AuctionEnd   uint64
	BeginPrice   *big.Int
	EndPrice     *big.Int
}, error) {
	return _Sale.Contract.Auctions(&_Sale.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions( uint256) constant returns(seller address, pepeId uint256, auctionBegin uint64, auctionEnd uint64, beginPrice uint256, endPrice uint256)
func (_Sale *SaleCallerSession) Auctions(arg0 *big.Int) (struct {
	Seller       common.Address
	PepeId       *big.Int
	AuctionBegin uint64
	AuctionEnd   uint64
	BeginPrice   *big.Int
	EndPrice     *big.Int
}, error) {
	return _Sale.Contract.Auctions(&_Sale.CallOpts, arg0)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_Sale *SaleCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sale.contract.Call(opts, out, "beneficiary")
	return *ret0, err
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_Sale *SaleSession) Beneficiary() (common.Address, error) {
	return _Sale.Contract.Beneficiary(&_Sale.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_Sale *SaleCallerSession) Beneficiary() (common.Address, error) {
	return _Sale.Contract.Beneficiary(&_Sale.CallOpts)
}

// CalculateBid is a free data retrieval call binding the contract method 0x2d4a170a.
//
// Solidity: function calculateBid(_pepeId uint256) constant returns(currentBid uint256)
func (_Sale *SaleCaller) CalculateBid(opts *bind.CallOpts, _pepeId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sale.contract.Call(opts, out, "calculateBid", _pepeId)
	return *ret0, err
}

// CalculateBid is a free data retrieval call binding the contract method 0x2d4a170a.
//
// Solidity: function calculateBid(_pepeId uint256) constant returns(currentBid uint256)
func (_Sale *SaleSession) CalculateBid(_pepeId *big.Int) (*big.Int, error) {
	return _Sale.Contract.CalculateBid(&_Sale.CallOpts, _pepeId)
}

// CalculateBid is a free data retrieval call binding the contract method 0x2d4a170a.
//
// Solidity: function calculateBid(_pepeId uint256) constant returns(currentBid uint256)
func (_Sale *SaleCallerSession) CalculateBid(_pepeId *big.Int) (*big.Int, error) {
	return _Sale.Contract.CalculateBid(&_Sale.CallOpts, _pepeId)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Sale *SaleCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Sale.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Sale *SaleSession) Fee() (*big.Int, error) {
	return _Sale.Contract.Fee(&_Sale.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_Sale *SaleCallerSession) Fee() (*big.Int, error) {
	return _Sale.Contract.Fee(&_Sale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sale *SaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sale.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sale *SaleSession) Owner() (common.Address, error) {
	return _Sale.Contract.Owner(&_Sale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Sale *SaleCallerSession) Owner() (common.Address, error) {
	return _Sale.Contract.Owner(&_Sale.CallOpts)
}

// PepeContract is a free data retrieval call binding the contract method 0x1887a4d0.
//
// Solidity: function pepeContract() constant returns(address)
func (_Sale *SaleCaller) PepeContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Sale.contract.Call(opts, out, "pepeContract")
	return *ret0, err
}

// PepeContract is a free data retrieval call binding the contract method 0x1887a4d0.
//
// Solidity: function pepeContract() constant returns(address)
func (_Sale *SaleSession) PepeContract() (common.Address, error) {
	return _Sale.Contract.PepeContract(&_Sale.CallOpts)
}

// PepeContract is a free data retrieval call binding the contract method 0x1887a4d0.
//
// Solidity: function pepeContract() constant returns(address)
func (_Sale *SaleCallerSession) PepeContract() (common.Address, error) {
	return _Sale.Contract.PepeContract(&_Sale.CallOpts)
}

// BuyPepe is a paid mutator transaction binding the contract method 0xb9a82f80.
//
// Solidity: function buyPepe(_pepeId uint256) returns()
func (_Sale *SaleTransactor) BuyPepe(opts *bind.TransactOpts, _pepeId *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "buyPepe", _pepeId)
}

// BuyPepe is a paid mutator transaction binding the contract method 0xb9a82f80.
//
// Solidity: function buyPepe(_pepeId uint256) returns()
func (_Sale *SaleSession) BuyPepe(_pepeId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.BuyPepe(&_Sale.TransactOpts, _pepeId)
}

// BuyPepe is a paid mutator transaction binding the contract method 0xb9a82f80.
//
// Solidity: function buyPepe(_pepeId uint256) returns()
func (_Sale *SaleTransactorSession) BuyPepe(_pepeId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.BuyPepe(&_Sale.TransactOpts, _pepeId)
}

// BuyPepeAffiliated is a paid mutator transaction binding the contract method 0xf49dec70.
//
// Solidity: function buyPepeAffiliated(_pepeId uint256, _affiliate address) returns()
func (_Sale *SaleTransactor) BuyPepeAffiliated(opts *bind.TransactOpts, _pepeId *big.Int, _affiliate common.Address) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "buyPepeAffiliated", _pepeId, _affiliate)
}

// BuyPepeAffiliated is a paid mutator transaction binding the contract method 0xf49dec70.
//
// Solidity: function buyPepeAffiliated(_pepeId uint256, _affiliate address) returns()
func (_Sale *SaleSession) BuyPepeAffiliated(_pepeId *big.Int, _affiliate common.Address) (*types.Transaction, error) {
	return _Sale.Contract.BuyPepeAffiliated(&_Sale.TransactOpts, _pepeId, _affiliate)
}

// BuyPepeAffiliated is a paid mutator transaction binding the contract method 0xf49dec70.
//
// Solidity: function buyPepeAffiliated(_pepeId uint256, _affiliate address) returns()
func (_Sale *SaleTransactorSession) BuyPepeAffiliated(_pepeId *big.Int, _affiliate common.Address) (*types.Transaction, error) {
	return _Sale.Contract.BuyPepeAffiliated(&_Sale.TransactOpts, _pepeId, _affiliate)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(_fee uint256) returns()
func (_Sale *SaleTransactor) ChangeFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "changeFee", _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(_fee uint256) returns()
func (_Sale *SaleSession) ChangeFee(_fee *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.ChangeFee(&_Sale.TransactOpts, _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(_fee uint256) returns()
func (_Sale *SaleTransactorSession) ChangeFee(_fee *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.ChangeFee(&_Sale.TransactOpts, _fee)
}

// GetFees is a paid mutator transaction binding the contract method 0xdb8d55f1.
//
// Solidity: function getFees() returns()
func (_Sale *SaleTransactor) GetFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "getFees")
}

// GetFees is a paid mutator transaction binding the contract method 0xdb8d55f1.
//
// Solidity: function getFees() returns()
func (_Sale *SaleSession) GetFees() (*types.Transaction, error) {
	return _Sale.Contract.GetFees(&_Sale.TransactOpts)
}

// GetFees is a paid mutator transaction binding the contract method 0xdb8d55f1.
//
// Solidity: function getFees() returns()
func (_Sale *SaleTransactorSession) GetFees() (*types.Transaction, error) {
	return _Sale.Contract.GetFees(&_Sale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sale *SaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sale *SaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sale.Contract.RenounceOwnership(&_Sale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sale *SaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sale.Contract.RenounceOwnership(&_Sale.TransactOpts)
}

// SavePepe is a paid mutator transaction binding the contract method 0xa8f5c673.
//
// Solidity: function savePepe(_pepeId uint256) returns()
func (_Sale *SaleTransactor) SavePepe(opts *bind.TransactOpts, _pepeId *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "savePepe", _pepeId)
}

// SavePepe is a paid mutator transaction binding the contract method 0xa8f5c673.
//
// Solidity: function savePepe(_pepeId uint256) returns()
func (_Sale *SaleSession) SavePepe(_pepeId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.SavePepe(&_Sale.TransactOpts, _pepeId)
}

// SavePepe is a paid mutator transaction binding the contract method 0xa8f5c673.
//
// Solidity: function savePepe(_pepeId uint256) returns()
func (_Sale *SaleTransactorSession) SavePepe(_pepeId *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.SavePepe(&_Sale.TransactOpts, _pepeId)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(_beneficiary address) returns()
func (_Sale *SaleTransactor) SetBeneficiary(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "setBeneficiary", _beneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(_beneficiary address) returns()
func (_Sale *SaleSession) SetBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _Sale.Contract.SetBeneficiary(&_Sale.TransactOpts, _beneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(_beneficiary address) returns()
func (_Sale *SaleTransactorSession) SetBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _Sale.Contract.SetBeneficiary(&_Sale.TransactOpts, _beneficiary)
}

// StartAuction is a paid mutator transaction binding the contract method 0x74a25d43.
//
// Solidity: function startAuction(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64) returns()
func (_Sale *SaleTransactor) StartAuction(opts *bind.TransactOpts, _pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "startAuction", _pepeId, _beginPrice, _endPrice, _duration)
}

// StartAuction is a paid mutator transaction binding the contract method 0x74a25d43.
//
// Solidity: function startAuction(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64) returns()
func (_Sale *SaleSession) StartAuction(_pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64) (*types.Transaction, error) {
	return _Sale.Contract.StartAuction(&_Sale.TransactOpts, _pepeId, _beginPrice, _endPrice, _duration)
}

// StartAuction is a paid mutator transaction binding the contract method 0x74a25d43.
//
// Solidity: function startAuction(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64) returns()
func (_Sale *SaleTransactorSession) StartAuction(_pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64) (*types.Transaction, error) {
	return _Sale.Contract.StartAuction(&_Sale.TransactOpts, _pepeId, _beginPrice, _endPrice, _duration)
}

// StartAuctionDirect is a paid mutator transaction binding the contract method 0xd024cd02.
//
// Solidity: function startAuctionDirect(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64, _seller address) returns()
func (_Sale *SaleTransactor) StartAuctionDirect(opts *bind.TransactOpts, _pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64, _seller common.Address) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "startAuctionDirect", _pepeId, _beginPrice, _endPrice, _duration, _seller)
}

// StartAuctionDirect is a paid mutator transaction binding the contract method 0xd024cd02.
//
// Solidity: function startAuctionDirect(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64, _seller address) returns()
func (_Sale *SaleSession) StartAuctionDirect(_pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64, _seller common.Address) (*types.Transaction, error) {
	return _Sale.Contract.StartAuctionDirect(&_Sale.TransactOpts, _pepeId, _beginPrice, _endPrice, _duration, _seller)
}

// StartAuctionDirect is a paid mutator transaction binding the contract method 0xd024cd02.
//
// Solidity: function startAuctionDirect(_pepeId uint256, _beginPrice uint256, _endPrice uint256, _duration uint64, _seller address) returns()
func (_Sale *SaleTransactorSession) StartAuctionDirect(_pepeId *big.Int, _beginPrice *big.Int, _endPrice *big.Int, _duration uint64, _seller common.Address) (*types.Transaction, error) {
	return _Sale.Contract.StartAuctionDirect(&_Sale.TransactOpts, _pepeId, _beginPrice, _endPrice, _duration, _seller)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Sale *SaleTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Sale *SaleSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Sale.Contract.TransferOwnership(&_Sale.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Sale *SaleTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Sale.Contract.TransferOwnership(&_Sale.TransactOpts, _newOwner)
}

// SaleAuctionFinalizedIterator is returned from FilterAuctionFinalized and is used to iterate over the raw logs and unpacked data for AuctionFinalized events raised by the Sale contract.
type SaleAuctionFinalizedIterator struct {
	Event *SaleAuctionFinalized // Event containing the contract specifics and raw log

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
func (it *SaleAuctionFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleAuctionFinalized)
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
		it.Event = new(SaleAuctionFinalized)
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
func (it *SaleAuctionFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleAuctionFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleAuctionFinalized represents a AuctionFinalized event raised by the Sale contract.
type SaleAuctionFinalized struct {
	Pepe   *big.Int
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionFinalized is a free log retrieval operation binding the contract event 0x95b73f79c6d7b09d4dd9a323589aec50a424621f53a70ece1cc21aa75554b519.
//
// Solidity: e AuctionFinalized(pepe indexed uint256, seller indexed address)
func (_Sale *SaleFilterer) FilterAuctionFinalized(opts *bind.FilterOpts, pepe []*big.Int, seller []common.Address) (*SaleAuctionFinalizedIterator, error) {

	var pepeRule []interface{}
	for _, pepeItem := range pepe {
		pepeRule = append(pepeRule, pepeItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Sale.contract.FilterLogs(opts, "AuctionFinalized", pepeRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &SaleAuctionFinalizedIterator{contract: _Sale.contract, event: "AuctionFinalized", logs: logs, sub: sub}, nil
}

// WatchAuctionFinalized is a free log subscription operation binding the contract event 0x95b73f79c6d7b09d4dd9a323589aec50a424621f53a70ece1cc21aa75554b519.
//
// Solidity: e AuctionFinalized(pepe indexed uint256, seller indexed address)
func (_Sale *SaleFilterer) WatchAuctionFinalized(opts *bind.WatchOpts, sink chan<- *SaleAuctionFinalized, pepe []*big.Int, seller []common.Address) (event.Subscription, error) {

	var pepeRule []interface{}
	for _, pepeItem := range pepe {
		pepeRule = append(pepeRule, pepeItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Sale.contract.WatchLogs(opts, "AuctionFinalized", pepeRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleAuctionFinalized)
				if err := _Sale.contract.UnpackLog(event, "AuctionFinalized", log); err != nil {
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

// SaleAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the Sale contract.
type SaleAuctionStartedIterator struct {
	Event *SaleAuctionStarted // Event containing the contract specifics and raw log

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
func (it *SaleAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleAuctionStarted)
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
		it.Event = new(SaleAuctionStarted)
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
func (it *SaleAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleAuctionStarted represents a AuctionStarted event raised by the Sale contract.
type SaleAuctionStarted struct {
	Pepe   *big.Int
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x16da476d7265fc95576888b93de4fa4849d6cea1228235887f569c6530ddfec1.
//
// Solidity: e AuctionStarted(pepe indexed uint256, seller indexed address)
func (_Sale *SaleFilterer) FilterAuctionStarted(opts *bind.FilterOpts, pepe []*big.Int, seller []common.Address) (*SaleAuctionStartedIterator, error) {

	var pepeRule []interface{}
	for _, pepeItem := range pepe {
		pepeRule = append(pepeRule, pepeItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Sale.contract.FilterLogs(opts, "AuctionStarted", pepeRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &SaleAuctionStartedIterator{contract: _Sale.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x16da476d7265fc95576888b93de4fa4849d6cea1228235887f569c6530ddfec1.
//
// Solidity: e AuctionStarted(pepe indexed uint256, seller indexed address)
func (_Sale *SaleFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *SaleAuctionStarted, pepe []*big.Int, seller []common.Address) (event.Subscription, error) {

	var pepeRule []interface{}
	for _, pepeItem := range pepe {
		pepeRule = append(pepeRule, pepeItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Sale.contract.WatchLogs(opts, "AuctionStarted", pepeRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleAuctionStarted)
				if err := _Sale.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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

// SaleAuctionWonIterator is returned from FilterAuctionWon and is used to iterate over the raw logs and unpacked data for AuctionWon events raised by the Sale contract.
type SaleAuctionWonIterator struct {
	Event *SaleAuctionWon // Event containing the contract specifics and raw log

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
func (it *SaleAuctionWonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleAuctionWon)
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
		it.Event = new(SaleAuctionWon)
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
func (it *SaleAuctionWonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleAuctionWonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleAuctionWon represents a AuctionWon event raised by the Sale contract.
type SaleAuctionWon struct {
	Pepe   *big.Int
	Winner common.Address
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionWon is a free log retrieval operation binding the contract event 0x94ffdfa810a5f08da0ec8ea0f74619814453cfc101a90504a3a8f77e0eb61986.
//
// Solidity: e AuctionWon(pepe indexed uint256, winner indexed address, seller indexed address)
func (_Sale *SaleFilterer) FilterAuctionWon(opts *bind.FilterOpts, pepe []*big.Int, winner []common.Address, seller []common.Address) (*SaleAuctionWonIterator, error) {

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

	logs, sub, err := _Sale.contract.FilterLogs(opts, "AuctionWon", pepeRule, winnerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &SaleAuctionWonIterator{contract: _Sale.contract, event: "AuctionWon", logs: logs, sub: sub}, nil
}

// WatchAuctionWon is a free log subscription operation binding the contract event 0x94ffdfa810a5f08da0ec8ea0f74619814453cfc101a90504a3a8f77e0eb61986.
//
// Solidity: e AuctionWon(pepe indexed uint256, winner indexed address, seller indexed address)
func (_Sale *SaleFilterer) WatchAuctionWon(opts *bind.WatchOpts, sink chan<- *SaleAuctionWon, pepe []*big.Int, winner []common.Address, seller []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Sale.contract.WatchLogs(opts, "AuctionWon", pepeRule, winnerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleAuctionWon)
				if err := _Sale.contract.UnpackLog(event, "AuctionWon", log); err != nil {
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

// SaleOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Sale contract.
type SaleOwnershipRenouncedIterator struct {
	Event *SaleOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *SaleOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleOwnershipRenounced)
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
		it.Event = new(SaleOwnershipRenounced)
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
func (it *SaleOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleOwnershipRenounced represents a OwnershipRenounced event raised by the Sale contract.
type SaleOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Sale *SaleFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*SaleOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Sale.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SaleOwnershipRenouncedIterator{contract: _Sale.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Sale *SaleFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *SaleOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Sale.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleOwnershipRenounced)
				if err := _Sale.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// SaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Sale contract.
type SaleOwnershipTransferredIterator struct {
	Event *SaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleOwnershipTransferred)
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
		it.Event = new(SaleOwnershipTransferred)
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
func (it *SaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleOwnershipTransferred represents a OwnershipTransferred event raised by the Sale contract.
type SaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Sale *SaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SaleOwnershipTransferredIterator{contract: _Sale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Sale *SaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleOwnershipTransferred)
				if err := _Sale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
