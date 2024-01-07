import 'package:spiza_customer/models/location.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/models/order.dart';

class Cart {
  List<Item> items = List<Item>.empty(growable: true);
  int addressId;
  int restaurantId;
  String restaurantName;
  Location restaurantLocation;
  Location? destinationLocation;
  num? deliveryTime;

  get totalPrice => items.fold<num>(0, (prev, i) => i.price + prev);

  Cart({
    required this.restaurantId,
    required this.restaurantName,
    required this.addressId,
    required this.restaurantLocation,
  });

  Cart.empty()
      : addressId = 0,
        restaurantName = '',
        restaurantId = 0,
        restaurantLocation = Location.empty();

  Order toOrder(int userId) {
    return Order(
      restaurantId: restaurantId,
      restaurantName: restaurantName,
      addressId: addressId,
      items: items.map((i) => i.id).toList(),
      userId: userId,
      restaurantLocation: restaurantLocation,
      destinationLocation: destinationLocation,
    );
  }
}
