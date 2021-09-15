
using Grpc.Core;
using Spiza.Services.Restaurant.Repositories;
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
            Id = x.Id,
            Name = x.Name
        }));

        return response;
    }
}
