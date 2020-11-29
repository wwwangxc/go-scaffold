#!/bin/bash

# 以当前系统环境构建grpc应用程序
# main函数目录： ./cmd/main.go
# 目标文件夹：./target/grpc

# 项目目录
PROJ_DIR=$(dirname $(dirname $0))

# 生成的目标目录
TARGET_DIR=${PROJ_DIR}/target/grpc

# 应用程序名称
APP_NAME="grpc"

# 配置文件名称
CONF_NAME="config.toml"

# 目标目录不存在，则创建目录
if [ ! -x "${TARGET_DIR}" ]; then
  mkdir -p "${TARGET_DIR}"
fi

# 目标目录内存在应用程序，则删除
if [ -f "${TARGET_DIR}/${APP_NAME}" ]; then
  rm "${TARGET_DIR}/${APP_NAME}"
fi

# 目标目录内存在配置文件，则删除
if [ -f "${TARGET_DIR}/${CONF_NAME}" ]; then
  rm "${TARGET_DIR}/${CONF_NAME}"
fi

# main函数所在目录
MAIN_DIR="${PROJ_DIR}/cmd"

# 编译
go build -o "${TARGET_DIR}/${APP_NAME}" "${MAIN_DIR}/main.go"

# 配置文件所在目录
CONF_DIR="${PROJ_DIR}/config"

# 存在config.toml文件时复制config.toml文件，否则复制dev.toml文件
if [ -f "${CONF_DIR}/config.toml" ]; then
  cp "${CONF_DIR}/config.toml" "${TARGET_DIR}/${CONF_NAME}"
else
  cp "${CONF_DIR}/dev.toml" "${TARGET_DIR}/${CONF_NAME}"
fi
