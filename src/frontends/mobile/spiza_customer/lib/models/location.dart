class Location {
  double lat;
  double lng;

  Location({required this.lat, required this.lng});

  Location.fromJson(Map<String, dynamic> json)
      : lat = json['lat'],
        lng = json['lng'];

  Location.empty()
      : lat = 0,
        lng = 0;
}
