import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:phimtor_app/helpers/time_helpers.dart';
import 'package:phimtor_app/routes/route_names.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

class ShowCard extends StatelessWidget {
  final phimtor_api.ModelShow show;

  const ShowCard({
    super.key,
    required this.show,
  });

  factory ShowCard.fake() {
    return ShowCard(
      show: phimtor_api.ModelShow(
        id: 1,
        title: "Bộ Chiến Tranh Bất Lịch Sự",
        originalTitle: "The Ministry of Ungentlemanly Warfare",
        posterLink:
            "https://image.tmdb.org/t/p/w300/avbU7Msx1O7387Lnh4N81Bq4gfC.jpg",
        type: phimtor_api.ShowType.movie,
        releaseYear: 2022,
        score: 8.5,
        durationInMinutes: 120,
        quantity: "4K",
        totalEpisodes: 1,
        currentEpisode: 1,
      ),
    );
  }

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
                        ShowLabel(title: show.releaseYear.toString()),
                        const SizedBox(width: 2),
                        ShowLabel(title: show.score.toString()),
                      ],
                    ),
                  ),
                  if (show.type == phimtor_api.ShowType.series) ...[
                    Positioned(
                      bottom: 2,
                      right: 2,
                      child: ShowLabel(
                        title: "${show.currentEpisode}/${show.totalEpisodes}",
                      ),
                    ),
                  ],
                  if (show.type == phimtor_api.ShowType.movie) ...[
                    Positioned(
                      bottom: 2,
                      left: 2,
                      child: ShowLabel(title: show.quantity),
                    ),
                    Positioned(
                      bottom: 2,
                      right: 2,
                      child: ShowLabel(
                          title: TimeHelpers.toHumanReadableDuration(
                              show.durationInMinutes)),
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
      onTap: () {
        AnalyticsService().sendEvent(
          name: "show_card_tap",
          parameters: {
            "show_id": show.id,
            "title": show.title,
          },
        );
        if (show.type == phimtor_api.ShowType.movie) {
          context.goNamed(
            routeNameMovieDetails,
            pathParameters: {
              "id": show.id.toString(),
              "title": show.title,
            },
          );
          return;
        }
        if (show.type == phimtor_api.ShowType.series) {
          context.goNamed(
            routeNameSeriesDetails,
            pathParameters: {
              "id": show.id.toString(),
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
