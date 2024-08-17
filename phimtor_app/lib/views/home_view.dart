import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart'
    as phimtor_api;
import 'package:phimtor_app/views/search_section.dart';
import 'package:phimtor_app/views/shows/shows_grid_view.dart';
import 'package:phimtor_app/views/shows/shows_list.dart';
import 'package:phimtor_openapi_client/api.dart';

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
              MoviesSection(),
              SizedBox(height: 32),
              TVSeriesSection(),
            ],
          ),
        ),
      ),
    );
  }
}


class MoviesSection extends StatelessWidget {
  const MoviesSection({super.key});

  Future<(List<ModelShow>, Pagination)> _loadMovies(int page, int pageSize) async {
    final resp = await phimtor_api.PhimtorService()
        .defaultApi
        .listShows(page: page, pageSize: pageSize, type: ShowType.movie);
    if (resp == null) {
      throw Exception("Null response");
    }
    return (resp.shows, resp.pagination);
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Text(
              context.loc.movies,
              style: Theme.of(context).textTheme.headlineLarge,
            ),
            const SizedBox(width: 8),
            ElevatedButton.icon(
              onPressed: () {
                Navigator.of(context).push(MaterialPageRoute(
                  builder: (context) => ShowsGridView(
                    title: context.loc.movies,
                    loadMore: _loadMovies,
                  ),
                ));
              },
              label:  Text(context.loc.load_more),
              icon: const Icon(Icons.arrow_forward),
              iconAlignment: IconAlignment.end,
            ),
          ],
        ),
        const SizedBox(height: 16),
        SizedBox(
          height: ShowsList.minHeight,
          child: FutureBuilder(
            future: phimtor_api.PhimtorService()
                .defaultApi
                .listShows(page: 1, pageSize: 10, type: ShowType.movie),
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return const Center(child: CircularProgressIndicator());
              }
              if (snapshot.hasError) {
                return Center(child: Text(context.loc.error(snapshot.error.toString())));
              }

              final response = snapshot.data as ListShowsResponse;
              return ShowsList(shows: response.shows);
            },
          ),
        ),
      ],
    );
  }
}

class TVSeriesSection extends StatelessWidget {
  const TVSeriesSection({super.key});

  Future<(List<ModelShow>, Pagination)> _loadSeries(int page, int pageSize) async {
    final resp = await phimtor_api.PhimtorService()
        .defaultApi
        .listShows(page: page, pageSize: pageSize, type: ShowType.series);
    if (resp == null) {
      throw Exception("Null response");
    }
    return (resp.shows, resp.pagination);
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Row(
          children: [
            Text(
              context.loc.tv_series,
              style: Theme.of(context).textTheme.headlineLarge,
            ),
            const SizedBox(width: 8),
            ElevatedButton.icon(
              onPressed: () {
                Navigator.of(context).push(MaterialPageRoute(
                  builder: (context) => ShowsGridView(
                    title: context.loc.tv_series,
                    loadMore: _loadSeries,
                  ),
                ));
              },
              label:  Text(context.loc.load_more),
              icon: const Icon(Icons.arrow_forward),
              iconAlignment: IconAlignment.end,
            ),
          ],
        ),
        const SizedBox(height: 16),
        SizedBox(
          height: ShowsList.minHeight,
          child: FutureBuilder(
            future: phimtor_api.PhimtorService()
                .defaultApi
                .listShows(page: 1, pageSize: 10, type: ShowType.series),
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return const Center(child: CircularProgressIndicator());
              }
              if (snapshot.hasError) {
                return Center(child: Text(context.loc.error(snapshot.error.toString())));
              }

              final response = snapshot.data as ListShowsResponse;
              return ShowsList(shows: response.shows);
            },
          ),
        ),
      ],
    );
  }
}
