import 'package:spiza_customer/models/restaurant.dart';
import 'package:spiza_customer/resources/restaurant_api_provider.dart';

class RestaurantRepository {
  final reustaurantApiProvider = RestaurantApiProvider();

  Future<List<Restaurant>> getRestaurants() {
    return reustaurantApiProvider.getRestaurants();
  }
}
