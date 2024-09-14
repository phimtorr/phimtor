// private naviagtors
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/views/account/account_view.dart';
import 'package:phimtor_app/views/home_view.dart';
import 'package:phimtor_app/views/settings_view.dart';

final _rootNavigatorKey = GlobalKey<NavigatorState>();
final _shellNavigatorShowKey =
    GlobalKey<NavigatorState>(debugLabel: 'ShellNavigatorShow');
final _shellNavigtorAccountKey =
    GlobalKey<NavigatorState>(debugLabel: 'ShellNavigatorAccount');
final _shellNavigatorSettingKey =
    GlobalKey<NavigatorState>(debugLabel: 'ShellNavigatorSetting');

List<(String, Icon)> _getDestinations(BuildContext context) {
  return [
    (context.loc.home, const Icon(Icons.home)),
    (context.loc.account, const Icon(Icons.account_circle)),
    (context.loc.setting, const Icon(Icons.settings)),
  ];
}

final goRouter = GoRouter(
  initialLocation: '/shows',
  // * Passing a navigatorKey causes an issue on hot reload:
  // * https://github.com/flutter/flutter/issues/113757#issuecomment-1518421380
  // * However it's still necessary otherwise the navigator pops back to
  // * root on hot reload
  navigatorKey: _rootNavigatorKey,
  debugLogDiagnostics: true,
  routes: [
    // Stateful navigation based on:
    // https://github.com/flutter/packages/blob/main/packages/go_router/example/lib/stateful_shell_route.dart
    StatefulShellRoute.indexedStack(
      builder: (
        BuildContext context,
        GoRouterState state,
        StatefulNavigationShell navigationShell,
      ) {
        return ScaffoldWithNestedNavigation(navigationShell: navigationShell);
      },
      branches: [
        StatefulShellBranch(
          navigatorKey: _shellNavigatorShowKey,
          routes: [
            GoRoute(
              path: "/shows",
              pageBuilder: (context, state) => const NoTransitionPage(
                child: HomeView(),
              ),
            )
          ],
        ),
        StatefulShellBranch(
          navigatorKey: _shellNavigtorAccountKey,
          routes: [
            GoRoute(
              path: "/accounts",
              pageBuilder: (context, state) => const NoTransitionPage(
                child: AccountView(),
              ),
            )
          ],
        ),
        StatefulShellBranch(
          navigatorKey: _shellNavigatorSettingKey,
          routes: [
            GoRoute(
              path: "/settings",
              pageBuilder: (context, state) => const NoTransitionPage(
                child: SettingsView(),
              ),
            )
          ],
        ),
      ],
    ),
  ],
);

class ScaffoldWithNestedNavigation extends StatelessWidget {
  const ScaffoldWithNestedNavigation({
    Key? key,
    required this.navigationShell,
  }) : super(
            key: key ?? const ValueKey<String>('ScaffoldWithNestedNavigation'));

  final StatefulNavigationShell navigationShell;

  void _goBranch(int index) {
    navigationShell.goBranch(
      index,
      // A common pattern when using bottom navigation bars is to support
      // navigating to the initial location when tapping the item that is
      // already active. This example demonstrates how to support this behavior,
      // using the initialLocation parameter of goBranch.
      initialLocation: index == navigationShell.currentIndex,
    );
  }

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(builder: (context, constrainst) {
      if (constrainst.maxWidth < 450) {
        return ScaffoldWithNavigationBar(
          body: navigationShell,
          selectedIndex: navigationShell.currentIndex,
          onDestinationSelected: _goBranch,
        );
      } else {
        return ScaffoldWithNavigationRail(
          body: navigationShell,
          selectedIndex: navigationShell.currentIndex,
          onDestinationSelected: _goBranch,
          extended: constrainst.maxWidth > 600,
        );
      }
    });
  }
}

class ScaffoldWithNavigationBar extends StatelessWidget {
  const ScaffoldWithNavigationBar({
    super.key,
    required this.body,
    required this.selectedIndex,
    required this.onDestinationSelected,
  });

  final Widget body;
  final int selectedIndex;
  final ValueChanged<int> onDestinationSelected;

  @override
  Widget build(BuildContext context) {
    final destinations = _getDestinations(context)
        .map((e) => NavigationDestination(icon: e.$2, label: e.$1))
        .toList();
    return Scaffold(
      body: body,
      bottomNavigationBar: NavigationBar(
        selectedIndex: selectedIndex,
        destinations: destinations,
        onDestinationSelected: onDestinationSelected,
      ),
    );
  }
}

class ScaffoldWithNavigationRail extends StatelessWidget {
  const ScaffoldWithNavigationRail({
    super.key,
    required this.body,
    required this.selectedIndex,
    required this.onDestinationSelected,
    this.extended = true,
  });

  final Widget body;
  final int selectedIndex;
  final ValueChanged<int> onDestinationSelected;
  final bool extended;

  @override
  Widget build(BuildContext context) {
    final destinations = _getDestinations(context)
        .map((e) => NavigationRailDestination(icon: e.$2, label: Text(e.$1)))
        .toList();
    return Scaffold(
      body: Row(
        children: [
          NavigationRail(
            extended: extended,
            selectedIndex: selectedIndex,
            destinations: destinations,
            onDestinationSelected: onDestinationSelected,
          ),
          const VerticalDivider(thickness: 1, width: 1),
          Expanded(
            child: body,
          ),
        ],
      ),
    );
  }
}
