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

  Future<AuthUser> createUser(
          {required String email, required String password}) =>
      authProvider.createUser(email: email, password: password);

  AuthUser? get currentUser => authProvider.currentUser;
  Future<void> initialize() => authProvider.initialize();

  Future<AuthUser> logIn({required String email, required String password}) =>
      authProvider.logIn(email: email, password: password);

  Future<void> logOut() => authProvider.logOut();

  Future<void> sendEmailVerification() => authProvider.sendEmailVerification();

  Future<void> sendPasswordReset({required String toEmail}) =>
      authProvider.sendPasswordReset(toEmail: toEmail);

  Future<String?> get authToken => authProvider.authToken;

  bool get isVerifiedUser => authProvider.isVerifiedUser;
}
