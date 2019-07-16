package rpc

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tellor-io/TellorMiner/config"
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
}

//clientInstance is the concrete implementation of the ETHClient
type clientInstance struct {
	ethClient *ethclient.Client
	timeout   time.Duration
}

var (
	//retry delays that range from 100ms to 2mins
	backoff = []uint64{100, 500, 1000, 2000, 5000, 10000, 15000, 30000, 60000, 120000}

	//rate to print errors if continue to occur in retry loop
	errorPrintTick = time.Duration(5000)
)

//NewClient creates a new client instance
func NewClient(url string) (ETHClient, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	timeout := time.Duration(time.Duration(cfg.EthClientTimeout) * time.Second)
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &clientInstance{ethClient: client, timeout: timeout}, nil
}

func (c *clientInstance) withTimeout(ctx context.Context, fn func(*context.Context) error) error {
	wTo, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	tryCount := 0
	nextTick := time.Now().Add(errorPrintTick)
	for {
		err := fn(&wTo)
		if err == nil {
			return nil
		}
		//pause for a bit and try again
		sleepTime := backoff[tryCount%len(backoff)]
		tryCount++
		if time.Now().After(nextTick) {
			fmt.Printf("Problem calling ethClient: %v\n", err)
			nextTick = time.Now().Add(errorPrintTick)
		}

		time.Sleep(time.Duration(sleepTime))
		dl, _ := wTo.Deadline()
		if dl.Before(time.Now()) {
			return err
		}
	}
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

func (c *clientInstance) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	var res []byte
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.CodeAt(*_ctx, contract, blockNumber)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var res []byte
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.CallContract(*_ctx, call, blockNumber)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	var res uint64
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.PendingNonceAt(*_ctx, address)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var res *big.Int
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.SuggestGasPrice(*_ctx)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	var res uint64
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.EstimateGas(*_ctx, call)
		res = r
		return e
	})
	return res, _err
}

func (c *clientInstance) BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error) {
	var res *big.Int
	_err := c.withTimeout(ctx, func(_ctx *context.Context) error {
		r, e := c.ethClient.BalanceAt(*_ctx, address, block)
		res = r
		return e
	})
	return res, _err
}
