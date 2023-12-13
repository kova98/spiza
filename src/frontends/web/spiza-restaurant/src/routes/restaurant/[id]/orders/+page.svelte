<script>
	import OrderCard from '$lib/components/order/order-card.svelte';
</script>

<h1 class="font-bold text-4xl mb-5">Orders</h1>
<div class="flex flex-row justify-between gap-3">
	<div class="w-full flex flex-col gap-3 items-center">
		<OrderCard />
		<OrderCard />
		<OrderCard />
	</div>
	<div class="w-full flex flex-col gap-3 items-center">
		<OrderCard />
		<OrderCard />
		<OrderCard />
	</div>
</div>
