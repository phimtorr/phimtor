import 'package:flutter/material.dart';
import 'package:phimtor_app/services/analytics/analytics_service.dart';
import 'package:phimtor_app/views/shows/movie_detail_view.dart';
import 'package:phimtor_app/views/shows/series_detail_view.dart';
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
          // height: 315.0,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
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
          Navigator.of(context).push(MaterialPageRoute(
            builder: (context) => MovieDetailView(
              movieId: show.id,
              title: show.title,
            ),
          ));
          return;
        }
        if (show.type == phimtor_api.ShowType.series) {
          Navigator.of(context).push(MaterialPageRoute(
            builder: (context) => SeriesDetailView(
              seriesId: show.id,
              title: show.title,
            ),
          ));
          return;
        }
      },
    );
  }
}
