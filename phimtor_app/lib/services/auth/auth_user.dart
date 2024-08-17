import 'package:flutter/foundation.dart';
import 'package:firebase_auth/firebase_auth.dart' show User;

@immutable
class AuthUser {
  final String uid;
  final String email;
  final String? displayName;
  final bool emailVerified;

  const AuthUser({
    required this.uid,
    required this.email,
    this.displayName,
    required this.emailVerified,
  });

  factory AuthUser.fromFireabaseUser(User user) {
    return AuthUser(
      uid: user.uid,
      email: user.email!,
      displayName: user.displayName,
      emailVerified: user.emailVerified,
    );
  }
}
