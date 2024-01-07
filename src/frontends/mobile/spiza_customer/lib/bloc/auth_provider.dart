import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/auth_bloc.dart';

class AuthProvider extends InheritedWidget {
  final AuthBloc bloc;

  const AuthProvider({Key? key, required Widget child, required this.bloc})
      : super(key: key, child: child);

  @override
  bool updateShouldNotify(oldWidget) => true;

  static AuthBloc of(BuildContext context) {
    return context.dependOnInheritedWidgetOfExactType<AuthProvider>()!.bloc;
  }
}
