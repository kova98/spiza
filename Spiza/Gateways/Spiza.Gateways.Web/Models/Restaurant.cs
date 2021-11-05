using System;

namespace Spiza.Gateways.Web.Models;

public class Restaurant
{
    public Guid? Id { get; set; }
    public string? Name { get; set; }
    public Menu Menu { get; set; } = new();
}
