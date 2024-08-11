import 'dart:async';

import 'package:flutter/material.dart';
import 'package:torrent/torrent.dart' as torrent;
import 'package:pretty_bytes/pretty_bytes.dart';

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
  int _verlocityBytesPerSecond = 0;

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
    final newStats = await torrent.LibTorrent()
        .torrentApi
        .getTorrentStats(widget.infoHash, widget.videoIndex);
    if (newStats == null || newStats.isEmpty) {
      resetStats();
      return;
    }
    setState(() {
      _verlocityBytesPerSecond = newStats.bytesCompleted - ( _stats?.bytesCompleted ?? 0);
      _stats = newStats;
    });
  }

  void resetStats() {
    setState(() {
      _verlocityBytesPerSecond = 0;
      _stats = null;
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

    return Text("Stats: ${_stats!.prettyBytesCompleted} / ${_stats!.prettyLength} bytes completed");
  }
}

extension on torrent.Stats {
  bool get isEmpty => length == 0;
  String get prettyBytesCompleted {
    return prettyBytes(bytesCompleted.toDouble());
  }
  String get prettyLength {
    return prettyBytes(length.toDouble());
  }
}
