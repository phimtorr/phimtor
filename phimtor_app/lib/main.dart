import 'dart:developer' show log;

import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:media_kit/media_kit.dart';
import 'package:phimtor_app/lifecycle_manager.dart';
import 'package:phimtor_app/locale_provider.dart';
import 'package:phimtor_app/routes/go_routes.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/services/auth/auth_service.dart';
import 'package:phimtor_app/services/preferences/preferences_service.dart';
import 'package:provider/provider.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  MediaKit.ensureInitialized();

  log("Initializing preferences service");
  await PreferencesService.ensureInitialized();

  try {
    log("Initializing auth service");
    await AuthService().initialize();
  } catch (e) {
    log("Error initializing auth service: $e");
  }

  log("Initializing analytics service");
  await AnalyticsService().initialize(
    userId: AuthService().currentUser?.uid,
  );

  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return LifecycleManager(
      child: ChangeNotifierProvider(
        create: (context) => LocaleProvider(),
        child: Builder(builder: (context) {
          final provider = Provider.of<LocaleProvider>(context);
          return MaterialApp.router(
            routerConfig: goRouter,
            debugShowCheckedModeBanner: false,
            supportedLocales: AppLocalizations.supportedLocales,
            localizationsDelegates: AppLocalizations.localizationsDelegates,
            locale: provider.locale,
            theme: ThemeData(
              colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepOrange),
              useMaterial3: true,
            ),
          );
        }),
      ),
    );
  }
}
