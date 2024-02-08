<script lang="ts">
  import OrderCard from "../lib/components/order/order-card.svelte";
  import { onMount } from "svelte";
  import { restaurantStore } from "../lib/stores";

  let orders: Order[] = [];

  let socket: WebSocket;

  function setupWebSocket() {
    const uri = `${import.meta.env.VITE_WS_ROOT}/api/restaurant/${$restaurantStore.id}/order-ws`;
    socket = new WebSocket(uri);

    socket.onopen = () => {
      console.log("WebSocket connection opened.");
    };

    socket.onclose = (event) => {
      console.log("WebSocket connection closed.", event);
      // Check if the close was intentional (1000) or due to error/network loss
      if (event.code !== 1000) {
        setTimeout(setupWebSocket, 3000);
      }
    };

    socket.onmessage = (e) => {
      let order = JSON.parse(e.data);
      if (order.id > 0) {
        orders = [order].concat(orders);
      }
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
      socket.close();
    };
  }

  onMount(async () => {
    const uri = `${import.meta.env.VITE_HTTP_ROOT}/api/restaurant/${$restaurantStore.id}/order`;
    const response = await fetch(uri);

    orders = await response.json();
    setupWebSocket();
  });

  let updateOrder = async (orderId: number, action: string) => {
    let order = orders.find((o) => o.id == orderId);
    if (!order) {
      return;
    }

    let status = action == "accept" ? 1 : action == "reject" ? 2 : action == "ready" ? 3 : 0;
    order.status = status;
    orders = orders;

    let msg = {
      status: status,
      id: orderId,
    };
    socket.send(JSON.stringify(msg));
  };
</script>

<h1 class="font-bold text-4xl mb-5">Orders</h1>
<div class="flex flex-col sm:flex-row justify-between gap-3">
  <div class="w-full flex flex-col gap-3 items-center">
    {#each orders ?? [] as order}
      <OrderCard {order} {updateOrder} />
    {/each}
  </div>
  <div class="w-full flex flex-col gap-3 items-center"></div>
</div>
