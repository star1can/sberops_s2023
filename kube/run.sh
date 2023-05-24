#!/bin/bash

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

# create certs
mkdir certs
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=pokemons Inc./CN=pokemons.com' -keyout certs/pokemons.com.key -out certs/pokemons.com.crt
openssl req -out certs/mipt.pokemons.com.csr -newkey rsa:2048 -nodes -keyout certs/mipt.pokemons.com.key -subj "/CN=mipt.pokemons.com/O=mipt organization"
openssl x509 -req -sha256 -days 365 -CA certs/pokemons.com.crt -CAkey certs/pokemons.com.key -set_serial 0 -in certs/mipt.pokemons.com.csr -out certs/mipt.pokemons.com.crt

kubectl create -n istio-system secret generic pokemons-secret \
  --from-file=tls.key=certs/mipt.pokemons.com.key \
  --from-file=tls.crt=certs/mipt.pokemons.com.crt \
  --from-file=ca.crt=certs/pokemons.com.crt

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