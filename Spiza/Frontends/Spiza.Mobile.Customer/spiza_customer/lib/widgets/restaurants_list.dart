import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/restaurants_provider.dart';
import 'package:spiza_customer/models/restaurant.dart';

class RestaurantsList extends StatelessWidget {
  Widget build(context) {
    final bloc = RestaurantsProvider.of(context);
    return StreamBuilder(
      stream: bloc.restaurants,
      builder: (context, AsyncSnapshot<List<Restaurant>> snapshot) {
        if (!snapshot.hasData) {
          bloc.getRestaurants();
          return Center(
            child: CircularProgressIndicator(),
          );
        } else {
          return ListView.builder(
            shrinkWrap: true,
            itemCount: snapshot.data.length,
            itemBuilder: (context, int index) {
              final item = snapshot.data[index];
              return Padding(
                padding: const EdgeInsets.all(8.0),
                child: Column(
                  children: [Text(item.name)],
                ),
              );
            },
          );
        }
      },
    );
  }
}
