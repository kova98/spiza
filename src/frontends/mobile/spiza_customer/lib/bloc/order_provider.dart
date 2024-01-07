import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/order_bloc.dart';

class OrderProvider extends InheritedWidget {
  final OrderBloc bloc;

  OrderProvider({Key? key, required Widget child})
      : bloc = OrderBloc(),
        super(key: key, child: child);

  @override
  bool updateShouldNotify(oldWidget) => true;

  static OrderBloc of(BuildContext context) {
    return context.dependOnInheritedWidgetOfExactType<OrderProvider>()!.bloc;
  }
}
