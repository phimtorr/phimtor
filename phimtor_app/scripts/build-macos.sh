#!/bin/sh
set -xe

readonly VERSION=${1:-0.0.1}
readonly BUILD_NUMBER=${2:-0}

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"
WORKING_DIR=$(cd "$SCRIPT_DIR"/.. && pwd)

(
    cd "${WORKING_DIR}/torrent/src"
    make macos
)

flutter build macos --release \
    --build-number=$BUILD_NUMBER \
    --build-name=$VERSION \
    --dart-define=API_URL=$SERVER_ADDR \
    --dart-define=APP_VERSION=$VERSION

APP_DIR="${WORKING_DIR}/build/macos/Build/Products/Release/PhimTor.app"

productbuild --component "$APP_DIR" /Applications "$WORKING_DIR/build/PhimTor-${VERSION}.pkg"

# DMG file 
mkdir -p "$WORKING_DIR/build/dmg"
cp -r "$APP_DIR" "$WORKING_DIR/build/dmg/PhimTor.app"

DMG_FILE_NAME="PhimTor-${VERSION}.dmg"
VOLUME_NAME="PhimTor Installer"
SOURCE_FOLDER="$WORKING_DIR/build/dmg"
DEST_FILE="$WORKING_DIR/build/$DMG_FILE_NAME"
APP_NAME="PhimTor"

[[ -f "${DEST_FILE}" ]] && rm "${DEST_FILE}"

create-dmg \
    --volname "${VOLUME_NAME}" \
    --volicon "$WORKING_DIR/assets/icon/icon.icns" \
    --window-pos 200 120 \
    --window-size 800 400 \
    --icon-size 100 \
    --icon "${APP_NAME}.app" 200 190 \
    --hide-extension "${APP_NAME}.app" \
    --app-drop-link 600 185 \
    "${DEST_FILE}" \
    "${SOURCE_FOLDER}"