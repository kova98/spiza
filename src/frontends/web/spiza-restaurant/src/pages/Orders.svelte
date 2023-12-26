<script lang="ts">
  import OrderCard from "../lib/components/order/order-card.svelte";
  import { onMount } from "svelte";
  import { restaurantStore } from "../lib/stores";

  let orders: Order[] = [];

  let socket: WebSocket;
  onMount(async () => {
    const uri = `${import.meta.env.VITE_HTTP_ROOT}/api/restaurant/${$restaurantStore.id}/order`;
    const response = await fetch(uri);

    orders = await response.json();
    socket = new WebSocket(`${import.meta.env.VITE_WS_ROOT}/api/restaurant/${$restaurantStore.id}/order-ws`);
    socket.onopen = () => {
      console.log("Opened");
    };
    socket.onmessage = (e) => {
      let order = JSON.parse(e.data) as Order;
      orders = [order].concat(orders);
    };
  });

  let updateOrder = async (orderId: number, action: string) => {
    let msg = {
      action: action,
      id: orderId,
    };
    socket.send(JSON.stringify(msg));
  };
</script>

<h1 class="font-bold text-4xl mb-5">Orders</h1>
<div class="flex flex-row justify-between gap-3">
  <div class="w-full flex flex-col gap-3 items-center">
    {#each orders ?? [] as order}
      <OrderCard {order} {updateOrder} />
    {/each}
  </div>
  <div class="w-full flex flex-col gap-3 items-center"></div>
</div>
