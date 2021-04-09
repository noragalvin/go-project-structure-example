FROM golang:1.15-alpine

RUN apk add bash ca-certificates git gcc g++ libc-dev curl

ENV CGO_ENABLED=0

ARG GO111MODULE=on

# RUN go get github.com/githubnemo/CompileDaemon
# RUN go get github.com/codegangsta/gin
RUN go get github.com/pilu/fresh

# ENTRYPOINT CompileDaemon -log-prefix=false -build="go build" -command="./gemtickets"
# ENTRYPOINT gin -p 8000 -a 8001 run server.go
ENTRYPOINT fresh -c runner.conf
