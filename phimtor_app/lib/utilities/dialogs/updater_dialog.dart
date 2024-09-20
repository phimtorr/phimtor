import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/updater/updater_version.dart';
import 'package:phimtor_app/utilities/dialogs/generic_dialog.dart';
import 'package:url_launcher/url_launcher.dart';

Future<void> showUpdaterDialog(
    BuildContext context, UpdaterVersion newVersion) async {
  // show dialog
  final ok = await showGenericDialog<bool>(
    context: context,
    title: context.loc.version_update_title,
    content: context.loc.version_update_message(newVersion.version),
    optionsBuilder: () {
      return {
        context.loc.close: null,
        context.loc.download: true,
      };
    },
  );
  if (ok == true) {
    launchUrl(newVersion.binaryUrl);
  }
}
