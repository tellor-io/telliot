#!/bin/sh

PKG="$1"
PSR=`pwd`/psr.json
CFG=`pwd`/../../../../localConfig.json
LOG=`pwd`/loggingConfig.json

if [ -z "${PKG}" ]; then
   echo "Missing package name";
   exit 2;
fi

echo "Running test ${NAME} from package  ${PKG}..."
go test -v "./${PKG}" -psrPath="${PSR}" -config="${CFG}" -logConfig="${LOG}"