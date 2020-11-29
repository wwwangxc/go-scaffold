#!/bin/bash

# 生成http swagger文件

# 项目目录
PROJ_DIR=$(dirname $(dirname $0))

# 跳转至Swagger目录
cd "${PROJ_DIR}/internal/apis/rest"

swag init -g "./swagger.go" --output "./swagger"