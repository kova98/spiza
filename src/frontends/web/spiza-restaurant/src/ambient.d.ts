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
  dateCreated: string;
  dateUpdated: string;
  items: Item[];
  status: number;
}

enum OrderStatus {
  Created,
  Accepted,
  Rejected,
  Ready,
  PickedUp,
  Delivered,
}
