namespace Spiza.Services.Restaurant.Entities
{
    public class Menu
    {
        public List<string> Categories { get; set; } = new();
        public List<Item> Items { get; set; } = new();
    }
}