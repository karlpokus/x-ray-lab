#!/bin/bash

cd k8s/app
# point to minikube dockerd
eval $(minikube -p minikube docker-env)
docker build . -t cf/xray-k8s-app
# reset to host dockerd
eval $(minikube docker-env -u)
