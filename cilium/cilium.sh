#!/bin/sh

helm repo add cilium https://helm.cilium.io/

docker pull quay.io/cilium/cilium:v1.16.1
kind load docker-image quay.io/cilium/cilium:v1.16.1

helm upgrade --install cilium cilium/cilium --version 1.16.1 \
    --namespace kube-system \
    --set image.pullPolicy=IfNotPresent \
    --set kubeProxyReplacement=true \
    --set k8sServiceHost=kind-control-plane \
    --set k8sServicePort=6443