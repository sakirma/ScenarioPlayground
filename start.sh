#!/bin/sh

# [x] cilium
# [-] -istio- *won't install it since EBPF and Cilium can outperform istio.
# [x] I want to connect some services with each other
#   [x] research what cilium service mesh offers.
# [x] connect two services with go code using TCP connection
# [ ] Check out about Kafka.


./kind/kind.sh ./kind/kind-cluster.yaml
./cilium/cilium.sh
./podinfo/podinfo.sh

./app/build-push.sh
./helm/apply.sh

./kafka/apply.sh

cloud-provider-kind