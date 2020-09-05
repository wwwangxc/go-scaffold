#!/bin/bash

# 项目目录
PROJ_DIR=$(dirname "$PWD")

# Protobuf目录
PB_DIR="${PROJ_DIR}/internal/grpc/pb"

cd "${PBDIR}"

protoc --go_out=plugins=grpc:. *.proto
