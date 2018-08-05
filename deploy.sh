#!/usr/bin/env bash
./build.sh
ssh root@206.189.79.199 '
docker rm -f mongo-server
docker run -d --name mongo-server -p 27017:27017 mongo
docker rm -f motiapp-server
docker pull kylews/motiapp-api
docker run -d \
-p 443:443 \
--name motiapp-server \
-v /etc/letsencrypt:/letsencrypt:ro \
-e TLSCERT=/etc/letsencrypt/live/api.motivationdashboard.app/fullchain.pem \
-e TLSKEY=/etc/letsencrypt/live/api.motivationdashboard.app/privkey.pem \
kylews/motiapp-api'
