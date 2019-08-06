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

//Flags holds all command line options
type Flags struct {
	ConfigPath        string
	LoggingConfigPath string
	PSRPath           string
	Miner             bool
	DataServer        bool
	Transfer          bool
	ToAddress         string
	Amount            int
	Deposit           bool
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
		path := flag.String(configPath, "", "Path to the primary JSON config file")
		logPath := flag.String(loggingConfigPath, "", "Path to a JSON logging config file")
		psr := flag.String(psrPath, "", "Path to the psr.json file for pre-specified requests")
		miner := flag.Bool(minerArg, false, "Whether to run the miner")
		dataServer := flag.Bool(dataServerArg, false, "Whether to run the data server")
		transfer := flag.Bool(transferArg, false, "Whether to transfer funds")
		deposit := flag.Bool(depositArg, false, "Whether to deposit funds")
		toAddress := flag.String(toAddressArg, "", "Address of party to transfer to")
		amount := flag.Int(amountArg, 0, "amount to transfer")

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
		sharedFlags = f
	}
	return sharedFlags
}
