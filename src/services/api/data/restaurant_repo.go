package data

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Restaurant struct {
	Id             int64          `json:"id,omitempty"`
	Name           string         `json:"name"`
	MenuCategories []MenuCategory `json:"menu_categories"`
}

type MenuCategory struct {
	Id           int64  `json:"id,omitempty"`
	Name         string `json:"name"`
	RestaurantId int64  `json:"restaurant_id"`
	Items        []Item `json:"items"`
}

type Item struct {
	Id          int64   `json:"id,omitempty"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Order       int32   `json:"order"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

type RestaurantRepo struct {
	db *sql.DB
}

func NewRestaurantRepo(db *sql.DB) *RestaurantRepo {
	return &RestaurantRepo{db}
}

func (r *RestaurantRepo) GetRestaurants() ([]Restaurant, error) {
	var restaurants []Restaurant
	l := log.New(os.Stdout, "data-web", log.LstdFlags)

	rows, err := r.db.Query(`SELECT id, name FROM restaurants`)
	if err != nil {
		l.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var restaurant Restaurant
		if err := rows.Scan(&restaurant.Id, &restaurant.Name); err != nil {
			l.Println(err)
			return nil, err
		}
		// You would need additional queries to populate the Menu and Items
		restaurants = append(restaurants, restaurant)
	}

	if err = rows.Err(); err != nil {
		l.Println(err)
		return nil, err
	}

	if restaurants == nil {
		restaurants = []Restaurant{}
	}

	return restaurants, nil
}

func (repo *RestaurantRepo) GetRestaurant(id int64) (*Restaurant, error) {
	var restaurant Restaurant

	// Query to get the restaurant
	restaurantQuery := `SELECT id, name FROM restaurants WHERE id = $1`
	err := repo.db.QueryRow(restaurantQuery, id).Scan(&restaurant.Id, &restaurant.Name)
	if err != nil {
		return nil, err
	}

	// Query to get the menu categories and items
	categoryQuery := `SELECT mc.id, mc.name, i.id, i.name, i.category_id, i.order_num, i.price, i.description, i.image 
                      FROM menu_categories mc 
                      LEFT JOIN items i ON mc.id = i.category_id 
                      WHERE mc.restaurant_id = $1`

	rows, err := repo.db.Query(categoryQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categoryMap := make(map[int64]*MenuCategory)
	for rows.Next() {
		var catId int64
		var catName string
		var item Item
		err := rows.Scan(&catId, &catName, &item.Id, &item.Name, &item.Category, &item.Order, &item.Price, &item.Description, &item.Image)
		if err != nil {
			return nil, err
		}

		if _, exists := categoryMap[catId]; !exists {
			categoryMap[catId] = &MenuCategory{Id: catId, Name: catName, RestaurantId: id}
		}

		if item.Id != 0 { // Check if item exists
			categoryMap[catId].Items = append(categoryMap[catId].Items, item)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Append the fully populated categories to restaurant.MenuCategories
	for _, category := range categoryMap {
		restaurant.MenuCategories = append(restaurant.MenuCategories, *category)
	}

	return &restaurant, nil
}

func (r *RestaurantRepo) CreateRestaurant(restaurant *Restaurant) (int64, error) {
	var restaurantID int64
	err := r.db.QueryRow("INSERT INTO restaurants (name) VALUES ($1) RETURNING id", restaurant.Name).Scan(&restaurantID)
	if err != nil {
		return 0, err
	}

	return restaurantID, nil
}

func (r *RestaurantRepo) UpdateRestaurant(restaurant *Restaurant) error {
	_, err := r.db.Exec("UPDATE restaurants SET name = $1 WHERE id = $2", restaurant.Name, restaurant.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *RestaurantRepo) DeleteRestaurant(id int64) error {
	_, err := r.db.Exec("DELETE FROM restaurants WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
