import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/bloc/auth_bloc.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/models/restaurant.dart';

class CartBloc {
  final _cart = BehaviorSubject<Cart>.seeded(Cart.empty());

  final AuthBloc authBloc;

  CartBloc(this.authBloc);

  Stream<Cart> get cart => _cart.stream;

  void addToCart(Item item) {
    final currentCart = _cart.value;
    final updatedCart = currentCart.copyWith(
      items: List.from(currentCart.items)..add(item),
    );
    _cart.sink.add(updatedCart);
  }

  void setRestaurant(Restaurant restaurant) {
    final updatedCart = _cart.value.copyWith(
      restaurantName: restaurant.name,
      restaurantId: restaurant.id,
      restaurantLocation: restaurant.address.getLocation(),
    );
    _cart.sink.add(updatedCart);
  }

  void dispose() {
    _cart.close();
  }

  void createCart(Restaurant res) {
    final address = authBloc.userSubject.value.address;

    final newCart = Cart(
      restaurantId: res.id,
      restaurantName: res.name,
      addressId: res.id,
      restaurantLocation: res.address.getLocation(),
      destinationLocation: address.getLocation(),
    );
    _cart.sink.add(newCart);
  }
}
