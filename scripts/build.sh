#!/bin/bash
set -e
OUTPUT_DIR="./bin"
BINARY_NAME="optipix-socket"

rm -rf ${OUTPUT_DIR}
mkdir -p ${OUTPUT_DIR}

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${OUTPUT_DIR}/${BINARY_NAME} ./cmd/main.go

echo "Build completed. Binary is available at ${OUTPUT_DIR}/${BINARY_NAME}."
