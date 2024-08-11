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
  phimtor_api.Subtitle? _selectedSubtitle;

  @override
  void initState() {
    super.initState();

    selectTorrentLink(widget.video.torrentLinks.first);
    if (widget.video.subtitles.getByLanguage("vi").isNotEmpty) {
      selectSubtitle(widget.video.subtitles.first);
    }
  }

  void selectTorrentLink(phimtor_api.TorrentLink torrentLink) {
    setState(() {
      _selectedTorrentLink = torrentLink;
    });
  }

  void selectSubtitle(phimtor_api.Subtitle subtitle) {
    setState(() {
      _selectedSubtitle = subtitle;
    });
  }

  @override
  Widget build(BuildContext context) {
    var titleStyle = Theme.of(context).textTheme.headlineMedium!;
    var subtitleStyle = Theme.of(context).textTheme.headlineSmall!;

    var vietnameseSubtitles = widget.video.subtitles.getByLanguage("vi");
    var englishSubtitles = widget.video.subtitles.getByLanguage("en");
    return SingleChildScrollView(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          SizedBox(
            width: MediaQuery.of(context).size.width,
            height: MediaQuery.of(context).size.width * 9.0 / 16.0,
            child: VideoPlayer(
              torrentLink: _selectedTorrentLink!,
              subtitle: _selectedSubtitle,
            ),
          ),
          const SizedBox(height: 16),
          Padding(
            padding: const EdgeInsets.all(16),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text("Torrent links", style: titleStyle),
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
                const SizedBox(height: 16),
                if (widget.video.subtitles.isNotEmpty)
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text("Subtitles", style: titleStyle),
                      const SizedBox(height: 8),
                      if (vietnameseSubtitles.isNotEmpty)
                        Text("Vietnamese", style: subtitleStyle),
                      Wrap(
                        spacing: 8,
                        children: widget.video.subtitles
                            .getByLanguage("vi")
                            .map((subtitle) {
                          VoidCallback? onPressed;
                          if (subtitle != _selectedSubtitle) {
                            onPressed = () => selectSubtitle(subtitle);
                          }
                          return ElevatedButton(
                            onPressed: onPressed,
                            child: Text(subtitle.name),
                          );
                        }).toList(),
                      ),
                      const SizedBox(height: 8),
                      if (englishSubtitles.isNotEmpty)
                        Text("English", style: subtitleStyle),
                      Wrap(
                        spacing: 8,
                        children: widget.video.subtitles
                            .getByLanguage("en")
                            .map((subtitle) {
                          VoidCallback? onPressed;
                          if (subtitle != _selectedSubtitle) {
                            onPressed = () => selectSubtitle(subtitle);
                          }
                          return ElevatedButton(
                            onPressed: onPressed,
                            child: Text(subtitle.name),
                          );
                        }).toList(),
                      ),
                    ],
                  ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}

extension on List<phimtor_api.Subtitle> {
  List<phimtor_api.Subtitle> getByLanguage(String language) {
    var result = where((subtitle) => subtitle.language == language).toList();
    result.sort((a, b) => b.priority.compareTo(a.priority));
    return result;
  }
}
