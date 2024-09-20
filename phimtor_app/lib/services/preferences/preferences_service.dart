import 'dart:async';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:path_provider/path_provider.dart';
import 'package:phimtor_app/utilities/platform/platform_utilities.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:path/path.dart' as p;

class PreferencesService {
  static PreferencesService? _instance;
  final SharedPreferences _prefs;

  static Future<void> ensureInitialized() async {
    if (_instance != null) {
      throw Exception('PreferencesRepository already initialized');
    }
    _instance = PreferencesService._(await SharedPreferences.getInstance());

    await _instance!._ensureDataDirPathIsExisted();
  }

  PreferencesService._(this._prefs);

  static PreferencesService getInstance() {
    if (_instance == null) {
      throw Exception('PreferencesRepository not initialized');
    }
    return _instance!;
  }

  static const _keyDataDirPath = 'data_dir_path';
  static const _keyDeleteAfterClose = 'delete_after_close';
  static const _keyLocale = 'locale';

  Future<void> _ensureDataDirPathIsExisted() async {
    final dataDirPath = _prefs.getString(_keyDataDirPath);
    if (dataDirPath != null) {
      final dataDir = Directory(dataDirPath);
      final isExisted = await dataDir.exists();
      if (!isExisted) {
        await dataDir.create(recursive: true);
      }
    } else {
      if (PlatformUtilities.isDesktop) {
        final dir = await getDownloadsDirectory();
        if (dir == null) {
          throw Exception('Could not get downloads directory');
        }
        final path = dir.path;
        final dataDirPath = p.join(path, 'PhimTor');

        final dataDir = Directory(dataDirPath);
        final isExisted = await dataDir.exists();
        if (!isExisted) {
          await dataDir.create(recursive: true);
        }

        await _prefs.setString(_keyDataDirPath, dataDirPath);
      } else if (Platform.isAndroid) {
        final dir = await getTemporaryDirectory();
        final path = dir.path;
        final dataDirPath = p.join(path, 'PhimTor');

        final dataDir = Directory(dataDirPath);
        final isExisted = await dataDir.exists();
        if (!isExisted) {
          await dataDir.create(recursive: true);
        }

        await _prefs.setString(_keyDataDirPath, dataDirPath);
      } else {
        throw UnsupportedError('Not support dataDir on this platform');
      }
    }
  }

  Future<void> setDataDirPath(String path) async {
    final dataDir = Directory(path);
    final isExisted = await dataDir.exists();
    if (!isExisted) {
      await dataDir.create(recursive: true);
    }

    await _prefs.setString(_keyDataDirPath, path);
  }

  String get dataDirPath {
    final dataDirPath = _prefs.getString(_keyDataDirPath);
    if (dataDirPath == null) {
      throw Exception('Data directory path not set');
    }
    return dataDirPath;
  }

  Future<void> setDeleteAfterClose(bool value) async {
    await _prefs.setBool(_keyDeleteAfterClose, value);
  }

  bool get deleteAfterClose {
    // if mobile, always delete after close
    if (PlatformUtilities.isMobile) {
      return true;
    }
    return _prefs.getBool(_keyDeleteAfterClose) ?? true;
  }

  Future<void> setLocale(Locale locale) async {
    await _prefs.setString(_keyLocale, locale.languageCode);
  }

  Locale get locale {
    final code = _prefs.getString(_keyLocale) ?? 'vi';
    return Locale(code);
  }
}
