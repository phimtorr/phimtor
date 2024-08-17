import 'dart:async';

import 'package:flutter/material.dart';
import 'package:media_kit/media_kit.dart';
import 'package:media_kit_video/media_kit_video.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/views/videos/stats_section.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;
import 'package:torrent/torrent.dart' as torrent;

class VideoPlayer extends StatefulWidget {
  const VideoPlayer({
    super.key,
    required this.torrentLink,
    this.subtitle,
  });

  final phimtor_api.TorrentLink torrentLink;
  final SubtitleTrack? subtitle;

  @override
  State<VideoPlayer> createState() => _VideoPlayerState();
}

class _VideoPlayerState extends State<VideoPlayer> {
  late final player = Player();
  late final controller = VideoController(player);

  bool isLoading = false;
  Exception? error;
  String? _infoHash;
  int? _videoIndex;
  String? _videoStreamUrl;

  @override
  void initState() {
    super.initState();
    // show log
    player.stream.log.listen((log) {
      debugPrint("Player log: $log");
    });
    // because the torrent file is not ready yet, retry after 3 seconds
    player.stream.error.listen((error) async {
      debugPrint("Player error: $error");
      if (error.contains('Failed to open ')) {
        if (_videoStreamUrl != null) {
          await Future.delayed(const Duration(seconds: 3));
          openVideo();
        }
      }
    });

    Future.delayed(Duration.zero, () async {
      await updateVideoStreamUrl();
    });
  }

  @override
  void dispose() {
    player.dispose();
    // remove torrent here
    super.dispose();
  }

  @override
  void didUpdateWidget(covariant VideoPlayer oldWidget) {
    super.didUpdateWidget(oldWidget);
    Future.delayed(Duration.zero, () async {
      if (oldWidget.torrentLink != widget.torrentLink) {
        await updateVideoStreamUrl();
      }
      if (oldWidget.subtitle != widget.subtitle) {
        updateSubtitle();
      }
    });
  }

  Future<void> updateVideoStreamUrl() async {
    setState(() {
      isLoading = true;
    });
    try {
      final torrentLink = widget.torrentLink;
      final torrentFile = await torrent.LibTorrent().torrentApi.addTorrent(
          torrent.AddTorrentRequest(
            link: torrentLink.link,
          ),
          dropOthers: true,
          deleteOthers: false);
      if (torrentFile == null) {
        throw Exception('Failed to add torrent');
      }

      final videoIndex = torrentFile.getVideoIndex(torrentLink.fileIndex);
      final fileName = torrentFile.files[videoIndex].name.split('/').last;

      final videoStreamUrl = torrent.LibTorrent().getStreamVideoURL(
        torrentFile.infoHash,
        videoIndex,
        fileName,
      );

      setState(() {
        _infoHash = torrentFile.infoHash;
        _videoIndex = videoIndex;
        _videoStreamUrl = videoStreamUrl;
      });

      openVideo();
    } on Exception catch (e) {
      debugPrint('Failed to open video: $e');
      setState(() {
        error = e;
      });
    } finally {
      setState(() {
        isLoading = false;
      });
    }
  }

  void openVideo() {
    if (_videoStreamUrl != null) {
      player.open(Media(_videoStreamUrl!));
      updateSubtitle();
    }
  }

  void updateSubtitle() {
    player.setSubtitleTrack(widget.subtitle ?? SubtitleTrack.no());
  }

  bool isTorrentIsAdded() {
    return _infoHash != null && _videoIndex != null;
  }

  @override
  Widget build(BuildContext context) {
    if (isLoading) {
      return const Center(
        child: CircularProgressIndicator(),
      );
    }
    if (error != null) {
      return Center(
        child: Text(context.loc.error(error.toString())),
      );
    }
    return Column(
      children: [
        SizedBox(
          width: MediaQuery.of(context).size.width,
          height: MediaQuery.of(context).size.width * 9.0 / 16.0,
          child: AspectRatio(
            aspectRatio: 16.0 / 9.0,
            child: Video(
              controller: controller,
              controls: MaterialDesktopVideoControls,
            ),
          ),
        ),
        Padding(
          padding: const EdgeInsets.fromLTRB(16, 8, 16, 0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              if (isTorrentIsAdded())
                StatsSection(
                  infoHash: _infoHash!,
                  videoIndex: _videoIndex!,
                ),
              if (widget.subtitle != null && widget.subtitle!.id != "no")
                Text(
                  "${context.loc.subtitle}: ${widget.subtitle!.title}",
                  style: Theme.of(context).textTheme.labelMedium,
                ),
            ],
          ),
        ),
      ],
    );
  }
}

extension on torrent.Torrent {
  bool _isVideoFile(torrent.File file) {
    final ext = file.name.split('.').last;
    return ext == 'mp4' || ext == 'mkv' || ext == 'avi';
  }

  int getVideoIndex(int suggestIndex) {
    var index = suggestIndex;
    if (index >= 0 && index < files.length) {
      if (_isVideoFile(files[index])) {
        return index;
      }
    }

    for (var i = 0; i < files.length; i++) {
      if (_isVideoFile(files[i])) {
        return i;
      }
    }
    throw Exception('No video file found');
  }
}
