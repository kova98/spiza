import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/models/restaurant.dart';

class CartBloc {
  final _cart = PublishSubject<Cart>();

  static Cart _lastCart = Cart(restaurantName: '');

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

  void createCart(Restaurant res) {
    _lastCart = Cart(restaurantId: res.id, restaurantName: res.name);
    _cart.sink.add(_lastCart);
  }
}
