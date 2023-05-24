#!/bin/bash

export INGRESS_NAME=istio-ingressgateway
export INGRESS_NS=istio-system
export INGRESS_HOST=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export SECURE_INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="https")].port}')
export TCP_INGRESS_PORT=$(kubectl -n "$INGRESS_NS" get service "$INGRESS_NAME" -o jsonpath='{.spec.ports[?(@.name=="tcp")].port}')

curl -v -HHost:randact.example.com --resolve "mipt.pokemons.com:$SECURE_INGRESS_PORT:$INGRESS_HOST" \
  --cacert kube/certs/pokemons.com.crt --cert kube/certs/mipt.pokemons.com.crt --key kube/certs/mipt.pokemons.com.key \
  "https://mipt.pokemons.com:$SECURE_INGRESS_PORT/pokemons/all"