import 'package:flutter/material.dart';
import 'package:phimtor_app/constants/enviroment_vars.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/updater/updater_service.dart';
import 'package:phimtor_app/services/updater/updater_version.dart';
import 'package:phimtor_app/utilities/dialogs/updater_dialog.dart';

class VersionWidget extends StatelessWidget {
  const VersionWidget({super.key});

  Future<void> alertNewVesion(
      BuildContext context, UpdaterVersion version) async {
    await showUpdaterDialog(context, version);
  }

  @override
  Widget build(BuildContext context) {
    final currentVersionDisplay = Text(
      Constants.appVersion,
      style: Theme.of(context)
          .textTheme
          .labelSmall
          ?.merge(const TextStyle(fontStyle: FontStyle.italic)),
    );
    return StreamBuilder<UpdaterVersion>(
      stream: UpdaterService().versionStream,
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          final version = snapshot.data!;

          Future.delayed(Duration.zero, () {
            // ignore: use_build_context_synchronously
            alertNewVesion(context, version);
          });

          return Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              OutlinedButton.icon(
                onPressed: () async {
                  alertNewVesion(context, version);
                },
                icon: const Icon(Icons.update),
                label: Text(
                  context.loc.has_new_version,
                  style: Theme.of(context).textTheme.labelSmall,
                ),
              ),
              currentVersionDisplay,
            ],
          );
        }
        return currentVersionDisplay;
      },
    );
  }
}
