#!/bin/sh
set -xe

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"
WORKING_DIR=$(cd "$SCRIPT_DIR"/.. && pwd)

if [ -z "$VERSION" ]; then
    echo "VERSION is not set"
    exit 1
fi
SHORT_VERSION=${VERSION%.*}
readonly APP=PhimTor
readonly APP_DIR="$WORKING_DIR/bin/$APP"

mkdir -p "$APP_DIR/usr/bin"
mkdir -p "$APP_DIR/usr/share/applications"
mkdir -p "$APP_DIR/usr/share/icons/hicolor/1024x1024/apps"
mkdir -p "$APP_DIR/usr/share/icons/hicolor/256x256/apps"
mkdir -p "$APP_DIR/DEBIAN"

"$SCRIPT_DIR/build.sh" "$APP_DIR/usr/bin/$APP"

cp "$SCRIPT_DIR/icons/icon.png" "$APP_DIR/usr/share/icons/hicolor/1024x1024/apps/${APP}.png"
cp "$SCRIPT_DIR/icons/icon.png" "$APP_DIR/usr/share/icons/hicolor/256x256/apps/${APP}.png"

cat > "$APP_DIR/usr/share/applications/${APP}.desktop" << EOF
[Desktop Entry]
Version=1.0
Type=Application
Name=$APP
Exec=$APP
Icon=$APP
Terminal=false
StartupWMClass=PhimTor
EOF

cat > "$APP_DIR/DEBIAN/control" << EOF
Package: ${APP}
Version: ${SHORT_VERSION}-0
Section: base
Priority: optional
Architecture: amd64
Maintainer: Chien Nguyen
Description: A tool for streaming films from the torrent network.
EOF

dpkg-deb --build "$APP_DIR"