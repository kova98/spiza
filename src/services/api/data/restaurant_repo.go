package data

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Restaurant struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
	Menu Menu   `json:"menu"`
}

type Menu struct {
	Categories []string `json:"categories"`
	Items      []Item   `json:"items"`
}

type Item struct {
	Id          sql.NullInt64   `json:"id,omitempty"`
	Name        sql.NullString  `json:"name"`
	Category    sql.NullString  `json:"category"`
	Order       sql.NullInt32   `json:"order"`
	Price       sql.NullFloat64 `json:"price"`
	Description sql.NullString  `json:"description"`
	Image       sql.NullString  `json:"image"`
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

func (r *RestaurantRepo) GetRestaurant(id int64) (Restaurant, error) {
	query := `SELECT r.id as restaurant_id, 
					 r.name as restaurant_name, 
					 mc.name as category_name, 
					 i.name as item_name, 
					 i.order_num, 
					 i.price, 
					 i.description, 
					 i.image
			  FROM restaurants r
			  LEFT JOIN menu_categories mc ON r.id = mc.restaurant_id
			  LEFT JOIN items i ON mc.id = i.category_id
			  WHERE r.id = $1`

	rows, err := r.db.Query(query, id)
	if err != nil {
		log.Fatal(err)
		return Restaurant{}, err
	}
	defer rows.Close()

	var restaurant Restaurant
	menuMap := make(map[string][]Item) // Temporary map to hold items by category

	for rows.Next() {
		var item Item
		var catName sql.NullString

		err := rows.Scan(&restaurant.Id, &restaurant.Name, &catName, &item.Name, &item.Order, &item.Price, &item.Description, &item.Image)
		if err != nil {
			log.Fatal(err)
			return restaurant, err
		}

		if item.Name.Valid && item.Name.String != "" {
			if catName.Valid {
				item.Category = catName
			}
			menuMap[item.Category.String] = append(menuMap[item.Category.String], item)
		}
	}

	var menu Menu

	// Convert the map to a slice of categories and items
	for categoryName, items := range menuMap {
		if categoryName != "" {
			menu.Categories = append(menu.Categories, categoryName)
		}
		menu.Items = append(menu.Items, items...)
	}
	restaurant.Menu = menu

	return restaurant, nil
}

func (r *RestaurantRepo) CreateRestaurant(restaurant *Restaurant) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var restaurantID int64
	err = tx.QueryRow("INSERT INTO restaurants (name) VALUES ($1) RETURNING id", restaurant.Name).Scan(&restaurantID)
	if err != nil {
		return err
	}

	for _, category := range restaurant.Menu.Categories {
		var categoryID int64
		err = tx.QueryRow(`INSERT INTO menu_categories (name, restaurant_id) VALUES ($1, $2) RETURNING id`,
			category, restaurantID).Scan(&categoryID)
		if err != nil {
			return err
		}

		for _, item := range restaurant.Menu.Items {
			if item.Category.Valid && item.Category.String == category {
				_, err = tx.Exec(`INSERT INTO items (category_id, name, order_num, price, description, image) 
										VALUES ($1, $2, $3, $4, $5, $6)`,
					categoryID, item.Name, item.Order, item.Price, item.Description, item.Image)
				if err != nil {
					return err
				}
			}
		}
	}

	return tx.Commit()
}

func (r *RestaurantRepo) UpdateRestaurant(restaurant *Restaurant) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE restaurants SET name = $1 WHERE id = $2", restaurant.Name, restaurant.Id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *RestaurantRepo) DeleteRestaurant(id int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM restaurants WHERE id = $1", id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
