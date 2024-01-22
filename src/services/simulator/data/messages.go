package data

import "time"

type CourierAssigned struct {
	OrderId   int64 `json:"order_id"`
	CourierId int64 `json:"courier_id"`
}

type OrderUpdated struct {
	Status       int       `json:"status"`
	DeliveryTime time.Time `json:"delivery_time"`
}
