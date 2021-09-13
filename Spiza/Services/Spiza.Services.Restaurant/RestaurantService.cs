
using Grpc.Core;
using Spiza.Services.Restaurant.Repositories;
using GrpcServices.Restaurant;

namespace GrpcServices.Restaurant;

public class RestaurantService : Restaurant.RestaurantBase
{
    private readonly IRestaurantsRepository restaurantsRepo;

    public RestaurantService(IRestaurantsRepository restaurantsRepo)
    {
        this.restaurantsRepo = restaurantsRepo;
    }

    public override async Task<RestaurantsResponse> GetRestaurants(RestaurantParameters request, ServerCallContext context)
    {
        var restaurants = restaurantsRepo.GetRestaurants();
        var response = new RestaurantsResponse();
        restaurants.ForEach(x => response.Data.Add(new RestaurantResponse
        {
            Id = x.Id.ToString(),
            Name = x.Name
        }));

        return response;
    }
}
