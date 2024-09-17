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
    required this.status,
    required this.tagline,
    this.genres = const [],
    required this.overview,
    required this.posterLink,
    required this.backdropLink,
    required this.releaseDate,
    required this.runtime,
    required this.voteAverage,
    required this.videoID,
  });

  int id;

  String title;

  String originalTitle;

  String status;

  String tagline;

  List<Genre> genres;

  String overview;

  String posterLink;

  String backdropLink;

  DateTime releaseDate;

  int runtime;

  num voteAverage;

  int videoID;

  @override
  bool operator ==(Object other) => identical(this, other) || other is Movie &&
    other.id == id &&
    other.title == title &&
    other.originalTitle == originalTitle &&
    other.status == status &&
    other.tagline == tagline &&
    _deepEquality.equals(other.genres, genres) &&
    other.overview == overview &&
    other.posterLink == posterLink &&
    other.backdropLink == backdropLink &&
    other.releaseDate == releaseDate &&
    other.runtime == runtime &&
    other.voteAverage == voteAverage &&
    other.videoID == videoID;

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (id.hashCode) +
    (title.hashCode) +
    (originalTitle.hashCode) +
    (status.hashCode) +
    (tagline.hashCode) +
    (genres.hashCode) +
    (overview.hashCode) +
    (posterLink.hashCode) +
    (backdropLink.hashCode) +
    (releaseDate.hashCode) +
    (runtime.hashCode) +
    (voteAverage.hashCode) +
    (videoID.hashCode);

  @override
  String toString() => 'Movie[id=$id, title=$title, originalTitle=$originalTitle, status=$status, tagline=$tagline, genres=$genres, overview=$overview, posterLink=$posterLink, backdropLink=$backdropLink, releaseDate=$releaseDate, runtime=$runtime, voteAverage=$voteAverage, videoID=$videoID]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'id'] = this.id;
      json[r'title'] = this.title;
      json[r'originalTitle'] = this.originalTitle;
      json[r'status'] = this.status;
      json[r'tagline'] = this.tagline;
      json[r'genres'] = this.genres;
      json[r'overview'] = this.overview;
      json[r'posterLink'] = this.posterLink;
      json[r'backdropLink'] = this.backdropLink;
      json[r'releaseDate'] = _dateFormatter.format(this.releaseDate.toUtc());
      json[r'runtime'] = this.runtime;
      json[r'voteAverage'] = this.voteAverage;
      json[r'videoID'] = this.videoID;
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
        status: mapValueOfType<String>(json, r'status')!,
        tagline: mapValueOfType<String>(json, r'tagline')!,
        genres: Genre.listFromJson(json[r'genres']),
        overview: mapValueOfType<String>(json, r'overview')!,
        posterLink: mapValueOfType<String>(json, r'posterLink')!,
        backdropLink: mapValueOfType<String>(json, r'backdropLink')!,
        releaseDate: mapDateTime(json, r'releaseDate', r'')!,
        runtime: mapValueOfType<int>(json, r'runtime')!,
        voteAverage: num.parse('${json[r'voteAverage']}'),
        videoID: mapValueOfType<int>(json, r'videoID')!,
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
    'status',
    'tagline',
    'genres',
    'overview',
    'posterLink',
    'backdropLink',
    'releaseDate',
    'runtime',
    'voteAverage',
    'videoID',
  };
}

