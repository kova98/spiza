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

  Restaurant.fromJson(Map<String, dynamic> parsedJson)
      : id = parsedJson['id'],
        name = parsedJson['name'],
        menu = Menu.fromJson(parsedJson['menu'] ?? {});
}
