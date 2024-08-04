//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;


class ShowType {
  /// Instantiate a new enum with the provided [value].
  const ShowType._(this.value);

  /// The underlying value of this enum member.
  final String value;

  @override
  String toString() => value;

  String toJson() => value;

  static const movie = ShowType._(r'movie');
  static const series = ShowType._(r'series');

  /// List of all possible values in this [enum][ShowType].
  static const values = <ShowType>[
    movie,
    series,
  ];

  static ShowType? fromJson(dynamic value) => ShowTypeTypeTransformer().decode(value);

  static List<ShowType> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <ShowType>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = ShowType.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }
}

/// Transformation class that can [encode] an instance of [ShowType] to String,
/// and [decode] dynamic data back to [ShowType].
class ShowTypeTypeTransformer {
  factory ShowTypeTypeTransformer() => _instance ??= const ShowTypeTypeTransformer._();

  const ShowTypeTypeTransformer._();

  String encode(ShowType data) => data.value;

  /// Decodes a [dynamic value][data] to a ShowType.
  ///
  /// If [allowNull] is true and the [dynamic value][data] cannot be decoded successfully,
  /// then null is returned. However, if [allowNull] is false and the [dynamic value][data]
  /// cannot be decoded successfully, then an [UnimplementedError] is thrown.
  ///
  /// The [allowNull] is very handy when an API changes and a new enum value is added or removed,
  /// and users are still using an old app with the old code.
  ShowType? decode(dynamic data, {bool allowNull = true}) {
    if (data != null) {
      switch (data) {
        case r'movie': return ShowType.movie;
        case r'series': return ShowType.series;
        default:
          if (!allowNull) {
            throw ArgumentError('Unknown enum value to decode: $data');
          }
      }
    }
    return null;
  }

  /// Singleton [ShowTypeTypeTransformer] instance.
  static ShowTypeTypeTransformer? _instance;
}

