<script lang="ts">
  import { Button } from "../lib/components/ui/button";
  import { cn } from "../lib/utils";
  import { Exit } from "radix-icons-svelte";
  import { restaurantStore } from "../lib/stores";
  import { links } from "svelte-routing";
</script>

<nav class={cn("flex flex-row items-center space-x-6 p-6 shadow-sm")} use:links>
  <h1 class="font-bold text-xl">{$restaurantStore.name}</h1>
  <a href="/restaurant" class="text-sm font-medium transition-colors hover:text-primary"> Overview </a>

  <a href="/restaurant/menu" class="text-sm font-medium transition-colors hover:text-primary"> Menu </a>
  <a href="/restaurant/orders" class="text-sm font-medium transition-colors hover:text-primary"> Orders </a>
  <Button class="ml-auto" variant="ghost" href="/"><Exit /></Button>
</nav>
<div class="p-6">
  <slot />
</div>
