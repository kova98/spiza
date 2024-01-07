import 'package:spiza_customer/models/location.dart';
import 'package:spiza_customer/models/item.dart';

class Order {
  List<int> items = List<int>.empty(growable: true);
  String? restaurantName;
  int addressId;
  int restaurantId = 0;
  int userId = 0;
  OrderStatus status = OrderStatus.created;

  String? deliveryTime;
  Location? restaurantLocation;
  Location? courierLocation;
  Location? destinationLocation;

  Order({
    required this.items,
    required this.addressId,
    required this.restaurantId,
    this.restaurantName,
    this.restaurantLocation,
    required this.userId,
    this.destinationLocation,
  });

  Order.empty()
      : userId = 0,
        restaurantId = 0,
        addressId = 0,
        items = List<int>.empty();

  Map<String, dynamic> toJson() => {
        'user_id': userId,
        'restaurant_id': restaurantId,
        'address': addressId,
        'items': items,
      };

  String getTime() {
    if (deliveryTime == null) {
      return '00:00';
    } else {
      final utcTime = DateTime.parse(deliveryTime ?? '').toUtc();
      final localTime = utcTime.toLocal();
      return '${localTime.hour}:${localTime.minute}';
    }
  }
}

enum OrderStatus {
  created,
  accepted,
  rejected,
  ready,
  pickedUp,
  delivered,
}

extension OrderStatusExtension on OrderStatus {
  String get description {
    switch (this) {
      case OrderStatus.created:
        return 'Your order has been placed.';
      case OrderStatus.accepted:
        return 'Your order is in progress.';
      case OrderStatus.rejected:
        return 'Your order has been rejected.';
      case OrderStatus.ready:
        return 'Your order is ready for pickup.';
      case OrderStatus.pickedUp:
        return 'Your order is on the way.';
      case OrderStatus.delivered:
        return 'Your order has been delivered.';
      default:
        return 'unknown status';
    }
  }
}
