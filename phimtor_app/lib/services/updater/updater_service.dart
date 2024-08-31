import 'dart:async';
import 'dart:io';

import 'package:http/http.dart' as http;

import 'package:phimtor_app/constants/enviroment_vars.dart';
import 'package:phimtor_app/services/updater/updater_version.dart';

class UpdaterService {
  // singleton
  static final UpdaterService _instance = UpdaterService._internal();
  factory UpdaterService() => _instance;
  UpdaterService._internal() {
    _versionController = StreamController<UpdaterVersion>.broadcast(
      onListen: () {
        if (_version != null) {
          _versionController.add(_version!);
        }
      },
    );
  }

  UpdaterVersion? _version;

  late final StreamController<UpdaterVersion> _versionController;

  Stream<UpdaterVersion> get versionStream => _versionController.stream;

  bool _isInitialized = false;
  Timer? _timer;

  void initialize() {
    if (_isInitialized) {
      return;
    }
    _isInitialized = true;
    _timer = Timer.periodic(const Duration(minutes: 10), (timer) async {
      await checkForUpdate();
    });

    Future.delayed(Duration.zero, () async {
      await checkForUpdate();
    });
  }

  void close() {
    _timer?.cancel();
    _versionController.close();
  }

  Future<void> checkForUpdate() async {
    // simulate network request
    await Future.delayed(const Duration(seconds: 2));
    const url = "${Constants.apiUrl}/public/phimtor-app/VERSION.txt";
    final response = await http.get(Uri.parse(url));

    if (response.statusCode != 200) {
      return;
    }

    final versionValue = response.body.toString().trim();

    if (Constants.appVersion.toString().trim() == versionValue.toLowerCase().trim()) {
      return;
    }

    final version = UpdaterVersion(
      version: versionValue,
      binaryUrl: binaryUrl,
    );

    _version = version;
    _versionController.add(version);
  }
}

Uri get binaryUrl {
  if (Platform.isLinux) {
    return Uri.parse(
        "${Constants.apiUrl}/public/phimtor-app/PhimTor-x86_64.AppImage");
  }
  if (Platform.isWindows) {
    return Uri.parse("${Constants.apiUrl}/public/phimtor-app/PhimTorSetup.exe");
  }
  if (Platform.isMacOS) {
    return Uri.parse("${Constants.apiUrl}/public/phimtor-app/PhimTor.dmg");
  }

  throw UnsupportedError("Unsupported platform");
}
