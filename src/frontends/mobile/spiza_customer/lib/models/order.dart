import 'package:spiza_customer/models/item.dart';

class Order {
  List<int> items = List<int>.empty(growable: true);
  String? address;
  String? restaurantName;
  num restaurantId = 0;
  num userId = 0;

  Order({
    required this.items,
    required this.address,
    required this.restaurantId,
    this.restaurantName,
    required this.userId,
  });

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
