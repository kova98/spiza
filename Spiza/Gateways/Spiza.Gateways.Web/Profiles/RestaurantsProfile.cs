
using AutoMapper;
using GrpcServices.Restaurant;
using Spiza.Gateways.Web.Models;

namespace Spiza.Gateways.Web.Profiles;
public class RestaurantsProfile : Profile
{
    public RestaurantsProfile()
    {
        // Source --> Target
        CreateMap<Models.Restaurant, RestaurantMessage>();
        CreateMap<RestaurantMessage, Models.Restaurant>();
    }
}
