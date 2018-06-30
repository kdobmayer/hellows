#!/usr/bin/env bash

external_ip=$(kubectl get svc hello-svc -n=test -o=json | jq -r '.status.loadBalancer.ingress[].ip')
curl http://${external_ip}/hello

