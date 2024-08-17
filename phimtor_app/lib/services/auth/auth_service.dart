import 'package:firebase_core/firebase_core.dart';
import 'package:phimtor_app/firebase_options.dart';

class AuthService {
  // singleton
  static final AuthService _instance = AuthService._internal();
  factory AuthService() => _instance;
  AuthService._internal();

  bool _isInitialized = false;

  Future<void> ensureInitialized() async {
    if (_isInitialized) {
      return;
    }
    await Firebase.initializeApp(
      options: DefaultFirebaseOptions.currentPlatform,
    );
    await Firebase.app().setAutomaticDataCollectionEnabled(true);
    _isInitialized = true;
  }
}