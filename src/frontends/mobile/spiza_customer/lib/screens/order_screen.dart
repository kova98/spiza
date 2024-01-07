import 'dart:convert';

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
  final LatLng _center = const LatLng(45.8150, 15.9819);

  final Map<String, BitmapDescriptor> _markerIcons = {
    'restaurant': BitmapDescriptor.defaultMarker,
    'courier': BitmapDescriptor.defaultMarker,
    'user': BitmapDescriptor.defaultMarker,
  };

  @override
  void initState() {
    _markerIcons.forEach((key, value) {
      AssetHelper.getBytesFromAsset('assets/$key-icon.png', 64).then((onValue) {
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
          if (snapshot.hasData == false) {
            orderBloc.refreshOrderUpdate();
            orderBloc.refreshOrder();
          }
          return !snapshot.hasData
              ? const CircularProgressIndicator()
              : Scaffold(
                  backgroundColor: Colors.yellow,
                  appBar: AppBar(
                    backgroundColor: Colors.white,
                    elevation: 0,
                    centerTitle: true,
                    title: Text(
                      snapshot.data!.restaurantName ?? "",
                      style: const TextStyle(fontWeight: FontWeight.bold),
                    ),
                  ),
                  body: Column(
                    children: [
                      Container(
                          alignment: Alignment.center,
                          height: 60,
                          decoration:
                              BoxDecoration(color: Colors.white, boxShadow: [
                            BoxShadow(
                                color: Colors.black.withOpacity(0.1),
                                spreadRadius: 3,
                                blurRadius: 3,
                                offset: const Offset(0, 1))
                          ]),
                          child: Text(
                            snapshot.data!.status.description,
                            style: const TextStyle(fontSize: 20),
                          )),
                      Expanded(
                        child: GoogleMap(
                          onMapCreated: (GoogleMapController c) {
                            changeMapMode(c, "assets/maps_style.json");
                          },
                          myLocationButtonEnabled: true,
                          zoomControlsEnabled: false,
                          initialCameraPosition: CameraPosition(
                            target: _center,
                            zoom: 18,
                          ),
                          markers: _getMarkers(context, snapshot.data!),
                        ),
                      ),
                      Container(
                          alignment: Alignment.center,
                          color: Colors.white,
                          height: 100,
                          child: Row(
                            children: [
                              Padding(
                                padding: const EdgeInsets.only(left: 20),
                                child: Text(
                                  snapshot.data!.getTime(),
                                  style: const TextStyle(
                                      fontSize: 50,
                                      fontWeight: FontWeight.w500),
                                ),
                              ),
                              const Padding(
                                padding: EdgeInsets.only(left: 20),
                                child: Text(
                                  'Estimated delivery time',
                                  style: TextStyle(
                                      fontSize: 20, color: Colors.black54),
                                ),
                              )
                            ],
                          )),
                    ],
                  ),
                );
        });
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

  void changeMapMode(GoogleMapController mapController, String stylePath) {
    _getJsonFile(stylePath).then((value) => _setMapStyle(value, mapController));
  }

  void _setMapStyle(String mapStyle, GoogleMapController mapController) {
    mapController.setMapStyle(mapStyle);
  }

  Future<String> _getJsonFile(String path) async {
    ByteData byte = await rootBundle.load(path);
    var list = byte.buffer.asUint8List(byte.offsetInBytes, byte.lengthInBytes);
    return utf8.decode(list);
  }
}
