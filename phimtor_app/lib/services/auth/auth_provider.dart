import 'package:phimtor_app/services/auth/auth_user.dart';

abstract class AuthProvider {
  Future<void> initialize();
  Future<AuthUser?> get currentUser;
  Future<AuthUser> logIn({
    required String email,
    required String password,
  });
  Future<AuthUser> createUser({
    required String email,
    required String password,
  });
  Future<void> logOut();
  Future<void> sendEmailVerification();
  Future<void> sendPasswordReset({required String toEmail});
  Future<String?> get authToken;
}
