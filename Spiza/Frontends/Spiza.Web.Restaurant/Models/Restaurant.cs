using System;

namespace Spiza.Web.Restaurant.Models;

public class Restaurant
{
    public Guid Id { get; set; }
    public string? Name { get; set; }
    public Menu Menu { get; set; }
} 

