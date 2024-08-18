import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_bloc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_event.dart';
import 'package:phimtor_app/services/auth/bloc/auth_state.dart';

class ForgotPasswordView extends StatefulWidget {
  const ForgotPasswordView({super.key});

  @override
  State<ForgotPasswordView> createState() => _ForgotPasswordViewState();
}

class _ForgotPasswordViewState extends State<ForgotPasswordView> {
  late final TextEditingController _email;

  @override
  void initState() {
    super.initState();
    _email = TextEditingController();
  }

  @override
  void dispose() {
    _email.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return BlocListener<AuthBloc, AuthState>(
      listener: (context, state) {
        if (state is AuthStateForgotPassword) {
          if (state.exception != null) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(
                content: Text(context.loc.error(state.exception.toString())),
              ),
            );
          }
        }
      },
      child: Scaffold(
          appBar: AppBar(
            title: Text(context.loc.forgot_password),
          ),
          body: Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              children: [
                Text(context.loc.forgot_password_note),
                const SizedBox(height: 8),
                Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: TextField(
                    controller: _email,
                    decoration: InputDecoration(
                      labelText: context.loc.email,
                    ),
                  ),
                ),
                const SizedBox(height: 8),
                ElevatedButton(
                  onPressed: () {
                    BlocProvider.of<AuthBloc>(context).add(
                      AuthEventForgotPassword(email: _email.text),
                    );
                  },
                  child: Text(context.loc.forgot_password_view_send_me_link),
                ),
                const SizedBox(height: 8),
                TextButton(
                  onPressed: () {
                    BlocProvider.of<AuthBloc>(context)
                        .add(const AuthEventLogOut());
                  },
                  child: Text(context.loc.forgot_password_view_back_to_login),
                ),
              ],
            ),
          )),
    );
  }
}
