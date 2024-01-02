import 'package:flutter/material.dart';
import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/bloc/order_provider.dart';
import 'package:spiza_customer/models/order.dart';
import 'package:spiza_customer/models/order_update.dart';

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
        stream: Rx.combineLatest2(orderBloc.order, orderBloc.orderUpdate,
            (Order o, OrderUpdate u) {
          o.status = u.status;
          o.deliveryTime = u.deliveryTime;
          return o;
        }),
        builder: (context, snapshot) {
          if (snapshot.hasData == false) {
            orderBloc.refreshOrderUpdate();
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
                          height: 60,
                          child: Text(
                            snapshot.data!.status.description,
                            style: const TextStyle(fontSize: 20),
                          ),
                          decoration:
                              BoxDecoration(color: Colors.white, boxShadow: [
                            BoxShadow(
                                color: Colors.black.withOpacity(0.1),
                                spreadRadius: 3,
                                blurRadius: 3,
                                offset: const Offset(0, 1))
                          ])),
                      const Expanded(child: Text('map')),
                      Container(
                          alignment: Alignment.center,
                          color: Colors.white,
                          height: 100,
                          child: Row(
                            children: [
                              Padding(
                                padding: const EdgeInsets.only(left: 20),
                                child: Text(
                                  getTime(snapshot.data!.deliveryTime),
                                  style: const TextStyle(
                                      fontSize: 50,
                                      fontWeight: FontWeight.w500),
                                ),
                              ),
                              const Padding(
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

  String getTime(String? deliveryTime) {
    if (deliveryTime == null) {
      return '00:00';
    } else {
      final utcTime = DateTime.parse(deliveryTime).toUtc();
      final localTime = utcTime.toLocal();
      return '${localTime.hour}:${localTime.minute}';
    }
  }
}
