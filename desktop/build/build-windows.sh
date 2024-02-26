#!/bin/sh
set -xe

if [ -z "$VERSION" ]; then
    echo "VERSION is not set"
    exit 1
fi

SHORT_VERSION=${VERSION%.*}

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"
WORKING_DIR=$(cd "$SCRIPT_DIR"/.. && pwd)

DEST_FILE=$WORKING_DIR/bin/PhimTor.exe
"$SCRIPT_DIR/build.sh" "$DEST_FILE"

rcedit "$DEST_FILE" --set-icon "$SCRIPT_DIR/icons/icon.ico" --set-file-version "$SHORT_VERSION" --set-product-version "$SHORT_VERSION"



