    namespace Spiza.Gateways.Web.Models
{
    public class Item
    {
        public string Name { get; set; } = "";
        public string Category { get; set; } = "";
        public int Order { get; set; }
        public double Price { get; set; }
        public string Description { get; set; } = "";
        public string Image { get; set; } = "";
    }
}