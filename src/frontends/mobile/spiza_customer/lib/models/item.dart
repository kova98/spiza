class Item {
  final int id;
  final String name;
  final String description;
  final String image;
  final int categoryId;
  final num price;
  final int order;
  int amount;

  Item.fromJson(Map<String, dynamic> parsedJson)
      : id = parsedJson['id'],
        name = parsedJson['name'],
        description = parsedJson['description'],
        image = parsedJson['image'],
        price = parsedJson['price'],
        categoryId = parsedJson['categoryId'],
        amount = parsedJson['amount'] ?? 1,
        order = parsedJson['order'];

  toJson() {
    return {
      'id': id,
      'name': name,
      'description': description,
      'image': image,
      'price': price,
      'categoryId': categoryId,
      'amount': amount,
      'order': order,
    };
  }
}
