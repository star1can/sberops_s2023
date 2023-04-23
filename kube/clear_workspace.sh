#!/bin/bash

minikube delete
../docker/clear-image.sh
rm -fr istio-1.17.2