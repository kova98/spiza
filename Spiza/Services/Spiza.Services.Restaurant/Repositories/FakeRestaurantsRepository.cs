
namespace Spiza.Services.Restaurant.Repositories;

using Entities;

public class FakeRestaurantsRepository : IRestaurantsRepository
{
    private List<Restaurant> restaurants = new();

    public FakeRestaurantsRepository()
    {
        restaurants.AddRange(new Restaurant[]
        {
            new Restaurant { Id = 1, Name = "Restaurant 1" },
            new Restaurant { Id = 2, Name = "Restaurant 2" },
            new Restaurant { Id = 3, Name = "Restaurant 3" },
            new Restaurant { Id = 4, Name = "Restaurant 4" },
            new Restaurant { Id = 5, Name = "Restaurant 5" },
        });
    }

    public List<Restaurant> GetRestaurants() => restaurants;
}
