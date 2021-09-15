namespace Spiza.Services.Restaurant.Repositories;

using Spiza.Services.Restaurant.Entities;

public interface IRestaurantsRepository
{
    public List<Restaurant> GetRestaurants();
    public void CreateRestaurant(Restaurant restaurant);
}
