import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/cart_provider.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/item.dart';

class CartScreen extends StatelessWidget {
  final menuText = TextStyle(fontSize: 24);

  @override
  Widget build(BuildContext context) {
    final cartBloc = CartProvider.of(context);

    return StreamBuilder<Cart>(
      stream: cartBloc.cart,
      builder: (context, snapshot) {
        if (!snapshot.hasData) {
          cartBloc.refreshCart();
          return CircularProgressIndicator();
        } else if (snapshot.data!.items.length == 0) {
          return SizedBox.shrink();
        } else {
          return Scaffold(
            appBar: AppBar(
              backgroundColor: Colors.transparent,
              elevation: 0,
              centerTitle: true,
              title: Text(snapshot.data!.restaurantName ?? ""),
            ),
            body: Container(
              padding: EdgeInsets.all(20),
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
                              style: TextStyle(fontSize: 24)),
                          Text("${item.price}", style: menuText),
                        ],
                      );
                    },
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text('Total:',
                          style: TextStyle(
                              fontSize: 24, fontWeight: FontWeight.bold)),
                      Text("${snapshot.data!.totalPrice}",
                          style: TextStyle(
                              fontSize: 24, fontWeight: FontWeight.bold))
                    ],
                  ),
                  Padding(
                    padding: const EdgeInsets.only(bottom: 10),
                    child: ConstrainedBox(
                      constraints: BoxConstraints.tightFor(
                        width: MediaQuery.of(context).size.width - 20,
                        height: 60,
                      ),
                      child: ElevatedButton(
                        onPressed: () {},
                        style: ElevatedButton.styleFrom(
                            textStyle: TextStyle(
                                color: Theme.of(context).primaryColor),
                            backgroundColor: Colors.amber,
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(30),
                            )),
                        child: Text(
                          'Confirm Order',
                          style: TextStyle(fontSize: 24, color: Colors.black),
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
}
