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
    if (oldWidget.infoHash != widget.infoHash || oldWidget.videoIndex != widget.videoIndex) {
      resetStats();
    }
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
      _verlocityBytesPerSecond =
          newStats.bytesCompleted - (_stats?.bytesCompleted ?? 0);
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

    final textStyle = Theme.of(context).textTheme.labelSmall;

    return Row(
      children: [
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text("Total peers: ${_stats!.totalPeers}", style: textStyle),
            Text("Active: ${_stats!.activePeers}", style: textStyle),
            Text("Connected: ${_stats!.connectedPeers}", style: textStyle),
          ],
        ),
        const SizedBox(width: 8),
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text("", style: textStyle,), // Spacer
            Text("Pending: ${_stats!.pendingPeers}", style: textStyle),
            Text("Half open: ${_stats!.halfOpenPeers}", style: textStyle),
          ],
        ),
        const SizedBox(width: 8),
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
                'Downloaded: ${_stats!.prettyBytesCompleted} / ${_stats!.prettyLength}',
                style: textStyle),
            Text('Progress: ${_stats!.progressPercentage.toStringAsFixed(2)}%',
                style: textStyle),
            Text(
                'Download speed: ${prettyBytes(_verlocityBytesPerSecond.toDouble())}/s',
                style: textStyle),
          ],
        ),
      ],
    );
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

  double get progressPercentage {
    return bytesCompleted.toDouble() / length.toDouble() * 100;
  }
}
