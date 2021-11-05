class Item {
  final String id;
  final String name;
  final String description;
  final String image;
  final String category;
  final num price;
  final int order;

  Item.fromJson(Map<String, dynamic> parsedJson)
      : id = parsedJson['id'],
        name = parsedJson['name'],
        description = parsedJson['description'],
        image = parsedJson['image'],
        price = parsedJson['price'],
        category = parsedJson['category'],
        order = parsedJson['order'];
}
