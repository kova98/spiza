import 'package:spiza_customer/models/restaurant.dart';
import 'package:spiza_customer/data/restaurant_api_provider.dart';

class RestaurantRepository {
  final reustaurantApiProvider = RestaurantApiProvider();

  Future<List<Restaurant>> getRestaurants() {
    return reustaurantApiProvider.getRestaurants();
  }

  Future<Restaurant> getRestaurantWithMenu(String id) {
    return reustaurantApiProvider.getRestaurantWithMenu(id);
  }
}
