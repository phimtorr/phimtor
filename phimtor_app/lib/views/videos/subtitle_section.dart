import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:media_kit/media_kit.dart';
import 'package:path/path.dart' as path;
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/views/components/buttons/need_verified_user_button.dart';
import 'package:phimtor_app/views/components/buttons/premium_button.dart';

import 'package:phimtor_openapi_client/api.dart' as phimtor_api;

typedef SelectSubtitleCallback = void Function(SubtitleTrack subtitle);

class SubtitleSection extends StatefulWidget {
  const SubtitleSection({
    super.key,
    required this.subtitles,
    required this.onSelectSubtitle,
  });

  final List<phimtor_api.Subtitle> subtitles;
  final SelectSubtitleCallback onSelectSubtitle;

  get vietnameseSubtitles => subtitles.getByLanguage("vi");
  get englishSubtitles => subtitles.getByLanguage("en");

  @override
  State<SubtitleSection> createState() => _SubtitleSectionState();
}

class _SubtitleSectionState extends State<SubtitleSection> {
  phimtor_api.Subtitle? _selectedSubtitle;

  @override
  void initState() {
    super.initState();

    Future.delayed(Duration.zero, () async {
      if (widget.vietnameseSubtitles.isNotEmpty) {
        selectSubtitle(widget.vietnameseSubtitles.first);
      }
    });
  }

  void selectSubtitle(phimtor_api.Subtitle subtitle) {
    setState(() {
      _selectedSubtitle = subtitle;
    });
    final subtitleTrack = SubtitleTrack.uri(
      subtitle.link,
      title: subtitle.name,
      language: subtitle.language,
    );
    widget.onSelectSubtitle(subtitleTrack);
  }

  void selectSubtitleFile() async {
    final result = await FilePicker.platform.pickFiles(
      type: FileType.custom,
      allowedExtensions: ["srt", "vtt", "ass"],
    );

    if (result == null) {
      return;
    }

    final subtitleFile = result.files.first;
    final subtitleTrack = SubtitleTrack.uri(
      path.toUri(subtitleFile.path!).toString(),
      title: subtitleFile.name,
    );

    setState(() {
      _selectedSubtitle = null;
    });

    widget.onSelectSubtitle(subtitleTrack);
  }

  void selectNoSubtitle() {
    setState(() {
      _selectedSubtitle = null;
    });
    widget.onSelectSubtitle(SubtitleTrack.no());
  }

  @override
  Widget build(BuildContext context) {
    var titleStyle = Theme.of(context).textTheme.headlineMedium!;
    var subtitleStyle = Theme.of(context).textTheme.headlineSmall!;

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Row(
          children: [
            Text(context.loc.subtitle, style: titleStyle),
            const Spacer(),
            ElevatedButton.icon(
              onPressed: selectNoSubtitle,
              label: Text(context.loc.unselect_subtitle),
              icon: const Icon(Icons.subtitles_off),
            ),
          ],
        ),
        const SizedBox(height: 8),
        PremiumButton(
          onPressed: selectSubtitleFile,
          label: Text(context.loc.select_subtitle_file),
          icon: const Icon(Icons.subtitles),
        ),
        const SizedBox(height: 8),
        if (widget.vietnameseSubtitles.isNotEmpty)
          Text(context.loc.subtitle_vietnamese, style: subtitleStyle),
        Wrap(
          spacing: 8,
          children: widget.vietnameseSubtitles.map<Widget>((subtitle) {
            VoidCallback? onPressed;
            if (subtitle != _selectedSubtitle) {
              onPressed = () => selectSubtitle(subtitle);
            }
            if (subtitle == widget.vietnameseSubtitles.first) {
              return ElevatedButton(
                onPressed: onPressed,
                child: Text(subtitle.name),
              );
            } else {
              return NeedVerifiedUserButton(
                onPressed: onPressed,
                child: Text(subtitle.name),
              );
            }
          }).toList(),
        ),
        const SizedBox(height: 8),
        if (widget.englishSubtitles.isNotEmpty)
          Text(context.loc.subtitle_english, style: subtitleStyle),
        Wrap(
          spacing: 8,
          children: widget.englishSubtitles.map<Widget>((subtitle) {
            VoidCallback? onPressed;
            if (subtitle != _selectedSubtitle) {
              onPressed = () => selectSubtitle(subtitle);
            }
            if (subtitle == widget.englishSubtitles.first) {
              return ElevatedButton(
                onPressed: onPressed,
                child: Text(subtitle.name),
              );
            } else {
              return NeedVerifiedUserButton(
                onPressed: onPressed,
                child: Text(subtitle.name),
              );
            }
          }).toList(),
        ),
      ],
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
