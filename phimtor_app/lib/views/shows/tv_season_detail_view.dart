import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:intl/intl.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/routes/route_names.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

class TVSeasonDetailView extends StatelessWidget {
  const TVSeasonDetailView({
    super.key,
    required this.seriesId,
    required this.seasonNumber,
    required this.title,
  });
  final int seriesId;
  final int seasonNumber;
  final String title;

  @override
  Widget build(BuildContext context) {
    AnalyticsService().sendEvent(
      name: "tv_season_detail_view",
      parameters: {
        "series_id": seriesId,
        "season_number": seasonNumber,
        "title": title,
      },
    );

    final infoTextStyte = Theme.of(context).textTheme.bodyMedium;

    return Scaffold(
      appBar: AppBar(
        title: Text(title),
      ),
      body: FutureBuilder(
        future: PhimtorService().defaultApi.getTvSeason(seriesId, seasonNumber),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          }
          if (snapshot.hasError) {
            return Center(
                child: Text(context.loc.error(snapshot.error.toString())));
          }

          final resp = snapshot.data as phimtor_api.GetTvSeasonResponse;
          final season = resp.tvSeason;
          return SingleChildScrollView(
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Row(
                    children: [
                      Expanded(
                        flex: 2,
                        child: Image.network(
                          season.posterLink,
                          width: 150,
                          fit: BoxFit.cover,
                        ),
                      ),
                      const SizedBox(width: 8),
                      Expanded(
                        flex: 3,
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              season.name,
                              style: Theme.of(context).textTheme.headlineMedium,
                            ),
                            Text(
                              season.overview,
                              style: infoTextStyte,
                            ),
                            const SizedBox(height: 16),
                            Text(
                              season.airDate?.year.toString() ?? "",
                              style: infoTextStyte,
                            ),
                            Text(
                              season.airDate == null
                                  ? ""
                                  : DateFormat.yMMMMd().format(season.airDate!),
                              style: infoTextStyte,
                            ),
                          ],
                        ),
                      ),
                    ],
                  ),
                  const SizedBox(height: 16),
                  Text(
                    context.loc.latest_episodes,
                    style: Theme.of(context).textTheme.headlineMedium,
                  ),
                  const SizedBox(height: 8),
                  ListView.separated(
                    shrinkWrap: true,
                    physics: const NeverScrollableScrollPhysics(),
                    itemCount: season.episodes.length,
                    itemBuilder: (context, index) {
                      return TVEpisodeInnerDisplay(
                        episode: season.episodes[index],
                      );
                    },
                    separatorBuilder: (context, index) => Divider(
                      color: Theme.of(context).dividerColor,
                    ),
                  ),
                ],
              ),
            ),
          );
        },
      ),
    );
  }
}

class TVEpisodeInnerDisplay extends StatelessWidget {
  const TVEpisodeInnerDisplay({
    super.key,
    required this.episode,
  });

  final phimtor_api.TVSeasonEpisodesInner episode;

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: () {
        if (episode.videoID == 0) {
          return;
        }
        context.goNamed(
          routeNameVideo,
          pathParameters: {
            "id": episode.videoID.toString(),
            "title": episode.name,
          },
        );
      },
      child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              episode.stillLink != ""
                  ? Image.network(
                      episode.stillLink,
                      width: 200,
                      fit: BoxFit.cover,
                    )
                  : const Icon(Icons.image_not_supported),
              const SizedBox(width: 8),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      episode.name,
                      style: Theme.of(context).textTheme.headlineMedium,
                    ),
                    Text(
                      episode.airDate != null
                          ? DateFormat.yMMMMd().format(episode.airDate!)
                          : "",
                      style: Theme.of(context).textTheme.bodyMedium,
                    ),
                    Text(
                      episode.overview,
                      style: Theme.of(context)
                          .textTheme
                          .bodyMedium!
                          .merge(const TextStyle(
                            fontStyle: FontStyle.italic,
                          )),
                    ),
                    Text(
                      "${context.loc.detail_score}: ${episode.voteAverage.toStringAsFixed(1)}",
                      style: Theme.of(context).textTheme.bodyMedium,
                    ),
                    if (episode.videoID == 0)
                      Text(
                        context.loc.not_available,
                        style: Theme.of(context).textTheme.bodyMedium!.merge(
                              const TextStyle(
                                fontStyle: FontStyle.italic,
                              ),
                            ),
                      ),
                  ],
                ),
              ),
            ],
          )),
    );
  }
}
