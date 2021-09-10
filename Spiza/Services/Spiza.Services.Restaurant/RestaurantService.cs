
using Grpc.Core;
using Spiza.Services.Restaurant.Repositories;

namespace Spiza.Services.Restaurant;
public class RestaurantService : Restaurant.RestaurantBase
{
    private readonly IRestaurantsRepository restaurantsRepo;

    public RestaurantService(IRestaurantsRepository restaurantsRepo)
    {
        this.restaurantsRepo = restaurantsRepo;
    }

    public override Task<RestaurantsResponse> GetRestaurants(RestaurantParameters request, ServerCallContext context)
    {
        return base.GetRestaurants(request, context);
    }
}
