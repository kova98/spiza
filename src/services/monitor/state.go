package main

import "time"

type Restaurant struct {
	Id   int64
	Name string
}

type Courier struct {
	Id   int64
	Name string
}

type Order struct {
	Id           int64      `db:"id"`
	RestaurantId int64      `db:"restaurant_id"`
	CourierId    int64      `db:"courier_id"`
	Status       int64      `db:"status"`
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
