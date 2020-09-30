package rpc

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/tellor-io/TellorMiner/contracts1"
	"github.com/tellor-io/TellorMiner/util"

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
	getRequestVars  = "0xe1eee6d6"
	didMineFN		= "0x63bb82ad"
)

var mockClientLog = util.NewLogger("rpc", "mockClient")

//CurrentChallenge holds details about the current mining challenge
type CurrentChallenge struct {
	ChallengeHash [32]byte
	RequestID     *big.Int
	Difficulty    *big.Int
	QueryString   string
	Granularity   *big.Int
	Tip           *big.Int
}

//MockQueryMeta is hardcoded query metadata to use for testing
type MockQueryMeta struct {
	QueryString string
	Granularity int
}

//MockOptions are config options for the mock client
type MockOptions struct {
	ETHBalance       *big.Int
	MiningStatus	 bool
	Nonce            uint64
	GasPrice         *big.Int
	TokenBalance     *big.Int
	Top50Requests    []*big.Int
	CurrentChallenge *CurrentChallenge
	DisputeStatus    *big.Int
	QueryMetadata    map[uint]*MockQueryMeta
}

type mockClient struct {
	balance          *big.Int
	nonce            uint64
	miningStatus	  bool
	gasPrice         *big.Int
	tokenBalance     *big.Int
	top50Requests    []*big.Int
	currentChallenge *CurrentChallenge
	disputeStatus    *big.Int
	mockQueryMeta    map[uint]*MockQueryMeta
	abiCodec         *ABICodec
}

type mockError struct {
	codeVal int
}

func (e *mockError) Error() string {
	return fmt.Sprintf("error value: %d",
		e.codeVal)
}

//NewMockClient returns instance of mock client
func NewMockClient() ETHClient {
	return &mockClient{}
}

//NewMockClientWithValues creates a mock client with default values to return for calls
func NewMockClientWithValues(opts *MockOptions) ETHClient {
	codec, err := BuildCodec()
	if err != nil {
		panic(err)
	}
	fmt.Println("Mining Status",opts.MiningStatus)
	return &mockClient{balance: opts.ETHBalance, miningStatus:opts.MiningStatus ,nonce: opts.Nonce,
		gasPrice: opts.GasPrice, tokenBalance: opts.TokenBalance,
		top50Requests: opts.Top50Requests, currentChallenge: opts.CurrentChallenge,
		disputeStatus: opts.DisputeStatus, mockQueryMeta: opts.QueryMetadata, abiCodec: codec}
}

func (c *mockClient) SetTokenBalance(bal *big.Int) {
	c.tokenBalance = bal
}

func (c *mockClient) Close() {
	mockClientLog.Info("Closing mock client")
}

