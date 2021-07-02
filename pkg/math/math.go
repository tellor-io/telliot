// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package math

import "math"

func PercentageChange(old, new int64) (delta float64) {
	diff := float64(new - old)
	return math.Abs((diff / float64(old)) * 100)
}
