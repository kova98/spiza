import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/models/order.dart';

class Cart {
  List<Item> items = List<Item>.empty(growable: true);
  String? address;
  String? restaurantName;
  num restaurantId = 0;
  num? deliveryTime;
  get totalPrice => items.fold<num>(0, (prev, i) => i.price + prev);

  Cart({this.restaurantId = 0});

  toOrder(num userId) {
    return Order(
        restaurantId: restaurantId,
        address: address,
        items: items.map((i) => i.id).toList(),
        userId: userId);
  }
}
