<script>
  export let restaurants = [];
  import { onMount } from "svelte";

  const apiRoot = "http://127.0.0.1:5002/api/restaurant";

  onMount(async () => {
    const response = await fetch(apiRoot);
    restaurants = await response.json();
  });
</script>

<main>
  <h1>Restaurants</h1>
  {#each restaurants as restaurant}
    <div>
      <a href="/restaurant/{restaurant.id}">{restaurant.name}</a>
    </div>
  {/each}
</main>

<style>
</style>
