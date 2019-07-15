package rpc

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type mockClient struct {
	balance  *big.Int
	nonce    uint64
	gasPrice *big.Int
}

//NewMockClient returns instance of mock client
func NewMockClient() ETHClient {
	return &mockClient{}
}

//NewMockClientWithValues creates a mock client with default values to return for calls
func NewMockClientWithValues(balance *big.Int, nonce uint64, gasPrice *big.Int) ETHClient {
	return &mockClient{balance, nonce, gasPrice}
}

func (c *mockClient) Close() {
	fmt.Println("Closing mock client")
}

func (c *mockClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return []byte{}, nil
}

func (c *mockClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return []byte{}, nil
}

func (c *mockClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return c.nonce, nil
}

func (c *mockClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.gasPrice, nil
}

func (c *mockClient) BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error) {
	return c.balance, nil
}
