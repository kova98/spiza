
using MongoDB.Bson.Serialization.Attributes;
using System;

namespace Spiza.Services.Restaurant.Entities;
public class Restaurant
{
    [BsonId]
    public Guid Id { get; set; } 
    public string? Name { get; set; }
    public Menu Menu { get; set; } = new();
}
