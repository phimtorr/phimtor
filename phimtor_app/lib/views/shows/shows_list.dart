import 'package:flutter/material.dart';
import 'package:phimtor_app/views/shows/show_card.dart';
import 'package:phimtor_openapi_client/api.dart';

class ShowsList extends StatelessWidget {
  final Iterable<ModelShow> shows;
  const ShowsList({
    super.key,
    required this.shows,
  });

  static const minHeight = 350.0; // for not overflow

  @override
  Widget build(BuildContext context) {
    return ListView.separated(
      scrollDirection: Axis.horizontal,
      itemCount: shows.length,
      itemBuilder: (context, i) {
        final show = shows.elementAt(i);
        return ShowCard(show: show);
      },
      separatorBuilder: (context, index) {
        return const SizedBox(width: 16);
      },
    );
  }
}
