import 'package:spiza_customer/models/address.dart';
import 'package:spiza_customer/models/menu.dart';

class Restaurant {
  final int id;
  final String name;
  final Menu menu;
  final Address address;
  final String image;
  final num rating;
  final num deliveryPrice;

  Restaurant(
      {required this.id,
      required this.image,
      required this.name,
      required this.menu,
      required this.address,
      required this.rating,
      required this.deliveryPrice});

  Restaurant.empty()
      : id = 0,
        name = '',
        menu = Menu.empty(),
        address = Address.empty(),
        image = '',
        rating = 0.0,
        deliveryPrice = 0.0;

  Restaurant.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        name = json['name'],
        menu = Menu.fromJson(json['menuCategories'] ?? []),
        address = Address.fromJson(json['address']),
        image = json['image'] ?? '',
        rating = json['rating'] ?? 0.0,
        deliveryPrice = json['deliveryPrice'] ?? 0.0;
}
