import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/helpers/time_helpers.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_app/views/videos/video_view.dart';
import 'package:phimtor_openapi_client/api.dart';

class SeriesDetailView extends StatelessWidget {
  final int seriesId;
  final String title;
  const SeriesDetailView({
    super.key,
    required this.seriesId,
    required this.title,
  });

  @override
  Widget build(BuildContext context) {
    AnalyticsService().sendEvent(
      name: "series_detail_view",
      parameters: {
        "series_id": seriesId,
        "title": title,
      },
    );

    final infoTextStyte = Theme.of(context).textTheme.bodyMedium;
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
            return Center(
                child: Text(context.loc.error(snapshot.error.toString())));
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
                        style: Theme.of(context).textTheme.headlineMedium,
                      ),
                      Text(
                        series.originalTitle,
                        style: Theme.of(context).textTheme.titleMedium!.merge(
                              const TextStyle(
                                fontStyle: FontStyle.italic,
                              ),
                            ),
                      ),
                      const SizedBox(height: 16),
                      Text(
                        "${context.loc.detail_release_year}: ${series.releaseYear}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        "${context.loc.detail_score}: ${series.score}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        "${context.loc.detail_duration(series.durationInMinutes)} (${TimeHelpers.toHumanReadableDuration(series.durationInMinutes)})",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        "${context.loc.detail_total_episodes}: ${series.totalEpisodes}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 16),
                      Text(
                        series.description,
                        style: infoTextStyte!.merge(const TextStyle(
                          fontStyle: FontStyle.italic,
                        )),
                      ),
                      const SizedBox(height: 16),
                      Wrap(
                        spacing: 8,
                        runSpacing: 8.0,
                        children: List.generate(series.episodes.length, (i) {
                          final episode = series.episodes[i];
                          return ElevatedButton(
                            onPressed: () {
                              Navigator.of(context).push(MaterialPageRoute(
                                builder: (context) => VideoView(
                                  videoId: episode.videoId,
                                  title: "${series.title} - ${episode.name}",
                                ),
                              ));
                            },
                            child: Text(episode.name),
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
