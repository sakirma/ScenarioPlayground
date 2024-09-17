#!/bin/zsh


./kind/kind.sh ./kind/kind-cluster.yaml
./cilium/cilium.sh

cloud-provider-kind