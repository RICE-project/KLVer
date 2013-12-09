#!/bin/sh

export GOPATH=$PWD/$(dirname $0)/
echo "Building...\nPath:$GOPATH"

go get github.com/tonychee7000/mysql
go build

echo "Done."
