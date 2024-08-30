import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/utilities/dialogs/generic_dialog.dart';

Future<void> showUpgradeToPremiumDialog(BuildContext context) async {
  return showGenericDialog<void>(
    context: context,
    title: context.loc.dialog_upgrade_to_premium_title,
    content: context.loc.dialog_upgrade_to_premium_message,
    optionsBuilder: () {
      return {
        context.loc.close: null,
      };
    },
  );
}
