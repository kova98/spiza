
using AutoMapper;
using GrpcServices.Restaurant;
using Spiza.Gateways.Web.Models;

namespace Spiza.Gateways.Web.Profiles;
public class RestaurantsProfile : Profile
{
    public RestaurantsProfile()
    {
        // Source --> Target
        CreateMap<Models.Restaurant, RestaurantMessage>()
            // Grpc auto-generated objects' Lists are read only. Manual mapping required.
            .ForPath(x => x.Menu.Items, o => o.Ignore());

        CreateMap<RestaurantMessage, Models.Restaurant>();
    }
}
