import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/cart_bloc.dart';
import 'package:spiza_customer/bloc/cart_provider.dart';
import 'package:spiza_customer/bloc/restaurants_provider.dart';
import 'package:spiza_customer/models/item.dart';
import 'package:spiza_customer/models/restaurant.dart';

class MenuList extends StatelessWidget {
  final int restaurantId;
  const MenuList(this.restaurantId, {super.key});

  @override
  Widget build(BuildContext context) {
    final restaurantsBloc = RestaurantsProvider.of(context);
    final cartBloc = CartProvider.of(context);

    return StreamBuilder(
      stream: restaurantsBloc.restaurants,
      builder: (context, AsyncSnapshot<List<Restaurant>> snapshot) {
        if (!snapshot.hasData) {
          restaurantsBloc.getRestaurantWithMenu(restaurantId);
          return const Center(
            child: CircularProgressIndicator(),
          );
        } else {
          final restaurant = snapshot.data!.firstWhere(
            (r) => r.id == restaurantId,
            orElse: () => Restaurant.empty(),
          );
          // tu je bug, treba dohvatit snapshot.restaurant bla bla
          return ListView.builder(
            shrinkWrap: true,
            physics: const ClampingScrollPhysics(),
            itemCount: restaurant.menu.categories.length,
            itemBuilder: (context, int index) {
              final category = restaurant.menu.categories[index];
              final items = category.items;
              return items.isEmpty
                  ? const SizedBox.shrink()
                  : Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Padding(
                          padding: const EdgeInsets.symmetric(vertical: 12),
                          child: Text(
                            category.name,
                            style: const TextStyle(
                              fontWeight: FontWeight.bold,
                              fontSize: 24,
                            ),
                          ),
                        ),
                        _buildMenuItems(items, restaurant, cartBloc)
                      ],
                    );
            },
          );
        }
      },
    );
  }

  Widget _buildMenuItems(
      List<Item> items, Restaurant restaurant, CartBloc cartBloc) {
    List<Widget> menuItems = List<Widget>.empty(growable: true);

    for (var i = 0; i < items.length; i++) {
      final item = items[i];
      menuItems.add(
        InkWell(
          onTap: () => addToCart(item, restaurant, cartBloc),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      item.name,
                      style: const TextStyle(
                        fontWeight: FontWeight.bold,
                        fontSize: 18,
                      ),
                    ),
                    const SizedBox(height: 5),
                    Text(
                      item.description,
                      style: const TextStyle(color: Colors.black54),
                    ),
                    const SizedBox(height: 5),
                    Text(
                      "${item.price} kn",
                    )
                  ],
                ),
              ),
              ClipRRect(
                borderRadius: BorderRadius.circular(5),
                child: const SizedBox(
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
      menuItems.add(const SizedBox(height: 15));
    }
    return Column(children: menuItems);
  }

  void addToCart(Item item, Restaurant restaurant, CartBloc cartBloc) {
    cartBloc.addToCart(item);
    cartBloc.setRestaurant(restaurant);
  }
}
