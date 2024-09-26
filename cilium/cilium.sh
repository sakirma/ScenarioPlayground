#!/bin/sh

helm repo add cilium https://helm.cilium.io/

docker pull quay.io/cilium/cilium:v1.16.1
kind load docker-image quay.io/cilium/cilium:v1.16.1

cilium install --version 1.16.1

# This costs extra resource to have lower latency and higher through output
#helm install cilium cilium/cilium --version 1.16.1 -n kube-system \
#    --set routingMode=native \
#    --set bpf.datapathMode=netkit \
#    --set bpf.masquerade=true \
#    --set ipv4.enabled=true \
#    --set enableIPv4BIGTCP=true \
#    --set ipam.mode=kubernetes \
#    --set autoDirectNodeRoutes=true \
#    --set ipv4NativeRoutingCIDR="10.0.0.0/8" \
#    --set kubeProxyReplacement=true \
#    --set bandwidthManager.enabled=true \
#    --set bandwidthManager.bbr=true \
#    --set k8sServiceHost=kind-control-plane \
#    --set k8sServicePort=6443