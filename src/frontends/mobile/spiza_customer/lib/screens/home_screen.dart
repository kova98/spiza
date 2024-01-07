import 'package:flutter/material.dart';
import 'package:spiza_customer/widgets/order_in_progress.dart';
import 'package:spiza_customer/widgets/restaurants_list.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final primaryColor = Theme.of(context).primaryColor;
    return Scaffold(
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 16),
          child: Stack(
            alignment: Alignment.bottomCenter,
            children: [
              ListView(
                physics: const BouncingScrollPhysics(),
                shrinkWrap: true,
                children: [
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      TextButton.icon(
                        style: TextButton.styleFrom(
                          padding: EdgeInsets.zero,
                        ),
                        label: const Text(
                          'A long address 5',
                          style: TextStyle(color: Colors.black87),
                        ),
                        icon: Icon(
                          Icons.location_on,
                          color: primaryColor,
                        ),
                        onPressed: () {},
                      ),
                      const SizedBox(height: 10),
                      const Text(
                        'What would\nyou like to eat?',
                        style: TextStyle(
                          fontWeight: FontWeight.bold,
                          fontSize: 32,
                        ),
                      ),
                    ],
                  ),
                  const SizedBox(height: 20),
                  TextField(
                    decoration: const InputDecoration(
                      prefixIcon: Icon(Icons.search),
                      border: UnderlineInputBorder(),
                    ),
                    onChanged: (value) {
                      // update
                    },
                  ),
                  const SizedBox(height: 20),
                  const Text(
                    'All Restaurants',
                    style: TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 24,
                    ),
                  ),
                  const SizedBox(height: 15),
                  const RestaurantsList()
                ],
              ),
              const OrderInProgress(),
            ],
          ),
        ),
      ),
    );
  }
}
