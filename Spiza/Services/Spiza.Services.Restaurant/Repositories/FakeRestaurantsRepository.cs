
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

    public void CreateRestaurant(Restaurant restaurant) => restaurants.Add(restaurant);

    public void DeleteRestaurant(long id)
    {
        var restaurantToRemove = restaurants.FirstOrDefault(x => x.Id == id);
        if (restaurantToRemove != null)
        {
            restaurants.Remove(restaurantToRemove);
        }
    }

    public void EditRestaurant(Restaurant restaurant)
    {
        var restaurantToEdit = restaurants.FirstOrDefault(x => x.Id == restaurant.Id);

        if (restaurantToEdit == null) return;

        restaurantToEdit.Name = restaurant.Name;
    }

    public List<Restaurant> GetRestaurants() => restaurants;
}
