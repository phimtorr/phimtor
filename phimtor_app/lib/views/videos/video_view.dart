import 'package:flutter/material.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_app/views/videos/video_screen.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

class VideoView extends StatelessWidget {
  final int videoId;
  final String title;
  const VideoView({
    super.key,
    required this.videoId,
    required this.title,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(title),
      ),
      body: FutureBuilder<phimtor_api.Video>(
        future: () async {
          final resp = await PhimtorService().defaultApi.getVideo(videoId);
          if (resp == null) {
            throw Exception('Failed to get video');
          }
          return resp.video;
        }(),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          }
          if (snapshot.hasError) {
            return Center(
              child: Text('Error: ${snapshot.error}'),
            );
          }
          if (!snapshot.hasData) {
            return const Center(
              child: Text('No data'),
            );
          }
          final video = snapshot.data!;
          return VideoScreen(video: video);
        },
      ),
    );
  }
}
