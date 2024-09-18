//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class TVSeasonEpisodesInner {
  /// Returns a new [TVSeasonEpisodesInner] instance.
  TVSeasonEpisodesInner({
    required this.id,
    required this.episodeNumber,
    required this.name,
    required this.overview,
    this.airDate,
    required this.runtime,
    required this.stillLink,
    required this.voteAverage,
    required this.videoID,
  });

  int id;

  int episodeNumber;

  String name;

  String overview;

  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  DateTime? airDate;

  int runtime;

  String stillLink;

  num voteAverage;

  int videoID;

  @override
  bool operator ==(Object other) => identical(this, other) || other is TVSeasonEpisodesInner &&
    other.id == id &&
    other.episodeNumber == episodeNumber &&
    other.name == name &&
    other.overview == overview &&
    other.airDate == airDate &&
    other.runtime == runtime &&
    other.stillLink == stillLink &&
    other.voteAverage == voteAverage &&
    other.videoID == videoID;

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (id.hashCode) +
    (episodeNumber.hashCode) +
    (name.hashCode) +
    (overview.hashCode) +
    (airDate == null ? 0 : airDate!.hashCode) +
    (runtime.hashCode) +
    (stillLink.hashCode) +
    (voteAverage.hashCode) +
    (videoID.hashCode);

  @override
  String toString() => 'TVSeasonEpisodesInner[id=$id, episodeNumber=$episodeNumber, name=$name, overview=$overview, airDate=$airDate, runtime=$runtime, stillLink=$stillLink, voteAverage=$voteAverage, videoID=$videoID]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'id'] = this.id;
      json[r'episodeNumber'] = this.episodeNumber;
      json[r'name'] = this.name;
      json[r'overview'] = this.overview;
    if (this.airDate != null) {
      json[r'airDate'] = _dateFormatter.format(this.airDate!.toUtc());
    } else {
      json[r'airDate'] = null;
    }
      json[r'runtime'] = this.runtime;
      json[r'stillLink'] = this.stillLink;
      json[r'voteAverage'] = this.voteAverage;
      json[r'videoID'] = this.videoID;
    return json;
  }

  /// Returns a new [TVSeasonEpisodesInner] instance and imports its values from
  /// [value] if it's a [Map], null otherwise.
  // ignore: prefer_constructors_over_static_methods
  static TVSeasonEpisodesInner? fromJson(dynamic value) {
    if (value is Map) {
      final json = value.cast<String, dynamic>();

      // Ensure that the map contains the required keys.
      // Note 1: the values aren't checked for validity beyond being non-null.
      // Note 2: this code is stripped in release mode!
      assert(() {
        requiredKeys.forEach((key) {
          assert(json.containsKey(key), 'Required key "TVSeasonEpisodesInner[$key]" is missing from JSON.');
          assert(json[key] != null, 'Required key "TVSeasonEpisodesInner[$key]" has a null value in JSON.');
        });
        return true;
      }());

      return TVSeasonEpisodesInner(
        id: mapValueOfType<int>(json, r'id')!,
        episodeNumber: mapValueOfType<int>(json, r'episodeNumber')!,
        name: mapValueOfType<String>(json, r'name')!,
        overview: mapValueOfType<String>(json, r'overview')!,
        airDate: mapDateTime(json, r'airDate', r''),
        runtime: mapValueOfType<int>(json, r'runtime')!,
        stillLink: mapValueOfType<String>(json, r'stillLink')!,
        voteAverage: num.parse('${json[r'voteAverage']}'),
        videoID: mapValueOfType<int>(json, r'videoID')!,
      );
    }
    return null;
  }

  static List<TVSeasonEpisodesInner> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <TVSeasonEpisodesInner>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = TVSeasonEpisodesInner.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }

  static Map<String, TVSeasonEpisodesInner> mapFromJson(dynamic json) {
    final map = <String, TVSeasonEpisodesInner>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = TVSeasonEpisodesInner.fromJson(entry.value);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  // maps a json object with a list of TVSeasonEpisodesInner-objects as value to a dart map
  static Map<String, List<TVSeasonEpisodesInner>> mapListFromJson(dynamic json, {bool growable = false,}) {
    final map = <String, List<TVSeasonEpisodesInner>>{};
    if (json is Map && json.isNotEmpty) {
      // ignore: parameter_assignments
      json = json.cast<String, dynamic>();
      for (final entry in json.entries) {
        map[entry.key] = TVSeasonEpisodesInner.listFromJson(entry.value, growable: growable,);
      }
    }
    return map;
  }

  /// The list of required keys that must be present in a JSON.
  static const requiredKeys = <String>{
    'id',
    'episodeNumber',
    'name',
    'overview',
    'runtime',
    'stillLink',
    'voteAverage',
    'videoID',
  };
}

