import 'package:spiza_customer/models/location.dart';

class Address {
  final int id;
  final String fullAddress;
  final String latLng;

  Address({required this.id, required this.fullAddress, required this.latLng});

  Address.fromJson(Map<String, dynamic> json)
      : id = json['id'] ?? 0,
        fullAddress = json['fullAddress'],
        latLng = json['latLng'];

  Address.empty()
      : id = 0,
        fullAddress = '',
        latLng = '';

  Location getLocation() {
    if (this.latLng.isEmpty) {
      return Location.empty();
    }
    final latLng = this.latLng.split(',');
    return Location(lat: double.parse(latLng[0]), lng: double.parse(latLng[1]));
  }
}
