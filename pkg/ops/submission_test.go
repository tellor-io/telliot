// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package ops

import (
	"testing"

	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/testutil"
)

// TestMain check for goroutine leaks.
// It ensures that at the end of the tests there are no remaining go routines.
func TestMain(m *testing.M) {
	testutil.TolerantVerifyLeakMain(m)
}

func TestSubmission(t *testing.T) {
	err := contracts.GetTellorCore()
	testutil.Ok(t, err)
}
