import 'package:spiza_customer/models/menu.dart';

class Restaurant {
  final String id;
  final String name;
  final Menu menu;

  Restaurant.fromJson(Map<String, dynamic> parsedJson)
      : id = parsedJson['id'],
        name = parsedJson['name'],
        menu = Menu.fromJson(parsedJson['menu']);
}
