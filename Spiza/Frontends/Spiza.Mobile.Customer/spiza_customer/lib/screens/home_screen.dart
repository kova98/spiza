import 'package:flutter/material.dart';
import 'package:spiza_customer/widgets/restaurants_list.dart';

class HomeScreen extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  @override
  Widget build(BuildContext context) {
    final primaryColor = Theme.of(context).primaryColor;

    return Scaffold(
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 16),
          child: ListView(
            physics: BouncingScrollPhysics(),
            shrinkWrap: true,
            children: [
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  TextButton.icon(
                    style: TextButton.styleFrom(
                      padding: EdgeInsets.zero,
                    ),
                    label: Text(
                      'A long address 5',
                      style: TextStyle(color: Colors.black87),
                    ),
                    icon: Icon(
                      Icons.location_on,
                      color: primaryColor,
                    ),
                    onPressed: () {},
                  ),
                  SizedBox(height: 10),
                  Text(
                    'What would\nyou like to eat?',
                    style: TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 32,
                    ),
                  ),
                ],
              ),
              SizedBox(height: 20),
              TextField(
                decoration: InputDecoration(
                  prefixIcon: Icon(Icons.search),
                  border: UnderlineInputBorder(),
                ),
                onChanged: (value) {
                  // update
                },
              ),
              SizedBox(height: 20),
              Text(
                'All Restaurants',
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                  fontSize: 24,
                ),
              ),
              SizedBox(height: 15),
              RestaurantsList()
            ],
          ),
        ),
      ),
    );
  }
}
