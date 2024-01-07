import 'package:spiza_customer/models/address.dart';

class User {
  int id = 0;
  String name = '';
  String email = '';
  Address address = Address.empty();
  String? token;

  User({
    required this.id,
    required this.name,
    required this.email,
    required this.address,
    this.token,
  });
}
