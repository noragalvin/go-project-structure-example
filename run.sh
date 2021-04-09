#!/bin/bash
if ! type fresh > /dev/null; then
  go get github.com/pilu/fresh
fi

source .env

fresh -c runner.conf
