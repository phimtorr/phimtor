#!/bin/sh
set -xe

readonly VERSION=${1:-0.0.1}
readonly BUILD_NUMBER=${2:-0}

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"
WORKING_DIR=$(cd "$SCRIPT_DIR"/.. && pwd)

(
    cd "${WORKING_DIR}/torrent/src"
    make android
)

flutter build apk --release \
    --build-number=$BUILD_NUMBER \
    --build-name=$VERSION \
    --dart-define=API_URL=$SERVER_ADDR \
    --dart-define=APP_VERSION=$VERSION

cp ${WORKING_DIR}/build/app/outputs/flutter-apk/app-release.apk ${WORKING_DIR}/build/PhimTor-$VERSION.apk