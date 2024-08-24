#!/bin/sh
set -xe

readonly VERSION=${1:-0.0.1}
readonly BUILD_NUMBER=${2:-0}

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"
WORKING_DIR=$(cd "$SCRIPT_DIR"/.. && pwd)
APP_DIR="$WORKING_DIR"/build/Phimtor.AppDir


cd "$WORKING_DIR"

echo "Working dir: $(pwd)"

flutter build linux --release \
    --build-number=$BUILD_NUMBER \
    --build-name=$VERSION \
    --dart-define=API_URL=$SERVER_ADDR \
    --dart-define=APP_VERSION=$VERSION

mkdir -p "$APP_DIR"
rm -rf "$APP_DIR"/*

cp -r build/linux/x64/release/bundle//* "$APP_DIR"/
cp assets/icon/icon.png "$APP_DIR"/

cat > "$APP_DIR"/AppRun <<EOL
#!/bin/sh
cd "\$(dirname "\$0")"
exec ./phimtorapp
EOL

chmod +x "$APP_DIR"/AppRun


cat > "$APP_DIR"/PhimTor.desktop <<EOL
[Desktop Entry]
Name=Phim Tor
Exec=phimtorapp
Icon=icon
Type=Application
Terminal=false
Categories=AudioVideo;Player;
EOL

if [ -f "appimagetool-x86_64.AppImage" ]; then
    echo "appimagetool-x86_64.AppImage already exists"
else
    wget "https://github.com/AppImage/AppImageKit/releases/download/continuous/appimagetool-x86_64.AppImage"
    chmod a+x appimagetool-x86_64.AppImage 
fi

./appimagetool-x86_64.AppImage "$APP_DIR" "build/PhimTor-x86_64.AppImage"

