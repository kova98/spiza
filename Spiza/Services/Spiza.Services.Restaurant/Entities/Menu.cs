using System.Collections.Generic;

namespace Spiza.Services.Restaurant.Entities
{
    public class Menu
    {
        public IList<string> Categories { get; set; } = new List<string>();
        public IList<Item> Items { get; set; } = new List<Item>();
    }
}