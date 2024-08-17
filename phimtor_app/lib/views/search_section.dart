import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart'
    as phimtor_api;
import 'package:phimtor_app/views/shows/shows_grid_view.dart';
import 'package:phimtor_openapi_client/api.dart';

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
          context.loc.search,
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
                title: context.loc.search_result_title(query),
                loadMore: _loadSearch,
              ),
            ));
          },
        ),
      ],
    );
  }
}