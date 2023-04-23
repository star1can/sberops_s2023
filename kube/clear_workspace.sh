#!/bin/bash

kubectl delete deployment pokemon-service-deployment nginx-deployment
kubectl delete se pokeapi
kubectl delete vs pokeapi-through-egress-gateway pokemon-service-ingress
kubectl delete cm nginx-cm pokemon-service-cm
kubectl delete svc pokemon-service-cip nginx-cip
kubectl delete gateway pokeapi-egressgateway pokemon-service-gateway
../docker/clear-image.sh
rm -fr istio-1.17.2