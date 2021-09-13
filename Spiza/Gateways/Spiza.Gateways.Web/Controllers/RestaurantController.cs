using Microsoft.AspNetCore.Mvc;
using Spiza.Gateways.Web.Models;
using Spiza.Gateways.Web.Services;

namespace Spiza.Gateways.Web.Controllers;

[Route("api/[controller]")]
public class RestaurantController : Controller
{
    private readonly IRestaurantService restaurantService;

    public RestaurantController(IRestaurantService restaurantService)
    {
        this.restaurantService = restaurantService;
    }

    [HttpGet("")]
    public ActionResult<Restaurant[]> GetRestaurants()
    {
        return Ok(restaurantService.GetRestaurants());  
    }
}
