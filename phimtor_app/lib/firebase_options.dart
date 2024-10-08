// File generated by FlutterFire CLI.
// ignore_for_file: type=lint
import 'package:firebase_core/firebase_core.dart' show FirebaseOptions;
import 'package:flutter/foundation.dart'
    show defaultTargetPlatform, kIsWeb, TargetPlatform;

/// Default [FirebaseOptions] for use with your Firebase apps.
///
/// Example:
/// ```dart
/// import 'firebase_options.dart';
/// // ...
/// await Firebase.initializeApp(
///   options: DefaultFirebaseOptions.currentPlatform,
/// );
/// ```
class DefaultFirebaseOptions {
  static FirebaseOptions get currentPlatform {
    if (kIsWeb) {
      return web;
    }
    switch (defaultTargetPlatform) {
      case TargetPlatform.android:
        return android;
      case TargetPlatform.iOS:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for ios - '
          'you can reconfigure this by running the FlutterFire CLI again.',
        );
      case TargetPlatform.macOS:
        return macos;
      case TargetPlatform.windows:
        return windows;
      case TargetPlatform.linux:
        return linux;
      default:
        throw UnsupportedError(
          'DefaultFirebaseOptions are not supported for this platform.',
        );
    }
  }

  static const FirebaseOptions android = FirebaseOptions(
    apiKey: 'AIzaSyA3_2E0gTvcCA6PYhSo3xlPWWI3tmL_KZ8',
    appId: '1:744867104175:android:271608753e2d8b594b1b5a',
    messagingSenderId: '744867104175',
    projectId: 'phimtor-d67b3',
    storageBucket: 'phimtor-d67b3.appspot.com',
  );

  static const FirebaseOptions macos = FirebaseOptions(
    apiKey: 'AIzaSyBKpApsYBPwcZZuJ8xeGJGXx0ylEu5dTGA',
    appId: '1:744867104175:ios:8b09eadd274b4b2b4b1b5a',
    messagingSenderId: '744867104175',
    projectId: 'phimtor-d67b3',
    storageBucket: 'phimtor-d67b3.appspot.com',
    iosBundleId: 'net.phimtor.phimtorapp',
  );

  static const FirebaseOptions windows = FirebaseOptions(
    apiKey: 'AIzaSyCSbDOQE6soH-DlOqD2ovrVUWjlzzg6pSA',
    appId: '1:744867104175:web:4d220c3f6a77442c4b1b5a',
    messagingSenderId: '744867104175',
    projectId: 'phimtor-d67b3',
    authDomain: 'phimtor-d67b3.firebaseapp.com',
    storageBucket: 'phimtor-d67b3.appspot.com',
    measurementId: 'G-N2M2ZZ2LJG',
  );

  static const linux = windows;

  static const FirebaseOptions web = FirebaseOptions(
    apiKey: 'AIzaSyCSbDOQE6soH-DlOqD2ovrVUWjlzzg6pSA',
    appId: '1:744867104175:web:339a4872024684b94b1b5a',
    messagingSenderId: '744867104175',
    projectId: 'phimtor-d67b3',
    authDomain: 'phimtor-d67b3.firebaseapp.com',
    storageBucket: 'phimtor-d67b3.appspot.com',
    measurementId: 'G-FNPK6JWFMJ',
  );

}