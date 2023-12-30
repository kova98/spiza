class Item {
  final int id;
  final String name;
  final String description;
  final String image;
  final int categoryId;
  final num price;
  final int amount;
  final int order;

  Item.fromJson(Map<String, dynamic> parsedJson)
      : id = parsedJson['id'],
        name = parsedJson['name'],
        description = parsedJson['description'],
        image = parsedJson['image'],
        price = parsedJson['price'],
        categoryId = parsedJson['category_id'],
        amount = parsedJson['amount'] ?? 0,
        order = parsedJson['order'];

  toJson() {
    return {
      'id': id,
      'name': name,
      'description': description,
      'image': image,
      'price': price,
      'category_id': categoryId,
      'amount': amount,
      'order': order,
    };
  }
}
