import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/utilities/dialogs/generic_dialog.dart';

Future<void> showNeedLoginDialog(BuildContext context) async {
  return showGenericDialog<void>(
    context: context,
    title: context.loc.dialog_need_login_title,
    content: context.loc.dialog_need_login_message,
    optionsBuilder: () {
      return {
        context.loc.close: null,
      };
    },
  );
}
