namespace Spiza.Services.Restaurant.Settings
{
    internal class SpizaDbSettings : ISpizaDbSettings
    {
        public string ConnectionString => "mongodb://localhost:27017";

        public string DatabaseName => "SpizaDb";

        public string RestaurantsCollectionName => "Restaurants";
    }
}