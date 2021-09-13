
using Spiza.Gateways.Web.Services;
using GrpcServices.Restaurant;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();

builder.Services.AddScoped<IRestaurantService, RestaurantService>();
builder.Services.AddGrpcClient<Restaurant.RestaurantClient>((services, options) => 
{
    options.Address = new Uri("http://localhost:5201");
});

var app = builder.Build();

// Configure the HTTP request pipeline.
if (builder.Environment.IsDevelopment())
{
    app.UseDeveloperExceptionPage();
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
