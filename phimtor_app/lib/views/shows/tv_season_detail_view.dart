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
          return buildTVSeasonDetailScreen(context, season);
        },
      ),
    );
  }

  Widget buildTVSeasonDetailScreen(
    BuildContext context,
    phimtor_api.TVSeason season,
  ) {
    return LayoutBuilder(
      builder: (context, constraints) {
        bool isWideScreen = constraints.maxWidth > 600;

        return SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Season details
              Padding(
                padding: const EdgeInsets.all(16.0),
                child: isWideScreen
                    ? Row(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          // Season poster
                          SizedBox(
                            width: 200.0,
                            height: 300.0,
                            child: Image.network(
                              season.posterLink,
                              fit: BoxFit.cover,
                            ),
                          ),
                          const SizedBox(width: 16.0),
                          // Season information
                          Expanded(
                            child: buildSeasonDetails(context, season),
                          ),
                        ],
                      )
                    : Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          // Season poster
                          SizedBox(
                            width: 150.0,
                            height: 200.0,
                            child: Image.network(
                              season.posterLink,
                              fit: BoxFit.cover,
                            ),
                          ),
                          const SizedBox(height: 16.0),
                          // Season information
                          buildSeasonDetails(context, season),
                        ],
                      ),
              ),
              Padding(
                padding: const EdgeInsets.all(16.0),
                child: Text(
                  season.overview,
                  style: Theme.of(context).textTheme.bodyMedium!.merge(
                        const TextStyle(
                          fontStyle: FontStyle.italic,
                        ),
                      ),
                ),
              ),
              // Episodes section
              buildEpisodeSection(context, season),
            ],
          ),
        );
      },
    );
  }

  Widget buildSeasonDetails(
    BuildContext context,
    phimtor_api.TVSeason season,
  ) {
    final infoTextStyte = Theme.of(context).textTheme.bodyMedium;
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          season.name,
          style: Theme.of(context).textTheme.headlineMedium,
        ),
        const SizedBox(height: 8.0),
        Text(
          season.airDate?.year.toString() ?? "",
          style: infoTextStyte,
        ),
        const SizedBox(height: 8.0),
        Text(
          season.airDate == null
              ? ""
              : DateFormat.yMMMMd().format(season.airDate!),
          style: infoTextStyte,
        ),
      ],
    );
  }

  Widget buildEpisodeSection(
    BuildContext context,
    phimtor_api.TVSeason season,
  ) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Padding(
          padding: const EdgeInsets.all(16.0),
          child: Text(
            context.loc.detail_episodes,
            style: Theme.of(context).textTheme.headlineMedium,
          ),
        ),
        ListView.separated(
          shrinkWrap: true,
          physics: const NeverScrollableScrollPhysics(),
          itemCount: season.episodes.length,
          itemBuilder: (context, index) {
            return buildEpisodeDetail(context, season, season.episodes[index]);
          },
          separatorBuilder: (context, index) => const SizedBox(height: 8.0),
        ),
      ],
    );
  }

  Widget buildEpisodeDetail(
    BuildContext context,
    phimtor_api.TVSeason season,
    phimtor_api.TVSeasonEpisodesInner episode,
  ) {
    return InkWell(
      onTap: () {
        if (episode.videoID == 0) {
          return;
        }
        context.goNamed(
          routeNameVideo,
          pathParameters: {
            "id": episode.videoID.toString(),
            "title": "$title - ${episode.name}",
          },
        );
      },
      child: Padding(
          padding: const EdgeInsets.symmetric(vertical: 8.0, horizontal: 16.0),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              SizedBox(
                width: 200.0,
                child: season.posterLink != ""
                    ? Image.network(
                        episode.stillLink,
                        fit: BoxFit.cover,
                      )
                    : const Center(
                        child: Icon(Icons.image_not_supported),
                      ),
              ),
              const SizedBox(width: 8),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      "${episode.episodeNumber}. ${episode.name}",
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
