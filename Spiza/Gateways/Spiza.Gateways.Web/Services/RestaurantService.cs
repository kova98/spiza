﻿using GrpcServices.Restaurant;

namespace Spiza.Gateways.Web.Services;
public class RestaurantService : IRestaurantService
{
    private readonly Restaurant.RestaurantClient client;

    public RestaurantService(Restaurant.RestaurantClient client)
    {
        this.client = client;
    }

    public IEnumerable<Models.Restaurant> GetRestaurants()
    {
        var response = client.GetRestaurants(new());

        return response.Restaurants.Select(MapToRestaurantModel);
    }

    private Models.Restaurant MapToRestaurantModel(RestaurantMessage restaurant)
    {
        return new Models.Restaurant
        {
            Id = restaurant.Id,
            Name = restaurant.Name
        };
    }
}
