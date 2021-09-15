
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

    public override async Task<CreateRestaurantResponse> CreateRestaurant(CreateRestaurantRequest request, ServerCallContext context)
    {
        var restaurant = new Entities.Restaurant { Name = request.Name };
        restaurantsRepo.CreateRestaurant(restaurant);

        return new();
    }

    public override async Task<DeleteRestaurantResponse> DeleteRestaurant(DeleteRestaurantRequest request, ServerCallContext context)
    {
        restaurantsRepo.DeleteRestaurant(request.Id);

        return new();
    }

    public override async Task<EditRestaurantResponse> EditRestaurant(EditRestaurantRequest request, ServerCallContext context)
    {
        restaurantsRepo.EditRestaurant(MapToEntity(request));
        return new();
    }

    public static Entities.Restaurant MapToEntity(EditRestaurantRequest request)
    {
        return new Entities.Restaurant
        {
            Id = request.Restaurant.Id,
            Name = request.Restaurant.Name
        };
    }
        
}
