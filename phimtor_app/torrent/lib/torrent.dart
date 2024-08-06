import 'dart:ffi' as ffi;
import 'dart:io' show Platform;

import 'package:ffi/ffi.dart' as ffi2;
import 'package:openapi_client/api.dart';
import 'package:path/path.dart' as p;

import 'torrent_bindings_generated.dart';

export 'package:openapi_client/api.dart';

const String _libName = 'torrent';

final ffi.DynamicLibrary _dylib = () {
  if (Platform.isMacOS || Platform.isIOS) {
    // Add from here...
    if (Platform.environment.containsKey('FLUTTER_TEST')) {
      return ffi.DynamicLibrary.open('build/macos/Build/Products/Debug'
          '/$_libName/$_libName.framework/$_libName');
    }
    // ...to here.
    return ffi.DynamicLibrary.open('$_libName.framework/$_libName');
  }
  if (Platform.isAndroid || Platform.isLinux) {
    // Add from here...
    if (Platform.environment.containsKey('FLUTTER_TEST')) {
      return ffi.DynamicLibrary.open(
          'build/linux/x64/debug/bundle/lib/lib$_libName.so');
    }
    // ...to here.
    return ffi.DynamicLibrary.open('lib$_libName.so');
  }
  if (Platform.isWindows) {
    // Add from here...
    if (Platform.environment.containsKey('FLUTTER_TEST')) {
      return ffi.DynamicLibrary.open(p.canonicalize(
          p.join(r'build\windows\runner\Debug', '$_libName.dll')));
    }
    // ...to here.
    return ffi.DynamicLibrary.open('$_libName.dll');
  }
  throw UnsupportedError('Unknown platform: ${Platform.operatingSystem}');
}();

class LibTorrent {
  // singleton
  static final LibTorrent _instance = LibTorrent._internal();

  factory LibTorrent() => _instance;

  LibTorrent._internal() : _torrent = TorrentBindings(_dylib);

  late final TorrentBindings _torrent;
  late final TorrentApi _torrentApi;

  TorrentApi get torrentApi {
    return _torrentApi;
  }

  void start(String dataDir) {
    final dataDirGoString = dataDir.toGoString();

    final listenPort = _torrent.Start(
      dataDirGoString,
      0,
    );

    _torrentApi =
        TorrentApi(ApiClient(basePath: 'http://localhost:$listenPort'));

    // ffi2.calloc.free(dataDirGoString.p);
  }

  Future<void> stop() async {
    _torrent.Stop();
  }
}

extension on String {
  GoString toGoString() {
    final goString = ffi2.calloc<GoString>();

    final ffi.Pointer<ffi.Char> charPtr = toNativeUtf8().cast();

    goString.ref.p = charPtr;
    goString.ref.n = length;

    return goString.ref;
  }
}
