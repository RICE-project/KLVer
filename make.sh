#!/bin/sh

export GOPATH=$PWD/$(dirname $0)/
echo "Building...\nPath:$GOPATH"

go get github.com/go-sql-driver/mysql
go build

echo "Done."
