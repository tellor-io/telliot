// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/util"
)

const (
	balanceAtFN    = "0x70a08231"
	top50FN        = "0xb5413029"
	currentVarsFN  = "0xa22e407a"
	getRequestVars = "0xe1eee6d6"
	didMineFN      = "0x63bb82ad"
	getUintVarFN   = "0x612c8f7f"
	decimalsFN     = "0x313ce567"
	// Balancerpool funcs.
	getCurentTokensFN = "0xcc77828d"
	getSpotPriceFN    = "0x15e84af9"
	// Balancertoken funcs.
	symbolFN = "0x95d89b41"
	// Uniswap pair funcs.
	getReservesFN = "0x0902f1ac"
	// Uniswap erc20 token funcs.
	token0FN = "0x0dfe1681"
	token1FN = "0xd21220a7"
)

// CurrentChallenge holds details about the current mining challenge.
type CurrentChallenge struct {
	ChallengeHash [32]byte
	RequestID     *big.Int
	Difficulty    *big.Int
	QueryString   string
	Granularity   *big.Int
	Tip           *big.Int
}

// CurrentReserves holds details about the current reserves on the uniswap pair contract.
type CurrentReserves struct {
	// Reserve0 is the amount of Token0 liquidity in the pool.
	Reserve0 *big.Int
	// Reserve1 is the amount of Token1 liquidity in the pool.
	Reserve1           *big.Int
	BlockTimestampLast uint32
}

// MockQueryMeta is hardcoded query metadata to use for testing.
type MockQueryMeta struct {
	QueryString string
	Granularity int
}

// MockOptions are config options for the mock client.
type MockOptions struct {
	ETHBalance       *big.Int
	MiningStatus     bool
	Nonce            uint64
	GasPrice         *big.Int
	TokenBalance     *big.Int
	Top50Requests    []*big.Int
	CurrentChallenge *CurrentChallenge
	QueryMetadata    map[uint]*MockQueryMeta

	// Balancer related.
	BPoolContractAddress common.Address
	BPoolCurrentTokens   []common.Address
	BPoolSpotPrice       *big.Int

	// Uniswap related.
	UniPairContractAddress common.Address
	UniReserves            *CurrentReserves
	UniToken0              common.Address
	UniToken1              common.Address

	// Decimals values for Uniswap, Balancer based on contract addresses.
	Decimals map[string]int
	// Token symbol map for Uniswap, Balancer based on contract addresses.
	TokenSymbols map[string]string
}

type mockClient struct {
	balance          *big.Int
	nonce            uint64
	miningStatus     bool
	gasPrice         *big.Int
	tokenBalance     *big.Int
	top50Requests    []*big.Int
	currentChallenge *CurrentChallenge
	logger           log.Logger

	mockQueryMeta map[uint]*MockQueryMeta

	// Balancer related.
	bPoolContractAddress common.Address
	bPoolCurrentTokens   []common.Address
	bPoolSpotPrice       *big.Int

	// Uniswap related.
	uniPairContractAddress common.Address
	uniReserves            *CurrentReserves
	uniToken0              common.Address
	uniToken1              common.Address

	// Decimals values for Uniswap, Balancer based on contract addresses.
	decimals map[string]int
	// Token symbol map for Uniswap, Balancer based on contract addresses.
	tokenSymbols map[string]string
	abiCodec     *ABICodec
}

type mockError struct {
	codeVal int
}

func (e *mockError) Error() string {
	return fmt.Sprintf("error value: %d",
		e.codeVal)
}

// NewMockClient returns instance of mock client.
func NewMockClient() contracts.ETHClient {
	return &mockClient{
		logger: log.With(logging.NewLogger(), "component", ComponentName),
	}
}

