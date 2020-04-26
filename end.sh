#!/bin/bash

kubectl delete deployments myapp-deployment
kubectl delete services myapp-service
minikube stop
minikube delete
