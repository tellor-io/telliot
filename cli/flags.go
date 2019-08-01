package cli

import (
	"flag"
)

const configPath = "config"
const loggingConfigPath = "logConfig"
const psrPath = "psrPath"
const minerArg = "miner"
const dataServerArg = "dataServer"

//Flags holds all command line options
type Flags struct {
	ConfigPath        string
	LoggingConfigPath string
	PSRPath           string
	Miner             bool
	DataServer        bool
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

		flag.Parse()
		//log.Printf("Path: %s, LogPath: %s, PSRPath: %s, Args: %v", *path, *logPath, *psr, flag.Args())
		f.ConfigPath = *path
		f.LoggingConfigPath = *logPath
		f.PSRPath = *psr
		f.Miner = *miner
		f.DataServer = *dataServer
		sharedFlags = f
	}
	return sharedFlags
}
