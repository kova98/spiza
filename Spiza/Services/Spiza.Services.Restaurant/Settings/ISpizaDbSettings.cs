namespace Spiza.Services.Restaurant.Settings
{
    public interface ISpizaDbSettings
    {
        string ConnectionString { get; set; }
        string DatabaseName { get; set; }
        string RestaurantsCollectionName { get; set; }
    }
}