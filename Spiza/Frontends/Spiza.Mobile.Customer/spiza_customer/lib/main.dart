import 'package:flutter/material.dart';
import 'package:spiza_customer/screens/home_screen.dart';

main() => runApp(SpizaApp());

class SpizaApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Spiza',
      theme: ThemeData(
        primaryColor: Colors.amber,
        accentColor: Colors.red,
      ),
      home: HomeScreen(),
    );
  }
}
