import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/restaurants_provider.dart';
import 'package:spiza_customer/models/restaurant.dart';
import 'package:spiza_customer/screens/menu_screen.dart';

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
            physics: ClampingScrollPhysics(),
            itemCount: snapshot.data.length,
            itemBuilder: (context, int index) {
              final item = snapshot.data[index];
              return InkWell(
                onTap: () => Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (context) => MenuScreen(restaurant: item),
                  ),
                ),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    new Container(
                      height: 150.0,
                      alignment: Alignment.center,
                      decoration: new BoxDecoration(
                        borderRadius: BorderRadius.all(Radius.circular(7)),
                        image: DecorationImage(
                            image: AssetImage('assets/burger.png'),
                            fit: BoxFit.cover),
                      ),
                    ),
                    SizedBox(height: 5),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          item.name,
                          style: TextStyle(
                            fontSize: 18,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                        Text(
                          '★ 4.5',
                          style: TextStyle(
                            fontSize: 18,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                      ],
                    ),
                    Text(
                      '5,00 kn',
                      style: TextStyle(color: Colors.black87),
                    ),
                    SizedBox(height: 17)
                  ],
                ),
              );
            },
          );
        }
      },
    );
  }
}
