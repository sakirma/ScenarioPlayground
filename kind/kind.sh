#!/bin/zsh

kind create cluster --config="$1"

kubectl label node kind-control-plane node.kubernetes.io/exclude-from-external-load-balancers-node/kind-control-plane unlabeled
