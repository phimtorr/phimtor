import 'dart:convert';

import 'package:firedart/auth/exceptions.dart';
import 'package:flutter/material.dart';
import 'package:phimtor_app/firebase_options.dart';
import 'package:phimtor_app/services/auth/auth_provider.dart';
import 'package:phimtor_app/services/auth/auth_user.dart';
import 'package:firedart/firedart.dart';

import 'package:shared_preferences/shared_preferences.dart';

class FiredartAuthProvider implements AuthProvider {
  late final TokenStore _tokenStore;

  @override
  Future<void> initialize() async {
    _tokenStore = await PreferencesStore.create();
    FirebaseAuth.initialize(
      DefaultFirebaseOptions.windows.apiKey,
      _tokenStore,
    );
    try {
      await syncCurrentUser();
    } catch (e) {
      debugPrint(e.toString());
    }
  }

  @override
  Future<void> createUser({
    required String email,
    required String password,
  }) async {
    await FirebaseAuth.instance.signUp(email, password);
    return;
  }

  @override
  Future<AuthUser?> syncCurrentUser() async {
    try {
      final user = await FirebaseAuth.instance.getUser();
      final token = _tokenStore.idToken;
      return AuthUser.fromFiredartUser(user, token);
    } on SignedOutException {
      return null;
    } catch (e) {
      debugPrint(e.toString());
      rethrow;
    }
  }

  @override
  Future<void> logIn({
    required String email,
    required String password,
  }) async {
    await FirebaseAuth.instance.signIn(email, password);
    return;
  }

  @override
  Future<void> logOut() async {
    FirebaseAuth.instance.signOut();
    _tokenStore.delete();
  }

  @override
  Future<void> sendEmailVerification() async {
    await FirebaseAuth.instance.requestEmailVerification();
  }

  @override
  Future<void> sendPasswordReset({required String toEmail}) async {
    await FirebaseAuth.instance.resetPassword(toEmail);
  }

  @override
  Future<String?> get authToken async {
    return _tokenStore.idToken;
  }
}

/// Stores tokens as preferences in Android and iOS.
/// Depends on the shared_preferences plugin: https://pub.dev/packages/shared_preferences
class PreferencesStore extends TokenStore {
  static const keyToken = "auth_token";

  static Future<PreferencesStore> create() async =>
      PreferencesStore._internal(await SharedPreferences.getInstance());

  final SharedPreferences _prefs;

  PreferencesStore._internal(this._prefs);

  @override
  Token? read() => _prefs.containsKey(keyToken)
      ? Token.fromMap(json.decode(_prefs.get(keyToken) as String))
      : null;

  @override
  void write(Token? token) => token != null
      ? _prefs.setString(keyToken, json.encode(token.toMap()))
      : null;

  @override
  void delete() => _prefs.remove(keyToken);
}
