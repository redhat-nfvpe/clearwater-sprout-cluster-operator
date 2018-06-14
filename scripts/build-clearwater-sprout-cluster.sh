#!/bin/bash
set -e

HERE=$(dirname "$(readlink -f "$0")")

if [ -z "$GOPATH" ]; then
	GOPATH="$HOME/go"
fi

PROJECT="$GOPATH/src/github.com/redhat-nfvpe/clearwater-sprout-cluster-operator"
IMAGE=localhost:5000/clearwater-sprout-cluster-operator:0.1

PATH="$GOPATH/bin:$PATH"

cd "$PROJECT"

set +e
docker rmi "$IMAGE"
set -e
operator-sdk build "$IMAGE"
docker push "$IMAGE"
