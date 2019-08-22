#!/bin/sh
PSR=`pwd`/psr.json
#CFG=`pwd`/../../../../localConfig.json
CFG=`pwd`/myconfig.json
LOG=`pwd`/loggingConfig.json
echo "Starting TellorMiner main"
go run . -psrPath="${PSR}" -config="${CFG}" -logConfig="${LOG}" "$@"


# ./TellorMiner -miner -dataServer -psrPath=./psr.json -config=./config.json -logConfig=./loggingconfig.json
# go run ./main.go -miner -dataServer -psrPath=./psr.json -config=./config.json -logConfig=./loggingconfig.json
#go run ./main.go -deposit -config=./config6.json -psrPath=./psr2.json -logConfig=./loggingConfig.json
#go run ./main.go -transfer -to=0x538d278e05a35c96bcdca1039e92c65b994256a0 -amount=1 -config=./config3.json -psrPath=./psr2.json -logConfig=./loggingConfig.json