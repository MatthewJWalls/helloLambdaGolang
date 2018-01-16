#!/usr/bin/env bash

set -u
set -e

GREEN="\033[0;32m"
PS_CLEAR="\033[0m"

function echo_g() {
    echo -e "\n$GREEN$1$PS_CLEAR"
}

echo_g "running tests"
go test

echo_g "building deployment"
GOOS=linux go build -o main
zip deployment.zip main
rm main

echo_g "uploading to aws lambda"
aws lambda update-function-code --function-name HelloGolang --zip-file fileb://./deployment.zip --region eu-west-1 --profile frontend > /dev/null
