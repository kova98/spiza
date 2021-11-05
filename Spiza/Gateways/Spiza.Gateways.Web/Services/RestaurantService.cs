using AutoMapper;
using GrpcServices.Restaurant;
using System;
using System.Collections.Generic;

namespace Spiza.Gateways.Web.Services;
public class RestaurantService : IRestaurantService
{
    private readonly Restaurant.RestaurantClient client;
    private readonly IMapper mapper;

    public RestaurantService(Restaurant.RestaurantClient client, IMapper mapper)
    {
        this.client = client;
        this.mapper = mapper;
    }

    public void CreateRestaurant(Models.Restaurant restaurant)
    {
        client.CreateRestaurant(new CreateRestaurantRequest { Name = restaurant.Name });
    }

    public void DeleteRestaurant(Guid id)
    {
        client.DeleteRestaurant(new DeleteRestaurantRequest { Id = id.ToString() });
    }

    public void EditRestaurant(Models.Restaurant restaurant)
    {
        var restaurantMessage = new RestaurantMessage
        {
            Id = restaurant.Id.ToString(),
            Name = restaurant.Name,
            Menu = mapper.Map<MenuMessage>(restaurant.Menu) ?? new()
        };

        client.UpdateRestaurant(new UpdateRestaurantRequest { Restaurant = restaurantMessage });
    }

    public Models.Restaurant GetRestaurant(Guid restaurantId)
    {
        var response = client.GetRestaurant(new GetRestaurantRequest { Id = restaurantId.ToString() });
        return mapper.Map<Models.Restaurant>(response.Restaurant);
    }

    public IEnumerable<Models.Restaurant> GetRestaurants()
    {
        var response = client.GetRestaurants(new());
        return mapper.Map<IEnumerable<Models.Restaurant>>(response.Restaurants);
    }
}
