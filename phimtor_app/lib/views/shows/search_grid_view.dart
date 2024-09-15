import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_app/views/shows/shows_grid.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

class SearchGridView extends StatelessWidget {
  const SearchGridView({super.key, required this.query})
      : assert(query.length > 0);

  final String query;

  Future<(List<phimtor_api.ModelShow>, phimtor_api.Pagination)> _loadSearch(
    int page,
    int pageSize,
  ) async {
    final resp =
        await PhimtorService().defaultApi.searchShows(query, page: page);
    if (resp == null) {
      throw Exception("Null response");
    }
    return (resp.shows, resp.pagination);
  }

  @override
  Widget build(BuildContext context) {
    AnalyticsService().sendEvent(
      name: "search_grid_view",
      parameters: {
        "query": query,
      },
    );
    return Scaffold(
      appBar: AppBar(
        title: Text(context.loc.search_result_title(query)),
      ),
      body: Padding(
        padding: const EdgeInsets.all(8.0),
        child: ShowsGrid(
          loadMore: _loadSearch,
        ),
      ),
    );
  }
}
