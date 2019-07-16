package rpc

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	balanceAtFN = "0x70a08231"
)

//MockOptions are config options for the mock client
type MockOptions struct {
	ETHBalance   *big.Int
	Nonce        uint64
	GasPrice     *big.Int
	TokenBalance *big.Int
}

type mockClient struct {
	balance      *big.Int
	nonce        uint64
	gasPrice     *big.Int
	tokenBalance *big.Int
}

//NewMockClient returns instance of mock client
func NewMockClient() ETHClient {
	return &mockClient{}
}

//NewMockClientWithValues creates a mock client with default values to return for calls
func NewMockClientWithValues(opts *MockOptions) ETHClient {
	return &mockClient{balance: opts.ETHBalance, nonce: opts.Nonce, gasPrice: opts.GasPrice, tokenBalance: opts.TokenBalance}
}

func (c *mockClient) SetTokenBalance(bal *big.Int) {
	c.tokenBalance = bal
}

func (c *mockClient) Close() {
	fmt.Println("Closing mock client")
}

func (c *mockClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return []byte("1234567890"), nil
}
func (c *mockClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return nil, nil
}
func (c *mockClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	fn := hexutil.Encode(call.Data[0:4])
	switch fn {
	case balanceAtFN:
		{
			log.Println("Getting balance from contract")
			return math.PaddedBigBytes(c.tokenBalance, 32), nil
		}
	}
	log.Printf("Call unhandled Fn: %s\n", fn)
	return []byte{}, nil
}

func (c *mockClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return c.nonce, nil
}

func (c *mockClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return 0, nil
}

func (c *mockClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.gasPrice, nil
}

func (c *mockClient) BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error) {
	return c.balance, nil
}

func (c *mockClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}
func (c *mockClient) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}

func (c *mockClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return nil
}

func (c *mockClient) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	return nil, nil
}
