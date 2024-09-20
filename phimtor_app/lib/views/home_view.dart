import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/routes/app_routes.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart'
    as phimtor_service;
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

class SearchSection extends StatefulWidget {
  const SearchSection({super.key});

  @override
  State<SearchSection> createState() => _SearchSectionState();
}

class _SearchSectionState extends State<SearchSection> {
  final _searchController = TextEditingController();

  @override
  void dispose() {
    _searchController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          context.loc.search,
          style: Theme.of(context).textTheme.headlineLarge,
        ),
        const SizedBox(height: 16),
        // a search text box with a button to search
        CupertinoSearchTextField(
          controller: _searchController,
          onSubmitted: (query) async {
            if (query.isEmpty) {
              return;
            }
            await context.pushNamed(
              AppRoutes.showSearch,
              pathParameters: {"query": query},
            );
          },
        ),
      ],
    );
  }
}

class LatestAddedMoviesSection extends StatelessWidget {
  const LatestAddedMoviesSection({super.key});

  Future<List<phimtor_api.ModelShow>> listShows(
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
  }

  @override
  Widget build(BuildContext context) {
    return buildListShowsSection(
      context,
      context.loc.latest_added_movies,
      AppRoutes.latestAddedMovies,
      listShows,
    );
  }
}

class MoviesSection extends StatelessWidget {
  const MoviesSection({super.key});

  Future<List<phimtor_api.ModelShow>> listShows(
    int page,
    int pageSize,
  ) async {
    final resp =
        await phimtor_service.PhimtorService().defaultApi.listLatestMovies(
              page: page,
              pageSize: pageSize,
            );
    if (resp == null) {
      throw Exception("Null response");
    }
    return resp.movies;
  }

  @override
  Widget build(BuildContext context) {
    return buildListShowsSection(
      context,
      context.loc.latest_movies,
      AppRoutes.movies,
      listShows,
    );
  }
}

class TVSeriesSection extends StatelessWidget {
  const TVSeriesSection({super.key});

  Future<List<phimtor_api.ModelShow>> listShows(
    int page,
    int pageSize,
  ) async {
    final resp =
        await phimtor_service.PhimtorService().defaultApi.listLatestTvSeries(
              page: page,
              pageSize: pageSize,
            );
    if (resp == null) {
      throw Exception("Null response");
    }
    return resp.tvSeries;
  }

  @override
  Widget build(BuildContext context) {
    return buildListShowsSection(
      context,
      context.loc.latest_tv_series,
      AppRoutes.tvSeries,
      listShows,
    );
  }
}

class TVEpisodesSection extends StatelessWidget {
  const TVEpisodesSection({super.key});
  
  Future<List<phimtor_api.ModelShow>> listShows(
    int page,
    int pageSize,
  ) async {
    final resp =
        await phimtor_service.PhimtorService().defaultApi.listLatestEpisodes(
              page: page,
              pageSize: pageSize,
            );
    if (resp == null) {
      throw Exception("Null response");
    }
    return resp.episodes;
  }

  @override
  Widget build(BuildContext context) {
    return buildListShowsSection(
      context,
      context.loc.latest_episodes,
      AppRoutes.tvLatestEpisodes,
      listShows,
    );
  }
}

Widget buildListShowsSection(
  BuildContext context,
  String title,
  String loadMoreRouteName,
  ListShowsFunction listShows,
) {
  return Column(
    crossAxisAlignment: CrossAxisAlignment.start,
    children: [
      buildHeadline(context, title, loadMoreRouteName),
      const SizedBox(height: 16),
      buildShowList(context, listShows),
    ],
  );
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
