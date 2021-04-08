// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package common

const (
	// GWEI constant is the multiplier from Wei.
	GWEI = 1e9
	// WEI constant is smallest unit of gas.
	WEI = GWEI * 1e9

	// PriceTXs is the key used to save transactions cost
	// These are used to calculate the profitability when submitting a solution.
	PriceTXs = "PriceTXSlot"
)
