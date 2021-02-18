// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// BalanceKey is the key to store/lookup account balance.
	BalancePrefix = "eth_balance_"

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

	TributeBalancePrefix = "trib_balance_"
	DisputeStatusPrefix  = "dispute_status_"

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
		BalancePrefix:        true,
		CurrentChallengeKey:  true,
		RequestIdKey:         true,
		RequestIdKey0:        true,
		RequestIdKey1:        true,
		RequestIdKey2:        true,
		RequestIdKey3:        true,
		RequestIdKey4:        true,
		DifficultyKey:        true,
		QueryStringKey:       true,
		GranularityKey:       true,
		TotalTipKey:          true,
		MiningStatusKey:      true,
		GasKey:               true,
		Top50Key:             true,
		TributeBalancePrefix: true,
		DisputeStatusPrefix:  true,
		LastNewValueKey:      true,
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

func DisputeStatusKeyFor(address common.Address) string {
	return fmt.Sprintf("%s%s", DisputeStatusPrefix, strings.ToLower(address.String()))
}

func TributeBalanceKeyFor(address common.Address) string {
	return fmt.Sprintf("%s%s", TributeBalancePrefix, strings.ToLower(address.String()))
}
