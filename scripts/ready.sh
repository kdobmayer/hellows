#!/bin/sh

status="$(curl -sS http://localhost:8080/health | jq -r '.status')" 
if [ "${status}" -ne "alive" ]; then
  echo "dead"
  exit 1
fi
echo "healthy"

