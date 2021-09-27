#!/bin/bash
if ! type air > /dev/null; then
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
fi

# source .env
air -c shopify.air.toml
# gin -p 8000 -a 8001 run server.go