﻿namespace Spiza.Services.Restaurant.Repositories;

using Spiza.Services.Restaurant.Entities;

public interface IRestaurantsRepository
{
    List<Restaurant> GetRestaurants();
    void CreateRestaurant(Restaurant restaurant);
    void DeleteRestaurant(long id);
}
