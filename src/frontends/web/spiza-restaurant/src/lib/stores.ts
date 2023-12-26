import { writable, type Writable } from "svelte/store";

const defaultRestaurant: Restaurant = {
  id: 0,
  name: "",
  menu_categories: [],
};

export const restaurantStore = writable<Restaurant>(defaultRestaurant);
