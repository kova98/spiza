import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/models/restaurant.dart';

class CartBloc {
  final _cart = PublishSubject<Cart>();

  static Cart _lastCart = Cart.empty();

  Stream<Cart> get cart => _cart.stream;

  void addToCart(Item item) {
    _lastCart.items.add(item);
    refreshCart();
  }

  void setRestaurant(Restaurant restaurant) {
    _lastCart.restaurantName = restaurant.name;
    _lastCart.restaurantId = restaurant.id;
    _lastCart.restaurantLocation = restaurant.address.getLocation();
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
    _lastCart = Cart(
      restaurantId: res.id,
      restaurantName: res.name,
      addressId: res.id,
      restaurantLocation: res.address.getLocation(),
    );
    _cart.sink.add(_lastCart);
  }
}
