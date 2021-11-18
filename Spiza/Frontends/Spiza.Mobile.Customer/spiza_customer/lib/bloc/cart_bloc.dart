import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/data/cart_api_provider.dart';

class CartBloc {
  final _api = CartApiProvider();
  final _cart = PublishSubject<Cart>();

  static Cart _lastCart = Cart();

  Stream<Cart> get cart => _cart.stream;

  void addToCart(Item item) {
    _lastCart.items.add(item);
    _cart.sink.add(_lastCart);
  }

  void refreshCart() {
    _cart.sink.add(_lastCart);
  }

  void dispose() {
    _cart.close();
  }
}
