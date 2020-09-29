// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package pow

import (
	"encoding/hex"
)

func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}
