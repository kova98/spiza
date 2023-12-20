<script lang="ts">
	import OrderCard from '$lib/components/order/order-card.svelte';
	import { onMount } from 'svelte';

	export let data: { restaurant: Restaurant; socket: WebSocket };

	let orders: Order[] = [];

	let socket: WebSocket;
	onMount(async () => {
		const ordersRoute = 'http://127.0.0.1:5002/api/restaurant/' + data.restaurant.id + '/order';
		const response = await fetch(ordersRoute);

		orders = await response.json();
		socket = new WebSocket(
			'ws://127.0.0.1:5002/api/restaurant/' + data.restaurant.id + '/order-ws'
		);
		socket.onopen = () => {
			console.log('Opened');
			socket.send('Hello!');
		};
		socket.onmessage = (e) => {
			let order = JSON.parse(e.data) as Order;
			orders = [order].concat(orders);
		};
	});
</script>

<h1 class="font-bold text-4xl mb-5">Orders</h1>
<div class="flex flex-row justify-between gap-3">
	<div class="w-full flex flex-col gap-3 items-center">
		{#each orders as order}
			<OrderCard {order} />
		{/each}
	</div>
	<div class="w-full flex flex-col gap-3 items-center"></div>
</div>
