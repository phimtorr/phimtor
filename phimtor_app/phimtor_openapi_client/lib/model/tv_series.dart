//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class TvSeries {
  /// Returns a new [TvSeries] instance.
  TvSeries({
    required this.id,
    required this.name,
    required this.originalName,
    required this.status,
    required this.tagline,
    this.genres = const [],
    required this.overview,
    required this.posterLink,
    required this.backdropLink,
    this.firstAirDate,
    this.lastAirDate,
    required this.voteAverage,
    required this.numberOfSeasons,
    required this.numberOfEpisodes,
    this.seasons = const [],
  });

  int id;

  String name;

  String originalName;

  String status;

  String tagline;

  List<Genre> genres;

  String overview;

  String posterLink;

  String backdropLink;

  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  DateTime? firstAirDate;

  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  DateTime? lastAirDate;

  num voteAverage;

  int numberOfSeasons;

  int numberOfEpisodes;

  List<TvSeriesSeasonsInner> seasons;

  @override
  bool operator ==(Object other) => identical(this, other) || other is TvSeries &&
    other.id == id &&
    other.name == name &&
    other.originalName == originalName &&
    other.status == status &&
    other.tagline == tagline &&
    _deepEquality.equals(other.genres, genres) &&
    other.overview == overview &&
    other.posterLink == posterLink &&
    other.backdropLink == backdropLink &&
    other.firstAirDate == firstAirDate &&
    other.lastAirDate == lastAirDate &&
    other.voteAverage == voteAverage &&
    other.numberOfSeasons == numberOfSeasons &&
    other.numberOfEpisodes == numberOfEpisodes &&
    _deepEquality.equals(other.seasons, seasons);

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (id.hashCode) +
    (name.hashCode) +
    (originalName.hashCode) +
    (status.hashCode) +
    (tagline.hashCode) +
    (genres.hashCode) +
    (overview.hashCode) +
    (posterLink.hashCode) +
    (backdropLink.hashCode) +
    (firstAirDate == null ? 0 : firstAirDate!.hashCode) +
    (lastAirDate == null ? 0 : lastAirDate!.hashCode) +
    (voteAverage.hashCode) +
    (numberOfSeasons.hashCode) +
    (numberOfEpisodes.hashCode) +
    (seasons.hashCode);

  @override
  String toString() => 'TvSeries[id=$id, name=$name, originalName=$originalName, status=$status, tagline=$tagline, genres=$genres, overview=$overview, posterLink=$posterLink, backdropLink=$backdropLink, firstAirDate=$firstAirDate, lastAirDate=$lastAirDate, voteAverage=$voteAverage, numberOfSeasons=$numberOfSeasons, numberOfEpisodes=$numberOfEpisodes, seasons=$seasons]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'id'] = this.id;
      json[r'name'] = this.name;
      json[r'originalName'] = this.originalName;
      json[r'status'] = this.status;
      json[r'tagline'] = this.tagline;
      json[r'genres'] = this.genres;
      json[r'overview'] = this.overview;
      json[r'posterLink'] = this.posterLink;
      json[r'backdropLink'] = this.backdropLink;
    if (this.firstAirDate != null) {
      json[r'firstAirDate'] = _dateFormatter.format(this.firstAirDate!.toUtc());
    } else {
      json[r'firstAirDate'] = null;
    }
    if (this.lastAirDate != null) {
      json[r'lastAirDate'] = _dateFormatter.format(this.lastAirDate!.toUtc());
    } else {
      json[r'lastAirDate'] = null;
    }
      json[r'voteAverage'] = this.voteAverage;
      json[r'numberOfSeasons'] = this.numberOfSeasons;
      json[r'numberOfEpisodes'] = this.numberOfEpisodes;
      json[r'seasons'] = this.seasons;
    return json;
  }

  /// Returns a new [TvSeries] instance and imports its values from
  /// [value] if it's a [Map], null otherwise.
  // ignore: prefer_constructors_over_static_methods
  static TvSeries? fromJson(dynamic value) {
    if (value is Map) {
      final json = value.cast<String, dynamic>();

      // Ensure that the map contains the required keys.
      // Note 1: the values aren't checked for validity beyond being non-null.
      // Note 2: this code is stripped in release mode!
      assert(() {
        requiredKeys.forEach((key) {
          assert(json.containsKey(key), 'Required key "TvSeries[$key]" is missing from JSON.');
          assert(json[key] != null, 'Required key "TvSeries[$key]" has a null value in JSON.');
        });
        return true;
      }());

      return TvSeries(
        id: mapValueOfType<int>(json, r'id')!,
        name: mapValueOfType<String>(json, r'name')!,
        originalName: mapValueOfType<String>(json, r'originalName')!,
        status: mapValueOfType<String>(json, r'status')!,
        tagline: mapValueOfType<String>(json, r'tagline')!,
        genres: Genre.listFromJson(json[r'genres']),
        overview: mapValueOfType<String>(json, r'overview')!,
        posterLink: mapValueOfType<String>(json, r'posterLink')!,
        backdropLink: mapValueOfType<String>(json, r'backdropLink')!,
        firstAirDate: mapDateTime(json, r'firstAirDate', r''),
        lastAirDate: mapDateTime(json, r'lastAirDate', r''),
        voteAverage: num.parse('${json[r'voteAverage']}'),
        numberOfSeasons: mapValueOfType<int>(json, r'numberOfSeasons')!,
        numberOfEpisodes: mapValueOfType<int>(json, r'numberOfEpisodes')!,
        seasons: TvSeriesSeasonsInner.listFromJson(json[r'seasons']),
      );
    }
    return null;
  }

  static List<TvSeries> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <TvSeries>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = TvSeries.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }

  static Map<String, TvSeries> mapFromJson(dynamic json) {
    final map = <String, TvSeries>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = TvSeries.fromJson(entry.value);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  // maps a json object with a list of TvSeries-objects as value to a dart map
  static Map<String, List<TvSeries>> mapListFromJson(dynamic json, {bool growable = false,}) {
    final map = <String, List<TvSeries>>{};
    if (json is Map && json.isNotEmpty) {
      // ignore: parameter_assignments
      json = json.cast<String, dynamic>();
      for (final entry in json.entries) {
        map[entry.key] = TvSeries.listFromJson(entry.value, growable: growable,);
      }
    }
    return map;
  }

  /// The list of required keys that must be present in a JSON.
  static const requiredKeys = <String>{
    'id',
    'name',
    'originalName',
    'status',
    'tagline',
    'genres',
    'overview',
    'posterLink',
    'backdropLink',
    'voteAverage',
    'numberOfSeasons',
    'numberOfEpisodes',
    'seasons',
  };
}

