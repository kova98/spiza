import 'dart:convert';

import 'package:http/http.dart' show Client;

import 'package:spiza_customer/models/order.dart';

class OrderApiProvider {
  final _root = "10.0.2.2:5002";
  Client _client = Client();

  Future<String> submitOrder(Order order) async {
    final body = json.encode(order.toJson());
    final response = await _client.post(Uri.http(_root, '/api/order'),
        body: body, headers: {'Content-Type': 'application/json'});

    if (response.statusCode == 200) {
      return "";
    } else {
      return "Error: ${response.body}";
    }
  }
}
