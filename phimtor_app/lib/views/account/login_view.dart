import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_bloc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_event.dart';

class LoginView extends StatefulWidget {
  const LoginView({super.key});

  @override
  State<LoginView> createState() => _LoginViewState();
}

class _LoginViewState extends State<LoginView> {
  late final TextEditingController _email;
  late final TextEditingController _password;

  @override
  void initState() {
    _email = TextEditingController();
    _password = TextEditingController();
    super.initState();
  }

  @override
  void dispose() {
    _email.dispose();
    _password.dispose();
    super.dispose();
  }

  void login() {
    if (_email.text.isEmpty) {
      showError(context.loc.email_invalid);
      return;
    }
    if (_password.text.isEmpty) {
      showError(context.loc.password_invalid);
      return;
    }

    BlocProvider.of<AuthBloc>(context).add(
      AuthEventLogIn(
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
    const smallSpacer = SizedBox(height: 8);
    return Scaffold(
      appBar: AppBar(
        title: Text(context.loc.login),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              TextField(
                controller: _email,
                decoration: InputDecoration(
                  labelText: context.loc.email,
                  hintText: context.loc.email_hint,
                ),
              ),
              smallSpacer,
              TextField(
                controller: _password,
                decoration: InputDecoration(
                  labelText: context.loc.password,
                  hintText: context.loc.password_hint,
                ),
                obscureText: true,
              ),
              smallSpacer,
              ElevatedButton(
                onPressed: login,
                child: Text(context.loc.login),
              ),
              smallSpacer,
              TextButton(
                onPressed: () {
                  BlocProvider.of<AuthBloc>(context)
                      .add(const AuthEventForgotPassword());
                },
                child: Text(context.loc.login_view_forgot_password),
              ),
              smallSpacer,
              TextButton(
                onPressed: () {
                  BlocProvider.of<AuthBloc>(context)
                      .add(const AuthEventShouldRegister());
                },
                child: Text(context.loc.login_view_not_register_yet),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
