import 'package:spiza_customer/models/item.dart';

class MenuCategory {
  final int id;
  final String name;
  final int restaurantId;
  final List<Item> items;

  MenuCategory({
    required this.id,
    required this.name,
    required this.restaurantId,
    required this.items,
  });

  factory MenuCategory.fromJson(Map<String, dynamic> json) {
    var itemsList = json['items'] as List;
    List<Item> items = itemsList.map((i) => Item.fromJson(i)).toList();
    return MenuCategory(
      id: json['id'],
      name: json['name'],
      restaurantId: json['restaurantId'],
      items: items,
    );
  }
}

class Menu {
  final List<MenuCategory> categories;

  Menu({required this.categories});

  factory Menu.empty() {
    return Menu(categories: List<MenuCategory>.empty());
  }

  factory Menu.fromJson(List<dynamic> jsonList) {
    List<MenuCategory> categories =
        jsonList.map((e) => MenuCategory.fromJson(e)).toList();
    return Menu(categories: categories);
  }
}
