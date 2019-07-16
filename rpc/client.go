package rpc

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//ETHClient is the main abstraction interface for client operations
type ETHClient interface {
	//close the client
	Close()

	// CodeAt returns the code of the given account. This is needed to differentiate
	// between contract internal errors and the local chain being out of sync.
	CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error)
	// ContractCall executes an Ethereum contract call with the specified data as the
	// input.
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)

	//PendingNonceAt gets the given address's nonce for submitting transactions
	PendingNonceAt(ctx context.Context, address common.Address) (uint64, error)
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error)
	//SuggestGasPrice retrieves the current gas price in the network
	SuggestGasPrice(ctx context.Context) (*big.Int, error)

	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	//BalanceAt gets the balance for the given address at the given block.
	BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error

	PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error)
}

//clientInstance is the concrete implementation of the ETHClient
type clientInstance struct {
	ethClient *ethclient.Client
}

//NewClient creates a new client instance
func NewClient(ct context.Context, url string) (ETHClient, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &clientInstance{ethClient: client}, nil
}

func (c *clientInstance) Close() {
	fmt.Println("Closing ETHClient")
	c.ethClient.Close()
}

func (c *clientInstance) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return nil
}

func (c *clientInstance) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	return nil, nil
}

func (c *clientInstance) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return nil, nil
}
func (c *clientInstance) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}
func (c *clientInstance) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}

func (c *clientInstance) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return 0, nil
}

func (c *clientInstance) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	//insert retry logic
	return c.ethClient.CodeAt(ctx, contract, blockNumber)
}

func (c *clientInstance) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return c.ethClient.CallContract(ctx, call, blockNumber)
}

func (c *clientInstance) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return c.ethClient.PendingNonceAt(ctx, address)
}

func (c *clientInstance) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.ethClient.SuggestGasPrice(ctx)
}

func (c *clientInstance) BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error) {
	return c.ethClient.BalanceAt(ctx, address, block)
}
