#!/bin/sh
set -xe

readonly VERSION=${1:-0.0.1}
readonly BUILD_NUMBER=${2:-0}

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"
WORKING_DIR=$(cd "$SCRIPT_DIR"/.. && pwd)

flutter build windows --release \
    --build-number=$BUILD_NUMBER \
    --build-name=$VERSION \
    --dart-define=API_URL=$SERVER_ADDR 