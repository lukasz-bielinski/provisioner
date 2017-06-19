#!/usr/bin/env bash

set -x

CONTAINER_NAME=lukaszbielinski/aela
#CONTAINER_TAG=2
# CONTAINER_NAME=${1}
CONTAINER_TAG=${1}

PROJECT_NAME='github.com/lukaszbielinski/go-test'
PROJECT_DIR="${PWD}/src"
VENDOR_DIR='Godeps/_workspace'

CONTAINER_GOPATH='/go'
CONTAINER_PROJECT_DIR="${CONTAINER_GOPATH}/src/${PROJECT_NAME}"
CONTAINER_PROJECT_GOPATH="${CONTAINER_PROJECT_DIR}/${VENDOR_DIR}:${CONTAINER_GOPATH}"

docker run --rm \
        -v ${PROJECT_DIR}:${CONTAINER_PROJECT_DIR} \
        -e GOPATH=${CONTAINER_PROJECT_GOPATH} \
        -e CGO_ENABLED=0 \
        -e GODEBUG=netdns=go \
        -w "${CONTAINER_PROJECT_DIR}" \
        golang:1.8.3-alpine \
        go build -v -o bin/test test.go

# Disable this to strip the debug information from the binary and shave off about 4Mb making the binary from 12mb to 8mb
# It means this can't be debugged by delve, gdb et al. but the side is even better
strip "${PROJECT_DIR}/test"

# docker build -f ${PROJECT_DIR}/Dockerfile \
#     -t ${CONTAINER_NAME}:${CONTAINER_TAG} \
#     --build-arg BINARY_FILE=./test \
#     "${PROJECT_DIR}"

docker build -t ${CONTAINER_NAME}:${CONTAINER_TAG} -f Dockerfile .
docker push ${CONTAINER_NAME}:${CONTAINER_TAG}
