import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/restaurants_bloc.dart';

class RestaurantsProvider extends InheritedWidget {
  final RestaurantsBloc bloc;

  RestaurantsProvider({Key? key, required Widget child})
      : bloc = RestaurantsBloc(),
        super(key: key, child: child);

  @override
  bool updateShouldNotify(oldWidget) => true;

  static RestaurantsBloc of(BuildContext context) {
    return context
        .dependOnInheritedWidgetOfExactType<RestaurantsProvider>()!
        .bloc;
  }
}
