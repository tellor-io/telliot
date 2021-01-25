package contracts

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts/master"
	"github.com/tellor-io/telliot/pkg/contracts/proxy"
)

type Tellor struct {
	Getter  *proxy.TellorGetters
	Caller  *master.Tellor
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

const (
	TellorMainnetAddress = "0x0Ba45A8b5d5575935B8158a88C631E9F9C95a2e5"
	TellorRinkebyAddress = "0xFe41Cb708CD98C5B20423433309E55b53F79134a"
)

// Get hard-coded Tellor contract address per network id.
func getContractAddress(client ETHClient) (string, error) {
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	switch networkID.Int64() {
	case 1:
		return TellorMainnetAddress, nil
	case 4:
		return TellorRinkebyAddress, nil
	default:
		return "", errors.New("contract address for current network id not found")
	}
}

func NewTellor(client ETHClient) (Tellor, error) {
	_contractAddress, err := getContractAddress(client)
	if err != nil {
		return Tellor{}, errors.Wrap(err, "getting tellor contract address")
	}
	contractAddress := common.HexToAddress(_contractAddress)
	contractTellorInstance, err := master.NewTellor(contractAddress, client)
	if err != nil {
		return Tellor{}, errors.Wrap(err, "creating telllor caller")
	}
	contractGetterInstance, err := proxy.NewTellorGetters(contractAddress, client)
	if err != nil {
		return Tellor{}, errors.Wrap(err, "creating telllor getter")
	}

	return Tellor{Address: contractAddress, Getter: contractGetterInstance, Caller: contractTellorInstance}, nil
}

func NewTellorGetters(client ETHClient) (*proxy.TellorGetters, error) {
	_contractAddress, err := getContractAddress(client)
	if err != nil {
		return nil, errors.Wrap(err, "getting tellor contract address")
	}
	contractAddress := common.HexToAddress(_contractAddress)
	contractGetterInstance, err := proxy.NewTellorGetters(contractAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "creating telllor getter")
	}

	return contractGetterInstance, nil
}
