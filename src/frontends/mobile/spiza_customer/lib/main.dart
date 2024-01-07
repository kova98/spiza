import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/auth_bloc.dart';
import 'package:spiza_customer/bloc/auth_provider.dart';
import 'package:spiza_customer/bloc/cart_bloc.dart';
import 'package:spiza_customer/bloc/cart_provider.dart';
import 'package:spiza_customer/bloc/order_bloc.dart';
import 'package:spiza_customer/bloc/order_provider.dart';
import 'package:spiza_customer/bloc/restaurants_bloc.dart';
import 'package:spiza_customer/bloc/restaurants_provider.dart';
import 'package:spiza_customer/screens/home_screen.dart';

main() {
  final cartBloc = CartBloc();
  final orderBloc = OrderBloc();
  final restaurantsBloc = RestaurantsBloc();
  final authBloc = AuthBloc();

  runApp(SpizaApp(
    cartBloc: cartBloc,
    orderBloc: orderBloc,
    restaurantsBloc: restaurantsBloc,
    authBloc: authBloc,
  ));
}

class SpizaApp extends StatelessWidget {
  final CartBloc cartBloc;
  final OrderBloc orderBloc;
  final RestaurantsBloc restaurantsBloc;
  final AuthBloc authBloc;

  const SpizaApp({
    Key? key,
    required this.cartBloc,
    required this.orderBloc,
    required this.restaurantsBloc,
    required this.authBloc,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AuthProvider(
      bloc: authBloc,
      child: RestaurantsProvider(
        bloc: restaurantsBloc,
        child: CartProvider(
          bloc: cartBloc,
          child: OrderProvider(
            bloc: orderBloc,
            child: MaterialApp(
              debugShowCheckedModeBanner: false,
              title: 'Spiza',
              theme: ThemeData(
                primaryColor: Colors.amber[600],
                colorScheme: ColorScheme.fromSwatch()
                    .copyWith(secondary: Colors.red, background: Colors.white),
              ),
              home: const HomeScreen(),
            ),
          ),
        ),
      ),
    );
  }
}
