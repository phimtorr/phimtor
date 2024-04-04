#!/bin/sh
set -xe

if [ -z "$VERSION" ]; then
    echo "VERSION is not set"
    exit 1
fi

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"
WORKING_DIR=$(cd "$SCRIPT_DIR"/.. && pwd)

readonly APP_NAME="PhimTor"
readonly APP_DIR="$WORKING_DIR/bin/source/${APP_NAME}.app"

mkdir -p $APP_DIR/Contents/{MacOS,Resources}

env GOARCH=amd64 "$SCRIPT_DIR/build.sh" "$WORKING_DIR/${APP_NAME}.amd64"
env GOARCH=arm64 "$SCRIPT_DIR/build.sh" "$WORKING_DIR/${APP_NAME}.arm64"

lipo -create -output "$APP_DIR/Contents/MacOS/${APP_NAME}" "$WORKING_DIR/${APP_NAME}.amd64" "$WORKING_DIR/${APP_NAME}.arm64"

cat > $APP_DIR/Contents/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>CFBundlePackageType</key>
        <string>APPL</string>
        <key>CFBundleName</key>
        <string>$APP_NAME</string>
        <key>CFBundleExecutable</key>
        <string>$APP_NAME</string>
        <key>CFBundleIdentifier</key>
        <string>net.phimtor</string>
        <key>CFBundleVersion</key>
        <string>$VERSION</string>
        <key>CFBundleGetInfoString</key>
        <string>Built using Go</string>
        <key>CFBundleShortVersionString</key>
        <string>$VERSION</string>
        <key>CFBundleIconFile</key>
        <string>icon.icns</string>
        <key>LSMinimumSystemVersion</key>
        <string>10.13.0</string>
        <key>NSHighResolutionCapable</key>
        <string>true</string>
        <key>NSHumanReadableCopyright</key>
        <string>Copyright.........</string>
    </dict>
</plist>
EOF

cp "$SCRIPT_DIR/icons/icon.icns" "$APP_DIR/Contents/Resources/icon.icns"
find "$APP_DIR"

productbuild --component "$APP_DIR" /Applications "$WORKING_DIR/bin/${APP_NAME}.pkg"

# DMG file
DMG_FILE_NAME="${APP_NAME}-Installer.dmg"
VOLUME_NAME="${APP_NAME} Installer"
SOURCE_FOLDER="$WORKING_DIR/bin/source"
DEST_FILE="$WORKING_DIR/bin/$DMG_FILE_NAME"

[[ -f "${DEST_FILE}" ]] && rm "${DEST_FILE}"

create-dmg \
    --volname "${VOLUME_NAME}" \
    --volicon "$SCRIPT_DIR/icons/icon.icns" \
    --window-pos 200 120 \
    --window-size 800 400 \
    --icon-size 100 \
    --icon "${APP_NAME}.app" 200 190 \
    --hide-extension "${APP_NAME}.app" \
    --app-drop-link 600 185 \
    "${DEST_FILE}" \
    "${SOURCE_FOLDER}"
