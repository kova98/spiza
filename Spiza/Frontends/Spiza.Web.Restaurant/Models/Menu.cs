using System.Collections.Generic;

namespace Spiza.Web.Restaurant.Models;

public class Menu
{
    public List<Category> Categories { get; set; } = new();
}
