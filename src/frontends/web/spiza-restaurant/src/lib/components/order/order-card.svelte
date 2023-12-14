<script lang="ts">
	import * as Card from '$lib/components/ui/card';

	export let order: Order;

	function formatDate(dateString: string) {
		const date = new Date(dateString);
		const secondsAgo = (Date.now() - date.getTime()) / 1000;

		return secondsAgo < 15
			? 'Just now'
			: secondsAgo < 60
			  ? Math.floor(secondsAgo) + ' seconds ago'
			  : secondsAgo < 3600
			    ? Math.floor(secondsAgo / 60) + ' minutes ago'
			    : secondsAgo < 86400
			      ? Math.floor(secondsAgo / 3600) + ' hours ago'
			      : secondsAgo < 604800
			        ? Math.floor(secondsAgo / 86400) + ' days ago'
			        : date.toLocaleDateString();
	}
</script>

<Card.Root class="w-full border-green-500 ">
	<Card.Header>
		<Card.Title>
			<p class="text-lg leading-none">Order #{order?.id}</p>
			<p class="text-sm font-normal text-muted-foreground">{formatDate(order?.date_created)}</p>
		</Card.Title>
	</Card.Header>
	<Card.Content>
		<ul class="text-sm">
			{#each order?.items ?? [] as item}
				<li>• {item.name}</li>
			{/each}
		</ul>
	</Card.Content>
</Card.Root>
