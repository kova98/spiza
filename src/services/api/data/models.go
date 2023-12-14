package data

type Restaurant struct {
	Id             int64          `json:"id,omitempty"`
	Name           string         `json:"name"`
	MenuCategories []MenuCategory `json:"menu_categories" db:"menu_categories"`
}

type MenuCategory struct {
	Id           int64  `json:"id,omitempty"`
	Name         string `json:"name"`
	RestaurantId int64  `json:"restaurant_id" db:"restaurant_id"`
	Items        []Item `json:"items"`
}

type Item struct {
	Id           int64   `json:"id,omitempty"`
	CategoryId   int64   `json:"category_id" db:"category_id"`
	CategoryName string  `json:"category_name,omitempty" db:"category_name"`
	Name         string  `json:"name"`
	Order        int32   `json:"order" db:"order_num"`
	Price        float64 `json:"price"`
	Description  string  `json:"description"`
	Image        string  `json:"image"`
}

type Order struct {
	Id           int64   `json:"id"`
	UserId       int64   `json:"user_id" db:"user_id"`
	RestaurantId int64   `json:"restaurant_id" db:"restaurant_id"`
	Status       int64   `json:"status"`
	Items        []int64 `json:"items" `
}
