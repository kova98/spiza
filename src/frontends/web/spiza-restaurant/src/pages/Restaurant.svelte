<script lang="ts">
  import { Button } from "../lib/components/ui/button";
  import { cn } from "../lib/utils";
  import { Exit } from "radix-icons-svelte";
  import { restaurantStore } from "../lib/stores";
  import { links, useLocation } from "svelte-routing";
  let location = useLocation();

  let paths = [
    { path: "/restaurant", name: "Overview", active: false },
    { path: "/restaurant/menu", name: "Menu", active: false },
    { path: "/restaurant/orders", name: "Orders", active: false },
  ];

  location.subscribe((value) => {
    paths.forEach((path) => {
      path.active = value.pathname === path.path;
    });
  });
</script>

<nav class={cn("flex flex-row items-center space-x-6 p-6 shadow-sm")} use:links>
  <h1 class="font-extrabold text-xl">{$restaurantStore.name}</h1>

  {#each paths as path}
    <a
      href={path.path}
      class="{path.active ? 'text-black' : 'text-gray-400'} font-medium text-sm transition-colors hover:text-"
      >{path.name}</a
    >
  {/each}
  <div class="flex-grow" />
  <Button class="ml-auto" variant="ghost" href="/"><Exit /></Button>
</nav>
<div class="p-6">
  <slot />
</div>
