#!/bin/bash

SCRIPT_DIR=$(dirname $0)

source "${SCRIPT_DIR}/gen_http_swagger.sh"
source "${SCRIPT_DIR}/go_build_http_local_env.sh"