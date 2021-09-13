using GrpcServices.Restaurant;

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

        return response.Data.Select(MapToRestaurantModel);
    }

    private Models.Restaurant MapToRestaurantModel(RestaurantResponse response)
    {
        return new Models.Restaurant
        {
            Id = Guid.Parse(response.Id),
            Name = response.Name
        };
    }
}
