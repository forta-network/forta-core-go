#!/bin/sh

set -xe

export PRIVATE_KEY="<dev-node-private-key>"
export ALERT_API="https://alerts-dev.forta.network"
export BATCH_FILE="$(pwd)/batch.json"

go run scripts/create-dev-alert/main.go
