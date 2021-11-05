import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/restaurants_provider.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/models/restaurant.dart';

class MenuList extends StatelessWidget {
  final Restaurant restaurant;

  const MenuList(this.restaurant);

  @override
  Widget build(BuildContext context) {
    final bloc = RestaurantsProvider.of(context);

    return StreamBuilder(
      stream: bloc.restaurants,
      builder: (context, AsyncSnapshot<List<Restaurant>> snapshot) {
        if (!snapshot.hasData) {
          bloc.getRestaurantWithMenu(restaurant.id);
          return Center(
            child: CircularProgressIndicator(),
          );
        } else {
          return ListView.builder(
            shrinkWrap: true,
            physics: ClampingScrollPhysics(),
            itemCount: restaurant.menu.categories.length,
            itemBuilder: (context, int index) {
              final category = restaurant.menu.categories[index];
              final items = restaurant.menu.items
                  .where((element) => element.category == category)
                  .toList();
              return items.length == 0
                  ? SizedBox.shrink()
                  : Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Padding(
                          padding: const EdgeInsets.symmetric(vertical: 12),
                          child: Text(
                            category,
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                              fontSize: 24,
                            ),
                          ),
                        ),
                        buildMenuItems(items)
                      ],
                    );
            },
          );
        }
      },
    );
  }
}

Widget buildMenuItems(List<Item> items) {
  List<Widget> menuItems = List<Widget>();

  for (var i = 0; i < items.length; i++) {
    final item = items[i];
    menuItems.add(
      Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  item.name,
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                    fontSize: 18,
                  ),
                ),
                SizedBox(height: 5),
                Text(
                  item.description,
                  style: TextStyle(color: Colors.black54),
                ),
                SizedBox(height: 5),
                Text(
                  "${item.price} kn",
                )
              ],
            ),
          ),
          ClipRRect(
            borderRadius: BorderRadius.circular(5),
            child: SizedBox(
              width: 150,
              child: Image(
                image: AssetImage('assets/burger.png'),
              ),
            ),
          )
        ],
      ),
    );

    menuItems.add(SizedBox(height: 15));
  }

  return Column(children: menuItems);
}
