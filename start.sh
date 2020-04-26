#!/bin/bash

minikube start --extra-config=apiserver.v=10 --extra-config "apiserver.cors-allowed-origins=["http://\*"]"
#kubectl run crasher --image=rossKukulinski/crashing-app
kubectl create -f documents/deployment-definition.yml
#kubectl run name --image=anyimagenamethatdoesntexist


