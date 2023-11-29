package data

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
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
	Id          int64   `json:"id,omitempty"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Order       int32   `json:"order"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

var db *sql.DB

func init() {
	var err error
	connStr := "user=spiza dbname=spiza password=spiza host=localhost port=5432 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func GetRestaurants() ([]Restaurant, error) {
	var restaurants []Restaurant
	l := log.New(os.Stdout, "data-web", log.LstdFlags)

	rows, err := db.Query(`SELECT id, name FROM restaurants`)
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
		restaurants = []Restaurant{} // Ensure we're sending an empty slice, not nil
	}

	return restaurants, nil
}

func GetRestaurant(id int64) (Restaurant, error) {
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

	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
		return Restaurant{}, err
	}
	defer rows.Close()

	var restaurant Restaurant
	menuMap := make(map[string][]Item) // Temporary map to hold items by category

	for rows.Next() {
		var catName, itemName, itemDescription, itemImage string
		var itemOrder int32
		var itemPrice float64

		err := rows.Scan(&restaurant.Id, &restaurant.Name, &catName, &itemName, &itemOrder, &itemPrice, &itemDescription, &itemImage)
		if err != nil {
			log.Fatal(err)
			return restaurant, err
		}

		if itemName != "" { // If there is an item, add it to the map
			item := Item{
				Name:        itemName,
				Category:    catName,
				Order:       itemOrder,
				Price:       itemPrice,
				Description: itemDescription,
				Image:       itemImage,
			}
			menuMap[catName] = append(menuMap[catName], item)
		}
	}

	var menu Menu

	// Convert the map to a slice of categories and items
	for categoryName, items := range menuMap {
		menu.Categories = append(menu.Categories, categoryName)
		menu.Items = append(menu.Items, items...)
	}
	restaurant.Menu = menu

	return restaurant, nil
}

func CreateRestaurant(restaurant *Restaurant) error {
	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Rollback in case of error
	defer tx.Rollback()

	// Insert the restaurant
	var restaurantID int64
	err = tx.QueryRow("INSERT INTO restaurants (name) VALUES ($1) RETURNING id", restaurant.Name).Scan(&restaurantID)
	if err != nil {
		return err
	}

	// Insert each category and its items
	for _, category := range restaurant.Menu.Categories {
		var categoryID int64
		err = tx.QueryRow(`INSERT INTO menu_categories (name, restaurant_id) VALUES ($1, $2) RETURNING id`,
			category, restaurantID).Scan(&categoryID)
		if err != nil {
			return err
		}

		// Insert items in this category
		for _, item := range restaurant.Menu.Items {
			if item.Category == category {
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

func UpdateRestaurant(restaurant *Restaurant) error {
	tx, err := db.Begin()
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

func DeleteRestaurant(id int64) error {
	tx, err := db.Begin()
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

func CreateItem(item *Item) error {
	_, err := db.Exec(`INSERT INTO items (category_id, name, order_num, price, description, image) 
								VALUES ($1, $2, $3, $4, $5, $6)`,
		item.Category, item.Name, item.Order, item.Price, item.Description, item.Image)
	return err
}

func UpdateItem(item *Item) error {
	_, err := db.Exec(`UPDATE items SET name = $1, order_num = $2, price = $3, description = $4, image = $5 
								WHERE id = $6`,
		item.Name, item.Order, item.Price, item.Description, item.Image, item.Id)
	return err
}
