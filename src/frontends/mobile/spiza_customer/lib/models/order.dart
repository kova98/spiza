import 'package:spiza_customer/models/location.dart';

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

  Order(
      {required this.items,
      required this.addressId,
      required this.restaurantId,
      this.restaurantName,
      this.restaurantLocation,
      required this.userId,
      this.destinationLocation,
      this.status = OrderStatus.created,
      this.deliveryTime,
      this.courierLocation});

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

  Order copyWith({
    List<int>? items,
    String? restaurantName,
    int? addressId,
    int? restaurantId,
    int? userId,
    OrderStatus? status,
    String? deliveryTime,
    Location? restaurantLocation,
    Location? courierLocation,
    Location? destinationLocation,
  }) {
    return Order(
      items: items ?? this.items,
      restaurantName: restaurantName ?? this.restaurantName,
      addressId: addressId ?? this.addressId,
      restaurantId: restaurantId ?? this.restaurantId,
      userId: userId ?? this.userId,
      status: status ?? this.status,
      deliveryTime: deliveryTime ?? this.deliveryTime,
      restaurantLocation: restaurantLocation ?? this.restaurantLocation,
      courierLocation: courierLocation ?? this.courierLocation,
      destinationLocation: destinationLocation ?? this.destinationLocation,
    );
  }

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
