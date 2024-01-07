import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/order_provider.dart';
import 'package:spiza_customer/models/order.dart';
import 'package:spiza_customer/screens/order_screen.dart';

class OrderInProgress extends StatelessWidget {
  const OrderInProgress({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    final orderBloc = OrderProvider.of(context);

    return Padding(
      padding: const EdgeInsets.only(bottom: 16),
      child: StreamBuilder<Order>(
        stream: orderBloc.order,
        builder: (context, snapshot) {
          if (!snapshot.hasData || snapshot.data!.restaurantId == 0) {
            return Container();
          }
          return Container(
            alignment: Alignment.bottomCenter,
            height: 80,
            decoration: BoxDecoration(
              boxShadow: [
                BoxShadow(
                  color: Colors.black.withOpacity(0.2),
                  spreadRadius: 3,
                  blurRadius: 3,
                  offset: const Offset(0, 1),
                ),
              ],
              color: Colors.white,
              borderRadius: BorderRadius.circular(10),
              border: Border.all(color: Colors.black, width: 2),
            ),
            child: InkWell(
              onTap: () => Navigator.push(context,
                  MaterialPageRoute(builder: (context) => const OrderScreen())),
              child: Row(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  const Padding(padding: EdgeInsets.only(left: 16)),
                  Text(
                    snapshot.data!.getTime(),
                    style: const TextStyle(
                        fontSize: 30, fontWeight: FontWeight.bold),
                  ),
                  const Padding(padding: EdgeInsets.only(left: 16)),
                  Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        snapshot.data!.restaurantName ?? "",
                        style: const TextStyle(fontWeight: FontWeight.w500),
                      ),
                      Text(snapshot.data!.status.description)
                    ],
                  ),
                ],
              ),
            ),
          );
        },
      ),
    );
  }
}
