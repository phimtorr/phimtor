import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_bloc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_event.dart';

class VerifyEmailView extends StatefulWidget {
  const VerifyEmailView({super.key});

  @override
  State<VerifyEmailView> createState() => _VerifyEmailViewState();
}

class _VerifyEmailViewState extends State<VerifyEmailView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(context.loc.verify_email),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(context.loc.verify_email_note),
            const SizedBox(height: 8),
            ElevatedButton(
              onPressed: () {
                BlocProvider.of<AuthBloc>(context)
                    .add(const AuthEventSendEmailVerification());
              },
              child: Text(context.loc.verify_email_resend),
            ),
            const SizedBox(height: 8),
            TextButton(
              onPressed: () {
                BlocProvider.of<AuthBloc>(context)
                    .add(const AuthEventLogOut());
              },
              child: Text(context.loc.verify_email_have_confirm),
            ),
          ],
        ),
      ),
    );
  }
}
