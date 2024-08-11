import 'package:flutter/material.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart'
    as phimtor_api;
import 'package:phimtor_app/views/shows/shows_grid_view.dart';
import 'package:phimtor_app/views/shows/shows_list.dart';
import 'package:phimtor_openapi_client/api.dart';

class HomeView extends StatelessWidget {
  const HomeView({super.key});

  Future<List<ModelShow>> _loadMovies(int page, int pageSize) async {
    final resp = await phimtor_api.PhimtorService()
        .defaultApi
        .listShows(page: page, pageSize: pageSize, type: ShowType.movie);
    if (resp == null) {
      throw Exception("Null response");
    }
    return resp.shows;
  }

  Future<List<ModelShow>> _loadSeries(int page, int pageSize) async {
    final resp = await phimtor_api.PhimtorService()
        .defaultApi
        .listShows(page: page, pageSize: pageSize, type: ShowType.series);
    if (resp == null) {
      throw Exception("Null response");
    }
    return resp.shows;
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Phim Tor"),
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
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
              const SizedBox(height: 32),
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
          ),
        ),
      ),
    );
  }
}
