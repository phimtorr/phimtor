//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class GetTvSeriesResponse {
  /// Returns a new [GetTvSeriesResponse] instance.
  GetTvSeriesResponse({
    required this.tvSeries,
  });

  TvSeries tvSeries;

  @override
  bool operator ==(Object other) => identical(this, other) || other is GetTvSeriesResponse &&
    other.tvSeries == tvSeries;

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (tvSeries.hashCode);

  @override
  String toString() => 'GetTvSeriesResponse[tvSeries=$tvSeries]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'tvSeries'] = this.tvSeries;
    return json;
  }

  /// Returns a new [GetTvSeriesResponse] instance and imports its values from
  /// [value] if it's a [Map], null otherwise.
  // ignore: prefer_constructors_over_static_methods
  static GetTvSeriesResponse? fromJson(dynamic value) {
    if (value is Map) {
      final json = value.cast<String, dynamic>();

      // Ensure that the map contains the required keys.
      // Note 1: the values aren't checked for validity beyond being non-null.
      // Note 2: this code is stripped in release mode!
      assert(() {
        requiredKeys.forEach((key) {
          assert(json.containsKey(key), 'Required key "GetTvSeriesResponse[$key]" is missing from JSON.');
          assert(json[key] != null, 'Required key "GetTvSeriesResponse[$key]" has a null value in JSON.');
        });
        return true;
      }());

      return GetTvSeriesResponse(
        tvSeries: TvSeries.fromJson(json[r'tvSeries'])!,
      );
    }
    return null;
  }

  static List<GetTvSeriesResponse> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <GetTvSeriesResponse>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = GetTvSeriesResponse.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }

  static Map<String, GetTvSeriesResponse> mapFromJson(dynamic json) {
    final map = <String, GetTvSeriesResponse>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = GetTvSeriesResponse.fromJson(entry.value);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  // maps a json object with a list of GetTvSeriesResponse-objects as value to a dart map
  static Map<String, List<GetTvSeriesResponse>> mapListFromJson(dynamic json, {bool growable = false,}) {
    final map = <String, List<GetTvSeriesResponse>>{};
    if (json is Map && json.isNotEmpty) {
      // ignore: parameter_assignments
      json = json.cast<String, dynamic>();
      for (final entry in json.entries) {
        map[entry.key] = GetTvSeriesResponse.listFromJson(entry.value, growable: growable,);
      }
    }
    return map;
  }

  /// The list of required keys that must be present in a JSON.
  static const requiredKeys = <String>{
    'tvSeries',
  };
}

