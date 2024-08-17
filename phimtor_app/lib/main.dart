import 'package:flutter/material.dart';
import 'package:media_kit/media_kit.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/lifecycle_manager.dart';
import 'package:phimtor_app/locale_provider.dart';
import 'package:phimtor_app/views/home_view.dart';
import 'package:phimtor_app/views/settings_view.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:provider/provider.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  MediaKit.ensureInitialized();
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
          return MaterialApp(
            supportedLocales: AppLocalizations.supportedLocales,
            localizationsDelegates: AppLocalizations.localizationsDelegates,
            locale: provider.locale,
            title: "Phim Tor",
            theme: ThemeData(
              colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepOrange),
              useMaterial3: true,
            ),
            home: const MyHomePage(),
          );
        }),
      ),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  var _selectedIndex = 0;

  @override
  Widget build(BuildContext context) {
    Widget page;
    final selectIndex = _selectedIndex;
    switch (selectIndex) {
      case 0:
        page = const HomeView();
      case 1:
        page = const SettingsView();
      default:
        throw UnimplementedError("no widget for index $selectIndex");
    }

    return LayoutBuilder(builder: (context, constraints) {
      return Row(
        children: [
          SafeArea(
            child: NavigationRail(
              extended: constraints.maxWidth > 600,
              destinations: [
                NavigationRailDestination(
                  icon: const Icon(Icons.home),
                  label: Text(context.loc.home),
                ),
                NavigationRailDestination(
                  icon: const Icon(Icons.settings),
                  label: Text(context.loc.setting),
                ),
              ],
              selectedIndex: selectIndex,
              onDestinationSelected: (value) {
                setState(() {
                  _selectedIndex = value;
                });
              },
              elevation: 5,
              backgroundColor:
                  Theme.of(context).colorScheme.surfaceContainerHigh,
            ),
          ),
          Expanded(
            child: page,
          )
        ],
      );
    });
  }
}
