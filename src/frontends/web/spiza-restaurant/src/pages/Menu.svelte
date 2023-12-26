<script lang="ts">
  import { Button } from "../lib/components/ui/button";
  import * as Card from "../lib/components/ui/card";
  import { Input } from "../lib/components/ui/input";
  import * as Accordion from "../lib/components/ui/accordion";
  import { restaurantStore } from "../lib/stores";

  const apiRoot = `${import.meta.env.VITE_HTTP_ROOT}/api`;

  let categoryName = "";
  let itemName = "";
  let itemDescription = "";
  let itemPrice = 0;

  const handleSubmitItem = async (categoryId: number) => {
    const item = { name: itemName, description: itemDescription, price: itemPrice * 1 };
    const response = await fetch(apiRoot + "/menu-category/" + categoryId + "/item", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(item),
    });

    if (!response.ok) {
      console.error("Error submitting item");
      return;
    }

    const itemResponse = await response.json();
    const category = $restaurantStore.menu_categories.find((category) => category.id === categoryId);
    if (!category) {
      console.error("Error finding category");
      return;
    }
    category.items.push(itemResponse);
    $restaurantStore = $restaurantStore;
    itemName = "";
    itemDescription = "";
    itemPrice = 0;
  };

  const handleSubmitCategory = async () => {
    const category = { name: categoryName, restaurant_id: $restaurantStore.id };
    const response = await fetch(apiRoot + "/menu-category", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(category),
    });

    if (!response.ok) {
      console.error("Error submitting category");
      return;
    }

    const categoryResponse = await response.json();
    $restaurantStore.menu_categories.push(categoryResponse);
    $restaurantStore = $restaurantStore;
    categoryName = "";
  };

  const deleteCategory = async function deleteCategory(categoryId: number) {
    const response = await fetch(`${apiRoot}/menu-category/${categoryId}`, {
      method: "DELETE",
    });

    if (!response.ok) {
      console.error("Error deleting category");
      return;
    }

    $restaurantStore.menu_categories = $restaurantStore.menu_categories.filter((cat) => cat.id !== categoryId);
    $restaurantStore = $restaurantStore;
  };

  async function deleteItem(itemId: number) {
    const response = await fetch(`${apiRoot}/item/${itemId}`, {
      method: "DELETE",
    });

    if (response.ok) {
      const category = $restaurantStore.menu_categories.find((cat) => cat.items.some((item) => item.id === itemId));
      if (category) {
        category.items = category.items.filter((item) => item.id !== itemId);
        $restaurantStore = $restaurantStore;
      }
    } else {
      console.error("Error deleting item");
    }
  }
</script>

<h1 class="font-bold text-4xl mb-5">Menu</h1>
<div class="grid grid-cols-3 justify-between gap-3">
  {#each $restaurantStore.menu_categories ?? [] as category}
    <Card.Root>
      <Card.Header>
        <div class="flex flex-row justify-between">
          <Card.Title class="font-bold text-3xl">{category.name}</Card.Title>
          <Button variant="destructive" on:click={() => deleteCategory(category.id)}>Delete</Button>
        </div>
      </Card.Header>
      <Card.Content>
        <Accordion.Root class="w-full">
          {#each category.items ?? [] as item}
            <Accordion.Item value="item-{item.id}">
              <Accordion.Trigger>
                <span>{item.name} €{item.price}</span>
              </Accordion.Trigger>
              <Accordion.Content>
                <div class="flex flex-row">
                  {item.description}
                  <Button on:click={() => deleteItem(item.id)}>Delete</Button>
                </div>
              </Accordion.Content>
            </Accordion.Item>
          {/each}
          <Accordion.Item value="item-add">
            <Accordion.Trigger>Add new item</Accordion.Trigger>
            <Accordion.Content>
              <div class="flex flex-col gap-1 p-1">
                <Input type="text" id="name" bind:value={itemName} placeholder="Name" />
                <Input type="text" id="description" bind:value={itemDescription} placeholder="Description" />
                <Input type="number" id="price" bind:value={itemPrice} placeholder="Price" />
                <Button on:click={(_) => handleSubmitItem(category.id)}>Create</Button>
              </div>
            </Accordion.Content>
          </Accordion.Item>
        </Accordion.Root>
      </Card.Content>
    </Card.Root>
  {/each}
  <Card.Root>
    <Card.Header>
      <Input type="text" class="font-bold text-3xl py-6" bind:value={categoryName} placeholder="New Category" />
      <Button on:click={(_) => handleSubmitCategory()}>Create</Button>
    </Card.Header>
  </Card.Root>
</div>
