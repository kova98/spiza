using Spiza.Gateways.Web.Models;

namespace Spiza.Gateways.Web.Services
{
    public interface IRestaurantService
    {
        IEnumerable<Restaurant> GetRestaurants();
        void CreateRestaurant(Restaurant restaurant);
        void EditRestaurant(Restaurant restaurant);
        void DeleteRestaurant(Guid id);
    }
}