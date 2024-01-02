import 'dart:convert';

import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/data/mqtt_provider.dart';
import 'package:spiza_customer/data/order_api_provider.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/order.dart';
import 'package:spiza_customer/models/order_update.dart';

class OrderBloc {
  final _orderSubject = PublishSubject<Order>();
  final _orderUpdateSubject = PublishSubject<OrderUpdate>();
  final _api = OrderApiProvider();

  Order _order = Order.empty();
  OrderUpdate _orderUpdate = OrderUpdate(status: OrderStatus.created);

  Stream<Order> get order => _orderSubject.stream;
  Stream<OrderUpdate> get orderUpdate => _orderUpdateSubject.stream;

  Future getOrderStatus(int orderId) async {
    print('connecting to mqtt with order id $orderId');
    final mqtt = MqttProvider();
    mqtt.connectToMQTT().then(
        (value) => {mqtt.subscribe('order/$orderId', _orderUpdateSubject)});
    _orderSubject.sink.add(_order);
  }

  void dispose() {
    _orderSubject.close();
  }

  void refreshOrder() {
    _orderSubject.sink.add(_order);
  }

  void refreshOrderUpdate() {
    _orderUpdateSubject.sink.add(_orderUpdate);
  }

  Future<(int id, String error)> confirmOrder(Cart cart) {
    final userId = 1; // TODO: get user id from auth
    _order = cart.toOrder(userId);
    refreshOrder();
    return _api.submitOrder(_order);
  }
}
