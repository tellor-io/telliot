#!/bin/sh
PSR=`pwd`/psr2.json
#CFG=`pwd`/../../../../localConfig.json
CFG1=`pwd`/config1.json
CFG2=`pwd`/config2.json
CFG3=`pwd`/config3.json
CFG4=`pwd`/config4.json
CFG5=`pwd`/config5.json
CFG6=`pwd`/config6.json
LOG=`pwd`/loggingConfig.json
echo "Starting TellorMiner main"
nohup go run . -psrPath="${PSR}" -config="${CFG6}" -logConfig="${LOG}" "$@" &
nohup go run . -psrPath="${PSR}" -config="${CFG1}" -logConfig="${LOG}" "$@" &
nohup go run . -psrPath="${PSR}" -config="${CFG2}" -logConfig="${LOG}" "$@" &
nohup go run . -psrPath="${PSR}" -config="${CFG3}" -logConfig="${LOG}" "$@" &
nohup go run . -psrPath="${PSR}" -config="${CFG4}" -logConfig="${LOG}" "$@" &
nohup go run . -psrPath="${PSR}" -config="${CFG5}" -logConfig="${LOG}" "$@" 
#go run . -miner -dataServer -psrPath=C:/company/code/go/src/github.com/tellor-io/TellorMiner/psr.json -config=C:/company/code/go/src/github.com/tellor-io/TellorMiner/config3.json -logConfig=C:/company/code/go/src/github.com/tellor-io/TellorMiner/loggingConfig.json
#go run ./main.go -deposit -config=./config6.json -psrPath=./psr2.json -logConfig=./loggingConfig.json
#go run ./main.go -transfer -to=0xe0d7bae200f0994b11423e8be8f386060bbdd808 -amount=1000000000000000000000 -config=./config3.json -psrPath=./psr2.json -logConfig=./loggingConfig.json

#0x8b73fa2c839ccea66e8eddf0aa95f6bc4c6aaa11e2fa126c1d9334985b0e7666
#0xe0d7bae200f0994b11423e8be8f386060bbdd808