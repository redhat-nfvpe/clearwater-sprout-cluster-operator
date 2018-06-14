#!/bin/bash
set -e

HERE=$(dirname "$(readlink -f "$0")")

if [ -z "$GOPATH" ]; then
	GOPATH="$HOME/go"
fi

PROJECT_BASE="$GOPATH/src/github.com/redhat-nfvpe"

PATH="$GOPATH/bin:$PATH"

mkdir --parents "$PROJECT_BASE"

cd "$PROJECT_BASE"

GOPATH=$GOPATH \
operator-sdk new clearwater-sprout-cluster-operator \
--kind=SproutCluster \
--api-version=projectclearwater.org/v1 \
--skip-git-init
