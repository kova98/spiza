import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/cart_provider.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/screens/cart_screen.dart';

class OrderButton extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final cartBloc = CartProvider.of(context);

    return StreamBuilder(
      stream: cartBloc.cart,
      builder: (context, AsyncSnapshot<Cart> snapshot) {
        if (!snapshot.hasData) {
          cartBloc.refreshCart();
          return SizedBox.shrink();
        } else if (snapshot.data!.items.length == 0) {
          return SizedBox.shrink();
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
                    builder: (context) => CartScreen(),
                  ),
                ),
                style: ElevatedButton.styleFrom(
                    textStyle: TextStyle(color: Theme.of(context).primaryColor),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(30),
                    ),
                    primary: Colors.amber),
                child: Text(
                  'Order ${snapshot.data!.items.length} for ${snapshot.data!.totalPrice} kn',
                  style: TextStyle(fontSize: 24, color: Colors.black),
                ),
              ),
            ),
          );
        }
      },
    );
  }
}
