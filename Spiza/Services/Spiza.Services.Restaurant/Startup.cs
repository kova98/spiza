using GrpcServices.Restaurant;
using Microsoft.AspNetCore.Mvc;
using Spiza.Services.Restaurant.Repositories;
using System.Net;

namespace Spiza.Services.Restaurant
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            services.AddGrpc();
            services.AddTransient<IRestaurantsRepository, FakeRestaurantsRepository>();
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            app.UseRouting();
            app.UseEndpoints(endpoints =>
            {
                endpoints.MapGrpcService<RestaurantService>();
            });
        }
    }
}

