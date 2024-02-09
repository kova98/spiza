import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/cart_provider.dart';
import 'package:spiza_customer/bloc/order_bloc.dart';
import 'package:spiza_customer/bloc/order_provider.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/screens/order_screen.dart';

class CartScreen extends StatelessWidget {
  final menuText = const TextStyle(fontSize: 24);

  const CartScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final cartBloc = CartProvider.of(context);
    final orderBloc = OrderProvider.of(context);

    return StreamBuilder<Cart>(
      stream: cartBloc.cart,
      builder: (context, snapshot) {
        if (!snapshot.hasData) {
          return const CircularProgressIndicator();
        } else if (snapshot.data!.items.isEmpty) {
          return const SizedBox.shrink();
        } else {
          return Scaffold(
            appBar: AppBar(
              backgroundColor: Colors.transparent,
              elevation: 0,
              centerTitle: true,
              title: Text(snapshot.data!.restaurantName),
            ),
            body: Container(
              padding: const EdgeInsets.all(20),
              child: Column(
                children: [
                  ListView.builder(
                    shrinkWrap: true,
                    itemCount: snapshot.data!.items.length,
                    itemBuilder: (context, index) {
                      final Item item = snapshot.data!.items[index];
                      return Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Text("${item.amount} x ${item.name}",
                              style: const TextStyle(fontSize: 24)),
                          Text("${getPrice(item)}€", style: menuText),
                        ],
                      );
                    },
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      const Text('Total:',
                          style: TextStyle(
                              fontSize: 24, fontWeight: FontWeight.bold)),
                      Text("${snapshot.data!.totalPrice}€",
                          style: const TextStyle(
                              fontSize: 24, fontWeight: FontWeight.bold))
                    ],
                  ),
                  const Spacer(),
                  Padding(
                    padding: const EdgeInsets.only(bottom: 10),
                    child: ConstrainedBox(
                      constraints: BoxConstraints.tightFor(
                        width: MediaQuery.of(context).size.width - 20,
                        height: 60,
                      ),
                      child: ElevatedButton(
                        onPressed: () =>
                            confirmOrder(snapshot.data!, context, orderBloc),
                        style: ElevatedButton.styleFrom(
                            textStyle: TextStyle(
                                color: Theme.of(context).primaryColor),
                            backgroundColor: Colors.amber,
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(30),
                            )),
                        child: const Text(
                          'Confirm Order',
                          style: TextStyle(
                            fontSize: 20,
                            color: Colors.black,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                      ),
                    ),
                  )
                ],
              ),
            ),
          );
        }
      },
    );
  }

  void confirmOrder(Cart cart, BuildContext context, OrderBloc orderBloc) {
    orderBloc.confirmOrder(cart).then((value) => {
          if (value.$2 == "")
            {
              orderBloc.getOrderStatus(value.$1),
              Navigator.pushReplacement(
                context,
                MaterialPageRoute(
                  builder: (context) => const OrderScreen(),
                ),
              ),
            }
          else
            {}
        });
  }

  getPrice(Item item) {
    return item.price * item.amount;
  }
}
