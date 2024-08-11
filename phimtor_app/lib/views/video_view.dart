import 'package:flutter/material.dart';
import 'package:media_kit/media_kit.dart';
import 'package:media_kit_video/media_kit_video.dart';
import 'package:phimtor_app/services/phimtor/phimtor_service.dart';
import 'package:phimtor_app/services/preferences/preferences_service.dart';
import 'package:phimtor_openapi_client/api.dart' as phimtor_api;
import 'package:torrent/torrent.dart' as torrent;

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

    updateVideoStreamUrl();
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
    if (oldWidget.torrentLink != widget.torrentLink) {
      updateVideoStreamUrl();
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
      child: Video(controller: controller),
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
