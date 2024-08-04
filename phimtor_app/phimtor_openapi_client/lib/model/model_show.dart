//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ModelShow {
  /// Returns a new [ModelShow] instance.
  ModelShow({
    required this.id,
    required this.title,
    required this.originalTitle,
    required this.posterLink,
    required this.type,
    required this.releaseYear,
    required this.score,
    required this.durationInMinutes,
    required this.quantity,
    required this.totalEpisodes,
    required this.currentEpisode,
  });

  int id;

  String title;

  String originalTitle;

  String posterLink;

  ShowType type;

  int releaseYear;

  num score;

  int durationInMinutes;

  String quantity;

  int totalEpisodes;

  int currentEpisode;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ModelShow &&
    other.id == id &&
    other.title == title &&
    other.originalTitle == originalTitle &&
    other.posterLink == posterLink &&
    other.type == type &&
    other.releaseYear == releaseYear &&
    other.score == score &&
    other.durationInMinutes == durationInMinutes &&
    other.quantity == quantity &&
    other.totalEpisodes == totalEpisodes &&
    other.currentEpisode == currentEpisode;

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (id.hashCode) +
    (title.hashCode) +
    (originalTitle.hashCode) +
    (posterLink.hashCode) +
    (type.hashCode) +
    (releaseYear.hashCode) +
    (score.hashCode) +
    (durationInMinutes.hashCode) +
    (quantity.hashCode) +
    (totalEpisodes.hashCode) +
    (currentEpisode.hashCode);

  @override
  String toString() => 'ModelShow[id=$id, title=$title, originalTitle=$originalTitle, posterLink=$posterLink, type=$type, releaseYear=$releaseYear, score=$score, durationInMinutes=$durationInMinutes, quantity=$quantity, totalEpisodes=$totalEpisodes, currentEpisode=$currentEpisode]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'id'] = this.id;
      json[r'title'] = this.title;
      json[r'originalTitle'] = this.originalTitle;
      json[r'posterLink'] = this.posterLink;
      json[r'type'] = this.type;
      json[r'releaseYear'] = this.releaseYear;
      json[r'score'] = this.score;
      json[r'durationInMinutes'] = this.durationInMinutes;
      json[r'quantity'] = this.quantity;
      json[r'totalEpisodes'] = this.totalEpisodes;
      json[r'currentEpisode'] = this.currentEpisode;
    return json;
  }

  /// Returns a new [ModelShow] instance and imports its values from
  /// [value] if it's a [Map], null otherwise.
  // ignore: prefer_constructors_over_static_methods
  static ModelShow? fromJson(dynamic value) {
    if (value is Map) {
      final json = value.cast<String, dynamic>();

      // Ensure that the map contains the required keys.
      // Note 1: the values aren't checked for validity beyond being non-null.
      // Note 2: this code is stripped in release mode!
      assert(() {
        requiredKeys.forEach((key) {
          assert(json.containsKey(key), 'Required key "ModelShow[$key]" is missing from JSON.');
          assert(json[key] != null, 'Required key "ModelShow[$key]" has a null value in JSON.');
        });
        return true;
      }());

      return ModelShow(
        id: mapValueOfType<int>(json, r'id')!,
        title: mapValueOfType<String>(json, r'title')!,
        originalTitle: mapValueOfType<String>(json, r'originalTitle')!,
        posterLink: mapValueOfType<String>(json, r'posterLink')!,
        type: ShowType.fromJson(json[r'type'])!,
        releaseYear: mapValueOfType<int>(json, r'releaseYear')!,
        score: num.parse('${json[r'score']}'),
        durationInMinutes: mapValueOfType<int>(json, r'durationInMinutes')!,
        quantity: mapValueOfType<String>(json, r'quantity')!,
        totalEpisodes: mapValueOfType<int>(json, r'totalEpisodes')!,
        currentEpisode: mapValueOfType<int>(json, r'currentEpisode')!,
      );
    }
    return null;
  }

  static List<ModelShow> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <ModelShow>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = ModelShow.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }

  static Map<String, ModelShow> mapFromJson(dynamic json) {
    final map = <String, ModelShow>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = ModelShow.fromJson(entry.value);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  // maps a json object with a list of ModelShow-objects as value to a dart map
  static Map<String, List<ModelShow>> mapListFromJson(dynamic json, {bool growable = false,}) {
    final map = <String, List<ModelShow>>{};
    if (json is Map && json.isNotEmpty) {
      // ignore: parameter_assignments
      json = json.cast<String, dynamic>();
      for (final entry in json.entries) {
        map[entry.key] = ModelShow.listFromJson(entry.value, growable: growable,);
      }
    }
    return map;
  }

  /// The list of required keys that must be present in a JSON.
  static const requiredKeys = <String>{
    'id',
    'title',
    'originalTitle',
    'posterLink',
    'type',
    'releaseYear',
    'score',
    'durationInMinutes',
    'quantity',
    'totalEpisodes',
    'currentEpisode',
  };
}

