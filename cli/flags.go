package cli

import (
	"flag"
	"log"
)

const configPath = "config"
const loggingConfigPath = "logConfig"

//Flags holds all command line options
type Flags struct {
	ConfigPath        string
	LoggingConfigPath string
}

var (
	sharedFlags *Flags
)

//GetFlags parses command line entries into a shared structure with values
func GetFlags() *Flags {
	if sharedFlags == nil {
		f := &Flags{}
		path := flag.String(configPath, "", "Path to the primary JSON config file")
		logPath := flag.String(loggingConfigPath, "", "Path to a JSON logging config file")
		flag.Parse()
		log.Printf("Path, LogPath: %s, %s, %v", *path, *logPath, flag.Args())
		f.ConfigPath = *path
		f.LoggingConfigPath = *logPath
		sharedFlags = f
	}
	return sharedFlags
}
