import 'package:spiza_customer/models/address.dart';
import 'package:spiza_customer/models/menu.dart';

class Restaurant {
  final int id;
  final String name;
  final Menu menu;
  final Address address;

  Restaurant(
      {required this.id,
      required this.name,
      required this.menu,
      required this.address});

  Restaurant.empty()
      : id = 0,
        name = '',
        menu = Menu.empty(),
        address = Address.empty();

  Restaurant.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        name = json['name'],
        menu = Menu.fromJson(json['menuCategories'] ?? []),
        address = Address.fromJson(json['address']);
}
