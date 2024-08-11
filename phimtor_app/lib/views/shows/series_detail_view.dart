import 'package:flutter/material.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_app/views/videos/video_view.dart';
import 'package:phimtor_openapi_client/api.dart';

class SeriesDetailView extends StatelessWidget {
  final int seriesId;
  final String title;
  const SeriesDetailView({super.key, required this.seriesId, required this.title});

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
        future: PhimtorService().defaultApi.getSeries(seriesId),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          }
          if (snapshot.hasError) {
            return Center(child: Text("Error: ${snapshot.error}"));
          }

          final resp = snapshot.data as GetSeriesResponse;
          final series = resp.series;
          return SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Image.network(
                  series.posterLink,
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
                        series.title,
                        style: const TextStyle(
                          fontSize: 24,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                      Text(
                        series.originalTitle,
                        style: const TextStyle(
                          fontSize: 16,
                        ),
                      ),
                      const SizedBox(height: 16),
                       Text(
                        "Release year: ${series.releaseYear}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        "Score: ${series.score}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        "Duration: ${series.durationInMinutes} minutes",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text( 
                        "Total episodes: ${series.totalEpisodes}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        series.description,
                        style: const TextStyle(
                          fontSize: 16,
                        ),
                      ),
                      const SizedBox(height: 16),
                      Wrap(
                        spacing: 8,
                        runSpacing: 8.0,
                        children: List.generate(series.totalEpisodes, (index) {
                          final episode = index + 1;
                          return ElevatedButton(
                            onPressed: () {
                              Navigator.of(context).push(MaterialPageRoute(
                                builder: (context) => VideoView(
                                  videoId: series.episodes[episode - 1].videoId,
                                  title: "${series.title} - Episode $episode",
                                ),
                              ));
                            },
                            child: Text("Episode $episode"),
                          );
                        }),
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