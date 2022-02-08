package data

type Restaurant struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Menu Menu   `json:"menu"`
}

type Menu struct {
	Categories []string `json:"categories"`
	Items      []Item   `json:"items"`
}

type Item struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Order       int32   `json:"order"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

func GetRestaurants() []*Restaurant {
	return []*Restaurant{
		{
			Id:   "test id",
			Name: "test name",
			Menu: Menu{
				Categories: []string{"test category"},
				Items: []Item{
					{
						Name:     "item name",
						Category: "item category",
						Order:    4,
						Price:    12.35,
					},
				},
			},
		},
	}
}
