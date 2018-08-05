#!/usr/bin/env bash
set -e CGO_ENABLED=0
go build -a
docker build -t kylews/motiapp-api .
docker push kylews/motiapp-api
go clean