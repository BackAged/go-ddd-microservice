#!/bin/sh
set -e

if ! [ -x "$(command -v protoc)" ]; then
    echo "protoc is not installed"
    exit 1
fi

if ! [ -x "$(command -v go)" ]; then
    echo "go is not installed"
    exit 1
fi

if ! [ -x "$(command -v protoc-gen-go)" ]; then
    echo "protoc-gen-go is not installed"
    exit 1
fi

proto_path="$PWD/proto/*"

for d in $proto_path; do
    if [ -d "${d}" ]; then 
        echo "========================= compiling $d ======================="
        echo "proto/$(basename "$d")/"
        protoc --proto_path=proto --go_out=plugins=grpc:pb $(find proto/$(basename "$d")/ -iname "*.proto")
        echo "========================= done compiling ======================="
    else
        echo "========================= compiling $d ======================="
        protoc --proto_path=./proto --go_out=plugins=grpc:$PWD "proto/$(basename "$d")"
        echo "========================= done compiling ======================="
    fi
done