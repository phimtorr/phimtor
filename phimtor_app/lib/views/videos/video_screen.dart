import 'package:flutter/material.dart';
import 'package:media_kit/media_kit.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/views/components/buttons/need_verified_user_button.dart';
import 'package:phimtor_app/views/components/buttons/premium_button.dart';
import 'package:phimtor_app/views/videos/subtitle_section.dart';
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
  SubtitleTrack _subtitleTrack = SubtitleTrack.no();

  @override
  void initState() {
    super.initState();

    selectTorrentLink(widget.video.torrentLinks.first);
  }

  void selectTorrentLink(phimtor_api.TorrentLink torrentLink) {
    setState(() {
      _selectedTorrentLink = torrentLink;
    });
  }

  void setSubtitleTrack(SubtitleTrack subtitleTrack) {
    setState(() {
      _subtitleTrack = subtitleTrack;
    });
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          VideoPlayer(
            torrentLink: _selectedTorrentLink!,
            subtitle: _subtitleTrack,
          ),
          buildTorrentLinksSection(context),
          SubtitleSection(
            subtitles: widget.video.subtitles,
            onSelectSubtitle: setSubtitleTrack,
          ),
        ],
      ),
    );
  }

  Widget buildTorrentLinksSection(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(context.loc.torrent_links,
              style: Theme.of(context).textTheme.headlineMedium),
          Wrap(
            spacing: 8,
            children: [
              ...widget.video.torrentLinks.map((link) {
                VoidCallback? onPressed;
                if (link != _selectedTorrentLink) {
                  onPressed = () => selectTorrentLink(link);
                }

                if (link == widget.video.torrentLinks.first) {
                  return ElevatedButton(
                    onPressed: onPressed,
                    child: Text(link.name),
                  );
                } else {
                  return NeedVerifiedUserButton(
                    onPressed: onPressed,
                    child: Text(link.name),
                  );
                }
              }),
              ...widget.video.premiumTorrentLinks.map((link) {
                return PremiumButton(
                  onPressed: () {},
                  label: Text(link.name),
                );
              }),
            ],
          ),
        ],
      ),
    );
  }
}
