using Spiza.Gateways.Web.Models;

namespace Spiza.Gateways.Web.Services
{
    public interface IRestaurantService
    {
        IEnumerable<Restaurant> GetRestaurants();
    }
}