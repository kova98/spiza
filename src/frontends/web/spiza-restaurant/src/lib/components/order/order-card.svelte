<script lang="ts">
  import { Button } from "../ui/button";
  import * as Card from "../ui/card";

  export let order: Order;
  export let updateOrder: (orderId: number, action: string) => void;

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
</script>

<Card.Root class="w-full">
  <Card.Header>
    <Card.Title>
      <p class="text-lg leading-none">Order #{order?.id}</p>
      <p class="text-sm font-normal text-muted-foreground">{formatDate(order?.date_created)}</p>
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
        <Button variant="default" on:click={() => updateOrder(order?.id, "accept")}>Accept</Button>
        <Button variant="default" on:click={() => updateOrder(order?.id, "refuse")}>Refuse</Button>
        <p class="text-sm font-bold">{order?.status}</p>
      </div>
    </div>
  </Card.Content>
</Card.Root>
