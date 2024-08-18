import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_bloc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_event.dart';

class VerifyEmailView extends StatefulWidget {
  const VerifyEmailView({super.key, this.needCooldown = false});

  final bool needCooldown;

  @override
  State<VerifyEmailView> createState() => _VerifyEmailViewState();
}

class _VerifyEmailViewState extends State<VerifyEmailView> {
  Timer? _timer;
  int _cooldown = 60;

  @override
  void initState() {
    super.initState();
    if (widget.needCooldown) {
      startTimer();
    } else {
      _timer = null;
      _cooldown = 0;
    }
  }

  void startTimer() {
    _timer?.cancel();
    _cooldown = 60;
    _timer = Timer.periodic(const Duration(seconds: 1), (timer) {
      if (_cooldown == 0) {
        timer.cancel();
      } else {
        setState(() {
          _cooldown--;
        });
      }
    });
  }

  @override
  void dispose() {
   _timer?.cancel();
    super.dispose();
  }

  void sendEmailVerification() {
    BlocProvider.of<AuthBloc>(context)
        .add(const AuthEventSendEmailVerification());
    startTimer();
  }

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
            if ( _cooldown > 0)
              Text(context.loc.verify_email_resend_in(_cooldown)),
            if ( _cooldown == 0)
              ElevatedButton(
                onPressed: sendEmailVerification,
                child: Text(context.loc.verify_email_resend),
              ),
            const SizedBox(height: 8),
            TextButton(
              onPressed: () {
                BlocProvider.of<AuthBloc>(context).add(const AuthEventLogOut());
              },
              child: Text(context.loc.verify_email_have_confirm),
            ),
          ],
        ),
      ),
    );
  }
}
