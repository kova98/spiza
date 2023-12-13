package data

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

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

type RestaurantRepo struct {
	db *sqlx.DB
}

func NewRestaurantRepo(db *sqlx.DB) *RestaurantRepo {
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
	restaurantQuery := `SELECT id, name FROM restaurants WHERE id = $1`
	if err := repo.db.Get(&restaurant, restaurantQuery, id); err != nil {
		return nil, err
	}

	itemQuery := `
		SELECT i.id, i.name, i.category_id, i.order_num, i.price, i.description, i.image
		FROM items i
		LEFT JOIN menu_categories mc ON mc.id = i.category_id
		WHERE mc.restaurant_id = $1`
	var items []Item
	if err := repo.db.Select(&items, itemQuery, id); err != nil {
		return nil, err
	}

	categoryQuery := "SELECT id, name FROM menu_categories WHERE restaurant_id = $1"
	var categories []MenuCategory
	if err := repo.db.Select(&categories, categoryQuery, id); err != nil {
		return nil, err
	}

	for i, category := range categories {
		for _, item := range items {
			if item.CategoryId == category.Id {
				categories[i].Items = append(categories[i].Items, item)
			}
		}
	}

	restaurant.MenuCategories = categories
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
