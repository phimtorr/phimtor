import 'package:flutter/material.dart';

@immutable
class UpdaterVersion {
  final String version;
  final String binaryUrl;

  const UpdaterVersion({
    required this.version,
    required this.binaryUrl,
  });

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;

    return other is UpdaterVersion &&
        other.version == version;
  }
  
  @override
  int get hashCode => version.hashCode;
  
}


