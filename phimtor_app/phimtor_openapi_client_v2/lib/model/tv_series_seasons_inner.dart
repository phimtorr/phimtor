//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class TvSeriesSeasonsInner {
  /// Returns a new [TvSeriesSeasonsInner] instance.
  TvSeriesSeasonsInner({
    required this.id,
    required this.seasonNumber,
    required this.name,
    required this.posterLink,
    required this.overview,
    this.airDate,
    required this.voteAverage,
  });

  int id;

  int seasonNumber;

  String name;

  String posterLink;

  String overview;

  ///
  /// Please note: This property should have been non-nullable! Since the specification file
  /// does not include a default value (using the "default:" property), however, the generated
  /// source code must fall back to having a nullable type.
  /// Consider adding a "default:" property in the specification file to hide this note.
  ///
  DateTime? airDate;

  num voteAverage;

  @override
  bool operator ==(Object other) => identical(this, other) || other is TvSeriesSeasonsInner &&
    other.id == id &&
    other.seasonNumber == seasonNumber &&
    other.name == name &&
    other.posterLink == posterLink &&
    other.overview == overview &&
    other.airDate == airDate &&
    other.voteAverage == voteAverage;

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (id.hashCode) +
    (seasonNumber.hashCode) +
    (name.hashCode) +
    (posterLink.hashCode) +
    (overview.hashCode) +
    (airDate == null ? 0 : airDate!.hashCode) +
    (voteAverage.hashCode);

  @override
  String toString() => 'TvSeriesSeasonsInner[id=$id, seasonNumber=$seasonNumber, name=$name, posterLink=$posterLink, overview=$overview, airDate=$airDate, voteAverage=$voteAverage]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'id'] = this.id;
      json[r'seasonNumber'] = this.seasonNumber;
      json[r'name'] = this.name;
      json[r'posterLink'] = this.posterLink;
      json[r'overview'] = this.overview;
    if (this.airDate != null) {
      json[r'airDate'] = _dateFormatter.format(this.airDate!.toUtc());
    } else {
      json[r'airDate'] = null;
    }
      json[r'voteAverage'] = this.voteAverage;
    return json;
  }

  /// Returns a new [TvSeriesSeasonsInner] instance and imports its values from
  /// [value] if it's a [Map], null otherwise.
  // ignore: prefer_constructors_over_static_methods
  static TvSeriesSeasonsInner? fromJson(dynamic value) {
    if (value is Map) {
      final json = value.cast<String, dynamic>();

      // Ensure that the map contains the required keys.
      // Note 1: the values aren't checked for validity beyond being non-null.
      // Note 2: this code is stripped in release mode!
      assert(() {
        requiredKeys.forEach((key) {
          assert(json.containsKey(key), 'Required key "TvSeriesSeasonsInner[$key]" is missing from JSON.');
          assert(json[key] != null, 'Required key "TvSeriesSeasonsInner[$key]" has a null value in JSON.');
        });
        return true;
      }());

      return TvSeriesSeasonsInner(
        id: mapValueOfType<int>(json, r'id')!,
        seasonNumber: mapValueOfType<int>(json, r'seasonNumber')!,
        name: mapValueOfType<String>(json, r'name')!,
        posterLink: mapValueOfType<String>(json, r'posterLink')!,
        overview: mapValueOfType<String>(json, r'overview')!,
        airDate: mapDateTime(json, r'airDate', r''),
        voteAverage: num.parse('${json[r'voteAverage']}'),
      );
    }
    return null;
  }

  static List<TvSeriesSeasonsInner> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <TvSeriesSeasonsInner>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = TvSeriesSeasonsInner.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }

  static Map<String, TvSeriesSeasonsInner> mapFromJson(dynamic json) {
    final map = <String, TvSeriesSeasonsInner>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = TvSeriesSeasonsInner.fromJson(entry.value);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  // maps a json object with a list of TvSeriesSeasonsInner-objects as value to a dart map
  static Map<String, List<TvSeriesSeasonsInner>> mapListFromJson(dynamic json, {bool growable = false,}) {
    final map = <String, List<TvSeriesSeasonsInner>>{};
    if (json is Map && json.isNotEmpty) {
      // ignore: parameter_assignments
      json = json.cast<String, dynamic>();
      for (final entry in json.entries) {
        map[entry.key] = TvSeriesSeasonsInner.listFromJson(entry.value, growable: growable,);
      }
    }
    return map;
  }

  /// The list of required keys that must be present in a JSON.
  static const requiredKeys = <String>{
    'id',
    'seasonNumber',
    'name',
    'posterLink',
    'overview',
    'voteAverage',
  };
}

