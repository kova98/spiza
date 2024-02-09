import 'dart:convert';
import 'package:spiza_customer/models/restaurant.dart';
import 'package:http/http.dart' show Client;

class RestaurantApiProvider {
  final _root = "10.0.2.2:5002";
  final Client _client = Client();

  Future<List<Restaurant>> getRestaurants() async {
    final response = await _client.get(Uri.http(_root, '/api/restaurant'));

    if (response.statusCode == 200) {
      final decoded = json.decode(utf8.decode(response.bodyBytes)) as List;
      final restaurants = decoded.map((i) => Restaurant.fromJson(i)).toList();
      return restaurants;
    } else {
      throw Exception('Failed to fetch restaurants');
    }
  }

  Future<Restaurant> getRestaurantWithMenu(int id) async {
    final response = await _client.get(
      Uri.http(_root, '/api/restaurant/$id'),
    );

    if (response.statusCode == 200) {
      final decoded = json.decode(utf8.decode(response.bodyBytes));
      final restaurant = Restaurant.fromJson(decoded);
      return restaurant;
    } else {
      throw Exception('Failed to fetch restaurant');
    }
  }
}
