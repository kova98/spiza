
namespace Spiza.Services.Restaurant.Repositories;

using Models;

public class FakeRestaurantsRepository : IRestaurantsRepository
{
    private List<Restaurant> restaurants = new();

    public FakeRestaurantsRepository()
    {
        restaurants.AddRange(new Restaurant[]
        {
            new Restaurant { Id = Guid.NewGuid(), Name = "Restaurant 1" },
            new Restaurant { Id = Guid.NewGuid(), Name = "Restaurant 2" },
            new Restaurant { Id = Guid.NewGuid(), Name = "Restaurant 3" },
            new Restaurant { Id = Guid.NewGuid(), Name = "Restaurant 4" },
            new Restaurant { Id = Guid.NewGuid(), Name = "Restaurant 5" },
        });
    }

    public List<Restaurant> GetRestaurants() => restaurants;
}
