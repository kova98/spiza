namespace Spiza.Services.Restaurant.Repositories;

using Spiza.Services.Restaurant.Entities;
using System;
using System.Collections.Generic;

public interface IRestaurantsRepository
{
    List<Restaurant> GetRestaurants();
    void CreateRestaurant(Restaurant restaurant);
    void DeleteRestaurant(Guid id);
    void UpdateRestaurant(Restaurant restaurant);
}
