import 'package:firebase_auth/firebase_auth.dart' show FirebaseAuth;
import 'package:firebase_core/firebase_core.dart';
import 'package:phimtor_app/firebase_options.dart';
import 'package:phimtor_app/services/auth/auth_provider.dart';
import 'package:phimtor_app/services/auth/auth_user.dart';

class FirebaseAuthProvider implements AuthProvider {
  @override
  Future<void> initialize() async {
    await Firebase.initializeApp(
      options: DefaultFirebaseOptions.currentPlatform,
    );
  }

  AuthUser? get _currentUser {
    final user = FirebaseAuth.instance.currentUser;
    return user == null ? null : AuthUser.fromFirebaseUser(user);
  }

  @override
  Future<AuthUser> createUser({
    required String email,
    required String password,
  }) async {
    try {
      await FirebaseAuth.instance.createUserWithEmailAndPassword(
        email: email,
        password: password,
      );

      final user = _currentUser;
      if (user == null) {
        throw Exception('User not found after creation');
      }
      return user;
    } catch (e) {
      rethrow;
    }
  }

  @override
  Future<AuthUser?> get currentUser async {
    return _currentUser;
  }

  @override
  Future<AuthUser> logIn({
    required String email,
    required String password,
  }) async {
    try {
      await FirebaseAuth.instance.signInWithEmailAndPassword(
        email: email,
        password: password,
      );

      final user = _currentUser;
      if (user == null) {
        throw Exception('User not found after login');
      }
      return user;
    } catch (e) {
      rethrow;
    }
  }

  @override
  Future<void> logOut() async {
    await FirebaseAuth.instance.signOut();
  }

  @override
  Future<void> sendEmailVerification() async {
    final user = FirebaseAuth.instance.currentUser;
    if (user == null) {
      throw Exception('User not found');
    }
    await user.sendEmailVerification();
  }

  @override
  Future<void> sendPasswordReset({required String toEmail}) async {
    await FirebaseAuth.instance.sendPasswordResetEmail(email: toEmail);
  }

  @override
  Future<String?> get authToken async {
    final user = FirebaseAuth.instance.currentUser;
    if (user == null) {
      return null;
    }
    return user.getIdToken();
  }
}
