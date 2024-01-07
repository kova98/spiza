import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/restaurants_bloc.dart';

class RestaurantsProvider extends InheritedWidget {
  final RestaurantsBloc bloc;

  const RestaurantsProvider(
      {Key? key, required Widget child, required this.bloc})
      : super(key: key, child: child);

  @override
  bool updateShouldNotify(oldWidget) => true;

  static RestaurantsBloc of(BuildContext context) {
    return context
        .dependOnInheritedWidgetOfExactType<RestaurantsProvider>()!
        .bloc;
  }
}
