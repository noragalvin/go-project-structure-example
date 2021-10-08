FROM golang:1.17-alpine

RUN apk add bash ca-certificates git gcc g++ libc-dev curl

ENV CGO_ENABLED=0

ARG GO111MODULE=on

# Add the keys
ARG github_user
ENV github_user=$github_user
ARG github_token
ENV github_token=$github_token

RUN git config \
  --global \
  url."https://${github_user}:${github_token}@github.com".insteadOf \
  "https://github.com"


CMD ["./scripts/shopify.sh"]
