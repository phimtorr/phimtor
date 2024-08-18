import 'dart:developer';
import 'dart:io';
import 'dart:ui';

import 'package:flutter/material.dart';
import 'package:phimtor_app/services/auth/auth_service.dart';
import 'package:phimtor_app/services/preferences/preferences_service.dart';
import 'package:torrent/torrent.dart' as torrent;

class LifecycleManager extends StatefulWidget {
  final Widget child;
  const LifecycleManager({super.key, required this.child});

  @override
  State<LifecycleManager> createState() => _LifecycleManagerState();
}

class _LifecycleManagerState extends State<LifecycleManager>
    with WidgetsBindingObserver {
  late final String _dataDirPath;
  bool _isInitialized = false;
  Exception? _error;
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addObserver(this);

    Future.delayed(Duration.zero, initServices).whenComplete(() {
      if (mounted) {
        setState(() {
          _isInitialized = true;
        });
      }
    }).catchError((error) {
      log("Error: $error");
      if (mounted) {
        setState(() {
          _error = error;
        });
      }
    });
  }

  @override
  void dispose() {
    WidgetsBinding.instance.removeObserver(this);
    super.dispose();
  }

  @override
  Future<AppExitResponse> didRequestAppExit() async {
    cleanUp();
    return super.didRequestAppExit();
  }

  Future<void> initServices() async {
    log("Initializing preferences service");
    await PreferencesService.ensureInitialized();

     log("Initializing auth service");
    await AuthService().initialize();

    _dataDirPath = PreferencesService.getInstance().dataDirPath;
    log("Starting libtorrent with dataDirPath: $_dataDirPath");
    torrent.LibTorrent().start(_dataDirPath);
  }

  void cleanUp() {
    log("Stopping libtorrent");
    torrent.LibTorrent().stop();
    _deleteDataDirIfNeed();
  }

  void _deleteDataDirIfNeed() {
    final deleteAfterClose = PreferencesService.getInstance().deleteAfterClose;
    if (!deleteAfterClose) {
      return;
    }
    
    log("Deleting dataDir");
    final dataDir = Directory(_dataDirPath);
    final contents = dataDir.listSync();
    for (final entity in contents) {
      try {
        if (entity is File) {
          entity.deleteSync();
        } else if (entity is Directory) {
          entity.deleteSync(recursive: true);
        }
      } catch (e) {
        log("Failed to delete entity: $entity, error: $e");
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_error != null) {
      return Center(
        child: Text('Error: ${_error.toString()}'),
      );
    }
    if (!_isInitialized) {
      return const CircularProgressIndicator();
    }
    return widget.child;
  }
}
