import 'package:flutter/material.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/views/shows/shows_grid.dart';

class ShowsGridView extends StatelessWidget {
  const ShowsGridView({
    super.key,
    required this.title,
    required this.loadMore,
  });

  final String title;
  final LoadMoreCallback loadMore;

  @override
  Widget build(BuildContext context) {
    AnalyticsService().sendEvent(
      name: "shows_grid_view",
      parameters: {
        "title": title,
      },
    );
    return Scaffold(
      appBar: AppBar(
        title: Text(title),
      ),
      body: Padding(
        padding: const EdgeInsets.all(8.0),
        child: ShowsGrid(loadMore: loadMore),
      ),
    );
  }
}