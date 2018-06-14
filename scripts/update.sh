#!/bin/bash
set -e

HERE=$(dirname "$(readlink -f "$0")")

if [ -z "$GOPATH" ]; then
	GOPATH="$HOME/go"
fi

PROJECT="$GOPATH/src/github.com/redhat-nfvpe/clearwater-sprout-cluster-operator"

PATH="$GOPATH/bin:$PATH"

cd "$PROJECT"

operator-sdk generate k8s 
