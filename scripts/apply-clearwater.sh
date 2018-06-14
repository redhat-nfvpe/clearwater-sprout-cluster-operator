#!/bin/bash
set -e

HERE=$(dirname "$(readlink -f "$0")")
ROOT=$(realpath "$HERE/..")

kubectl apply -f "$ROOT/assets/crd.yaml"
kubectl apply -f "$ROOT/assets/clearwater.yaml"
