package db

//BalanceKey is the key to store/lookup account balance
const BalanceKey = "eth_balance"

//CurrentChallengeKey DB key
const CurrentChallengeKey = "current_challenge"
const RequestIdKey = "current_requestId"
const DifficultyKey = "current_difficulty"
const QueryStringKey = "current_query_string"
const GranularityKey = "current_granularity"
const TotalTipKey = "current_total_tip"

//Gas
const GasKey = "wei_gas_price"

//Top 50
const Top50Key = "top_50_requestIds"

//TributeBalance
const TributeBalanceKey = "trib_balance"

//Dispute Status
const DisputeStatusKey = "dispute_status"

//RequestID's are stored with this prefix and the id itself
//e.g. "qm_2" represents request ID 2
const QueryMetadataPrefix = "qm_"

//Request values are stored with this prefix plus request id
const QueriedValuePrefix = "qv_"
