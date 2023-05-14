#!/bin/env sh
set -ex 

minikube start

eval $(minikube -p minikube docker-env) 

minikube addons enable ingress

docker build -t hi-er:1.0 .

kubectl apply -f k8s/

sleep 10

curl --resolve "hi.info:80:$( minikube ip )"  hi.info