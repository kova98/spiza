import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/cart_bloc.dart';

class CartProvider extends InheritedWidget {
  final CartBloc bloc;

  const CartProvider({Key? key, required Widget child, required this.bloc})
      : super(key: key, child: child);

  @override
  bool updateShouldNotify(oldWidget) => true;

  static CartBloc of(BuildContext context) {
    return context.dependOnInheritedWidgetOfExactType<CartProvider>()!.bloc;
  }
}
