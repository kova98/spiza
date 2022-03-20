import 'package:spiza_customer/models/item.dart';

class Menu {
  final List<String> categories;
  final List<Item> items;

  Menu.fromJson(Map<String, dynamic> parsedJson)
      : categories = parsedJson != null
            ? (parsedJson['categories'] as List).cast<String>()
            : List<String>(),
        items = parsedJson != null
            ? (parsedJson['items'] as List)
                .map<Item>((e) => Item.fromJson(e))
                .toList()
            : List<Item>();
}
