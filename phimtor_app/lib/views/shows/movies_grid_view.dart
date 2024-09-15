import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_app/views/shows/shows_grid.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

class MoviesGridView extends StatelessWidget {
  const MoviesGridView({super.key});

  Future<(List<phimtor_api.ModelShow>, phimtor_api.Pagination)> _loadMovies(
    int page,
    int pageSize,
  ) async {
    final resp = await PhimtorService().defaultApi.listShows(
          page: page,
          pageSize: pageSize,
          type: phimtor_api.ShowType.movie,
        );
    if (resp == null) {
      throw Exception("Null response");
    }
    return (resp.shows, resp.pagination);
  }

  @override
  Widget build(BuildContext context) {
    AnalyticsService().sendEvent(
      name: "movies_grid_view",
    );
    return Scaffold(
      appBar: AppBar(
        title: Text(context.loc.movies),
      ),
      body: Padding(
        padding: const EdgeInsets.all(8.0),
        child: ShowsGrid(
          loadMore: _loadMovies,
        ),
      ),
    );
  }
}