import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/order_provider.dart';
import 'package:spiza_customer/models/order.dart';

class OrderScreen extends StatefulWidget {
  const OrderScreen({super.key});

  @override
  State<OrderScreen> createState() => _OrderScreenState();
}

class _OrderScreenState extends State<OrderScreen> {
  @override
  Widget build(BuildContext context) {
    final orderBloc = OrderProvider.of(context);

    return StreamBuilder<Order>(
        stream: orderBloc.order,
        builder: (context, snapshot) {
          if (snapshot.hasData == false) {
            orderBloc.refreshOrder();
          }
          return !snapshot.hasData
              ? const CircularProgressIndicator()
              : Scaffold(
                  backgroundColor: Colors.yellow,
                  appBar: AppBar(
                    backgroundColor: Colors.white,
                    elevation: 0,
                    centerTitle: true,
                    title: Text(
                      snapshot.data!.restaurantName ?? "",
                      style: const TextStyle(fontWeight: FontWeight.bold),
                    ),
                  ),
                  body: Column(
                    children: [
                      Container(
                        alignment: Alignment.center,
                        color: Colors.white,
                        height: 40,
                        child: const Text(
                          'current status',
                          style: const TextStyle(fontSize: 20),
                        ),
                      ),
                      const Expanded(child: Text('map')),
                      Container(
                          alignment: Alignment.center,
                          color: Colors.white,
                          height: 100,
                          child: const Row(
                            children: [
                              Padding(
                                padding: EdgeInsets.only(left: 20),
                                child: Text(
                                  '14:33',
                                  style: TextStyle(
                                      fontSize: 50,
                                      fontWeight: FontWeight.w500),
                                ),
                              ),
                              Padding(
                                padding: EdgeInsets.only(left: 20),
                                child: Text(
                                  'Estimated delivery time',
                                  style: TextStyle(
                                      fontSize: 20, color: Colors.black54),
                                ),
                              )
                            ],
                          )),
                    ],
                  ),
                );
        });
  }
}
