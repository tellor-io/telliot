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
# Hack to get manualData.json inside package scope.
# Not sure if this counts as "test instrumentation" or just a hack.
# Go seems to be designed for package folders to serve as "root" in a 
# test context. If we want this json file to be accessible without this,
# either [1] a config for the manual data path is needed, where the 
# path can be set, or [2] a refactor is needed to utilize absolute paths
cp manualData.json ${PKG}
go test -coverprofile trackercoverage.out -v "./${PKG}"
testResult=$?
go tool cover -func=trackercoverage.out
rm ${PKG}/manualData.json

exit $testResult
