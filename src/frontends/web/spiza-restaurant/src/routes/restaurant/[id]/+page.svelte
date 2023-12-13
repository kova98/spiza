<script lang="ts">
	export let data: { restaurant: Restaurant };
	export let restaurant = data.restaurant;

	const apiRoot = 'http://127.0.0.1:5002/api';

	let categoryName = '';
	let itemName = '';
	let itemDescription = '';
	let itemPrice = 0;

	const handleSubmitItem = async (categoryId: number) => {
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
		if (!category) {
			console.error('Error finding category');
			return;
		}
		category.items.push(itemResponse);
		restaurant = restaurant;
		itemName = '';
		itemDescription = '';
		itemPrice = 0;
	};

	const handleSubmitCategory = async () => {
		const category = { name: categoryName, restaurant_id: restaurant.id };
		const response = await fetch(apiRoot + '/menu-category', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(category)
		});

		if (!response.ok) {
			console.error('Error submitting category');
			return;
		}

		const categoryResponse = await response.json();
		restaurant.menu_categories.push(categoryResponse);
		restaurant = restaurant;
		categoryName = '';
	};

	const deleteCategory = async function deleteCategory(categoryId: number) {
		const response = await fetch(`${apiRoot}/menu-category/${categoryId}`, {
			method: 'DELETE'
		});

		if (!response.ok) {
			console.error('Error deleting category');
			return;
		}

		restaurant.menu_categories = restaurant.menu_categories.filter(
			(category) => category.id !== categoryId
		);
		restaurant = restaurant;
	};

	async function deleteItem(itemId: number) {
		const response = await fetch(`${apiRoot}/item/${itemId}`, {
			method: 'DELETE'
		});

		if (response.ok) {
			const category = restaurant.menu_categories.find((category) =>
				category.items.some((item) => item.id === itemId)
			);
			if (category) {
				category.items = category.items.filter((item) => item.id !== itemId);
				restaurant = restaurant;
			}
		} else {
			console.error('Error deleting item');
		}
	}
</script>

<h1>{restaurant.name}</h1>
<div>
	{#each restaurant.menu_categories ?? [] as category}
		<h2>{category.name}</h2>
		<button on:click={() => deleteCategory(category.id)}>Delete</button>
		<ul>
			{#each category.items ?? [] as item}
				<li>
					<div>{item.name}</div>
					<div>{item.description}</div>
					<div>{item.price}</div>
					<button on:click={() => deleteItem(item.id)}>Delete</button>
				</li>
			{/each}
			<li>
				<form on:submit|preventDefault={(e) => handleSubmitItem(category.id)}>
					<input type="text" bind:value={itemName} placeholder="Item Name" />
					<input type="text" bind:value={itemDescription} placeholder="Description" />
					<input type="number" bind:value={itemPrice} placeholder="Price" />
					<button type="submit">Create Item</button>
				</form>
			</li>
		</ul>
	{/each}
	<form on:submit|preventDefault={handleSubmitCategory}>
		<input type="text" bind:value={categoryName} placeholder="Category Name" />
		<button type="submit">Create Category</button>
	</form>
</div>
