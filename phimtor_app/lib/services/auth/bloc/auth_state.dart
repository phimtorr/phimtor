import 'package:flutter/foundation.dart' show immutable;
import 'package:equatable/equatable.dart';
import 'package:phimtor_app/services/auth/auth_user.dart';

@immutable
abstract class AuthState {
  final bool isLoading;
  final String? loadingText;

  const AuthState({
    required this.isLoading,
    this.loadingText = "Loading...",
  });
}

class AuthStateLoggedOut extends AuthState with EquatableMixin {
  final Exception? exception;

  const AuthStateLoggedOut({
    required this.exception,
    required super.isLoading,
    super.loadingText,
  });

  @override
  List<Object?> get props => [exception, isLoading];
}

class AuthStateLoggedIn extends AuthState {
  final AuthUser user;

  const AuthStateLoggedIn({
    required this.user,
    required super.isLoading,
  });
}

class AuthStateRegistering extends AuthState {
  final Exception? exception;
  const AuthStateRegistering({
    required this.exception,
    required super.isLoading,
  });
}

class AuthStateNeedsVerification extends AuthState {
  final bool needCooldown;
  const AuthStateNeedsVerification({
    required super.isLoading,
    this.needCooldown = false,
  });
}

class AuthStateForgotPassword extends AuthState {
  final Exception? exception;
  final bool emailSent;
  const AuthStateForgotPassword({
    required this.exception,
    required super.isLoading,
    required this.emailSent,
  });
}
