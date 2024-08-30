import 'package:flutter/material.dart';
import 'package:phimtor_app/services/auth/auth_service.dart';
import 'package:phimtor_app/utilities/dialogs/need_login_dialog.dart';

class NeedVerifiedUserButton extends StatelessWidget {
  const NeedVerifiedUserButton({
    super.key,
    required this.onPressed,
    required this.child,
    this.icon,
  });

  final VoidCallback? onPressed;
  final Widget child;
  final Widget? icon;

  void hanldePressed(BuildContext context) async {
    if (!AuthService().isVerifiedUser) {
      await showNeedLoginDialog(context);
      return;
    }

    onPressed!.call();
  }

  bool get isVerifired => AuthService().isVerifiedUser;

  @override
  Widget build(BuildContext context) {
    if (!isVerifired) {
      return ElevatedButton.icon(
        onPressed: onPressed == null
            ? null
            : () async {
                await showNeedLoginDialog(context);
              },
        label: child,
        icon: const Icon(Icons.lock_outline),
      );
    }
    return ElevatedButton.icon(
      onPressed: onPressed == null ? null : () => hanldePressed(context),
      label: child,
      icon: icon,
    );
  }
}
