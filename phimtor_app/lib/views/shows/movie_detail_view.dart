import 'package:flutter/material.dart';
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
    const infoTextStyte = TextStyle(
      fontSize: 16,
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
            return Center(child: Text("Error: ${snapshot.error}"));
          }

          final resp = snapshot.data as GetMovieResponse;
          final movie = resp.movie;
          return Column(
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
                      style: const TextStyle(
                        fontSize: 24,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    Text(
                      movie.originalTitle,
                      style: const TextStyle(
                        fontSize: 16,
                        fontStyle: FontStyle.italic,
                      ),
                    ),
                    const SizedBox(height: 8),
                    Text(
                      "Release year: ${movie.releaseYear}",
                      style: infoTextStyte,
                    ),
                    const SizedBox(height: 8),
                    Text(
                      "Score: ${movie.score}",
                      style: infoTextStyte,
                    ),
                    const SizedBox(height: 8),
                    Text(
                      "Duration: ${movie.durationInMinutes} minutes",
                      style: infoTextStyte,
                    ),
                    const SizedBox(height: 8),
                    Text(
                      "Quantity: ${movie.quantity}",
                      style: infoTextStyte,
                    ),
                    const SizedBox(height: 8),
                    Text(
                      movie.description,
                      style: infoTextStyte,
                    ),
                    const SizedBox(height: 8),
                    ElevatedButton.icon(
                      onPressed: () {},
                      label: const Text("Watch now"),
                      icon: const Icon(Icons.play_arrow),
                    ),
                  ],
                ),
              ),
            ],
          );
        },
      ),
    );
  }
}
