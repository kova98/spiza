import 'dart:convert';
import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:spiza_customer/bloc/order_provider.dart';
import 'package:spiza_customer/models/location.dart';
import 'package:spiza_customer/models/order.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:spiza_customer/utility/asset_helper.dart';

class OrderScreen extends StatefulWidget {
  const OrderScreen({super.key});

  @override
  State<OrderScreen> createState() => _OrderScreenState();
}

class _OrderScreenState extends State<OrderScreen> {
  final Map<String, BitmapDescriptor> _markerIcons = {
    'restaurant': BitmapDescriptor.defaultMarker,
    'courier': BitmapDescriptor.defaultMarker,
    'user': BitmapDescriptor.defaultMarker,
  };

  GoogleMapController? _mapController;

  @override
  void initState() {
    _markerIcons.forEach((key, value) {
      AssetHelper.getBytesFromAsset('assets/$key-icon.png', 128)
          .then((onValue) {
        setState(() {
          _markerIcons[key] = BitmapDescriptor.fromBytes(onValue);
        });
      });
    });
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final orderBloc = OrderProvider.of(context);

    return StreamBuilder<Order>(
      stream: orderBloc.order,
      builder: (context, snapshot) {
        if (!snapshot.hasData) {
          return const CircularProgressIndicator();
        }
        final restLoc = toLatLng(snapshot.data!.restaurantLocation!);
        final destLoc = toLatLng(snapshot.data!.destinationLocation!);
        _updateMapBounds(snapshot.data!);
        return Scaffold(
          backgroundColor: Colors.white,
          appBar: AppBar(
            backgroundColor: Colors.white,
            elevation: 0,
            centerTitle: true,
            title: Text(
              snapshot.data!.restaurantName ?? "",
              style: const TextStyle(fontWeight: FontWeight.bold),
            ),
          ),
          body: Stack(
            children: [
              map(snapshot, restLoc, destLoc, context),
              Column(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  statusBar(snapshot),
                  deliveryTimeBar(snapshot),
                ],
              ),
            ],
          ),
        );
      },
    );
  }

  Widget map(AsyncSnapshot<Order> snapshot, LatLng restLoc, LatLng destLoc,
      BuildContext context) {
    return GoogleMap(
      onMapCreated: (GoogleMapController c) {
        changeMapStyle(c, "assets/maps_style.json");
        _mapController = c;
        _updateMapBounds(snapshot.data!);
      },
      myLocationButtonEnabled: true,
      zoomControlsEnabled: false,
      initialCameraPosition: CameraPosition(
        target: _getCentralPoint(restLoc, destLoc),
        zoom: 16,
      ),
      markers: _getMarkers(context, snapshot.data!),
    );
  }

  Widget deliveryTimeBar(AsyncSnapshot<Order> snapshot) {
    return Container(
      decoration: BoxDecoration(
        color: Colors.white,
      ),
      alignment: Alignment.center,
      height: 80,
      child: snapshot.data!.deliveryTime != null
          ? Row(
              children: [
                Padding(
                  padding: const EdgeInsets.only(left: 20),
                  child: Text(
                    snapshot.data!.getTime(),
                    style: const TextStyle(
                        fontSize: 40, fontWeight: FontWeight.w500),
                  ),
                ),
                const Flexible(
                  child: Padding(
                    padding: EdgeInsets.symmetric(horizontal: 20),
                    child: Text(
                      'Estimated delivery time',
                      textAlign: TextAlign.end,
                      style: TextStyle(fontSize: 20, color: Colors.black54),
                    ),
                  ),
                )
              ],
            )
          : const Text(
              'Waiting for confirmation...',
              textAlign: TextAlign.end,
              style: TextStyle(fontSize: 20, color: Colors.black),
            ),
    );
  }

  Widget statusBar(AsyncSnapshot<Order> snapshot) {
    return Container(
      alignment: Alignment.center,
      height: 60,
      decoration: BoxDecoration(
        color: Colors.white,
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.2),
            spreadRadius: 1,
            blurRadius: 1,
            offset: const Offset(0, 2),
          ),
        ],
      ),
      child: Text(
        snapshot.data!.status.description,
        style: const TextStyle(fontSize: 20),
      ),
    );
  }

  Set<Marker> _getMarkers(BuildContext context, Order order) {
    return <Marker>{
      if (order.restaurantLocation != null)
        Marker(
          markerId: const MarkerId('restaurant'),
          position: toLatLng(order.restaurantLocation!),
          icon: _markerIcons['restaurant']!,
        ),
      if (order.courierLocation != null)
        Marker(
          markerId: const MarkerId('courier'),
          position: toLatLng(order.courierLocation!),
          icon: _markerIcons['courier']!,
        ),
      if (order.destinationLocation != null)
        Marker(
          markerId: const MarkerId('user'),
          position: toLatLng(order.destinationLocation!),
          icon: _markerIcons['user']!,
        ),
    };
  }

  LatLng toLatLng(Location location) {
    return LatLng(location.lat, location.lng);
  }

  void changeMapStyle(GoogleMapController mapController, String stylePath) {
    _getJsonFile(stylePath).then((value) => _setMapStyle(value, mapController));
  }

  void _setMapStyle(String mapStyle, GoogleMapController mapController) {
    mapController.setMapStyle(mapStyle);
  }

  Future<String> _getJsonFile(String path) async {
    var byte = await rootBundle.load(path);
    var list = byte.buffer.asUint8List(byte.offsetInBytes, byte.lengthInBytes);
    return utf8.decode(list);
  }

  LatLng _getCentralPoint(LatLng point1, LatLng point2) {
    var centralLat = (point1.latitude + point2.latitude) / 2;
    var centralLng = (point1.longitude + point2.longitude) / 2;
    return LatLng(centralLat, centralLng);
  }

  LatLngBounds _createBounds(List<LatLng> positions) {
    final southwestLat = positions.map((p) => p.latitude).reduce(min);
    final southwestLon = positions.map((p) => p.longitude).reduce(min);
    final northeastLat = positions.map((p) => p.latitude).reduce(max);
    final northeastLon = positions.map((p) => p.longitude).reduce((max));
    return LatLngBounds(
      southwest: LatLng(southwestLat, southwestLon),
      northeast: LatLng(northeastLat, northeastLon),
    );
  }

  void _updateMapBounds(Order order) {
    if (_mapController == null) return;

    final locs = [
      toLatLng(order.destinationLocation!),
      toLatLng(order.restaurantLocation!)
    ];
    if (order.courierLocation != null) {
      locs.add(toLatLng(order.courierLocation!));
    }

    final bounds = _createBounds(locs);
    _mapController!.animateCamera(CameraUpdate.newLatLngBounds(bounds, 30));
  }
}
