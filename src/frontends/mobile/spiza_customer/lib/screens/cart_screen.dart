import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:spiza_customer/bloc/cart_provider.dart';
import 'package:spiza_customer/bloc/order_bloc.dart';
import 'package:spiza_customer/bloc/order_provider.dart';
import 'package:spiza_customer/models/cart.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/screens/order_screen.dart';
import 'package:google_static_maps_controller/google_static_maps_controller.dart';

class CartScreen extends StatelessWidget {
  const CartScreen({super.key});
  final houseIcon = "https://i.imgur.com/PARCU0w.png";
  final restaurantIcon = "https://i.imgur.com/1suBNlj.png";

  @override
  Widget build(BuildContext context) {
    final cartBloc = CartProvider.of(context);
    final orderBloc = OrderProvider.of(context);

    return StreamBuilder<Cart>(
      stream: cartBloc.cart,
      builder: (context, snapshot) {
        if (!snapshot.hasData) {
          return const CircularProgressIndicator();
        } else if (snapshot.data!.items.isEmpty) {
          return const SizedBox.shrink();
        } else {
          return Scaffold(
            appBar: AppBar(
              backgroundColor: Colors.transparent,
              elevation: 0,
              centerTitle: true,
              title: Text(snapshot.data!.restaurantName),
            ),
            body: Container(
              padding: const EdgeInsets.all(20),
              child: Column(
                children: [
                  StaticMap(
                    styles: const [
                      MapStyle(
                        element: StyleElement.labels,
                        feature: StyleFeature.transit,
                        rules: [StyleRule.visibility(VisibilityRule.off)],
                      ),
                      MapStyle(
                        element: StyleElement.labels,
                        feature: StyleFeature.all,
                        rules: [StyleRule.visibility(VisibilityRule.off)],
                      ),
                      MapStyle(
                        element: StyleElement.labels,
                        feature: StyleFeature.road,
                        rules: [StyleRule.visibility(VisibilityRule.on)],
                      ),
                    ],
                    googleApiKey: "AIzaSyB7VBHijYya2Wd49MPLBXG_BaQw-1jtU0Y",
                    width: MediaQuery.of(context).size.width,
                    height: 300,
                    scaleToDevicePixelRatio: true,
                    zoom: 16,
                    visible: [
                      GeocodedLocation.latLng(
                        snapshot.data!.restaurantLocation.lat,
                        snapshot.data!.restaurantLocation.lng,
                      ),
                      GeocodedLocation.latLng(
                          snapshot.data!.destinationLocation!.lat,
                          snapshot.data!.destinationLocation!.lng),
                    ],
                    markers: <Marker>[
                      Marker.custom(locations: [
                        Location(
                          snapshot.data!.destinationLocation!.lat,
                          snapshot.data!.destinationLocation!.lng,
                        ),
                      ], icon: houseIcon, anchor: MarkerAnchor.center),
                      Marker.custom(locations: [
                        Location(
                          snapshot.data!.restaurantLocation.lat,
                          snapshot.data!.restaurantLocation.lng,
                        ),
                      ], icon: restaurantIcon),
                    ],
                  ),
                  SizedBox(
                    height: 20,
                  ),
                  Row(
                    children: [
                      Icon(
                        Icons.home,
                        size: 20,
                      ),
                      const SizedBox(
                        width: 5,
                      ),
                      Text(
                        snapshot.data!.addressName,
                        style: TextStyle(fontSize: 20),
                      ),
                    ],
                  ),
                  const Row(
                    children: [
                      Icon(
                        Icons.timer_outlined,
                        size: 20,
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      Text(
                        "30-40 min",
                        style: TextStyle(fontSize: 20),
                      ),
                    ],
                  ),
                  const Spacer(),
                  ListView.builder(
                    shrinkWrap: true,
                    itemCount: snapshot.data!.items.length,
                    itemBuilder: (context, index) {
                      final Item item = snapshot.data!.items[index];
                      return Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Text("${item.amount}  ${item.name}",
                              style: const TextStyle(
                                fontSize: 20,
                                color: Colors.black,
                              )),
                          Text(
                            "${getPrice(item)}€",
                            style: const TextStyle(fontSize: 20),
                          ),
                        ],
                      );
                    },
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      const Text('Total:',
                          style: TextStyle(
                              fontSize: 20, fontWeight: FontWeight.bold)),
                      Text("${snapshot.data!.totalPrice}€",
                          style: const TextStyle(
                              fontSize: 20, fontWeight: FontWeight.bold))
                    ],
                  ),
                  Padding(
                    padding: const EdgeInsets.only(top: 10),
                    child: ConstrainedBox(
                      constraints: BoxConstraints.tightFor(
                        width: MediaQuery.of(context).size.width - 20,
                        height: 60,
                      ),
                      child: ElevatedButton(
                        onPressed: () =>
                            confirmOrder(snapshot.data!, context, orderBloc),
                        style: ElevatedButton.styleFrom(
                            textStyle: TextStyle(
                                color: Theme.of(context).primaryColor),
                            backgroundColor: Colors.amber,
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(30),
                            )),
                        child: const Text(
                          'Confirm Order',
                          style: TextStyle(
                            fontSize: 20,
                            color: Colors.black,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                      ),
                    ),
                  )
                ],
              ),
            ),
          );
        }
      },
    );
  }

  void confirmOrder(Cart cart, BuildContext context, OrderBloc orderBloc) {
    orderBloc.confirmOrder(cart).then((value) => {
          if (value.$2 == "")
            {
              orderBloc.getOrderStatus(value.$1),
              Navigator.pushReplacement(
                context,
                MaterialPageRoute(
                  builder: (context) => const OrderScreen(),
                ),
              ),
            }
          else
            {}
        });
  }

  getPrice(Item item) {
    return item.price * item.amount;
  }
}
