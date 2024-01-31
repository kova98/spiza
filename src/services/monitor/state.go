package main

type Restaurant struct {
	Id   int64
	Name string
}

type Courier struct {
	Id   int64
	Name string
}

type Order struct {
	Id           int64
	RestaurantId int64
	CourierId    int64
	Status       int64
	DateCreated  int64
}

type State struct {
	Restaurants  []Restaurant
	Couriers     []Courier
	ActiveOrders []Order
}
