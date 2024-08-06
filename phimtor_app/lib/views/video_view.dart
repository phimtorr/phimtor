import 'package:flutter/material.dart';

class VideoView extends StatefulWidget {
  final int videoId;
  final String title;
  const VideoView({
    super.key,
    required this.videoId,
    required this.title,
  });

  @override
  State<VideoView> createState() => _VideoViewState();
}

class _VideoViewState extends State<VideoView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Text("Video ID: ${widget.videoId}"),
      ),
    );
  }
}
