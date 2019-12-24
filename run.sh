#!/bin/bash
if ! type fresh > /dev/null; then
    # go get github.com/githubnemo/CompileDaemon
    go get github.com/pilu/fresh
fi

source .env

# CompileDaemon -command="./gemtickets" 
fresh -c runner.conf cmd/gemtracker/main.go
# gin -p 8000 -a 8001 run server.go