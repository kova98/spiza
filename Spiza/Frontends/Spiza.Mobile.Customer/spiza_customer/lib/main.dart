import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/cart_provider.dart';
import 'package:spiza_customer/bloc/restaurants_provider.dart';
import 'package:spiza_customer/screens/home_screen.dart';

main() => runApp(SpizaApp());

class SpizaApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return RestaurantsProvider(
      child: CartProvider(
        child: MaterialApp(
          debugShowCheckedModeBanner: false,
          title: 'Spiza',
          theme: ThemeData(
              primaryColor: Colors.amber[600],
              accentColor: Colors.red,
              backgroundColor: Colors.white),
          home: HomeScreen(),
        ),
      ),
    );
  }
}
