import 'package:flutter/material.dart';

class ShowCard extends StatelessWidget {
  final String title;
  final String? originalTitle;
  final String imageUrl;

  const ShowCard({
    super.key,
    required this.title,
    this.originalTitle,
    required this.imageUrl,
  });

  factory ShowCard.fake() {
    return const ShowCard(
      title: "Bộ Chiến Tranh Bất Lịch Sự",
      originalTitle: "The Ministry of Ungentlemanly Warfare",
      imageUrl: "https://image.tmdb.org/t/p/w300/avbU7Msx1O7387Lnh4N81Bq4gfC.jpg",
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
              imageUrl,
              width: 150,
              height: 200,
              fit: BoxFit.cover,
            ),
          ),
          const SizedBox(height: 8),
          Text(
            title,
            style: const TextStyle(
              fontSize: 16,
              fontWeight: FontWeight.bold,
            ),
            maxLines: 2,
            overflow: TextOverflow.ellipsis,
          ),
          if (originalTitle != null) ...[
            const SizedBox(height: 4),
            Text(
              originalTitle!,
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
