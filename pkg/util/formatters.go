// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package util

import (
	"encoding/json"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func FormatERC20Balance(balance *big.Int) string {
	f := 0.0
	if balance != nil {
		divisor := big.NewInt(1e18 / 100)
		divisor.Div(balance, divisor)
		f = float64(divisor.Uint64()) / 100
	}
	return strconv.FormatFloat(f, 'f', 2, 64)
}

// Unfortunate hack to enable json parsing of human readable time strings
// see https://github.com/golang/go/issues/10275
// code from https://stackoverflow.com/questions/48050945/how-to-unmarshal-json-into-durations.
type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.String() + "\""), nil
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value * float64(time.Second))
		return nil
	case string:
		dur, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		d.Duration = dur
		return nil
	default:
		return errors.Errorf("invalid duration")
	}
}

func SanitizeMetricName(input string) string {
	return strings.ReplaceAll(input, "/", "_")
}
