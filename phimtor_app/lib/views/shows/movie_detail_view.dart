import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/helpers/time_helpers.dart';
import 'package:phimtor_app/routes/route_names.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

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

          final resp = snapshot.data as phimtor_api.GetMovieResponse;
          final movie = resp.movie;
          return buildMovieScreen(movie);
        },
      ),
    );
  }

  Widget buildMovieScreen(phimtor_api.Movie movie) {
    return LayoutBuilder(
      builder: (context, constraints) {
        // Check if the screen is wide (e.g., desktop/tablet)
        bool isWideScreen = constraints.maxWidth > 600;

        return SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Backdrop and Play Button
              Stack(
                children: [
                  // Backdrop image
                  Image.network(
                    movie.backdropLink,
                    width: double.infinity,
                    height: isWideScreen ? 400.0 : 250.0,
                    fit: BoxFit.cover,
                  ),
                ],
              ),
              // Movie details layout: conditional based on screen width
              Padding(
                padding: const EdgeInsets.all(16.0),
                child: isWideScreen
                    ? Row(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          // Poster image
                          SizedBox(
                            width: 200.0,
                            height: 300.0,
                            child: Image.network(
                              movie.posterLink,
                              fit: BoxFit.cover,
                            ),
                          ),
                          const SizedBox(width: 16.0),
                          // Movie details
                          Expanded(
                            child: buildMovieDetails(context, movie),
                          ),
                        ],
                      )
                    : Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          // Poster image
                          SizedBox(
                            width: 100.0,
                            height: 150.0,
                            child: Image.network(
                              movie.posterLink,
                              fit: BoxFit.cover,
                            ),
                          ),
                          const SizedBox(height: 16.0),
                          // Movie details
                          buildMovieDetails(context, movie),
                        ],
                      ),
              ),
              // Overview
              Padding(
                padding: const EdgeInsets.all(16.0),
                child: Text(
                  movie.overview,
                  style: Theme.of(context).textTheme.bodyMedium!.merge(
                        const TextStyle(
                          fontStyle: FontStyle.italic,
                        ),
                      ),
                ),
              ),
            ],
          ),
        );
      },
    );
  }

  // Method to build the movie details section
  Widget buildMovieDetails(BuildContext context, phimtor_api.Movie movie) {
    final infoTextStyte = Theme.of(context).textTheme.bodyMedium;
    return Column(
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
        const SizedBox(height: 8.0),
        Text(
          "${context.loc.detail_release_year}: ${movie.releaseDate.year}",
          style: infoTextStyte,
        ),
        const SizedBox(height: 8.0),
        Row(
          children: [
            Text(
              "${context.loc.detail_duration(movie.runtime)} (${TimeHelpers.toHumanReadableDuration(movie.runtime)})",
              style: infoTextStyte,
            ),
            const SizedBox(width: 16.0),
            Text(
              "${context.loc.detail_score}: ${movie.voteAverage.toStringAsFixed(1)}",
              style: infoTextStyte,
            ),
          ],
        ),
        const SizedBox(height: 16.0),
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
    );
  }
}
