class Restaurant {
  final String id;
  final String name;

  Restaurant({this.id, this.name});

  Restaurant.fromJson(Map<String, dynamic> parsedJson)
      : id = parsedJson['id'],
        name = parsedJson['name'];
}
