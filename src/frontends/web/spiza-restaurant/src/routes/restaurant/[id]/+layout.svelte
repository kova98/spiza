<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { cn } from '$lib/utils';
	import { Exit } from 'radix-icons-svelte';

	let className: string | undefined | null = undefined;
	export { className as class };

	export let data: { restaurant: Restaurant };
	export let restaurant = data.restaurant;
</script>

<nav class={cn('flex flex-row items-center space-x-6 p-6 shadow-sm', className)}>
	<h1 class="font-bold text-xl">{restaurant.name}</h1>
	<a
		href="/restaurant/{restaurant.id}"
		class="text-sm font-medium transition-colors hover:text-primary"
	>
		Overview
	</a>

	<a
		href="/restaurant/{restaurant.id}/menu"
		class="text-sm font-medium transition-colors hover:text-primary"
	>
		Menu
	</a>
	<Button class="ml-auto" variant="ghost" href="/"><Exit /></Button>
</nav>
<div class="p-6">
	<slot />
</div>
