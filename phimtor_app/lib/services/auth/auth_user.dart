import 'package:firedart/auth/user_gateway.dart' as firedart_auth show User;
import 'package:flutter/foundation.dart';
import 'package:firebase_auth/firebase_auth.dart' as firebase_auth show User;

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

  factory AuthUser.fromFirebaseUser(firebase_auth.User user) {
    return AuthUser(
      uid: user.uid,
      email: user.email!,
      displayName: user.displayName,
      emailVerified: user.emailVerified,
    );
  }
  
  factory AuthUser.fromFiredartUser(firedart_auth.User user) {
    return AuthUser(
      uid: user.id,
      email: user.email ?? '',
      displayName: user.displayName,
      emailVerified: user.emailVerified ?? false,
    );
  }
}
