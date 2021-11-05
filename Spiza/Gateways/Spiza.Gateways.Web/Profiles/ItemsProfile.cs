using AutoMapper;
using GrpcServices.Restaurant;

namespace Spiza.Gateways.Web.Profiles
{
    public class ItemsProfile : Profile
    {
        public ItemsProfile()
        {
            CreateMap<Models.Item, ItemMessage>();
            CreateMap<ItemMessage, Models.Item>();
        }
    }
}
