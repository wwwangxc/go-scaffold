#!/bin/bash

SCRIPT_DIR=$(dirname $0)

source "${SCRIPT_DIR}/gen_swagger.sh"
source "${SCRIPT_DIR}/go_build_rest_local_env.sh"