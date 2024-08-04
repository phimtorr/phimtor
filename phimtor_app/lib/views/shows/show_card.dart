import 'package:flutter/material.dart';
import 'package:phimtor_openapi_client/api.dart';

class ShowCard extends StatelessWidget {
  final ModelShow show;

  const ShowCard({
    super.key,
    required this.show,
  });

  factory ShowCard.fake() {
    return ShowCard(
      show: ModelShow(
        id: 1,
        title: "Bộ Chiến Tranh Bất Lịch Sự",
        originalTitle: "The Ministry of Ungentlemanly Warfare",
        posterLink:
            "https://image.tmdb.org/t/p/w300/avbU7Msx1O7387Lnh4N81Bq4gfC.jpg",
        type: ShowType.movie,
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
    return SizedBox(
      width: 150,
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
            style: const TextStyle(
              fontSize: 16,
              fontWeight: FontWeight.bold,
            ),
            maxLines: 2,
            overflow: TextOverflow.ellipsis,
          ),
          if (show.originalTitle != "") ...[
            const SizedBox(height: 4),
            Text(
              show.originalTitle,
              style: const TextStyle(
                fontSize: 14,
                color: Colors.grey,
              ),
              maxLines: 2,
              overflow: TextOverflow.ellipsis,
            ),
          ],
        ],
      ),
    );
  }
}
