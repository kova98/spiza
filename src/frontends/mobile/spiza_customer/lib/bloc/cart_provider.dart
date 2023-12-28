import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/cart_bloc.dart';

class CartProvider extends InheritedWidget {
  final CartBloc bloc;

  CartProvider({Key? key, required Widget child})
      : bloc = CartBloc(),
        super(key: key, child: child);

  bool updateShouldNotify(_) => true;

  static CartBloc of(BuildContext context) {
    return context.dependOnInheritedWidgetOfExactType<CartProvider>()!.bloc;
  }
}