// NewMockClientWithValues creates a mock client with default values to return for calls.
func NewMockClientWithValues(opts *MockOptions) contracts.ETHClient {
	codec, err := BuildCodec(logging.NewLogger())
	if err != nil {
		panic(err)
	}

	logger := logging.NewLogger()
	level.Info(logger).Log("msg", "check mining status", "status", opts.MiningStatus)
	return &mockClient{
		balance:                opts.ETHBalance,
		miningStatus:           opts.MiningStatus,
		nonce:                  opts.Nonce,
		gasPrice:               opts.GasPrice,
		tokenBalance:           opts.TokenBalance,
		top50Requests:          opts.Top50Requests,
		currentChallenge:       opts.CurrentChallenge,
		mockQueryMeta:          opts.QueryMetadata,
		bPoolContractAddress:   opts.BPoolContractAddress,
		bPoolCurrentTokens:     opts.BPoolCurrentTokens,
		bPoolSpotPrice:         opts.BPoolSpotPrice,
		tokenSymbols:           opts.TokenSymbols,
		uniPairContractAddress: opts.UniPairContractAddress,
		uniReserves:            opts.UniReserves,
		uniToken0:              opts.UniToken0,
		uniToken1:              opts.UniToken1,
		decimals:               opts.Decimals,
		abiCodec:               codec,
		logger:                 log.With(logger, "component", ComponentName),
	}
}

func (c *mockClient) SetTokenBalance(bal *big.Int) {
	c.tokenBalance = bal
}

func (c *mockClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	return nil, nil
}

func (c *mockClient) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return nil, true, nil
}

func (c *mockClient) Close() {
	level.Info(c.logger).Log("msg", "closing mock client")
}

func (c *mockClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return []byte("1234567890"), nil
}

func (c *mockClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return &types.Receipt{Status: 1}, nil
}
func (c *mockClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return nil, nil
}

func (c *mockClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	fn := hexutil.Encode(call.Data[0:4])
	meth := c.abiCodec.methods[fn]
	if meth == nil {
		return []byte{}, errors.Errorf("unknown function signature:%v", fn)
	}

	switch fn {
	case balanceAtFN:
		{
			return meth.Outputs.Pack(c.tokenBalance)
			//level.Debug(c.logger).Log("msg", "getting balance from contract")
			//return math.PaddedBigBytes(math.U256(c.tokenBalance), 32), nil
		}
	case didMineFN:
		{
			level.Info(c.logger).Log("msg", "getting mining status", "status", c.miningStatus)
			return meth.Outputs.Pack(c.miningStatus)
		}
	case top50FN:
		{
			return meth.Outputs.Pack(c.top50Requests)
			/*
				level.Debug(c.logger).Log("msg","getting top-50")
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
				level.Debug(c.logger).Log("msg", "getting current variables")
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
				level.Warn(c.logger).Log("msg", "no params found", "reqId", reqID)
				return []byte{}, nil
			}
			level.Debug(c.logger).Log(
				"msg", "using mock",
				"params", fmt.Sprintf("%+v", params),
			)
			return meth.Outputs.Pack(params.QueryString, "", [32]byte{}, big.NewInt(int64(params.Granularity)), big.NewInt(1), big.NewInt(0))

			/*
				level.Debug(c.logger).Log("msg","getting request vars")

				hex := hexutil.Encode(reqIDData)
				level.Debug(c.logger).Log("msg", "encoded call params", "data", hex)
				reqIDNum, err := hexutil.DecodeBig(hex)
				if err != nil {
					return nil, err
				}

				reqID := reqIDNum.Uint64()
				mockClientLog.Debug("Mocking response for request id: %d\n", reqID)
				level.Debug(c.logger).Log(
					   "msg", "mocking response for request",
					   "reqID", reqID,
				)
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
				level.Debug(c.logger).Log(
					   "msg","encoded request",
					   "vars", fmt.Sprint("%v\n", b.Bytes()),
				)
				return b.Bytes(), nil
			*/
		}
	// Balancer related.
	case getUintVarFN:
		{
			// Return 10 minutes ago time for the _TIME_OF_LAST_NEW_VALUE key.
			inputdata := call.Data[4:]
			timeOfLastValueData := util.Keccak256([]byte("_TIME_OF_LAST_NEW_VALUE"))
			if bytes.Equal(inputdata, timeOfLastValueData[:]) {
				return meth.Outputs.Pack(big.NewInt(time.Now().Unix() - 10*60))
			}
			return meth.Outputs.Pack(big.NewInt(1))
		}
	case getCurentTokensFN:
		{
			return meth.Outputs.Pack(c.bPoolCurrentTokens)
		}
	case getSpotPriceFN:
		{
			return meth.Outputs.Pack(c.bPoolSpotPrice)
		}
	case symbolFN:
		{
			return meth.Outputs.Pack(c.tokenSymbols[call.To.Hex()])
		}
	// Uniswap related.
	case getReservesFN:
		{
			return meth.Outputs.Pack(c.uniReserves.Reserve0, c.uniReserves.Reserve1, c.uniReserves.BlockTimestampLast)
		}
	case token0FN:
		{
			return meth.Outputs.Pack(c.uniToken0)
		}
	case token1FN:
		{
			return meth.Outputs.Pack(c.uniToken1)
		}
	// Handle "decimals" func for different contracts.
	case decimalsFN:
		outValue := c.decimals[call.To.Hex()]
		switch meth.Outputs[0].Type.String() {
		case "uint8":
			return meth.Outputs.Pack(uint8(outValue))
		default:
			return meth.Outputs.Pack(big.NewInt(int64(outValue)))
		}
	}

	level.Warn(c.logger).Log("msg", "call unhandled", "fn", fn)
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
	var logs []types.Log

	abi, _ := abi.JSON(strings.NewReader(contracts.ITellorABI))
	ev, ok := abi.Events["NonceSubmitted"]
	if !ok {
		return logs, errors.New("NonceSubmitted event not foind in the ABI")
	}
	event := contracts.TellorNonceSubmitted{
		Miner:            common.Address{0},
		Nonce:            "0",
		RequestId:        [5]*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(1)},
		Value:            [5]*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(1)},
		CurrentChallenge: [32]byte{0},
	}
	test, err := ev.Inputs.NonIndexed().Pack(event.Nonce, event.RequestId, event.Value)
	if err != nil {
		return logs, err
	}

	log := types.Log{
		Address:     common.Address{0},
		Topics:      []common.Hash{ev.ID, common.BigToHash(common.Big0), common.BigToHash(common.Big1)},
		Data:        test,
		BlockNumber: 9,
	}
	logs = append(logs, log)

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
		Number:     math.BigPow(1, 0),
		GasLimit:   12345678,
		GasUsed:    1476322,
		Extra:      []byte("coolest block on chain"),
	}
	header.Time = uint64(time.Now().Unix())
	return &header, nil
}

