import 'package:flutter/material.dart';

class ShowComponents {
  static Widget buildTagline(BuildContext context, String title) {
    final backgroundColor =
        Theme.of(context).colorScheme.surfaceContainerLow.withOpacity(0.6);
    final textStyle = Theme.of(context).textTheme.bodyMedium!.merge(
          TextStyle(
            color: Theme.of(context).colorScheme.onSurface,
            fontStyle: FontStyle.italic,
          ),
        );
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 4, vertical: 4),
      decoration: BoxDecoration(
        color: backgroundColor,
        borderRadius: BorderRadius.circular(4),
      ),
      child: Text(
        title,
        style: textStyle,
      ),
    );
  }

}