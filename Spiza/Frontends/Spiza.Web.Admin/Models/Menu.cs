using System.Collections.Generic;

namespace Spiza.Web.Admin.Models
{
    public class Menu
    {
        public List<string> Categories { get; set; } = new();
        public List<Item> Items { get; set; } = new();
    }
}