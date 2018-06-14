#!/bin/bash
set -e

if [ -z "$GOPATH" ]; then
	GOPATH="$HOME/go"
fi

PROJECT="$GOPATH/src/github.com/operator-framework/operator-sdk"

cd "$GOPATH"

go get -u github.com/golang/dep/cmd/dep

mkdir --parents "$PROJECT"
cd "$PROJECT"

git clone https://github.com/operator-framework/operator-sdk.git .
"$GOPATH/bin/dep" ensure
go install github.com/operator-framework/operator-sdk/commands/operator-sdk
