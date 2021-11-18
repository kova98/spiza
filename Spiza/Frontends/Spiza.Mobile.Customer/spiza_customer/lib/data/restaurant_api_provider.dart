import 'dart:convert';
import 'package:spiza_customer/models/restaurant.dart';
import 'package:http/http.dart' show Client;

class RestaurantApiProvider {
  final _root = "http://10.0.2.2:5101/api";
  Client _client = Client();

  Future<List<Restaurant>> getRestaurants() async {
    final response = await _client.get(
      '$_root/restaurant',
    );

    if (response.statusCode == 200) {
      final decoded = json.decode(response.body) as List;
      final restaurants = decoded.map((i) => Restaurant.fromJson(i)).toList();
      return restaurants;
    } else {
      throw Exception('Failed to fetch restaurants');
    }
  }

  Future<Restaurant> getRestaurantWithMenu(String id) async {
    final response = await _client.get(
      '$_root/restaurant/$id',
    );

    if (response.statusCode == 200) {
      final decoded = json.decode(response.body);
      final restaurant = Restaurant.fromJson(decoded);
      return restaurant;
    } else {
      throw Exception('Failed to fetch restaurant');
    }
  }
}
