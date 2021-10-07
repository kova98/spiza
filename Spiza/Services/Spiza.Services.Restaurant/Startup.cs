using GrpcServices.Restaurant;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Options;
using Spiza.Services.Restaurant.Repositories;
using Spiza.Services.Restaurant.Settings;
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
            services.Configure<SpizaDbSettings>(
                Configuration.GetSection(nameof(SpizaDbSettings)));
            services.AddSingleton<ISpizaDbSettings>(sp =>
                sp.GetRequiredService<IOptions<SpizaDbSettings>>().Value);

            services.AddSingleton<IRestaurantsRepository, MongoDBRestaurantsRepository>();
            
            services.AddGrpc();
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

