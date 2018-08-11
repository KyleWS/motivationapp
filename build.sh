#!/bin/bash
# apparently the . after cgo is important? This has sucked so far
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
echo "remember apprently the fucking login will say it succeeded when it didnt"
docker login
docker build -t kylews/motiapp-api .
docker push kylews/motiapp-api
