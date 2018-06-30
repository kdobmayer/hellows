#!/usr/bin/env bash

kubens() {
  kubectl config set-context $(kubectl config current-context) --namespace="$1"
}

get_ip() {
  external_ip=$(kubectl get svc hello-svc -o=json \
    | jq -r '.status.loadBalancer.ingress[].ip' 2>/dev/null) || return 1
  echo "${external_ip}"
}

#docker build -t kdobmayer/hellows .
#docker push kdobmayer/hellows:latest

kubens test
kubectl delete svc hello-svc
kubectl delete deploy hello-deploy
#kubectl delete quota test-quotas
kubectl delete ns test

while kubectl get ns | grep -qc test; do
  printf "\rDeleting namespace..."
  sleep 5
done
echo "Done!"

#kubectl create -f ./hello.yaml
#kubectl get pods
#
#until get_ip >/dev/null; do
#  printf "\rWaiting for external ip..."
#  sleep 5 
#done
#echo "Done!"
#
#external_ip=$(get_ip)
#for ((i=0; i<10; i++)); do
#  pod="$(curl -sS http://${external_ip}/hello | awk -F- '{print $NF}')"
#  # remove the trailing exclamation mark '!'
#  echo "${pod::-1}"
#  sleep 1
#done

kubens default

# | awk -F- '{print substr($NF, 1, length($NF)-1)}'
