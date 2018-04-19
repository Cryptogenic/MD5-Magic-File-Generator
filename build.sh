#!/bin/sh
set -e
cd ./src/main
go build
mv ./main ../../md5magic
cd ../../
