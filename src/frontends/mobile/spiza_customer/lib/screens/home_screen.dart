import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/auth_provider.dart';
import 'package:spiza_customer/models/user.dart';
import 'package:spiza_customer/widgets/order_in_progress.dart';
import 'package:spiza_customer/widgets/restaurants_list.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});
  @override
  Widget build(BuildContext context) {
    final primaryColor = Theme.of(context).primaryColor;
    final authBloc = AuthProvider.of(context);
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
                      StreamBuilder<User>(
                          stream: authBloc.user,
                          builder: (context, snapshot) {
                            final address = snapshot.hasData
                                ? snapshot.data!.address.fullAddress.isEmpty
                                    ? 'Add address'
                                    : snapshot.data!.address.fullAddress
                                : '';
                            if (address.isEmpty) {}
                            return TextButton.icon(
                              style: TextButton.styleFrom(
                                padding: EdgeInsets.zero,
                              ),
                              label: Text(
                                address,
                                style: const TextStyle(color: Colors.black87),
                              ),
                              icon: Icon(
                                Icons.location_on,
                                color: primaryColor,
                              ),
                              onPressed: () {},
                            );
                          }),
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
