import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/routes/app_routes.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart'
    as phimtor_service;
import 'package:phimtor_app/views/search_section.dart';
import 'package:phimtor_app/views/shows/shows_list.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

class HomeView extends StatelessWidget {
  const HomeView({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(context.loc.title),
      ),
      body: const SingleChildScrollView(
        child: Padding(
          padding: EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              SearchSection(),
              SizedBox(height: 32),
              LatestAddedMoviesSection(),
              SizedBox(height: 32),
              MoviesSection(),
              SizedBox(height: 32),
              TVSeriesSection(),
              SizedBox(height: 32),
              TVEpisodesSection(),
            ],
          ),
        ),
      ),
    );
  }
}

class LatestAddedMoviesSection extends StatelessWidget {
  const LatestAddedMoviesSection({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        buildHeadline(context, context.loc.latest_added_movies,
            AppRoutes.latestAddedMovies),
        const SizedBox(height: 16),
        buildShowList(context, (
          int page,
          int pageSize,
        ) async {
          final resp = await phimtor_service.PhimtorService()
              .defaultApi
              .listRecentlyAddedMovies(
                page: page,
                pageSize: pageSize,
              );
          if (resp == null) {
            throw Exception("Null response");
          }
          return resp.movies;
        }),
      ],
    );
  }
}

class MoviesSection extends StatelessWidget {
  const MoviesSection({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        buildHeadline(context, context.loc.latest_movies, AppRoutes.movies),
        const SizedBox(height: 16),
        buildShowList(context, (
          int page,
          int pageSize,
        ) async {
          final resp = await phimtor_service.PhimtorService()
              .defaultApi
              .listLatestMovies(
                page: page,
                pageSize: pageSize,
              );
          if (resp == null) {
            throw Exception("Null response");
          }
          return resp.movies;
        }),
      ],
    );
  }
}

class TVSeriesSection extends StatelessWidget {
  const TVSeriesSection({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        buildHeadline(
            context, context.loc.latest_tv_series, AppRoutes.tvSeries),
        const SizedBox(height: 16),
        buildShowList(context, (
          int page,
          int pageSize,
        ) async {
          final resp = await phimtor_service.PhimtorService()
              .defaultApi
              .listLatestTvSeries(
                page: page,
                pageSize: pageSize,
              );
          if (resp == null) {
            throw Exception("Null response");
          }
          return resp.tvSeries;
        }),
      ],
    );
  }
}

class TVEpisodesSection extends StatelessWidget {
  const TVEpisodesSection({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        buildHeadline(
            context, context.loc.latest_episodes, AppRoutes.tvLatestEpisodes),
        const SizedBox(height: 16),
        buildShowList(context, (
          int page,
          int pageSize,
        ) async {
          final resp = await phimtor_service.PhimtorService()
              .defaultApi
              .listLatestEpisodes(
                page: page,
                pageSize: pageSize,
              );
          if (resp == null) {
            throw Exception("Null response");
          }
          return resp.episodes;
        }),
      ],
    );
  }
}

Widget buildHeadline(
    BuildContext context, String title, String loadMoreRouteName) {
  return LayoutBuilder(builder: (context, constraints) {
    final isWideScreen = constraints.maxWidth > 600;
    if (isWideScreen) {
      return Row(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Text(
            title,
            style: Theme.of(context).textTheme.headlineLarge,
          ),
          const SizedBox(width: 8),
          ElevatedButton.icon(
            onPressed: () async {
              await context.pushNamed(loadMoreRouteName);
            },
            label: Text(context.loc.load_more),
            icon: const Icon(Icons.arrow_forward),
            iconAlignment: IconAlignment.end,
          ),
        ],
      );
    } else {
      return Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            title,
            style: Theme.of(context).textTheme.headlineLarge,
          ),
          const SizedBox(height: 8),
          ElevatedButton.icon(
            onPressed: () async {
              await context.pushNamed(loadMoreRouteName);
            },
            label: Text(context.loc.load_more),
            icon: const Icon(Icons.arrow_forward),
            iconAlignment: IconAlignment.end,
          ),
        ],
      );
    }
  });
}

typedef ListShowsFunction = Future<List<phimtor_api.ModelShow>> Function(
  int page,
  int pageSize,
);

Widget buildShowList(BuildContext context, ListShowsFunction listShows) {
  return SizedBox(
    height: 320.0,
    child: FutureBuilder(
      future: listShows(1, 10),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return const Center(child: CircularProgressIndicator());
        }
        if (snapshot.hasError) {
          return Center(
              child: Text(context.loc.error(snapshot.error.toString())));
        }

        final shows = snapshot.data ?? [];
        return ShowsList(shows: shows);
      },
    ),
  );
}
