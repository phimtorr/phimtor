import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

class TVSeriesDetailView extends StatelessWidget {
  final int seriesId;
  final String title;
  const TVSeriesDetailView({
    super.key,
    required this.seriesId,
    required this.title,
  });

  @override
  Widget build(BuildContext context) {
    AnalyticsService().sendEvent(
      name: "tv_series_detail_view",
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
        future: PhimtorService().defaultApi.getTvSeries(seriesId),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          }
          if (snapshot.hasError) {
            return Center(
                child: Text(context.loc.error(snapshot.error.toString())));
          }

          final resp = snapshot.data as phimtor_api.GetTvSeriesResponse;
          final series = resp.tvSeries;
          return SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Image.network(
                  series.backdropLink,
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
                        series.name,
                        style: Theme.of(context).textTheme.headlineMedium,
                      ),
                      Text(
                        series.originalName,
                        style: Theme.of(context).textTheme.titleMedium!.merge(
                              const TextStyle(
                                fontStyle: FontStyle.italic,
                              ),
                            ),
                      ),
                      const SizedBox(height: 16),
                      Text(
                        "${context.loc.detail_release_year}: ${series.firstAirDate?.year}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      Text(
                        "${context.loc.detail_score}: ${series.voteAverage.toStringAsFixed(1)}",
                        style: infoTextStyte,
                      ),
                      const SizedBox(height: 8),
                      const SizedBox(height: 8),
                      // Text(
                      //   "${context.loc.detail_total_episodes}: ${series.totalEpisodes}",
                      //   style: infoTextStyte,
                      // ),
                      const SizedBox(height: 16),
                      Text(
                        series.overview,
                        style: infoTextStyte!.merge(const TextStyle(
                          fontStyle: FontStyle.italic,
                        )),
                      ),
                      const SizedBox(height: 16),
                      // display list of seasons here
                      Text(
                        context.loc.detail_seasons,
                        style: Theme.of(context).textTheme.headlineMedium,
                      ),
                      const SizedBox(height: 8),
                      ListView.builder(
                        shrinkWrap: true,
                        physics: const NeverScrollableScrollPhysics(),
                        itemCount: series.seasons.length,
                        itemBuilder: (context, index) {
                          final season = series.seasons[index];
                          return TVSeasonInnerDisplay(season: season);
                        },
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

class TVSeasonInnerDisplay extends StatelessWidget {
  const TVSeasonInnerDisplay({super.key, required this.season});

  final phimtor_api.TvSeriesSeasonsInner season;

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: () {
        // context.goNamed(
        //   routeNameTvSeriesSeason,
        //   pathParameters: {
        //     "seriesId": season.tvSeriesId.toString(),
        //     "seasonNumber": season.seasonNumber.toString(),
        //     "title": season.name,
        //   },
        // );
      },
      child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Row(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              season.posterLink != ""
                  ? Image.network(
                      season.posterLink,
                      width: 100,
                      fit: BoxFit.cover,
                    )
                  : const Icon(Icons.image_not_supported),
              const SizedBox(width: 8),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      season.name,
                      style: Theme.of(context).textTheme.headlineMedium,
                    ),
                    Text(
                      season.airDate.toString(),
                      style: Theme.of(context).textTheme.bodyMedium,
                    ),
                    Text(
                      season.overview,
                      style: Theme.of(context).textTheme.bodyMedium!.merge(const TextStyle(
                            fontStyle: FontStyle.italic,
                          )),
                    ),
                    Text(
                      "${context.loc.detail_score}: ${season.voteAverage.toStringAsFixed(1)}",
                      style: Theme.of(context).textTheme.bodyMedium,
                    ),
                  ],
                ),
              ),
            ],
          )),
    );
  }
}
