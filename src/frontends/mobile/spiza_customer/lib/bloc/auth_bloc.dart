import 'package:rxdart/rxdart.dart';
import 'package:spiza_customer/models/address.dart';
import 'package:spiza_customer/models/user.dart';

class AuthBloc {
  final userSubject = BehaviorSubject<User>.seeded(
    User(
      id: 1,
      name: 'Placeholder User',
      email: 'placeholder@mail.com',
      address: Address(
        id: 2,
        fullAddress: 'Dobojska ul. 9, 10000, Zagreb',
        latLng: '45.79916506137784,15.95298740195445',
      ),
    ),
  );

  Stream<User> get user => userSubject.stream;

  void dispose() {
    userSubject.close();
  }

  void setUser(User user) {
    userSubject.add(user);
  }
}
