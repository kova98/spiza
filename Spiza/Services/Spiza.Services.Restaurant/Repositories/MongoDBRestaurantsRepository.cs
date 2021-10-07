
using MongoDB.Driver;
using Spiza.Services.Restaurant.Entities;
using Spiza.Services.Restaurant.Settings;
using System;
using System.Collections.Generic;

namespace Spiza.Services.Restaurant.Repositories;
public class MongoDBRestaurantsRepository : IRestaurantsRepository
{
    private readonly IMongoCollection<Entities.Restaurant> restaurants;

    public MongoDBRestaurantsRepository(ISpizaDbSettings settings)
    {
        var client = new MongoClient(settings.ConnectionString);
        var database = client.GetDatabase(settings.DatabaseName);

        restaurants = database.GetCollection<Entities.Restaurant>(settings.RestaurantsCollectionName);
    }

    public void CreateRestaurant(Entities.Restaurant restaurant) => 
        restaurants.InsertOne(restaurant);

    public void DeleteRestaurant(Guid id) => 
        restaurants.DeleteOne(x => x.Id == id);

    public void EditRestaurant(Entities.Restaurant restaurant) =>
        restaurants.ReplaceOne(x => x.Id == restaurant.Id, restaurant);

    public List<Entities.Restaurant> GetRestaurants() => 
        restaurants.Find(x => true).ToList();
}
