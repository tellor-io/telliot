// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rpc

import (
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/telliot/pkg/contracts/master"
	"github.com/tellor-io/telliot/pkg/contracts/proxy"
	balancerpool "github.com/tellor-io/telliot/pkg/tracker/balancer/balancerpool"
	balancerToken "github.com/tellor-io/telliot/pkg/tracker/balancer/balancertoken"
	uniswapcontract "github.com/tellor-io/telliot/pkg/tracker/uniswap"

	"github.com/tellor-io/telliot/pkg/util"
)

var abiCodecLog = util.NewLogger("rpc", "ABICodec")

// ABICodec holds abi definitions for encoding/decoding contract methods and events.
type ABICodec struct {
	abiStruct abi.ABI
	methods   map[string]*abi.Method
	Events    map[string]*abi.Event
}

// BuildCodec constructs a merged abi structure representing all methods/events for Tellor tellor. This is primarily
// used for mock encoding/decoding parameters but could also be used for manual RPC operations that do not rely on geth's contract impl.
func BuildCodec() (*ABICodec, error) {
	all := []string{
		master.TellorDisputeABI,
		master.TellorLibraryABI,
		master.TellorGettersLibraryABI,
		master.TellorStakeABI,
		master.TellorTransferABI,
		proxy.TellorGettersABI,
		balancerpool.BalancerpoolABI,
		balancerToken.BalancertokenABI,
		uniswapcontract.IERC20ABI,
		uniswapcontract.IUniswapV2PairABI,
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
		abiCodecLog.Debug("Mapping method sig: %s to method: %s", sig, a.Name)
		methodMap[sig] = &abi.Method{Name: a.Name, Constant: a.Constant, Inputs: a.Inputs, Outputs: a.Outputs}
	}
	for _, e := range abiStruct.Events {
		sig := hexutil.Encode(e.ID.Bytes())
		abiCodecLog.Debug("Mapping event sig: %s to event %s", sig, e.Name)
		eventMap[sig] = &abi.Event{Name: e.Name, Anonymous: e.Anonymous, Inputs: e.Inputs}
	}

	return &ABICodec{abiStruct, methodMap, eventMap}, nil
}

// AllEventsthis lets you quickly find the type of each event. It is helpful for debugging.
func AllEvents() (map[[32]byte]abi.Event, error) {
	all := []string{
		master.TellorABI,
		master.TellorDisputeABI,
		master.TellorGettersLibraryABI,
		master.TellorLibraryABI,
		master.TellorStorageABI,
		master.TellorTransferABI,
		balancerpool.BalancerpoolABI,
		balancerToken.BalancertokenABI,
		uniswapcontract.IERC20ABI,
		uniswapcontract.IUniswapV2PairABI,
	}

	parsed := make([]interface{}, 0)
	for _, abi := range all {
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
	eventMap := make(map[[32]byte]abi.Event)
	for _, e := range abiStruct.Events {
		eventMap[e.ID] = e
	}

	return eventMap, nil
}
