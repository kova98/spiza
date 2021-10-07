
using Google.Protobuf.Collections;
using Grpc.Core;
using Spiza.Services.Restaurant.Repositories;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Entities = Spiza.Services.Restaurant.Entities;

namespace GrpcServices.Restaurant;

public class RestaurantService : Restaurant.RestaurantBase
{
    private readonly IRestaurantsRepository restaurantsRepo;

    public RestaurantService(IRestaurantsRepository restaurantsRepo)
    {
        this.restaurantsRepo = restaurantsRepo;
    }

    public override async Task<GetRestaurantsResponse> GetRestaurants(GetRestaurantsRequest request, ServerCallContext context)
    {
        var restaurants = restaurantsRepo.GetRestaurants();
        var response = new GetRestaurantsResponse();
        restaurants.ForEach(x => response.Restaurants.Add(new RestaurantMessage
        {
            Id = x.Id.ToString(),
            Name = x.Name
        }));

        return response;
    }

    public override async Task<CreateRestaurantResponse> CreateRestaurant(CreateRestaurantRequest request, ServerCallContext context)
    {
        var restaurant = new Entities.Restaurant { Name = request.Name };
        restaurantsRepo.CreateRestaurant(restaurant);

        return new();
    }

    public override async Task<DeleteRestaurantResponse> DeleteRestaurant(DeleteRestaurantRequest request, ServerCallContext context)
    {
        restaurantsRepo.DeleteRestaurant(Guid.Parse(request.Id));

        return new();
    }

    public override async Task<UpdateRestaurantResponse> UpdateRestaurant(UpdateRestaurantRequest request, ServerCallContext context)
    {
        restaurantsRepo.UpdateRestaurant(MapToEntity(request));
        return new();
    }

    public override async Task<UpdateMenuResponse> UpdateMenu(UpdateMenuRequest request, ServerCallContext context)
    {
        var restaurantId = Guid.Parse(request.RestaurantId);
        var menu = MapToEntity(request.Menu);

        restaurantsRepo.UpdateMenu(restaurantId, menu);

        return new();
    }

    private Entities.Menu MapToEntity(MenuMessage menu) => new Entities.Menu
    {
        Categories = menu.Categories,
        Items = MapToEntity(menu.Items)
    };

    private IList<Entities.Item> MapToEntity(RepeatedField<ItemMessage> items)
    {
        var itemEntities = new List<Entities.Item>();

        foreach(var item in items)
        {
            itemEntities.Add(new Entities.Item
            {
                Category = item.Category,
                Name = item.Name,
                Order = item.Order,
                Price = item.Price,
            });
        }

        return itemEntities;
    }

    public static Entities.Restaurant MapToEntity(UpdateRestaurantRequest request)
    {
        return new Entities.Restaurant
        {
            Id = Guid.Parse(request.Restaurant.Id),
            Name = request.Restaurant.Name,
        };
    }
}
