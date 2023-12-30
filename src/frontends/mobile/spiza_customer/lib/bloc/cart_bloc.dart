import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/data/order_api_provider.dart';

class CartBloc {
  final _api = OrderApiProvider();
  final _cart = PublishSubject<Cart>();

  static Cart _lastCart = Cart();

  Stream<Cart> get cart => _cart.stream;

  void addToCart(Item item) {
    _lastCart.items.add(item);
    refreshCart();
  }

  void setRestaurantName(String name) {
    _lastCart.restaurantName = name;
    refreshCart();
  }

  void refreshCart() {
    // TODO: support multiple restaurants
    _cart.sink.add(_lastCart);
  }

  void dispose() {
    _cart.close();
  }

  void confirmOrder() {
    var order = _lastCart.toOrder(1);
    _api.submitOrder(order);
  }

  void createCart(int restaurantId) {
    _lastCart = Cart(restaurantId: restaurantId);
  }
}
