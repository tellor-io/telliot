// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ethereum

import (
	"testing"

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
}
