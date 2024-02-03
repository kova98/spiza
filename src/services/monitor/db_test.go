package main_test

import (
	"github.com/kova98/spiza/services/monitor"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var pg *main.PostgresDb

func TestMain(m *testing.M) {
	connStr := os.Getenv("SPIZA_TEST_DB_CONN_STR")
	if connStr == "" {
		log.Fatal("SPIZA_TEST_DB_CONN_STR environment variable empty")
	}
	pg = main.NewPostgresDb(connStr)

	clearTables()
	pg.Db.Exec("INSERT INTO users (id, name) VALUES (1, 'Test User')")
	m.Run()
	clearTables()
}

func clearTables() {
	pg.Db.Exec("DELETE FROM orders")
	pg.Db.Exec("DELETE FROM restaurants")
	pg.Db.Exec("DELETE FROM couriers")
	pg.Db.Exec("DELETE FROM addresses")
}

func TestGetCurrentState(t *testing.T) {
	addRestaurant(t, 1, "Test Restaurant")
	addCourier(t, 1, "Test Courier")
	addAddress(t, 1, "0.0,0.0", "Test Address")
	addOrder(t, 1, 1, main.OrderStatusCreated, "2024-01-01 00:00:00")
	addOrder(t, 1, 1, main.OrderStatusAccepted, "2024-01-01 00:00:00")
	addOrder(t, 1, 1, main.OrderStatusRejected, "2024-01-01 00:00:00")
	addOrder(t, 1, 1, main.OrderStatusReady, "2024-01-01 00:00:00")
	addOrder(t, 1, 1, main.OrderStatusPickedUp, "2024-01-01 00:00:00")
	addOrder(t, 1, 1, main.OrderStatusDelivered, "2024-01-01 00:00:00")

	state, err := pg.GetCurrentState()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(state.Restaurants))
	assert.Equal(t, 1, len(state.Couriers))
	assert.Equal(t, 4, len(state.ActiveOrders))
	statuses := make([]int64, 6)
	for _, order := range state.ActiveOrders {
		statuses = append(statuses, order.Status)
	}
	assert.NotContains(t, statuses, main.OrderStatusDelivered)
	assert.NotContains(t, statuses, main.OrderStatusRejected)
}

func addOrder(t *testing.T, restaurantId int, courierId int, status int, dateCreated string) {
	_, err := pg.Db.Exec(`INSERT INTO orders (restaurant_id, courier_id, status, date_created, user_id, items) 
	     				  VALUES ($1, $2, $3, $4, $5, $6)`,
		restaurantId, courierId, status, dateCreated, 1, "{}")
	if err != nil {
		t.Fatal(err)
	}
}

func addRestaurant(t *testing.T, id int, name string) {
	_, err := pg.Db.Exec("INSERT INTO restaurants (id, name) VALUES ($1, $2)", id, name)
	if err != nil {
		t.Fatal(err)
	}
}

func addCourier(t *testing.T, id int, name string) {
	_, err := pg.Db.Exec("INSERT INTO couriers (id, name) VALUES ($1, $2)", id, name)
	if err != nil {
		t.Fatal(err)
	}
}

func addAddress(t *testing.T, id int, latLng string, fullAddress string) {
	_, err := pg.Db.Exec("INSERT INTO addresses (id, lat_lng, full_address) VALUES ($1, $2, $3)", id, latLng, fullAddress)
	if err != nil {
		t.Fatal(err)
	}
}
