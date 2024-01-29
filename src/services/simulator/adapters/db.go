package adapters

import (
	"github.com/kova98/spiza/services/simulator/domain"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDb struct {
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

func NewPostgresDb(connStr string) *PostgresDb {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &PostgresDb{db}
}

func (r *PostgresDb) GetOrderDestinationLatLng(orderId int64) (domain.Location, error) {
	var latLng string
	sql := `SELECT a.lat_lng 
			FROM orders o 
			JOIN addresses a ON o.destination_id = a.id
			WHERE o.id = $1;`
	if err := r.db.Get(&latLng, sql, orderId); err != nil {
		return domain.Location{}, err
	}
	return domain.LatLngToLocation(latLng), nil
}

func (r *PostgresDb) GetOrderRestaurantLocation(orderId int64) (domain.Location, error) {
	var latLng string
	sql := `SELECT a.lat_lng 
			FROM orders o 
			JOIN restaurants r ON o.restaurant_id = r.id
			JOIN addresses a ON r.address_id = a.id
			WHERE o.id = $1;`
	if err := r.db.Get(&latLng, sql, orderId); err != nil {
		return domain.Location{}, err
	}
	return domain.LatLngToLocation(latLng), nil
}
