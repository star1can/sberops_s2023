#!/bin/bash

# start cluster
minikube start --driver=docker --memory 4096

# download ISTIO
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.17.2 sh -

# export path to bin
cd istio-1.17.2
export PATH=$PWD/bin:$PATH

# apply istio demo profile
istioctl manifest apply --set profile=demo -y

# ban any outgoing traffic from service mesh
istioctl install --set profile=demo --set meshConfig.outboundTrafficPolicy.mode=REGISTRY_ONLY -y
cd ..

# add sidecar proxies auto injection for any pod in ns default
kubectl label namespace default istio-injection=enabled

# add pokemon-service
kubectl apply -f pokemon-service/service.yml -f pokemon-service/configmap/configmap.yml

# add ingress policy
kubectl apply -f ingress/gateway.yml -f ingress/virtual-service.yml

# add egress policy
kubectl apply -f egress/gateway.yml -f egress/virtual-service.yml -f egress/service-entry.yml

# add nginx
kubectl apply -f nginx/service.yml -f nginx/configmap/configmap.yml

# get external IP for ingress via tunnel
minikube tunnel