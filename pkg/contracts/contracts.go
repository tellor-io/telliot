package contracts

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts/balancer"
	"github.com/tellor-io/telliot/pkg/contracts/lens"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/contracts/tellorAccess"
	"github.com/tellor-io/telliot/pkg/contracts/uniswap"
)

const (
	TellorAddress               = "0x88dF592F8eb5D7Bd38bFeF7dEb0fBc02cf3778a0"
	TellorAccessAddressRinkeby  = "0x5a991dd4f646ed7efdd090b1ba5b68d222273f7e"
	TellorAccessAddressArbitrum = "0xCf26Ce0a3a9EF0125FA53a05A00b6B68F5ddb27A"
	TellorAccessAddress         = "0x5a991dd4f646ed7efdd090b1ba5b68d222273f7e"
	LensAddressMainnet          = "0x577417CFaF319a1fAD90aA135E3848D2C00e68CF"
	LensAddressRinkeby          = "0xebEF7ceB7C43850898e258be0a1ea5ffcdBc3205"
)

type (
	ITellorNewDispute    = tellor.ITellorNewDispute
	TellorNonceSubmitted = tellor.TellorNonceSubmitted
)

const (
	BPoolABI          = balancer.BPoolABI
	BTokenABI         = balancer.BTokenABI
	IERC20ABI         = uniswap.IERC20ABI
	IUniswapV2PairABI = uniswap.IUniswapV2PairABI
	ITellorABI        = tellor.ITellorABI
)

type ITellorAccess struct {
	Address common.Address
	*tellorAccess.TellorAccess
}

type ITellor struct {
	*tellor.ITellor
	*tellor.ITellorNewDispute
	*lens.Main
	Address common.Address
}

// ETHClient is the main abstraction interface for client operations.
type ETHClient interface {

	// Close the client.
	Close()

	// CodeAt returns the code of the given account. This is needed to differentiate
	// between contract internal errors and the local chain being out of sync.
	CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error)

	// TransactionReceipt implements the geth backend DeployBackend interface.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)

	// ContractCall executes an Ethereum contract call with the specified data as the
	// input.
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	NonceAt(ctx context.Context, address common.Address) (uint64, error)
	PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error)

	// PendingCodeAt returns the code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)

	//PendingNonceAt gets the given address's nonce for submitting transactions
	PendingNonceAt(ctx context.Context, address common.Address) (uint64, error)
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)

	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	IsSyncing(ctx context.Context) (bool, error)
	NetworkID(ctx context.Context) (*big.Int, error)
	HeaderByNumber(ctx context.Context, num *big.Int) (*types.Header, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
}

func getLensAddress(client ETHClient) (common.Address, error) {
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Address{}, err
	}
	switch netID := networkID.Int64(); netID {
	case 1:
		return common.HexToAddress(LensAddressMainnet), nil
	case 4:
		return common.HexToAddress(LensAddressRinkeby), nil
	default:
		return common.Address{}, errors.Errorf("contract address for current network id not found:%v", netID)
	}
}

func NewITellor(client ETHClient) (*ITellor, error) {
	tellorInstance, err := tellor.NewITellor(common.HexToAddress(TellorAddress), client)
	if err != nil {
		return nil, errors.Wrap(err, "creating telllor interface")
	}
	contractAddr, err := getLensAddress(client)
	if err != nil {
		return nil, errors.Wrap(err, "creating lens address")
	}

	lensInstance, err := lens.NewMain(contractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating telllor interface")
	}

	return &ITellor{Address: common.HexToAddress(TellorAddress), ITellor: tellorInstance, Main: lensInstance}, nil
}

func NewITellorAccess(client ETHClient) (*ITellorAccess, error) {
	conractAddr, err := getTellorAccessAddress(client)
	if err != nil {
		return nil, errors.Wrap(err, "creating lens address")
	}
	tellorInstance, err := tellorAccess.NewTellorAccess(conractAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating telllor interface")
	}

	return &ITellorAccess{Address: common.HexToAddress(TellorAccessAddress), TellorAccess: tellorInstance}, nil
}

func getTellorAccessAddress(client ETHClient) (common.Address, error) {
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Address{}, err
	}
	switch netID := networkID.Int64(); netID {
	case 144545313136048:
		return common.HexToAddress(TellorAccessAddressArbitrum), nil
	case 4:
		return common.HexToAddress(TellorAccessAddressRinkeby), nil
	default:
		return common.Address{}, errors.Errorf("contract address for current network id not found:%v", netID)
	}
}
