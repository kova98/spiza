import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/cart_provider.dart';
import 'package:spiza_customer/bloc/order_provider.dart';
import 'package:spiza_customer/bloc/restaurants_provider.dart';
import 'package:spiza_customer/screens/home_screen.dart';

main() => runApp(const SpizaApp());

class SpizaApp extends StatelessWidget {
  const SpizaApp({super.key});

  @override
  Widget build(BuildContext context) {
    return RestaurantsProvider(
      child: CartProvider(
        child: OrderProvider(
          child: MaterialApp(
            debugShowCheckedModeBanner: false,
            title: 'Spiza',
            theme: ThemeData(
                primaryColor: Colors.amber[600],
                colorScheme: ColorScheme.fromSwatch()
                    .copyWith(secondary: Colors.red, background: Colors.white)),
            home: const HomeScreen(),
          ),
        ),
      ),
    );
  }
}
