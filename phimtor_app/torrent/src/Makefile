LIB_NAME=libtorrent
macos:
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=arm64 \
	SDK=macos \
	go build -trimpath -buildmode=c-shared -o ${LIB_NAME}_arm64.dylib main.go
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=amd64 \
	SDK=macos \
	go build -trimpath -buildmode=c-shared -o ${LIB_NAME}_amd64.dylib main.go
	lipo \
		-create \
		${LIB_NAME}_arm64.dylib \
		${LIB_NAME}_amd64.dylib \
		-output ${LIB_NAME}.dylib
	install_name_tool -id "@rpath/libtorrent.dylib" libtorrent.dylib
	mv ${LIB_NAME}.dylib ../macos/${LIB_NAME}.dylib
	rm ${LIB_NAME}_*.h
	rm ${LIB_NAME}_*.dylib