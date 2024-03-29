import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/bloc/auth_bloc.dart';
import 'package:spiza_customer/data/mqtt_provider.dart';
import 'package:spiza_customer/data/order_api_provider.dart';
import 'package:spiza_customer/models/location.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/order.dart';
import 'package:spiza_customer/models/order_update.dart';

class OrderBloc {
  final BehaviorSubject<Order> _orderSubject =
      BehaviorSubject.seeded(Order.empty());
  final BehaviorSubject<OrderUpdate> _orderUpdateSubject =
      BehaviorSubject.seeded(OrderUpdate(status: OrderStatus.created));
  final BehaviorSubject<Location> _courierLocationSubject =
      BehaviorSubject.seeded(Location.empty());
  final _api = OrderApiProvider();

  final MqttProvider _mqtt = MqttProvider();
  final AuthBloc authBloc;

  Stream<Order> get order => _orderSubject.stream;
  Stream<OrderUpdate> get orderUpdate => _orderUpdateSubject.stream;
  Stream<Location> get courierLocation => _courierLocationSubject.stream;

  OrderBloc(this.authBloc) {
    _orderUpdateSubject.stream.listen((event) {
      var updatedOrder = _orderSubject.value.copyWith(
        deliveryTime: event.deliveryTime,
        status: event.status,
      );
      _orderSubject.add(updatedOrder);
    });

    _courierLocationSubject.stream.listen((event) {
      var updatedOrder = _orderSubject.value.copyWith(
        courierLocation: event,
      );
      _orderSubject.add(updatedOrder);
    });
  }

  Future getOrderStatus(int orderId) async {
    try {
      await _mqtt.connectToMQTT();
      _mqtt.subscribe(
        'order/$orderId',
        _orderUpdateSubject,
        OrderUpdate.fromJson,
      );
      _mqtt.subscribe(
        'order/$orderId/courier-location',
        _courierLocationSubject,
        Location.fromJson,
      );
    } catch (e) {
      print(e); // TODO: Replace with proper error handling
    }
  }

  void dispose() {
    _orderSubject.close();
    _orderUpdateSubject.close();
    _courierLocationSubject.close();
  }

  Future<(int id, String error)> confirmOrder(Cart cart) async {
    final user = authBloc.userSubject.value;
    var newOrder = cart.toOrder(user.id);
    _orderSubject.add(newOrder);
    return _api.submitOrder(newOrder);
  }
}
