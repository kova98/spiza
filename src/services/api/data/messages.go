package data

import "time"

type OrderStatusUpdated struct {
	Status       int       `json:"status"`
	DeliveryTime time.Time `json:"delivery_time"`
}

type CourierLocationUpdated struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type CourierAssigned struct {
	OrderId   int64 `json:"order_id"`
	CourierId int64 `json:"courier_id"`
}

type OrderDelivered struct {
	OrderId   int64 `json:"order_id"`
	CourierId int64 `json:"courier_id"`
}
