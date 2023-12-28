<script lang="ts">
  import { Button } from "../ui/button";
  import * as Card from "../ui/card";

  export let order: Order;
  export let updateOrder: (orderId: number, status: string) => void;

  function formatDate(dateString: string) {
    const date = new Date(dateString);
    const secondsAgo = (new Date().getTime() - date.getTime()) / 1000;

    return secondsAgo < 15
      ? "Just now"
      : secondsAgo < 60
        ? Math.floor(secondsAgo) + " seconds ago"
        : secondsAgo < 3600
          ? Math.floor(secondsAgo / 60) + " minutes ago"
          : secondsAgo < 86400
            ? Math.floor(secondsAgo / 3600) + " hours ago"
            : secondsAgo < 604800
              ? Math.floor(secondsAgo / 86400) + " days ago"
              : date.toLocaleDateString();
  }

  function formatStatus(status: number) {
    switch (status) {
      case 0:
        return "Created";
      case 1:
        return "Accepted";
      case 2:
        return "Rejected";
      case 3:
        return "Ready";
      case 4:
        return "Picked up";
      case 5:
        return "Delivered";
      default:
        return "Unknown";
    }
  }
</script>

<Card.Root class="w-full">
  <Card.Header>
    <Card.Title class="flex flex-row justify-between">
      <div>
        <p class="text-lg leading-none">Order #{order?.id}</p>
        <p class="text-sm font-normal text-muted-foreground">{formatDate(order?.date_created)}</p>
      </div>
      {#if order.status != 0}
        <p class="text-sm font-bold">{formatStatus(order?.status)}</p>
      {/if}
    </Card.Title>
  </Card.Header>
  <Card.Content>
    <div class="flex flex-row justify-between">
      <div class="flex flex-col gap-1">
        <ul class="text-sm">
          {#each order?.items ?? [] as item}
            <li>• {item.name}</li>
          {/each}
        </ul>
      </div>
      <div class="flex flex-col gap-1">
        {#if order.status == 0}
          <Button variant="default" on:click={() => updateOrder(order?.id, "accept")}>Accept</Button>
          <Button variant="destructive" on:click={() => updateOrder(order?.id, "reject")}>Reject</Button>
        {/if}
        {#if order.status == 1}
          <Button variant="default" class="mt-auto" on:click={() => updateOrder(order?.id, "ready")}>Ready</Button>
        {/if}
      </div>
    </div>
  </Card.Content>
</Card.Root>
