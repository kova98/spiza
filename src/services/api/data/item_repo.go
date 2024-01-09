package data

import (
	"github.com/jmoiron/sqlx"
)

type ItemRepo struct {
	db *sqlx.DB
}

func NewItemRepo(db *sqlx.DB) *ItemRepo {
	return &ItemRepo{db}
}

func (r *ItemRepo) CreateItem(item *Item) (int64, error) {
	var itemId int64
	err := r.db.QueryRow(`
			INSERT INTO items (category_id, name, order_num, price, description, image) 
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id`,
		item.CategoryId, item.Name, item.Order, item.Price, item.Description, item.Image).Scan(&itemId)

	if err != nil {
		return 0, err
	}

	return itemId, err
}

func (r *ItemRepo) UpdateItem(item *Item) error {
	_, err := r.db.Exec(`UPDATE items SET name = $1, order_num = $2, price = $3, description = $4, image = $5 
						 WHERE id = $6`,
		item.Name, item.Order, item.Price, item.Description, item.Image, item.Id)
	return err
}

func (r *ItemRepo) DeleteItem(id int64) error {
	_, err := r.db.Exec(`DELETE FROM items WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ItemRepo) GetByOrder(orderId int64) ([]Item, error) {
	var items []Item
	sql := `SELECT i.*
			FROM orders o
			JOIN items i ON i.id = ANY(o.items)
			WHERE o.id = $1`
	if err := r.db.Select(&items, sql, orderId); err != nil {
		return nil, err
	}
	return items, nil
}
