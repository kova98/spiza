package adapters

import (
	"github.com/jmoiron/sqlx"
	"github.com/kova98/spiza/services/monitor/domain"
	_ "github.com/lib/pq"
	"log"
)

type PostgresDb struct {
	Db *sqlx.DB
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

func (p PostgresDb) GetCurrentState() (domain.State, error) {

	var state domain.State
	err := p.Db.Select(&state.Restaurants, `
		SELECT r.id, r.name, a.lat_lng 
		FROM restaurants r 
		JOIN public.addresses a on r.address_id = a.id`)
	if err != nil {
		return state, err
	}
	err = p.Db.Select(&state.Couriers, "SELECT id, name FROM couriers")
	if err != nil {
		return state, err
	}
	err = p.Db.Select(&state.ActiveOrders, `SELECT id, restaurant_id, COALESCE(courier_id,0) AS courier_id, status, date_created 
								   FROM orders 
								   WHERE status NOT IN ($1, $2)`, domain.OrderStatusDelivered, domain.OrderStatusRejected)
	if err != nil {
		return state, err
	}

	return state, nil
}
