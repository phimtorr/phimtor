#!/bin/sh
set -xe

if [ -z "$VERSION" ]; then
    echo "VERSION is not set"
    exit 1
fi

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"
WORKING_DIR=$(cd "$SCRIPT_DIR"/.. && pwd)

readonly APPDIR="$WORKING_DIR/bin/PhimTor.app"

mkdir -p $APPDIR/Contents/{MacOS,Resources}

"$SCRIPT_DIR/build.sh" "$APPDIR/Contents/MacOS/PhimTor"

cat > $APPDIR/Contents/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>CFBundlePackageType</key>
        <string>APPL</string>
        <key>CFBundleName</key>
        <string>PhimTor</string>
        <key>CFBundleExecutable</key>
        <string>PhimTor</string>
        <key>CFBundleIdentifier</key>
        <string>com.PhimTor</string>
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

cp "$SCRIPT_DIR/icons/icon.icns" "$APPDIR/Contents/Resources/icon.icns"
find "$APPDIR"

productbuild --component "$APPDIR" "$WORKING_DIR/bin/PhimTor.pkg"