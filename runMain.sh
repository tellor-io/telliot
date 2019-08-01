#!/bin/sh

PSR=`pwd`/psr.json
CFG=`pwd`/../../../../localConfig.json
LOG=`pwd`/loggingConfig.json


echo "Starting TellorMiner main"
go run . -psrPath="${PSR}" -config="${CFG}" -logConfig="${LOG}" "$@"