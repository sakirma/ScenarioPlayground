#!/bin/sh

# Team A
kubectl create namespace team-a

helm repo add podinfo https://stefanprodan.github.io/podinfo || true

helm upgrade --install --wait frontend \
--namespace team-a \
--set replicaCount=2 \
--set backend=http://backend-podinfo:9898/echo \
podinfo/podinfo

helm upgrade --install --wait backend \
--namespace team-a \
--set redis.enabled=true \
podinfo/podinfo

# Team B
kubectl create namespace team-b

helm upgrade --install --wait frontend \
--namespace team-b \
--set replicaCount=2 \
--set backend=http://backend-podinfo:9898/echo \
podinfo/podinfo

helm upgrade --install --wait backend \
--namespace team-b \
--set redis.enabled=true \
podinfo/podinfo