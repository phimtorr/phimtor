import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/views/shows/show_card.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

typedef LoadMoreCallback
    = Future<(List<phimtor_api.ModelShow>, phimtor_api.Pagination)> Function(
        int page, int pageSize);

class ShowsGridView extends StatefulWidget {
  const ShowsGridView({
    super.key,
    required this.title,
    required this.loadMore,
  });

  final String title;
  final LoadMoreCallback loadMore;

  @override
  State<ShowsGridView> createState() => _ShowsGridViewState();
}

class _ShowsGridViewState extends State<ShowsGridView> {
  final scrollController = ScrollController();
  List<phimtor_api.ModelShow> shows = [];
  bool isLoading = false;
  bool hasMore = true;
  int currentPage = 1;
  int? totalItems;

  static int pageSize = 20;

  @override
  void initState() {
    super.initState();
    scrollController.addListener(_scrollListener);
    _loadMore();
  }

  @override
  void dispose() {
    scrollController.dispose();
    super.dispose();
  }

  void _scrollListener() {
    if (scrollController.position.pixels ==
        scrollController.position.maxScrollExtent) {
      _loadMore();
    }
  }

  void _loadMore() async {
    if (isLoading) {
      return;
    }
    if (!hasMore) {
      return;
    }
    try {
      setState(() {
        isLoading = true;
      });

      final (newShows, pagination) =
          await widget.loadMore(currentPage, pageSize);
      if (newShows.length < pageSize) {
        setState(() {
          hasMore = false;
        });
      }

      setState(() {
        totalItems = pagination.totalResults;
        shows.addAll(newShows);
        currentPage++;
      });
    } catch (e) {
      debugPrint("Error: $e");
    } finally {
      setState(() {
        isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    AnalyticsService().sendEvent(
      name: "shows_grid_view",
      parameters: {
        "title": widget.title,
      },
    );
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            if (totalItems == 0)
              Center(child: Text(context.loc.search_no_result)),
            if (totalItems != null && totalItems! > 0)
              Padding(
                padding: const EdgeInsets.all(8.0),
                child: Text(context.loc.search_count(shows.length, totalItems!)),
              ),
            Expanded(
              child: SingleChildScrollView(
                controller: scrollController,
                child: Wrap(
                  alignment: WrapAlignment.start,
                  crossAxisAlignment: WrapCrossAlignment.center,
                  spacing: 16,
                  runSpacing: 24,
                  children: shows.map((show) => ShowCard(show: show)).toList(),
                ),
              ),
            ),
            if (isLoading)
              const Center(
                child: CircularProgressIndicator(),
              ),
          ],
        ),
      ),
    );
  }
}
