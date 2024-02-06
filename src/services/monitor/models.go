package main

import "time"

type Restaurant struct {
	Id     int64
	Name   string
	LatLng string `db:"lat_lng"`
}

type Courier struct {
	Id   int64
	Name string
}

type Order struct {
	Id           int64      `db:"id"`
	RestaurantId int64      `db:"restaurant_id"`
	CourierId    int64      `db:"courier_id"`
	Status       int        `db:"status"`
	DateCreated  *time.Time `db:"date_created"`
}

type State struct {
	Restaurants  []Restaurant
	Couriers     []Courier
	ActiveOrders []Order
}

const OrderStatusCreated = 0
const OrderStatusAccepted = 1
const OrderStatusRejected = 2
const OrderStatusReady = 3
const OrderStatusPickedUp = 4
const OrderStatusDelivered = 5

type CourierAssigned struct {
	OrderId   int64 `json:"order_id"`
	CourierId int64 `json:"courier_id"`
}

type OrderUpdated struct {
	Id           int64     `json:"id"`
	Status       int       `json:"status"`
	DeliveryTime time.Time `json:"delivery_time"`
}
