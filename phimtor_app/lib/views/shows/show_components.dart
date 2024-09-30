import 'package:flutter/material.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

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
    return buildOverlayLabel(
      context,
      title,
      backgroundColor: backgroundColor,
      textStyle: textStyle,
    );
  }

  static Widget buildOverlayLabel(
    BuildContext context,
    String title, {
    Color? backgroundColor,
    TextStyle? textStyle,
  }) {
    backgroundColor ??=
        Theme.of(context).colorScheme.surfaceContainerLow.withOpacity(0.6);
    textStyle ??= Theme.of(context).textTheme.labelSmall!.merge(
          TextStyle(
            color: Theme.of(context).colorScheme.onSurface,
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

  static Widget buildLable(
    BuildContext context,
    String title, {
    Color? backgroundColor,
    TextStyle? textStyle,
  }) {
    backgroundColor ??=
        Theme.of(context).colorScheme.surfaceContainerHigh;
    textStyle ??= Theme.of(context).textTheme.labelSmall!.merge(
          TextStyle(
            color: Theme.of(context).colorScheme.onSurface,
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

  static Widget buildGenres(
    BuildContext context,
    List<phimtor_api.Genre> genres,
  ) {
    return Wrap(
      spacing: 4,
      runSpacing: 4,
      children: genres
          .map(
            (e) => ShowComponents.buildLable(
              context,
              e.name,
            ),
          )
          .toList(),
    );
  }

  static String formatReleaseDate(DateTime releaseDate) {
    return "${releaseDate.year.toString()}-${releaseDate.month.toString().padLeft(2, '0')}-${releaseDate.day.toString().padLeft(2, '0')}";
  }
}
