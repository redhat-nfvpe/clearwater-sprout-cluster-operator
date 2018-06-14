#!/bin/bash
set -e

# See: https://blog.hasura.io/sharing-a-local-registry-for-minikube-37c7240d0615

HERE=$(dirname "$(readlink -f "$0")")
ROOT=$(realpath "$HERE/..")

if [ -z "$GOPATH" ]; then
	GOPATH="$HOME/go"
fi

PATH="$GOPATH/bin:$PATH"

kubectl apply -f "$ROOT/assets/kube-registry.yaml"

POD=$(kubectl get pods -n kube-system -l app=kube-registry -l role=registry -o jsonpath='{.items[0].metadata.name}')
kubectl port-forward -n kube-system $POD 5000:5000 &
