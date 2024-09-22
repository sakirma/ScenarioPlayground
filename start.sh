#!/bin/sh

# [x] cilium
# [-] -istio- *won't install it since EBPF and Cilium can outperform istio.
# [ ] I want to connect some services with each other
#   [ ] research what cilium service mesh offers.
# [ ] connect two services with go code using TCP connection
#   [ ] connect with a different service if the current one is dead.


./kind/kind.sh ./kind/kind-cluster.yaml
./cilium/cilium.sh
./podinfo/podinfo.sh

./app/build-push.sh
./app/deploy.sh

cloud-provider-kind