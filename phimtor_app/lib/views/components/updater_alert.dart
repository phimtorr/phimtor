import 'dart:developer' show log;

import 'package:flutter/material.dart';
import 'package:phimtor_app/services/updater/updater_service.dart';
import 'package:phimtor_app/services/updater/updater_version.dart';
import 'package:phimtor_app/utilities/dialogs/updater_dialog.dart';

class UpdaterAlert extends StatefulWidget {
  UpdaterAlert({
    super.key,
    this.navigatorKey,
    this.child,
  });

  final UpdaterService updaterService = UpdaterService();

  final GlobalKey<NavigatorState>? navigatorKey;

  final Widget? child;

  @override
  State<UpdaterAlert> createState() => _UpdaterAlertState();
}

class _UpdaterAlertState extends State<UpdaterAlert> {
  /// Is the alert dialog being displayed right now?
  bool displayed = false;

  @override
  void initState() {
    super.initState();

    log("Initializing updater service");
    UpdaterService().initialize();
  }

  @override
  void dispose() {
    log("Closing updater service");
    UpdaterService().close();

    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return StreamBuilder(
      stream: widget.updaterService.versionStream,
      builder: (context, snapshot) {
        if  ((snapshot.connectionState == ConnectionState.waiting ||
                snapshot.connectionState == ConnectionState.active) && 
                (snapshot.hasData)) {

          final checkContext = widget.navigatorKey?.currentContext ?? context;
          final newVersion = snapshot.data;

          if (newVersion != null && !displayed) {
            alertNewVersion(checkContext, newVersion);
          }
          
        }
        return widget.child ?? const SizedBox.shrink();
      },
    );
  }

  void alertNewVersion(BuildContext context, UpdaterVersion version) {
    if (displayed) {
      return;
    }

    displayed = true;
    Future.delayed(Duration.zero, () async {
      await showUpdaterDialog(context, version); // ignore: use_build_context_synchronously
      displayed = false;
    });
  }
    
}
