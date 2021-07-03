// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package math

func PercentageDiff(old, new float64) (delta float64) {
	diff := float64(new - old)

	if old > new {
		return (diff / float64(old)) * 100
	}
	return (diff / float64(new)) * 100
}
