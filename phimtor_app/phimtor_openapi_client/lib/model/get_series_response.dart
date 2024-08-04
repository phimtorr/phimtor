//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class GetSeriesResponse {
  /// Returns a new [GetSeriesResponse] instance.
  GetSeriesResponse({
    required this.series,
  });

  Series series;

  @override
  bool operator ==(Object other) => identical(this, other) || other is GetSeriesResponse &&
    other.series == series;

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (series.hashCode);

  @override
  String toString() => 'GetSeriesResponse[series=$series]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'series'] = this.series;
    return json;
  }

  /// Returns a new [GetSeriesResponse] instance and imports its values from
  /// [value] if it's a [Map], null otherwise.
  // ignore: prefer_constructors_over_static_methods
  static GetSeriesResponse? fromJson(dynamic value) {
    if (value is Map) {
      final json = value.cast<String, dynamic>();

      // Ensure that the map contains the required keys.
      // Note 1: the values aren't checked for validity beyond being non-null.
      // Note 2: this code is stripped in release mode!
      assert(() {
        requiredKeys.forEach((key) {
          assert(json.containsKey(key), 'Required key "GetSeriesResponse[$key]" is missing from JSON.');
          assert(json[key] != null, 'Required key "GetSeriesResponse[$key]" has a null value in JSON.');
        });
        return true;
      }());

      return GetSeriesResponse(
        series: Series.fromJson(json[r'series'])!,
      );
    }
    return null;
  }

  static List<GetSeriesResponse> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <GetSeriesResponse>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = GetSeriesResponse.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }

  static Map<String, GetSeriesResponse> mapFromJson(dynamic json) {
    final map = <String, GetSeriesResponse>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = GetSeriesResponse.fromJson(entry.value);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  // maps a json object with a list of GetSeriesResponse-objects as value to a dart map
  static Map<String, List<GetSeriesResponse>> mapListFromJson(dynamic json, {bool growable = false,}) {
    final map = <String, List<GetSeriesResponse>>{};
    if (json is Map && json.isNotEmpty) {
      // ignore: parameter_assignments
      json = json.cast<String, dynamic>();
      for (final entry in json.entries) {
        map[entry.key] = GetSeriesResponse.listFromJson(entry.value, growable: growable,);
      }
    }
    return map;
  }

  /// The list of required keys that must be present in a JSON.
  static const requiredKeys = <String>{
    'series',
  };
}

