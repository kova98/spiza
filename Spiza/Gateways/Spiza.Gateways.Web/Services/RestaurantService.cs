using AutoMapper;
using GrpcServices.Restaurant;

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
        var message = mapper.Map<RestaurantMessage>(restaurant);
        client.EditRestaurant(new EditRestaurantRequest { Restaurant = message });
    }

    public IEnumerable<Models.Restaurant> GetRestaurants()
    {
        var response = client.GetRestaurants(new());
        return mapper.Map<IEnumerable<Models.Restaurant>>(response.Restaurants);
    }
}
