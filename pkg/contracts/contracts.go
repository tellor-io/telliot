package contracts

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/config"
	"github.com/tellor-io/telliot/pkg/contracts/balancer"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/contracts/uniswap"
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

type ITellor struct {
	*tellor.ITellor
	*tellor.ITellorNewDispute
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
	BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	IsSyncing(ctx context.Context) (bool, error)
	NetworkID(ctx context.Context) (*big.Int, error)
	HeaderByNumber(ctx context.Context, num *big.Int) (*types.Header, error)
}

// Get hard-coded Tellor contract address per network id.
func getContractAddress(client ETHClient) (string, error) {
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	switch networkID.Int64() {
	case 1:
		return config.TellorMainnetAddress, nil
	case 4:
		return config.TellorRinkebyAddress, nil
	default:
		return "", errors.New("contract address for current network id not found")
	}
}

func NewITellor(client ETHClient) (*ITellor, error) {
	_contractAddress, err := getContractAddress(client)
	if err != nil {
		return nil, errors.Wrap(err, "getting tellor contract address")
	}
	contractAddress := common.HexToAddress(_contractAddress)

	contractInterfaceInstance, err := tellor.NewITellor(contractAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating telllor interface")
	}

	return &ITellor{Address: contractAddress, ITellor: contractInterfaceInstance}, nil
}
