package db

import "strings"

//BalanceKey is the key to store/lookup account balance
const (
	BalanceKey = "eth_balance"

	//CurrentChallengeKey DB key
	CurrentChallengeKey = "current_challenge"
	RequestIdKey        = "current_requestId"
	DifficultyKey       = "current_difficulty"
	QueryStringKey      = "current_query_string"
	GranularityKey      = "current_granularity"
	TotalTipKey         = "current_total_tip"
	MiningStatusKey     = "mining_status"
	PendingChallengeKey = "pending_challenge"

	//GasKey
	GasKey = "wei_gas_price"

	//Top 50
	Top50Key = "top_50_requestIds"

	//TributeBalance
	TributeBalanceKey = "trib_balance"

	//Dispute Status
	DisputeStatusKey = "dispute_status"

	//RequestID's are stored with this prefix and the id itself
	//e.g. "qm_2" represents request ID 2
	QueryMetadataPrefix = "qm_"

	//Request values are stored with this prefix plus request id
	QueriedValuePrefix = "qv_"
)

var knownKeys map[string]bool

func initKeyLook() {
	knownKeys = map[string]bool{
		BalanceKey:          true,
		CurrentChallengeKey: true,
		RequestIdKey:        true,
		DifficultyKey:       true,
		QueryStringKey:      true,
		GranularityKey:      true,
		TotalTipKey:         true,
		MiningStatusKey:     true,
		PendingChallengeKey: true,
		GasKey:              true,
		Top50Key:            true,
		TributeBalanceKey:   true,
		DisputeStatusKey:    true,
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
