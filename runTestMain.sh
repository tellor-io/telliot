#!/bin/sh
PSR=`pwd`/psr.json
#CFG=`pwd`/../../../../localConfig.json
CFG1=`pwd`/config1.json
CFG2=`pwd`/config2.json
CFG3=`pwd`/config3.json
CFG4=`pwd`/config4.json
CFG5=`pwd`/config5.json
LOG=`pwd`/loggingConfig.json
echo "Starting TellorMiner main"
go run . -psrPath="${PSR}" -config="${CFG1}" -logConfig="${LOG}" "$@" &
go run . -psrPath="${PSR}" -config="${CFG2}" -logConfig="${LOG}" "$@" &
go run . -psrPath="${PSR}" -config="${CFG3}" -logConfig="${LOG}" "$@" &
go run . -psrPath="${PSR}" -config="${CFG4}" -logConfig="${LOG}" "$@" &
go run . -psrPath="${PSR}" -config="${CFG5}" -logConfig="${LOG}" "$@"
#go run . -miner -dataServer -psrPath=C:/company/code/go/src/github.com/tellor-io/TellorMiner/psr.json -config=C:/company/code/go/src/github.com/tellor-io/TellorMiner/config4.json -logConfig=C:/company/code/go/src/github.com/tellor-io/TellorMiner/loggingConfig.json
