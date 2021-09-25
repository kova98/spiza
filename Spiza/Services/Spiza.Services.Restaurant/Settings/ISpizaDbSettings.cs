namespace Spiza.Services.Restaurant.Settings
{
    public interface ISpizaDbSettings
    {
        string ConnectionString { get; }
        string DatabaseName { get; }
        string RestaurantsCollectionName { get; }
    }
}