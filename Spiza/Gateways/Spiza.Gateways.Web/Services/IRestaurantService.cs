using Spiza.Gateways.Web.Models;
using System;
using System.Collections.Generic;

namespace Spiza.Gateways.Web.Services
{
    public interface IRestaurantService
    {
        IEnumerable<Restaurant> GetRestaurants();
        void CreateRestaurant(Restaurant restaurant);
        void EditRestaurant(Restaurant restaurant);
        void DeleteRestaurant(Guid id);
        Restaurant GetRestaurant(Guid restaurantId);
    }
}