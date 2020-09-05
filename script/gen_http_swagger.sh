#!/bin/bash

# 生成http swagger文件

# 项目目录
PROJ_DIR=$(dirname "$PWD")

# Swagger目录
SWAG_DIR="${PROJ_DIR}/internal/http"

cd "${SWAG_DIR}"

swag init -g "./swagger.go" --output "./swagger"