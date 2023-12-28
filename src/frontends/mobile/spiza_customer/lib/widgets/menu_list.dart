import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/cart_bloc.dart';
import 'package:spiza_customer/bloc/cart_provider.dart';
import 'package:spiza_customer/bloc/restaurants_provider.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/models/restaurant.dart';

class MenuList extends StatelessWidget {
  final int restaurantId;
  const MenuList(this.restaurantId);

  @override
  Widget build(BuildContext context) {
    final restaurantsBloc = RestaurantsProvider.of(context);
    final cartBloc = CartProvider.of(context);

    return StreamBuilder(
      stream: restaurantsBloc.restaurants,
      builder: (context, AsyncSnapshot<List<Restaurant>> snapshot) {
        if (!snapshot.hasData) {
          restaurantsBloc.getRestaurantWithMenu(restaurantId);
          return Center(
            child: CircularProgressIndicator(),
          );
        } else {
          final restaurant =
              snapshot.data!.firstWhere((r) => r.id == restaurantId);
          // tu je bug, treba dohvatit snapshot.restaurant bla bla
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
                        _buildMenuItems(items, restaurant.name, cartBloc)
                      ],
                    );
            },
          );
        }
      },
    );
  }

  Widget _buildMenuItems(
      List<Item> items, String restaurantName, CartBloc cartBloc) {
    List<Widget> menuItems = List<Widget>.empty();

    for (var i = 0; i < items.length; i++) {
      final item = items[i];
      menuItems.add(
        InkWell(
          onTap: () => addToCart(item, restaurantName, cartBloc),
          child: Row(
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
        ),
      );
      menuItems.add(SizedBox(height: 15));
    }
    return Column(children: menuItems);
  }

  void addToCart(Item item, String restaurantName, CartBloc cartBloc) {
    cartBloc.addToCart(item);
    cartBloc.setRestaurantName(restaurantName);
  }
}
