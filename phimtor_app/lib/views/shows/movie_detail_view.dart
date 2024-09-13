import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/helpers/time_helpers.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_app/views/videos/video_view.dart';
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
                  movie.posterLink,
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
                        "${context.loc.detail_release_year}: ${movie.releaseYear}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        "${context.loc.detail_score}: ${movie.score}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        "${context.loc.detail_duration(movie.durationInMinutes)} (${TimeHelpers.toHumanReadableDuration(movie.durationInMinutes)})",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        "${context.loc.detail_quality}: ${movie.quantity}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 16),
                      Text(
                        movie.description,
                        style: infoTextStyte!.merge(const TextStyle(
                          fontStyle: FontStyle.italic,
                        )),
                      ),
                      const SizedBox(height: 16),
                      ElevatedButton.icon(
                        onPressed: () {
                          Navigator.of(context).push(MaterialPageRoute(
                            builder: (context) => VideoView(
                              videoId: movie.videoId,
                              title: movie.title,
                            ),
                          ));
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
