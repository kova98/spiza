package data

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
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

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) GetOrderDestinationLatLng(orderId int64) (string, error) {
	var latLng string
	sql := `SELECT a.lat_lng 
			FROM orders o 
			JOIN addresses a ON o.destination_id = a.id
			WHERE o.id = $1;`
	if err := r.db.Get(&latLng, sql, orderId); err != nil {
		return "", err
	}
	return latLng, nil
}

func (r *Repo) GetOrderRestaurantLocationLatLng(orderId int64) (string, error) {
	var latLng string
	sql := `SELECT a.lat_lng 
			FROM orders o 
			JOIN restaurants r ON o.restaurant_id = r.id
			JOIN addresses a ON r.address_id = a.id
			WHERE o.id = $1;`
	if err := r.db.Get(&latLng, sql, orderId); err != nil {
		return "", err
	}
	return latLng, nil
}
