import 'package:spiza_customer/models/item.dart';

class Cart {
  List<Item> items = List<Item>();
  String address;
  String restaurantName;
  num deliveryTime;

  get totalPrice => items.fold(0, (prev, i) => i.price + prev);
}
