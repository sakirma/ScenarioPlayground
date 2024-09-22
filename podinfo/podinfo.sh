#!/bin/sh

kubectl create namespace podinfo

helm repo add podinfo https://stefanprodan.github.io/podinfo || true

helm upgrade --install --wait frontend \
--namespace podinfo \
--set replicaCount=2 \
--set backend=http://backend-podinfo:9898/echo \
podinfo/podinfo

helm upgrade --install --wait backend \
--namespace podinfo \
--set redis.enabled=true \
podinfo/podinfo