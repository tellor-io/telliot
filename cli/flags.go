package cli

import (
	"flag"
)

const configPath = "config"
const loggingConfigPath = "logConfig"
const psrPath = "psrPath"
const minerArg = "miner"
const dataServerArg = "dataServer"
const toAddressArg = "to"
const amountArg = "amount"
const transferArg = "transfer"
const depositArg = "deposit"
const approveArg = "approve"
const disputeArg = "dispute"
const requestIdArg = "requestId"
const timestampArg = "timestamp"
const minerIndexArg = "minerIndex"
const requestStakingWithdrawArg = "requestStakingWithdraw"
const withdrawStakeArg = "withdrawStake"
const voteArg = "vote"
const disputeIdArg = "disputeId"
const supportsDisputeArg = "supportsDispute"

//Flags holds all command line options
type Flags struct {
	ConfigPath             string
	LoggingConfigPath      string
	PSRPath                string
	Miner                  bool
	DataServer             bool
	Transfer               bool
	ToAddress              string
	Amount                 string
	Deposit                bool
	Approve                bool
	Dispute                bool
	RequestId              string
	Timestamp              string
	MinerIndex             string
	RequestStakingWithdraw bool
	WithdrawStake          bool
	Vote                   bool
	DisputeId              string
	SupportsDispute        bool
}

var (
	sharedFlags *Flags
)

func init() {
	GetFlags()
}

//GetFlags parses command line entries into a shared structure with values
func GetFlags() *Flags {
	if sharedFlags == nil {
		f := &Flags{}
		// home, e := os.UserHomeDir()
		// if e != nil {
		// 	log.Fatal("Could not get the base file path for app", e)
		// }
		// logConfigPath := filepath.Join(home, "LoggingConfig.json")
		// psrConfigPath := filepath.Join(home, "psr.json")
		// cfgPath := filepath.Join(home, "config.json")
		logConfigPath := "./loggingConfig.json"
		psrConfigPath := "./psr.json"
		cfgPath := "./config.json"

		path := flag.String(configPath, cfgPath, "Path to the primary JSON config file")
		logPath := flag.String(loggingConfigPath, logConfigPath, "Path to a JSON logging config file")
		psr := flag.String(psrPath, psrConfigPath, "Path to the psr.json file for pre-specified requests")
		miner := flag.Bool(minerArg, false, "Whether to run the miner")
		dataServer := flag.Bool(dataServerArg, false, "Whether to run the data server")
		transfer := flag.Bool(transferArg, false, "Whether to transfer funds")
		deposit := flag.Bool(depositArg, false, "Whether to deposit funds")
		toAddress := flag.String(toAddressArg, "", "Address of party to transfer/approve to")
		amount := flag.String(amountArg, "0", "amount to transfer")
		approve := flag.Bool(approveArg, false, "Whether to transfer funds")
		dispute := flag.Bool(disputeArg, false, "Whether to dispute")
		requestId := flag.String(requestIdArg, "0", "requestId to dispute")
		timestamp := flag.String(timestampArg, "0", "timestamp to dispute")
		minerIndex := flag.String(minerIndexArg, "6", "minerIndex to dispute")
		requestStakingWithdraw := flag.Bool(requestStakingWithdrawArg, false, "Whether to request a staking withdraw")
		withdrawStake := flag.Bool(withdrawStakeArg, false, "Whether to withdrawstake")
		vote := flag.Bool(voteArg, false, "Whether to vote")
		disputeId := flag.String(disputeIdArg, "0", "dispute id to vote on")
		supportsDispute := flag.Bool(supportsDisputeArg, false, "Whether to withdrawstake")
		flag.Parse()
		//log.Printf("Path: %s, LogPath: %s, PSRPath: %s, Args: %v", *path, *logPath, *psr, flag.Args())
		f.ConfigPath = *path
		f.LoggingConfigPath = *logPath
		f.PSRPath = *psr
		f.Miner = *miner
		f.DataServer = *dataServer
		f.Deposit = *deposit
		f.Transfer = *transfer
		f.ToAddress = *toAddress
		f.Amount = *amount
		f.Approve = *approve
		f.Dispute = *dispute
		f.RequestId = *requestId
		f.Timestamp = *timestamp
		f.MinerIndex = *minerIndex
		f.RequestStakingWithdraw = *requestStakingWithdraw
		f.WithdrawStake = *withdrawStake
		f.Vote = *vote
		f.DisputeId = *disputeId
		f.SupportsDispute = *supportsDispute
		sharedFlags = f
	}
	return sharedFlags
}
