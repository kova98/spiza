import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/order_bloc.dart';

class OrderProvider extends InheritedWidget {
  final OrderBloc bloc;

  const OrderProvider({Key? key, required Widget child, required this.bloc})
      : super(key: key, child: child);

  @override
  bool updateShouldNotify(oldWidget) => true;

  static OrderBloc of(BuildContext context) {
    return context.dependOnInheritedWidgetOfExactType<OrderProvider>()!.bloc;
  }
}
