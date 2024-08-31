import 'package:flutter/material.dart';

@immutable
class UpdaterVersion {
  final String version;
  final Uri binaryUrl;

  const UpdaterVersion({
    required this.version,
    required this.binaryUrl,
  });

  @override
  bool operator ==(Object other) {
    if (identical(this, other)) return true;

    return other is UpdaterVersion &&
        other.version.toString().trim() == version.toLowerCase().trim();
  }
  
  @override
  int get hashCode => version.hashCode;
  
}