func (c *mockClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return nil, nil
}

func (c *mockClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return nil, nil
}

// ABICodec holds abi definitions for encoding/decoding contract methods and events.
type ABICodec struct {
	abiStruct abi.ABI
	methods   map[string]*abi.Method
	Events    map[string]*abi.Event
}

// BuildCodec constructs a merged abi structure representing all methods/events for Tellor tellor. This is primarily
// used for mock encoding/decoding parameters but could also be used for manual RPC operations that do not rely on geth's contract impl.
func BuildCodec(logger log.Logger) (*ABICodec, error) {
	all := []string{
		contracts.ITellorABI,
		contracts.BPoolABI,
		contracts.BTokenABI,
		contracts.IERC20ABI,
		contracts.IUniswapV2PairABI,
	}

	parsed := make([]interface{}, 0)
	for _, abi := range all {
		var f interface{}
		if err := json.Unmarshal([]byte(abi), &f); err != nil {
			return nil, err
		}
		asList := f.([]interface{})
		parsed = append(parsed, asList...)
	}
	j, err := json.Marshal(parsed)
	if err != nil {
		return nil, err
	}
	abiStruct, err := abi.JSON(strings.NewReader(string(j)))
	if err != nil {
		return nil, err
	}
	methodMap := make(map[string]*abi.Method)
	eventMap := make(map[string]*abi.Event)
	for _, a := range abiStruct.Methods {
		sig := hexutil.Encode(a.ID)
		methodMap[sig] = &abi.Method{Name: a.Name, Constant: a.Constant, Inputs: a.Inputs, Outputs: a.Outputs}
	}
	for _, e := range abiStruct.Events {
		sig := hexutil.Encode(e.ID.Bytes())
		eventMap[sig] = &abi.Event{Name: e.Name, Anonymous: e.Anonymous, Inputs: e.Inputs}
	}

	return &ABICodec{abiStruct, methodMap, eventMap}, nil
}
