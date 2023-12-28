import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/models/restaurant.dart';
import 'package:spiza_customer/data/restaurant_repository.dart';

class RestaurantsBloc {
  final _repository = RestaurantRepository();
  final _restaurants = PublishSubject<List<Restaurant>>();

  List<Restaurant> _restaurantsList = List<Restaurant>.empty();

  Stream<List<Restaurant>> get restaurants => _restaurants.stream;

  Future getRestaurants() async {
    _restaurantsList = await _repository.getRestaurants();
    _restaurants.sink.add(_restaurantsList);
  }

  Future getRestaurantWithMenu(int id) async {
    final restaurant = await _repository.getRestaurantWithMenu(id);
    final existingRestaurant = _restaurantsList.firstWhere(
      (x) => x.id == id,
      orElse: () => Restaurant.empty(),
    );
    if (existingRestaurant != null) {
      final index = _restaurantsList.indexOf(existingRestaurant);
      _restaurantsList[index] = restaurant;
    }
    _restaurants.sink.add(_restaurantsList);
  }

  void dispose() {
    _restaurants.close();
  }
}
