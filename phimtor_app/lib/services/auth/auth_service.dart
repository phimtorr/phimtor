import 'package:phimtor_app/services/auth/auth_provider.dart';
import 'package:phimtor_app/services/auth/auth_user.dart';
import 'package:flutter/foundation.dart'
    show defaultTargetPlatform, TargetPlatform;
import 'package:phimtor_app/services/auth/firebase_auth_provider.dart';
import 'package:phimtor_app/services/auth/firedart_auth_provider.dart';

class AuthService {
  // singleton
  static final AuthService _instance = AuthService._internal();

  factory AuthService() {
    return _instance;
  }

  AuthService._internal()
      : authProvider = defaultTargetPlatform == TargetPlatform.linux
            ? FiredartAuthProvider()
            : FirebaseAuthProvider();

  final AuthProvider authProvider;
  AuthUser? _currentUser;

  Future<AuthUser?> createUser({
    required String email,
    required String password,
  }) async {
    await authProvider.createUser(email: email, password: password);
    await _syncCurrentUser();
    return _currentUser;
  }

  AuthUser? get currentUser => _currentUser;

  Future<void> initialize() => authProvider.initialize();

  Future<AuthUser> logIn({
    required String email,
    required String password,
  }) async {
    await authProvider.logIn(email: email, password: password);
    await _syncCurrentUser();
    return _currentUser!;
  }

  Future<void> logOut() async {
    await authProvider.logOut();
    await _syncCurrentUser();
  }

  Future<void> sendEmailVerification() => authProvider.sendEmailVerification();

  Future<void> sendPasswordReset({required String toEmail}) =>
      authProvider.sendPasswordReset(toEmail: toEmail);

  Future<String?> get authToken => authProvider.authToken;

  bool get isVerifiedUser => _currentUser?.emailVerified ?? false;

  Future<void> _syncCurrentUser() async {
    _currentUser = await authProvider.syncCurrentUser();
  }

  bool get isPremiumUser {
    if (_currentUser == null) {
      return false;
    }

    if (_currentUser!.premiumUntil == null) {
      return false;
    }

    return DateTime.now().isBefore(_currentUser!.premiumUntil!);
  }
}
