import 'package:flutter/material.dart';
import 'package:phimtor_app/services/auth/auth_service.dart';
import 'package:phimtor_app/utilities/dialogs/need_login_dialog.dart';

class NeedVerifiedUserButton extends StatelessWidget {
  const NeedVerifiedUserButton({
    super.key,
    required this.onPressed,
    required this.child,
  });

  final Widget child;
  final VoidCallback? onPressed;

  void hanldePressed(BuildContext context) async {
    if (!AuthService().isVerifiedUser) {
      await showNeedLoginDialog(context);
      return;
    }

    onPressed!.call();
  }

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: onPressed == null ? null : () => hanldePressed(context),
      child: child,
    );
  }
}
