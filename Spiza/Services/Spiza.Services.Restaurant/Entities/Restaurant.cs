
using MongoDB.Bson.Serialization.Attributes;

namespace Spiza.Services.Restaurant.Entities;
public class Restaurant
{
    [BsonId]
    public Guid Id { get; set; } 
    public string? Name { get; set; }
    public Menu Menu { get; set; }
}
