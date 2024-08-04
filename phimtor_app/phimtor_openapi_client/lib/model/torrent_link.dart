//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.18

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class TorrentLink {
  /// Returns a new [TorrentLink] instance.
  TorrentLink({
    required this.id,
    required this.name,
    required this.link,
    required this.fileIndex,
    required this.priority,
  });

  int id;

  String name;

  String link;

  int fileIndex;

  int priority;

  @override
  bool operator ==(Object other) => identical(this, other) || other is TorrentLink &&
    other.id == id &&
    other.name == name &&
    other.link == link &&
    other.fileIndex == fileIndex &&
    other.priority == priority;

  @override
  int get hashCode =>
    // ignore: unnecessary_parenthesis
    (id.hashCode) +
    (name.hashCode) +
    (link.hashCode) +
    (fileIndex.hashCode) +
    (priority.hashCode);

  @override
  String toString() => 'TorrentLink[id=$id, name=$name, link=$link, fileIndex=$fileIndex, priority=$priority]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
      json[r'id'] = this.id;
      json[r'name'] = this.name;
      json[r'link'] = this.link;
      json[r'fileIndex'] = this.fileIndex;
      json[r'priority'] = this.priority;
    return json;
  }

  /// Returns a new [TorrentLink] instance and imports its values from
  /// [value] if it's a [Map], null otherwise.
  // ignore: prefer_constructors_over_static_methods
  static TorrentLink? fromJson(dynamic value) {
    if (value is Map) {
      final json = value.cast<String, dynamic>();

      // Ensure that the map contains the required keys.
      // Note 1: the values aren't checked for validity beyond being non-null.
      // Note 2: this code is stripped in release mode!
      assert(() {
        requiredKeys.forEach((key) {
          assert(json.containsKey(key), 'Required key "TorrentLink[$key]" is missing from JSON.');
          assert(json[key] != null, 'Required key "TorrentLink[$key]" has a null value in JSON.');
        });
        return true;
      }());

      return TorrentLink(
        id: mapValueOfType<int>(json, r'id')!,
        name: mapValueOfType<String>(json, r'name')!,
        link: mapValueOfType<String>(json, r'link')!,
        fileIndex: mapValueOfType<int>(json, r'fileIndex')!,
        priority: mapValueOfType<int>(json, r'priority')!,
      );
    }
    return null;
  }

  static List<TorrentLink> listFromJson(dynamic json, {bool growable = false,}) {
    final result = <TorrentLink>[];
    if (json is List && json.isNotEmpty) {
      for (final row in json) {
        final value = TorrentLink.fromJson(row);
        if (value != null) {
          result.add(value);
        }
      }
    }
    return result.toList(growable: growable);
  }

  static Map<String, TorrentLink> mapFromJson(dynamic json) {
    final map = <String, TorrentLink>{};
    if (json is Map && json.isNotEmpty) {
      json = json.cast<String, dynamic>(); // ignore: parameter_assignments
      for (final entry in json.entries) {
        final value = TorrentLink.fromJson(entry.value);
        if (value != null) {
          map[entry.key] = value;
        }
      }
    }
    return map;
  }

  // maps a json object with a list of TorrentLink-objects as value to a dart map
  static Map<String, List<TorrentLink>> mapListFromJson(dynamic json, {bool growable = false,}) {
    final map = <String, List<TorrentLink>>{};
    if (json is Map && json.isNotEmpty) {
      // ignore: parameter_assignments
      json = json.cast<String, dynamic>();
      for (final entry in json.entries) {
        map[entry.key] = TorrentLink.listFromJson(entry.value, growable: growable,);
      }
    }
    return map;
  }

  /// The list of required keys that must be present in a JSON.
  static const requiredKeys = <String>{
    'id',
    'name',
    'link',
    'fileIndex',
    'priority',
  };
}

