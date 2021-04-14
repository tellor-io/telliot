// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package rpc

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/tellor-io/telliot/pkg/logging"
	"github.com/tellor-io/telliot/pkg/testutil"
)

func TestABICodec(t *testing.T) {
	codec, err := BuildCodec(logging.NewLogger())
	if err != nil {
		testutil.Ok(t, err)
	}
	m := codec.methods[getRequestVars]
	if m == nil {
		testutil.Ok(t, errors.New("Missing expected method matching test sig"))
	} else if m.Name != "getRequestVars" {
		testutil.Ok(t, errors.Errorf("Method name is unexpected. %s != getRequestVars", m.Name))
	}

	data, err := m.Outputs.Pack(big.NewInt(0), big.NewInt(0))
	testutil.Ok(t, err)

	for i := 0; i < len(data); i += 32 {
		hex := hexutil.Encode(data[i : i+32])
		fmt.Println(hex)
	}
}
