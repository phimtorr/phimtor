import 'package:firedart/auth/user_gateway.dart' as firedart_auth show User;
import 'package:flutter/foundation.dart';
import 'package:firebase_auth/firebase_auth.dart' as firebase_auth show User;
import 'package:jwt_decoder/jwt_decoder.dart';

@immutable
class AuthUser {
  // constants
  static const String premiumUntilKey = 'premium_until';

  final String uid;
  final String email;
  final String? displayName;
  final bool emailVerified;
  final DateTime? premiumUntil;

  const AuthUser({
    required this.uid,
    required this.email,
    this.displayName,
    required this.emailVerified,
    this.premiumUntil,
  });

  factory AuthUser.fromFirebaseUser(firebase_auth.User user, Map<String, dynamic>? claims) {
    DateTime? premiumUntil;
    if (claims != null) {
      final premiumUntilEpoch = claims[premiumUntilKey] as int?;
      if (premiumUntilEpoch != null) {
        premiumUntil = DateTime.fromMillisecondsSinceEpoch(premiumUntilEpoch * 1000);
      }
    }
    return AuthUser(
      uid: user.uid,
      email: user.email!,
      displayName: user.displayName,
      emailVerified: user.emailVerified,
      premiumUntil: premiumUntil,
    );
  }
  
  factory AuthUser.fromFiredartUser(firedart_auth.User user, String? jwtToken) {
    DateTime? premiumUntil;
    if (jwtToken != null) {
      final decodedToken = JwtDecoder.decode(jwtToken);
      final premiumUntilEpoch = decodedToken[premiumUntilKey] as int?;
      if (premiumUntilEpoch != null) {
        premiumUntil = DateTime.fromMillisecondsSinceEpoch(premiumUntilEpoch * 1000);
      }
    }

    return AuthUser(
      uid: user.id,
      email: user.email ?? '',
      displayName: user.displayName,
      emailVerified: user.emailVerified ?? false,
      premiumUntil: premiumUntil,
    );
  }
}
