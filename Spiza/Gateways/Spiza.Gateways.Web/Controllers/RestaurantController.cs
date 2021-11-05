using Microsoft.AspNetCore.Mvc;
using Spiza.Gateways.Web.Models;
using Spiza.Gateways.Web.Services;
using System;

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

    [HttpGet("{id}")]
    public ActionResult<Restaurant[]> GetRestaurant(string id)
    {
        var idValid = Guid.TryParse(id, out Guid restaurantId);
        
        if (idValid == false)
        {
            return BadRequest();
        }

        var restaurant = restaurantService.GetRestaurant(restaurantId);
        
        if (restaurant == null)
        {
            return NotFound();
        }

        return Ok(restaurant);
    }

    [HttpPost("")]
    public ActionResult<Restaurant> CreateRestaurant([FromBody] Restaurant restaurant)
    {
        if (restaurant == null)
        {
            return BadRequest();
        }

        restaurantService.CreateRestaurant(restaurant);

        return Ok();
    }

    [HttpPut("")]
    public ActionResult<Restaurant> EditRestaurant([FromBody] Restaurant restaurant)
    {
        if (restaurant == null)
        {
            return BadRequest();
        }

        restaurantService.EditRestaurant(restaurant);

        return Ok();
    }

    [HttpDelete("{id}")]
    public ActionResult<Restaurant> DeleteRestaurant(Guid id)
    {
        restaurantService.DeleteRestaurant(id);
        return Ok();
    }
}
