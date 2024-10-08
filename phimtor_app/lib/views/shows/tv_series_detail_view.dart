import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/routes/app_routes.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_app/views/shows/show_components.dart';
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
          return buildTVSeriesDetailScreen(context, series);
        },
      ),
    );
  }

  Widget buildTVSeriesDetailScreen(
      BuildContext context, phimtor_api.TvSeries series) {
    return LayoutBuilder(
      builder: (context, constraints) {
        bool isWideScreen = constraints.maxWidth > 600;

        return SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Backdrop and Series Details
              Stack(
                children: [
                  // Backdrop image
                  Image.network(
                    series.backdropLink,
                    width: double.infinity,
                    height: isWideScreen ? 400.0 : 250.0,
                    fit: BoxFit.cover,
                  ),
                  if (series.tagline != "")
                    Positioned(
                      bottom: 16.0,
                      left: 16.0,
                      child:
                          ShowComponents.buildTagline(context, series.tagline),
                    ),
                ],
              ),
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
                              series.posterLink,
                              fit: BoxFit.cover,
                            ),
                          ),
                          const SizedBox(width: 16.0),
                          // Series details
                          Expanded(
                            child: buildSeriesDetails(context, series),
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
                              series.posterLink,
                              fit: BoxFit.cover,
                            ),
                          ),
                          const SizedBox(height: 16.0),
                          // Series details
                          buildSeriesDetails(context, series),
                        ],
                      ),
              ),
              Padding(
                padding: const EdgeInsets.all(16.0),
                child: Text(
                  series.overview,
                  style: Theme.of(context).textTheme.bodyMedium!.merge(
                        const TextStyle(
                          fontStyle: FontStyle.italic,
                        ),
                      ),
                ),
              ),
              // Season list
              // display list of seasons here
              buildSeasonSection(context, series)
            ],
          ),
        );
      },
    );
  }

  Widget buildSeriesDetails(BuildContext context, phimtor_api.TvSeries series) {
    final infoTextStyte = Theme.of(context).textTheme.bodyMedium;
    return Column(
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
        const SizedBox(height: 16.0),
        ShowComponents.buildGenres(context, series.genres),
        const SizedBox(height: 8.0),
        if (series.firstAirDate != null)
          Row(
            children: [
              Text(
                "${context.loc.detail_release_year}:",
                style: infoTextStyte,
              ),
              const SizedBox(width: 4.0),
              ShowComponents.buildLable(
                context,
                ShowComponents.formatReleaseDate(series.firstAirDate!),
              ),
            ],
          ),
        const SizedBox(height: 8),
        Row(
          children: [
            Text(
              "${context.loc.detail_score}:",
              style: infoTextStyte,
            ),
            const SizedBox(width: 8.0),
            ShowComponents.buildLable(
              context,
              series.voteAverage.toStringAsFixed(1),
            ),
          ],
        ),
      ],
    );
  }

  Widget buildSeasonSection(
    BuildContext context,
    phimtor_api.TvSeries series,
  ) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            context.loc.detail_seasons,
            style: Theme.of(context).textTheme.headlineMedium,
          ),
          const SizedBox(height: 16.0),
          ListView.separated(
            shrinkWrap: true,
            physics: const NeverScrollableScrollPhysics(),
            itemCount: series.seasons.length,
            itemBuilder: (context, index) {
              final season = series.seasons[index];
              return buildSeasonDetail(context, series, season);
            },
            separatorBuilder: (context, index) => const SizedBox(height: 8.0),
          ),
        ],
      ),
    );
  }

  Widget buildSeasonDetail(
    BuildContext context,
    phimtor_api.TvSeries series,
    phimtor_api.TvSeriesSeasonsInner season,
  ) {
    return InkWell(
      onTap: () async {
        await context.pushNamed(
          AppRoutes.tvSeriesSeasonDetails,
          pathParameters: {
            "id": seriesId.toString(),
            "seasonNumber": season.seasonNumber.toString(),
            "title": "${series.name} - ${season.name}",
          },
        );
      },
      child: Ink(
        padding: const EdgeInsets.all(8.0),
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(8.0),
          color: Theme.of(context)
              .colorScheme
              .surfaceContainerLow
              .withOpacity(0.7),
        ),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            SizedBox(
              width: 100.0,
              height: 150.0,
              child: season.posterLink != ""
                  ? Image.network(
                      season.posterLink,
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
                    season.name,
                    style: Theme.of(context).textTheme.headlineMedium,
                  ),
                  const SizedBox(height: 8.0),
                  Row(
                    children: [
                      if (season.airDate != null) ...[
                        ShowComponents.buildLable(
                          context,
                          ShowComponents.formatReleaseDate(season.airDate!),
                        ),
                        const SizedBox(width: 8.0),
                      ],
                      Row(
                        children: [
                          Text(
                            "${context.loc.detail_score}:",
                            style: Theme.of(context).textTheme.bodyMedium,
                          ),
                          const SizedBox(width: 8.0),
                          ShowComponents.buildLable(
                            context,
                            season.voteAverage.toStringAsFixed(1),
                          ),
                        ],
                      ),
                    ],
                  ),
                  const SizedBox(height: 8.0),
                  Text(
                    season.overview,
                    style: Theme.of(context)
                        .textTheme
                        .bodyMedium!
                        .merge(const TextStyle(
                          fontStyle: FontStyle.italic,
                        )),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
