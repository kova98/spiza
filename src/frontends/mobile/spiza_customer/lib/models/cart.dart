import 'package:spiza_customer/models/item.dart';

class Cart {
  List<Item> items = List<Item>.empty();
  String? address;
  String? restaurantName;
  num? deliveryTime;

  get totalPrice => items.fold<num>(0, (prev, i) => i.price + prev);
}
