import 'package:spiza_customer/models/location.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/models/order.dart';

class Cart {
  List<Item> items = List<Item>.empty(growable: true);
  int addressId;
  String addressName;
  int restaurantId;
  int destinationId;
  String restaurantName;
  Location restaurantLocation;
  Location? destinationLocation;
  num? deliveryTime;

  get totalPrice => items.fold<num>(0, (prev, i) => prev + i.price * i.amount);
  get totalAmount => items.fold<int>(0, (prev, i) => i.amount + prev);

  Cart({
    this.items = const [],
    required this.addressId,
    required this.restaurantId,
    required this.restaurantName,
    required this.restaurantLocation,
    required this.destinationId,
    this.destinationLocation,
    this.deliveryTime,
    required this.addressName,
  });

  Cart copyWith({
    List<Item>? items,
    int? addressId,
    int? restaurantId,
    int? destinationId,
    String? restaurantName,
    Location? restaurantLocation,
    Location? destinationLocation,
    num? deliveryTime,
    String? addressName,
  }) {
    return Cart(
      addressId: addressId ?? this.addressId,
      restaurantId: restaurantId ?? this.restaurantId,
      destinationId: destinationId ?? this.destinationId,
      restaurantName: restaurantName ?? this.restaurantName,
      restaurantLocation: restaurantLocation ?? this.restaurantLocation,
      destinationLocation: destinationLocation ?? this.destinationLocation,
      items: items ?? this.items,
      deliveryTime: deliveryTime ?? this.deliveryTime,
      addressName: addressName ?? this.addressName,
    );
  }

  Cart.empty()
      : addressId = 0,
        addressName = '',
        restaurantName = '',
        restaurantId = 0,
        destinationId = 0,
        restaurantLocation = Location.empty();

  Order toOrder(int userId) {
    return Order(
      restaurantId: restaurantId,
      restaurantName: restaurantName,
      addressId: addressId,
      destinationId: destinationId,
      items: items.map((i) => i.id).toList(),
      userId: userId,
      restaurantLocation: restaurantLocation,
      destinationLocation: destinationLocation,
    );
  }
}
