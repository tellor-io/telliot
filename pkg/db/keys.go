// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"github.com/prometheus/prometheus/pkg/labels"
)

var GasPriceLabel = labels.Label{Name: "__name__", Value: "gas_price"}

const (
	// CurrentChallengeKey DB key.
	RequestIdKey    = "current_requestId"
	RequestIdKey0   = "current_requestId0"
	RequestIdKey1   = "current_requestId1"
	RequestIdKey2   = "current_requestId2"
	RequestIdKey3   = "current_requestId3"
	RequestIdKey4   = "current_requestId4"
	DifficultyKey   = "current_difficulty"
	QueryStringKey  = "current_query_string"
	GranularityKey  = "current_granularity"
	TotalTipKey     = "current_total_tip"
	MiningStatusKey = "mining_status"

	// QueriedValuePrefix is for request values that are stored with this prefix plus request id.
	QueriedValuePrefix = "qv_"
)

var knownKeys map[string]bool

func initKeyLook() {
	knownKeys = map[string]bool{
		RequestIdKey:    true,
		RequestIdKey0:   true,
		RequestIdKey1:   true,
		RequestIdKey2:   true,
		RequestIdKey3:   true,
		RequestIdKey4:   true,
		DifficultyKey:   true,
		QueryStringKey:  true,
		GranularityKey:  true,
		TotalTipKey:     true,
		MiningStatusKey: true,
	}
}
func isKnownKey(key string) bool {
	if knownKeys == nil {
		initKeyLook()
	}
	if !knownKeys[key] {
		return false
	}
	return true
}
