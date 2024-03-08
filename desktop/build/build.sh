#!/bin/sh
set -xe

DEST_FILE=${1}
if [ -z "$DEST_FILE" ]; then
    echo "Usage: $0 <destination file>"
    exit 1
fi

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"
WORKING_DIR=$(cd "$SCRIPT_DIR"/.. && pwd)
PACKAGE="github.com/phimtorr/phimtor/desktop"

LDFLAGS="-s -w"

if [ -n "$VERSION" ]; then
    LDFLAGS="$LDFLAGS -X ${PACKAGE}/build.Version=$VERSION"
fi

if [ -n "$APP_NAME" ]; then
    LDFLAGS="$LDFLAGS -X ${PACKAGE}/build.AppName=$APP_NAME"
fi

if [ -n "$SERVER_ADDR" ]; then
    LDFLAGS="$LDFLAGS -X ${PACKAGE}/build.ServerAddr=$SERVER_ADDR"
fi

if [ -n "$FIREBASE_API_KEY" ]; then
    LDFLAGS="$LDFLAGS -X ${PACKAGE}/build.FirebaseAPIKey=$FIREBASE_API_KEY"
fi

GOOS=$(go env GOOS)
if [ "$GOOS" = "windows" ]; then
    LDFLAGS="$LDFLAGS -H windowsgui -extldflags=-static"
fi

(
    cd "$WORKING_DIR"
    go build -ldflags "$LDFLAGS" -o "$DEST_FILE"
    chmod +x "$DEST_FILE"
)



