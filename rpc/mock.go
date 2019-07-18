package rpc

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	balanceAtFN     = "0x70a08231"
	top50FN         = "0xb5413029"
	currentVarsFN   = "0xa22e407a"
	disputeStatusFN = "0x733bdef0"
)

//CurrentChallenge holds details about the current mining challenge
type CurrentChallenge struct {
	ChallengeHash [32]byte
	RequestID     *big.Int
	Difficulty    *big.Int
	QueryString   string
	Granularity   *big.Int
	Tip           *big.Int
}

//MockOptions are config options for the mock client
type MockOptions struct {
	ETHBalance       *big.Int
	Nonce            uint64
	GasPrice         *big.Int
	TokenBalance     *big.Int
	Top50Requests    []*big.Int
	CurrentChallenge *CurrentChallenge
	DisputeStatus    *big.Int
}

type mockClient struct {
	balance          *big.Int
	nonce            uint64
	gasPrice         *big.Int
	tokenBalance     *big.Int
	top50Requests    []*big.Int
	currentChallenge *CurrentChallenge
	disputeStatus    *big.Int
}

//NewMockClient returns instance of mock client
func NewMockClient() ETHClient {
	return &mockClient{}
}

//NewMockClientWithValues creates a mock client with default values to return for calls
func NewMockClientWithValues(opts *MockOptions) ETHClient {
	return &mockClient{balance: opts.ETHBalance, nonce: opts.Nonce,
		gasPrice: opts.GasPrice, tokenBalance: opts.TokenBalance,
		top50Requests: opts.Top50Requests, currentChallenge: opts.CurrentChallenge,
		disputeStatus: opts.DisputeStatus}
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
			return math.PaddedBigBytes(math.U256(c.tokenBalance), 32), nil
		}
	case top50FN:
		{
			log.Println("Getting top-50")
			b := new(bytes.Buffer)
			for _, t := range c.top50Requests {
				hex := math.PaddedBigBytes(math.U256(t), 32)
				b.Write(hex)
			}
			return b.Bytes(), nil
		}
	case currentVarsFN:
		{
			log.Println("Getting current variables")
			//bytes32, uint, uint,string memory,uint,uint
			//current challenge hash, curretnRequestId, level of difficulty, api/query string, and granularity(number of decimals requested
			b := new(bytes.Buffer)

			_, err := b.Write(c.currentChallenge.ChallengeHash[:])
			if err != nil {
				return nil, err
			}

			if err := paddedInt(b, c.currentChallenge.RequestID); err != nil {
				return nil, err
			}
			if err := paddedInt(b, c.currentChallenge.Difficulty); err != nil {
				return nil, err
			}
			asBytes := []byte(c.currentChallenge.QueryString)
			origLength := len(asBytes)
			diff := 32 - (len(asBytes) % 32)
			asBytes = common.RightPadBytes(asBytes, len(asBytes)+diff)
			//strings are dynamic types and therefore have an 'offset' in the position
			//where string resides in results. We will then write the length-prefixed string
			//at that position later
			if err := paddedInt(b, big.NewInt(192)); err != nil {
				return nil, err
			}

			if err := paddedInt(b, c.currentChallenge.Granularity); err != nil {
				return nil, err
			}
			if err := paddedInt(b, c.currentChallenge.Tip); err != nil {
				return nil, err
			}
			//now we write the length of the string
			if err := paddedInt(b, big.NewInt(int64(origLength))); err != nil {
				return nil, err
			}
			//now the bytes
			_, err = b.Write(asBytes)
			return b.Bytes(), nil

		}

	case disputeStatusFN:
		{
			b := new(bytes.Buffer)
			b.Write(math.PaddedBigBytes(math.U256(c.disputeStatus), 32))
			b.Write(math.PaddedBigBytes(math.U256(big.NewInt(time.Now().Unix())), 32))
			return b.Bytes(), nil
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

func paddedRLP(w *bytes.Buffer, val interface{}) error {
	b, err := rlp.EncodeToBytes(val)
	if err != nil {
		return err
	}
	log.Printf("Val: %v Enc: %v\n", val, b)
	_, err = w.Write(common.LeftPadBytes(b, 32))
	return err
}

func paddedInt(w *bytes.Buffer, val *big.Int) error {
	hex := math.PaddedBigBytes(math.U256(val), 32)
	_, err := w.Write(hex)
	return err
}
