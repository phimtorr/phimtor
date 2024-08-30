import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';
import 'package:phimtor_app/services/auth/auth_user.dart';
import 'package:phimtor_app/services/auth/bloc/auth_bloc.dart';
import 'package:phimtor_app/services/auth/bloc/auth_event.dart';

class LoggedInView extends StatelessWidget {
  const LoggedInView({
    super.key,
    required this.user,
  });

  final AuthUser user;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(context.loc.account_title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(context.loc.account_message(user.email)),
            const SizedBox(height: 8),
            if (user.premiumUntil != null)
              Text(context.loc.account_preminum_expire(user.premiumUntil!)),
            const SizedBox(height: 8),
            ElevatedButton(
              onPressed: () {
                BlocProvider.of<AuthBloc>(context).add(const AuthEventLogOut());
              },
              child: Text(context.loc.logout),
            ),
          ],
        ),
      ),
    );
  }
}
