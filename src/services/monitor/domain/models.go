package domain

import "time"

type Restaurant struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	LatLng string `db:"lat_lng" json:"latLng"`
}

type Courier struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	LatLng string `db:"lat_lng" json:"latLng"`
}

type Order struct {
	Id           int64      `db:"id" json:"id" json:"id"`
	RestaurantId int64      `db:"restaurant_id" json:"restaurantId"`
	CourierId    int64      `db:"courier_id" json:"courierId"`
	Status       int        `db:"status" json:"status"`
	DateCreated  *time.Time `db:"date_created" json:"dateCreated"`
}

type State struct {
	Restaurants  []Restaurant `json:"restaurants"`
	Couriers     []Courier    `json:"couriers"`
	ActiveOrders []Order      `json:"activeOrders"`
}

const OrderStatusCreated = 0
const OrderStatusAccepted = 1
const OrderStatusRejected = 2
const OrderStatusReady = 3
const OrderStatusPickedUp = 4
const OrderStatusDelivered = 5

type CourierAssigned struct {
	OrderId   int64 `json:"orderId"`
	CourierId int64 `json:"courierId"`
}

type OrderUpdated struct {
	Id           int64     `json:"id"`
	Status       int       `json:"status"`
	DeliveryTime time.Time `json:"deliveryTime"`
}

type CourierLocationUpdated struct {
	OrderId int64   `json:"orderId"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}
