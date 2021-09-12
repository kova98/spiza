using Spiza.Gateways.Web.Models;

namespace Spiza.Gateways.Web.Services
{
    public interface IRestaurantService
    {
        Task<IEnumerable<Restaurant>> GetRestaurants();
    }
}