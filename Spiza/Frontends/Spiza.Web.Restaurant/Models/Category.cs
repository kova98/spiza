using System;
using System.Collections.Generic;

namespace Spiza.Web.Restaurant.Models;
  
public class Category
{
    public Guid Id { get; set; }
    public string? Name { get; set; }
    public List<Item> Items { get; set; } = new();
}