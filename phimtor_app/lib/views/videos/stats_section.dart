import 'dart:async';

import 'package:flutter/material.dart';
import 'package:torrent/torrent.dart' as torrent;

class StatsSection extends StatefulWidget {
  const StatsSection({
    super.key,
    required this.infoHash,
    required this.videoIndex,
  });

  final String infoHash;
  final int videoIndex;

  @override
  State<StatsSection> createState() => _StatsSectionState();
}

class _StatsSectionState extends State<StatsSection> {
  late final Timer _timer;

  torrent.Stats? _stats;

  @override
  void initState() {
    super.initState();

    _timer = Timer.periodic(const Duration(seconds: 1), (timer) async {
      await updateStats();
    });
  }

  @override
  void didUpdateWidget(covariant StatsSection oldWidget) {
    super.didUpdateWidget(oldWidget);
    setState(() {
      _stats = null;
    });
  }

  @override
  void dispose() {
    _timer.cancel();
    super.dispose();
  }

  Future<void> updateStats() async {
    final stats = await torrent.LibTorrent()
        .torrentApi
        .getTorrentStats(widget.infoHash, widget.videoIndex);
    setState(() {
      _stats = stats;
    });
  }

  bool isStatsAvailable() {
    return _stats != null && _stats!.length > 0;
  }

  @override
  Widget build(BuildContext context) {
    if (!isStatsAvailable()) {
      return const SizedBox();
    }

    return Text("Stats: ${_stats!.bytesCompleted} bytes completed");
  }
}
