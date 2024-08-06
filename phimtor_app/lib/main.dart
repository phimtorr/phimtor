import 'package:flutter/material.dart';
import 'package:phimtor_app/lifecycle_manager.dart';
import 'package:phimtor_app/views/home_view.dart';
import 'package:phimtor_app/views/settings_view.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return LifecycleManager(
      child: MaterialApp(
        title: 'Phim Tor',
        theme: ThemeData(
          colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepOrange),
          useMaterial3: true,
        ),
        home: const MyHomePage(),
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
              destinations: const [
                NavigationRailDestination(
                  icon: Icon(Icons.home),
                  label: Text("Home"),
                ),
                NavigationRailDestination(
                  icon: Icon(Icons.settings),
                  label: Text("Setting"),
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
