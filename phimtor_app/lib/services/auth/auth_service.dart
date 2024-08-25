import 'package:phimtor_app/services/auth/auth_provider.dart';
import 'package:phimtor_app/services/auth/auth_user.dart';
import 'package:flutter/foundation.dart'
    show defaultTargetPlatform, TargetPlatform;
import 'package:phimtor_app/services/auth/firebase_auth_provider.dart';
import 'package:phimtor_app/services/auth/firedart_auth_provider.dart';

class AuthService implements AuthProvider {
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

  @override
  Future<AuthUser> createUser(
          {required String email, required String password}) =>
      authProvider.createUser(email: email, password: password);

  @override
  AuthUser? get currentUser => authProvider.currentUser;
  @override
  Future<void> initialize() => authProvider.initialize();

  @override
  Future<AuthUser> logIn({required String email, required String password}) =>
      authProvider.logIn(email: email, password: password);

  @override
  Future<void> logOut() => authProvider.logOut();

  @override
  Future<void> sendEmailVerification() => authProvider.sendEmailVerification();

  @override
  Future<void> sendPasswordReset({required String toEmail}) =>
      authProvider.sendPasswordReset(toEmail: toEmail);

  @override
  Future<String?> get authToken => authProvider.authToken;

  @override
  bool get isVerifiedUser => authProvider.isVerifiedUser;
}
