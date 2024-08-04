//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class Series {
  /// Returns a new [Series] instance.
  Series({
    required this.id,
    required this.title,
    required this.originalTitle,
    required this.description,
    required this.posterLink,
    required this.releaseYear,
    required this.score,
    required this.durationInMinutes,
    required this.totalEpisodes,
    required this.currentEpisode,
    this.episodes = const [],
  });

  int id;

  String title;

  String originalTitle;

  String description;

  String posterLink;

  int releaseYear;

  num score;

  int durationInMinutes;

  int totalEpisodes;

  int currentEpisode;

  List<Episode> episodes;

  @override
  bool operator ==(Object other) => identical(this, other) || other is Series &&
    other.id == id &&
    other.title == title &&
    other.originalTitle == originalTitle &&
    other.description == description &&
    other.posterLink == posterLink &&
    other.releaseYear == releaseYear &&
    other.score == score &&
    other.durationInMinutes == durationInMinutes &&
    other.totalEpisodes == totalEpisodes &&
    other.currentEpisode == currentEpisode &&
    _deepEquality.equals(other.episodes, episodes);

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (id.hashCode) +
    (title.hashCode) +
    (originalTitle.hashCode) +
    (description.hashCode) +
    (posterLink.hashCode) +
    (releaseYear.hashCode) +
    (score.hashCode) +
    (durationInMinutes.hashCode) +
    (totalEpisodes.hashCode) +
    (currentEpisode.hashCode) +
    (episodes.hashCode);

  @override
  String toString() => 'Series[id=$id, title=$title, originalTitle=$originalTitle, description=$description, posterLink=$posterLink, releaseYear=$releaseYear, score=$score, durationInMinutes=$durationInMinutes, totalEpisodes=$totalEpisodes, currentEpisode=$currentEpisode, episodes=$episodes]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'id'] = this.id;
      json[r'title'] = this.title;
      json[r'originalTitle'] = this.originalTitle;
      json[r'description'] = this.description;
      json[r'posterLink'] = this.posterLink;
      json[r'releaseYear'] = this.releaseYear;
      json[r'score'] = this.score;
      json[r'durationInMinutes'] = this.durationInMinutes;
      json[r'totalEpisodes'] = this.totalEpisodes;
      json[r'currentEpisode'] = this.currentEpisode;
      json[r'episodes'] = this.episodes;
    return json;
  }

  /// Returns a new [Series] instance and imports its values from
  /// [value] if it's a [Map], null otherwise.
  // ignore: prefer_constructors_over_static_methods
  static Series? fromJson(dynamic value) {
    if (value is Map) {
      final json = value.cast<String, dynamic>();

      // Ensure that the map contains the required keys.
      // Note 1: the values aren't checked for validity beyond being non-null.
      // Note 2: this code is stripped in release mode!
      assert(() {
        requiredKeys.forEach((key) {
          assert(json.containsKey(key), 'Required key "Series[$key]" is missing from JSON.');
          assert(json[key] != null, 'Required key "Series[$key]" has a null value in JSON.');
        });
        return true;
      }());

      return Series(
        id: mapValueOfType<int>(json, r'id')!,
        title: mapValueOfType<String>(json, r'title')!,
        originalTitle: mapValueOfType<String>(json, r'originalTitle')!,
        description: mapValueOfType<String>(json, r'description')!,
        posterLink: mapValueOfType<String>(json, r'posterLink')!,
        releaseYear: mapValueOfType<int>(json, r'releaseYear')!,
        score: num.parse('${json[r'score']}'),
        durationInMinutes: mapValueOfType<int>(json, r'durationInMinutes')!,
        totalEpisodes: mapValueOfType<int>(json, r'totalEpisodes')!,
        currentEpisode: mapValueOfType<int>(json, r'currentEpisode')!,
        episodes: Episode.listFromJson(json[r'episodes']),
      );
    }
    return null;
  }

  static List<Series> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <Series>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = Series.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }

  static Map<String, Series> mapFromJson(dynamic json) {
    final map = <String, Series>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = Series.fromJson(entry.value);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  // maps a json object with a list of Series-objects as value to a dart map
  static Map<String, List<Series>> mapListFromJson(dynamic json, {bool growable = false,}) {
    final map = <String, List<Series>>{};
    if (json is Map && json.isNotEmpty) {
      // ignore: parameter_assignments
      json = json.cast<String, dynamic>();
      for (final entry in json.entries) {
        map[entry.key] = Series.listFromJson(entry.value, growable: growable,);
      }
    }
    return map;
  }

  /// The list of required keys that must be present in a JSON.
  static const requiredKeys = <String>{
    'id',
    'title',
    'originalTitle',
    'description',
    'posterLink',
    'releaseYear',
    'score',
    'durationInMinutes',
    'totalEpisodes',
    'currentEpisode',
    'episodes',
  };
}

