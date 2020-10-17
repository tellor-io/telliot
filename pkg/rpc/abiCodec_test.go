// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rpc

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tellor-io/TellorMiner/pkg/config"
)

func TestABICodec(t *testing.T) {
	config.OpenTestConfig(t)
	codec, err := BuildCodec()
	if err != nil {
		t.Fatal(err)
	}
	m := codec.methods["0xe1eee6d6"]
	if m == nil {
		t.Fatal("Missing expected method matching test sig")
	}
	if m.Name != "getRequestVars" {
		t.Fatalf("Method name is unexpected. %s != getRequestVars", m.Name)
	}

	// string, string, bytes32,  uint, uint, uint.
	var hash [32]byte
	copy([]byte("12345"), hash[:])
	data, err := m.Outputs.Pack("someQueryString", "ETH/USD", hash, big.NewInt(1000), big.NewInt(0), big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(data); i += 32 {
		hex := hexutil.Encode(data[i : i+32])
		fmt.Println(hex)
	}
}
