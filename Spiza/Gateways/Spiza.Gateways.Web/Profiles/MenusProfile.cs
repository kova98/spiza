using AutoMapper;
using GrpcServices.Restaurant;

namespace Spiza.Gateways.Web.Profiles
{
    public class MenusProfile : Profile
    {
        public MenusProfile()
        {
            CreateMap<Models.Menu, MenuMessage>();
            CreateMap<MenuMessage, Models.Menu>();
        }
    }
}
