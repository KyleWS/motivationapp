#!/bin/bash

export ADDR=:443
export TLSCERT=/etc/letsencrypt/live/motivationdashboard.app/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/your-host-name.com/privkey.pem
docker run -d --name apiserv -p 443:443 -v /etc/letsencrypt:/etc/letsencrypt:ro -e TLSCERT=$TLSCERT -e TLSKEY=$TLSKEY CONTAINER-NAME-HERE-972359
