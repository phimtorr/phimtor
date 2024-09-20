import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:phimtor_app/helpers/time_helpers.dart';
import 'package:phimtor_app/routes/app_routes.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

class ShowCard extends StatelessWidget {
  final phimtor_api.ModelShow show;

  const ShowCard({
    super.key,
    required this.show,
  });

  @override
  Widget build(BuildContext context) {
    return InkWell(
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: SizedBox(
          width: 150.0,
          height: 315.0,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Stack(
                children: [
                  ClipRRect(
                    borderRadius: BorderRadius.circular(8),
                    child: Image.network(
                      show.posterLink,
                      width: 150,
                      height: 200,
                      fit: BoxFit.cover,
                    ),
                  ),
                  Positioned(
                    top: 2,
                    left: 2,
                    child: Row(
                      children: [
                        ShowLabel(title: show.airDate.year.toString()),
                        const SizedBox(width: 2),
                        ShowLabel(title: show.voteAverage.toStringAsFixed(1)),
                      ],
                    ),
                  ),
                  if (show.type == phimtor_api.ModelShowTypeEnum.tvSeries) ...[
                    // Positioned(
                    //   bottom: 2,
                    //   right: 2,
                    //   child: ShowLabel(
                    //     title: "${show.currentEpisode}/${show.totalEpisodes}",
                    //   ),
                    // ),
                  ],
                  if (show.type == phimtor_api.ModelShowTypeEnum.movie) ...[
                    // Positioned(
                    //   bottom: 2,
                    //   left: 2,
                    //   child: ShowLabel(title: show.quality),
                    // ),
                    Positioned(
                      bottom: 2,
                      right: 2,
                      child: ShowLabel(
                        title:
                            TimeHelpers.toHumanReadableDuration(show.runtime),
                      ),
                    ),
                  ],
                ],
              ),
              const SizedBox(height: 8),
              Text(
                show.title,
                style: Theme.of(context).textTheme.titleMedium,
                maxLines: 2,
                overflow: TextOverflow.ellipsis,
              ),
              if (show.originalTitle != "") ...[
                const SizedBox(height: 4),
                Text(
                  show.originalTitle,
                  style: Theme.of(context)
                      .textTheme
                      .bodyMedium!
                      .merge(const TextStyle(fontStyle: FontStyle.italic)),
                  maxLines: 2,
                  overflow: TextOverflow.ellipsis,
                ),
              ],
            ],
          ),
        ),
      ),
      onTap: () async {
        AnalyticsService().sendEvent(
          name: "show_card_tap",
          parameters: {
            "show_id": show.id,
            "title": show.title,
          },
        );
        if (show.type == phimtor_api.ModelShowTypeEnum.movie) {
          await context.pushNamed(
            AppRoutes.movieDetails,
            pathParameters: {
              "id": show.showId.toString(),
              "title": show.title,
            },
          );
          return;
        }
        if (show.type == phimtor_api.ModelShowTypeEnum.tvSeries) {
          await context.pushNamed(
            AppRoutes.tvSeriesDetails,
            pathParameters: {
              "id": show.showId.toString(),
              "title": show.title,
            },
          );
          return;
        }
        if (show.type == phimtor_api.ModelShowTypeEnum.episode) {
          await context.pushNamed(
            AppRoutes.tvSeriesSeasonDetails,
            pathParameters: {
              "id": show.showId.toString(),
              "seasonNumber": show.seasonNumber.toString(),
              "title": show.title,
            },
          );
          return;
        }
      },
    );
  }
}

class ShowLabel extends StatelessWidget {
  const ShowLabel({
    super.key,
    required this.title,
  });

  final String title;

  @override
  Widget build(BuildContext context) {
    final backgroundColor =
        Theme.of(context).colorScheme.surfaceContainerLow.withOpacity(0.6);
    final textStyle = Theme.of(context).textTheme.labelSmall!.merge(
          TextStyle(
            color: Theme.of(context).colorScheme.onSurface,
          ),
        );
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 4, vertical: 4),
      decoration: BoxDecoration(
        color: backgroundColor,
        borderRadius: BorderRadius.circular(4),
      ),
      child: Text(
        title,
        style: textStyle,
      ),
    );
  }
}
