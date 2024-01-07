import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/data/mqtt_provider.dart';
import 'package:spiza_customer/data/order_api_provider.dart';
import 'package:spiza_customer/models/location.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/order.dart';
import 'package:spiza_customer/models/order_update.dart';

class OrderBloc {
  final _orderSubject = PublishSubject<Order>();
  final _orderUpdateSubject = PublishSubject<OrderUpdate>();
  final _courierLocationSubject = PublishSubject<Location>();
  final _api = OrderApiProvider();

  Order _order = Order.empty();
  OrderUpdate _orderUpdate = OrderUpdate(status: OrderStatus.created);

  Stream<Order> get order => _orderSubject.stream;
  Stream<OrderUpdate> get orderUpdate => _orderUpdateSubject.stream;
  Stream<Location> get courierLocation => _courierLocationSubject.stream;

  OrderBloc() {
    _orderUpdateSubject.listen((event) {
      _orderUpdate = event;
      _order.deliveryTime = event.deliveryTime;
      _order.status = event.status;
      refreshOrder();
    });

    _courierLocationSubject.listen((event) {
      _order.courierLocation = event;
      refreshOrder();
    });
  }

  Future getOrderStatus(int orderId) async {
    final mqtt = MqttProvider();
    mqtt.connectToMQTT().then((value) => {
          mqtt.subscribe(
            'order/$orderId',
            _orderUpdateSubject,
            OrderUpdate.fromJson,
          ),
          mqtt.subscribe(
            'order/$orderId/courier-location',
            _courierLocationSubject,
            Location.fromJson,
          )
        });
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
    const userId = 1; // TODO: get user id from auth
    _order = cart.toOrder(userId);
    refreshOrder();
    return _api.submitOrder(_order);
  }
}
