#!/bin/zsh

helm repo add cilium https://helm.cilium.io/


docker pull quay.io/cilium/cilium:v1.16.1
kind load docker-image quay.io/cilium/cilium:v1.16.1

helm install cilium cilium/cilium --version 1.16.1 \
   --namespace kube-system \
   --set image.pullPolicy=IfNotPresent \
   --set ipam.mode=kubernetes