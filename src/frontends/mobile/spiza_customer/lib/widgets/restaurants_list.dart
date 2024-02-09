import 'package:flutter/material.dart';
import 'package:spiza_customer/bloc/restaurants_provider.dart';
import 'package:spiza_customer/models/restaurant.dart';
import 'package:spiza_customer/screens/menu_screen.dart';

class RestaurantsList extends StatelessWidget {
  const RestaurantsList({super.key});

  @override
  Widget build(context) {
    final bloc = RestaurantsProvider.of(context);
    return StreamBuilder(
      stream: bloc.restaurants,
      builder: (context, AsyncSnapshot<List<Restaurant>> snapshot) {
        if (!snapshot.hasData) {
          bloc.getRestaurants();
          return const Center(
            child: CircularProgressIndicator(),
          );
        } else {
          return snapshot.data!.isNotEmpty
              ? ListView.builder(
                  shrinkWrap: true,
                  physics: const ClampingScrollPhysics(),
                  itemCount: snapshot.data!.length,
                  itemBuilder: (context, int index) {
                    final item = snapshot.data![index];
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
                          Container(
                            height: 150.0,
                            alignment: Alignment.center,
                            decoration: const BoxDecoration(
                              borderRadius:
                                  BorderRadius.all(Radius.circular(7)),
                              image: DecorationImage(
                                  image: AssetImage('assets/burger.png'),
                                  fit: BoxFit.cover),
                            ),
                          ),
                          const SizedBox(height: 5),
                          Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              Text(
                                item.name,
                                style: const TextStyle(
                                  fontSize: 18,
                                  fontWeight: FontWeight.bold,
                                ),
                              ),
                              const Text(
                                '★ 4.5',
                                style: TextStyle(
                                  fontSize: 18,
                                  fontWeight: FontWeight.bold,
                                ),
                              ),
                            ],
                          ),
                          const Text(
                            '0,99€',
                            style: TextStyle(color: Colors.black87),
                          ),
                          const SizedBox(height: 17)
                        ],
                      ),
                    );
                  },
                )
              : const Center(
                  child: CircularProgressIndicator(),
                );
        }
      },
    );
  }
}
