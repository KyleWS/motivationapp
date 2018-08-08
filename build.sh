#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
docker build -t kylews/motiapp-api .
docker push kylews/motiapp-api
