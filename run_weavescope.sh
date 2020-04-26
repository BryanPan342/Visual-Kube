#!/bin/bash
kubectl apply -f test/examples/k8s
kubectl port-forward svc/weave-scope-app -n weave 4040:80
