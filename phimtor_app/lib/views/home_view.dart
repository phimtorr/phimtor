import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart'
    as phimtor_api;
import 'package:phimtor_app/views/shows/shows_grid_view.dart';
import 'package:phimtor_app/views/shows/shows_list.dart';
import 'package:phimtor_openapi_client/api.dart';

class HomeView extends StatelessWidget {
  const HomeView({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Phim Tor"),
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

class SearchSection extends StatefulWidget {
  const SearchSection({super.key});

  @override
  State<SearchSection> createState() => _SearchSectionState();
}

class _SearchSectionState extends State<SearchSection> {
  final _searchController = TextEditingController();

  Future<(List<ModelShow>, Pagination)> _loadSearch(int page, int pageSize) async {
    final resp = await phimtor_api.PhimtorService()
        .defaultApi
        .searchShows(_searchController.text, page: page);
    if (resp == null) {
      throw Exception("Null response");
    }
    return (resp.shows, resp.pagination);
  }

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
          "Search",
          style: Theme.of(context).textTheme.headlineLarge,
        ),
        const SizedBox(height: 16),
        // a search text box with a button to search
        CupertinoSearchTextField(
          controller: _searchController,
          onSubmitted: (query) {
            if (query.isEmpty) {
              return;
            }
            Navigator.of(context).push(MaterialPageRoute(
              builder: (context) => ShowsGridView(
                title: "Search results for '$query'",
                loadMore: _loadSearch,
              ),
            ));
          },
        ),
      ],
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
              "Movies",
              style: Theme.of(context).textTheme.headlineLarge,
            ),
            const SizedBox(width: 8),
            ElevatedButton.icon(
              onPressed: () {
                Navigator.of(context).push(MaterialPageRoute(
                  builder: (context) => ShowsGridView(
                    title: "Movies",
                    loadMore: _loadMovies,
                  ),
                ));
              },
              label: const Text("Load more"),
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
                return Center(child: Text("Error: ${snapshot.error}"));
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
              "TV Series",
              style: Theme.of(context).textTheme.headlineLarge,
            ),
            const SizedBox(width: 8),
            ElevatedButton.icon(
              onPressed: () {
                Navigator.of(context).push(MaterialPageRoute(
                  builder: (context) => ShowsGridView(
                    title: "TV Series",
                    loadMore: _loadSeries,
                  ),
                ));
              },
              label: const Text("Load more"),
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
                return Center(child: Text("Error: ${snapshot.error}"));
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
