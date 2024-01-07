import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/models/restaurant.dart';
import 'package:spiza_customer/data/restaurant_repository.dart';

class RestaurantsBloc {
  final _repository = RestaurantRepository();
  final BehaviorSubject<List<Restaurant>> _restaurants =
      BehaviorSubject.seeded(List<Restaurant>.empty());

  Stream<List<Restaurant>> get restaurants => _restaurants.stream;

  Future getRestaurants() async {
    try {
      var restaurantsList = await _repository.getRestaurants();
      _restaurants.sink.add(restaurantsList);
    } catch (e) {
      // TODO: handle the error or forward it to the UI
      print(e);
    }
  }

  Future getRestaurantWithMenu(int id) async {
    try {
      final restaurant = await _repository.getRestaurantWithMenu(id);
      var existingRestaurantIndex =
          _restaurants.value.indexWhere((x) => x.id == id);
      if (existingRestaurantIndex >= 0) {
        var updatedList = List<Restaurant>.from(_restaurants.value);
        updatedList[existingRestaurantIndex] = restaurant;
        _restaurants.sink.add(updatedList);
      }
    } catch (e) {
      // TODO: Handle the error or forward it to the UI
      print(e);
    }
  }

  void dispose() {
    _restaurants.close();
  }
}
