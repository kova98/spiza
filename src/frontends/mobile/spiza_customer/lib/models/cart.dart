import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/models/order.dart';

class Cart {
  List<Item> items = List<Item>.empty(growable: true);
  String? address;
  String restaurantName;
  num restaurantId = 0;
  num? deliveryTime;
  get totalPrice => items.fold<num>(0, (prev, i) => i.price + prev);

  Cart({this.restaurantId = 0, required this.restaurantName, this.address});

  Order toOrder(num userId) {
    print(restaurantName);
    return Order(
        restaurantId: restaurantId,
        restaurantName: restaurantName,
        address: address,
        items: items.map((i) => i.id).toList(),
        userId: userId);
  }
}
