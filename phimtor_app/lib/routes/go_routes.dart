import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:phimtor_app/routes/route_names.dart';
import 'package:phimtor_app/routes/scaffold_with_nested_navigation.dart';
import 'package:phimtor_app/views/account/account_view.dart';
import 'package:phimtor_app/views/home_view.dart';
import 'package:phimtor_app/views/settings_view.dart';
import 'package:phimtor_app/views/shows/movie_detail_view.dart';
import 'package:phimtor_app/views/shows/movies_grid_view.dart';
import 'package:phimtor_app/views/shows/search_grid_view.dart';
import 'package:phimtor_app/views/shows/series_detail_view.dart';
import 'package:phimtor_app/views/shows/series_grid_view.dart';
import 'package:phimtor_app/views/videos/video_view.dart';

// private naviagtors
final _rootNavigatorKey = GlobalKey<NavigatorState>();
final _shellNavigatorShowKey =
    GlobalKey<NavigatorState>(debugLabel: 'ShellNavigatorShow');
final _shellNavigtorAccountKey =
    GlobalKey<NavigatorState>(debugLabel: 'ShellNavigatorAccount');
final _shellNavigatorSettingKey =
    GlobalKey<NavigatorState>(debugLabel: 'ShellNavigatorSetting');

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
              name: routeNameShows,
              path: "/shows",
              pageBuilder: (context, state) => const NoTransitionPage(
                child: HomeView(),
              ),
              routes: [
                // search
                GoRoute(
                  name: routeNameShowSearch,
                  path: "search/:query",
                  builder: (context, state) {
                    final query = state.pathParameters['query']!;
                    return SearchGridView(query: query);
                  },
                ),
                // movies
                GoRoute(
                  name: routeNameMovies,
                  path: "movies",
                  builder: (context, state) => const MoviesGridView(),
                ),
                GoRoute(
                  name: routeNameMovieDetails,
                  path: "movies/:id/:title",
                  builder: (context, state) {
                    final id = int.parse(state.pathParameters['id']!);
                    final title = state.pathParameters['title']!;
                    return MovieDetailView(movieId: id, title: title);
                  },
                ),
                // series
                GoRoute(
                  name: routeNameSeries,
                  path: "series",
                  builder: (context, state) => const SeriesGridView(),
                ),
                GoRoute(
                  name: routeNameSeriesDetails,
                  path: "series/:id/:title",
                  builder: (context, state) {
                    final id = int.parse(state.pathParameters['id']!);
                    final title = state.pathParameters['title']!;
                    return SeriesDetailView(seriesId: id, title: title);
                  },
                ),
                // video
                GoRoute(
                  name: routeNameVideo,
                  path: "video/:id/:title",
                  builder: (context, state) {
                    final id = int.parse(state.pathParameters['id']!);
                    final title = state.pathParameters['title']!;
                    return VideoView(videoId: id, title: title);
                  },
                )
              ],
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
