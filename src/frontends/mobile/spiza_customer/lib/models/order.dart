import 'package:spiza_customer/models/item.dart';

class Order {
  List<int> items = List<int>.empty(growable: true);
  String? address;
  String? restaurantName;
  num restaurantId = 0;
  num userId = 0;
  OrderStatus status = OrderStatus.created;

  String? deliveryTime;
  String? restaurantLocation;
  String? driverLocation;
  String? destinationLocation;

  Order({
    required this.items,
    required this.address,
    required this.restaurantId,
    this.restaurantName,
    required this.userId,
  });

  String getTime() {
    if (deliveryTime == null) {
      return '00:00';
    } else {
      final utcTime = DateTime.parse(deliveryTime ?? '').toUtc();
      final localTime = utcTime.toLocal();
      return '${localTime.hour}:${localTime.minute}';
    }
  }

  factory Order.fromJson(Map<String, dynamic> json) {
    var itemsList = json['items'] as List;
    List<Item> items = itemsList.map((i) => Item.fromJson(i)).toList();
    return Order(
      userId: json['user_id'],
      restaurantId: json['restaurant_id'],
      address: json['address'],
      items: items.map((i) => i.id).toList(),
    );
  }

  Map<String, dynamic> toJson() => {
        'user_id': userId,
        'restaurant_id': restaurantId,
        'address': address,
        'items': items,
      };

  static Order empty() {
    return Order(
      userId: 0,
      restaurantId: 0,
      address: '',
      items: List<int>.empty(),
    );
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
