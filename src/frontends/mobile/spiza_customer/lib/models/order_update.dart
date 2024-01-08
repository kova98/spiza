import 'package:spiza_customer/models/order.dart';

class OrderUpdate {
  final int orderId;
  final String? deliveryTime;
  final OrderStatus status;

  OrderUpdate({this.orderId = 0, required this.status, this.deliveryTime});

  OrderUpdate.fromJson(Map<String, dynamic> json)
      : orderId = json['order_id'] ?? 0,
        status = OrderStatus.values[json['status']],
        deliveryTime = json['delivery_time'];
}
