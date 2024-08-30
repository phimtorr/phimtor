import 'package:flutter/material.dart';
import 'package:phimtor_app/services/auth/auth_service.dart';
import 'package:phimtor_app/utilities/dialogs/upgrade_to_preimum_dialog.dart';

class PremiumButton extends StatelessWidget {
  const PremiumButton({
    super.key,
    required this.onPressed,
    required this.label,
    this.icon,
  });

  final VoidCallback? onPressed;
  final Widget label;
  final Widget? icon;

  bool get isPremium => AuthService().isPremiumUser;

  @override
  Widget build(BuildContext context) {
    if (!isPremium) {
      return ElevatedButton.icon(
        onPressed: onPressed == null
            ? null
            : () async {
                await showUpgradeToPremiumDialog(context);
              },
        label: label,
        icon: const Icon(Icons.star),
      );
    }

    return ElevatedButton.icon(
      onPressed: onPressed,
      label: label,
      icon: icon,
    );
  }
}
