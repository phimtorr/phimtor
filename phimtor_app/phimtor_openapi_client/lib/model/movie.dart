//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class Movie {
  /// Returns a new [Movie] instance.
  Movie({
    required this.id,
    required this.title,
    required this.originalTitle,
    required this.description,
    required this.posterLink,
    required this.quantity,
    required this.releaseYear,
    required this.score,
    required this.durationInMinutes,
    required this.videoId,
  });

  int id;

  String title;

  String originalTitle;

  String description;

  String posterLink;

  String quantity;

  int releaseYear;

  num score;

  int durationInMinutes;

  int videoId;

  @override
  bool operator ==(Object other) => identical(this, other) || other is Movie &&
    other.id == id &&
    other.title == title &&
    other.originalTitle == originalTitle &&
    other.description == description &&
    other.posterLink == posterLink &&
    other.quantity == quantity &&
    other.releaseYear == releaseYear &&
    other.score == score &&
    other.durationInMinutes == durationInMinutes &&
    other.videoId == videoId;

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (id.hashCode) +
    (title.hashCode) +
    (originalTitle.hashCode) +
    (description.hashCode) +
    (posterLink.hashCode) +
    (quantity.hashCode) +
    (releaseYear.hashCode) +
    (score.hashCode) +
    (durationInMinutes.hashCode) +
    (videoId.hashCode);

  @override
  String toString() => 'Movie[id=$id, title=$title, originalTitle=$originalTitle, description=$description, posterLink=$posterLink, quantity=$quantity, releaseYear=$releaseYear, score=$score, durationInMinutes=$durationInMinutes, videoId=$videoId]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'id'] = this.id;
      json[r'title'] = this.title;
      json[r'originalTitle'] = this.originalTitle;
      json[r'description'] = this.description;
      json[r'posterLink'] = this.posterLink;
      json[r'quantity'] = this.quantity;
      json[r'releaseYear'] = this.releaseYear;
      json[r'score'] = this.score;
      json[r'durationInMinutes'] = this.durationInMinutes;
      json[r'videoId'] = this.videoId;
    return json;
  }

  /// Returns a new [Movie] instance and imports its values from
  /// [value] if it's a [Map], null otherwise.
  // ignore: prefer_constructors_over_static_methods
  static Movie? fromJson(dynamic value) {
    if (value is Map) {
      final json = value.cast<String, dynamic>();

      // Ensure that the map contains the required keys.
      // Note 1: the values aren't checked for validity beyond being non-null.
      // Note 2: this code is stripped in release mode!
      assert(() {
        requiredKeys.forEach((key) {
          assert(json.containsKey(key), 'Required key "Movie[$key]" is missing from JSON.');
          assert(json[key] != null, 'Required key "Movie[$key]" has a null value in JSON.');
        });
        return true;
      }());

      return Movie(
        id: mapValueOfType<int>(json, r'id')!,
        title: mapValueOfType<String>(json, r'title')!,
        originalTitle: mapValueOfType<String>(json, r'originalTitle')!,
        description: mapValueOfType<String>(json, r'description')!,
        posterLink: mapValueOfType<String>(json, r'posterLink')!,
        quantity: mapValueOfType<String>(json, r'quantity')!,
        releaseYear: mapValueOfType<int>(json, r'releaseYear')!,
        score: num.parse('${json[r'score']}'),
        durationInMinutes: mapValueOfType<int>(json, r'durationInMinutes')!,
        videoId: mapValueOfType<int>(json, r'videoId')!,
      );
    }
    return null;
  }

  static List<Movie> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <Movie>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = Movie.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }

  static Map<String, Movie> mapFromJson(dynamic json) {
    final map = <String, Movie>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = Movie.fromJson(entry.value);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  // maps a json object with a list of Movie-objects as value to a dart map
  static Map<String, List<Movie>> mapListFromJson(dynamic json, {bool growable = false,}) {
    final map = <String, List<Movie>>{};
    if (json is Map && json.isNotEmpty) {
      // ignore: parameter_assignments
      json = json.cast<String, dynamic>();
      for (final entry in json.entries) {
        map[entry.key] = Movie.listFromJson(entry.value, growable: growable,);
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
    'quantity',
    'releaseYear',
    'score',
    'durationInMinutes',
    'videoId',
  };
}

