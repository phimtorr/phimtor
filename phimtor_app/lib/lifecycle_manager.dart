import 'dart:developer';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:phimtor_app/services/preferences/preferences_service.dart';
import 'package:torrent/torrent.dart' as torrent;

class LifecycleManager extends StatefulWidget {
  final Widget child;
  const LifecycleManager({super.key, required this.child});

  @override
  State<LifecycleManager> createState() => _LifecycleManagerState();
}

class _LifecycleManagerState extends State<LifecycleManager> {
  bool _isInitialized = false;
  Exception? _error;
  @override
  void initState() {
    Future.delayed(Duration.zero, () async {
      await PreferencesService.ensureInitialized();
      final dataDirPath = PreferencesService.getInstance().dataDirPath;
      log("Starting libtorrent with dataDirPath: $dataDirPath");
      torrent.LibTorrent().start(dataDirPath);
    }).whenComplete(() {
      setState(() {
        _isInitialized = true;
      });
    }).catchError((error) {
      setState(() {
        _error = error as Exception;
      });
    });

    super.initState();
  }

  @override
  void dispose() {
    log("Stopping libtorrent");
    torrent.LibTorrent().stop();
    _deleteDataDirIfNeed();
    super.dispose();
  }

  void _deleteDataDirIfNeed() {
    final deleteAfterClose = PreferencesService.getInstance().deleteAfterClose;
    if (deleteAfterClose) {
      log("Deleting dataDir");
      final dataDirPath = PreferencesService.getInstance().dataDirPath;
      final dataDir = Directory(dataDirPath);
      dataDir.deleteSync(recursive: true);
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_error != null) {
      return Center(
        child: Text('Error: $_error'),
      );
    }
    if (!_isInitialized) {
      return const CircularProgressIndicator();
    }
    return widget.child;
  }
}
