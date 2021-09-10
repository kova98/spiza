namespace Spiza.Services.Restaurant.Repositories;

using Models;

public interface IRestaurantsRepository
{
    public List<Restaurant> GetRestaurants();
}
