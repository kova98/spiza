package data

import "database/sql"

type ItemRepo struct {
	db *sql.DB
}

func (r *ItemRepo) CreateItem(item *Item) error {
	_, err := r.db.Exec(`INSERT INTO items (category_id, name, order_num, price, description, image) 
						 VALUES ($1, $2, $3, $4, $5, $6)`,
		item.Category, item.Name, item.Order, item.Price, item.Description, item.Image)
	return err
}

func (r *ItemRepo) UpdateItem(item *Item) error {
	_, err := r.db.Exec(`UPDATE items SET name = $1, order_num = $2, price = $3, description = $4, image = $5 
						 WHERE id = $6`,
		item.Name, item.Order, item.Price, item.Description, item.Image, item.Id)
	return err
}
