interface Item {
	id: number;
	name: string;
	description: string;
	price: number;
}

interface MenuCategory {
	id: number;
	name: string;
	items: Item[];
}

interface Restaurant {
	id: number;
	name: string;
	menu_categories: MenuCategory[];
}

interface Order {
	id: number;
	date_created: string;
	date_updated: string;
	items: Item[];
}