func (c *mockClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return []byte("1234567890"), nil
}
func (c *mockClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return nil, nil
}
func (c *mockClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	fn := hexutil.Encode(call.Data[0:4])
	meth := c.abiCodec.methods[fn]
	if meth == nil {
		mockClientLog.Error("Unknown function with sig: %s\n", fn)
		return []byte{}, nil
	}

	switch fn {
	case balanceAtFN:
		{
			return meth.Outputs.Pack(c.tokenBalance)
			//mockClientLog.Debug("Getting balance from contract")
			//return math.PaddedBigBytes(math.U256(c.tokenBalance), 32), nil
		}
	case didMineFN:
		{
		    fmt.Println("getting Mining Status",c.miningStatus)
			return meth.Outputs.Pack(c.miningStatus)
		}
	case top50FN:
		{
			return meth.Outputs.Pack(c.top50Requests)
			/*
				mockClientLog.Debug("Getting top-50")
				b := new(bytes.Buffer)
				for _, t := range c.top50Requests {
					hex := math.PaddedBigBytes(math.U256(t), 32)
					b.Write(hex)
				}
				return b.Bytes(), nil
			*/
		}
	case currentVarsFN:
		{
			return meth.Outputs.Pack(c.currentChallenge.ChallengeHash,
				c.currentChallenge.RequestID,
				c.currentChallenge.Difficulty,
				c.currentChallenge.QueryString,
				c.currentChallenge.Granularity,
				c.currentChallenge.Tip)
			/*
				mockClientLog.Debug("Getting current variables")
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
			*/

		}

	case disputeStatusFN:
		{
			return meth.Outputs.Pack(c.disputeStatus, big.NewInt(time.Now().Unix()))
			/*
				b := new(bytes.Buffer)
				b.Write(math.PaddedBigBytes(math.U256(c.disputeStatus), 32))
				b.Write(math.PaddedBigBytes(math.U256(, 32))
				return b.Bytes(), nil
			*/
		}

	case getRequestVars:
		{
			reqIDData := call.Data[4:]
			vals, err := meth.Inputs.UnpackValues(reqIDData)
			if err != nil {
				return nil, err
			}
			reqID := vals[0].(*big.Int)
			params := c.mockQueryMeta[uint(reqID.Uint64())]
			if params == nil {
				mockClientLog.Warn("No params found with id: %v\n", reqID)
				return []byte{}, nil
			}
			mockClientLog.Debug("Using mock params: %+v\n", params)
			return meth.Outputs.Pack(params.QueryString, "", [32]byte{}, big.NewInt(int64(params.Granularity)), big.NewInt(1), big.NewInt(0))

			/*
				mockClientLog.Debug("Getting request vars")

				hex := hexutil.Encode(reqIDData)
				mockClientLog.Debug("Encoded call params data: %s\n", hex)
				reqIDNum, err := hexutil.DecodeBig(hex)
				if err != nil {
					return nil, err
				}

				reqID := reqIDNum.Uint64()
				mockClientLog.Debug("Mocking response for request id: %d\n", reqID)
				b := new(bytes.Buffer)
				mockParams := c.mockQueryMeta[uint(reqID)]
				if mockParams == nil {
					return b.Bytes(), nil
				}
				qsBytes := []byte(mockParams.QueryString)
				origLength := len(qsBytes)
				diff := 32 - (len(qsBytes) % 32)
				qsBytes = common.RightPadBytes(qsBytes, len(qsBytes)+diff)

				//first return variable is a string, but we just insert offset position, which
				//comes AFTER all the padded number vars
				if err := paddedInt(b, big.NewInt(128)); err != nil {
					return nil, err
				}
				//second return variable is another string, but it will come AFTER the first
				//string (i.e. after its padded length and character bytes) and all padded number vars
				if err := paddedInt(b, big.NewInt(int64(128+origLength+32))); err != nil {
					return nil, err
				}

				//now the hash
				if err := paddedInt(b, big.NewInt(0)); err != nil {
					return nil, err
				}

				//now the granularity
				if err := paddedInt(b, big.NewInt(int64(mockParams.Granularity))); err != nil {
					return nil, err
				}

				//now fake index
				if err := paddedInt(b, big.NewInt(0)); err != nil {
					return nil, err
				}

				//and tip
				if err := paddedInt(b, big.NewInt(0)); err != nil {
					return nil, err
				}

				//NOW we right our query string length
				if err := paddedInt(b, big.NewInt(int64(origLength))); err != nil {
					return nil, err
				}
				//then bytes
				_, err = b.Write(qsBytes)
				if err != nil {
					return nil, err
				}

				//Finally, the length of symbol (0)
				if err := paddedInt(b, big.NewInt(0)); err != nil {
					return nil, err
				}
				mockClientLog.Debug("Encoded request vars: %v\n", b.Bytes())
				return b.Bytes(), nil
			*/
		}
	}

	mockClientLog.Warn("Call unhandled Fn: %s\n", fn)
	return []byte{}, nil
}

func (c *mockClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return c.nonce, nil
}

func (c *mockClient) NonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return c.nonce, nil
}

func (c *mockClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return 0, nil
}

func (c *mockClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.gasPrice, nil
}

func (c *mockClient) BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error) {
	if c.balance.Cmp(big.NewInt(0)) < 0 {
		return nil, &mockError{1}
	}
	return c.balance, nil
}

func (c *mockClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	tokenAbi, _ := abi.JSON(strings.NewReader(contracts1.TellorLibraryABI))
	ev := tokenAbi.Events["NonceSubmitted"]
	event1 := contracts1.TellorLibraryNonceSubmitted{
		Miner: common.Address{0}, 
		Nonce: "0", 
		RequestId: big.NewInt(1),
		Value: big.NewInt(1), 
		CurrentChallenge: [32]byte{0},
		}
	//eventResult := contracts1.TellorLibraryNonceSubmitted{}
	test, _ := ev.Inputs.NonIndexed().Pack(event1.Nonce, event1.Value, event1.CurrentChallenge)
	//test, err := ev.Inputs.Pack(event1.Miner, event1.Nonce, event1.RequestId, event1.Value, event1.CurrentChallenge)
	//fmt.Print("\nresult: ", test," Error: ",  err, "\n", "test: ", common.BigToHash(common.Big1))

	log := types.Log{
		Address: common.Address{0},
		Topics: []common.Hash{ev.ID(), common.BigToHash(common.Big0), common.BigToHash(common.Big1)},
		Data: test,
		BlockNumber: 9,
	}

	var logs []types.Log

	logs = append(logs,log)

	return logs, nil
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

func (c *mockClient) IsSyncing(ctx context.Context) (bool, error) {
	return false, nil
}

func (c *mockClient) NetworkID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1), nil
}

func (c *mockClient) HeaderByNumber(ctx context.Context, num *big.Int) (*types.Header, error) {
	header := types.Header{
		Difficulty: math.BigPow(11, 11),
		Number:     math.BigPow(1,0),
		GasLimit:   12345678,
		GasUsed:    1476322,
		Extra:      []byte("coolest block on chain"),
	}
	header.Time = uint64(time.Now().Unix())
	return &header, nil
}

func paddedRLP(w *bytes.Buffer, val interface{}) error {
	b, err := rlp.EncodeToBytes(val)
	if err != nil {
		return err
	}
	_, err = w.Write(common.LeftPadBytes(b, 32))
	return err
}

func paddedInt(w *bytes.Buffer, val *big.Int) error {
	hex := math.PaddedBigBytes(math.U256(val), 32)
	_, err := w.Write(hex)
	return err
}
