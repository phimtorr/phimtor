//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class Video {
  /// Returns a new [Video] instance.
  Video({
    required this.id,
    required this.title,
    this.torrentLinks = const [],
    this.premiumTorrentLinks = const [],
    this.subtitles = const [],
  });

  int id;

  String title;

  List<TorrentLink> torrentLinks;

  List<PremiumTorrentLink> premiumTorrentLinks;

  List<Subtitle> subtitles;

  @override
  bool operator ==(Object other) => identical(this, other) || other is Video &&
    other.id == id &&
    other.title == title &&
    _deepEquality.equals(other.torrentLinks, torrentLinks) &&
    _deepEquality.equals(other.premiumTorrentLinks, premiumTorrentLinks) &&
    _deepEquality.equals(other.subtitles, subtitles);

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (id.hashCode) +
    (title.hashCode) +
    (torrentLinks.hashCode) +
    (premiumTorrentLinks.hashCode) +
    (subtitles.hashCode);

  @override
  String toString() => 'Video[id=$id, title=$title, torrentLinks=$torrentLinks, premiumTorrentLinks=$premiumTorrentLinks, subtitles=$subtitles]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'id'] = this.id;
      json[r'title'] = this.title;
      json[r'torrentLinks'] = this.torrentLinks;
      json[r'premiumTorrentLinks'] = this.premiumTorrentLinks;
      json[r'subtitles'] = this.subtitles;
    return json;
  }

  /// Returns a new [Video] instance and imports its values from
  /// [value] if it's a [Map], null otherwise.
  // ignore: prefer_constructors_over_static_methods
  static Video? fromJson(dynamic value) {
    if (value is Map) {
      final json = value.cast<String, dynamic>();

      // Ensure that the map contains the required keys.
      // Note 1: the values aren't checked for validity beyond being non-null.
      // Note 2: this code is stripped in release mode!
      assert(() {
        requiredKeys.forEach((key) {
          assert(json.containsKey(key), 'Required key "Video[$key]" is missing from JSON.');
          assert(json[key] != null, 'Required key "Video[$key]" has a null value in JSON.');
        });
        return true;
      }());

      return Video(
        id: mapValueOfType<int>(json, r'id')!,
        title: mapValueOfType<String>(json, r'title')!,
        torrentLinks: TorrentLink.listFromJson(json[r'torrentLinks']),
        premiumTorrentLinks: PremiumTorrentLink.listFromJson(json[r'premiumTorrentLinks']),
        subtitles: Subtitle.listFromJson(json[r'subtitles']),
      );
    }
    return null;
  }

  static List<Video> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <Video>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = Video.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }

  static Map<String, Video> mapFromJson(dynamic json) {
    final map = <String, Video>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = Video.fromJson(entry.value);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  // maps a json object with a list of Video-objects as value to a dart map
  static Map<String, List<Video>> mapListFromJson(dynamic json, {bool growable = false,}) {
    final map = <String, List<Video>>{};
    if (json is Map && json.isNotEmpty) {
      // ignore: parameter_assignments
      json = json.cast<String, dynamic>();
      for (final entry in json.entries) {
        map[entry.key] = Video.listFromJson(entry.value, growable: growable,);
      }
    }
    return map;
  }

  /// The list of required keys that must be present in a JSON.
  static const requiredKeys = <String>{
    'id',
    'title',
    'torrentLinks',
    'premiumTorrentLinks',
    'subtitles',
  };
}

