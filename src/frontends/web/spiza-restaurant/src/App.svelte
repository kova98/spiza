<script lang="ts">
  import Home from "./pages/Home.svelte";
  import RestaurantPage from "./pages/Restaurant.svelte";
  import MenuPage from "./pages/Menu.svelte";
  import OrdersPage from "./pages/Orders.svelte";
  import { Router, Route, Link } from "svelte-routing";
  import { restaurantStore } from "./lib/stores";

  let restaurant: Restaurant;
  restaurantStore.subscribe((value) => {
    restaurant = value;
  });

  // let restaurant: Restaurant = {
  //   name: "test rest",
  //   id: 1,
  //   menu_categories: [
  //     {
  //       id: 1,
  //       name: "test cat",
  //       items: [
  //         {
  //           id: 1,
  //           name: "test item",
  //           description: "test desc",
  //           price: 1,
  //         },
  //       ],
  //     },
  //   ],
  // };
</script>

<main>
  <Router>
    <div>
      <Route path="/"><Home /></Route>
      <Route path="/restaurant"><RestaurantPage /></Route>
      <Route path="/restaurant/menu">
        <RestaurantPage>
          <MenuPage />
        </RestaurantPage>
      </Route>
      <Route path="/restaurant/orders">
        <RestaurantPage>
          <OrdersPage />
        </RestaurantPage>
      </Route>
    </div>
  </Router>
</main>
