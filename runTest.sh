#!/bin/sh

NAME="$1"
PKG="$2"

if [ -z "${NAME}" ]; then
  echo "Missing test name";
  exit 1;
fi

if [ -z "${PKG}" ]; then
   echo "Missing package name";
   exit 2;
fi

echo "Running test ${NAME} from package  ${PKG}..."
go test -v -run "${NAME}" "./${PKG}" 
