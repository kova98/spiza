import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/models/restaurant.dart';
import 'package:spiza_customer/resources/restaurant_repository.dart';

class RestaurantsBloc {
  final _repository = RestaurantRepository();
  final _restaurants = PublishSubject<List<Restaurant>>();

  String searchQuery = "";
  bool isSearching = false;

  Stream<List<Restaurant>> get restaurants => _restaurants.stream;

  Future getRestaurants() async {
    final restaurants = await _repository.getRestaurants();
    _restaurants.sink.add(restaurants);
  }

  void dispose() {
    _restaurants.close();
  }
}
