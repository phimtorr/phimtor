import 'package:flutter/material.dart';
import 'package:phimtor_app/extensions/buildcontext/loc.dart';

class AccountView extends StatelessWidget {
  const AccountView({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(context.loc.account),
      ),
      body: const SingleChildScrollView(
        child: Padding(
          padding: EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text("Account"),
              SizedBox(height: 32),
              Text("Account"),
              SizedBox(height: 32),
              Text("Account"),
            ],
          ),
        ),
      ),
    );
  }
}
