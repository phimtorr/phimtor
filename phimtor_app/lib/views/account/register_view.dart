import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_bloc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_event.dart';
import 'package:phimtor_app/services/auth/bloc/auth_state.dart';

class RegisterView extends StatefulWidget {
  const RegisterView({super.key});

  @override
  State<RegisterView> createState() => _RegisterViewState();
}

class _RegisterViewState extends State<RegisterView> {
  late final TextEditingController _email;
  late final TextEditingController _password;
  late final TextEditingController _confirmPassword;

  @override
  void initState() {
    _email = TextEditingController();
    _password = TextEditingController();
    _confirmPassword = TextEditingController();
    super.initState();
  }

  @override
  void dispose() {
    _email.dispose();
    _password.dispose();
    _confirmPassword.dispose();
    super.dispose();
  }

  void register() {
    if (_email.text.isEmpty) {
      showError(context.loc.email_invalid);
      return;
    }
    if (_password.text.isEmpty) {
      showError(context.loc.password_invalid);
      return;
    }
    if (_password.text != _confirmPassword.text) {
      showError(context.loc.register_view_password_not_match);
      return;
    }

    BlocProvider.of<AuthBloc>(context).add(
      AuthEventReigister(
        email: _email.text,
        password: _password.text,
      ),
    );
  }

  void showError(String message) {
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(
        content: Text(message),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(context.loc.register),
      ),
      body: BlocListener<AuthBloc, AuthState>(
        listener: (context, state) {
          if (state is AuthStateRegistering) {
            if (state.exception != null) {
              showError(context.loc.error(state.exception.toString()));
            }
          }
        },
        child: Padding(
          padding: const EdgeInsets.all(16),
          child: Column(
            children: [
              TextField(
                controller: _email,
                decoration: InputDecoration(
                  labelText: context.loc.email,
                ),
              ),
              const SizedBox(height: 8),
              TextField(
                controller: _password,
                decoration: InputDecoration(
                  labelText: context.loc.password,
                ),
                obscureText: true,
              ),
              const SizedBox(height: 8),
              TextField(
                controller: _confirmPassword,
                decoration: InputDecoration(
                  labelText: context.loc.confirm_password,
                ),
                obscureText: true,
              ),
              const SizedBox(height: 8),
              ElevatedButton(
                onPressed: register,
                child: Text(context.loc.register),
              ),
              const SizedBox(height: 8),
              TextButton(
                onPressed: () {
                  BlocProvider.of<AuthBloc>(context)
                      .add(const AuthEventLogOut());
                },
                child: Text(context.loc.register_view_already_have_account),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
