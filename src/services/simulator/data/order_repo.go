package data

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

type Order struct {
	Id                int64     ``
	UserId            int64     `db:"user_id"`
	RestaurantId      int64     `db:"restaurant_id"`
	DestinationId     int64     `db:"destination_id"`
	Status            int64     ``
	DateCreated       time.Time `db:"date_created"`
	DestinationLatLng string    `db:"lat_lng"`
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db}
}

func (r *OrderRepo) GetOrderDestinationLatLng(orderId int64) (string, error) {
	var order string
	sql := `SELECT a.lat_lng 
			FROM orders o 
			JOIN addresses a ON o.destination_id = a.id
			WHERE o.id = $1;`
	if err := r.db.Get(&order, sql, orderId); err != nil {
		return "", err
	}
	return order, nil
}
