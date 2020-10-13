#!/bin/sh

PKG="$1"
PSR=`pwd`/psr.json
CFG=`pwd`/../../../../localConfig.json
LOG=`pwd`/loggingConfig.json

if [ -z "${PKG}" ]; then
   echo "Missing package name";
   exit 2;
fi
 #// TODO Add to make.......
echo "Running test ${NAME} from package  ${PKG}..."
go test -coverprofile trackercoverage.out -v "./${PKG}"
testResult=$?
go tool cover -func=trackercoverage.out
rm ${PKG}/possible-dispute*.txt

exit $testResult
