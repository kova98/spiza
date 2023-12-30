import 'package:spiza_customer/models/menu.dart';

class Restaurant {
  final int id;
  final String name;
  final Menu menu;

  Restaurant({required this.id, required this.name, required this.menu});

  Restaurant.empty()
      : id = 0,
        name = '',
        menu = Menu.empty();

  factory Restaurant.fromJson(Map<String, dynamic> json) {
    return Restaurant(
      id: json['id'],
      name: json['name'],
      menu: Menu.fromJson(json['menu_categories'] ?? [] as List),
    );
  }
}
