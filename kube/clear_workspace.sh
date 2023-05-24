#!/bin/bash

kubectl delete deployment pokemon-service-deployment nginx-deployment
kubectl delete se pokeapi
kubectl delete vs pokeapi-through-egress-gateway pokemon-service-ingress
kubectl delete cm nginx-cm pokemon-service-cm
kubectl delete svc pokemon-service-cip nginx-cip
kubectl delete gateway pokeapi-egressgateway pokemon-service-gateway
kubectl label namespace default istio-injection=disabled

cd istio-1.17.2
export PATH=$PWD/bin:$PATH
istioctl manifest generate --set profile=demo | kubectl delete -f -
cd ..

../docker/clear-image.sh
rm -fr istio-1.17.2
rm -fr certs