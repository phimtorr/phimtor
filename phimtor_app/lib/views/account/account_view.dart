import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/auth/auth_service.dart';
import 'package:phimtor_app/services/auth/bloc/auth_bloc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_state.dart';
import 'package:phimtor_app/views/account/forgot_password_view.dart';
import 'package:phimtor_app/views/account/logged_in_view.dart';
import 'package:phimtor_app/views/account/login_view.dart';
import 'package:phimtor_app/views/account/register_view.dart';
import 'package:phimtor_app/views/account/verify_email_view.dart';

class AccountView extends StatelessWidget {
  const AccountView({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider<AuthBloc>(
      create: (context) => AuthBloc(AuthService()),
      child: Builder(builder: (context) {
        return BlocConsumer<AuthBloc, AuthState>(
          listener: (context, state) {
            if (state is AuthStateLoggedOut) {
              if (state.exception != null) {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(
                    content:
                        Text(context.loc.error(state.exception.toString())),
                  ),
                );
              }
            }
          },
          builder: (context, state) {
            if (state.isLoading) {
              return const Scaffold(
                body: Center(child: CircularProgressIndicator()),
              );
            }

            if (state is AuthStateLoggedOut) {
              return const LoginView();
            }

            if (state is AuthStateRegistering) {
              return const RegisterView();
            }

            if (state is AuthStateNeedsVerification) {
              return VerifyEmailView(
                needCooldown: state.needCooldown,
              );
            }

            if (state is AuthStateForgotPassword) {
              return const ForgotPasswordView();
            }

            if (state is AuthStateLoggedIn) {
              return LoggedInView(user: state.user);
            }

            return const Scaffold(
              body: Center(child: CircularProgressIndicator()),
            );
          },
        );
      }),
    );
  }
}
