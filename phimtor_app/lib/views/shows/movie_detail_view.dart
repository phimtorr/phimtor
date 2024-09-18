import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/helpers/time_helpers.dart';
import 'package:phimtor_app/routes/route_names.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_openapi_client/api.dart';

class MovieDetailView extends StatelessWidget {
  final int movieId;
  final String title;
  const MovieDetailView({
    super.key,
    required this.movieId,
    required this.title,
  });

  @override
  Widget build(BuildContext context) {
    AnalyticsService().sendEvent(
      name: "movie_detail_view",
      parameters: {
        "movie_id": movieId,
        "title": title,
      },
    );
    final infoTextStyte = Theme.of(context).textTheme.bodyMedium;
    return Scaffold(
      appBar: AppBar(
        title: Text(title),
      ),
      body: FutureBuilder(
        future: PhimtorService().defaultApi.getMovie(movieId),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          }
          if (snapshot.hasError) {
            return Center(
                child: Text(context.loc.error(snapshot.error.toString())));
          }

          final resp = snapshot.data as GetMovieResponse;
          final movie = resp.movie;
          return SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Image.network(
                  movie.backdropLink,
                  width: double.infinity,
                  height: 300,
                  fit: BoxFit.cover,
                ),
                Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        movie.title,
                        style: Theme.of(context).textTheme.headlineMedium,
                      ),
                      Text(
                        movie.originalTitle,
                        style: Theme.of(context).textTheme.titleMedium!.merge(
                              const TextStyle(
                                fontStyle: FontStyle.italic,
                              ),
                            ),
                      ),
                      const SizedBox(height: 16),
                      Text(
                        "${context.loc.detail_release_year}: ${movie.releaseDate.year}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        "${context.loc.detail_score}: ${movie.voteAverage.toStringAsFixed(1)}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        "${context.loc.detail_duration(movie.runtime)} (${TimeHelpers.toHumanReadableDuration(movie.runtime)})",
                        style: infoTextStyte,
                      ),
                      // TODO: Update quality
                      // const SizedBox(height: 8),
                      // Text(
                      //   "${context.loc.detail_quality}: ${movie.}",
                      //   style: infoTextStyte,
                      // ),
                      const SizedBox(height: 16),
                      Text(
                        movie.overview,
                        style: infoTextStyte!.merge(const TextStyle(
                          fontStyle: FontStyle.italic,
                        )),
                      ),
                      const SizedBox(height: 16),
                      ElevatedButton.icon(
                        onPressed: () {
                          context.goNamed(routeNameVideo, pathParameters: {
                            "id": movie.videoID.toString(),
                            "title": movie.title,
                          });
                        },
                        label: Text(context.loc.watch_now),
                        icon: const Icon(Icons.play_arrow),
                      ),
                    ],
                  ),
                ),
              ],
            ),
          );
        },
      ),
    );
  }
}
