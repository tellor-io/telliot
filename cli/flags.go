package cli

import (
	"flag"
)

const configPath = "config"
const loggingConfigPath = "logConfig"
const psrPath = "psrPath"

//Flags holds all command line options
type Flags struct {
	ConfigPath        string
	LoggingConfigPath string
	PSRPath           string
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
		flag.Parse()
		//log.Printf("Path: %s, LogPath: %s, PSRPath: %s, Args: %v", *path, *logPath, *psr, flag.Args())
		f.ConfigPath = *path
		f.LoggingConfigPath = *logPath
		f.PSRPath = *psr
		sharedFlags = f
	}
	return sharedFlags
}
