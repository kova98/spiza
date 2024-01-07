import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/cart_provider.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/restaurant.dart';
import 'package:spiza_customer/screens/cart_screen.dart';

class OrderButton extends StatelessWidget {
  final Restaurant restaurant;

  const OrderButton(this.restaurant, {super.key});

  @override
  Widget build(BuildContext context) {
    final cartBloc = CartProvider.of(context);

    return StreamBuilder(
      stream: cartBloc.cart,
      builder: (context, AsyncSnapshot<Cart> snapshot) {
        if (!snapshot.hasData) {
          cartBloc.createCart(restaurant);
          return const SizedBox.shrink();
        } else if (snapshot.data!.items.isEmpty) {
          return const SizedBox.shrink();
        } else {
          return Padding(
            padding: const EdgeInsets.only(bottom: 10),
            child: ConstrainedBox(
              constraints: BoxConstraints.tightFor(
                width: MediaQuery.of(context).size.width - 20,
                height: 60,
              ),
              child: ElevatedButton(
                onPressed: () => Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (context) => const CartScreen(),
                  ),
                ),
                style: ElevatedButton.styleFrom(
                    textStyle: TextStyle(color: Theme.of(context).primaryColor), backgroundColor: Colors.amber,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(30),
                    )),
                child: Text(
                  'Order ${snapshot.data!.items.length} for ${snapshot.data!.totalPrice} kn',
                  style: const TextStyle(fontSize: 24, color: Colors.black),
                ),
              ),
            ),
          );
        }
      },
    );
  }
}
