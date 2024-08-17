import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:phimtor_app/services/preferences/preferences_service.dart';

class LocaleProvider extends ChangeNotifier {
  Locale _locale = PreferencesService.getInstance().locale;

  Locale get locale => _locale;

  void setLocale(Locale locale) {
    if (!AppLocalizations.supportedLocales
        .map((e) => e.languageCode)
        .toList()
        .contains(locale.languageCode)) {
      return;
    }
    if (locale == _locale) {
      return;
    }

    PreferencesService.getInstance().setLocale(locale);
    _locale = locale;
    notifyListeners();
  }
}
