package db

//BalanceKey is the key to store/lookup account balance
const BalanceKey = "eth_balance"

//Current Variables
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

//RequestID's are stored as the string of the number and the normalized timestamp
//e.g. "2" is the key for request ID 2
