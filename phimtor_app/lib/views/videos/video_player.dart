import 'package:flutter/material.dart';
import 'package:media_kit/media_kit.dart';
import 'package:media_kit_video/media_kit_video.dart';
import 'package:phimtor_app/services/preferences/preferences_service.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;
import 'package:torrent/torrent.dart' as torrent;

class VideoPlayer extends StatefulWidget {
  const VideoPlayer({
    super.key,
    required this.torrentLink,
    this.subtitle,
  });

  final phimtor_api.TorrentLink torrentLink;
  final phimtor_api.Subtitle? subtitle;

  @override
  State<VideoPlayer> createState() => _VideoPlayerState();
}

class _VideoPlayerState extends State<VideoPlayer> {
  late final player = Player();
  late final controller = VideoController(player);

  bool isLoading = false;
  Exception? error;
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
          player.open(Media(_videoStreamUrl!));
        }
      }
    });

    Future.delayed(Duration.zero, () async {
      await updateVideoStreamUrl();
      await updateSubtitle();
    });
  }

  @override
  void dispose() {
    player.dispose();
    // remove torrent here
    super.dispose();
  }

  @override
  void didUpdateWidget(covariant VideoPlayer oldWidget) async {
    super.didUpdateWidget(oldWidget);
    if (oldWidget.torrentLink != widget.torrentLink) {
      await updateVideoStreamUrl();
    }
    if (oldWidget.subtitle != widget.subtitle) {
      await updateSubtitle();
    }
  }

  Future<void> updateVideoStreamUrl() async {
    setState(() {
      isLoading = true;
    });
    try {
      final torrentLink = widget.torrentLink;
      final addTorrentResp = await torrent.LibTorrent().torrentApi.addTorrent(
          torrent.AddTorrentRequest(
            url: torrentLink.link,
          ),
          dropOthers: true,
          deleteOthers: PreferencesService.getInstance().deleteAfterClose);
      if (addTorrentResp == null) {
        throw Exception('Failed to add torrent');
      }

      final videoIndex =
          addTorrentResp.torrent.getVideoIndex(torrentLink.fileIndex);
      final fileName =
          addTorrentResp.torrent.files[videoIndex].name.split('/').last;

      final videoStreamUrl = torrent.LibTorrent().getStreamVideoURL(
        addTorrentResp.torrent.infoHash,
        videoIndex,
        fileName,
      );

      _videoStreamUrl = videoStreamUrl;

      player.open(Media(videoStreamUrl));
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

  Future<void> updateSubtitle() async {
    print("updateSubtitle ${widget.subtitle.toString()}");
    if (widget.subtitle == null) {
      player.setSubtitleTrack(SubtitleTrack.no());
      return;
    }
    final subtitle = widget.subtitle!;
    final subtitleUrl = subtitle.link;
    player.setSubtitleTrack(
      SubtitleTrack.uri(
        subtitleUrl,
        title: subtitle.name,
        language: subtitle.language,
      ),
    );
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
        child: Text('Error: $error'),
      );
    }
    return AspectRatio(
      aspectRatio: 16.0 / 9.0,
      child: Video(
        controller: controller,
        controls: MaterialDesktopVideoControls,
      ),
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
