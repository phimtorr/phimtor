import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:phimtor_app/routes/app_routes.dart';
import 'package:phimtor_app/routes/scaffold_with_nested_navigation.dart';
import 'package:phimtor_app/views/account/account_view.dart';
import 'package:phimtor_app/views/home_view.dart';
import 'package:phimtor_app/views/settings_view.dart';
import 'package:phimtor_app/views/shows/movie_detail_view.dart';
import 'package:phimtor_app/views/shows/movies_grid_view.dart';
import 'package:phimtor_app/views/shows/search_grid_view.dart';
import 'package:phimtor_app/views/shows/tv_latest_episodes_grid_view.dart';
import 'package:phimtor_app/views/shows/tv_season_detail_view.dart';
import 'package:phimtor_app/views/shows/tv_series_detail_view.dart';
import 'package:phimtor_app/views/shows/tv_series_grid_view.dart';
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
              name: AppRoutes.shows,
              path: "/shows",
              pageBuilder: (context, state) => const NoTransitionPage(
                child: HomeView(),
              ),
              routes: [
                // search
                GoRoute(
                  name: AppRoutes.showSearch,
                  path: "search/:query",
                  builder: (context, state) {
                    final query = state.pathParameters['query']!;
                    return SearchGridView(query: query);
                  },
                ),
                // movies
                GoRoute(
                  name: AppRoutes.movies,
                  path: "movies",
                  builder: (context, state) => const MoviesGridView(),
                ),
                GoRoute(
                  name: AppRoutes.movieDetails,
                  path: "movies/:id/:title",
                  builder: (context, state) {
                    final id = int.parse(state.pathParameters['id']!);
                    final title = state.pathParameters['title']!;
                    return MovieDetailView(movieId: id, title: title);
                  },
                ),
                // series
                GoRoute(
                  name: AppRoutes.tvSeries,
                  path: "series",
                  builder: (context, state) => const TVSeriesGridView(),
                ),
                GoRoute(
                  name: AppRoutes.tvSeriesDetails,
                  path: "series/:id/:title",
                  builder: (context, state) {
                    final id = int.parse(state.pathParameters['id']!);
                    final title = state.pathParameters['title']!;
                    return TVSeriesDetailView(seriesId: id, title: title);
                  },
                ),
                GoRoute(
                  name: AppRoutes.tvSeriesSeasonDetails,
                  path: "series/:id/season/:seasonNumber/:title",
                  builder: (context, state) {
                    final seriesId = int.parse(state.pathParameters['id']!);
                    final seasonNumber =
                        int.parse(state.pathParameters['seasonNumber']!);
                    final title = state.pathParameters['title']!;
                    return TVSeasonDetailView(
                      seriesId: seriesId,
                      seasonNumber: seasonNumber,
                      title: title,
                    );
                  },
                ),
                GoRoute(
                  name: AppRoutes.tvLatestEpisodes,
                  path: "latest-episodes",
                  pageBuilder: (context, state) => const NoTransitionPage(
                    child: TvLatestEpisodesGridView(),
                  ),
                ),
                // video
                GoRoute(
                  name: AppRoutes.video,
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
