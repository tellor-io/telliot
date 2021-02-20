// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import "strings"

const (
	// BalanceKey is the key to store/lookup account balance.
	BalanceKey = "eth_balance"

	// CurrentChallengeKey DB key.
	CurrentChallengeKey = "current_challenge"
	RequestIdKey        = "current_requestId"
	RequestIdKey0       = "current_requestId0"
	RequestIdKey1       = "current_requestId1"
	RequestIdKey2       = "current_requestId2"
	RequestIdKey3       = "current_requestId3"
	RequestIdKey4       = "current_requestId4"
	DifficultyKey       = "current_difficulty"
	QueryStringKey      = "current_query_string"
	GranularityKey      = "current_granularity"
	TotalTipKey         = "current_total_tip"
	MiningStatusKey     = "mining_status"

	GasKey   = "wei_gas_price"
	Top50Key = "top_50_requestIds"

	TributeBalanceKey = "trib_balance"
	DisputeStatusKey  = "dispute_status"

	// QueryMetadataPrefix is for RequestID's that are stored with this prefix and the id itself
	// e.g. "qm_2" represents request ID 2.
	QueryMetadataPrefix = "qm_"

	// QueriedValuePrefix is for request values that are stored with this prefix plus request id.
	QueriedValuePrefix = "qv_"
	LastNewValueKey    = "lastnewvalue"
)

var knownKeys map[string]bool

func initKeyLook() {
	knownKeys = map[string]bool{
		BalanceKey:          true,
		CurrentChallengeKey: true,
		RequestIdKey:        true,
		RequestIdKey0:       true,
		RequestIdKey1:       true,
		RequestIdKey2:       true,
		RequestIdKey3:       true,
		RequestIdKey4:       true,
		DifficultyKey:       true,
		QueryStringKey:      true,
		GranularityKey:      true,
		TotalTipKey:         true,
		MiningStatusKey:     true,
		GasKey:              true,
		Top50Key:            true,
		TributeBalanceKey:   true,
		DisputeStatusKey:    true,
		LastNewValueKey:     true,
	}
}
func isKnownKey(key string) bool {
	if knownKeys == nil {
		initKeyLook()
	}
	if !knownKeys[key] {
		if !strings.HasPrefix(key, QueryMetadataPrefix) &&
			!strings.HasPrefix(key, QueriedValuePrefix) {
			return false
		}
	}
	return true
}
