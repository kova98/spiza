import 'dart:convert';

import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/data/order_api_provider.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/order.dart';

class OrderBloc {
  final _orderSubject = PublishSubject<Order>();
  final _api = OrderApiProvider();

  Order _order = Order.empty();

  Stream<Order> get order => _orderSubject.stream;

  Future getOrderStatus() async {
    // get order stream from MQTT broker
    _orderSubject.sink.add(_order);
  }

  void dispose() {
    _orderSubject.close();
  }

  void refreshOrder() {
    _orderSubject.sink.add(_order);
  }

  Future<String> confirmOrder(Cart cart) {
    final userId = 1; // TODO: get user id from auth
    _order = cart.toOrder(userId);
    refreshOrder();
    return _api.submitOrder(_order);
  }
}
