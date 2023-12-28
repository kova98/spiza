import 'package:spiza_customer/models/item.dart';

class Menu {
  final List<String> categories;
  final List<Item> items;

  Menu.empty()
      : categories = List<String>.empty(),
        items = List<Item>.empty();

  Menu.fromJson(Map<String, dynamic> parsedJson)
      : categories = parsedJson["categories"] != null
            ? (parsedJson['categories'] as List).cast<String>()
            : List<String>.empty(),
        items = parsedJson['items'] != null
            ? (parsedJson['items'] as List)
                .map<Item>((e) => Item.fromJson(e))
                .toList()
            : List<Item>.empty();
}
