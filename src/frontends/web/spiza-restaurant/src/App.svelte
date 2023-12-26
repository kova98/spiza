<script lang="ts">
  import Home from "./pages/Home.svelte";
  import RestaurantPage from "./pages/Restaurant.svelte";
  import MenuPage from "./pages/Menu.svelte";
  import OrdersPage from "./pages/Orders.svelte";
  import { Router, Route, Link } from "svelte-routing";
  import { restaurantStore } from "./lib/stores";
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
