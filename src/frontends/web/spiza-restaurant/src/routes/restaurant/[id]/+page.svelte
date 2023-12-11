<script>
	export let data;
	export let restaurant = data.restaurant;

	let itemName = '';
	const apiRoot = 'http://127.0.0.1:5002/api';

	let itemDescription = '';
	let itemPrice = 0;

	const handleSubmitItem = async (categoryId) => {
		const item = { name: itemName, description: itemDescription, price: itemPrice };
		const response = await fetch(apiRoot + '/menu-category/' + categoryId + '/item', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(item)
		});

		if (!response.ok) {
			console.error('Error submitting item');
			return;
		}

		const itemResponse = await response.json();
		const category = restaurant.menu_categories.find((category) => category.id === categoryId);
		category.items.push(itemResponse);
		restaurant = restaurant;
		itemName = '';
		itemDescription = '';
		itemPrice = 0;
	};

	async function deleteItem(itemId) {
		const response = await fetch(`${apiRoot}/item/${itemId}`, {
			method: 'DELETE'
		});

		if (response.ok) {
			items = items.filter((item) => item.id !== itemId);
		} else {
			console.error('Error deleting item');
		}
	}
</script>

<h1>{restaurant.name}</h1>
<div>
	{#each restaurant.menu_categories ?? [] as category}
		<h2>{category.name}</h2>
		<ul>
			{#each category.items ?? [] as item}
				<li>
					<div>{item.name}</div>
					<div>{item.description}</div>
					<div>{item.price}</div>
					<button on:click={() => deleteItem(item.id)}>Delete</button>
				</li>
			{/each}
		</ul>
		<form on:submit|preventDefault={handleSubmitItem(category.id)}>
			<input type="text" bind:value={itemName} placeholder="Item Name" />
			<input type="text" bind:value={itemDescription} placeholder="Description" />
			<input type="number" bind:value={itemPrice} placeholder="Price" />
			<button type="submit">Create Item</button>
		</form>
	{/each}
</div>
