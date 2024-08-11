import 'package:flutter/material.dart';
import 'package:phimtor_app/views/videos/video_player.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

class VideoScreen extends StatefulWidget {
  final phimtor_api.Video video;
  const VideoScreen({super.key, required this.video});

  @override
  State<VideoScreen> createState() => _VideoScreenState();
}

class _VideoScreenState extends State<VideoScreen> {
  phimtor_api.TorrentLink? _selectedTorrentLink;

  @override
  void initState() {
    super.initState();

    selectTorrentLink(widget.video.torrentLinks.first);
  }

  Future<void> selectTorrentLink(phimtor_api.TorrentLink torrentLink) async {
    setState(() {
      _selectedTorrentLink = torrentLink;
    });
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SizedBox(
            width: MediaQuery.of(context).size.width,
            height: MediaQuery.of(context).size.width * 9.0 / 16.0,
            child: VideoPlayer(
              torrentLink: _selectedTorrentLink!,
            ),
          ),
          const SizedBox(height: 16),
          Padding(
            padding: const EdgeInsets.all(16),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const Text("Torrent links"),
                Wrap(
                  spacing: 8,
                  children: widget.video.torrentLinks.map((link) {
                    VoidCallback? onPressed;
                    if (link != _selectedTorrentLink) {
                      onPressed = () => selectTorrentLink(link);
                    }
                    return ElevatedButton(
                      onPressed: onPressed,
                      child: Text(link.name),
                    );
                  }).toList(),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
