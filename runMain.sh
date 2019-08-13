#!/bin/sh
PSR=`pwd`/psr.json
#CFG=`pwd`/../../../../localConfig.json
CFG=`pwd`/config.json
LOG=`pwd`/loggingConfig.json
echo "Starting TellorMiner main"
go run . -psrPath="${PSR}" -config="${CFG}" -logConfig="${LOG}" "$@"


./TellorMiner -miner -dataServer -psrPath=./psr.json -config=./config.json -logConfig=./loggingconfig.json


#go run . -psrPah=C:/company/code/go/src/github.com/tellor-io/TellorMiner/psr.json -config=C:/company/code/go/src/github.com/tellor-io/TellorMiner/config.json -logConfig=C:/company/code/go/src/github.com/tellor-io/TellorMiner/loggingConfig.json