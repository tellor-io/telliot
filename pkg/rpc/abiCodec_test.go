// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rpc

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/tellor-io/TellorMiner/pkg/config"
	"github.com/tellor-io/TellorMiner/pkg/testutil"
)

func TestABICodec(t *testing.T) {
	config.OpenTestConfig(t)
	codec, err := BuildCodec()
	if err != nil {
		testutil.Ok(t, err)
	}
	m := codec.methods["0xe1eee6d6"]
	if m == nil {
		testutil.Ok(t, errors.New("Missing expected method matching test sig"))
	} else if m.Name != "getRequestVars" {
		testutil.Ok(t, errors.New(fmt.Sprintf("Method name is unexpected. %s != getRequestVars", m.Name)))
	}

	// string, string, bytes32,  uint, uint, uint.
	var hash [32]byte
	copy([]byte("12345"), hash[:])
	data, err := m.Outputs.Pack("someQueryString", "ETH/USD", hash, big.NewInt(1000), big.NewInt(0), big.NewInt(0))
	if err != nil {
		testutil.Ok(t, err)
	}

	for i := 0; i < len(data); i += 32 {
		hex := hexutil.Encode(data[i : i+32])
		fmt.Println(hex)
	}
}
