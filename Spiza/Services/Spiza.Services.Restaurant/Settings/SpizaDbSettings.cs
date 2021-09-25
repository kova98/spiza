namespace Spiza.Services.Restaurant.Settings
{
    internal class SpizaDbSettings : ISpizaDbSettings
    {
        public string ConnectionString { get; set; }

        public string DatabaseName { get; set; }

        public string RestaurantsCollectionName { get; set; }
    }
}