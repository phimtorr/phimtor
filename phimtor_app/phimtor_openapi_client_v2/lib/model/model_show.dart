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
    required this.showId,
    required this.type,
    required this.title,
    required this.originalTitle,
    required this.posterLink,
    required this.airDate,
    required this.runtime,
    required this.voteAverage,
    required this.quality,
    required this.seasonNumber,
    required this.episodeNumber,
  });

  int id;

  int showId;

  ModelShowTypeEnum type;

  String title;

  String originalTitle;

  String posterLink;

  DateTime airDate;

  int runtime;

  num voteAverage;

  String quality;

  int seasonNumber;

  int episodeNumber;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ModelShow &&
    other.id == id &&
    other.showId == showId &&
    other.type == type &&
    other.title == title &&
    other.originalTitle == originalTitle &&
    other.posterLink == posterLink &&
    other.airDate == airDate &&
    other.runtime == runtime &&
    other.voteAverage == voteAverage &&
    other.quality == quality &&
    other.seasonNumber == seasonNumber &&
    other.episodeNumber == episodeNumber;

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (id.hashCode) +
    (showId.hashCode) +
    (type.hashCode) +
    (title.hashCode) +
    (originalTitle.hashCode) +
    (posterLink.hashCode) +
    (airDate.hashCode) +
    (runtime.hashCode) +
    (voteAverage.hashCode) +
    (quality.hashCode) +
    (seasonNumber.hashCode) +
    (episodeNumber.hashCode);

  @override
  String toString() => 'ModelShow[id=$id, showId=$showId, type=$type, title=$title, originalTitle=$originalTitle, posterLink=$posterLink, airDate=$airDate, runtime=$runtime, voteAverage=$voteAverage, quality=$quality, seasonNumber=$seasonNumber, episodeNumber=$episodeNumber]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'id'] = this.id;
      json[r'showId'] = this.showId;
      json[r'type'] = this.type;
      json[r'title'] = this.title;
      json[r'originalTitle'] = this.originalTitle;
      json[r'posterLink'] = this.posterLink;
      json[r'airDate'] = _dateFormatter.format(this.airDate.toUtc());
      json[r'runtime'] = this.runtime;
      json[r'voteAverage'] = this.voteAverage;
      json[r'quality'] = this.quality;
      json[r'seasonNumber'] = this.seasonNumber;
      json[r'episodeNumber'] = this.episodeNumber;
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
        showId: mapValueOfType<int>(json, r'showId')!,
        type: ModelShowTypeEnum.fromJson(json[r'type'])!,
        title: mapValueOfType<String>(json, r'title')!,
        originalTitle: mapValueOfType<String>(json, r'originalTitle')!,
        posterLink: mapValueOfType<String>(json, r'posterLink')!,
        airDate: mapDateTime(json, r'airDate', r'')!,
        runtime: mapValueOfType<int>(json, r'runtime')!,
        voteAverage: num.parse('${json[r'voteAverage']}'),
        quality: mapValueOfType<String>(json, r'quality')!,
        seasonNumber: mapValueOfType<int>(json, r'seasonNumber')!,
        episodeNumber: mapValueOfType<int>(json, r'episodeNumber')!,
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
    'showId',
    'type',
    'title',
    'originalTitle',
    'posterLink',
    'airDate',
    'runtime',
    'voteAverage',
    'quality',
    'seasonNumber',
    'episodeNumber',
  };
}


class ModelShowTypeEnum {
  /// Instantiate a new enum with the provided [value].
  const ModelShowTypeEnum._(this.value);

  /// The underlying value of this enum member.
  final String value;

  @override
  String toString() => value;

  String toJson() => value;

  static const movie = ModelShowTypeEnum._(r'movie');
  static const tvSeries = ModelShowTypeEnum._(r'tv-series');
  static const episode = ModelShowTypeEnum._(r'episode');

  /// List of all possible values in this [enum][ModelShowTypeEnum].
  static const values = <ModelShowTypeEnum>[
    movie,
    tvSeries,
    episode,
  ];

  static ModelShowTypeEnum? fromJson(dynamic value) => ModelShowTypeEnumTypeTransformer().decode(value);

  static List<ModelShowTypeEnum> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <ModelShowTypeEnum>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = ModelShowTypeEnum.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }
}

/// Transformation class that can [encode] an instance of [ModelShowTypeEnum] to String,
/// and [decode] dynamic data back to [ModelShowTypeEnum].
class ModelShowTypeEnumTypeTransformer {
  factory ModelShowTypeEnumTypeTransformer() => _instance ??= const ModelShowTypeEnumTypeTransformer._();

  const ModelShowTypeEnumTypeTransformer._();

  String encode(ModelShowTypeEnum data) => data.value;

  /// Decodes a [dynamic value][data] to a ModelShowTypeEnum.
  ///
  /// If [allowNull] is true and the [dynamic value][data] cannot be decoded successfully,
  /// then null is returned. However, if [allowNull] is false and the [dynamic value][data]
  /// cannot be decoded successfully, then an [UnimplementedError] is thrown.
  ///
  /// The [allowNull] is very handy when an API changes and a new enum value is added or removed,
  /// and users are still using an old app with the old code.
  ModelShowTypeEnum? decode(dynamic data, {bool allowNull = true}) {
    if (data != null) {
      switch (data) {
        case r'movie': return ModelShowTypeEnum.movie;
        case r'tv-series': return ModelShowTypeEnum.tvSeries;
        case r'episode': return ModelShowTypeEnum.episode;
        default:
          if (!allowNull) {
            throw ArgumentError('Unknown enum value to decode: $data');
          }
      }
    }
    return null;
  }

  /// Singleton [ModelShowTypeEnumTypeTransformer] instance.
  static ModelShowTypeEnumTypeTransformer? _instance;
}


